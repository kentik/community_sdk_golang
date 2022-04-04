package models

import "time"

type Device struct {
	// Read-write properties

	PlanID              *ID
	SiteID              *ID
	DeviceDescription   *string
	DeviceSampleRate    int
	SendingIPs          []string
	DeviceSNMPIP        *string
	DeviceSNMPCommunity *string
	MinimizeSNMP        *bool
	// Note: for DeviceBGPType = DeviceBGPTypeDevice, either DeviceBGPNeighborIP or DeviceBGPNeighborIPv6 is required
	DeviceBGPType         *DeviceBGPType
	DeviceBGPNeighborIP   *string
	DeviceBGPNeighborIPv6 *string
	DeviceBGPNeighborASN  *string
	DeviceBGPFlowSpec     *bool
	DeviceBGPPassword     *string
	UseBGPDeviceID        *ID
	DeviceSNMPv3Conf      *SNMPv3Conf
	CDNAttr               *CDNAttribute

	// Read-only properties

	ID              ID
	DeviceName      string
	DeviceType      DeviceType
	DeviceSubType   DeviceSubtype
	DeviceStatus    *string
	DeviceFlowType  *string
	CompanyID       ID
	SNMPLastUpdated *string
	CreatedDate     time.Time
	UpdatedDate     time.Time
	BGPPeerIP4      *string
	BGPPeerIP6      *string
	Plan            DevicePlan
	Site            *DeviceSite
	Labels          []DeviceLabel
	AllInterfaces   []AllInterfaces
}

// DeviceRouterRequiredFields is a subset of Device fields required to create a DeviceRouter.
type DeviceRouterRequiredFields struct {
	DeviceName       string
	DeviceSubType    DeviceSubtype
	DeviceSampleRate int
	PlanID           ID
	SendingIPs       []string
	MinimizeSNMP     bool
}

// NewDeviceRouter creates a new Router with all necessary fields set
// Optional fields that can be set for a router include:
// - DeviceSNMPIP
// - DeviceSNMPCommunity
// - DeviceSNMPv3Conf  // when set, overwrites "device_snmp_community"
// Optional fields that can be always set include:
// - DeviceDescription
// - SiteID
// - DeviceBGPFlowSpec.
func NewDeviceRouter(r DeviceRouterRequiredFields) *Device {
	bgpType := DeviceBGPTypeNone // default
	return &Device{
		DeviceType:       DeviceTypeRouter,
		DeviceName:       r.DeviceName,
		DeviceSubType:    r.DeviceSubType,
		DeviceSampleRate: r.DeviceSampleRate,
		PlanID:           &r.PlanID,
		DeviceBGPType:    &bgpType,
		SendingIPs:       r.SendingIPs,
		MinimizeSNMP:     &r.MinimizeSNMP,
	}
}

// DeviceDNSRequiredFields is a subset of Device fields required to create a DeviceDNS.
type DeviceDNSRequiredFields struct {
	DeviceName       string
	DeviceSubType    DeviceSubtype
	DeviceSampleRate int
	PlanID           ID
	CDNAttr          CDNAttribute
}

// NewDeviceDNS creates a new DSN with all necessary fields set
// Optional fields that can be set include:
// - DeviceDescription
// - SiteID
// - DeviceBGPFlowSpec.
func NewDeviceDNS(d DeviceDNSRequiredFields) *Device {
	bgpType := DeviceBGPTypeNone // default
	return &Device{
		DeviceType:       DeviceTypeHostNProbeDNSWWW,
		DeviceName:       d.DeviceName,
		DeviceSubType:    d.DeviceSubType,
		DeviceSampleRate: d.DeviceSampleRate,
		PlanID:           &d.PlanID,
		DeviceBGPType:    &bgpType,
		CDNAttr:          &d.CDNAttr,
	}
}

// WithBGPTypeDevice is alternative to WithBGPTypeOtherDevice
// Optional fields that can be set for BGPTypeDevice include:
// - DeviceBGPPassword
// Note: either DeviceBGPNeighborIP or DeviceBGPNeighborIPv6 is required for DeviceBGPTypeDevice.
func (d *Device) WithBGPTypeDevice(deviceBGPNeighborASN string) *Device {
	bgpType := DeviceBGPTypeDevice
	d.DeviceBGPType = &bgpType
	d.DeviceBGPNeighborASN = &deviceBGPNeighborASN
	return d
}

// WithBGPTypeOtherDevice is alternative to WithBGPTypeDevice.
func (d *Device) WithBGPTypeOtherDevice(useBGPDeviceID ID) *Device {
	bgpType := DeviceBGPTypeOtherDevice
	d.DeviceBGPType = &bgpType
	d.UseBGPDeviceID = &useBGPDeviceID
	return d
}

type AllInterfaces struct {
	InterfaceDescription string
	DeviceID             ID
	SNMPSpeed            float64
	InitialSNMPSpeed     *float64
}

type SNMPv3Conf struct {
	UserName                 string
	AuthenticationProtocol   *AuthenticationProtocol
	AuthenticationPassphrase *string
	PrivacyProtocol          *PrivacyProtocol
	PrivacyPassphrase        *string
}

