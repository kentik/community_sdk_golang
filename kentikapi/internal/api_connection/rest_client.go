package api_connection

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	kentikErrors "github.com/kentik/community_sdk_golang/kentikapi/internal/errors"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/httputil"
)

//nolint:gosec
const (
	authEmailKey    = "X-CH-Auth-Email"
	authAPITokenKey = "X-CH-Auth-API-Token"
)

type RestClient struct {
	config     RestClientConfig
	httpClient *retryablehttp.Client
}

type RestClientConfig struct {
	APIURL    string
	AuthEmail string
	AuthToken string
	RetryCfg  httputil.RetryConfig
	Timeout   *time.Duration
}

func NewRestClient(c RestClientConfig) *RestClient {
	return &RestClient{
		config: c,
		httpClient: httputil.NewRetryingClient(
			httputil.ClientConfig{
				RetryCfg: c.RetryCfg,
			},
		),
	}
}

// Get sends GET request to the API and returns raw response body.
//nolint:dupl
func (c *RestClient) Get(ctx context.Context, path string) (responseBody json.RawMessage, err error) {
	ctx, cancel := context.WithTimeout(ctx, *c.config.Timeout)
	defer cancel()

	request, err := c.newRequest(ctx, http.MethodGet, path, json.RawMessage{})
	if err != nil {
		return nil, kentikErrors.New(kentikErrors.InvalidRequest, fmt.Sprintf("new request: %v", err))
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, kentikErrors.KentikErrorFromHTTP(response, fmt.Errorf("do request: %w", err))
	}

	defer func() {
		cErr := response.Body.Close()
		if err == nil && cErr != nil {
			err = fmt.Errorf("close response body: %v", cErr)
		}
	}()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %v", err)
	}

	return body, errorFromResponseStatus(response, string(body))
}

// Post sends POST request to the API and returns raw response body
//nolint:dupl
func (c *RestClient) Post(
	ctx context.Context, path string, payload json.RawMessage,
) (responseBody json.RawMessage, err error) {
	ctx, cancel := context.WithTimeout(ctx, *c.config.Timeout)
	defer cancel()

	request, err := c.newRequest(ctx, http.MethodPost, path, payload)
	if err != nil {
		return nil, kentikErrors.New(kentikErrors.InvalidRequest, fmt.Sprintf("new request: %v", err))
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, kentikErrors.KentikErrorFromHTTP(response, fmt.Errorf("do request: %w", err))
	}
	defer func() {
		cErr := response.Body.Close()
		if err == nil && cErr != nil {
			err = fmt.Errorf("close response body: %v", cErr)
		}
	}()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %v", err)
	}

	return body, errorFromResponseStatus(response, string(body))
}

// Put sends PUT request to the API and returns raw response body
//nolint:dupl
func (c *RestClient) Put(
	ctx context.Context, path string, payload json.RawMessage,
) (responseBody json.RawMessage, err error) {
	ctx, cancel := context.WithTimeout(ctx, *c.config.Timeout)
	defer cancel()

	request, err := c.newRequest(ctx, http.MethodPut, path, payload)
	if err != nil {
		return nil, kentikErrors.New(kentikErrors.InvalidRequest, fmt.Sprintf("new request: %v", err))
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, kentikErrors.KentikErrorFromHTTP(response, fmt.Errorf("do request: %w", err))
	}
	defer func() {
		cErr := response.Body.Close()
		if err == nil && cErr != nil {
			err = fmt.Errorf("close response body: %v", cErr)
		}
	}()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %v", err)
	}

	return body, errorFromResponseStatus(response, string(body))
}

// Delete sends DELETE request to the API and returns raw response body.
//nolint:dupl
func (c *RestClient) Delete(ctx context.Context, path string) (responseBody json.RawMessage, err error) {
	ctx, cancel := context.WithTimeout(ctx, *c.config.Timeout)
	defer cancel()

	request, err := c.newRequest(ctx, http.MethodDelete, path, json.RawMessage{})
	if err != nil {
		return nil, kentikErrors.New(kentikErrors.InvalidRequest, fmt.Sprintf("new request: %v", err))
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, kentikErrors.KentikErrorFromHTTP(response, fmt.Errorf("do request: %w", err))
	}

	defer func() {
		cErr := response.Body.Close()
		if err == nil && cErr != nil {
			err = fmt.Errorf("close response body: %v", cErr)
		}
	}()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %v", err)
	}

	return body, errorFromResponseStatus(response, string(body))
}

func errorFromResponseStatus(r *http.Response, responseBody string) error {
	if r.StatusCode >= http.StatusBadRequest {
		return kentikErrors.KentikErrorFromHTTP(r,
			fmt.Errorf("API response error, status: %v, response body: %v", r.Status, responseBody))
	}
	return nil
}

func (c *RestClient) newRequest(ctx context.Context, method string, path string, payload json.RawMessage,
) (*retryablehttp.Request, error) {
	request, err := http.NewRequestWithContext(ctx, method, c.makeFullURL(path), bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	request.Header.Set(authEmailKey, c.config.AuthEmail)
	request.Header.Set(authAPITokenKey, c.config.AuthToken)
	request.Header.Set("Content-Type", "application/json")

	return retryablehttp.FromRequest(request)
}

func (c *RestClient) makeFullURL(path string) string {
	return c.config.APIURL + path
}
