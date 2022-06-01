package synthetics

import (
	"context"
	"fmt"
	"log"

	syntheticspb "github.com/kentik/api-schema-public/gen/go/kentik/synthetics/v202202"
	kentikerrors "github.com/kentik/community_sdk_golang/kentikapi/internal/errors"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/kentik/community_sdk_golang/kentikapi/synthetics"
	"google.golang.org/grpc"
)

// API aggregates Synthetics API methods.
type API struct {
	client syntheticspb.SyntheticsAdminServiceClient
}

// NewAPI creates new API.
func NewAPI(cc grpc.ClientConnInterface) *API {
	return &API{
		client: syntheticspb.NewSyntheticsAdminServiceClient(cc),
	}
}

// GetAllAgents lists synthetics agents.
func (a *API) GetAllAgents(ctx context.Context) (*synthetics.GetAllAgentsResponse, error) {
	response, err := a.client.ListAgents(ctx, &syntheticspb.ListAgentsRequest{})
	if err != nil {
		return nil, kentikerrors.StatusErrorFromGRPC(err)
	}

	return (*listAgentsResponse)(response).ToModel()
}

// GetAgent retrieves synthetics agent with given ID.
func (a *API) GetAgent(ctx context.Context, id models.ID) (*synthetics.Agent, error) {
	response, err := a.client.GetAgent(ctx, &syntheticspb.GetAgentRequest{Id: id})
	if err != nil {
		return nil, kentikerrors.StatusErrorFromGRPC(err)
	}

	return agentFromPayload(response.GetAgent())
}

// UpdateAgent updates the synthetics agent.
func (a *API) UpdateAgent(ctx context.Context, agent *synthetics.Agent) (*synthetics.Agent, error) {
	payload, err := agentToPayload(agent)
	if err != nil {
		return nil, kentikerrors.StatusErrorFromGRPC(err)
	}

	response, err := a.client.UpdateAgent(ctx, &syntheticspb.UpdateAgentRequest{
		Agent: payload,
	})
	if err != nil {
		return nil, kentikerrors.StatusErrorFromGRPC(err)
	}

	return agentFromPayload(response.GetAgent())
}

// DeleteAgent removes synthetics agent with given ID.
func (a *API) DeleteAgent(ctx context.Context, id models.ID) error {
	_, err := a.client.DeleteAgent(ctx, &syntheticspb.DeleteAgentRequest{Id: id})
	return kentikerrors.StatusErrorFromGRPC(err)
}

// ActivateAgent activates pending (waiting) synthetics agent with given ID.
func (a *API) ActivateAgent(ctx context.Context, id models.ID) (*synthetics.Agent, error) {
	agent, err := a.GetAgent(ctx, id)
	if err != nil {
		return nil, kentikerrors.StatusErrorFromGRPC(err)
	}

	if agent.Status != synthetics.AgentStatusWait {
		log.Printf("Agent %q is not pending (status: %v) - skipping activation", agent.ID, agent.Status)
		return agent, nil
	}

	agent.Status = synthetics.AgentStatusOK
	agent, err = a.UpdateAgent(ctx, agent)
	if err != nil {
		return nil, kentikerrors.StatusErrorFromGRPC(err)
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
		return nil, kentikerrors.StatusErrorFromGRPC(err)
	}

	if agent.Status != synthetics.AgentStatusOK {
		log.Printf("Agent %q is not active (status: %v) - skipping deactivation", agent.ID, agent.Status)
		return agent, nil
	}

	agent.Status = synthetics.AgentStatusWait
	agent, err = a.UpdateAgent(ctx, agent)
	if err != nil {
		return nil, kentikerrors.StatusErrorFromGRPC(err)
	}

	if agent.Status != synthetics.AgentStatusWait {
		return nil, fmt.Errorf("failed to deactivate the agent %q (status: %v)", id, agent.Status)
	}

	return agent, nil
}
