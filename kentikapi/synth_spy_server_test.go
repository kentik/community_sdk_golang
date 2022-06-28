package kentikapi_test

import (
	"context"
	"errors"
	"net"
	"testing"

	"github.com/kentik/api-schema-public/gen/go/kentik/synthetics/v202202"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type spySyntheticsServer struct {
	synthetics.UnimplementedSyntheticsAdminServiceServer
	server *grpc.Server

	url  string
	done chan struct{}
	t    testing.TB
	// responses to return to the client
	responses syntheticsResponses

	// requests spied by the server
	requests syntheticsRequests
}

type syntheticsRequests struct {
	listAgentsRequests    []listAgentsRequest
	getAgentRequests      []getAgentRequest
	updateAgentRequests   []updateAgentRequest
	deleteAgentRequests   []deleteAgentRequest
	listTestsRequests     []listTestsRequest
	getTestRequests       []getTestRequest
	createTestRequests    []createTestRequest
	updateTestRequests    []updateTestRequest
	deleteTestRequests    []deleteTestRequest
	setTestStatusRequests []setTestStatusRequest
}

type listAgentsRequest struct {
	metadata metadata.MD
	data     *synthetics.ListAgentsRequest
}

type getAgentRequest struct {
	metadata metadata.MD
	data     *synthetics.GetAgentRequest
}

type updateAgentRequest struct {
	metadata metadata.MD
	data     *synthetics.UpdateAgentRequest
}

type deleteAgentRequest struct {
	metadata metadata.MD
	data     *synthetics.DeleteAgentRequest
}

type listTestsRequest struct {
	metadata metadata.MD
	data     *synthetics.ListTestsRequest
}

type getTestRequest struct {
	metadata metadata.MD
	data     *synthetics.GetTestRequest
}

type createTestRequest struct {
	metadata metadata.MD
	data     *synthetics.CreateTestRequest
}

type updateTestRequest struct {
	metadata metadata.MD
	data     *synthetics.UpdateTestRequest
}

type deleteTestRequest struct {
	metadata metadata.MD
	data     *synthetics.DeleteTestRequest
}

type setTestStatusRequest struct {
	metadata metadata.MD
	data     *synthetics.SetTestStatusRequest
}

type syntheticsResponses struct {
	listAgentsResponse    listAgentsResponse
	getAgentResponse      getAgentResponse
	updateAgentResponse   updateAgentResponse
	deleteAgentResponse   deleteAgentResponse
	listTestsResponse     listTestsResponse
	getTestResponse       getTestResponse
	createTestResponse    createTestResponse
	updateTestResponse    updateTestResponse
	deleteTestResponse    deleteTestResponse
	setTestStatusResponse setTestStatusResponse
}

type listAgentsResponse struct {
	data *synthetics.ListAgentsResponse
	err  error
}

type getAgentResponse struct {
	data *synthetics.GetAgentResponse
	err  error
}

type updateAgentResponse struct {
	data *synthetics.UpdateAgentResponse
	err  error
}

type deleteAgentResponse struct {
	data *synthetics.DeleteAgentResponse
	err  error
}

type listTestsResponse struct {
	data *synthetics.ListTestsResponse
	err  error
}

type getTestResponse struct {
	data *synthetics.GetTestResponse
	err  error
}

type createTestResponse struct {
	data *synthetics.CreateTestResponse
	err  error
}

type updateTestResponse struct {
	data *synthetics.UpdateTestResponse
	err  error
}

type deleteTestResponse struct {
	data *synthetics.DeleteTestResponse
	err  error
}

type setTestStatusResponse struct {
	data *synthetics.SetTestStatusResponse
	err  error
}

func newSpySyntheticsServer(t testing.TB, r syntheticsResponses) *spySyntheticsServer {
	return &spySyntheticsServer{
		done:      make(chan struct{}),
		t:         t,
		responses: r,
	}
}

func (s *spySyntheticsServer) Start() {
	l, err := net.Listen("tcp", "localhost:0")
	require.NoError(s.t, err)

	s.url = l.Addr().String()
	s.server = grpc.NewServer()
	synthetics.RegisterSyntheticsAdminServiceServer(s.server, s)

	go func() {
		err = s.server.Serve(l)
		if !errors.Is(err, grpc.ErrServerStopped) {
			assert.NoError(s.t, err)
		}
		s.done <- struct{}{}
	}()
}

// Stop blocks until the server is stopped.
func (s *spySyntheticsServer) Stop() {
	s.server.GracefulStop()
	<-s.done
}

func (s *spySyntheticsServer) ListAgents(
	ctx context.Context, req *synthetics.ListAgentsRequest,
) (*synthetics.ListAgentsResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	s.requests.listAgentsRequests = append(s.requests.listAgentsRequests, listAgentsRequest{
		metadata: md,
		data:     req,
	})

	return s.responses.listAgentsResponse.data, s.responses.listAgentsResponse.err
}

func (s *spySyntheticsServer) GetAgent(
	ctx context.Context, req *synthetics.GetAgentRequest,
) (*synthetics.GetAgentResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	s.requests.getAgentRequests = append(s.requests.getAgentRequests, getAgentRequest{
		metadata: md,
		data:     req,
	})

	return s.responses.getAgentResponse.data, s.responses.getAgentResponse.err
}

func (s *spySyntheticsServer) UpdateAgent(
	ctx context.Context, req *synthetics.UpdateAgentRequest,
) (*synthetics.UpdateAgentResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	s.requests.updateAgentRequests = append(s.requests.updateAgentRequests, updateAgentRequest{
		metadata: md,
		data:     req,
	})

	return s.responses.updateAgentResponse.data, s.responses.updateAgentResponse.err
}

func (s *spySyntheticsServer) DeleteAgent(
	ctx context.Context, req *synthetics.DeleteAgentRequest,
) (*synthetics.DeleteAgentResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	s.requests.deleteAgentRequests = append(s.requests.deleteAgentRequests, deleteAgentRequest{
		metadata: md,
		data:     req,
	})

	return s.responses.deleteAgentResponse.data, s.responses.deleteAgentResponse.err
}

func (s *spySyntheticsServer) ListTests(
	ctx context.Context, req *synthetics.ListTestsRequest,
) (*synthetics.ListTestsResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	s.requests.listTestsRequests = append(s.requests.listTestsRequests, listTestsRequest{
		metadata: md,
		data:     req,
	})

	return s.responses.listTestsResponse.data, s.responses.listTestsResponse.err
}

func (s *spySyntheticsServer) GetTest(
	ctx context.Context, req *synthetics.GetTestRequest,
) (*synthetics.GetTestResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	s.requests.getTestRequests = append(s.requests.getTestRequests, getTestRequest{
		metadata: md,
		data:     req,
	})

	return s.responses.getTestResponse.data, s.responses.getTestResponse.err
}

func (s *spySyntheticsServer) CreateTest(
	ctx context.Context, req *synthetics.CreateTestRequest,
) (*synthetics.CreateTestResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	s.requests.createTestRequests = append(s.requests.createTestRequests, createTestRequest{
		metadata: md,
		data:     req,
	})

	return s.responses.createTestResponse.data, s.responses.createTestResponse.err
}

func (s *spySyntheticsServer) UpdateTest(
	ctx context.Context, req *synthetics.UpdateTestRequest,
) (*synthetics.UpdateTestResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	s.requests.updateTestRequests = append(s.requests.updateTestRequests, updateTestRequest{
		metadata: md,
		data:     req,
	})

	return s.responses.updateTestResponse.data, s.responses.updateTestResponse.err
}

func (s *spySyntheticsServer) DeleteTest(
	ctx context.Context, req *synthetics.DeleteTestRequest,
) (*synthetics.DeleteTestResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	s.requests.deleteTestRequests = append(s.requests.deleteTestRequests, deleteTestRequest{
		metadata: md,
		data:     req,
	})

	return s.responses.deleteTestResponse.data, s.responses.deleteTestResponse.err
}

func (s *spySyntheticsServer) SetTestStatus(
	ctx context.Context, req *synthetics.SetTestStatusRequest,
) (*synthetics.SetTestStatusResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	s.requests.setTestStatusRequests = append(s.requests.setTestStatusRequests, setTestStatusRequest{
		metadata: md,
		data:     req,
	})

	return s.responses.setTestStatusResponse.data, s.responses.setTestStatusResponse.err
}