func NewSNMPv3Conf(userName string) *SNMPv3Conf {
	return &SNMPv3Conf{UserName: userName}
}

func (c *SNMPv3Conf) WithAuthentication(protocol AuthenticationProtocol, pass string) *SNMPv3Conf {
	c.AuthenticationProtocol = &protocol
	c.AuthenticationPassphrase = &pass
	return c
}

func (c *SNMPv3Conf) WithPrivacy(protocol PrivacyProtocol, pass string) *SNMPv3Conf {
	c.PrivacyProtocol = &protocol
	c.PrivacyPassphrase = &pass
	return c
}

// DeviceSite embedded under Device differs from a regular Site in that all fields are optional.
type DeviceSite struct {
	ID        *ID
	CompanyID *ID
	Latitude  *float64
	Longitude *float64
	SiteName  *string
}

// DevicePlan embedded under Device differs from a regular Plan in that all fields are optional.
type DevicePlan struct {
	ID            *ID
	CompanyID     *ID
	Name          *string
	Description   *string
	Active        *bool
	MaxDevices    *int
	MaxFPS        *int
	BGPEnabled    *bool
	FastRetention *int
	FullRetention *int
	CreatedDate   *time.Time
	UpdatedDate   *time.Time
	MaxBigdataFPS *int
	DeviceTypes   []PlanDeviceType
	Devices       []PlanDevice
}

type AppliedLabels struct {
	// Read-only properties

	ID         ID
	DeviceName string
	Labels     []DeviceLabel
}

type DeviceType string

const (
	DeviceTypeRouter           DeviceType = "router"
	DeviceTypeHostNProbeDNSWWW DeviceType = "host-nprobe-dns-www"
)

type DeviceSubtype string

// Device subtypes for type DeviceTypeRouter.
const (
	DeviceSubtypeRouter                 DeviceSubtype = "router"
	DeviceSubtypeCiscoASA               DeviceSubtype = "cisco_asa"
	DeviceSubtypePaloAlto               DeviceSubtype = "paloalto"
	DeviceSubtypeSilverpeak             DeviceSubtype = "silverpeak"
	DeviceSubtypeMPLS                   DeviceSubtype = "mpls"
	DeviceSubtypeViptela                DeviceSubtype = "viptela"
	DeviceSubtypePFESyslog              DeviceSubtype = "pfe_syslog"
	DeviceSubtypeSyslog                 DeviceSubtype = "syslog"
	DeviceSubtypeMeraki                 DeviceSubtype = "meraki"
	DeviceSubtypeIstio                  DeviceSubtype = "istio"
	DeviceSubtypeIOSXR                  DeviceSubtype = "ios_xr"
	DeviceSubtypeCiscoZoneBasedFirewall DeviceSubtype = "cisco_zone_based_firewall"
	DeviceSubtypeCiscoNBAR              DeviceSubtype = "cisco_nbar"
	DeviceSubtypeCiscoASASyslog         DeviceSubtype = "cisco_asa_syslog"
	DeviceSubtypeAdvancedSFlow          DeviceSubtype = "advanced_sflow"
	DeviceSubtypeA10CGN                 DeviceSubtype = "a10_cgn"
)

// Device subtypes for type DeviceTypeHostNProbeDNSWWW.
const (
	DeviceSubtypeKProbe      DeviceSubtype = "kprobe"
	DeviceSubtypeNProbe      DeviceSubtype = "nprobe"
	DeviceSubtypeAWSSubnet   DeviceSubtype = "aws_subnet"
	DeviceSubtypeAzureSubnet DeviceSubtype = "azure_subnet"
	DeviceSubtypeGCPSubnet   DeviceSubtype = "gcp_subnet"
	DeviceSubtypeKappa       DeviceSubtype = "kappa"      // not in the API documentation
	DeviceSubtypeIBMSubnet   DeviceSubtype = "ibm_subnet" // not in the API documentation
)

type DeviceBGPType string

const (
	DeviceBGPTypeNone        DeviceBGPType = "none"
	DeviceBGPTypeDevice      DeviceBGPType = "device"
	DeviceBGPTypeOtherDevice DeviceBGPType = "other_device"
)

func DeviceBGPTypePtr(d DeviceBGPType) *DeviceBGPType {
	return &d
}

type AuthenticationProtocol string

const (
	AuthenticationProtocolNoAuth AuthenticationProtocol = "NoAuth"
	AuthenticationProtocolMD5    AuthenticationProtocol = "MD5"
	AuthenticationProtocolSHA    AuthenticationProtocol = "SHA"
)

type PrivacyProtocol string

const (
	PrivacyProtocolNoPriv PrivacyProtocol = "NoPriv"
	PrivacyProtocolDES    PrivacyProtocol = "DES"
	PrivacyProtocolAES    PrivacyProtocol = "AES"
)

type CDNAttribute string

const (
	CDNAttributeNone CDNAttribute = "None"
	CDNAttributeYes  CDNAttribute = "Y"
	CDNAttributeNo   CDNAttribute = "N"
)

func CDNAttributePtr(c CDNAttribute) *CDNAttribute {
	return &c
}
