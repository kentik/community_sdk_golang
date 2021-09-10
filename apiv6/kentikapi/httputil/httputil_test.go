package httputil_test

import (
	"context"
	"errors"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/httputil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Sequence of calls for single request:
// 1. retryingClient.Do()
// 2. retryingClient.retryableRoundTripper.RoundTrip()
// 3. retryingClient.retryableRoundTripper.retryableClient.Do()
// 4. retryingClient.retryableRoundTripper.retryableClient.httpClient.Do()
// 5. retryingClient.retryableRoundTripper.retryableClient.httpClient.httpTransport.RoundTrip()

func TestRetryingClient_Do_ReturnsHTTPTransportError(t *testing.T) {
	// arrange
	c := httputil.NewRetryingClient(httputil.ClientConfig{})

	req, err := retryablehttp.NewRequest(http.MethodGet, "https://invalid.url", nil)
	require.NoError(t, err)

	// act
	//nolint:bodyclose
	resp, err := c.Do(req.WithContext(context.Background()))

	// assert
	t.Logf("Got response: %v, err: %v", resp, err)

	var dnsErr *net.DNSError
	require.True(t, errors.As(err, &dnsErr))
	assert.Equal(t, "no such host", dnsErr.Err)
}

func TestRetryingClientWithSpyHTTPTransport_Do(t *testing.T) {
	const retryMax = 5

	tests := []struct {
		name                  string
		transportError        error
		expectedRequestsCount int
	}{
		{
			name: "retries when underlying client returns temporary URL error",
			transportError: &url.Error{
				Err: &net.OpError{
					Err: &net.DNSError{
						Err:         "fake error",
						IsTemporary: true,
					},
				},
			},
			expectedRequestsCount: retryMax + 1,
		}, {
			name: "does not retry when underlying client returns non-temporary URL error",
			transportError: &url.Error{
				Err: &net.OpError{
					Err: &net.DNSError{
						Err:         "fake error",
						IsTemporary: false,
					},
				},
			},
			expectedRequestsCount: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			st := spyTransport{transportError: tt.transportError}
			c := httputil.NewRetryingClient(httputil.ClientConfig{
				HTTPClient: &http.Client{
					Transport: &st,
				},
				RetryCfg: httputil.RetryConfig{
					MaxAttempts: intPtr(retryMax),
					MinDelay:    durationPtr(1 * time.Microsecond),
					MaxDelay:    durationPtr(10 * time.Microsecond),
				},
			})

			req, err := retryablehttp.NewRequest(http.MethodGet, "https://dummy.url", nil)
			require.NoError(t, err)

			// act
			//nolint:bodyclose
			resp, err := c.Do(req.WithContext(context.Background()))

			// assert
			t.Logf("Got response: %v, err: %v", resp, err)
			assert.Equal(t, tt.expectedRequestsCount, st.requestsCount)

			var dnsErr *net.DNSError
			require.True(t, errors.As(err, &dnsErr))
			assert.Equal(t, "fake error", dnsErr.Err)
		})
	}
}

func TestRetryingClientRequestTimeout(t *testing.T) {
	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Second)
		_, err := io.WriteString(w, "done")
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	c := NewRetryingClient(ClientConfig{
		RetryCfg: RetryConfig{
			MaxAttempts: intPtr(2),
			MinDelay:    durationPtr(750 * time.Millisecond),
			MaxDelay:    durationPtr(10000 * time.Millisecond),
		},
	})

	backend := httptest.NewServer(handlerFunc)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	testUrl := backend.URL
	req, err := retryablehttp.NewRequest(http.MethodGet, testUrl, nil)
	if err != nil {
		t.Error("Request error", err)
		return
	}

	resp, err := c.Do(req.WithContext(ctx))
	assert.Error(t, err)
	t.Logf("Got response: %v, err: %v", resp, err)

	assert.True(t, isTimeoutError(err), "timeout error expected")
}

type spyTransport struct {
	transportError error
	requestsCount  int
}

func (t *spyTransport) RoundTrip(_ *http.Request) (*http.Response, error) {
	t.requestsCount++
	return nil, t.transportError
}

func intPtr(v int) *int {
	return &v
}

func durationPtr(v time.Duration) *time.Duration {
	return &v
}
