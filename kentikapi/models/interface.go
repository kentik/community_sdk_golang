package models

import "time"

// Note: InterfacesAPI belongs under DevicesAPI, but it is vast, so it lives in a separate file.

type Interface struct {
	// Read-write properties

	SNMPID               ID
	SNMPSpeed            int
	InterfaceDescription string // if fact, this is interface name, that's why not optional
	SNMPAlias            *string
	InterfaceIP          *string
	InterfaceIPNetmask   *string
	VRF                  *VRFAttributes
	VRFID                *ID
	SecondaryIPS         []SecondaryIP

	// Read-only properties

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

// InterfaceRequiredFields is a subset of Interface fields required to create an Interface.
type InterfaceRequiredFields struct {
	DeviceID             ID
	SNMPID               ID
	SNMPSpeed            int
	InterfaceDescription string
}

// NewInterface creates a new Interface with all necessary fields set.
func NewInterface(i InterfaceRequiredFields) *Interface {
	return &Interface{
		DeviceID:             i.DeviceID,
		SNMPID:               i.SNMPID,
		SNMPSpeed:            i.SNMPSpeed,
		InterfaceDescription: i.InterfaceDescription,
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

// VRFAttributesRequiredFields is a subset of VRFAttributes fields required to create a VRFAttributes.
type VRFAttributesRequiredFields struct {
	Name               string
	RouteTarget        string
	RouteDistinguisher string
}

// NewVRFAttributes creates new VRFAttributes with all necessary fields set.
func NewVRFAttributes(v VRFAttributesRequiredFields) *VRFAttributes {
	return &VRFAttributes{
		Name:               v.Name,
		RouteTarget:        v.RouteTarget,
		RouteDistinguisher: v.RouteDistinguisher,
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
