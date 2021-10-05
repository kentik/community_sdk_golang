package kentikapi

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	grpccloudesxport "github.com/kentik/api-schema-public/gen/go/kentik/cloud_export/v202101beta1"
	grpcsynthetics "github.com/kentik/api-schema-public/gen/go/kentik/synthetics/v202101beta1"
	"github.com/kentik/community_sdk_golang/kentikapi/cloudexport"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/httputil"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/resources"
	"github.com/kentik/community_sdk_golang/kentikapi/synthetics"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

//nolint:gosec
const (
	authAPITokenKey = "X-CH-Auth-API-Token"
	authEmailKey    = "X-CH-Auth-Email"
)

// Kentik API URLs.
const (
	APIURLUS                = "https://api.kentik.com/api/v5"
	APIURLEU                = "https://api.kentik.eu/api/v5"
	cloudExportAPIURL       = "https://cloudexports.api.kentik.com"
	syntheticsAPIURL        = "https://synthetics.api.kentik.com"
	syntheticsGRPCHostPort  = "synthetics.api.kentik.com:443"
	cloudExportGRPCHostPort = "cloudexport.api.kentik.com:443"
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

	// CloudExportAdminServiceAPI, SyntheticsAdminServiceAPI and SyntheticsDataServiceAPI are http clients
	// for Kentik API Cloud Export and Synthetics services.
	// They will be deprecated and using gRPC client is recommended
	CloudExportAdminServiceAPI *cloudexport.CloudExportAdminServiceApiService

	SyntheticsAdminServiceAPI *synthetics.SyntheticsAdminServiceApiService
	SyntheticsDataServiceAPI  *synthetics.SyntheticsDataServiceApiService

	// CloudExportAdmin, SyntheticsAdmin and SyntheticsData are gRPC clients
	// for Kentik API Cloud Export and Synthetics services.
	CloudExportAdmin grpccloudesxport.CloudExportAdminServiceClient
	SyntheticsAdmin  grpcsynthetics.SyntheticsAdminServiceClient
	SyntheticsData   grpcsynthetics.SyntheticsDataServiceClient

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
	// SyntheticsGRPCHostPort defaults to "synthetics.api.kentik.com:443".
	SyntheticsGRPCHostPort string
	// CloudExportGRPCHostPort defaults to "cloudexport.api.kentik.com:443".
	CloudExportGRPCHostPort string
	AuthEmail               string
	AuthToken               string
	// RetryCfg has no effect on CloudExportAdmin, SyntheticsAdmin and SyntheticsData.
	RetryCfg RetryConfig

	// LogPayloads enables logging of request and response payloads to Cloud Export and Synthetics APIs.
	// LogPayloads has no effect on CloudExportAdmin, SyntheticsAdmin and SyntheticsData.
	LogPayloads bool
	// Timeout specifies a limit of a total time of a single client call, including redirects and retries.
	// A Timeout of zero means no timeout. Currently it works only for v5 Admin APIs (e.g. users, devices).
	// Timeout has no effect on CloudExportAdmin, SyntheticsAdmin and SyntheticsData.
	// Default: 100 seconds.
	Timeout *time.Duration
	// DisableTLS disables TLS for client connections.
	// It has effect only on gRPC services: CloudExportAdmin, SyntheticsAdmin and SyntheticsData.
	DisableTLS bool
}

type RetryConfig = httputil.RetryConfig

// NewClient creates a new Kentik API client.
func NewClient(c Config) (*Client, error) {
	if c.APIURL == "" {
		c.APIURL = APIURLUS
	}

	if c.CloudExportAPIURL == "" {
		c.CloudExportAPIURL = cloudExportAPIURL
	}

	if c.SyntheticsAPIURL == "" {
		c.SyntheticsAPIURL = syntheticsAPIURL
	}

	if c.CloudExportGRPCHostPort == "" {
		c.CloudExportGRPCHostPort = cloudExportGRPCHostPort
	}

	if c.SyntheticsGRPCHostPort == "" {
		c.SyntheticsGRPCHostPort = syntheticsGRPCHostPort
	}

	cloudexportClient := cloudexport.NewAPIClient(makeCloudExportConfig(c))
	syntheticsClient := synthetics.NewAPIClient(makeSyntheticsConfig(c))

	syntheticsConnection, err := c.makeConnForGRPC(c.SyntheticsGRPCHostPort)
	if err != nil {
		return nil, fmt.Errorf("grpc synthetics connection: %v", err)
	}
	cloudExportConnection, err := c.makeConnForGRPC(c.CloudExportGRPCHostPort)
	if err != nil {
		return nil, fmt.Errorf("grpc cloud export connection: %v", err)
	}

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
		SyntheticsAdmin:            grpcsynthetics.NewSyntheticsAdminServiceClient(syntheticsConnection),
		SyntheticsData:             grpcsynthetics.NewSyntheticsDataServiceClient(syntheticsConnection),
		CloudExportAdmin:           grpccloudesxport.NewCloudExportAdminServiceClient(cloudExportConnection),
		config:                     c,
	}, nil
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

func (c Config) makeConnForGRPC(hostPort string) (grpc.ClientConnInterface, error) {
	return grpc.Dial(
		hostPort,
		c.makeTLSOption(),
		grpc.WithUnaryInterceptor(c.makeAuthInterceptor()),
	)
}

func (c Config) makeTLSOption() grpc.DialOption {
	if c.DisableTLS {
		return grpc.WithInsecure()
	}
	return grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
		MinVersion: tls.VersionTLS13,
	}))
}

func (c Config) makeAuthInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{},
		cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption,
	) error {
		ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs(
			authEmailKey, c.AuthEmail,
			authAPITokenKey, c.AuthToken,
		))
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
