package api_payloads

import (
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/utils"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

// CreateDeviceRequest represents DevicesAPI Create JSON request
type CreateDeviceRequest struct {
	Payload DevicePayload `json:"device"`
}

// UpdateDeviceRequest represents DevicesAPI Create JSON request
type UpdateDeviceRequest CreateDeviceRequest

// ApplyLabelsRequest represents DevicesAPI ApplyLabels JSON request
type ApplyLabelsRequest struct {
	Labels []labelIDPayload `json:"labels,omitempty"`
}

// GetAllDevicesResponse represents DevicesAPI GetAll JSON response
type GetAllDevicesResponse struct {
	Payload []DevicePayload `json:"devices"`
}

func (r GetAllDevicesResponse) ToDevices() (result []models.Device, err error) {
	err = utils.ConvertList(r.Payload, payloadToDevice, &result)
	return result, err
}

// GetDeviceResponse represents DevicesAPI Get JSON response
type GetDeviceResponse struct {
	Payload DevicePayload `json:"device"`
}

func (r GetDeviceResponse) ToDevice() (result models.Device, err error) {
	return payloadToDevice(r.Payload)
}

// CreateDeviceResponse represents DevicesAPI Create JSON response
type CreateDeviceResponse = GetDeviceResponse

// UpdateDeviceResponse represents DevicesAPI Update JSON response
type UpdateDeviceResponse = GetDeviceResponse

// ApplyLabelsResponse represents JSON ApplyLabelsResponse payload as it is transmitted from KentikAPI
type ApplyLabelsResponse struct {
	ID         models.ID            `json:"id,string"`
	DeviceName string               `json:"device_name"`
	Labels     []deviceLabelPayload `json:"labels"`
}

func (r *ApplyLabelsResponse) ToAppliedLabels() (models.AppliedLabels, error) {
	var labels []models.DeviceLabel
	if err := utils.ConvertList(r.Labels, payloadToDeviceLabel, &labels); err != nil {
		return models.AppliedLabels{}, err
	}
	return models.AppliedLabels{ID: r.ID, DeviceName: r.DeviceName, Labels: labels}, nil
}

// DevicePayload represents JSON Device payload as it is transmitted to and from KentikAPI
type DevicePayload struct {
	// following fields can appear in request: post/put, response: get/post/put
	PlanID                *models.ID         `json:"plan_id,omitempty" request:"post"`
	SiteID                *models.ID         `json:"site_id,omitempty"`
	DeviceDescription     *string            `json:"device_description,omitempty"`
	DeviceSampleRate      *int               `json:"device_sample_rate,string,omitempty" request:"post" response:"get,post,put"`
	SendingIPS            []string           `json:"sending_ips,omitempty"`
	DeviceSNMNPIP         *string            `json:"device_snmp_ip,omitempty"`
	DeviceSNMPCommunity   *string            `json:"device_snmp_community,omitempty"`
	MinimizeSNMP          *bool              `json:"minimize_snmp,omitempty"`
	DeviceBGPType         *string            `json:"device_bgp_type,omitempty"`
	DeviceBGPNeighborIP   *string            `json:"device_bgp_neighbor_ip,omitempty"`
	DeviceBGPNeighborIPv6 *string            `json:"device_bgp_neighbor_ip6,omitempty"`
	DeviceBGPNeighborASN  *string            `json:"device_bgp_neighbor_asn,omitempty"`
	DeviceBGPFlowSpec     *bool              `json:"device_bgp_flowspec,omitempty"`
	DeviceBGPPassword     *string            `json:"device_bgp_password,omitempty"`
	UseBGPDeviceID        *models.ID         `json:"use_bgp_device_id,omitempty"`
	DeviceSNMPv3Conf      *snmpv3ConfPayload `json:"device_snmp_v3_conf,omitempty"`
	CDNAttr               *string            `json:"cdn_attr,omitempty"`

	// following fields can appear in request: post, response: get/post/put
	DeviceName    *string `json:"device_name,omitempty" request:"post" response:"get,post,put"`
	DeviceType    *string `json:"device_type,omitempty" request:"post" response:"get,post,put"`
	DeviceSubType *string `json:"device_subtype,omitempty" request:"post" response:"get,post,put"`

	// following fields can appear in request: none, response: get/post/put
	ID              *models.ID             `json:"id,string,omitempty" response:"get,post,put"`
	Plan            *devicePlanPayload     `json:"plan,omitempty" response:"get,post,put"`
	Site            *deviceSitePayload     `json:"site,omitempty"`
	Labels          []deviceLabelPayload   `json:"labels,omitempty"`
	AllInterfaces   []allInterfacesPayload `json:"all_interfaces,omitempty"`
	DeviceStatus    *string                `json:"device_status,omitempty"`
	DeviceFlowType  *string                `json:"device_flow_type,omitempty"`
	CompanyID       *models.ID             `json:"company_id,string,omitempty" response:"get,post,put"`
	SNMPLastUpdated *string                `json:"snmp_last_updated,omitempty"`
	CreatedDate     *time.Time             `json:"created_date,omitempty" response:"get,post,put"`
	UpdatedDate     *time.Time             `json:"updated_date,omitempty" response:"get,post,put"`
	BGPPeerIP4      *string                `json:"bgpPeerIP4,omitempty"`
	BGPPeerIP6      *string                `json:"bgpPeerIP6,omitempty"`
}

// payloadToDevice transforms GET/POST/PUT response payload into Device
func payloadToDevice(p DevicePayload) (models.Device, error) {
	deviceType, err := models.DeviceTypeString(*p.DeviceType)
	if err != nil {
		return models.Device{}, err
	}

	deviceSubtype, err := models.DeviceSubtypeString(*p.DeviceSubType)
	if err != nil {
		return models.Device{}, err
	}

	plan, err := payloadToDevicePlan(*p.Plan)
	if err != nil {
		return models.Device{}, err
	}

	var cdn *models.CDNAttribute
	if err := utils.ConvertOrNone(p.CDNAttr, models.CDNAttributeString, &cdn); err != nil {
		return models.Device{}, err
	}

	snmp, err := payloadToSNMPv3Conf(p.DeviceSNMPv3Conf)
	if err != nil {
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
		Site:                  payloadToDeviceSite(p.Site),
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

// DeviceToPayload prepares POST/PUT request payload: fill only the user-provided fields
func DeviceToPayload(d models.Device) (DevicePayload, error) {
	deviceType := d.DeviceType.String()
	deviceSubtype := d.DeviceSubType.String()

	var cdn *string
	if d.CDNAttr != nil {
		s := d.CDNAttr.String()
		cdn = &s
	}

	var bgpType *string
	if d.DeviceBGPType != nil {
		bgpType = new(string)
		*bgpType = d.DeviceBGPType.String()
	}

	return DevicePayload{
		DeviceName:            &d.DeviceName,
		DeviceType:            &deviceType,
		DeviceSubType:         &deviceSubtype,
		DeviceSampleRate:      &d.DeviceSampleRate,
		SendingIPS:            d.SendingIPS,
		PlanID:                d.PlanID,
		SiteID:                d.SiteID,
		CDNAttr:               cdn,
		DeviceDescription:     d.DeviceDescription,
		DeviceSNMNPIP:         d.DeviceSNMNPIP,
		DeviceSNMPCommunity:   d.DeviceSNMPCommunity,
		DeviceSNMPv3Conf:      snmp3ConfToPayload(d.DeviceSNMPv3Conf),
		MinimizeSNMP:          d.MinimizeSNMP,
		DeviceBGPType:         bgpType,
		DeviceBGPNeighborIP:   d.DeviceBGPNeighborIP,
		DeviceBGPNeighborIPv6: d.DeviceBGPNeighborIPv6,
		DeviceBGPNeighborASN:  d.DeviceBGPNeighborASN,
		DeviceBGPFlowSpec:     d.DeviceBGPFlowSpec,
		DeviceBGPPassword:     d.DeviceBGPPassword,
		UseBGPDeviceID:        d.UseBGPDeviceID,
	}, nil
}

// allInterfacesPayload represents JSON Device.AllInterfaces payload as it is transmitted from KentikAPI
type allInterfacesPayload struct {
	DeviceID             models.ID `json:"device_id,string,omitempty"`
	SNMPSpeed            float64   `json:"snmp_speed,string,omitempty"`
	InterfaceDesctiption string    `json:"interface_description,omitempty"`
	InitialSNMPSpeed     *float64  `json:"initial_snmp_speed,string,omitempty"`
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
	UserName                 string  `json:"UserName,omitempty"`
	AuthenticationProtocol   *string `json:"AuthenticationProtocol,omitempty"`
	AuthenticationPassphrase *string `json:"AuthenticationPassphrase,omitempty"`
	PrivacyProtocol          *string `json:"PrivacyProtocol,omitempty"`
	PrivacyPassphrase        *string `json:"PrivacyPassphrase,omitempty"`
}

func payloadToSNMPv3Conf(p *snmpv3ConfPayload) (*models.SNMPv3Conf, error) {
	if p == nil {
		return nil, nil
	}

	var auth *models.AuthenticationProtocol
	err := utils.ConvertOrNone(p.AuthenticationProtocol, models.AuthenticationProtocolString, &auth)
	if err != nil {
		return nil, err
	}

	var priv *models.PrivacyProtocol
	err = utils.ConvertOrNone(p.PrivacyProtocol, models.PrivacyProtocolString, &priv)
	if err != nil {
		return nil, err
	}

	return &models.SNMPv3Conf{
		UserName:                 p.UserName,
		AuthenticationProtocol:   auth,
		AuthenticationPassphrase: p.AuthenticationPassphrase,
		PrivacyProtocol:          priv,
		PrivacyPassphrase:        p.PrivacyPassphrase,
	}, nil
}

func snmp3ConfToPayload(d *models.SNMPv3Conf) *snmpv3ConfPayload {
	if d == nil {
		return nil
	}

	var auth *string
	if d.AuthenticationProtocol != nil {
		auth = new(string)
		*auth = d.AuthenticationProtocol.String()
	}
	var priv *string
	if d.PrivacyProtocol != nil {
		priv = new(string)
		*priv = d.PrivacyProtocol.String()
	}

	return &snmpv3ConfPayload{
		UserName:                 d.UserName,
		AuthenticationProtocol:   auth,
		AuthenticationPassphrase: d.AuthenticationPassphrase,
		PrivacyProtocol:          priv,
		PrivacyPassphrase:        d.PrivacyPassphrase,
	}
}

// deviceLabelPayload represents JSON Device.Label payload as it is transmitted from KentikAPI.
// deviceLabelPayload embedded under Device differs from standalone LabelPayload in that it lacks devices list,
// and differs in field names, eg. cdate vs created_date, edate vs updated_date
type deviceLabelPayload struct {
	ID          models.ID  `json:"id"`
	Color       string     `json:"color"`
	Name        string     `json:"name"`
	UserID      *models.ID `json:"user_id,string,omitempty"`
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

func deviceLabelToPayload(d models.DeviceLabel) (deviceLabelPayload, error) {
	return deviceLabelPayload{
		ID:          d.ID,
		Name:        d.Name,
		Color:       d.Color,
		UserID:      d.UserID,
		CompanyID:   d.CompanyID,
		CreatedDate: d.CreatedDate,
		UpdatedDate: d.UpdatedDate,
	}, nil
}

// deviceSitePayload represents JSON Device.Site payload as it is transmitted from KentikAPI.
// deviceSitePayload embeddedd under Device differs from regular SitePayload in that all fields are optional.
type deviceSitePayload struct {
	CompanyID *models.ID `json:"company_id,omitempty"`
	ID        *models.ID `json:"id,omitempty"`
	Latitude  *float64   `json:"lat,omitempty"`
	Longitude *float64   `json:"lon,omitempty"`
	SiteName  *string    `json:"site_name,omitempty"`
}

func payloadToDeviceSite(p *deviceSitePayload) *models.DeviceSite {
	return &models.DeviceSite{
		ID:        p.ID,
		CompanyID: p.CompanyID,
		Latitude:  p.Latitude,
		Longitude: p.Longitude,
		SiteName:  p.SiteName,
	}
}

// devicePlanPayload represents JSON Device.Plan payload as it is transmitted from KentikAPI.
// devicePlanPayload embeddedd under Device differs from regular PlanPayload in that all fields are optional.
type devicePlanPayload struct {
	// following fields can appear in request: none, response: get
	ID            *models.ID              `json:"id,omitempty"`
	CompanyID     *models.ID              `json:"company_id,omitempty"`
	Name          *string                 `json:"name,omitempty"`
	Description   *string                 `json:"description,omitempty"`
	Active        *bool                   `json:"active,omitempty"`
	MaxDevices    *int                    `json:"max_devices,omitempty"`
	MaxFPS        *int                    `json:"max_fps,omitempty"`
	BGPEnabled    *bool                   `json:"bgp_enabled,omitempty"`
	FastRetention *int                    `json:"fast_retention,omitempty"`
	FullRetention *int                    `json:"full_retention,omitempty"`
	CreatedDate   *time.Time              `json:"cdate,omitempty"`
	UpdatedDate   *time.Time              `json:"edate,omitempty"`
	MaxBigdataFPS *int                    `json:"max_bigdata_fps,omitempty"`
	DeviceTypes   []planDeviceTypePayload `json:"deviceTypes"`
	Devices       []planDevicePayload     `json:"devices"`
}

func payloadToDevicePlan(p devicePlanPayload) (models.DevicePlan, error) {
	var deviceTypes []models.PlanDeviceType
	err := utils.ConvertList(p.DeviceTypes, payloadToPlanDeviceType, &deviceTypes)
	if err != nil {
		return models.DevicePlan{}, err
	}

	var devices []models.PlanDevice
	err = utils.ConvertList(p.Devices, payloadToPlanDevice, &devices)
	if err != nil {
		return models.DevicePlan{}, err
	}

	return models.DevicePlan{
		ID:            p.ID,
		CompanyID:     p.CompanyID,
		Name:          p.Name,
		Description:   p.Description,
		Active:        p.Active,
		MaxDevices:    p.MaxDevices,
		MaxFPS:        p.MaxFPS,
		BGPEnabled:    p.BGPEnabled,
		FastRetention: p.FastRetention,
		FullRetention: p.FullRetention,
		CreatedDate:   p.CreatedDate,
		UpdatedDate:   p.UpdatedDate,
		MaxBigdataFPS: p.MaxBigdataFPS,
		DeviceTypes:   deviceTypes,
		Devices:       devices,
	}, nil
}

// labelIDPayload represents JSON ApplyLabels.LabelID payload as it is transmitted to KentikAPI
type labelIDPayload struct {
	ID int `json:"id"`
}

func LabelIDsToPayload(ids []models.ID) []labelIDPayload {
	result := make([]labelIDPayload, 0, len(ids))
	for _, id := range ids {
		result = append(result, labelIDPayload{ID: id})
	}
	return result
}
