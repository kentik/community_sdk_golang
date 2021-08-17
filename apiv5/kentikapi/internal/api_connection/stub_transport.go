package api_connection

import (
	"context"
	"encoding/json"
	"net/http"
)

// StubTransport provides Transport interface with preconfigured ResponseBody
type StubTransport struct {
	ResponseBody string

	RequestMethod string
	RequestPath   string
	RequestBody   string
}

func (t *StubTransport) Get(ctx context.Context, path string) (json.RawMessage, error) {
	t.RequestMethod = http.MethodGet
	t.RequestPath = path
	return json.RawMessage(t.ResponseBody), nil
}

func (t *StubTransport) Post(ctx context.Context, path string, payload json.RawMessage) (json.RawMessage, error) {
	t.RequestMethod = http.MethodPost
	t.RequestPath = path
	t.RequestBody = string(payload)
	return json.RawMessage(t.ResponseBody), nil
}

func (t *StubTransport) Put(ctx context.Context, path string, payload json.RawMessage) (json.RawMessage, error) {
	t.RequestMethod = http.MethodPut
	t.RequestPath = path
	t.RequestBody = string(payload)
	return json.RawMessage(t.ResponseBody), nil
}

func (t *StubTransport) Delete(ctx context.Context, path string) (responseBody json.RawMessage, err error) {
	t.RequestMethod = http.MethodDelete
	t.RequestPath = path
	return json.RawMessage(t.ResponseBody), nil
}
