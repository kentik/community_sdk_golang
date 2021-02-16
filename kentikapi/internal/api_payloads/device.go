package api_payloads

import (
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/utils"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

// GetAllDevicesResponse represents DevicesAPI GetAll JSON response
type GetAllDevicesResponse struct {
	Devices []DevicePayload `json:"devices"`
}

func (p GetAllDevicesResponse) ToDevices() (result []models.Device, err error) {
	convertFunc := func(d DevicePayload) (models.Device, error) {
		return d.ToDevice()
	}
	err = utils.ConvertList(p.Devices, convertFunc, &result)
	return result, err
}

// GetDeviceResponse represents DevicesAPI Get JSON response
type GetDeviceResponse struct {
	Device DevicePayload `json:"device"`
}

func (p GetDeviceResponse) ToDevice() (result models.Device, err error) {
	return p.Device.ToDevice()
}

// DevicePayload represents JSON Device payload as it is transmitted to and from KentikAPI
type DevicePayload struct {
	// following fields can appear in request: post/put, response: get/post/put
	PlanID                *models.ID         `json:"plan_id"`
	SiteID                *models.ID         `json:"site_id"`
	DeviceDescription     *string            `json:"device_description"`
	DeviceSampleRate      *int               `json:"device_sample_rate,string" request:"post" response:"get,post,put"`
	SendingIPS            []string           `json:"sending_ips"`
	DeviceSNMNPIP         *string            `json:"device_snmp_ip"`
	DeviceSNMPCommunity   *string            `json:"device_snmp_community"`
	MinimizeSNMP          *bool              `json:"minimize_snmp"`
	DeviceBGPType         *string            `json:"device_bgp_type"`
	DeviceBGPNeighborIP   *string            `json:"device_bgp_neighbor_ip"`
	DeviceBGPNeighborIPv6 *string            `json:"device_bgp_neighbor_ip6"`
	DeviceBGPNeighborASN  *string            `json:"device_bgp_neighbor_asn"`
	DeviceBGPFlowSpec     *bool              `json:"device_bgp_flowspec"`
	DeviceBGPPassword     *string            `json:"device_bgp_password"`
	UseBGPDeviceID        *models.ID         `json:"use_bgp_device_id"`
	DeviceSNMPv3Conf      *snmpv3ConfPayload `json:"device_snmp_v3_conf"`
	CDNAttr               *string            `json:"cdn_attr"`

	// following fields can appear in request: post, response: get/post/put
	DeviceName    *string `json:"device_name" request:"post" response:"get,post,put"`
	DeviceType    *string `json:"device_type" request:"post" response:"get,post,put"`
	DeviceSubType *string `json:"device_subtype" request:"post" response:"get,post,put"`

	// following fields can appear in request: none, response: get/post/put
	ID              *models.ID             `json:"id,string" response:"get,post,put"`
	Plan            *PlanPayload           `json:"plan" response:"get,post,put"`
	Site            *deviceSitePayload     `json:"site"`
	Labels          []deviceLabelPayload   `json:"labels"`
	AllInterfaces   []allInterfacesPayload `json:"all_interfaces"`
	DeviceStatus    *string                `json:"device_status"`
	DeviceFlowType  *string                `json:"device_flow_type"`
	CompanyID       *models.ID             `json:"company_id,string" response:"get,post,put"`
	SNMPLastUpdated *string                `json:"snmp_last_updated"`
	CreatedDate     *time.Time             `json:"created_date" response:"get,post,put"`
	UpdatedDate     *time.Time             `json:"updated_date" response:"get,post,put"`
	BGPPeerIP4      *string                `json:"bgpPeerIP4"`
	BGPPeerIP6      *string                `json:"bgpPeerIP6"`
}

// DeviceToPayload prepares POST/PUT request payload: fill only the user-provided fields
func DeviceToPayload(d models.Device) DevicePayload {
	var cdn *string = nil
	if d.CDNAttr != nil {
		s := d.CDNAttr.String()
		cdn = &s
	}
	return DevicePayload{
		CDNAttr: cdn,
		// more to come...
	}
}

// ToDevice transforms GET/POST/PUT response payload into Device
func (p DevicePayload) ToDevice() (models.Device, error) {
	deviceType, err := models.DeviceTypeString(*p.DeviceType)
	if err != nil {
		return models.Device{}, err
	}

	deviceSubtype, err := models.DeviceSubtypeString(*p.DeviceSubType)
	if err != nil {
		return models.Device{}, err
	}

	plan, err := p.Plan.ToPlan()
	if err != nil {
		return models.Device{}, err
	}

	var cdn *models.CDNAttribute
	if err := utils.ConvertOrNone(p.CDNAttr, models.CDNAttributeString, &cdn); err != nil {
		return models.Device{}, err
	}

	var snmp *models.SNMPv3Conf
	if err = utils.ConvertOrNone(p.DeviceSNMPv3Conf, payloadToSNMPv3Conf, &snmp); err != nil {
		return models.Device{}, err
	}

	var site *models.DeviceSite
	if err = utils.ConvertOrNone(p.Site, payloadToDeviceSite, &site); err != nil {
		return models.Device{}, err
	}

	var bgpType *models.DeviceBGPType
	err = utils.ConvertOrNone(p.DeviceBGPType, models.DeviceBGPTypeString, &bgpType)
	if err != nil {
		return models.Device{}, err
	}

	var labels []models.DeviceLabel
	err = utils.ConvertList(p.Labels, payloadToDeviceLabel, &labels)
	if err != nil {
		return models.Device{}, err
	}

	var allInterfaces []models.AllInterfaces
	err = utils.ConvertList(p.AllInterfaces, payloadToAllInterfaces, &allInterfaces)
	if err != nil {
		return models.Device{}, err
	}

	return models.Device{
		ID:                    *p.ID,
		CompanyID:             *p.CompanyID,
		CreatedDate:           *p.CreatedDate,
		UpdatedDate:           *p.UpdatedDate,
		DeviceName:            *p.DeviceName,
		DeviceType:            deviceType,
		DeviceSubType:         deviceSubtype,
		DeviceSampleRate:      *p.DeviceSampleRate,
		Plan:                  plan,
		SendingIPS:            p.SendingIPS,
		Site:                  site,
		PlanID:                p.PlanID,
		SiteID:                p.SiteID,
		Labels:                labels,
		AllInterfaces:         allInterfaces,
		CDNAttr:               cdn,
		DeviceDescription:     p.DeviceDescription,
		DeviceSNMNPIP:         p.DeviceSNMNPIP,
		DeviceSNMPCommunity:   p.DeviceSNMPCommunity,
		DeviceSNMPv3Conf:      snmp,
		MinimizeSNMP:          p.MinimizeSNMP,
		DeviceBGPType:         bgpType,
		DeviceBGPNeighborIP:   p.DeviceBGPNeighborIP,
		DeviceBGPNeighborIPv6: p.DeviceBGPNeighborIPv6,
		DeviceBGPNeighborASN:  p.DeviceBGPNeighborASN,
		DeviceBGPFlowSpec:     p.DeviceBGPFlowSpec,
		DeviceBGPPassword:     p.DeviceBGPPassword,
		UseBGPDeviceID:        p.UseBGPDeviceID,
		DeviceStatus:          p.DeviceStatus,
		DeviceFlowType:        p.DeviceFlowType,
		SNMPLastUpdated:       p.SNMPLastUpdated,
		BGPPeerIP4:            p.BGPPeerIP4,
		BGPPeerIP6:            p.BGPPeerIP6,
	}, nil
}

// allInterfacesPayload represents JSON Device.AllInterfaces payload as it is transmitted from KentikAPI
type allInterfacesPayload struct {
	DeviceID             models.ID `json:"device_id,string"`
	SNMPSpeed            float64   `json:"snmp_speed,string"`
	InterfaceDesctiption string    `json:"interface_description"`
	InitialSNMPSpeed     *float64  `json:"initial_snmp_speed,string"`
}

func payloadToAllInterfaces(p allInterfacesPayload) (models.AllInterfaces, error) {
	return models.AllInterfaces{
		InterfaceDescription: p.InterfaceDesctiption,
		DeviceID:             p.DeviceID,
		SNMPSpeed:            p.SNMPSpeed,
		InitialSNMPSpeed:     p.InitialSNMPSpeed,
	}, nil
}

// snmpv3ConfPayload represents JSON Device.SNMPv3Conf payload as it is transmitted to and from KentikAPI
type snmpv3ConfPayload struct {
	UserName                 string  `json:"UserName"`
	AuthenticationProtocol   *string `json:"AuthenticationProtocol"`
	AuthenticationPassphrase *string `json:"AuthenticationPassphrase"`
	PrivacyProtocol          *string `json:"PrivacyProtocol"`
	PrivacyPassphrase        *string `json:"PrivacyPassphrase"`
}

func payloadToSNMPv3Conf(p snmpv3ConfPayload) (models.SNMPv3Conf, error) {
	var auth *models.AuthenticationProtocol
	err := utils.ConvertOrNone(p.AuthenticationProtocol, models.AuthenticationProtocolString, &auth)
	if err != nil {
		return models.SNMPv3Conf{}, err
	}

	var priv *models.PrivacyProtocol
	err = utils.ConvertOrNone(p.PrivacyProtocol, models.PrivacyProtocolString, &priv)
	if err != nil {
		return models.SNMPv3Conf{}, err
	}

	return models.SNMPv3Conf{
		UserName:                 p.UserName,
		AuthenticationProtocol:   auth,
		AuthenticationPassphrase: p.AuthenticationPassphrase,
		PrivacyProtocol:          priv,
		PrivacyPassphrase:        p.PrivacyPassphrase,
	}, nil
}

// deviceLabelPayload represents JSON Device.Label payload as it is transmitted from KentikAPI.
// deviceLabelPayload embedded under Device differs from standalone LabelPayload in that it lacks devices list,
// and differs in field names, eg. cdate vs created_date, edate vs updated_date
type deviceLabelPayload struct {
	ID          models.ID  `json:"id"`
	Color       string     `json:"color"`
	Name        string     `json:"name"`
	UserID      *models.ID `json:"user_id,string"`
	CompanyID   models.ID  `json:"company_id,string"`
	CreatedDate time.Time  `json:"cdate"`
	UpdatedDate time.Time  `json:"edate"`
}

func payloadToDeviceLabel(p deviceLabelPayload) (models.DeviceLabel, error) {
	return models.DeviceLabel{
		ID:          p.ID,
		Name:        p.Name,
		Color:       p.Color,
		UserID:      p.UserID,
		CompanyID:   p.CompanyID,
		CreatedDate: p.CreatedDate,
		UpdatedDate: p.UpdatedDate,
	}, nil
}

// deviceSitePayload represents JSON Device.Site payload as it is transmitted from KentikAPI.
// deviceSitePayload embeddedd under Device differs from regular SitePayload in that all fields are optional.
type deviceSitePayload struct {
	CompanyID *models.ID `json:"company_id"`
	ID        *models.ID `json:"id"`
	Latitude  *float64   `json:"lat"`
	Longitude *float64   `json:"lon"`
	SiteName  *string    `json:"site_name"`
}

func payloadToDeviceSite(p deviceSitePayload) (models.DeviceSite, error) {
	return models.DeviceSite{
		ID:        p.ID,
		CompanyID: p.CompanyID,
		Latitude:  p.Latitude,
		Longitude: p.Longitude,
		SiteName:  p.SiteName,
	}, nil
}
