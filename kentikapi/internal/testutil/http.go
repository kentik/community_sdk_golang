package testutil

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
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
