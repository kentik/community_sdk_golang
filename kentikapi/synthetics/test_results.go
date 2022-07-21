package synthetics

import (
	"net"
	"net/url"
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

// GetResultsForTestsRequest contains parameters for GetResultsForTests request.
type GetResultsForTestsRequest struct {
	// TestIDs is a list of IDs of the tests to get health for.
	TestIDs []models.ID
	// StartTime is a start of the time interval for this query.
	StartTime time.Time
	// EndTime is an end of the time interval for this query.
	EndTime time.Time
	// AgentIDs is an optional subset of agents to retrieve results for.
	AgentIDs []models.ID
	// Targets is an optional subset of destination IP addresses to retrieve results for.
	Targets []net.IP
}

// TestResults contains test observations for one synthetic test for a point in time.
type TestResults struct {
	// TestID is an ID of the test which results are for.
	TestID models.ID
	// Time is a timestamp of the observation.
	Time time.Time
	// Health is an evaluation of health for this set of observations.
	Health Health
	// Agents contains actual observations for each agent.
	Agents []AgentResults
}

// AgentResults contains observations for a particular agent.
type AgentResults struct {
	// AgentID is an ID of the reporting agent.
	AgentID models.ID
	// Health is an evaluation of health for this set of observations.
	Health Health
	// Tasks contains actual observations for particular agent.
	Tasks []TaskResults
}

// TaskResults contains observations.
type TaskResults struct {
	// Health is an evaluation of health for this set of observations.
	Health Health
	// TaskType indicates the type of task results held in Task field.
	// Supported values: TaskTypePing, TaskTypeHTTP, TaskTypeDNS.
	// Note that both URL and page load tests produce TaskTypeHTTP task results.
	// For example, if TaskType is TaskTypePing, then Task holds PingResults.
	TaskType TaskType
	// Task contains actual observations.
	Task TaskSpecificResults
}

func (r TaskResults) GetPingResults() *PingResults {
	ping, _ := r.Task.(PingResults) // nolint:errcheck // user can check the pointer
	return &ping
}

func (r TaskResults) GetHTTPResults() *HTTPResults {
	http, _ := r.Task.(HTTPResults) // nolint:errcheck // user can check the pointer
	return &http
}

func (r TaskResults) GetDNSResults() *DNSResults {
	dns, _ := r.Task.(DNSResults) // nolint:errcheck // user can check the pointer
	return &dns
}

// TaskSpecificResults emulates a union of PingResults, HTTPResults and DNSResults.
type TaskSpecificResults interface {
	isTaskSpecificResults()
}

// PingResults contains observations for ping task.
type PingResults struct {
	// Target is a task target (IP address, hostname)
	Target string
	// PacketLoss contains packet loss observed in the current period.
	PacketLoss PacketLossData
	// Latency contains the observed latency.
	Latency MetricData
	// Jitter contains the observed jitter.
	Jitter MetricData
	// DstIP is a destination IP address.
	DstIP net.IP
}

func (r PingResults) isTaskSpecificResults() {}

// HTTPResults contains observations for HTTP or page load task.
type HTTPResults struct {
	// Target is a URL of the task target.
	Target url.URL
	// Latency is the observed HTTP response latency.
	Latency MetricData
	// Response is an object carrying information about received HTTP response.
	Response HTTPResponseData
	// DstIP is a destination IP address.
	DstIP net.IP
}

func (r HTTPResults) isTaskSpecificResults() {}

// HTTPResponseData is an object carrying information about received HTTP response.
type HTTPResponseData struct {
	// Status is an HTTP status code.
	Status uint32
	// Size is a size of received response body in bytes.
	Size uint32
	// Data contains details about request delivery and response processing timing.
	// It differs for HTTP and page load tasks.
	// TODO(dfurman): point to the documentation with HTTP / page load response details
	//  or model as an union of structs.
	Data []map[string]interface{}
}

// DNSResults contains observations for DNS task.
type DNSResults struct {
	// Target is a DNS name.
	Target string
	// Server is an IP address of server queried.
	Server net.IP
	// Latency is the observed DNS response latency.
	Latency MetricData
	// Response holds information about received DNS response.
	Response DNSResponseData
}

func (r DNSResults) isTaskSpecificResults() {}

// DNSResponseData holds information about received DNS response.
type DNSResponseData struct {
	// Status is a DNS status code.
	Status uint32
	// Data is a query result data (i.e. the DNS entity the query returned).
	Data string
}

// PacketLossData contains packet loss observation.
type PacketLossData struct {
	// Current is a packet loss observed in the interval as a number with range [0, 1].
	Current float64
	// Health is a health evaluation for the observation.
	Health Health
}

// MetricData contains latency/jitter observation.
type MetricData struct {
	// Current is an observation in the current interval.
	Current time.Duration
	// RollingAvg is a rolling average.
	RollingAvg time.Duration
	// RollingStdDev is a rolling standard deviation.
	RollingStdDev time.Duration
	// Health is a health evaluation of the observation.
	Health Health
}

// Health is a health evaluation of the observation.
type Health string

const (
	HealthHealthy  Health = "healthy"
	HealthWarning  Health = "warning"
	HealthFailing  Health = "failing"
	HealthCritical Health = "critical"
)
