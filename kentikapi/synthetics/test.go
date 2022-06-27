package synthetics

import (
	"net"
	"net/url"
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

// NewIPTest creates a new IP test with all required fields set.
func NewIPTest(obj IPTestRequiredFields) *Test {
	t := newBasePingTraceTest(obj.BasePingTraceTestRequiredFields)
	t.Type = TestTypeIP
	t.Settings.Definition = &obj.Definition
	t.Settings.Tasks = []TaskType{TaskTypePing, TaskTypeTraceroute}
	return t
}

// NewNetworkGridTest creates a new NetworkGrid test with all required fields set.
func NewNetworkGridTest(obj NetworkGridTestRequiredFields) *Test {
	t := newBasePingTraceTest(obj.BasePingTraceTestRequiredFields)
	t.Type = TestTypeNetworkGrid
	t.Settings.Definition = &obj.Definition
	t.Settings.Tasks = []TaskType{TaskTypePing, TaskTypeTraceroute}
	return t
}

// NewHostnameTest creates a new hostname test with all required fields set.
func NewHostnameTest(obj HostnameTestRequiredFields) *Test {
	t := newBasePingTraceTest(obj.BasePingTraceTestRequiredFields)
	t.Type = TestTypeHostname
	t.Settings.Definition = &obj.Definition
	t.Settings.Tasks = []TaskType{TaskTypePing, TaskTypeTraceroute}
	return t
}

// NewAgentTest creates a new agent test with all required fields set.
func NewAgentTest(obj AgentTestRequiredFields) *Test {
	t := newBasePingTraceTest(obj.BasePingTraceTestRequiredFields)
	t.Type = TestTypeAgent
	t.Settings.Definition = &TestDefinitionAgent{
		Target: obj.Definition.Target,
	}
	t.Settings.Tasks = []TaskType{TaskTypePing, TaskTypeTraceroute}
	return t
}

// NewNetworkMeshTest creates a new network mesh test with all required fields set.
func NewNetworkMeshTest(obj NetworkMeshTestRequiredFields) *Test {
	t := newBasePingTraceTest(obj.BasePingTraceTestRequiredFields)
	t.Type = TestTypeNetworkMesh
	t.Settings.Definition = &TestDefinitionNetworkMesh{}
	t.Settings.Tasks = []TaskType{TaskTypePing, TaskTypeTraceroute}
	return t
}

// NewFlowTest creates a new flow test with all required fields set.
func NewFlowTest(obj FlowTestRequiredFields) *Test {
	t := newBasePingTraceTest(obj.BasePingTraceTestRequiredFields)
	t.Type = TestTypeFlow
	t.Settings.Definition = &TestDefinitionFlow{
		Type:          obj.Definition.Type,
		Target:        obj.Definition.Target,
		Direction:     obj.Definition.Direction,
		InetDirection: obj.Definition.InetDirection,
	}
	t.Settings.Tasks = []TaskType{TaskTypePing, TaskTypeTraceroute}
	return t
}

// NewURLTest creates a new URL test with all required fields set.
func NewURLTest(obj URLTestRequiredFields) *Test {
	t := newBaseTest(obj.BaseTestRequiredFields)
	t.Type = TestTypeURL
	t.Settings.Definition = &TestDefinitionURL{
		Target:  obj.Definition.Target,
		Timeout: obj.Definition.Timeout,
	}
	t.Settings.Tasks = []TaskType{TaskTypeHTTP} // ping and traceroute tasks are optional
	return t
}

// NewPageLoadTest creates a new page load test with all required fields set.
func NewPageLoadTest(obj PageLoadTestRequiredFields) *Test {
	t := newBaseTest(obj.BaseTestRequiredFields)
	t.Type = TestTypePageLoad
	t.Settings.Definition = &TestDefinitionPageLoad{
		Target:  obj.Definition.Target,
		Timeout: obj.Definition.Timeout,
	}
	t.Settings.Tasks = []TaskType{TaskTypePageLoad} // ping and traceroute tasks are optional
	return t
}

// NewDNSTest creates a new DNS test with all required fields set.
func NewDNSTest(obj DNSTestRequiredFields) *Test {
	t := newBaseTest(obj.BaseTestRequiredFields)
	t.Type = TestTypeDNS
	t.Settings.Definition = &obj.Definition
	t.Settings.Tasks = []TaskType{TaskTypeDNS}
	return t
}

// NewDNSGridTest creates a new DNS grid test with all required fields set.
func NewDNSGridTest(obj DNSGridTestRequiredFields) *Test {
	t := newBaseTest(obj.BaseTestRequiredFields)
	t.Type = TestTypeDNSGrid
	t.Settings.Definition = &obj.Definition
	t.Settings.Tasks = []TaskType{TaskTypeDNS}
	return t
}

func newBasePingTraceTest(obj BasePingTraceTestRequiredFields) *Test {
	t := newBaseTest(obj.BaseTestRequiredFields)
	t.Settings.Ping = &PingSettings{
		Timeout:  obj.Ping.Timeout,
		Count:    obj.Ping.Count,
		Protocol: obj.Ping.Protocol,
	}
	t.Settings.Traceroute = &TracerouteSettings{
		Timeout:  obj.Traceroute.Timeout,
		Count:    obj.Traceroute.Count,
		Delay:    obj.Traceroute.Delay,
		Protocol: obj.Traceroute.Protocol,
		Limit:    obj.Traceroute.Limit,
	}
	return t
}

func newBaseTest(obj BaseTestRequiredFields) *Test {
	return &Test{
		Name: obj.Name,
		Settings: TestSettings{
			AgentIDs: obj.AgentIDs,
		},
	}
}

// GetAllTestsResponse model.
type GetAllTestsResponse struct {
	// Tests holds all tests.
	Tests []Test
	// InvalidTestsCount is a number of invalid tests.
	InvalidTestsCount uint32
}

// Test is synthetics test model.
type Test struct {
	// Read-write properties

	// Name is user selected name for the test.
	Name string
	// Type is the specified type of the test. It must be provided on test creation and becomes read-only after that.
	Type TestType
	// Status is the life-cycle status of the test.
	Status TestStatus
	// Settings contains test configuration attributes.
	Settings TestSettings

	// Read-only properties

	// ID is unique test identification. It is read-only.
	ID models.ID
	// CreateDate is the creation timestamp. It is read-only.
	CreateDate time.Time
	// UpdateDate is the lst modification timestamp. It is read-only.
	UpdateDate time.Time
	// CreatedBy is an identity of the user that has created the test. It is read-only.
	CreatedBy UserInfo
	// LastUpdatedBy is the identity of the user that has modified the test last. It is read-only.
	LastUpdatedBy *UserInfo
}

// TestSettings contains test configuration attributes.
type TestSettings struct {
	// Definition contains test type specific configuration attributes.
	Definition TestDefinition
	// AgentIDs contains IDs of agents that shall execute tasks for this test. Only existing agents in the account
	// are allowed.
	AgentIDs []models.ID
	// Period is a test execution period. Default: 60s. Allowed values range: [1 s, 900 s].
	Period time.Duration
	// Family selects which type of DNS resource is queried for resolving hostname to target address.
	// It is used only for DNS and HTTP class of tests. Default: IPFamilyDual.
	Family IPFamily
	// NotificationChannels is a list of notifications channels for the tests. It must contain IDs of existing
	// notification channels.
	NotificationChannels []string
	// Health is a configuration of thresholds, acceptable status codes for evaluating test health
	// and activation conditions for alarms.
	Health HealthSettings
	// Ping is a configuration of ping task for the test.
	Ping *PingSettings
	// Traceroute if a configuration of traceroute task for the test.
	Traceroute *TracerouteSettings
	// Tasks is a list of names of tasks that shall be executed on behalf of this test.
	// Valid combinations of tasks depend on test type:
	// - IP, network grid, hostname, agent, network mesh and flow test types - ping and traceroute tasks (required)
	// - URL test type - HTTP task (required); ping and traceroute tasks (optional)
	// - page load test type - page load task (required); ping and traceroute tasks (optional)
	// - DNS and DNS grid test types - DNS task (required)
	// The system supports only running both ping and traceroute tasks together (as opposed to ping or traceroute task
	// individually).
	Tasks []TaskType
}

func (s TestSettings) GetIPDefinition() *TestDefinitionIP {
	d, _ := s.Definition.(*TestDefinitionIP) //nolint:errcheck // user can check the pointer
	return d
}

func (s TestSettings) GetNetworkGridDefinition() *TestDefinitionNetworkGrid {
	d, _ := s.Definition.(*TestDefinitionNetworkGrid) //nolint:errcheck // user can check the pointer
	return d
}

func (s TestSettings) GetHostnameDefinition() *TestDefinitionHostname {
	d, _ := s.Definition.(*TestDefinitionHostname) //nolint:errcheck // user can check the pointer
	return d
}

func (s TestSettings) GetAgentDefinition() *TestDefinitionAgent {
	d, _ := s.Definition.(*TestDefinitionAgent) //nolint:errcheck // user can check the pointer
	return d
}

func (s TestSettings) GetNetworkMeshDefinition() *TestDefinitionNetworkMesh {
	d, _ := s.Definition.(*TestDefinitionNetworkMesh) //nolint:errcheck // user can check the pointer
	return d
}

func (s TestSettings) GetFlowDefinition() *TestDefinitionFlow {
	d, _ := s.Definition.(*TestDefinitionFlow) //nolint:errcheck // user can check the pointer
	return d
}

func (s TestSettings) GetURLDefinition() *TestDefinitionURL {
	d, _ := s.Definition.(*TestDefinitionURL) //nolint:errcheck // user can check the pointer
	return d
}

func (s TestSettings) GetPageLoadDefinition() *TestDefinitionPageLoad {
	d, _ := s.Definition.(*TestDefinitionPageLoad) //nolint:errcheck // user can check the pointer
	return d
}

func (s TestSettings) GetDNSDefinition() *TestDefinitionDNS {
	d, _ := s.Definition.(*TestDefinitionDNS) //nolint:errcheck // user can check the pointer
	return d
}

func (s TestSettings) GetDNSGridDefinition() *TestDefinitionDNSGrid {
	d, _ := s.Definition.(*TestDefinitionDNSGrid) //nolint:errcheck // user can check the pointer
	return d
}

// TestDefinition emulates a union of:
// - TestDefinitionIP
// - TestDefinitionNetworkGrid
// - TestDefinitionHostname
// - TestDefinitionAgent
// - TestDefinitionNetworkMesh
// - TestDefinitionFlow
// - TestDefinitionURL
// - TestDefinitionPageLoad
// - TestDefinitionDNS
// - TestDefinitionDNSGrid
// Note that the interface is implemented with pointer receiver, so that definition's fields can be updated easily.
type TestDefinition interface {
	isTestDefinition()
}

// TestDefinitionIP contains the definition of TestTypeIP test.
type TestDefinitionIP struct {
	// Targets define target IP addresses.
	Targets []net.IP
}

func (d *TestDefinitionIP) isTestDefinition() {}

// TestDefinitionNetworkGrid contains the definition of TestTypeNetworkGrid test.
type TestDefinitionNetworkGrid struct {
	// Targets define target IP addresses.
	Targets []net.IP
}

func (d *TestDefinitionNetworkGrid) isTestDefinition() {}

// TestDefinitionHostname contains the definition of TestTypeHostname test.
type TestDefinitionHostname struct {
	// Target defines target fully qualified DNS name.
	Target string
}

func (d *TestDefinitionHostname) isTestDefinition() {}

// TestDefinitionAgent contains the definition of TestTypeAgent test.
type TestDefinitionAgent struct {
	// Target is an ID of target agent. Valid ID of an existing agent must be provided.
	Target models.ID
	// UseLocalIP indicates whether to use the configured "local" (internal) IP address as the source
	// when initiating probes.
	UseLocalIP bool
}

func (d *TestDefinitionAgent) isTestDefinition() {}

// TestDefinitionNetworkMesh contains the definition of TestTypeNetworkMesh test.
type TestDefinitionNetworkMesh struct {
	// UseLocalIP indicates whether to use the configured "local" (internal) IP address as the source
	// when initiating probes.
	UseLocalIP bool
}

func (d *TestDefinitionNetworkMesh) isTestDefinition() {}

// TestDefinitionFlow contains the definition of TestTypeFlow test.
type TestDefinitionFlow struct {
	// Type defines subtype of the test, which also specifies the type of target.
	Type FlowTestType
	// Target is a value to be matched in the query. Depending on Type, it must be: AS number, CDN name, country name,
	// region name or city name.
	Target string
	// TargetRefreshInterval is a period between regenerating list of targets based on flow data query.
	// Default: 0. Allowed values: 0 or range [1 hour, 168 hours].
	TargetRefreshInterval time.Duration
	// MaxIPTargets is maximum number of target IP addresses to select based on flow data query.
	// Default: 10. Allowed values range: [1, 20].
	MaxIPTargets uint32
	// MaxProviders is a maximum number of providers tracked for selection of target IP addresses.
	// Default: 3. Allowed values range: [1, 10].
	MaxProviders uint32
	// Direction specifies whether to match the (sub) type attribute in source or destination of flows
	// in the flow data query.
	Direction Direction
	// InetDirection specifies whether to use source address in inbound flows or destination addresses
	// in outbound flows.
	InetDirection Direction
}

func (d *TestDefinitionFlow) isTestDefinition() {}

// TestDefinitionURL contains the definition of TestTypeURL test.
type TestDefinitionURL struct {
	// Target is a URL to use in the HTTP request.
	Target url.URL
	// Timeout is an HTTP request timeout. Allowed values range: [5 s, 60 s].
	Timeout time.Duration
	// Method is an HTTP method to use in the request. Default: GET. Allowed values: GET, PATCH, POST, PUT.
	Method string
	// Headers is a set of key-value pairs to be included among HTTP headers in the request. Valid HTTP header names
	// and values are expected.
	Headers map[string]string
	// Body is a content to be placed in the body of the request.
	Body string
	// IgnoreTLSErrors is an indication whether to ignore errors reported in TLS session establishment.
	IgnoreTLSErrors bool
}

func (d *TestDefinitionURL) isTestDefinition() {}

// TestDefinitionPageLoad contains the definition of TestTypePageLoad test.
type TestDefinitionPageLoad struct {
	// Target is a URL to use in the HTTP request.
	Target url.URL
	// Timeout is an HTTP request timeout. Allowed values range: [5 s, 60 s].
	Timeout time.Duration
	// Headers is a set of key-value pairs to be included among HTTP headers in the request. Valid HTTP header names
	// and values are expected.
	Headers map[string]string
	// CSSSelectors is a set of key-value pairs to set as CSS selectors in the request. Valid HTTP CSS selector keys
	// and values are expected.
	CSSSelectors map[string]string
	// IgnoreTLSErrors is an indication whether to ignore errors reported in TLS session establishment.
	IgnoreTLSErrors bool
}

func (d *TestDefinitionPageLoad) isTestDefinition() {}

// TestDefinitionDNS contains the definition of TestTypeDNS test.
type TestDefinitionDNS struct {
	// Target is a fully qualified DNS name to resolve.
	Target string
	// Timeout is a DNS request timeout. Allowed values range: [5 s, 60 s].
	Timeout time.Duration
	// RecordType is a type of DNS record to query.
	RecordType DNSRecord
	// Servers is a list of addresses of servers to query. At least one entry is required.
	Servers []net.IP
	// Port is a server port to use. Allowed values range: [1, 65535].
	Port uint32
}

func (d *TestDefinitionDNS) isTestDefinition() {}

// TestDefinitionDNSGrid contains the definition of TestTypeDNSGrid test.
type TestDefinitionDNSGrid = TestDefinitionDNS

// HealthSettings is a configuration of thresholds, acceptable status codes for evaluating test health
// and activation conditions for alarms.
type HealthSettings struct {
	// LatencyCritical is a threshold for critical level of the average value of latency.
	// 0 means no health check. Default: 0. Allowed values: >= 0.
	LatencyCritical time.Duration
	// LatencyWarning is a threshold for warning level of the average value of latency.
	// 0 means no health check. Default: 0. Allowed values: >= 0.
	LatencyWarning time.Duration
	// LatencyCriticalStdDev is a threshold for critical level of the standard deviation of latency.
	// 0 means no health check. Default: 0. Allowed values range: [0, 100 ms].
	LatencyCriticalStdDev time.Duration
	// LatencyWarningStdDev is a threshold for warning level of the standard deviation of latency.
	// 0 means no health check. Default: 0. Allowed values range: [0, 100 ms].
	LatencyWarningStdDev time.Duration
	// JitterCritical is a threshold for critical level of the average value of jitter.
	// 0 means no health check. Default: 0. Allowed values: >= 0.
	JitterCritical time.Duration
	// JitterWarning is a threshold for warning level of the average value of jitter.
	// 0 means no health check. Default: 0. Allowed values: >= 0.
	JitterWarning time.Duration
	// JitterCriticalStdDev is a threshold for critical level of the standard deviation of jitter.
	// 0 means no health check. Default: 0. Allowed values range: [0, 100 ms].
	JitterCriticalStdDev time.Duration
	// JitterWarningStdDev is a threshold for warning level of the standard deviation of jitter.
	// 0 means no health check. Default: 0. Allowed values range: [0, 100 ms].
	JitterWarningStdDev time.Duration
	// PacketLossCritical is a threshold for critical level of packet loss (in percents).
	// 0 means no health check. Default: 0. Allowed values range: [0, 100].
	PacketLossCritical float32
	// PacketLossWarning is a threshold for warning level of packet loss (in percents).
	// 0 means no health check. Default: 0. Allowed values range: [0, 100].
	PacketLossWarning float32
	// HTTPLatencyCritical is a threshold for critical level of the average value of HTTP response latency.
	// 0 means no health check. Default: 0. Allowed values: >= 0.
	HTTPLatencyCritical time.Duration
	// HTTPLatencyWarning is a threshold for warning level of the average value of HTTP response latency.
	// 0 means no health check. Default: 0. Allowed values: >= 0.
	HTTPLatencyWarning time.Duration
	// HTTPLatencyCriticalStdDev is a threshold for critical level of the standard deviation of HTTP response latency.
	// 0 means no health check. Default: 0. Allowed values range: [0, 100 ms].
	HTTPLatencyCriticalStdDev time.Duration
	// HTTPLatencyWarningStdDev is a threshold for warning level of the standard deviation of HTTP response latency
	// 0 means no health check. Default: 0. Allowed values range: [0, 100 ms].
	HTTPLatencyWarningStdDev time.Duration
	// HTTPValidCodes is a list of HTTP result codes indicating success. Only valid HTTP result codes are accepted.
	// Empty list means result code is not checked.
	HTTPValidCodes []uint32
	// DNSValidCodes is a list of DNS result codes indicating success. Only valid DNS result codes are accepted.
	// Empty list means result code is not checked.
	DNSValidCodes []uint32
	// UnhealthySubtestThreshold is a number of tasks that has to be declared unhealthy in order for the test
	// to be declared unhealthy. Default: 1. Allowed values: > 0.
	UnhealthySubtestThreshold uint32
	// AlarmActivation sets activation conditions for alarms generated based on health thresholds.
	AlarmActivation *AlarmActivationSettings
}

// AlarmActivationSettings sets activation conditions for alarms generated based on health thresholds.
type AlarmActivationSettings struct {
	// TimeWindow is an activation window. The value in seconds must be greater or equal
	// to TestSettings.Period * (times + 1). Default: 5 minutes.
	TimeWindow time.Duration
	// Times is a minimum number of unhealthy test events within the TimeWindow for alarm activation.
	// Default: 3. Allowed values: > 0.
	Times uint
	// Grace period is a maximum duration of continuous test healthy state not canceling test activation. Default: 2.
	GracePeriod uint
}

// PingSettings is a configuration of ping task for the test.
type PingSettings struct {
	// Timeout is a total timeout for one execution of the task. Allowed values range: [1 ms, 10000 ms].
	Timeout time.Duration
	// Count is a number of probe packets per one task execution. Allowed values: [1, 10].
	Count uint32
	// Delay is a delay between sending of individual probe packets. Default: 0. Allowed values: >= 0.
	Delay time.Duration
	// Protocol is a type of probe packets. Ping task sends either ICMP echo request or performs TCP half-connect
	// to specified destination port (SYN, SYN-ACK, RST).
	Protocol PingProtocol
	// Port is a destination TCP port to use in probe packets. It is required for TCP protocol.
	// Allowed values range: [1, 65535] for TCP; 0 for ICMP.
	Port uint32
}

// TracerouteSettings if a configuration of traceroute task for the test.
type TracerouteSettings struct {
	// Timeout is a total timeout for one execution of the task.
	// Allowed values range: [1 ms, 5 m]. It must be lower than TestSettings.Period.
	Timeout time.Duration
	// Count is a number of probe packets per one patch hop. Allowed values: [1, 5].
	Count uint32
	// Delay is a delay between sending of individual probe packets. Allowed values: >= 0.
	Delay time.Duration
	// Protocol is a type of probe packets.
	Protocol TracerouteProtocol
	// Port is a destination TCP or UDP port to use in probe packets. Default: 33434.
	// Allowed values range: [1, 65535] for TCP or UDP; 0 for ICMP.
	Port uint32
	// Limit is maximum number of hops to probe (e.e. maximum TTL). Allowed values range: [1, 255].
	Limit uint32
}

// UserInfo contains user identity information.
type UserInfo struct {
	// ID is unique identification of the user. It is read-only.
	ID models.ID
	// Email is e-mail address of the user. It is read-only.
	Email string
	// FullName is full name of the user. It is read-only.
	FullName string
}

// IPTestRequiredFields is a subset of Test fields required to create an IP test.
type IPTestRequiredFields struct {
	BasePingTraceTestRequiredFields
	Definition TestDefinitionIPRequiredFields
}

// TestDefinitionIPRequiredFields is a subset of TestDefinitionIP fields required to create an IP test.
// Currently, it contains all TestDefinitionIP fields.
type TestDefinitionIPRequiredFields = TestDefinitionIP

// NetworkGridTestRequiredFields is a subset of Test fields required to create a NetworkGrid test.
type NetworkGridTestRequiredFields struct {
	BasePingTraceTestRequiredFields
	Definition TestDefinitionNetworkGridRequiredFields
}

// TestDefinitionNetworkGridRequiredFields is a subset of TestDefinitionNetworkGrid fields required to create
// a NetworkGrid test. Currently, it contains all TestDefinitionNetworkGrid fields.
type TestDefinitionNetworkGridRequiredFields = TestDefinitionNetworkGrid

// HostnameTestRequiredFields is a subset Test of fields required to create a hostname test.
type HostnameTestRequiredFields struct {
	BasePingTraceTestRequiredFields
	Definition TestDefinitionHostnameRequiredFields
}

// TestDefinitionHostnameRequiredFields is a subset of TestDefinitionHostname fields required to create a hostname test.
// Currently, it contains all TestDefinitionHostname fields.
type TestDefinitionHostnameRequiredFields = TestDefinitionHostname

// AgentTestRequiredFields is a subset Test of fields required to create an agent test.
type AgentTestRequiredFields struct {
	BasePingTraceTestRequiredFields
	Definition TestDefinitionAgentRequiredFields
}

// TestDefinitionAgentRequiredFields is a subset of TestDefinitionAgent fields required to create an agent test.
type TestDefinitionAgentRequiredFields struct {
	Target models.ID
}

// NetworkMeshTestRequiredFields is a subset Test of fields required to create a network mesh test.
type NetworkMeshTestRequiredFields struct {
	BasePingTraceTestRequiredFields
	// Definition contains no required fields
}

// FlowTestRequiredFields is a subset Test of fields required to create a flow test.
type FlowTestRequiredFields struct {
	BasePingTraceTestRequiredFields
	Definition TestDefinitionFlowRequiredFields
}

// TestDefinitionFlowRequiredFields is a subset of TestDefinitionFlow fields required to create a flow test.
type TestDefinitionFlowRequiredFields struct {
	Type          FlowTestType
	Target        string
	Direction     Direction
	InetDirection Direction
}

// URLTestRequiredFields is a subset Test of fields required to create a URL test.
type URLTestRequiredFields struct {
	BaseTestRequiredFields
	Definition TestDefinitionURLRequiredFields
}

// TestDefinitionURLRequiredFields is a subset of TestDefinitionURL fields required to create a URL test.
type TestDefinitionURLRequiredFields struct {
	Target  url.URL
	Timeout time.Duration
}

// PageLoadTestRequiredFields is a subset Test of fields required to create a page load test.
type PageLoadTestRequiredFields struct {
	BaseTestRequiredFields
	Definition TestDefinitionPageLoadRequiredFields
}

// TestDefinitionPageLoadRequiredFields is a subset of TestDefinitionPageLoad fields required to create
// a page load test.
type TestDefinitionPageLoadRequiredFields struct {
	Target  url.URL
	Timeout time.Duration
}

// DNSTestRequiredFields is a subset Test of fields required to create a DNS test.
type DNSTestRequiredFields struct {
	BaseTestRequiredFields
	Definition TestDefinitionDNSRequiredFields
}

// TestDefinitionDNSRequiredFields is a subset of TestDefinitionDNS fields required to create a DNS test.
// Currently, it contains all TestDefinition fields.
type TestDefinitionDNSRequiredFields = TestDefinitionDNS

// DNSGridTestRequiredFields is a subset Test of fields required to create a DNS grid test.
type DNSGridTestRequiredFields struct {
	BaseTestRequiredFields
	Definition TestDefinitionDNSGridRequiredFields
}

// TestDefinitionDNSGridRequiredFields is a subset of TestDefinitionDNSGrid fields required to create a DNS grid test.
// Currently, it contains all TestDefinition fields.
type TestDefinitionDNSGridRequiredFields = TestDefinitionDNSGrid

// BasePingTraceTestRequiredFields is a subset of Test fields required to create a test with ping and traceroute tasks.
type BasePingTraceTestRequiredFields struct {
	BaseTestRequiredFields
	Ping       PingSettingsRequiredFields
	Traceroute TracerouteSettingsRequiredFields
}

// BaseTestRequiredFields is a subset of Test fields required to create any test.
type BaseTestRequiredFields struct {
	Name     string
	AgentIDs []models.ID
}

// PingSettingsRequiredFields is a subset of PingSettings fields required for create.
type PingSettingsRequiredFields struct {
	Timeout  time.Duration
	Count    uint32
	Protocol PingProtocol
}

// TracerouteSettingsRequiredFields is a subset of TracerouteSettings fields required for create.
type TracerouteSettingsRequiredFields struct {
	Timeout  time.Duration
	Count    uint32
	Delay    time.Duration
	Protocol TracerouteProtocol
	Limit    uint32
}

// TestType is the specified type of the test.
type TestType string

const (
	// TestTypeIP allows testing of a multiple target addresses from one or more agents.
	TestTypeIP TestType = "ip"
	// TestTypeNetworkGrid allows testing of a multiple target addresses from one or more agents.
	// It differs from the TestTypeIP only in presentation of results in the UI.
	TestTypeNetworkGrid TestType = "network_grid"
	// TestTypeHostname allows testing of a single target defined by DNS name. It resolves the
	// DNS name to IP address and selects destination address(es) based on the family setting. If
	// the target name has both IPv4 and IPv6 resolution and the family is IP_FAMILY_DUAL, ping
	// and trace are run against both.
	TestTypeHostname TestType = "hostname"
	// TestTypeAgent allows probing from one or more agents to another agent.
	TestTypeAgent TestType = "agent"
	// TestTypeNetworkMesh allows to probe paths between a set of agents. Every agent probes
	// every other agent.The common setting agentIds attribute is used as a list of targets for the test.
	TestTypeNetworkMesh TestType = "network_mesh"
	// TestTypeFlow (called “Autonomous Tests” in the UI) allow to automatically select test targets
	// based on a query to flow data. The test configuration specifies parameters for the query.
	TestTypeFlow TestType = "flow"
	// TestTypeURL is an HTTP application test that verifies ability to execute HTTP request against
	// specific URL end-point and collect observations on latency in various stages of processing
	// (DNS resolution, TCP connection establishment, response latency). In addition to that it allows
	// to run ping and trace tasks targeting addresses to which the hostname in the URL resolves.
	TestTypeURL TestType = "url"
	// TestTypePageLoad is similar to TestTypeURL but provides more detailed information
	// about requests processing.
	TestTypePageLoad TestType = "page_load"
	// TestTypeDNS allows to test availability and latency of resolutions for a set of target DNS names
	// using a set of DNS servers. There is no functional difference between this and TestTypeDNSGrid
	// test, other than visual representation of results in the UI.
	TestTypeDNS TestType = "dns"
	// TestTypeDNSGrid is similar to TestTypeDNS described above.
	TestTypeDNSGrid TestType = "dns_grid"
)

// FlowTestType is a subtype of a synthetics flow test.
type FlowTestType string

const (
	FlowTestTypeASN     FlowTestType = "asn"
	FlowTestTypeCDN     FlowTestType = "cdn"
	FlowTestTypeCountry FlowTestType = "country"
	FlowTestTypeRegion  FlowTestType = "region"
	FlowTestTypeCity    FlowTestType = "city"
)

// TestStatus is the life-cycle status of the test.
type TestStatus string

const (
	TestStatusActive  TestStatus = "TEST_STATUS_ACTIVE"
	TestStatusPaused  TestStatus = "TEST_STATUS_PAUSED"
	TestStatusDeleted TestStatus = "TEST_STATUS_DELETED"
)

// TaskType is a name of the task that shall be executed on behalf of synthetics test.
type TaskType string

const (
	TaskTypeDNS        TaskType = "dns"
	TaskTypeHTTP       TaskType = "http"
	TaskTypePageLoad   TaskType = "page-load"
	TaskTypePing       TaskType = "ping"
	TaskTypeTraceroute TaskType = "traceroute"
)

// Direction is a source and destination enumeration.
type Direction string

const (
	DirectionSrc Direction = "src"
	DirectionDst Direction = "dst"
)

// DNSRecord is a DNS record type.
type DNSRecord string

const (
	DNSRecordUnspecified DNSRecord = "DNS_RECORD_UNSPECIFIED"
	DNSRecordA           DNSRecord = "DNS_RECORD_A"
	DNSRecordAAAA        DNSRecord = "DNS_RECORD_AAAA"
	DNSRecordCName       DNSRecord = "DNS_RECORD_CNAME"
	DNSRecordDName       DNSRecord = "DNS_RECORD_DNAME"
	DNSRecordNS          DNSRecord = "DNS_RECORD_NS"
	DNSRecordMX          DNSRecord = "DNS_RECORD_MX"
	DNSRecordPTR         DNSRecord = "DNS_RECORD_PTR"
	DNSRecordSOA         DNSRecord = "DNS_RECORD_SOA"
)

// PingProtocol is a type of probe packets for ping task.
type PingProtocol string

const (
	PingProtocolICMP PingProtocol = "icmp"
	PingProtocolTCP  PingProtocol = "tcp"
)

// TracerouteProtocol is a type of probe packets for traceroute task.
type TracerouteProtocol string

const (
	TracerouteProtocolICMP TracerouteProtocol = "icmp"
	TracerouteProtocolTCP  TracerouteProtocol = "tcp"
	TracerouteProtocolUDP  TracerouteProtocol = "udp"
)
