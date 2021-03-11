package testutil

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SpyHTTPHandler struct {
	t            testing.TB
	responseCode int
	responseBody []byte

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

func (s *SpyHTTPHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	s.RequestsCount++
	s.LastMethod = r.Method
	s.LastURL = r.URL
	s.LastHeader = r.Header

	body, err := ioutil.ReadAll(r.Body)
	assert.NoError(s.t, err)
	s.LastRequestBody = string(body)

	err = r.Body.Close()
	assert.NoError(s.t, err)

	rw.WriteHeader(s.responseCode)
	_, err = rw.Write(s.responseBody)
	assert.NoError(s.t, err)
}
