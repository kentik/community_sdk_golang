package api_payloads

import (
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

// PopulatorPayload represents JSON Populator payload as it is transmitted to and from KentikAPI
type PopulatorPayload struct {
	// following fields can appear in request: post/put, response: get/post/put
	Direction     string  `json:"direction"`   // direction is always required
	Value         string  `json:"value"`       // value is always required
	DeviceName    string  `json:"device_name"` // device_name is always required
	InterfaceName *string `json:"interface_name"`
	Addr          *string `json:"addr"`
	Port          *string `json:"port"`
	TCPFlags      *string `json:"tcp_flags"`
	Protocol      *string `json:"protocol"`
	ASN           *string `json:"asn"`
	NextHopASN    *string `json:"nexthop_asn"`
	NextHop       *string `json:"nexthop"`
	BGPAsPath     *string `json:"bgp_aspath"`
	BGPCommunity  *string `json:"bgp_community"`
	DeviceType    *string `json:"device_type"`
	Site          *string `json:"site"`
	LastHopAsName *string `json:"lasthop_as_name"`
	NextHopAsName *string `json:"nexthop_as_name"`
	MAC           *string `json:"mac"`
	Country       *string `json:"country"`
	VLans         *string `json:"vlans"`

	// following fields can appear in request: none, response: get/post/put
	ID          *models.ID `json:"id" response:"get,post,put"`
	DimensionID *models.ID `json:"dimension_id" response:"get,post,put"`
	CompanyID   *models.ID `json:"company_id,string" response:"get,post,put"`
	User        *string    `json:"user"` // not always returned
	MACCount    *int       `json:"mac_count" response:"get,post,put"`
	AddrCount   *int       `json:"addr_count" response:"get,post,put"`
	CreatedDate *time.Time `json:"created_date" response:"get,post,put"`
	UpdatedDate *time.Time `json:"updated_date" response:"get,post,put"`
}

// payloadToPopulator transforms GET/POST/PUT response payload into Populator
func payloadToPopulator(p PopulatorPayload) (models.Populator, error) {
	direction, err := models.PopulatorDirectionString(p.Direction)
	if err != nil {
		return models.Populator{}, err
	}
	return models.Populator{
		Direction:     direction,
		Value:         p.Value,
		DeviceName:    p.DeviceName,
		InterfaceName: p.InterfaceName,
		Addr:          p.Addr,
		Port:          p.Port,
		TCPFlags:      p.TCPFlags,
		Protocol:      p.Protocol,
		ASN:           p.ASN,
		NextHopASN:    p.NextHopASN,
		NextHop:       p.NextHop,
		BGPAsPath:     p.BGPAsPath,
		BGPCommunity:  p.BGPCommunity,
		DeviceType:    p.DeviceType,
		Site:          p.Site,
		LastHopAsName: p.LastHopAsName,
		NextHopAsName: p.NextHopAsName,
		MAC:           p.MAC,
		Country:       p.Country,
		VLans:         p.VLans,
		ID:            *p.ID,
		DimensionID:   *p.DimensionID,
		CompanyID:     *p.CompanyID,
		User:          p.User,
		MACCount:      *p.MACCount,
		AddrCount:     *p.AddrCount,
		CreatedDate:   *p.CreatedDate,
		UpdatedDate:   *p.UpdatedDate,
	}, nil
}
