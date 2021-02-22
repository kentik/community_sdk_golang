package api_payloads

import (
	"encoding/json"
	"fmt"
	"strconv"
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

// CreateInterfaceRequest represents DevicesAPI.InterfacesAPI Create JSON request
type CreateInterfaceRequest InterfacePayload

// CreateInterfaceResponse represents DevicesAPI.InterfacesAPI Create JSON response
type CreateInterfaceResponse InterfacePayload

func (r CreateInterfaceResponse) ToInterface() (result models.Interface, err error) {
	return payloadToInterface(InterfacePayload(r))
}

// InterfacePayload represents JSON Interface payload as it is transmitted to and from KentikAPI
type InterfacePayload struct {
	// following fields can appear in request: post/put, response: get/post/put
	SNMPID               *models.ID            `json:"snmp_id,string,omitempty"`
	SNMPSpeed            interface{}           `json:"snmp_speed"` // caveat, GET returns snmp_speed as string but POST and PUT as int; handle manually
	InterfaceDescription *string               `json:"interface_description,omitempty"`
	SNMPAlias            *string               `json:"snmp_alias,omitempty"`
	InterfaceIP          *string               `json:"interface_ip,omitempty"`
	InterfaceIPNetmask   *string               `json:"interface_ip_netmask,omitempty"`
	VRF                  *vrfAttributesPayload `json:"vrf,omitempty"`           // caveat, GET returns vrf as valid object, but POST and PUT as empty object
	VRFID                interface{}           `json:"vrf_id,string,omitempty"` // caveat, GET returns snmp_speed as string but POST and PUT as int and it is optional; handle manually
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

	// "snmp_speed" is returned as string for get, but as int for post/put
	speed, err := stringOrNumberToInt(p.SNMPSpeed)
	if err != nil {
		return models.Interface{}, err
	}

	// "vrf_id" is returned as string for get, but as int for post/put. And it is optional
	var vrfID *int
	if p.VRFID != nil {
		vrfID = new(int)
		*vrfID, err = stringOrNumberToInt(p.VRFID)
		if err != nil {
			return models.Interface{}, err
		}
	}

	return models.Interface{
		SNMPID:               *p.SNMPID,
		SNMPSpeed:            speed,
		SNMPAlias:            p.SNMPAlias,
		InterfaceDescription: p.InterfaceDescription,
		InterfaceIP:          p.InterfaceIP,
		InterfaceIPNetmask:   p.InterfaceIPNetmask,
		VRFID:                vrfID,
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

// InterfaceToPayload prepares POST/PUT request payload: fill only the user-provided fields
func InterfaceToPayload(i models.Interface) (InterfacePayload, error) {
	var vrf *vrfAttributesPayload
	err := utils.ConvertOrNone(i.VRF, vrfAttributesToPayload, &vrf)
	if err != nil {
		return InterfacePayload{}, err
	}

	var secondaryIPs []secondaryIPPayload
	err = utils.ConvertList(i.SecondaryIPS, secondaryIPToPayload, &secondaryIPs)
	if err != nil {
		return InterfacePayload{}, err
	}

	speed := strconv.Itoa(i.SNMPSpeed)
	return InterfacePayload{
		SNMPID:               &i.SNMPID,
		SNMPSpeed:            &speed,
		InterfaceDescription: i.InterfaceDescription,
		SNMPAlias:            i.SNMPAlias,
		InterfaceIP:          i.InterfaceIP,
		InterfaceIPNetmask:   i.InterfaceIPNetmask,
		VRF:                  vrf,
		VRFID:                i.VRFID,
		SecondaryIPs:         secondaryIPs,
	}, nil
}

// vrfAttributesPayload represents JSON Interface.VRFAttributes payload as it is transmitted to and from KentikAPI
// Note: it is returned only in get response, for post and put responses empty object is returned but vrf_id is set
type vrfAttributesPayload struct {
	// following fields can appear in request: post/put, response: get
	Name               string  `json:"name"`
	RouteTarget        string  `json:"route_target"`
	RouteDistinguisher string  `json:"route_distinguisher"`
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

// vrfAttributesToPayload prepares POST/PUT request payload: fill only the user-provided fields
func vrfAttributesToPayload(a models.VRFAttributes) (vrfAttributesPayload, error) {
	return vrfAttributesPayload{
		Name:                  a.Name,
		RouteTarget:           a.RouteTarget,
		RouteDistinguisher:    a.RouteDistinguisher,
		Description:           a.Description,
		ExtRouteDistinguisher: a.ExtRouteDistinguisher,
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

// secondaryIPToPayload prepares POST/PUT request payload: fill only the user-provided fields
func secondaryIPToPayload(s models.SecondaryIP) (secondaryIPPayload, error) {
	return secondaryIPPayload{
		Address: s.Address,
		Netmask: s.Netmask,
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

// special treatment of InterfacePayload.SNMPSpeed which sometime comes as string, sometimes as float
func stringOrNumberToInt(i interface{}) (result int, err error) {
	switch val := i.(type) {
	case string:
		result, err = strconv.Atoi(val)
		if err != nil {
			return 0, fmt.Errorf("stringOrFloatToInt Atoi conversion: %v", err)
		}
	case float64: // json.Unmarshall recognizes numbers as float64
		return int(val), nil
	default:
		return 0, fmt.Errorf("stringOrFloatToInt input should be string or float64, got {%T}", i)
	}

	return
}
