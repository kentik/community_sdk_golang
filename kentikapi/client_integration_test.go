package kentikapi_test

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/AlekSi/pointer"
	syntheticspb "github.com/kentik/api-schema-public/gen/go/kentik/synthetics/v202101beta1"
	"github.com/kentik/community_sdk_golang/kentikapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	testAgentID = "968"
)

func TestClient_GetAgent(t *testing.T) {
	tests := []struct {
		name              string
		retryMax          *uint
		timeout           *time.Duration
		request           *syntheticspb.GetAgentRequest
		expectedRequestID string
		responses         []gRPCResponse
		handlingDelay     time.Duration
		expectedResult    *syntheticspb.GetAgentResponse
		expectedErrorCode codes.Code
		expectedErrorMsg  string
		expectedError     bool
	}{
		{
			name:              "empty request, status InvalidArgument received",
			request:           &syntheticspb.GetAgentRequest{},
			expectedRequestID: "",
			responses: []gRPCResponse{
				newErrorGRPCResponse(codes.InvalidArgument),
			},
			expectedResult:    nil,
			expectedErrorCode: codes.InvalidArgument,
			expectedErrorMsg:  codes.InvalidArgument.String(),
			expectedError:     true,
		}, {
			name:              "request with invalid id, status NotFound received",
			request:           &syntheticspb.GetAgentRequest{Id: "1000"},
			expectedRequestID: "1000",
			responses: []gRPCResponse{
				newErrorGRPCResponse(codes.NotFound),
			},
			expectedResult:    nil,
			expectedErrorCode: codes.NotFound,
			expectedErrorMsg:  codes.NotFound.String(),
			expectedError:     true,
		}, {
			name:              "valid request, agent returned",
			request:           &syntheticspb.GetAgentRequest{Id: "968"},
			expectedRequestID: "968",
			responses: []gRPCResponse{
				{
					nil,
					&syntheticspb.GetAgentResponse{Agent: newDummyAgentForGRPC()},
				},
			},
			expectedResult: &syntheticspb.GetAgentResponse{Agent: newDummyAgentForGRPC()},
		}, {
			name:              "retry 2 times till success on code Unavailable",
			request:           &syntheticspb.GetAgentRequest{Id: "968"},
			expectedRequestID: "968",
			responses: []gRPCResponse{
				newErrorGRPCResponse(codes.Unavailable),
				newErrorGRPCResponse(codes.Unavailable),
				{
					nil,
					&syntheticspb.GetAgentResponse{Agent: newDummyAgentForGRPC()},
				},
			},
			expectedResult: &syntheticspb.GetAgentResponse{Agent: newDummyAgentForGRPC()},
			expectedError:  false,
		}, {
			name:    "retry 4 times when code Unavailable received and last code is Unavailable",
			request: &syntheticspb.GetAgentRequest{},
			responses: []gRPCResponse{
				newErrorGRPCResponse(codes.Unavailable),
				newErrorGRPCResponse(codes.Unavailable),
				newErrorGRPCResponse(codes.Unavailable),
				newErrorGRPCResponse(codes.Unavailable),
				newErrorGRPCResponse(codes.Unavailable),
			},
			expectedResult:    nil,
			expectedErrorCode: codes.Unavailable,
			expectedErrorMsg:  codes.Unavailable.String(),
			expectedError:     true,
		}, {
			name:    "retry 4 times when code Unavailable received and last code is Unknown",
			request: &syntheticspb.GetAgentRequest{},
			responses: []gRPCResponse{
				newErrorGRPCResponse(codes.Unavailable),
				newErrorGRPCResponse(codes.Unavailable),
				newErrorGRPCResponse(codes.Unavailable),
				newErrorGRPCResponse(codes.Unavailable),
				newErrorGRPCResponse(codes.Unknown),
			},
			expectedResult:    nil,
			expectedErrorCode: codes.Unknown,
			expectedErrorMsg:  codes.Unknown.String(),
			expectedError:     true,
		}, {
			name:    "do not retry when code Unknown received",
			request: &syntheticspb.GetAgentRequest{},
			responses: []gRPCResponse{
				newErrorGRPCResponse(codes.Unknown),
			},
			expectedResult:    nil,
			expectedErrorCode: codes.Unknown,
			expectedErrorMsg:  codes.Unknown.String(),
			expectedError:     true,
		}, {
			name:     "do not retry when retries disabled and code Unavailable received",
			retryMax: pointer.ToUint(0),
			request:  &syntheticspb.GetAgentRequest{},
			responses: []gRPCResponse{
				newErrorGRPCResponse(codes.Unavailable),
			},
			expectedResult:    nil,
			expectedErrorCode: codes.Unavailable,
			expectedErrorMsg:  codes.Unavailable.String(),
			expectedError:     true,
		}, {
			name:     "retry specified number of times when code Unavailable received",
			retryMax: pointer.ToUint(2),
			request:  &syntheticspb.GetAgentRequest{},
			responses: []gRPCResponse{
				newErrorGRPCResponse(codes.Unavailable),
				newErrorGRPCResponse(codes.Unavailable),
				newErrorGRPCResponse(codes.Unavailable),
			},
			expectedResult:    nil,
			expectedErrorCode: codes.Unavailable,
			expectedErrorMsg:  codes.Unavailable.String(),
			expectedError:     true,
		}, {
			name:              "timeout during first request",
			timeout:           pointer.ToDuration(1 * time.Millisecond),
			request:           &syntheticspb.GetAgentRequest{},
			responses:         []gRPCResponse{},
			handlingDelay:     1 * time.Second,
			expectedResult:    nil,
			expectedErrorCode: codes.DeadlineExceeded,
			expectedErrorMsg:  "context deadline exceeded",
			expectedError:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			s := newSpySyntheticsServer(t, tt.responses)
			s.handlingDelay = tt.handlingDelay
			s.Start()
			defer s.Stop()

			client, err := kentikapi.NewClient(kentikapi.Config{
				SyntheticsHostPort: s.url,
				AuthToken:          dummyAuthToken,
				AuthEmail:          dummyAuthEmail,
				DisableTLS:         true,
				RetryCfg: kentikapi.RetryConfig{
					MaxAttempts: tt.retryMax,
					MinDelay:    pointer.ToDuration(10 * time.Microsecond),
				},
				Timeout: tt.timeout,
			})
			require.NoError(t, err)

			// act
			result, err := client.SyntheticsAdmin.GetAgent(
				context.Background(),
				tt.request,
			)

			// assert
			t.Logf("Got response: %v, err: %v", result, err)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if e, ok := status.FromError(err); ok {
				assert.Equal(t, tt.expectedErrorCode, e.Code())
				assert.Equal(t, tt.expectedErrorMsg, e.Message())
			}

			assert.Equal(t, len(s.responses), len(s.requests), "invalid number of requests")
			for i, r := range s.requests {
				assert.Equal(t, dummyAuthEmail, s.headers[i].Get(authEmailKey)[0])
				assert.Equal(t, dummyAuthToken, s.headers[i].Get(authAPITokenKey)[0])
				assert.Equal(t, tt.expectedRequestID, r.GetId())
			}

			assert.True(
				t,
				proto.Equal(tt.expectedResult, result),
				fmt.Sprintf("Protobuf messages are not equal:\nexpected: %v\nactual:  %v", tt.expectedResult, result),
			)
		})
	}
}

