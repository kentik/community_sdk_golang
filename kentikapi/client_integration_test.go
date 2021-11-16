package kentikapi_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	syntheticspb "github.com/kentik/api-schema-public/gen/go/kentik/synthetics/v202101beta1"
	"github.com/kentik/community_sdk_golang/kentikapi"
	httpSynthetics "github.com/kentik/community_sdk_golang/kentikapi/synthetics"
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

func TestClient_PatchAgentHTTP(t *testing.T) {
	tests := []struct {
		name     string
		retryMax *int
		request  httpSynthetics.V202101beta1PatchAgentRequest
		// expectedRequestBody is map for the granularity of assertion diff
		expectedRequestBody map[string]interface{}
		responses           []httpResponse
		expectedResult      httpSynthetics.V202101beta1PatchAgentResponse
		expectedError       bool
	}{
		{
			name:                "empty request, status 400 Bad Request received",
			expectedRequestBody: map[string]interface{}{},
			responses:           []httpResponse{newErrorHTTPResponse(http.StatusBadRequest)},
			expectedError:       true,
		}, {
			name:                "status 400 Bad Request received",
			request:             newPatchAgentNameRequest(),
			expectedRequestBody: newPatchAgentNameRequestBody(),
			responses:           []httpResponse{newErrorHTTPResponse(http.StatusBadRequest)},
			expectedError:       true,
		}, {
			name:                "name updated",
			request:             newPatchAgentNameRequest(),
			expectedRequestBody: newPatchAgentNameRequestBody(),
			responses: []httpResponse{{
				statusCode: http.StatusOK,
				body:       dummyAgentResponseBody,
			}},
			expectedResult: httpSynthetics.V202101beta1PatchAgentResponse{Agent: newDummyAgent()},
		}, {
			name:                "retry till success when status 429 Too Many Requests received",
			request:             newPatchAgentNameRequest(),
			expectedRequestBody: newPatchAgentNameRequestBody(),
			responses: []httpResponse{
				newErrorHTTPResponse(http.StatusTooManyRequests),
				newErrorHTTPResponse(http.StatusTooManyRequests),
				newErrorHTTPResponse(http.StatusTooManyRequests),
				{
					statusCode: http.StatusOK,
					body:       dummyAgentResponseBody,
				},
			},
			expectedResult: httpSynthetics.V202101beta1PatchAgentResponse{Agent: newDummyAgent()},
		}, {
			name:                "retry 4 times when status 429, 500, 502, 503 or 504 received and last status is 429",
			request:             newPatchAgentNameRequest(),
			expectedRequestBody: newPatchAgentNameRequestBody(),
			responses: []httpResponse{
				newErrorHTTPResponse(http.StatusInternalServerError),
				newErrorHTTPResponse(http.StatusBadGateway),
				newErrorHTTPResponse(http.StatusServiceUnavailable),
				newErrorHTTPResponse(http.StatusGatewayTimeout),
				newErrorHTTPResponse(http.StatusTooManyRequests),
			},
			expectedError: true,
		}, {
			name:                "retry 4 times when status 429, 500, 502, 503 or 504 received and last status is 504",
			request:             newPatchAgentNameRequest(),
			expectedRequestBody: newPatchAgentNameRequestBody(),
			responses: []httpResponse{
				newErrorHTTPResponse(http.StatusTooManyRequests),
				newErrorHTTPResponse(http.StatusInternalServerError),
				newErrorHTTPResponse(http.StatusBadGateway),
				newErrorHTTPResponse(http.StatusServiceUnavailable),
				newErrorHTTPResponse(http.StatusGatewayTimeout),
			},
			expectedError: true,
		}, {
			name:                "do not retry when status 505 HTTP Version Not Supported received",
			request:             newPatchAgentNameRequest(),
			expectedRequestBody: newPatchAgentNameRequestBody(),
			responses: []httpResponse{
				newErrorHTTPResponse(http.StatusHTTPVersionNotSupported),
			},
			expectedError: true,
		}, {
			name:                "do not retry when retries disabled and status 429 Too Many Requests received",
			retryMax:            intPtr(0),
			request:             newPatchAgentNameRequest(),
			expectedRequestBody: newPatchAgentNameRequestBody(),
			responses: []httpResponse{
				newErrorHTTPResponse(http.StatusTooManyRequests),
			},
			expectedError: true,
		}, {
			name:                "retry specified number of times when status 429 Too Many Requests received",
			retryMax:            intPtr(2),
			request:             newPatchAgentNameRequest(),
			expectedRequestBody: newPatchAgentNameRequestBody(),
			responses: []httpResponse{
				newErrorHTTPResponse(http.StatusTooManyRequests),
				newErrorHTTPResponse(http.StatusTooManyRequests),
				newErrorHTTPResponse(http.StatusTooManyRequests),
			},
			expectedError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			h := newSpyHTTPHandler(t, tt.responses)
			s := httptest.NewServer(h)
			defer s.Close()

			c, err := kentikapi.NewClient(kentikapi.Config{
				SyntheticsAPIURL: s.URL,
				AuthEmail:        dummyAuthEmail,
				AuthToken:        dummyAuthToken,
				RetryCfg: kentikapi.RetryConfig{
					MaxAttempts: tt.retryMax,
					MinDelay:    durationPtr(1 * time.Microsecond),
					MaxDelay:    durationPtr(10 * time.Microsecond),
				},
				LogPayloads: true,
			})
			require.NoError(t, err)

			// act
			result, httpResp, err := c.SyntheticsAdminServiceAPI.
				AgentPatch(context.Background(), testAgentID).
				Body(tt.request).
				Execute()
			//nolint:errcheck // Additional info: https://github.com/kisielk/errcheck/issues/55
			defer httpResp.Body.Close()

			// assert
			t.Logf("Got result: %v, httpResp: %v, err: %v", result, httpResp, err)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, len(h.responses), len(h.requests), "invalid number of requests")
			for _, r := range h.requests {
				assert.Equal(t, http.MethodPatch, r.method)
				assert.Equal(t, fmt.Sprintf("/synthetics/v202101beta1/agents/%v", testAgentID), r.url.Path)
				assert.Equal(t, dummyAuthEmail, r.header.Get(authEmailKey))
				assert.Equal(t, dummyAuthToken, r.header.Get(authAPITokenKey))
				assert.Equal(t, tt.expectedRequestBody, unmarshalJSONToIf(t, r.body))
			}

			assert.Equal(t, tt.expectedResult, result)

			if assert.NotNil(t, httpResp) {
				assert.Equal(t, tt.responses[len(tt.responses)-1].statusCode, httpResp.StatusCode,
					"invalid HTTP response status code",
				)
			}
		})
	}
}

