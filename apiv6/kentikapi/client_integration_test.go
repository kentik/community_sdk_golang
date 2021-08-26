package kentikapi_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/kentik/community_sdk_golang/apiv6/kentikapi"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/synthetics"
	"github.com/stretchr/testify/assert"
)

const (
	authEmailKey    = "X-CH-Auth-Email"
	authAPITokenKey = "X-CH-Auth-API-Token"
	dummyAuthEmail  = "email@example.com"
	dummyAuthToken  = "api-test-token"
	testAgentID     = "968"
)

func TestClient_PatchAgent(t *testing.T) {
	tests := []struct {
		name             string
		retryMax         *int
		retryStatusCodes []int
		retryMethods     []string
		request          synthetics.V202101beta1PatchAgentRequest
		// expectedRequestBody is map for the granularity of assertion diff
		expectedRequestBody map[string]interface{}
		responses           []httpResponse
		expectedResult      synthetics.V202101beta1PatchAgentResponse
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
			expectedResult: synthetics.V202101beta1PatchAgentResponse{Agent: newDummyAgent()},
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
			expectedResult: synthetics.V202101beta1PatchAgentResponse{Agent: newDummyAgent()},
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
		}, {
			name:                "retry only on configured HTTP response status codes if given any",
			retryStatusCodes:    []int{777, 888, 999},
			request:             newPatchAgentNameRequest(),
			expectedRequestBody: newPatchAgentNameRequestBody(),
			responses: []httpResponse{
				newErrorHTTPResponse(777),
				newErrorHTTPResponse(888),
				newErrorHTTPResponse(999),
				newErrorHTTPResponse(http.StatusTooManyRequests),
			},
			expectedError: true,
		}, {
			name:                "retry only on configured HTTP request methods if given any - retry case",
			retryMethods:        []string{http.MethodTrace, http.MethodPatch},
			request:             newPatchAgentNameRequest(),
			expectedRequestBody: newPatchAgentNameRequestBody(),
			responses: []httpResponse{
				newErrorHTTPResponse(http.StatusTooManyRequests),
				{
					statusCode: http.StatusOK,
					body:       dummyAgentResponseBody,
				},
			},
			expectedResult: synthetics.V202101beta1PatchAgentResponse{Agent: newDummyAgent()},
		}, {
			name:                "retry only on configured HTTP request methods if given any - no retry case",
			retryMethods:        []string{http.MethodTrace},
			request:             newPatchAgentNameRequest(),
			expectedRequestBody: newPatchAgentNameRequestBody(),
			responses: []httpResponse{
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

			c := kentikapi.NewClient(kentikapi.Config{
				SyntheticsAPIURL: s.URL,
				AuthEmail:        dummyAuthEmail,
				AuthToken:        dummyAuthToken,
				RetryCfg: kentikapi.RetryConfig{
					MaxAttempts:          tt.retryMax,
					MinDelay:             durationPtr(1 * time.Microsecond),
					MaxDelay:             durationPtr(10 * time.Microsecond),
					RetryableStatusCodes: tt.retryStatusCodes,
					RetryableMethods:     tt.retryMethods,
				},
				LogPayloads: true,
			})

			// act
			result, httpResp, err := c.SyntheticsAdminServiceApi.
				AgentPatch(context.Background(), testAgentID).
				Body(tt.request).
				Execute()

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
				assert.Equal(t, fmt.Sprintf("/synthetics/v202101beta1/agents/%v", testAgentID), r.url_.Path)
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

func newPatchAgentNameRequest() synthetics.V202101beta1PatchAgentRequest {
	return synthetics.V202101beta1PatchAgentRequest{
		Agent: newDummyAgent(),
		Mask:  stringPtr("agent.name"),
	}
}

func newDummyAgent() *synthetics.V202101beta1Agent {
	status := synthetics.V202101BETA1AGENTSTATUS_WAIT
	family := synthetics.V202101BETA1IPFAMILY_DUAL
	agent := &synthetics.V202101beta1Agent{
		Id:         stringPtr(testAgentID),
		Name:       stringPtr("dummy-agent"),
		Status:     &status,
		Alias:      stringPtr("probe-4-ams-1"),
		Type:       stringPtr("global"),
		Os:         stringPtr("I use Manjaro BTW"),
		Ip:         stringPtr("95.179.136.58"),
		Lat:        float64Ptr(52.374031),
		Long:       float64Ptr(4.88969),
		LastAuthed: timePtr(time.Date(2020, time.July, 9, 21, 37, 00, 826*1000000, time.UTC)),
		Family:     &family,
		Asn:        int64Ptr(20473),
		SiteId:     stringPtr("2137"),
		Version:    stringPtr("0.0.2"),
		Challenge:  stringPtr("dummy-challenge"),
		City:       stringPtr("Amsterdam"),
		Region:     stringPtr("Noord-Holland"),
		Country:    stringPtr("Netherlands"),
		TestIds:    &[]string{"13", "133", "1337"},
		LocalIp:    stringPtr("10.10.10.10"),
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
		url_:   r.URL,
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
	url_   *url.URL
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