type spySyntheticsServer struct {
	syntheticspb.UnimplementedSyntheticsAdminServiceServer
	server *grpc.Server
	url    string
	done   chan struct{}
	t      testing.TB
	// responses to return to the client
	responses []gRPCResponse
	// handlingDelay specifies the delay applied while handling the request
	handlingDelay time.Duration

	// requests spied by the server
	requests []*syntheticspb.GetAgentRequest
	headers  []metadata.MD
}

type gRPCResponse struct {
	err  error
	body *syntheticspb.GetAgentResponse
}

func newSpySyntheticsServer(t testing.TB, responses []gRPCResponse) *spySyntheticsServer {
	return &spySyntheticsServer{
		done:      make(chan struct{}),
		t:         t,
		responses: responses,
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

// Stop blocks until the server is stopped. Graceful stop is not used to make tests quicker.
func (s *spySyntheticsServer) Stop() {
	s.server.Stop()
	<-s.done
}

func (s *spySyntheticsServer) GetAgent(ctx context.Context, req *syntheticspb.GetAgentRequest,
) (*syntheticspb.GetAgentResponse, error) {
	time.Sleep(s.handlingDelay)

	header, ok := metadata.FromIncomingContext(ctx)
	assert.True(s.t, ok)

	s.headers = append(s.headers, header)

	s.requests = append(s.requests, req)

	response := s.response()

	return response.body, response.err
}

func (s *spySyntheticsServer) response() gRPCResponse {
	if len(s.requests) > len(s.responses) {
		return gRPCResponse{
			status.Errorf(
				codes.Unknown,
				"spySyntheticsServer: unexpected request, requests count: %v, expected: %v",
				len(s.requests),
				len(s.responses),
			),
			nil,
		}
	}
	return s.responses[len(s.requests)-1]
}

func newDummyAgentForGRPC() *syntheticspb.Agent {
	return &syntheticspb.Agent{
		Id:     testAgentID,
		Name:   "dummy-agent",
		Status: syntheticspb.AgentStatus_AGENT_STATUS_WAIT,
		Alias:  "probe-4-ams-1",
		Type:   "global",
		Os:     "I use Manjaro BTW",
		Ip:     "95.179.136.58",
		Lat:    52.374031,
		Long:   4.88969,
		LastAuthed: timestamppb.New(time.Date(2020,
			time.July,
			9,
			21,
			37,
			0,
			826*1000000,
			time.UTC,
		)),
		Family:    syntheticspb.IPFamily_IP_FAMILY_DUAL,
		Asn:       20473,
		SiteId:    "2137",
		Version:   "0.0.2",
		Challenge: "dummy-challenge",
		City:      "Amsterdam",
		Region:    "Noord-Holland",
		Country:   "Netherlands",
		TestIds:   []string{"13", "133", "1337"},
		LocalIp:   "10.10.10.10",
	}
}

func newErrorGRPCResponse(c codes.Code) gRPCResponse {
	return gRPCResponse{
		err: status.Errorf(
			c,
			c.String(),
		),
		body: nil,
	}
}
