package kentikapi

import (
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/cloudexport"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/synthetics"
)

const CLOUDEXPORT_API_URL = "https://cloudexports.api.kentik.com"
const STYNTHETICS_API_URL = "https://synthetics.api.kentik.com"

const authEmailKey = "X-CH-Auth-Email"
const authAPITokenKey = "X-CH-Auth-API-Token"

// Client is the root object for manipulating all the Kentik API resources.
type Client struct {
	// cloudexport
	CloudExportAdminServiceApi *cloudexport.CloudExportAdminServiceApiService

	// synthetics
	SyntheticsAdminServiceApi *synthetics.SyntheticsAdminServiceApiService
	SyntheticsDataServiceApi  *synthetics.SyntheticsDataServiceApiService
}

// NewClient creates kentikapi client with provided config
func NewClient(c Config) *Client {
	if c.CloudExportAPIURL == "" {
		c.CloudExportAPIURL = CLOUDEXPORT_API_URL
	}

	if c.SyntheticsAPIURL == "" {
		c.SyntheticsAPIURL = STYNTHETICS_API_URL
	}

	cloudexportClient := cloudexport.NewAPIClient(makeCloudExportConfig(c))
	syntheticsClient := synthetics.NewAPIClient(makeSyntheticsConfig(c))

	return &Client{
		CloudExportAdminServiceApi: cloudexportClient.CloudExportAdminServiceApi,
		SyntheticsAdminServiceApi:  syntheticsClient.SyntheticsAdminServiceApi,
		SyntheticsDataServiceApi:   syntheticsClient.SyntheticsDataServiceApi,
	}
}

func makeCloudExportConfig(c Config) *cloudexport.Configuration {
	cfg := cloudexport.NewConfiguration()

	// setup authorization
	cfg.DefaultHeader[authEmailKey] = c.AuthEmail
	cfg.DefaultHeader[authAPITokenKey] = c.AuthToken

	// setup target API server
	cfg.Servers[0].URL = c.CloudExportAPIURL
	cfg.Servers[0].Description = "Kentik CloudExport server"

	cfg.Debug = c.Debug
	return cfg
}

func makeSyntheticsConfig(c Config) *synthetics.Configuration {
	cfg := synthetics.NewConfiguration()

	// setup authorization
	cfg.DefaultHeader[authEmailKey] = c.AuthEmail
	cfg.DefaultHeader[authAPITokenKey] = c.AuthToken

	// setup target API server
	cfg.Servers[0].URL = c.SyntheticsAPIURL
	cfg.Servers[0].Description = "Kentik Synthetics server"

	cfg.Debug = c.Debug
	return cfg
}

// Config holds configuration of the client.
type Config struct {
	CloudExportAPIURL string // if blank, default api url will be used
	SyntheticsAPIURL  string // if blank, default api url will be used
	AuthEmail         string
	AuthToken         string
	// Debug logs for requests and responses
	Debug             bool
}
