package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"time"

	"github.com/kentik/community_sdk_golang/examples/demos"
	"github.com/kentik/community_sdk_golang/kentikapi"
)

func main() {
	err := showRetryingOnMultipleCodes()
	if err != nil {
		log.Fatal(err)
	}
}

func showRetryingOnMultipleCodes() error {
	demos.Step("Create fake Kentik API server")
	h := newSpyHTTPHandler([]httpResponse{
		newErrorHTTPResponse(http.StatusBadGateway),
		newErrorHTTPResponse(http.StatusBadGateway),
		newErrorHTTPResponse(http.StatusBadGateway),
		newErrorHTTPResponse(http.StatusBadGateway),
		newErrorHTTPResponse(http.StatusServiceUnavailable),
		newErrorHTTPResponse(http.StatusServiceUnavailable),
		newErrorHTTPResponse(http.StatusBadGateway),
		newErrorHTTPResponse(http.StatusTooManyRequests),
		newErrorHTTPResponse(http.StatusTooManyRequests),
		{
			statusCode: http.StatusOK,
			body:       dummyAgentsResponseBody,
		},
	})
	s := httptest.NewServer(h)
	fmt.Printf("Running fake server on URL %v\n", s.URL)

	demos.Step("Create Kentik API v6 client")
	c := kentikapi.NewClient(kentikapi.Config{
		SyntheticsAPIURL: s.URL,
		RetryCfg: kentikapi.RetryConfig{
			MaxAttempts:          intPtr(42),
			MinDelay:             durationPtr(100 * time.Millisecond),
			MaxDelay:             durationPtr(10 * time.Second),
			RetryableStatusCodes: []int{http.StatusTooManyRequests, http.StatusBadGateway, http.StatusServiceUnavailable},
			RetryableMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		},
	})

	demos.Step("List synthetic agents")
	result, _, err := c.SyntheticsAdminServiceAPI.
		AgentsList(context.Background()).
		Execute()

	fmt.Println("Received result.Agents:")
	demos.PrettyPrint(result.Agents)

	return err
}

const (
	retryAfterHeaderValue = "3"
)

type spyHTTPHandler struct {
	// responses to return to the client
	responses []httpResponse

	// requests spied by the handler
	requests []httpRequest
}

func newSpyHTTPHandler(responses []httpResponse) *spyHTTPHandler {
	return &spyHTTPHandler{
		responses: responses,
	}
}

func (h *spyHTTPHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeResponse(
			rw, http.StatusInternalServerError,
			fmt.Sprintf("spyHTTPHandler: failed to read request body: %v", err),
		)
	}

	err = r.Body.Close()
	if err != nil {
		log.Printf("spyHTTPHandler: failed to close request body: %v", err)
	}

	h.requests = append(h.requests, httpRequest{
		method: r.Method,
		url_:   r.URL,
		header: r.Header,
		body:   string(body),
	})

	retryAfterValue, _ := strconv.Atoi(retryAfterHeaderValue)
	retryAfterDate := time.Now().Add(time.Duration(retryAfterValue) * time.Second)

	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Retry-After", retryAfterDate.Format(time.RFC1123))
	response := h.response()
	writeResponse(rw, response.statusCode, response.body)
}

func writeResponse(rw http.ResponseWriter, statusCode int, body string) {
	rw.WriteHeader(statusCode)
	_, err := rw.Write([]byte(body))
	if err != nil {
		log.Printf("spyHTTPHandler: failed to write response body: %v", err)
	}
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

const dummyAgentsResponseBody string = `{
	"agents": [{
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
	}]
}`

func intPtr(v int) *int {
	return &v
}

func durationPtr(v time.Duration) *time.Duration {
	return &v
}
