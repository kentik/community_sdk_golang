package synthetics

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	syntheticspb "github.com/kentik/api-schema-public/gen/go/kentik/synthetics/v202202"
	kentikerrors "github.com/kentik/community_sdk_golang/kentikapi/internal/errors"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/kentik/community_sdk_golang/kentikapi/synthetics"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// API aggregates synthetics API methods.
type API struct {
	adminService syntheticspb.SyntheticsAdminServiceClient
	dataService  syntheticspb.SyntheticsDataServiceClient
}

// NewAPI creates new synthetics API.
func NewAPI(cc grpc.ClientConnInterface) *API {
	return &API{
		adminService: syntheticspb.NewSyntheticsAdminServiceClient(cc),
		dataService:  syntheticspb.NewSyntheticsDataServiceClient(cc),
	}
}

// GetAllAgents lists synthetics agents.
func (a *API) GetAllAgents(ctx context.Context) (*synthetics.GetAllAgentsResponse, error) {
	respPayload, err := a.adminService.ListAgents(ctx, &syntheticspb.ListAgentsRequest{})
	if err != nil {
		return nil, kentikerrors.StatusErrorFromGRPC(err)
	}

	resp, err := (*listAgentsResponse)(respPayload).ToModel()
	if err != nil {
		return nil, kentikerrors.New(kentikerrors.InvalidResponse, err.Error())
	}

	return resp, err
}

// GetAgent retrieves synthetics agent with given ID.
func (a *API) GetAgent(ctx context.Context, id models.ID) (*synthetics.Agent, error) {
	respPayload, err := a.adminService.GetAgent(ctx, &syntheticspb.GetAgentRequest{Id: id})
	if err != nil {
		return nil, kentikerrors.StatusErrorFromGRPC(err)
	}

	resp, err := agentFromPayload(respPayload.GetAgent())
	if err != nil {
		return nil, kentikerrors.New(kentikerrors.InvalidResponse, err.Error())
	}

	return resp, err
}

// UpdateAgent updates the synthetics agent.
func (a *API) UpdateAgent(ctx context.Context, agent *synthetics.Agent) (*synthetics.Agent, error) {
	reqPayload, err := agentToPayload(agent)
	if err != nil {
		return nil, kentikerrors.New(kentikerrors.InvalidRequest, err.Error())
	}

	respPayload, err := a.adminService.UpdateAgent(ctx, &syntheticspb.UpdateAgentRequest{Agent: reqPayload})
	if err != nil {
		return nil, kentikerrors.StatusErrorFromGRPC(err)
	}

	resp, err := agentFromPayload(respPayload.GetAgent())
	if err != nil {
		return nil, kentikerrors.New(kentikerrors.InvalidResponse, err.Error())
	}

	return resp, nil
}

// DeleteAgent removes synthetics agent with given ID.
func (a *API) DeleteAgent(ctx context.Context, id models.ID) error {
	_, err := a.adminService.DeleteAgent(ctx, &syntheticspb.DeleteAgentRequest{Id: id})
	return kentikerrors.StatusErrorFromGRPC(err)
}

// ActivateAgent activates pending (waiting) synthetics agent with given ID.
func (a *API) ActivateAgent(ctx context.Context, id models.ID) (*synthetics.Agent, error) {
	agent, err := a.GetAgent(ctx, id)
	if err != nil {
		return nil, err
	}

	if agent.Status != synthetics.AgentStatusWait {
		log.Printf("Agent %q is not pending (status: %v) - skipping activation", agent.ID, agent.Status)
		return agent, nil
	}

	agent.Status = synthetics.AgentStatusOK
	agent, err = a.UpdateAgent(ctx, agent)
	if err != nil {
		return nil, err
	}

	if agent.Status != synthetics.AgentStatusOK {
		return nil, fmt.Errorf("failed to activate the agent %q (status: %v)", id, agent.Status)
	}

	return agent, nil
}

// DeactivateAgent deactivates the active client with given ID.
func (a *API) DeactivateAgent(ctx context.Context, id models.ID) (*synthetics.Agent, error) {
	agent, err := a.GetAgent(ctx, id)
	if err != nil {
		return nil, err
	}

	if agent.Status != synthetics.AgentStatusOK {
		log.Printf("Agent %q is not active (status: %v) - skipping deactivation", agent.ID, agent.Status)
		return agent, nil
	}

	agent.Status = synthetics.AgentStatusWait
	agent, err = a.UpdateAgent(ctx, agent)
	if err != nil {
		return nil, err
	}

	if agent.Status != synthetics.AgentStatusWait {
		return nil, fmt.Errorf("failed to deactivate the agent %q (status: %v)", id, agent.Status)
	}

	return agent, nil
}

// GetAllTests lists synthetics tests.
func (a *API) GetAllTests(ctx context.Context) (*synthetics.GetAllTestsResponse, error) {
	respPayload, err := a.adminService.ListTests(ctx, &syntheticspb.ListTestsRequest{})
	if err != nil {
		return nil, kentikerrors.StatusErrorFromGRPC(err)
	}

	resp, err := (*listTestsResponse)(respPayload).ToModel()
	if err != nil {
		return nil, kentikerrors.New(kentikerrors.InvalidResponse, err.Error())
	}

	return resp, nil
}

