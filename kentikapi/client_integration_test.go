package kentikapi_test

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/AlekSi/pointer"
	syntheticspb "github.com/kentik/api-schema-public/gen/go/kentik/synthetics/v202202"
	"github.com/kentik/community_sdk_golang/kentikapi"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/testutil"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	testAgentID = "968"
)

func TestClient_GetUserWithRetries(t *testing.T) {
	tests := []struct {
		name                string
		responses           []testutil.HTTPResponse
		serverHandlingDelay time.Duration
		timeout             time.Duration
		expectedResult      *models.User
		expectedError       bool
		expectedTimeout     bool
	}{
		{
			name: "retry on status 502 Bad Gateway until invalid response format received",
			responses: []testutil.HTTPResponse{
				testutil.NewErrorHTTPResponse(http.StatusBadGateway),
				testutil.NewErrorHTTPResponse(http.StatusBadGateway),
				{StatusCode: http.StatusOK, Body: "invalid JSON"},
			},
			expectedError: true,
		}, {
			name: "retry till success when status 429 Too Many Requests received",
			responses: []testutil.HTTPResponse{
				testutil.NewErrorHTTPResponse(http.StatusTooManyRequests),
				{
					StatusCode: http.StatusOK,
					Body: `{
							"user": {
								"id": "145999",
								"username": "testuser",
								"user_full_name": "Test User",
								"user_email": "test@user.example",
								"role": "Member",
								"email_service": true,
								"email_product": true,
								"last_login": null,
								"created_date": "2020-12-09T14:48:42.187Z",
								"updated_date": "2020-12-09T14:48:43.243Z",
								"company_id": "74333",
								"user_api_token": "****************************a997",
								"filters": {},
								"saved_filters": []
							}
						}`,
				},
			},
			expectedResult: &models.User{
				ID:           "145999",
				Username:     "testuser",
				UserFullName: "Test User",
				UserEmail:    "test@user.example",
				Role:         "Member",
				EmailService: true,
				EmailProduct: true,
				LastLogin:    nil,
				CreatedDate:  *testutil.ParseISO8601Timestamp(t, "2020-12-09T14:48:42.187Z"),
				UpdatedDate:  *testutil.ParseISO8601Timestamp(t, "2020-12-09T14:48:43.243Z"),
				CompanyID:    "74333",
				UserAPIToken: pointer.ToString("****************************a997"),
			},
		}, {
			name: "default timeout is longer than 30 ms",
			responses: []testutil.HTTPResponse{
				testutil.NewErrorHTTPResponse(http.StatusBadGateway),
				testutil.NewErrorHTTPResponse(http.StatusBadGateway),
				{StatusCode: http.StatusBadRequest, Body: `{"error":"Bad Request"}`},
			},
			serverHandlingDelay: 10 * time.Millisecond,
			expectedError:       true,
			expectedTimeout:     false,
		}, {
			name: "timeout is longer than the wait for response with retries",
			responses: []testutil.HTTPResponse{
				testutil.NewErrorHTTPResponse(http.StatusBadGateway),
				testutil.NewErrorHTTPResponse(http.StatusBadGateway),
				{StatusCode: http.StatusBadRequest, Body: `{"error":"Bad Request"}`},
			},
			serverHandlingDelay: 10 * time.Millisecond,
			timeout:             10 * time.Second,
			expectedError:       true,
			expectedTimeout:     false,
		}, {
			name: "timeout during first request",
			responses: []testutil.HTTPResponse{
				testutil.NewErrorHTTPResponse(http.StatusBadGateway),
				testutil.NewErrorHTTPResponse(http.StatusBadGateway),
				{StatusCode: http.StatusBadRequest, Body: `{"error":"Bad Request"}`},
			},
			serverHandlingDelay: 10 * time.Millisecond,
			timeout:             5 * time.Millisecond,
			expectedError:       true,
			expectedTimeout:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			h := testutil.NewMultipleResponseSpyHTTPHandler(t, tt.responses, tt.serverHandlingDelay)
			s := httptest.NewServer(h)
			defer s.Close()

			c, err := kentikapi.NewClient(
				kentikapi.WithAPIURL(s.URL),
				kentikapi.WithCredentials(dummyAuthEmail, dummyAuthToken),
				kentikapi.WithTimeout(tt.timeout),
				kentikapi.WithRetryMinDelay(1*time.Microsecond),
				kentikapi.WithRetryMaxDelay(10*time.Microsecond),
				kentikapi.WithLogPayloads(),
			)
			assert.NoError(t, err)

			// act
			result, err := c.Users.Get(context.Background(), testUserID)

			// assert
			t.Logf("Got result: %v, err: %v", result, err)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if tt.expectedTimeout {
				assert.True(t, errors.Is(err, context.DeadlineExceeded))
			} else {
				assert.False(t, errors.Is(err, context.DeadlineExceeded))
			}

			if !tt.expectedTimeout {
				assert.Equal(t, len(tt.responses), len(h.Requests), "invalid number of requests")
			}

			for _, r := range h.Requests {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, fmt.Sprintf("/api/v5/user/%v", testUserID), r.URL.Path)
				assert.Equal(t, dummyAuthEmail, r.Header.Get(authEmailKey))
				assert.Equal(t, dummyAuthToken, r.Header.Get(authAPITokenKey))
			}

			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestClient_GetAgentWithRetries(t *testing.T) {
	tests := []struct {
		name              string
		options           []kentikapi.ClientOption
		request           *syntheticspb.GetAgentRequest
		expectedRequestID string
		responses         []gRPCGetAgentResponse
		handlingDelay     time.Duration
		expectedResult    *syntheticspb.GetAgentResponse
		expectedErrorCode codes.Code
		expectedErrorMsg  string
		expectedError     bool
	}{
		{
			name:              "retry 2 times till success on code Unavailable",
			request:           &syntheticspb.GetAgentRequest{Id: testAgentID},
			expectedRequestID: testAgentID,
			responses: []gRPCGetAgentResponse{
				newErrorGRPCGetAgentResponse(codes.Unavailable),
				newErrorGRPCGetAgentResponse(codes.Unavailable),
				{
					nil,
					&syntheticspb.GetAgentResponse{Agent: newDummyAgent()},
				},
			},
			expectedResult: &syntheticspb.GetAgentResponse{Agent: newDummyAgent()},
			expectedError:  false,
		}, {
			name:    "retry 4 times when code Unavailable received and last code is Unavailable",
			request: &syntheticspb.GetAgentRequest{},
			responses: []gRPCGetAgentResponse{
				newErrorGRPCGetAgentResponse(codes.Unavailable),
				newErrorGRPCGetAgentResponse(codes.Unavailable),
				newErrorGRPCGetAgentResponse(codes.Unavailable),
				newErrorGRPCGetAgentResponse(codes.Unavailable),
				newErrorGRPCGetAgentResponse(codes.Unavailable),
			},
			expectedResult:    nil,
			expectedErrorCode: codes.Unavailable,
			expectedErrorMsg:  codes.Unavailable.String(),
			expectedError:     true,
		}, {
			name:    "retry 4 times when code Unavailable received and last code is Unknown",
			request: &syntheticspb.GetAgentRequest{},
			responses: []gRPCGetAgentResponse{
				newErrorGRPCGetAgentResponse(codes.Unavailable),
				newErrorGRPCGetAgentResponse(codes.Unavailable),
				newErrorGRPCGetAgentResponse(codes.Unavailable),
				newErrorGRPCGetAgentResponse(codes.Unavailable),
				newErrorGRPCGetAgentResponse(codes.Unknown),
			},
			expectedResult:    nil,
			expectedErrorCode: codes.Unknown,
			expectedErrorMsg:  codes.Unknown.String(),
			expectedError:     true,
		}, {
			name:    "do not retry when code Unknown received",
			request: &syntheticspb.GetAgentRequest{},
			responses: []gRPCGetAgentResponse{
				newErrorGRPCGetAgentResponse(codes.Unknown),
			},
			expectedResult:    nil,
			expectedErrorCode: codes.Unknown,
			expectedErrorMsg:  codes.Unknown.String(),
			expectedError:     true,
		}, {
			name:    "do not retry when retries disabled and code Unavailable received",
			options: []kentikapi.ClientOption{kentikapi.WithRetryMaxAttempts(0)},
			request: &syntheticspb.GetAgentRequest{},
			responses: []gRPCGetAgentResponse{
				newErrorGRPCGetAgentResponse(codes.Unavailable),
			},
			expectedResult:    nil,
			expectedErrorCode: codes.Unavailable,
			expectedErrorMsg:  codes.Unavailable.String(),
			expectedError:     true,
		}, {
			name:    "retry specified number of times when code Unavailable received",
			options: []kentikapi.ClientOption{kentikapi.WithRetryMaxAttempts(2)},
			request: &syntheticspb.GetAgentRequest{},
			responses: []gRPCGetAgentResponse{
				newErrorGRPCGetAgentResponse(codes.Unavailable),
				newErrorGRPCGetAgentResponse(codes.Unavailable),
				newErrorGRPCGetAgentResponse(codes.Unavailable),
			},
			expectedResult:    nil,
			expectedErrorCode: codes.Unavailable,
			expectedErrorMsg:  codes.Unavailable.String(),
			expectedError:     true,
		}, {
			name:              "timeout during first request",
			options:           []kentikapi.ClientOption{kentikapi.WithTimeout(5 * time.Millisecond)},
			request:           &syntheticspb.GetAgentRequest{},
			responses:         []gRPCGetAgentResponse{},
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
			server := newSpySyntheticsServerWithDelays(t, tt.responses)
			server.handlingDelay = tt.handlingDelay
			server.Start()
			defer server.Stop()

			options := []kentikapi.ClientOption{
				kentikapi.WithAPIURL("http://" + server.url),
				kentikapi.WithCredentials(dummyAuthEmail, dummyAuthToken),
				kentikapi.WithRetryMinDelay(10 * time.Microsecond),
				kentikapi.WithLogPayloads(),
			}
			client, err := kentikapi.NewClient(append(options, tt.options...)...)
			require.NoError(t, err)

			// act
			result, err := client.SyntheticsAdmin.GetAgent(context.Background(), tt.request)

			// assert
			t.Logf("Got result: %+v, err: %v", result, err)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if e, ok := status.FromError(err); ok {
				assert.Equal(t, tt.expectedErrorCode, e.Code())
				assert.Equal(t, tt.expectedErrorMsg, e.Message())
			}

			assert.Equal(t, len(server.responses), len(server.requests), "invalid number of requests")
			for i, r := range server.requests {
				assert.Equal(t, dummyAuthEmail, server.metadataSlice[i].Get(authEmailKey)[0])
				assert.Equal(t, dummyAuthToken, server.metadataSlice[i].Get(authAPITokenKey)[0])
				assert.Equal(t, tt.expectedRequestID, r.GetId())
			}

			testutil.AssertProtoEqual(t, tt.expectedResult, result)
		})
	}
}

type spySyntheticsServerWithDelays struct {
	syntheticspb.UnimplementedSyntheticsAdminServiceServer
	server *grpc.Server

	url  string
	done chan struct{}
	t    testing.TB
	// responses to return to the client
	responses []gRPCGetAgentResponse
	// handlingDelay specifies the delay applied while handling the request
	handlingDelay time.Duration

	// requests spied by the server
	requests      []*syntheticspb.GetAgentRequest
	metadataSlice []metadata.MD
}

type gRPCGetAgentResponse struct {
	err  error
	body *syntheticspb.GetAgentResponse
}

func newSpySyntheticsServerWithDelays(t testing.TB, responses []gRPCGetAgentResponse) *spySyntheticsServerWithDelays {
	return &spySyntheticsServerWithDelays{
		done:      make(chan struct{}),
		t:         t,
		responses: responses,
	}
}

func (s *spySyntheticsServerWithDelays) Start() {
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
func (s *spySyntheticsServerWithDelays) Stop() {
	s.server.Stop()
	<-s.done
}

func (s *spySyntheticsServerWithDelays) GetAgent(
	ctx context.Context, req *syntheticspb.GetAgentRequest,
) (*syntheticspb.GetAgentResponse, error) {
	time.Sleep(s.handlingDelay)

	md, _ := metadata.FromIncomingContext(ctx)
	s.metadataSlice = append(s.metadataSlice, md)

	s.requests = append(s.requests, req)

	response := s.response()
	return response.body, response.err
}

func (s *spySyntheticsServerWithDelays) response() gRPCGetAgentResponse {
	if len(s.requests) > len(s.responses) {
		return gRPCGetAgentResponse{
			status.Errorf(
				codes.Unknown,
				"spySyntheticsServerWithDelays: unexpected request, requests count: %v, expected: %v",
				len(s.requests),
				len(s.responses),
			),
			nil,
		}
	}
	return s.responses[len(s.requests)-1]
}

func newDummyAgent() *syntheticspb.Agent {
	return &syntheticspb.Agent{
		Id:            testAgentID,
		SiteName:      "dummy-site-name",
		Status:        syntheticspb.AgentStatus_AGENT_STATUS_WAIT,
		Alias:         "probe-4-ams-1",
		Type:          "global",
		Os:            "I use Manjaro BTW",
		Ip:            "95.179.136.58",
		Lat:           52.374031,
		Long:          4.88969,
		LastAuthed:    timestamppb.New(time.Date(2020, time.July, 9, 21, 37, 0, 826*1000000, time.UTC)),
		Family:        syntheticspb.IPFamily_IP_FAMILY_DUAL,
		Asn:           20473,
		SiteId:        "2137",
		Version:       "0.0.2",
		City:          "Amsterdam",
		Region:        "Noord-Holland",
		Country:       "Netherlands",
		TestIds:       []string{"13", "133", "1337"},
		LocalIp:       "10.10.10.10",
		CloudRegion:   "dummy-region",
		CloudProvider: "dummy-cloud-provider",
		AgentImpl:     syntheticspb.ImplementType_IMPLEMENT_TYPE_RUST,
	}
}

func newErrorGRPCGetAgentResponse(c codes.Code) gRPCGetAgentResponse {
	return gRPCGetAgentResponse{
		err: status.Errorf(
			c,
			c.String(),
		),
		body: nil,
	}
}
