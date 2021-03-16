package models

import "time"

// Tag model.
type Tag struct {
	// read-write properties (can be updated in update call)
	FlowTag       string
	DeviceName    *string
	DeviceType    *string
	Site          *string
	InterfaceName *string
	Addr          *string
	Port          *string
	TCPFlags      *string
	Protocol      *string
	ASN           *string
	LasthopAsName *string
	NexthopAsn    *string
	NexthopAsName *string
	Nexthop       *string
	BGPAspath     *string
	BGPCommunity  *string
	MAC           *string
	Country       *string
	VLANs         *string

	// read-only properties (can't be updated in update call)
	ID          ID
	UserID      ID
	CompanyID   ID
	AddrCount   int
	MACCount    int
	EditedBy    string
	CreatedDate time.Time
	UpdatedDate time.Time
}

// NewTag creates a new Tag with all required fields set.
// Creating a tag requires specifying the FlowTag parameter, and also specifying at least one additional
// optional property that can be found in read-write group of Tag properties above.
func NewTag(flowTag string) *Tag {
	return &Tag{
		FlowTag: flowTag,
	}
}
