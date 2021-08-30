package testutil

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	serverSleepTime = 10 * time.Millisecond
)

type SpyHTTPHandler struct {
	t testing.TB
	// Response to return to the client
	responseCode int
	responseBody []byte

	// Data spied by the handler
	RequestsCount   int
	LastMethod      string
	LastURL         *url.URL
	LastHeader      http.Header
	LastRequestBody string
}

func NewSpyHTTPHandler(t testing.TB, responseCode int, responseBody []byte) *SpyHTTPHandler {
	return &SpyHTTPHandler{
		t:            t,
		responseCode: responseCode,
		responseBody: responseBody,
	}
}

func (h *SpyHTTPHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.RequestsCount++
	h.LastMethod = r.Method
	h.LastURL = r.URL
	h.LastHeader = r.Header

	body, err := ioutil.ReadAll(r.Body)
	assert.NoError(h.t, err)
	h.LastRequestBody = string(body)

	err = r.Body.Close()
	assert.NoError(h.t, err)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(h.responseCode)
	_, err = rw.Write(h.responseBody)
	assert.NoError(h.t, err)
}

type MultipleResponseSpyHTTPHandler struct {
	t testing.TB
	// Responses to return to the client
	responses []HTTPResponse

	// Requests spied by the handler
	Requests []HTTPRequest
}

func NewMultipleResponseSpyHTTPHandler(t testing.TB, responses []HTTPResponse) *MultipleResponseSpyHTTPHandler {
	return &MultipleResponseSpyHTTPHandler{
		t:         t,
		responses: responses,
	}
}

func (h *MultipleResponseSpyHTTPHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	assert.NoError(h.t, err)

	err = r.Body.Close()
	assert.NoError(h.t, err)

	h.Requests = append(h.Requests, HTTPRequest{
		Method: r.Method,
		URL:    r.URL,
		Header: r.Header,
		Body:   string(body),
	})

	rw.Header().Set("Content-Type", "application/json")
	response := h.response()
	rw.WriteHeader(response.StatusCode)
	_, err = rw.Write([]byte(response.Body))
	assert.NoError(h.t, err)
}

func (h *MultipleResponseSpyHTTPHandler) response() HTTPResponse {
	if len(h.Requests) > len(h.responses) {
		return HTTPResponse{
			StatusCode: http.StatusGone,
			Body: fmt.Sprintf(
				"spyHTTPHandler: unexpected request, requests count: %v, expected: %v",
				len(h.Requests), len(h.responses),
			),
		}
	}

	return h.responses[len(h.Requests)-1]
}

type HTTPRequest struct {
	Method string
	URL    *url.URL
	Header http.Header
	Body   string
}

type HTTPResponse struct {
	StatusCode int
	Body       string
}

func NewErrorHTTPResponse(statusCode int) HTTPResponse {
	return HTTPResponse{
		StatusCode: statusCode,
		Body:       fmt.Sprintf(`{"error":"%v"}`, http.StatusText(statusCode)),
	}
}

type MultipleResponseSpyHTTPHandlerSleep struct {
	t testing.TB
	// Responses to return to the client
	responses []HTTPResponse

	// Requests spied by the handler
	Requests []HTTPRequest
}

func NewMultipleResponseSpyHTTPHandlerSleep(t testing.TB, responses []HTTPResponse) *MultipleResponseSpyHTTPHandlerSleep {
	return &MultipleResponseSpyHTTPHandlerSleep{
		t:         t,
		responses: responses,
	}
}

func (h *MultipleResponseSpyHTTPHandlerSleep) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	assert.NoError(h.t, err)

	err = r.Body.Close()
	assert.NoError(h.t, err)

	h.Requests = append(h.Requests, HTTPRequest{
		Method: r.Method,
		URL:    r.URL,
		Header: r.Header,
		Body:   string(body),
	})

	rw.Header().Set("Content-Type", "application/json")
	response := h.response()
	time.Sleep(serverSleepTime)
	rw.WriteHeader(response.StatusCode)
	_, err = rw.Write([]byte(response.Body))
	assert.NoError(h.t, err)
}

func (h *MultipleResponseSpyHTTPHandlerSleep) response() HTTPResponse {
	if len(h.Requests) > len(h.responses) {
		return HTTPResponse{
			StatusCode: http.StatusGone,
			Body: fmt.Sprintf(
				"spyHTTPHandler: unexpected request, requests count: %v, expected: %v",
				len(h.Requests), len(h.responses),
			),
		}
	}

	return h.responses[len(h.Requests)-1]
}
