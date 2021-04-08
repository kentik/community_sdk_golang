package kentikapi

import (
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/cloudexport"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/synthetics"
)

// APIURLUS specifies the US cloud export http server
const APIURLUS = "https://cloudexports.api.kentik.com"

const authEmailKey = "X-CH-Auth-Email"
const authAPITokenKey = "X-CH-Auth-API-Token"

// Client is the root object for manipulating all the Kentik API resources.
type Client struct {
	// API Services
	CloudExportAdminServiceApi *cloudexport.CloudExportAdminServiceApiService

	SyntheticsAdminServiceApi *synthetics.SyntheticsAdminServiceApiService
	SyntheticsDataServiceApi  *synthetics.SyntheticsDataServiceApiService
}

// NewClient creates kentikapi client with provided config
func NewClient(c Config) *Client {
	if c.APIURL == "" {
		c.APIURL = APIURLUS
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
	cfg.Servers[0].URL = c.APIURL
	cfg.Servers[0].Description = "Kentik CloudExport server"

	return cfg
}

func makeSyntheticsConfig(c Config) *synthetics.Configuration {
	cfg := synthetics.NewConfiguration()

	// setup authorization
	cfg.DefaultHeader[authEmailKey] = c.AuthEmail
	cfg.DefaultHeader[authAPITokenKey] = c.AuthToken

	// setup target API server
	cfg.Servers[0].URL = c.APIURL
	cfg.Servers[0].Description = "Kentik Synthetics server"

	return cfg
}

// Config holds configuration of the client.
type Config struct {
	APIURL    string // if blank, default api url will be used
	AuthEmail string
	AuthToken string
}
