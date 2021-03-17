package kentikapi

import "github.com/kentik/community_sdk_golang/apiv6/kentikapi/cloudexport"

// Client is the root object for manipulating all the Kentik API resources.
type Client = cloudexport.APIClient

// NewClient creates kentikapi client with provided credentials
func NewClient(authEmail, authToken string) *cloudexport.APIClient {
	cfg := makeKentikConfig(authEmail, authToken)
	return cloudexport.NewAPIClient(cfg)
}

func makeKentikConfig(email, token string) *cloudexport.Configuration {
	const authEmailKey = "X-CH-Auth-Email"
	const authAPITokenKey = "X-CH-Auth-API-Token"

	cfg := cloudexport.NewConfiguration()

	// setup authorization
	cfg.DefaultHeader[authEmailKey] = email
	cfg.DefaultHeader[authAPITokenKey] = token

	// setup target API server
	cfg.Servers[0].URL = "https://cloudexports.api.kentik.com"
	cfg.Servers[0].Description = "Kentik CloudExport server"

	return cfg
}
