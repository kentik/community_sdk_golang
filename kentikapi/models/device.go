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
	DeviceBGPType         *DeviceBGPType
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
	Plan            Plan
	Site            *DeviceSite
	Labels          []DeviceLabel
	AllInterfaces   []AllInterfaces
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

// DeviceSite embedded under Device differs from regular Site in that all fields are optional
type DeviceSite struct {
	ID        *ID
	CompanyID *ID
	Latitude  *float64
	Longitude *float64
	SiteName  *string
}

type DeviceType int

const (
	DeviceTypeRouter           DeviceType = iota // "router"
	DeviceTypeHostNProbeDNSWWW                   // "host-nprobe-dns-www"
)

type DeviceSubtype int

const (
	// for DeviceType = DeviceTypeRouter
	DeviceSubtypeRouter                 DeviceSubtype = iota // "router"
	DeviceSubtypeCiscoAsa                                    // "cisco_asa"
	DeviceSubtypePaloalto                                    // "paloalto"
	DeviceSubtypeSilverpeak                                  // "silverpeak"
	DeviceSubtypeMpls                                        // "mpls"
	DeviceSubtypeViptela                                     // "viptela"
	DeviceSubtypePfeSyslog                                   // "pfe_syslog"
	DeviceSubtypeSyslog                                      // "syslog"
	DeviceSubtypeMeraki                                      // "meraki"
	DeviceSubtypeIstio                                       // "istio"
	DeviceSubtypeIosxr                                       // "ios_xr"
	DeviceSubtypeCiscoZoneBasedFirewall                      // "cisco_zone_based_firewall"
	DeviceSubtypeCiscoNbar                                   // "cisco_nbar"
	DeviceSubtypeCiscoAsaSyslog                              // "cisco_asa_syslog"
	DeviceSubtypeAdvancedSflow                               // "advanced_sflow"
	DeviceSubtypeA10Cgn                                      // "a10_cgn"

	// for DeviceType = DeviceTypeHostNProbeDNSWWW
	DeviceSubtypeKprobe      // "kprobe"
	DeviceSubtypeNprobe      // "nprobe"
	DeviceSubtypeAwsSubnet   // "aws_subnet"
	DeviceSubtypeAzureSubnet // "azure_subnet"
	DeviceSubtypeGcpSubnet   // "gcp_subnet"
	DeviceSubtypeKappa       // "kappa"  # not in api documentation
	DeviceSubtypeIbmSubnet   // "ibm_subnet"  # not in api documentation
)

type DeviceBGPType int

const (
	DeviceBGPTypeNone        DeviceBGPType = iota // "none"
	DeviceBGPTypeDevice                           // "device
	DeviceBGPTypeOtherDevice                      // "other_device"
)

type AuthenticationProtocol int

const (
	AuthenticationProtocolNoAuth AuthenticationProtocol = iota // "NoAuth"
	AuthenticationProtocolMD5                                  // "MD5"
	AuthenticationProtocolSHA                                  // "SHA"
)

type PrivacyProtocol int

const (
	PrivacyProtocolNoPriv PrivacyProtocol = iota // "NoPriv"
	PrivacyProtocolDES                           // "DES"
	PrivacyProtocolAES                           // "AES"
)

type CDNAttribute int

const (
	CDNAttributeNone CDNAttribute = iota // "None"
	CDNAttributeYes                      // "Y""
	CDNAttributeNo                       // "N""
)
