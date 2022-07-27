package kentikapi

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/url"
	"time"

	"github.com/AlekSi/pointer"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/cloud"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/httputil"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/resources"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/synthetics"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

//nolint:gosec
const (
	authAPITokenKey    = "X-CH-Auth-API-Token"
	authEmailKey       = "X-CH-Auth-Email"
	defaultTimeout     = 100 * time.Second
	defaultMaxAttempts = 4
	defaultMinDelay    = 1 * time.Second
	defaultMaxDelay    = 30 * time.Second
)

// Kentik API URLs.
const (
	APIURLUS = "https://api.kentik.com"
	APIURLEU = "https://api.kentik.eu"
)

// Client is the root object for manipulating all the Kentik API resources.
type Client struct {
	Alerting           *resources.AlertingAPI
	Cloud              *cloud.API
	CustomApplications *resources.CustomApplicationsAPI
	CustomDimensions   *resources.CustomDimensionsAPI
	DeviceLabels       *resources.DeviceLabelsAPI
	Devices            *resources.DevicesAPI
	MyKentikPortal     *resources.MyKentikPortalAPI
	Plans              *resources.PlansAPI
	Query              *resources.QueryAPI
	SavedFilters       *resources.SavedFiltersAPI
	Sites              *resources.SitesAPI
	Synthetics         *synthetics.API
	Tags               *resources.TagsAPI
	Users              *resources.UsersAPI

	config config
}

type config struct {
	APIURL    string
	AuthEmail string
	AuthToken string
	RetryCfg  RetryConfig

	LogPayloads bool
	Timeout     time.Duration
}

type RetryConfig = httputil.RetryConfig

// ClientOption defines an option for a Client.
type ClientOption func(*config)

// WithAPIURL sets the Kentik API URL. Default: "https://api.kentik.com".
func WithAPIURL(apiURL string) ClientOption {
	return func(c *config) {
		c.APIURL = apiURL
	}
}

// WithCredentials specifies authentication email and token for the Client.
func WithCredentials(authEmail, authToken string) ClientOption {
	return func(c *config) {
		c.AuthEmail = authEmail
		c.AuthToken = authToken
	}
}

// WithTimeout specifies a limit of a total time of a single client call, including redirects and retries.
// Default: 100 seconds.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *config) {
		if timeout != 0 {
			c.Timeout = timeout
		}
	}
}

// WithRetryMaxAttempts specifies maximum number of request retry attempts. Set to 0 to disable retrying. Default: 4.
func WithRetryMaxAttempts(maxAttempts uint) ClientOption {
	return func(c *config) {
		c.RetryCfg.MaxAttempts = &maxAttempts
	}
}

// WithRetryMinDelay specifies a minimum delay before request retry. Default: 1 second.
func WithRetryMinDelay(minDelay time.Duration) ClientOption {
	return func(c *config) {
		c.RetryCfg.MinDelay = &minDelay
	}
}

// WithRetryMaxDelay specifies a maximum delay before request retry. Default: 30 seconds.
func WithRetryMaxDelay(maxDelay time.Duration) ClientOption {
	return func(c *config) {
		c.RetryCfg.MaxDelay = &maxDelay
	}
}

// WithLogPayloads enables logging of request and response payloads.
func WithLogPayloads() ClientOption {
	return func(c *config) {
		c.LogPayloads = true
	}
}

// NewClient creates a new Kentik API client.
func NewClient(opts ...ClientOption) (*Client, error) {
	c := config{
		APIURL:  APIURLUS,
		Timeout: defaultTimeout,
		RetryCfg: RetryConfig{
			MaxAttempts: pointer.ToUint(defaultMaxAttempts),
			MinDelay:    pointer.ToDuration(defaultMinDelay),
			MaxDelay:    pointer.ToDuration(defaultMaxDelay),
		},
	}
	for _, opt := range opts {
		opt(&c)
	}
	apiV5URL, err := makeAPIV5URL(c.APIURL)
	if err != nil {
		return nil, fmt.Errorf("make API v5 URL: %v", err)
	}

	grpcConnection, err := makeConnForGRPC(c)
	if err != nil {
		return nil, fmt.Errorf("grpc connection: %v", err)
	}
	rc := api_connection.NewRestClient(api_connection.RestClientConfig{
		APIURL:    apiV5URL,
		AuthEmail: c.AuthEmail,
		AuthToken: c.AuthToken,
		RetryCfg:  c.RetryCfg,
		Timeout:   &c.Timeout,
	})
	return &Client{
		Alerting:           resources.NewAlertingAPI(rc, c.LogPayloads),
		Cloud:              cloud.NewAPI(grpcConnection),
		CustomApplications: resources.NewCustomApplicationsAPI(rc, c.LogPayloads),
		CustomDimensions:   resources.NewCustomDimensionsAPI(rc, c.LogPayloads),
		DeviceLabels:       resources.NewDeviceLabelsAPI(rc, c.LogPayloads),
		Devices:            resources.NewDevicesAPI(rc, c.LogPayloads),
		MyKentikPortal:     resources.NewMyKentikPortalAPI(rc, c.LogPayloads),
		Plans:              resources.NewPlansAPI(rc, c.LogPayloads),
		Query:              resources.NewQueryAPI(rc, c.LogPayloads),
		SavedFilters:       resources.NewSavedFiltersAPI(rc, c.LogPayloads),
		Sites:              resources.NewSitesAPI(rc, c.LogPayloads),
		Synthetics:         synthetics.NewAPI(grpcConnection),
		Tags:               resources.NewTagsAPI(rc, c.LogPayloads),
		Users:              resources.NewUsersAPI(rc, c.LogPayloads),

		config: c,
	}, nil
}

