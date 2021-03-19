package kentikapi

import "github.com/kentik/community_sdk_golang/apiv6/kentikapi/cloudexport"

// APIURLUS specifies the US cloud export http server
const APIURLUS = "https://cloudexports.api.kentik.com"

// Client is the root object for manipulating all the Kentik API resources.
type Client = cloudexport.APIClient

// NewClient creates kentikapi client with provided config
func NewClient(c Config) *cloudexport.APIClient {
	if c.APIURL == "" {
		c.APIURL = APIURLUS
	}
	cfg := makeKentikConfig(c)
	return cloudexport.NewAPIClient(cfg)
}

func makeKentikConfig(c Config) *cloudexport.Configuration {
	const authEmailKey = "X-CH-Auth-Email"
	const authAPITokenKey = "X-CH-Auth-API-Token"

	cfg := cloudexport.NewConfiguration()

	// setup authorization
	cfg.DefaultHeader[authEmailKey] = c.AuthEmail
	cfg.DefaultHeader[authAPITokenKey] = c.AuthToken

	// setup target API server
	cfg.Servers[0].URL = c.APIURL
	cfg.Servers[0].Description = "Kentik CloudExport server"

	return cfg
}

// Config holds configuration of the client.
type Config struct {
	APIURL    string // if blank, default api url will be used
	AuthEmail string
	AuthToken string
}
