package synthetics

import (
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/cloud"
)

// GetAllAgentsResponse model.
type GetAllAgentsResponse struct {
	// Agents holds all synthetics agents.
	Agents []Agent
	// InvalidAgentsCount is a number of invalid agents.
	InvalidAgentsCount uint32
}

// Agent is synthetics agent model.
// The Kentik synthetics monitoring system allows customers to deploy private test agents in
// their infrastructure and/or use one of the global agents maintained by Kentik. All
// information provided by the API for global agents is read-only, so the read-write status below
// applies only to private agents.
type Agent struct {
	// Read-write properties (for private agents)

	// Status is a life-cycle status of the agent. Only AgentStatusOK and AgentStatusWait can be set via the API.
	Status AgentStatus
	// Alias is user selected name of the agent.
	Alias string
	// SiteID is unique identification of the site where the agent is located.
	// Allowed values: a valid identifier of existing site (in Kentik configuration).
	// The field is ignored if CloudProvider and CloudRegion are set.
	SiteID string
	// LocalIP is a private/internal address of an agent behind address translation.
	// Allowed values: valid IPv4 or IPv6 address in standard notation.
	LocalIP string
	// IPFamily is the IP address family the agent supports (for running tests).
	IPFamily IPFamily
	// CloudProvider is the name of the cloud provider for agents hosted in a public cloud (otherwise an empty string).
	CloudProvider cloud.Provider
	// CloudRegion is the name of the cloud region for agents hosted in a public cloud (otherwise an empty string).
	// Allowed values: a valid name of a region for the cloud provider.
	CloudRegion string

	// Read-only properties

	// ID is unique agent identification. It is read-only.
	ID string
	// Type is an indication whether agent is private or global (public). It is read-only.
	Type AgentType
	// SiteName is the name of the site where the agent is located. It is read-only.
	// The value is imported based on SiteID or CloudProvider and CloudRegion.
	SiteName string
	// IP is the public IP address of the agent (as seen by the Kentik system). It is read-only.
	// Allowed values: valid IPv4 or IPv6 address in standard notation.
	IP string
	// ASN is an autonomous system number derived from the public address of the agent. It is read-only.
	ASN uint32
	// Latitude is the latitude of agent's location (derived from site). It is read-only.
	Latitude float64
	// Longitude is the longitude of agent's location (derived from site). It is read-only.
	Longitude float64
	// City is the name of the city where the agent is located (derived from site). It is read-only.
	City string
	// Region is the name of the geographic region where the agent is located (derived from site). It is read-only.
	Region string
	// Country is the name of the country where the agent is located (derived from site). It is read-only.
	Country string
	// Version is the software version of the agent. It is read-only.
	Version string
	// OS is the version of the operating system hosting the agent. It is read-only.
	OS string
	// ImplementationType is the implementation type of the agent. It affects types of test the agent can execute.
	// It is read-only.
	ImplementationType AgentImplementationType
	// LastAuthed is a timestamp of last authentication of the agent to the Kentik system. It is read-only.
	LastAuthed time.Time
	// TestIDs is a list of test IDs the agent is  currently servicing. It is read-only.
	TestIDs []string
}

// AgentStatus is a life-cycle status of the agent.
type AgentStatus string

const (
	AgentStatusUnspecified AgentStatus = "AGENT_STATUS_UNSPECIFIED"
	AgentStatusOK          AgentStatus = "AGENT_STATUS_OK"
	AgentStatusWait        AgentStatus = "AGENT_STATUS_WAIT"
	AgentStatusDeleted     AgentStatus = "AGENT_STATUS_DELETED"
)

// AgentType is an indication whether agent is private or global (public).
type AgentType string

const (
	AgentTypePrivate AgentType = "private"
	AgentTypeGlobal  AgentType = "global"
)

// AgentImplementationType is the implementation type of the agent.
type AgentImplementationType string

const (
	AgentImplementationTypeUnspecified AgentImplementationType = "IMPLEMENT_TYPE_UNSPECIFIED"
	// AgentImplementationTypeRust specifies synthetics agent implemented in Rust capable of running all tasks
	// except for page-load.
	AgentImplementationTypeRust AgentImplementationType = "IMPLEMENT_TYPE_RUST"
	// AgentImplementationTypeNode specifies synthetics agent implemented in NodeJS capable of running ping,
	// traceroute and page-load tasks.
	AgentImplementationTypeNode AgentImplementationType = "IMPLEMENT_TYPE_NODE"
)

// IPFamily is the IP address family.
type IPFamily string

const (
	IPFamilyUnspecified IPFamily = "IP_FAMILY_UNSPECIFIED"
	IPFamilyV4          IPFamily = "IP_FAMILY_V4"
	IPFamilyV6          IPFamily = "IP_FAMILY_V6"
	IPFamilyDual        IPFamily = "IP_FAMILY_DUAL"
)
