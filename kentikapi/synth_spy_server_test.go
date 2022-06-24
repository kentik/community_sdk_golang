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
	listAgentsRequests  []listAgentsRequest
	getAgentRequests    []getAgentRequest
	updateAgentRequests []updateAgentRequest
	deleteAgentRequests []deleteAgentRequest
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

type syntheticsResponses struct {
	listAgentsResponse  listAgentsResponse
	getAgentResponse    getAgentResponse
	updateAgentResponse updateAgentResponse
	deleteAgentResponse deleteAgentResponse
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
