package kentikapi

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"github.com/AlekSi/pointer"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	grpcsynthetics "github.com/kentik/api-schema-public/gen/go/kentik/synthetics/v202101beta1"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/httputil"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/resources"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
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
	Alerting           *resources.AlertingAPI
	CloudExports       *resources.CloudExportsAPI
	CustomApplications *resources.CustomApplicationsAPI
	CustomDimensions   *resources.CustomDimensionsAPI
	DeviceLabels       *resources.DeviceLabelsAPI
	Devices            *resources.DevicesAPI
	MyKentikPortal     *resources.MyKentikPortalAPI
	Plans              *resources.PlansAPI
	Query              *resources.QueryAPI
	SavedFilters       *resources.SavedFiltersAPI
	Sites              *resources.SitesAPI
	Tags               *resources.TagsAPI
	Users              *resources.UsersAPI

	// SyntheticsAdmin and SyntheticsData are gRPC clients
	// for Kentik API Cloud Export and Synthetics services.
	SyntheticsAdmin grpcsynthetics.SyntheticsAdminServiceClient
	SyntheticsData  grpcsynthetics.SyntheticsDataServiceClient

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
	LogPayloads bool
	// Timeout specifies a limit of a total time of a single client call, including redirects and retries.
	// A Timeout of zero means no timeout.
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

	syntheticsConnection, err := makeConnForGRPC(c, c.SyntheticsHostPort)
	if err != nil {
		return nil, fmt.Errorf("grpc synthetics connection: %v", err)
	}
	cloudExportConnection, err := makeConnForGRPC(c, c.CloudExportHostPort)
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
		Alerting:           resources.NewAlertingAPI(rc),
		CloudExports:       resources.NewCloudExportsAPI(cloudExportConnection),
		CustomApplications: resources.NewCustomApplicationsAPI(rc),
		CustomDimensions:   resources.NewCustomDimensionsAPI(rc),
		DeviceLabels:       resources.NewDeviceLabelsAPI(rc),
		Devices:            resources.NewDevicesAPI(rc),
		MyKentikPortal:     resources.NewMyKentikPortalAPI(rc),
		Plans:              resources.NewPlansAPI(rc),
		Query:              resources.NewQueryAPI(rc),
		SavedFilters:       resources.NewSavedFiltersAPI(rc),
		Sites:              resources.NewSitesAPI(rc),
		Tags:               resources.NewTagsAPI(rc),
		Users:              resources.NewUsersAPI(rc),

		SyntheticsAdmin: grpcsynthetics.NewSyntheticsAdminServiceClient(syntheticsConnection),
		SyntheticsData:  grpcsynthetics.NewSyntheticsDataServiceClient(syntheticsConnection),

		config: c,
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
		c.Timeout = pointer.ToDuration(defaultTimeout)
	}

	c.RetryCfg.FillDefaults()
}

func makeConnForGRPC(c Config, hostPort string) (grpc.ClientConnInterface, error) {
	return grpc.Dial(
		hostPort,
		makeTLSOption(c),
		grpc.WithUnaryInterceptor(
			grpcmiddleware.ChainUnaryClient(
				makeTimeoutInterceptor(c),
				makeAuthInterceptor(c),
				makeRetryInterceptor(c),
				makeLoggerInterceptor(c),
			),
		),
	)
}

func makeTLSOption(c Config) grpc.DialOption {
	if c.DisableTLS {
		return grpc.WithTransportCredentials(insecure.NewCredentials())
	}
	return grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
		MinVersion: tls.VersionTLS13,
	}))
}

func makeTimeoutInterceptor(c Config) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{},
		cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption,
	) error {
		ctx, cancel := context.WithTimeout(ctx, *c.Timeout)
		defer cancel()
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func makeAuthInterceptor(c Config) grpc.UnaryClientInterceptor {
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

func makeRetryInterceptor(c Config) grpc.UnaryClientInterceptor {
	return grpcretry.UnaryClientInterceptor(
		grpcretry.WithBackoff(grpcretry.BackoffExponential(*c.RetryCfg.MinDelay)),
		grpcretry.WithCodes(codes.Unavailable),
		// grpcretry.WithMax specifies the number of total requests sent and not retries
		grpcretry.WithMax(*c.RetryCfg.MaxAttempts+1),
	)
}

func makeLoggerInterceptor(c Config) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{},
		cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption,
	) error {
		if c.LogPayloads {
			log.Printf("Kentik API request - %s - %+v\n", method, req)
		}
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			if c.LogPayloads {
				log.Printf("grpc connection error: %v", err)
			}
			return err
		}
		if c.LogPayloads {
			log.Printf("Kentik API response - %s - %+v\n", method, reply)
		}
		return err
	}
}
