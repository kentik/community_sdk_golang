package connection

import (
	"context"
	"encoding/json"
)

type Transport interface {
	Get(ctx context.Context, path string) (responseBody json.RawMessage, err error)
	Post(ctx context.Context, path string, payload json.RawMessage) (responseBody json.RawMessage, err error)
	Put(ctx context.Context, path string, payload json.RawMessage) (responseBody json.RawMessage, err error)
	Delete(ctx context.Context, path string) (responseBody json.RawMessage, err error)
}
