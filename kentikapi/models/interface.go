package models

import "time"

// Note: InterfacesAPI belong under DevicesAPI but it is wast so it lives in a separate file

type Interface struct {
	// read-write properties (can be updated in update call)
	SNMPID               ID
	SNMPSpeed            int
	InterfaceDescription string // if fact, this is interface name, that's why not optional
	SNMPAlias            *string
	InterfaceIP          *string
	InterfaceIPNetmask   *string
	VRF                  *VRFAttributes
	VRFID                *ID
	SecondaryIPS         []SecondaryIP

	// read-only properties (can't be updated in update call)
	ID                          ID
	CompanyID                   ID
	DeviceID                    ID
	CreatedDate                 time.Time
	UpdatedDate                 time.Time
	InitialSNMPID               *string
	InitialSNMPAlias            *string
	InitialInterfaceDescription *string
	InitialSNMPSpeed            *int
	TopNextHopASNs              []TopNextHopASN
}

// NewInterface creates a new Interface with all necessary fields set
func NewInterface(deviceID ID, snmpID ID, snmpSpeed int, interfaceDescription string) *Interface {
	return &Interface{
		DeviceID:             deviceID,
		SNMPID:               snmpID,
		SNMPSpeed:            snmpSpeed,
		InterfaceDescription: interfaceDescription,
	}
}

type VRFAttributes struct {
	// read-write
	Name                  string
	RouteTarget           string
	RouteDistinguisher    string
	Description           *string
	ExtRouteDistinguisher *string

	// read-only
	ID        ID
	CompanyID ID
	DeviceID  ID
}

// NewVRFAttributes creates new VRFAttributes with all necessary fields set
func NewVRFAttributes(name string, routeTarget string, routeDistinguisher string) *VRFAttributes {
	return &VRFAttributes{
		Name:               name,
		RouteTarget:        routeTarget,
		RouteDistinguisher: routeDistinguisher,
	}
}

type SecondaryIP struct {
	Address string
	Netmask string
}

type TopNextHopASN struct {
	ASN     int
	Packets int
}