func makeAPIV5URL(apiURL string) (string, error) {
	u, err := url.Parse(apiURL)
	if err != nil {
		return "", err
	}
	u.Path += "/api/v5"
	return u.String(), nil
}

func makeConnForGRPC(c config) (grpc.ClientConnInterface, error) {
	grpcHostPort, tlsEnabled, err := makeGRPCHostPort(c.APIURL)
	if err != nil {
		return nil, err
	}
	return grpc.Dial(
		grpcHostPort,
		makeTLSOption(tlsEnabled),
		grpc.WithUnaryInterceptor(
			grpcmiddleware.ChainUnaryClient(
				makeTimeoutInterceptor(c),
				makeLoggerInterceptor(c),
				makeAuthInterceptor(c),
				makeRetryInterceptor(c),
			),
		),
	)
}

func makeGRPCHostPort(apiURL string) (grpcHostPort string, tlsEnabled bool, err error) {
	tlsEnabled = false
	u, err := url.Parse(apiURL)
	if err != nil {
		return "", false, err
	}
	grpcHostPort = u.Host
	if u.Scheme == "https" {
		tlsEnabled = true
		if u.Port() == "" {
			grpcHostPort += ":443"
		}
	}
	if u.Scheme == "http" && u.Port() == "" {
		grpcHostPort += ":80"
	}
	hostIP := net.ParseIP(u.Hostname())
	if hostIP != nil {
		return grpcHostPort, tlsEnabled, nil
	}
	return "grpc." + grpcHostPort, tlsEnabled, nil
}

func makeTLSOption(tlsEnabled bool) grpc.DialOption {
	if !tlsEnabled {
		return grpc.WithTransportCredentials(insecure.NewCredentials())
	}
	return grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
		MinVersion: tls.VersionTLS13,
	}))
}

func makeTimeoutInterceptor(c config) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{},
		cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption,
	) error {
		ctx, cancel := context.WithTimeout(ctx, c.Timeout)
		defer cancel()
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func makeAuthInterceptor(c config) grpc.UnaryClientInterceptor {
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

func makeRetryInterceptor(c config) grpc.UnaryClientInterceptor {
	return grpcretry.UnaryClientInterceptor(
		grpcretry.WithBackoff(grpcretry.BackoffExponential(*c.RetryCfg.MinDelay)),
		grpcretry.WithCodes(
			codes.Unavailable,
			codes.ResourceExhausted,
			codes.FailedPrecondition,
			codes.Aborted,
		),
		// grpcretry.WithMax specifies the number of total requests sent and not retries
		grpcretry.WithMax(*c.RetryCfg.MaxAttempts+1),
	)
}

func makeLoggerInterceptor(c config) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{},
		cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption,
	) error {
		if c.LogPayloads {
			log.Printf(
				"Kentik API request: target=%s method=%s payload=%v",
				cc.Target(), method, sanitizePayload(fmt.Sprint(req)),
			)
		}
		err := invoker(ctx, method, req, reply, cc, opts...)
		if c.LogPayloads {
			log.Printf(
				"Kentik API response: target=%s method=%s payload=%v error=%v",
				cc.Target(), method, sanitizePayload(fmt.Sprint(reply)), err,
			)
		}
		return err
	}
}

func sanitizePayload(payload string) string {
	if payload == "" {
		return "<empty>"
	}

	const maxLoggedPayloadSize = 10000
	if len(payload) > maxLoggedPayloadSize {
		return fmt.Sprintf("<size: %v bytes>", len(payload))
	}
	return payload
}
