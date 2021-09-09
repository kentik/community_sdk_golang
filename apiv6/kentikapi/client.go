package kentikapi

import (
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/cloudexport"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/httputil"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/synthetics"
)

//nolint:gosec
const (
	authAPITokenKey   = "X-CH-Auth-API-Token"
	authEmailKey      = "X-CH-Auth-Email"
	cloudExportAPIURL = "https://cloudexports.api.kentik.com"
	syntheticsAPIURL  = "https://synthetics.api.kentik.com"
)

// Client is the root object for manipulating all the Kentik API resources.
type Client struct {
	// cloudexport
	CloudExportAdminServiceAPI *cloudexport.CloudExportAdminServiceApiService

	// synthetics
	SyntheticsAdminServiceAPI *synthetics.SyntheticsAdminServiceApiService
	SyntheticsDataServiceAPI  *synthetics.SyntheticsDataServiceApiService
}

// Config holds configuration of the Client.
type Config struct {
	// CloudExportAPIURL defaults to "https://cloudexports.api.kentik.com".
	CloudExportAPIURL string
	// SyntheticsAPIURL defaults to "https://synthetics.api.kentik.com".
	SyntheticsAPIURL string
	AuthEmail        string
	AuthToken        string
	RetryCfg         RetryConfig
	// LogPayloads enables logging of request and response api_payloads.
	LogPayloads bool
}

type RetryConfig = httputil.RetryConfig

// NewClient creates Kentik API client with provided Config.
func NewClient(c Config) *Client {
	if c.CloudExportAPIURL == "" {
		c.CloudExportAPIURL = cloudExportAPIURL
	}

	if c.SyntheticsAPIURL == "" {
		c.SyntheticsAPIURL = syntheticsAPIURL
	}

	cloudexportClient := cloudexport.NewAPIClient(makeCloudExportConfig(c))
	syntheticsClient := synthetics.NewAPIClient(makeSyntheticsConfig(c))

	return &Client{
		CloudExportAdminServiceAPI: cloudexportClient.CloudExportAdminServiceApi,
		SyntheticsAdminServiceAPI:  syntheticsClient.SyntheticsAdminServiceApi,
		SyntheticsDataServiceAPI:   syntheticsClient.SyntheticsDataServiceApi,
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

	cfg.HTTPClient = httputil.NewRetryingStdClient(makeRetryingClientConfig(c))
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

	cfg.HTTPClient = httputil.NewRetryingStdClient(makeRetryingClientConfig(c))
	cfg.Debug = c.LogPayloads
	return cfg
}

func makeRetryingClientConfig(c Config) httputil.ClientConfig {
	return httputil.ClientConfig{
		RetryCfg: c.RetryCfg,
	}
}