func newPatchAgentNameRequest() httpSynthetics.V202101beta1PatchAgentRequest {
	return httpSynthetics.V202101beta1PatchAgentRequest{
		Agent: newDummyAgent(),
		Mask:  stringPtr("agent.name"),
	}
}

func newDummyAgent() *httpSynthetics.V202101beta1Agent {
	status := httpSynthetics.V202101BETA1AGENTSTATUS_WAIT
	family := httpSynthetics.V202101BETA1IPFAMILY_DUAL
	agent := &httpSynthetics.V202101beta1Agent{
		Id:     stringPtr(testAgentID),
		Name:   stringPtr("dummy-agent"),
		Status: &status,
		Alias:  stringPtr("probe-4-ams-1"),
		Type:   stringPtr("global"),
		Os:     stringPtr("I use Manjaro BTW"),
		Ip:     stringPtr("95.179.136.58"),
		Lat:    float64Ptr(52.374031),
		Long:   float64Ptr(4.88969),
		LastAuthed: timePtr(time.Date(2020,
			time.July,
			9,
			21,
			37,
			0,
			826*1000000,
			time.UTC,
		)),
		Family:    &family,
		Asn:       int64Ptr(20473),
		SiteId:    stringPtr("2137"),
		Version:   stringPtr("0.0.2"),
		Challenge: stringPtr("dummy-challenge"),
		City:      stringPtr("Amsterdam"),
		Region:    stringPtr("Noord-Holland"),
		Country:   stringPtr("Netherlands"),
		TestIds:   &[]string{"13", "133", "1337"},
		LocalIp:   stringPtr("10.10.10.10"),
	}

	return agent
}

func newPatchAgentNameRequestBody() map[string]interface{} {
	return map[string]interface{}{
		"agent": newDummyAgentRequestBody(),
		"mask":  "agent.name",
	}
}

func newDummyAgentRequestBody() map[string]interface{} {
	return map[string]interface{}{
		"id":         "968",
		"name":       "dummy-agent",
		"status":     "AGENT_STATUS_WAIT",
		"alias":      "probe-4-ams-1",
		"type":       "global",
		"os":         "I use Manjaro BTW",
		"ip":         "95.179.136.58",
		"lat":        52.374031,
		"long":       4.88969,
		"lastAuthed": "2020-07-09T21:37:00.826Z",
		"family":     "IP_FAMILY_DUAL",
		"asn":        20473.0,
		"siteId":     "2137",
		"version":    "0.0.2",
		"challenge":  "dummy-challenge",
		"city":       "Amsterdam",
		"region":     "Noord-Holland",
		"country":    "Netherlands",
		"testIds":    []interface{}{"13", "133", "1337"},
		"localIp":    "10.10.10.10",
	}
}

