package kentikapi

import (
	"time"

	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/cloudexport"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/httputil"
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

// Config holds configuration of the Client.
// See httputil.NewRetryingClient for default retry policy description.
type Config struct {
	CloudExportAPIURL string // if blank, default api url will be used
	SyntheticsAPIURL  string // if blank, default api url will be used
	AuthEmail         string
	AuthToken         string
	// RetryMax is a maximum number of request retries. Default: 4.
	RetryMax *int
	// RetryWaitMin is a minimum time to wait before request retry. Default: 1 second.
	RetryWaitMin *time.Duration
	// RetryWaitMax is a maximum time to wait before request retry. Default. 30 seconds.
	RetryWaitMax *time.Duration
	// LogPayloads enables logging of request and response payloads.
	LogPayloads bool
}

// NewClient creates Kentik API client with provided Config.
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

	cfg.HTTPClient = httputil.NewRetryingClient(makeRetryingClientConfig(c))
	cfg.Debug = c.LogPayloads
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

	cfg.HTTPClient = httputil.NewRetryingClient(makeRetryingClientConfig(c))
	cfg.Debug = c.LogPayloads
	return cfg
}

func makeRetryingClientConfig(c Config) httputil.ClientConfig {
	return httputil.ClientConfig{
		RetryMax:     c.RetryMax,
		RetryWaitMin: c.RetryWaitMin,
		RetryWaitMax: c.RetryWaitMax,
	}
}
