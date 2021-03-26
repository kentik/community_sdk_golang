package models

import "time"

type Device struct {
	// read-write properties (can be updated in update call)
	PlanID                *ID
	SiteID                *ID
	DeviceDescription     *string
	DeviceSampleRate      int
	SendingIPS            []string
	DeviceSNMNPIP         *string
	DeviceSNMPCommunity   *string
	MinimizeSNMP          *bool
	DeviceBGPType         *DeviceBGPType // Note: for DeviceBGPType = DeviceBGPTypeDevice, either DeviceBGPNeighborIP or DeviceBGPNeighborIPv6 is required
	DeviceBGPNeighborIP   *string
	DeviceBGPNeighborIPv6 *string
	DeviceBGPNeighborASN  *string
	DeviceBGPFlowSpec     *bool
	DeviceBGPPassword     *string
	UseBGPDeviceID        *ID
	DeviceSNMPv3Conf      *SNMPv3Conf
	CDNAttr               *CDNAttribute

	// read-only properties (can't be updated in update call)
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

// NewDeviceRouter creates a new Router with all necessary fields set
// Optional fields that can be set for router include:
// - DeviceSNMNPIP
// - DeviceSNMPCommunity
// - DeviceSNMPv3Conf  // when set, overwrites "device_snmp_community"
// Optional fields that can be always set include:
// - DeviceDescription
// - SiteID
// - DeviceBGPFlowSpec
func NewDeviceRouter(
	// common required
	DeviceName string,
	DeviceSubType DeviceSubtype,
	DeviceSampleRate int,
	PlanID ID,
	// router required
	SendingIPS []string,
	MinimizeSNMP bool,
) *Device {
	bgpType := DeviceBGPTypeNone // default
	return &Device{
		DeviceType:       DeviceTypeRouter,
		DeviceName:       DeviceName,
		DeviceSubType:    DeviceSubType,
		DeviceSampleRate: DeviceSampleRate,
		PlanID:           &PlanID,
		DeviceBGPType:    &bgpType,
		SendingIPS:       SendingIPS,
		MinimizeSNMP:     &MinimizeSNMP,
	}
}

// NewDeviceDNS creates a new DSN with all necessary fields set
// Optional fields that can be set include:
// - DeviceDescription
// - SiteID
// - DeviceBGPFlowSpec
func NewDeviceDNS(
	// common required
	DeviceName string,
	DeviceSubType DeviceSubtype,
	DeviceSampleRate int,
	PlanID ID,
	// dns required
	CDNAttr CDNAttribute,
) *Device {
	bgpType := DeviceBGPTypeNone // default
	return &Device{
		DeviceType:       DeviceTypeHostNProbeDNSWWW,
		DeviceName:       DeviceName,
		DeviceSubType:    DeviceSubType,
		DeviceSampleRate: DeviceSampleRate,
		PlanID:           &PlanID,
		DeviceBGPType:    &bgpType,
		CDNAttr:          &CDNAttr,
	}
}

// WithBGPTypeDevice is alternative to WithBGPTypeOtherDevice
// Optional fields that can be set for BGPTypeDevice include:
// - DeviceBGPPassword
// Note: either DeviceBGPNeighborIP or DeviceBGPNeighborIPv6 is required for DeviceBGPTypeDevice
func (d *Device) WithBGPTypeDevice(deviceBGPNeighborASN string) *Device {
	bgpType := DeviceBGPTypeDevice
	d.DeviceBGPType = &bgpType
	d.DeviceBGPNeighborASN = &deviceBGPNeighborASN
	return d
}

// WithBGPTypeOtherDevice is alternative to WithBGPTypeDevice
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

// DeviceSite embedded under Device differs from regular Site in that all fields are optional
type DeviceSite struct {
	ID        *ID
	CompanyID *ID
	Latitude  *float64
	Longitude *float64
	SiteName  *string
}

// DevicePlan embedded under Device differs from regular Plan in that all fields are optional
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
	// read-only properties
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

const (
	// for DeviceType = DeviceTypeRouter
	DeviceSubtypeRouter                 DeviceSubtype = "router"
	DeviceSubtypeCiscoAsa               DeviceSubtype = "cisco_asa"
	DeviceSubtypePaloalto               DeviceSubtype = "paloalto"
	DeviceSubtypeSilverpeak             DeviceSubtype = "silverpeak"
	DeviceSubtypeMpls                   DeviceSubtype = "mpls"
	DeviceSubtypeViptela                DeviceSubtype = "viptela"
	DeviceSubtypePfeSyslog              DeviceSubtype = "pfe_syslog"
	DeviceSubtypeSyslog                 DeviceSubtype = "syslog"
	DeviceSubtypeMeraki                 DeviceSubtype = "meraki"
	DeviceSubtypeIstio                  DeviceSubtype = "istio"
	DeviceSubtypeIosxr                  DeviceSubtype = "ios_xr"
	DeviceSubtypeCiscoZoneBasedFirewall DeviceSubtype = "cisco_zone_based_firewall"
	DeviceSubtypeCiscoNbar              DeviceSubtype = "cisco_nbar"
	DeviceSubtypeCiscoAsaSyslog         DeviceSubtype = "cisco_asa_syslog"
	DeviceSubtypeAdvancedSflow          DeviceSubtype = "advanced_sflow"
	DeviceSubtypeA10Cgn                 DeviceSubtype = "a10_cgn"

	// for DeviceType = DeviceTypeHostNProbeDNSWWW
	DeviceSubtypeKprobe      DeviceSubtype = "kprobe"
	DeviceSubtypeNprobe      DeviceSubtype = "nprobe"
	DeviceSubtypeAwsSubnet   DeviceSubtype = "aws_subnet"
	DeviceSubtypeAzureSubnet DeviceSubtype = "azure_subnet"
	DeviceSubtypeGcpSubnet   DeviceSubtype = "gcp_subnet"
	DeviceSubtypeKappa       DeviceSubtype = "kappa"      // not in the API documentation
	DeviceSubtypeIbmSubnet   DeviceSubtype = "ibm_subnet" // not in the API documentation
)

type DeviceBGPType string

const (
	DeviceBGPTypeNone        DeviceBGPType = "none"
	DeviceBGPTypeDevice      DeviceBGPType = "device"
	DeviceBGPTypeOtherDevice DeviceBGPType = "other_device"
)

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
