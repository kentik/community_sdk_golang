package api_payloads

import (
	"encoding/json"
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/utils"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

// Note: InterfacesAPI belong under DevicesAPI but it is wast so it lives in a separate file

// GetInterfaceResponse represents DevicesAPI.InterfacesAPI GetAll JSON response
type GetAllInterfacesResponse []InterfacePayload

func (p GetAllInterfacesResponse) ToInterfaces() (result []models.Interface, err error) {
	err = utils.ConvertList(p, payloadToInterface, &result)
	return result, err
}

// GetInterfaceResponse represents DevicesAPI.InterfacesAPI Get JSON response
type GetInterfaceResponse struct {
	Interface InterfacePayload `json:"interface"`
}

func (r GetInterfaceResponse) ToInterface() (result models.Interface, err error) {
	return payloadToInterface(r.Interface)
}

// InterfacePayload represents JSON Interface payload as it is transmitted to and from KentikAPI
type InterfacePayload struct {
	// following fields can appear in request: post/put, response: get/post/put
	SNMPID               *models.ID            `json:"snmp_id,string,omitempty"`
	SNMPSpeed            *int                  `json:"snmp_speed,string,omitempty"` // caveat, GET returns snmp_speed as string but POST and PUT as int
	InterfaceDescription *string               `json:"interface_description,omitempty"`
	SNMPAlias            *string               `json:"snmp_alias,omitempty"`
	InterfaceIP          *string               `json:"interface_ip,omitempty"`
	InterfaceIPNetmask   *string               `json:"interface_ip_netmask,omitempty"`
	VRF                  *vrfAttributesPayload `json:"vrf,omitempty"`           // special case; vrf is returned only in get response
	VRFID                *models.ID            `json:"vrf_id,string,omitempty"` // and vrf_id is returned normally in get/post/put
	SecondaryIPs         []secondaryIPPayload  `json:"secondary_ips,omitempty"`

	// following fields can appear in request: none, response: get/post/put
	ID                          *models.ID             `json:"id,string,omitempty" response:"get,post,put"`
	CompanyID                   *models.ID             `json:"company_id,string,omitempty"`
	DeviceID                    *models.ID             `json:"device_id,string,omitempty"`
	CreatedDate                 *time.Time             `json:"cdate,omitempty"`
	UpdatedDate                 *time.Time             `json:"edate,omitempty"`
	InitialSNMPID               *string                `json:"initial_snmp_id,omitempty"` // API happens to return empty string ""
	InitialSNMPAlias            *string                `json:"initial_snmp_alias,omitempty"`
	InitialInterfaceDescription *string                `json:"initial_interface_description,omitempty"`
	InitialSNMPSpeed            *int                   `json:"initial_snmp_speed,string,omitempty"`
	TopNexthopASNs              []topNextHopASNPayload `json:"top_nexthop_asns,omitempty"`
}

func (p *InterfacePayload) UnmarshalJSON(data []byte) error {
	type tmp InterfacePayload
	if err := json.Unmarshal(data, (*tmp)(p)); err != nil {
		return err
	}

	// postprocessing
	// API returns non-set VRF as empty object which presumes all VRF fields must be optional.
	// make empty VRF field a nil so no need to make everything optional
	var emptyVRF vrfAttributesPayload
	if p.VRF != nil && *p.VRF == emptyVRF {
		p.VRF = nil
	}
	return nil
}

func payloadToInterface(p InterfacePayload) (models.Interface, error) {
	var err error

	var vrf *models.VRFAttributes
	err = utils.ConvertOrNone(p.VRF, payloadToVRFAttributes, &vrf)
	if err != nil {
		return models.Interface{}, err
	}

	var secondaryIPs []models.SecondaryIP
	err = utils.ConvertList(p.SecondaryIPs, payloadToSecondaryIP, &secondaryIPs)
	if err != nil {
		return models.Interface{}, err
	}

	var topNextHopASNs []models.TopNextHopASN
	err = utils.ConvertList(p.TopNexthopASNs, payloadToTopNextHopASN, &topNextHopASNs)
	if err != nil {
		return models.Interface{}, err
	}

	return models.Interface{
		SNMPID:               *p.SNMPID,
		SNMPSpeed:            *p.SNMPSpeed,
		SNMPAlias:            p.SNMPAlias,
		InterfaceDescription: p.InterfaceDescription,
		InterfaceIP:          p.InterfaceIP,
		InterfaceIPNetmask:   p.InterfaceIPNetmask,
		VRFID:                p.VRFID,
		VRF:                  vrf,
		SecondaryIPS:         secondaryIPs,

		ID:                          *p.ID,
		CompanyID:                   *p.CompanyID,
		DeviceID:                    *p.DeviceID,
		CreatedDate:                 *p.CreatedDate,
		UpdatedDate:                 *p.UpdatedDate,
		InitialSNMPID:               p.InitialSNMPID,
		InitialSNMPAlias:            p.InitialSNMPAlias,
		InitialInterfaceDescription: p.InitialInterfaceDescription,
		InitialSNMPSpeed:            p.InitialSNMPSpeed,
		TopNextHopASNs:              topNextHopASNs,
	}, nil
}

// vrfAttributesPayload represents JSON Interface.VRFAttributes payload as it is transmitted to and from KentikAPI
// Note: it is returned only in get response, for post and put responses empty object is returned but vrf_id is set
type vrfAttributesPayload struct {
	// following fields can appear in request: post/put, response: get
	Name               string  `json:"name" request:"post" response:"get"`
	RouteTarget        string  `json:"route_target" request:"post" response:"get"`
	RouteDistinguisher string  `json:"route_distinguisher" request:"post" response:"get"`
	Description        *string `json:"description,omitempty"`

	// following fields can appear in request: post/put, response: none
	ExtRouteDistinguisher *string `json:"ext_route_distinguisher,omitempty"` // not returned in any response

	// following fields can appear in request: none, response: get
	ID        *models.ID `json:"id,omitempty" response:"get"`
	CompanyID *models.ID `json:"company_id,string,omitempty" response:"get"`
	DeviceID  *models.ID `json:"device_id,string,omitempty" response:"get"`
}

func payloadToVRFAttributes(p vrfAttributesPayload) (models.VRFAttributes, error) {
	return models.VRFAttributes{
		Name:                  p.Name,
		RouteTarget:           p.RouteTarget,
		RouteDistinguisher:    p.RouteDistinguisher,
		Description:           p.Description,
		ExtRouteDistinguisher: p.ExtRouteDistinguisher,
		ID:                    *p.ID,
		CompanyID:             *p.CompanyID,
		DeviceID:              *p.DeviceID,
	}, nil
}

// secondaryIPPayload represents JSON Interface.SecondaryIPPayload payload as it is transmitted to and from KentikAPI
type secondaryIPPayload struct {
	// following fields can appear in request: post/put, response: get/post/put
	Address string `json:"address"`
	Netmask string `json:"netmask"`
}

func payloadToSecondaryIP(p secondaryIPPayload) (models.SecondaryIP, error) {
	return models.SecondaryIP{
		Address: p.Address,
		Netmask: p.Netmask,
	}, nil
}

// topNextHopASNPayload represents JSON Interface.TopNextHopASNPayload payload as it is transmitted from KentikAPI
type topNextHopASNPayload struct {
	// following fields can appear in request: post/put, response: get/post/put
	ASN     int `json:"ASN"`
	Packets int `json:"packets"`
}

func payloadToTopNextHopASN(p topNextHopASNPayload) (models.TopNextHopASN, error) {
	return models.TopNextHopASN{
		ASN:     p.ASN,
		Packets: p.Packets,
	}, nil
}
