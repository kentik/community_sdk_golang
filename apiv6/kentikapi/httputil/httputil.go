package httputil

import (
	"context"
	"crypto/x509"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

// NewRetryingClient returns new http.Client with request retry strategy.
// Exponential backoff algorithm is used to calculate delay between retries.
// Retry-After header of HTTP 429 response is respected while calculating the retry delay.
//
// By default following retry policy is used:
// - Retry on following HTTP status codes: [429, 500, 502, 503, 504],
// - Retry on following HTTP request methods: [GET, HEAD, POST, PUT, PATCH, DELETE, CONNECT, OPTIONS, TRACE].
// - Retry on underlying http.Client.Do() error (except: too many redirects, invalid protocol scheme,
//   TLS cert verification failure).
//
// The client performs logging to os.Stderr.
func NewRetryingClient(cfg ClientConfig) *http.Client {
	cfg.FillDefaults()

	c := retryablehttp.NewClient()
	c.HTTPClient = cfg.HTTPClient
	if cfg.RetryMax != nil {
		c.RetryMax = *cfg.RetryMax
	}
	if cfg.RetryWaitMin != nil {
		c.RetryWaitMin = *cfg.RetryWaitMin
	}
	if cfg.RetryWaitMax != nil {
		c.RetryWaitMax = *cfg.RetryWaitMax
	}
	c.CheckRetry = makeRetryPolicy(cfg.RetryableStatusCodes, cfg.RetryableMethods)
	c.ErrorHandler = retryablehttp.PassthroughErrorHandler

	return c.StandardClient()
}

// ClientConfig holds configuration for retrying client.
type ClientConfig struct {
	HTTPClient           *http.Client
	RetryMax             *int
	RetryWaitMin         *time.Duration
	RetryWaitMax         *time.Duration
	RetryableStatusCodes []int
	RetryableMethods     []string
}

func (cfg *ClientConfig) FillDefaults() {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = defaultHTTPClient()
	}

	if len(cfg.RetryableStatusCodes) == 0 {
		cfg.RetryableStatusCodes = []int{429, 500, 502, 503, 504}
	}

	if len(cfg.RetryableMethods) == 0 {
		cfg.RetryableMethods = []string{
			http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete,
			http.MethodConnect, http.MethodOptions, http.MethodTrace,
		}
	}
}

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
	redirectsErrorRe := redirectsErrorRegexp()
	schemeErrorRe := schemeErrorRegexp()

	return func(ctx context.Context, resp *http.Response, err error) (bool, error) {
		// do not retry on context.Canceled or context.DeadlineExceeded
		if ctx.Err() != nil {
			return false, ctx.Err()
		}

		if !isRequestRetrayable(resp.Request, methodsSet) {
			return false, err
		}

		if err != nil {
			// TODO(dfurman): unit test that branch
			return isErrorRecoverable(err, redirectsErrorRe, schemeErrorRe)
		}

		if _, ok := statusCodesSet[resp.StatusCode]; ok {
			return true, nil
		}

		return false, nil
	}
}

func isRequestRetrayable(r *http.Request, methodsSet map[string]struct{}) bool {
	_, ok := methodsSet[r.Method]
	return ok
}

func isErrorRecoverable(err error, redirectsErrorRe, schemeErrorRe *regexp.Regexp) (bool, error) {
	if v, ok := err.(*url.Error); ok {
		// Don't retry if the error was due to too many redirects.
		if redirectsErrorRe.MatchString(v.Error()) {
			return false, v
		}

		// Don't retry if the error was due to an invalid protocol scheme.
		if schemeErrorRe.MatchString(v.Error()) {
			return false, v
		}

		// Don't retry if the error was due to TLS cert verification failure.
		if _, ok = v.Err.(x509.UnknownAuthorityError); ok {
			return false, v
		}
	}

	// The error is likely recoverable so retry.
	return true, nil
}

// redirectsErrorRegexp returns a regular expression to match the error returned by net/http when the
// configured number of redirects is exhausted. This error isn't typed
// specifically so we resort to matching on the error string.
func redirectsErrorRegexp() *regexp.Regexp {
	return regexp.MustCompile(`stopped after \d+ redirects\z`)
}

// schemeErrorRegexp returns a regular expression to match the error returned by net/http when the
// scheme specified in the URL is invalid. This error isn't typed
// specifically so we resort to matching on the error string.
func schemeErrorRegexp() *regexp.Regexp {
	return regexp.MustCompile(`unsupported protocol scheme`)
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
