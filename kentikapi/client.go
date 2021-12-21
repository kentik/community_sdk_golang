package kentikapi

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	grpccloudesxport "github.com/kentik/api-schema-public/gen/go/kentik/cloud_export/v202101beta1"
	grpcsynthetics "github.com/kentik/api-schema-public/gen/go/kentik/synthetics/v202101beta1"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/httputil"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/resources"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

//nolint:gosec
const (
	authAPITokenKey = "X-CH-Auth-API-Token"
	authEmailKey    = "X-CH-Auth-Email"
	defaultTimeout  = 100 * time.Second
)

// Kentik API URLs.
const (
	APIURLUS            = "https://api.kentik.com/api/v5"
	APIURLEU            = "https://api.kentik.eu/api/v5"
	syntheticsHostPort  = "synthetics.api.kentik.com:443"
	cloudExportHostPort = "cloudexport.api.kentik.com:443"
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
	// SyntheticsHostPort defaults to "synthetics.api.kentik.com:443".
	SyntheticsHostPort string
	// CloudExportHostPort defaults to "cloudexport.api.kentik.com:443".
	CloudExportHostPort string
	AuthEmail           string
	AuthToken           string
	RetryCfg            RetryConfig

	// LogPayloads enables logging of request and response payloads to Cloud Export and Synthetics APIs.
	// LogPayloads has no effect on CloudExportAdmin, SyntheticsAdmin and SyntheticsData.
	LogPayloads bool
	// Timeout specifies a limit of a total time of a single client call, including redirects and retries.
	// A Timeout of zero means no timeout. Currently it works only for v5 Admin APIs (e.g. users, devices).
	// Default: 100 seconds.
	Timeout *time.Duration
	// DisableTLS disables TLS for client connections.
	// It has effect only on gRPC services: CloudExportAdmin, SyntheticsAdmin and SyntheticsData.
	DisableTLS bool
}

type RetryConfig = httputil.RetryConfig

// NewClient creates a new Kentik API client.
func NewClient(c Config) (*Client, error) {
	c.FillDefaults()

	syntheticsConnection, err := c.makeConnForGRPC(c.SyntheticsHostPort)
	if err != nil {
		return nil, fmt.Errorf("grpc synthetics connection: %v", err)
	}
	cloudExportConnection, err := c.makeConnForGRPC(c.CloudExportHostPort)
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
		Alerting:           resources.NewAlertingAPI(rc),
		SyntheticsAdmin:    grpcsynthetics.NewSyntheticsAdminServiceClient(syntheticsConnection),
		SyntheticsData:     grpcsynthetics.NewSyntheticsDataServiceClient(syntheticsConnection),
		CloudExportAdmin:   grpccloudesxport.NewCloudExportAdminServiceClient(cloudExportConnection),
		config:             c,
	}, nil
}

func (c *Config) FillDefaults() {
	if c.APIURL == "" {
		c.APIURL = APIURLUS
	}

	if c.CloudExportHostPort == "" {
		c.CloudExportHostPort = cloudExportHostPort
	}

	if c.SyntheticsHostPort == "" {
		c.SyntheticsHostPort = syntheticsHostPort
	}

	if c.Timeout == nil {
		c.Timeout = durationPtr(defaultTimeout)
	}

	c.RetryCfg.FillDefaults()
}

func (c Config) makeConnForGRPC(hostPort string) (grpc.ClientConnInterface, error) {
	return grpc.Dial(
		hostPort,
		c.makeTLSOption(),
		grpc.WithUnaryInterceptor(
			grpcmiddleware.ChainUnaryClient(
				c.makeTimeoutInterceptor(),
				c.makeAuthInterceptor(),
				c.makeRetryInterceptor(),
			),
		),
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

func (c Config) makeTimeoutInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{},
		cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption,
	) error {
		ctx, cancel := context.WithTimeout(ctx, *c.Timeout)
		defer cancel()
		return invoker(ctx, method, req, reply, cc, opts...)
	}
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

func (c Config) makeRetryInterceptor() grpc.UnaryClientInterceptor {
	return grpcretry.UnaryClientInterceptor(
		grpcretry.WithBackoff(grpcretry.BackoffExponential(*c.RetryCfg.MinDelay)),
		grpcretry.WithCodes(codes.Unavailable),
		// grpcretry.WithMax specifies the number of total requests sent and not retries
		grpcretry.WithMax(*c.RetryCfg.MaxAttempts+1),
	)
}

func durationPtr(v time.Duration) *time.Duration {
	return &v
}
