package kentikapi

import (
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_resources"
)

// Public constants.
const (
	APIURLUS = "https://api.kentik.com/api/v5"
	APIURLEU = "https://api.kentik.eu/api/v5"
)

// Client is the root object for manipulating all the Kentik API resources.
type Client struct {
	UsersAPI   *api_resources.UsersAPI
	DevicesAPI *api_resources.DevicesAPI
	config     Config
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
	rc := api_connection.NewRestClient(api_connection.RestClientConfig{
		APIURL:    c.APIURL,
		AuthEmail: c.AuthEmail,
		AuthToken: c.AuthToken,
	})
	return &Client{
		UsersAPI:   api_resources.NewUsersAPI(rc),
		DevicesAPI: api_resources.NewDevicesAPI(rc),
		config:     c,
	}
}