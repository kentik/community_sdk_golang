package api_connection

import (
	"context"
	"encoding/json"
)

// StubTransport provides Transport interface with preconfigured ResponseBody
type StubTransport struct {
	ResponseBody string
}

func (t StubTransport) Get(ctx context.Context, path string) (json.RawMessage, error) {
	return json.RawMessage(t.ResponseBody), nil
}