package kentikapi_test

import (
	"context"
	"net"
	"testing"

	syntheticspb "github.com/kentik/api-schema-public/gen/go/kentik/synthetics/v202202"
	"github.com/kentik/community_sdk_golang/kentikapi"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestClient_GetAgent(t *testing.T) {
	tests := []struct {
		name string
		// TODO(dfurman): create public Agent model
		request           *syntheticspb.GetAgentRequest
		expectedRequest   *syntheticspb.GetAgentRequest
		response          getAgentResponse
		expectedResult    *syntheticspb.GetAgentResponse
		expectedErrorCode *codes.Code
		expectedError     bool
	}{
		{
			name:            "empty request, status InvalidArgument received",
			request:         &syntheticspb.GetAgentRequest{},
			expectedRequest: &syntheticspb.GetAgentRequest{Id: ""},
			response: getAgentResponse{
				err: status.Errorf(codes.InvalidArgument, codes.InvalidArgument.String()),
			},
			expectedResult:    nil,
			expectedErrorCode: codePtr(codes.InvalidArgument),
			expectedError:     true,
		}, {
			name:            "status NotFound received",
			request:         &syntheticspb.GetAgentRequest{Id: "1000"},
			expectedRequest: &syntheticspb.GetAgentRequest{Id: "1000"},
			response: getAgentResponse{
				err: status.Errorf(codes.NotFound, codes.NotFound.String()),
			},
			expectedResult:    nil,
			expectedErrorCode: codePtr(codes.NotFound),
			expectedError:     true,
		}, {
			name:            "agent returned",
			request:         &syntheticspb.GetAgentRequest{Id: testAgentID},
			expectedRequest: &syntheticspb.GetAgentRequest{Id: testAgentID},
			response: getAgentResponse{
				data: &syntheticspb.GetAgentResponse{Agent: newDummyAgent()},
			},
			expectedResult: &syntheticspb.GetAgentResponse{Agent: newDummyAgent()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			server := newSpySyntheticsServer(t, syntheticsResponses{
				getAgentResponse: tt.response,
			})
			server.Start()
			defer server.Stop()

			client, err := kentikapi.NewClient(kentikapi.Config{
				APIURL:      "http://" + server.url,
				AuthToken:   dummyAuthToken,
				AuthEmail:   dummyAuthEmail,
				LogPayloads: true,
			})
			require.NoError(t, err)

			// act
			result, err := client.SyntheticsAdmin.GetAgent(context.Background(), tt.request)

			// assert
			t.Logf("Got result: %+v, err: %v", result, err)
			if tt.expectedError {
				assert.Error(t, err)
				if tt.expectedErrorCode != nil {
					s, ok := status.FromError(err)
					assert.True(t, ok)
					assert.Equal(t, *tt.expectedErrorCode, s.Code())
				}
			} else {
				assert.NoError(t, err)
			}

			if assert.Equal(t, 1, len(server.requests.getAgentRequests), "invalid number of requests") {
				r := server.requests.getAgentRequests[0]
				assert.Equal(t, dummyAuthEmail, r.metadata.Get(authEmailKey)[0])
				assert.Equal(t, dummyAuthToken, r.metadata.Get(authAPITokenKey)[0])
				testutil.AssertProtoEqual(t, tt.expectedRequest, r.data)
			}

			testutil.AssertProtoEqual(t, tt.expectedResult, result)
		})
	}
}

type spySyntheticsServer struct {
	syntheticspb.UnimplementedSyntheticsAdminServiceServer
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
	getAgentRequests []getAgentRequest
}

type getAgentRequest struct {
	metadata metadata.MD
	data     *syntheticspb.GetAgentRequest
}

type syntheticsResponses struct {
	getAgentResponse getAgentResponse
}

type getAgentResponse struct {
	data *syntheticspb.GetAgentResponse
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
	syntheticspb.RegisterSyntheticsAdminServiceServer(s.server, s)

	go func() {
		err = s.server.Serve(l)
		assert.NoError(s.t, err)
		s.done <- struct{}{}
	}()
}

// Stop blocks until the server is stopped.
func (s *spySyntheticsServer) Stop() {
	s.server.GracefulStop()
	<-s.done
}

func (s *spySyntheticsServer) GetAgent(
	ctx context.Context, req *syntheticspb.GetAgentRequest,
) (*syntheticspb.GetAgentResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	s.requests.getAgentRequests = append(s.requests.getAgentRequests, getAgentRequest{
		metadata: md,
		data:     req,
	})

	return s.responses.getAgentResponse.data, s.responses.getAgentResponse.err
}
