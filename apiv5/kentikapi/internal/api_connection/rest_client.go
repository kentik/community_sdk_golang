package api_connection

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const (
	authEmailKey    = "X-CH-Auth-Email"
	authAPITokenKey = "X-CH-Auth-API-Token"
)

type restClient struct {
	config     RestClientConfig
	httpClient *http.Client
}

type RestClientConfig struct {
	APIURL    string
	AuthEmail string
	AuthToken string
}

func NewRestClient(c RestClientConfig) *restClient {
	return &restClient{
		config: c,
		httpClient: &http.Client{
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
		},
	}
}

// Get sends GET request to the API and returns raw response body.
func (c *restClient) Get(ctx context.Context, path string) (responseBody json.RawMessage, err error) {
	request, err := c.newRequest(ctx, http.MethodGet, path, json.RawMessage{})
	if err != nil {
		return nil, fmt.Errorf("new request: %v", err)
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("do request: %v", err)
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
func (c *restClient) Post(ctx context.Context, path string, payload json.RawMessage) (responseBody json.RawMessage, err error) {
	request, err := c.newRequest(ctx, http.MethodPost, path, payload)
	if err != nil {
		return nil, fmt.Errorf("new request: %v", err)
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("do request: %v", err)
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
func (c *restClient) Put(ctx context.Context, path string, payload json.RawMessage) (responseBody json.RawMessage, err error) {
	request, err := c.newRequest(ctx, http.MethodPut, path, payload)
	if err != nil {
		return nil, fmt.Errorf("new request: %v", err)
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("do request: %v", err)
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

// Delete sends DELETE request to the API and returns raw response body
func (c *restClient) Delete(ctx context.Context, path string) (responseBody json.RawMessage, err error) {
	request, err := c.newRequest(ctx, http.MethodDelete, path, json.RawMessage{})
	if err != nil {
		return nil, fmt.Errorf("new request: %v", err)
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("do request: %v", err)
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
	// TODO(dfurman): return more specific errors
	if r.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("API response error, status: %v, response body: %v", r.Status, responseBody)
	}
	return nil
}

func (c *restClient) newRequest(ctx context.Context, method string, path string, payload json.RawMessage) (*http.Request, error) {
	request, err := http.NewRequestWithContext(ctx, method, c.makeFullURL(path), bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	request.Header.Set(authEmailKey, c.config.AuthEmail)
	request.Header.Set(authAPITokenKey, c.config.AuthToken)
	request.Header.Set("Content-Type", "application/json")

	return request, nil
}

func (c *restClient) makeFullURL(path string) string {
	return c.config.APIURL + path
}