// GetTest retrieves synthetics test with given ID.
func (a *API) GetTest(ctx context.Context, id models.ID) (*synthetics.Test, error) {
	respPayload, err := a.adminService.GetTest(ctx, &syntheticspb.GetTestRequest{Id: id})
	if err != nil {
		return nil, kentikerrors.StatusErrorFromGRPC(err)
	}

	resp, err := testFromPayload(respPayload.GetTest())
	if err != nil {
		return nil, kentikerrors.New(kentikerrors.InvalidResponse, err.Error())
	}

	return resp, nil
}

// CreateTest creates the synthetics test.
func (a *API) CreateTest(ctx context.Context, test *synthetics.Test) (*synthetics.Test, error) {
	test, err := testWithDefaultFields(test)
	if err != nil {
		return nil, kentikerrors.New(kentikerrors.InvalidRequest, err.Error())
	}

	// TODO(dfurman): validate create request

	reqPayload, err := testToPayload(test)
	if err != nil {
		return nil, kentikerrors.New(kentikerrors.InvalidRequest, err.Error())
	}

	respPayload, err := a.adminService.CreateTest(ctx, &syntheticspb.CreateTestRequest{Test: reqPayload})
	if err != nil {
		return nil, kentikerrors.StatusErrorFromGRPC(err)
	}

	resp, err := testFromPayload(respPayload.GetTest())
	if err != nil {
		return nil, kentikerrors.New(kentikerrors.InvalidResponse, err.Error())
	}

	return resp, nil
}

func testWithDefaultFields(t *synthetics.Test) (*synthetics.Test, error) {
	if t == nil {
		return nil, errors.New("test object is nil")
	}

	if t.Status == "" {
		t.Status = synthetics.TestStatusActive // field required by the server (but ignored on create)
	}

	if t.Settings.Period == 0 {
		t.Settings.Period = time.Minute // field required by the server
	}

	if t.Settings.Family == "" {
		t.Settings.Family = synthetics.IPFamilyDual // field required by the server
	}

	if t.Settings.Health.UnhealthySubtestThreshold == 0 {
		t.Settings.Health.UnhealthySubtestThreshold = 1 // field required by the server
	}

	if t.Settings.Traceroute != nil && t.Settings.Traceroute.Protocol != synthetics.TracerouteProtocolICMP {
		t.Settings.Traceroute.Port = 33434
	}

	return t, nil
}

// UpdateTest updates the synthetics test.
func (a *API) UpdateTest(ctx context.Context, test *synthetics.Test) (*synthetics.Test, error) {
	// TODO(dfurman): validate update request

	reqPayload, err := testToPayload(test)
	if err != nil {
		return nil, kentikerrors.New(kentikerrors.InvalidRequest, err.Error())
	}

	respPayload, err := a.adminService.UpdateTest(ctx, &syntheticspb.UpdateTestRequest{Test: reqPayload})
	if err != nil {
		return nil, kentikerrors.StatusErrorFromGRPC(err)
	}

	resp, err := testFromPayload(respPayload.GetTest())
	if err != nil {
		return nil, kentikerrors.New(kentikerrors.InvalidResponse, err.Error())
	}

	return resp, nil
}

// DeleteTest removes synthetics test with given ID.
func (a *API) DeleteTest(ctx context.Context, id models.ID) error {
	_, err := a.adminService.DeleteTest(ctx, &syntheticspb.DeleteTestRequest{Id: id})
	return kentikerrors.StatusErrorFromGRPC(err)
}

// SetTestStatus modifies status of the synthetics test with given ID.
func (a *API) SetTestStatus(ctx context.Context, id models.ID, ts synthetics.TestStatus) error {
	_, err := a.adminService.SetTestStatus(ctx, &syntheticspb.SetTestStatusRequest{
		Id:     id,
		Status: syntheticspb.TestStatus(syntheticspb.TestStatus_value[string(ts)]),
	})
	return kentikerrors.StatusErrorFromGRPC(err)
}

// GetResultsForTests returns measurement results for a set of tests for specified period of time, or the latest
// available data. It returns one TestResults object for each requested test.
// TODO(dfurman): implement integration test for the endpoint.
func (a *API) GetResultsForTests(
	ctx context.Context, req synthetics.GetResultsForTestsRequest,
) ([]synthetics.TestResults, error) {
	respPayload, err := a.dataService.GetResultsForTests(ctx, getResultsForTestsRequestToPayload(req))
	if err != nil {
		return nil, kentikerrors.StatusErrorFromGRPC(err)
	}

	resp, err := (*getResultsForTestsResponse)(respPayload).ToModel()
	if err != nil {
		return nil, kentikerrors.New(kentikerrors.InvalidResponse, err.Error())
	}

	return resp, nil
}

// GetTraceForTest retrieves trace route results for specified synthetic test.
// TODO(dfurman): implement integration test for the endpoint.
func (a *API) GetTraceForTest(
	ctx context.Context, req synthetics.GetTraceForTestRequest,
) (synthetics.GetTraceForTestResponse, error) {
	respPayload, err := a.dataService.GetTraceForTest(ctx, &syntheticspb.GetTraceForTestRequest{
		Id:        req.TestID,
		StartTime: timestamppb.New(req.StartTime),
		EndTime:   timestamppb.New(req.EndTime),
		AgentIds:  req.AgentIDs,
		TargetIps: req.TargetIPs,
	})
	if err != nil {
		return synthetics.GetTraceForTestResponse{}, kentikerrors.StatusErrorFromGRPC(err)
	}

	resp, err := (*getTraceForTestResponse)(respPayload).ToModel()
	if err != nil {
		return synthetics.GetTraceForTestResponse{}, kentikerrors.New(kentikerrors.InvalidResponse, err.Error())
	}

	return resp, nil
}
