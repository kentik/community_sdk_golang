package httputil

import (
	"context"
	"errors"
	"net"
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

// NewRetryingClient returns new retryablehttp.Client with request retry strategy.
// Exponential backoff algorithm is used to calculate delay between retries.
// Retry-After header of HTTP 429 response is respected while calculating the retry delay.
//
// By default following retry policy is used:
// - Retry on following HTTP status codes: [429, 500, 502, 503, 504],
// - Retry on following HTTP request methods: [GET, HEAD, POST, PUT, PATCH, DELETE, CONNECT, OPTIONS, TRACE].
// - Retry on underlying http.Client.Do() temporary errors.
//
// The client performs logging to os.Stderr.
func NewRetryingClient(cfg ClientConfig) *retryablehttp.Client {
	cfg.FillDefaults()

	c := retryablehttp.NewClient()
	c.HTTPClient = cfg.HTTPClient
	if cfg.RetryCfg.MaxAttempts != nil {
		c.RetryMax = *cfg.RetryCfg.MaxAttempts
	}
	if cfg.RetryCfg.MinDelay != nil {
		c.RetryWaitMin = *cfg.RetryCfg.MinDelay
	}
	if cfg.RetryCfg.MaxDelay != nil {
		c.RetryWaitMax = *cfg.RetryCfg.MaxDelay
	}
	c.CheckRetry = makeRetryPolicy(cfg.RetryCfg.RetryableStatusCodes, cfg.RetryCfg.RetryableMethods)
	c.ErrorHandler = retryablehttp.PassthroughErrorHandler

	return c
}

// NewRetryingStdClient returns new http.Client with request retry strategy.
// See NewRetryingClient for more information.
func NewRetryingStdClient(cfg ClientConfig) *http.Client {
	return NewRetryingClient(cfg).StandardClient()
}

// ClientConfig holds configuration for retrying client.
type ClientConfig struct {
	HTTPClient *http.Client
	RetryCfg   RetryConfig
}

// RetryConfig groups Client's configuration related to request retry functionality.
// See httputil.NewRetryingClient for retry policy description.
type RetryConfig struct {
	// MaxAttempts is a maximum number of request retry attempts. Set to 0 to disable retrying. Default: 4.
	MaxAttempts *int
	// MinDelay is a minimum delay before request retry. Default: 1 second.
	MinDelay *time.Duration
	// MaxDelay is a maximum delay before request retry. Default: 30 seconds.
	MaxDelay *time.Duration
	// RetryableStatusCodes are HTTP response status codes to retry on. Default: [429, 500, 502, 503, 504].
	RetryableStatusCodes []int
	// RetryableMethods are HTTP request retry methods, which the retry strategy is enabled for.
	// Default: [GET, HEAD, POST, PUT, PATCH, DELETE, CONNECT, OPTIONS, TRACE].
	RetryableMethods []string
}

func (cfg *ClientConfig) FillDefaults() {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = defaultHTTPClient()
	}

	if len(cfg.RetryCfg.RetryableStatusCodes) == 0 {
		cfg.RetryCfg.RetryableStatusCodes = []int{429, 500, 502, 503, 504}
	}

	if len(cfg.RetryCfg.RetryableMethods) == 0 {
		cfg.RetryCfg.RetryableMethods = []string{
			http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete,
			http.MethodConnect, http.MethodOptions, http.MethodTrace,
		}
	}
}

//nolint:gomnd // This is the only place for these numbers to turn up.
func defaultHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
}

// makeRetryPolicy creates customized retry policy.
// Its implementation is based on retryablehttp.ErrorPropagatedRetryPolicy.
// Retry policy function returns true if the request should be retried.
func makeRetryPolicy(statusCodes []int, methods []string) retryablehttp.CheckRetry {
	statusCodesSet := makeIntSet(statusCodes)
	methodsSet := makeStringSet(methods)

	return func(ctx context.Context, resp *http.Response, err error) (bool, error) {
		if ctx.Err() != nil {
			return false, ctx.Err()
		}

		if err != nil {
			if isErrorTemporary(err) {
				return true, nil
			}
			return false, nil
		}

		if !isRequestRetryable(resp.Request, methodsSet) {
			return false, nil
		}

		if _, ok := statusCodesSet[resp.StatusCode]; ok {
			return true, nil
		}

		return false, nil
	}
}

func isRequestRetryable(r *http.Request, methodsSet map[string]struct{}) bool {
	_, ok := methodsSet[r.Method]
	return ok
}

func isErrorTemporary(err error) bool {
	var tErr interface {
		Temporary() bool
	}
	if ok := errors.As(err, &tErr); ok {
		if tErr.Temporary() {
			return true
		}
	}

	return false
}

func makeIntSet(s []int) map[int]struct{} {
	result := make(map[int]struct{})
	for _, sc := range s {
		result[sc] = struct{}{}
	}
	return result
}

func makeStringSet(s []string) map[string]struct{} {
	result := make(map[string]struct{})
	for _, sc := range s {
		result[sc] = struct{}{}
	}
	return result
}
