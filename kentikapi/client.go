package kentikapi

import (
	"time"

	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/cloudexport"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/httputil"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/synthetics"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/resources"
)

// Public constants.
//nolint:gosec
const (
	authAPITokenKey   = "X-CH-Auth-API-Token"
	authEmailKey      = "X-CH-Auth-Email"
	cloudExportAPIURL = "https://cloudexports.api.kentik.com"
	syntheticsAPIURL  = "https://synthetics.api.kentik.com"
)

// Kentik API URLs.
const (
	APIURLUS = "https://api.kentik.com/api/v5"
	APIURLEU = "https://api.kentik.eu/api/v5"
)

// Client is the root object for manipulating all the Kentik API resources.
type Client struct {
	Users              *resources.UsersAPI
	Devices            *resources.DevicesAPI
	DeviceLabels       *resources.DeviceLabelsAPI
	Sites              *resources.SitesAPI
	Tags               *resources.TagsAPI
	SavedFilters       *resources.SavedFiltersAPI
	CustomDimensions   *resources.CustomDimensionsAPI
	CustomApplications *resources.CustomApplicationsAPI
	Query              *resources.QueryAPI
	MyKentikPortal     *resources.MyKentikPortalAPI
	Plans              *resources.PlansAPI
	Alerting           *resources.AlertingAPI

	CloudExportAdminServiceAPI *cloudexport.CloudExportAdminServiceApiService

	SyntheticsAdminServiceAPI *synthetics.SyntheticsAdminServiceApiService
	SyntheticsDataServiceAPI  *synthetics.SyntheticsDataServiceApiService

	config Config
}

// Config holds configuration of the client.
type Config struct {
	// APIURL defaults to "https://api.kentik.com/api/v5"
	APIURL string
	// CloudExportAPIURL defaults to "https://cloudexports.api.kentik.com".
	CloudExportAPIURL string
	// SyntheticsAPIURL defaults to "https://synthetics.api.kentik.com".
	SyntheticsAPIURL string
	AuthEmail        string
	AuthToken        string
	RetryCfg         RetryConfig

	// LogPayloads enables logging of request and response payloads to Cloud Export and Synthetics APIs.
	LogPayloads bool
	// Timeout specifies a limit of a total time of a single client call, including redirects and retries.
	// A Timeout of zero means no timeout. Currently it works only for v5 Admin APIs (e.g. users, devices).
	// Default: 100 seconds.
	Timeout *time.Duration
}

type RetryConfig = httputil.RetryConfig

// NewClient creates a new Kentik API client.
func NewClient(c Config) *Client {
	if c.APIURL == "" {
		c.APIURL = APIURLUS
	}

	if c.CloudExportAPIURL == "" {
		c.CloudExportAPIURL = cloudExportAPIURL
	}

	if c.SyntheticsAPIURL == "" {
		c.SyntheticsAPIURL = syntheticsAPIURL
	}

	cloudexportClient := cloudexport.NewAPIClient(makeCloudExportConfig(c))
	syntheticsClient := synthetics.NewAPIClient(makeSyntheticsConfig(c))

	rc := api_connection.NewRestClient(api_connection.RestClientConfig{
		APIURL:    c.APIURL,
		AuthEmail: c.AuthEmail,
		AuthToken: c.AuthToken,
		RetryCfg:  c.RetryCfg,
		Timeout:   c.Timeout,
	})
	return &Client{
		Users:                      resources.NewUsersAPI(rc),
		Devices:                    resources.NewDevicesAPI(rc),
		DeviceLabels:               resources.NewDeviceLabelsAPI(rc),
		Sites:                      resources.NewSitesAPI(rc),
		Tags:                       resources.NewTagsAPI(rc),
		SavedFilters:               resources.NewSavedFiltersAPI(rc),
		CustomDimensions:           resources.NewCustomDimensionsAPI(rc),
		CustomApplications:         resources.NewCustomApplicationsAPI(rc),
		Query:                      resources.NewQueryAPI(rc),
		MyKentikPortal:             resources.NewMyKentikPortalAPI(rc),
		Plans:                      resources.NewPlansAPI(rc),
		Alerting:                   resources.NewAlertingAPI(rc),
		CloudExportAdminServiceAPI: cloudexportClient.CloudExportAdminServiceApi,
		SyntheticsAdminServiceAPI:  syntheticsClient.SyntheticsAdminServiceApi,
		SyntheticsDataServiceAPI:   syntheticsClient.SyntheticsDataServiceApi,
		config:                     c,
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
