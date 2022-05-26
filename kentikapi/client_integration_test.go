package kentikapi_test

import (
	"context"
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
)

const (
	warsawAgentID = "41910"
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
		errorPredicates     []func(error) bool
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
			errorPredicates:     []func(error) bool{kentikapi.IsInvalidRequestError},
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
			errorPredicates:     []func(error) bool{kentikapi.IsInvalidRequestError},
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
			errorPredicates:     []func(error) bool{kentikapi.IsTimeoutError, kentikapi.IsTemporaryError},
		}, {
			name: "authorization error",
			responses: []testutil.HTTPResponse{
				testutil.NewErrorHTTPResponse(http.StatusUnauthorized),
			},
			expectedError:   true,
			expectedTimeout: false,
			errorPredicates: []func(error) bool{kentikapi.IsAuthError},
		}, {
			name: "too many requests error",
			responses: []testutil.HTTPResponse{
				testutil.NewErrorHTTPResponse(http.StatusTooManyRequests),
				testutil.NewErrorHTTPResponse(http.StatusTooManyRequests),
				testutil.NewErrorHTTPResponse(http.StatusTooManyRequests),
				testutil.NewErrorHTTPResponse(http.StatusTooManyRequests),
				testutil.NewErrorHTTPResponse(http.StatusTooManyRequests),
			},
			expectedError:   true,
			expectedTimeout: false,
			errorPredicates: []func(error) bool{kentikapi.IsRateLimitExhaustedError, kentikapi.IsTemporaryError},
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
				for _, isErr := range tt.errorPredicates {
					assert.True(t, isErr(err))
				}
			} else {
				assert.NoError(t, err)
			}

			if tt.expectedTimeout {
				assert.True(t, kentikapi.IsTimeoutError(err))
			} else {
				assert.False(t, kentikapi.IsTimeoutError(err))
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
		name            string
		options         []kentikapi.ClientOption
		responses       []gRPCGetAgentResponse
		handlingDelay   time.Duration
		expectedResult  *models.SyntheticsAgent
		expectedError   bool
		errorPredicates []func(error) bool
	}{
		{
			name: "retry 2 times till success on code Unavailable",
			responses: []gRPCGetAgentResponse{
				newErrorGRPCGetAgentResponse(codes.Unavailable),
				newErrorGRPCGetAgentResponse(codes.Unavailable),
				{
					nil,
					&syntheticspb.GetAgentResponse{Agent: newWarsawAgentPayload()},
				},
			},
			expectedResult: newWarsawAgent(),
			expectedError:  false,
		}, {
			name: "retry 4 times when code Unavailable received and last code is Unavailable",
			responses: []gRPCGetAgentResponse{
				newErrorGRPCGetAgentResponse(codes.Unavailable),
				newErrorGRPCGetAgentResponse(codes.Unavailable),
				newErrorGRPCGetAgentResponse(codes.Unavailable),
				newErrorGRPCGetAgentResponse(codes.Unavailable),
				newErrorGRPCGetAgentResponse(codes.Unavailable),
			},
			expectedResult:  nil,
			expectedError:   true,
			errorPredicates: []func(error) bool{kentikapi.IsTemporaryError},
		}, {
			name: "retry 4 times when code Unavailable received and last code is Unknown",
			responses: []gRPCGetAgentResponse{
				newErrorGRPCGetAgentResponse(codes.Unavailable),
				newErrorGRPCGetAgentResponse(codes.Unavailable),
				newErrorGRPCGetAgentResponse(codes.Unavailable),
				newErrorGRPCGetAgentResponse(codes.Unavailable),
				newErrorGRPCGetAgentResponse(codes.Unknown),
			},
			expectedResult: nil,
			expectedError:  true,
		}, {
			name: "do not retry when code Unknown received",
			responses: []gRPCGetAgentResponse{
				newErrorGRPCGetAgentResponse(codes.Unknown),
			},
			expectedResult: nil,
			expectedError:  true,
		}, {
			name:    "do not retry when retries disabled and code Unavailable received",
			options: []kentikapi.ClientOption{kentikapi.WithRetryMaxAttempts(0)},
			responses: []gRPCGetAgentResponse{
				newErrorGRPCGetAgentResponse(codes.Unavailable),
			},
			expectedResult:  nil,
			expectedError:   true,
			errorPredicates: []func(error) bool{kentikapi.IsTemporaryError},
		}, {
			name:    "retry specified number of times when code Unavailable received",
			options: []kentikapi.ClientOption{kentikapi.WithRetryMaxAttempts(2)},
			responses: []gRPCGetAgentResponse{
				newErrorGRPCGetAgentResponse(codes.Unavailable),
				newErrorGRPCGetAgentResponse(codes.Unavailable),
				newErrorGRPCGetAgentResponse(codes.Unavailable),
			},
			expectedResult:  nil,
			expectedError:   true,
			errorPredicates: []func(error) bool{kentikapi.IsTemporaryError},
		}, {
			name:            "timeout during first request",
			options:         []kentikapi.ClientOption{kentikapi.WithTimeout(5 * time.Millisecond)},
			responses:       []gRPCGetAgentResponse{},
			handlingDelay:   1 * time.Second,
			expectedResult:  nil,
			expectedError:   true,
			errorPredicates: []func(error) bool{kentikapi.IsTimeoutError, kentikapi.IsTemporaryError},
		}, {
			name: "Unauthenticated error",
			responses: []gRPCGetAgentResponse{
				newErrorGRPCGetAgentResponse(codes.Unauthenticated),
			},
			expectedResult:  nil,
			expectedError:   true,
			errorPredicates: []func(error) bool{kentikapi.IsAuthError},
		}, {
			name:    "ResourceExhausted error",
			options: []kentikapi.ClientOption{kentikapi.WithRetryMaxAttempts(3)},
			responses: []gRPCGetAgentResponse{
				newErrorGRPCGetAgentResponse(codes.ResourceExhausted),
				newErrorGRPCGetAgentResponse(codes.ResourceExhausted),
				newErrorGRPCGetAgentResponse(codes.ResourceExhausted),
				newErrorGRPCGetAgentResponse(codes.ResourceExhausted),
			},
			expectedResult: nil,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			server := newSpySyntheticsServerWithDelays(t, tt.responses, tt.handlingDelay)
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
			result, err := client.Synthetics.Agents.Get(context.Background(), warsawAgentID)

			// assert
			t.Logf("Got result: %+v, err: %v", result, err)
			if tt.expectedError {
				assert.Error(t, err)
				for _, isErr := range tt.errorPredicates {
					assert.True(t, isErr(err))
				}
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, len(server.responses), len(server.requests), "invalid number of requests")
			for _, r := range server.requests {
				assert.Equal(t, dummyAuthEmail, r.metadata.Get(authEmailKey)[0])
				assert.Equal(t, dummyAuthToken, r.metadata.Get(authAPITokenKey)[0])
				assert.Equal(t, warsawAgentID, r.data.GetId())
			}

			assert.Equal(t, tt.expectedResult, result)
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
	requests []getAgentRequest
}

type gRPCGetAgentResponse struct {
	err  error
	body *syntheticspb.GetAgentResponse
}

func newSpySyntheticsServerWithDelays(
	t testing.TB, responses []gRPCGetAgentResponse, handlingDelay time.Duration,
) *spySyntheticsServerWithDelays {
	return &spySyntheticsServerWithDelays{
		done:          make(chan struct{}),
		t:             t,
		responses:     responses,
		handlingDelay: handlingDelay,
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
	s.requests = append(s.requests, getAgentRequest{
		metadata: md,
		data:     req,
	})

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

func newErrorGRPCGetAgentResponse(c codes.Code) gRPCGetAgentResponse {
	return gRPCGetAgentResponse{
		err: status.Errorf(
			c,
			c.String(),
		),
		body: nil,
	}
}
