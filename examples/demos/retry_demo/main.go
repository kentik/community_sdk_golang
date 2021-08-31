package main

import (
	"context"
	"fmt"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/httputil"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
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
		newErrorHTTPResponse(http.StatusTooManyRequests),
		newErrorHTTPResponse(http.StatusTooManyRequests),
		{
			statusCode: http.StatusOK,
			body:       dummyUsersResponseBody,
		},
	})
	s := httptest.NewServer(h)
	fmt.Printf("Running fake server on URL %v\n", s.URL)

	demos.Step("Create Kentik API v5 client")
	c := kentikapi.NewClient(kentikapi.Config{
		APIURL: s.URL,
		RetryCfg: httputil.RetryConfig{
			MaxAttempts:          intPtr(42),
			MinDelay:             durationPtr(100 * time.Millisecond),
			MaxDelay:             durationPtr(10 * time.Second),
			RetryableStatusCodes: []int{http.StatusTooManyRequests, http.StatusBadGateway, http.StatusServiceUnavailable},
			RetryableMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		},
	})

	demos.Step("List users")
	result, err := c.Users.GetAll(context.Background())

	fmt.Println("Received result:")
	demos.PrettyPrint(result)

	return err
}

const (
	retryAfterHeaderValue = "2"
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

	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Retry-After", retryAfterHeaderValue)
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

const dummyUsersResponseBody string = `{
	"users": [
		{
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
			"user_api_token": null,
			"filters": {},
			"saved_filters": []
		},
		{
			"id": "666666",
			"username": "Alice",
			"user_full_name": "Alice Awesome",
			"user_email": "alice.awesome@company.com",
			"role": "Administrator",
			"email_service": false,
			"email_product": false,
			"last_login": "2021-02-05T11:40:09.257Z",
			"created_date": "2021-01-05T12:49:21.306Z",
			"updated_date": "2021-02-05T11:40:09.258Z",
			"company_id": "74333",
			"user_api_token": null,
			"filters": {},
			"saved_filters": []
		}
	]
}`

func intPtr(v int) *int {
	return &v
}

func durationPtr(v time.Duration) *time.Duration {
	return &v
}