const dummyAgentResponseBody string = `{
	"agent": {
		"id": "968",
		"name": "dummy-agent",
		"status": "AGENT_STATUS_WAIT",
		"alias": "probe-4-ams-1",
		"type": "global",
		"os": "I use Manjaro BTW",
		"ip": "95.179.136.58",
		"lat": 52.374031,
		"long": 4.88969,
		"lastAuthed": "2020-07-09T21:37:00.826Z",
		"family": "IP_FAMILY_DUAL",
		"asn": 20473,
		"siteId": "2137",
		"version": "0.0.2",
		"challenge": "dummy-challenge",
		"city": "Amsterdam",
		"region": "Noord-Holland",
		"country": "Netherlands",
		"testIds": [
			"13",
			"133",
			"1337"
		],
		"localIp": "10.10.10.10"
	}
}`

type spyHTTPHandler struct {
	t testing.TB
	// responses to return to the client
	responses []httpResponse

	// requests spied by the handler
	requests []httpRequest
}

func newSpyHTTPHandler(t testing.TB, responses []httpResponse) *spyHTTPHandler {
	return &spyHTTPHandler{
		t:         t,
		responses: responses,
	}
}

func (h *spyHTTPHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	assert.NoError(h.t, err)

	err = r.Body.Close()
	assert.NoError(h.t, err)

	h.requests = append(h.requests, httpRequest{
		method: r.Method,
		url:    r.URL,
		header: r.Header,
		body:   string(body),
	})

	rw.Header().Set("Content-Type", "application/json")
	response := h.response()
	rw.WriteHeader(response.statusCode)
	_, err = rw.Write([]byte(response.body))
	assert.NoError(h.t, err)
}

func (h *spyHTTPHandler) response() httpResponse {
	if len(h.requests) > len(h.responses) {
		return httpResponse{
			statusCode: http.StatusGone,
			body: fmt.Sprintf(
				"spyHTTPHandler: unexpected request, requests count: %v, expected: %v",
				len(h.requests), len(h.responses),
			),
		}
	}

	return h.responses[len(h.requests)-1]
}

type httpRequest struct {
	method string
	url    *url.URL
	header http.Header
	body   string
}

type httpResponse struct {
	statusCode int
	body       string
}

func newErrorHTTPResponse(statusCode int) httpResponse {
	return httpResponse{
		statusCode: statusCode,
		body:       fmt.Sprintf(`{"error":"%v"}`, http.StatusText(statusCode)),
	}
}

func unmarshalJSONToIf(t testing.TB, jsonString string) interface{} {
	var data interface{}
	err := json.Unmarshal([]byte(jsonString), &data)
	assert.NoError(t, err)
	return data
}

func stringPtr(s string) *string {
	return &s
}

func intPtr(v int) *int {
	return &v
}

func int64Ptr(v int64) *int64 {
	return &v
}

func float64Ptr(v float64) *float64 {
	return &v
}

func timePtr(v time.Time) *time.Time {
	return &v
}

func durationPtr(v time.Duration) *time.Duration {
	return &v
}

func TestClient_GetAgent(t *testing.T) {
	tests := []struct {
		name              string
		retryMax          *int
		request           *syntheticspb.GetAgentRequest
		expectedRequestID string
		responses         []gRPCResponse
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
			expectedResult:    &syntheticspb.GetAgentResponse{},
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
			expectedResult:    &syntheticspb.GetAgentResponse{},
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
			expectedResult:    &syntheticspb.GetAgentResponse{Agent: newDummyAgentForGRPC()},
			expectedErrorCode: codes.OK,
			expectedErrorMsg:  "",
			expectedError:     false,
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
			retryMax: intPtr(0),
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
			retryMax: intPtr(2),
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
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			s := newSpySyntheticsServer(t, tt.responses)
			s.Start()
			defer s.Stop()

			client, err := kentikapi.NewClient(kentikapi.Config{
				SyntheticsGRPCHostPort: s.url,
				AuthToken:              dummyAuthToken,
				AuthEmail:              dummyAuthEmail,
				DisableTLS:             true,
				RetryCfg: kentikapi.RetryConfig{
					MaxAttempts: tt.retryMax,
					MinDelay:    durationPtr(10 * time.Millisecond),
				},
			})
			require.NoError(t, err)

			// act
			response, err := client.SyntheticsAdmin.GetAgent(
				context.Background(),
				tt.request,
			)

			// assert
			t.Logf("Got response: %v, err: %v", response, err)
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

			assert.Equal(t, tt.expectedResult.GetAgent().String(), response.GetAgent().String())
		})
	}
}

type spySyntheticsServer struct {
	syntheticspb.UnimplementedSyntheticsAdminServiceServer
	server *grpc.Server
	url    string
	t      testing.TB
	// responses to return to the client
	responses []gRPCResponse

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
		require.NoError(s.t, s.server.Serve(l))
	}()
}

func (s *spySyntheticsServer) Stop() {
	s.server.GracefulStop()
}

func (s *spySyntheticsServer) GetAgent(ctx context.Context, req *syntheticspb.GetAgentRequest,
) (*syntheticspb.GetAgentResponse, error) {
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
