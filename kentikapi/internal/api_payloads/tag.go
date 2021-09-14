package api_payloads

import (
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type GetAllTagsResponse struct {
	Tags []tagPayload `json:"tags"`
}

func (r GetAllTagsResponse) ToTags() []models.Tag {
	var tags []models.Tag
	for _, up := range r.Tags {
		tags = append(tags, *up.ToTag())
	}
	return tags
}

type GetTagResponse struct {
	Tag tagPayload `json:"tag"`
}

func (r GetTagResponse) ToTag() *models.Tag {
	return r.Tag.ToTag()
}

type CreateTagRequest struct {
	Tag tagPayload `json:"tag"`
}

type CreateTagResponse = GetTagResponse

type UpdateTagRequest = CreateTagRequest

type UpdateTagResponse = GetTagResponse

type tagPayload struct {
	// following fields can appear in request: post/put, response: get/post/put
	FlowTag       string  `json:"flow_tag"`
	DeviceName    *string `json:"device_name,omitempty"`
	DeviceType    *string `json:"device_type,omitempty"`
	Site          *string `json:"site,omitempty"`
	InterfaceName *string `json:"interface_name,omitempty"`
	Addr          *string `json:"addr,omitempty"`
	Port          *string `json:"port,omitempty"`
	TCPFlags      *string `json:"tcp_flags,omitempty"`
	Protocol      *string `json:"protocol,omitempty"`
	ASN           *string `json:"asn,omitempty"`
	LasthopAsName *string `json:"lasthop_as_name,omitempty"`
	NexthopAsn    *string `json:"nexthop_asn,omitempty"`
	NexthopAsName *string `json:"nexthop_as_name,omitempty"`
	Nexthop       *string `json:"nexthop,omitempty"`
	BGPAspath     *string `json:"bgp_aspath,omitempty"`
	BGPCommunity  *string `json:"bgp_community,omitempty"`
	MAC           *string `json:"mac,omitempty"`
	Country       *string `json:"country,omitempty"`
	VLANs         *string `json:"vlans,omitempty"`

	// following fields can appear in request: none, response: get/post/put
	ID          IntAsString `json:"id,omitempty"`
	UserID      IntAsString `json:"user,omitempty"`
	CompanyID   IntAsString `json:"company_id,omitempty"`
	AddrCount   int         `json:"addr_count,omitempty"`
	MACCount    int         `json:"mac_count,omitempty"`
	EditedBy    string      `json:"edited_by,omitempty"`
	CreatedDate *time.Time  `json:"created_date,omitempty" response:"get,post,put"`
	UpdatedDate *time.Time  `json:"updated_date,omitempty" response:"get,post,put"`
}

func (p tagPayload) ToTag() *models.Tag {
	return &models.Tag{
		FlowTag:       p.FlowTag,
		DeviceName:    p.DeviceName,
		DeviceType:    p.DeviceType,
		Site:          p.Site,
		InterfaceName: p.InterfaceName,
		Addr:          p.Addr,
		Port:          p.Port,
		TCPFlags:      p.TCPFlags,
		Protocol:      p.Protocol,
		ASN:           p.ASN,
		LasthopAsName: p.LasthopAsName,
		NexthopAsn:    p.NexthopAsn,
		NexthopAsName: p.NexthopAsName,
		Nexthop:       p.Nexthop,
		BGPAspath:     p.BGPAspath,
		BGPCommunity:  p.BGPCommunity,
		MAC:           p.MAC,
		Country:       p.Country,
		VLANs:         p.VLANs,
		ID:            models.ID(p.ID),
		UserID:        models.ID(p.UserID),
		CompanyID:     models.ID(p.CompanyID),
		AddrCount:     p.AddrCount,
		MACCount:      p.MACCount,
		EditedBy:      p.EditedBy,
		CreatedDate:   *p.CreatedDate,
		UpdatedDate:   *p.UpdatedDate,
	}
}

// TagToPayload prepares POST/PUT request payload: fill only the user-provided fields.
//nolint:revive // tagPayLoad doesn't need to be exported
func TagToPayload(u models.Tag) tagPayload {
	return tagPayload{
		FlowTag:       u.FlowTag,
		DeviceName:    u.DeviceName,
		DeviceType:    u.DeviceType,
		Site:          u.Site,
		InterfaceName: u.InterfaceName,
		Addr:          u.Addr,
		Port:          u.Port,
		TCPFlags:      u.TCPFlags,
		Protocol:      u.Protocol,
		ASN:           u.ASN,
		LasthopAsName: u.LasthopAsName,
		NexthopAsn:    u.NexthopAsn,
		NexthopAsName: u.NexthopAsName,
		Nexthop:       u.Nexthop,
		BGPAspath:     u.BGPAspath,
		BGPCommunity:  u.BGPCommunity,
		MAC:           u.MAC,
		Country:       u.Country,
		VLANs:         u.VLANs,
	}
}
