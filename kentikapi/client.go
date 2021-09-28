package kentikapi

import (
	"time"

	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/httputil"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/resources"
)

// Public constants.
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
	// Batch
	Alerting *resources.AlertingAPI

	config Config
}

// Config holds configuration of the client.
type Config struct {
	// APIURL defaults to "https://api.kentik.com/api/v5"
	APIURL    string
	AuthEmail string
	AuthToken string
	RetryCfg  httputil.RetryConfig
	// Timeout specifies a limit of a total time of a single client call, including redirects and retries.
	// A Timeout of zero means no timeout. Default: 100 seconds.
	Timeout *time.Duration
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
		RetryCfg:  c.RetryCfg,
		Timeout:   c.Timeout,
	})
	return &Client{
		Users:              resources.NewUsersAPI(rc),
		Devices:            resources.NewDevicesAPI(rc),
		DeviceLabels:       resources.NewDeviceLabelsAPI(rc),
		Sites:              resources.NewSitesAPI(rc),
		Tags:               resources.NewTagsAPI(rc),
		SavedFilters:       resources.NewSavedFiltersAPI(rc),
		CustomDimensions:   resources.NewCustomDimensionsAPI(rc),
		CustomApplications: resources.NewCustomApplicationsAPI(rc),
		Query:              resources.NewQueryAPI(rc),
		MyKentikPortal:     resources.NewMyKentikPortalAPI(rc),
		Plans:              resources.NewPlansAPI(rc),
		// Batch
		Alerting: resources.NewAlertingAPI(rc),
		config:   c,
	}
}
