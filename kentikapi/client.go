package kentikapi

import (
	"context"
	"encoding/json"
)

// Public constants.
const (
	APIURLUS = "https://api.kentik.com/api/v5"
	APIURLEU = "https://api.kentik.eu/api/v5"
)

// Client is the root object for manipulating all the Kentik API resources.
type Client struct {
	UsersAPI usersAPI

	config Config
}

type transport interface {
	Get(ctx context.Context, path string) (responseBody json.RawMessage, err error)
}

// Config holds configuration of the client.
type Config struct {
	// APIURL defaults to "https://api.kentik.com/api/v5"
	APIURL    string
	AuthEmail string
	AuthToken string
}

// NewClient creates a new Kentik API client.
func NewClient(c Config) *Client {
	if c.APIURL == "" {
		c.APIURL = APIURLUS
	}

	rc := newRestClient(restClientConfig{
		APIURL:    c.APIURL,
		AuthEmail: c.AuthEmail,
		AuthToken: c.AuthToken,
	})
	return &Client{
		UsersAPI: usersAPI{transport: rc},
		config:   c,
	}
}
