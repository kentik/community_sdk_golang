package httputil

import (
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

// ClientConfig holds configuration for retrying client.
type ClientConfig struct {
	RetryMax     *int
	RetryWaitMin *time.Duration
	RetryWaitMax *time.Duration
}

// NewRetryingClient returns new http.Client with request retry strategy.
// Exponential backoff algorithm is used to calculate delay between retries.
// Retry-After header of HTTP 429 response is respected while calculating the retry delay.
//
// By default following retry policy is used:
// - Retry on following HTTP status codes: 429, all 5xx except 501.
// - Retry on all HTTP request methods.
// - Don't retry if the error was due to too many redirects.
// - Don't retry if the error was due to an invalid protocol scheme.
// - Don't retry if the error was due to TLS cert verification failure.
//
// The client performs logging to os.Stderr.
// TODO(dfurman): configurable retry policy: HTTP request methods, HTTP response status codes
func NewRetryingClient(cfg ClientConfig) *http.Client {
	c := retryablehttp.NewClient()

	if cfg.RetryMax != nil {
		c.RetryMax = *cfg.RetryMax
	}
	if cfg.RetryWaitMin != nil {
		c.RetryWaitMin = *cfg.RetryWaitMin
	}
	if cfg.RetryWaitMax != nil {
		c.RetryWaitMax = *cfg.RetryWaitMax
	}
	c.CheckRetry = retryablehttp.ErrorPropagatedRetryPolicy
	c.ErrorHandler = retryablehttp.PassthroughErrorHandler

	return c.StandardClient()
}
