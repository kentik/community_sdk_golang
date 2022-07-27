package synthetics

import (
	"net"
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

// GetTraceForTestRequest contains parameters for GetTraceForTest request.
type GetTraceForTestRequest struct {
	// TestID is and ID of the test to get traceroute results for.
	TestID models.ID
	// StartTime is a start of the time interval for this query.
	StartTime time.Time
	// EndTime is an end of the time interval for this query.
	EndTime time.Time
	// AgentIDs is an optional subset of agents to retrieve results for.
	AgentIDs []models.ID
	// Targets is an optional subset of destination IP addresses to retrieve results for.
	Targets []net.IP
}

// GetTraceForTestResponse contains results for GetTraceForTest request.
type GetTraceForTestResponse struct {
	// Nodes maps node IDs to network nodes.
	Nodes map[string]NetworkNode
	// Paths is a sequence of network path objects.
	Paths []Path
}

// NetworkNode represents a single trace hop.
type NetworkNode struct {
	// IP is an IP address of the ingress interface of the device.
	IP net.IP
	// ASN is a AS number owning the IP address.
	ASN uint32
	// AsName is he name of the AS owning the IP address.
	AsName string
	// Location is the geographic location of the device
	Location *Location
	// DNSName is a DNS name returned by reverse resolution of the IP address.
	DNSName string
	// DeviceID is an ID of device in Kentik configuration (empty if not available).
	DeviceID models.ID
	// SiteID is an ID of the site containing the device in Kentik configuration (empty if not available).
	SiteID models.ID
}

// Location represents the geographical location of a devices.
type Location struct {
	// Latitude is geographical latitude in decimal degrees.
	Latitude float64
	// Longitude is geographical longitude in decimal degrees.
	Longitude float64
	// Country is a country name associated with geographical location.
	Country string
	// Region is a region name associated with geographical location.
	Region string
	// City is a city name associated with geographical location.
	City string
}

// Path object describes results of one iteration of execution of a network path mapping task.
// It consists of one or more path traces (depending on the task configuration). Origin agent
// and target address are common to all path traces included in the path object. Path traces
// associated with a path may differ due to changes to network routing between attempts
// and/or due to equal cost multipath routing (ECMP).
// Each path trace contains one or more network hops which refer (via unique ID) to network nodes provided in
// the map of network nodes. Network nodes typically correspond to devices forwarding packets along the path.
type Path struct {
	// Time is a timestamp of start of execution of the trace task.
	Time time.Time
	// AgentID is an ID of the agent that executed the trace task.
	AgentID models.ID
	// TargetIP is an IP address of the target.
	TargetIP net.IP
	// HopCount contains hop count statistics across all path traces.
	HopCount Stats
	// MaxASPathLength is maximum length of the AS path across all traces.
	MaxASPathLength int32
	// Traces is a list of traces observed in the interval.
	Traces []PathTrace
}

// Stats stores minimum, maximum and average value of sequence of integer values.
type Stats struct {
	// Average value of sequence of integer values.
	Average int32
	// Min is minimum value of sequence of integer values.
	Min int32
	// Max is maximum value of sequence of integer values.
	Max int32
}

// PathTrace object describes the result of a single attempt to map a network path from the
// origin agent to the target address. The path trace consists of a sequence of trace hops.
type PathTrace struct {
	// ASPath is a sequence of AS numbers.
	ASPath []int32
	// IsComplete is an indication whether path trace reached the target.
	IsComplete bool
	// Hops is a sequence of hops describing the network path.
	Hops []TraceHop
}

// TraceHop represents a single forwarding hop on the path from origin agent to the target.
// The trace hop may either be resolved to a specific device or be unknown (if no ICMP Time Exceeded response
// has been received for the hop). If known, the device may either be present in customers’ Kentik configuration
// or be foreign.
type TraceHop struct {
	// Latency is a latency of response from this hop. It is 0 if no response is received.
	Latency int32
	// NodeID is a unique ID of the network node associated with the hop. It is empty if no response is received.
	NodeID models.ID
}
