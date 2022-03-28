package models

import "time"

type Populator struct {
	// read-write properties (can be updated in update call)
	Value         string
	Direction     PopulatorDirection
	DeviceName    string
	InterfaceName *string
	Addr          *string
	Port          *string
	TCPFlags      *string
	Protocol      *string
	ASN           *string
	NextHopASN    *string
	NextHop       *string
	BGPAsPath     *string
	BGPCommunity  *string
	DeviceType    *string
	Site          *string
	LastHopAsName *string
	NextHopAsName *string
	MAC           *string
	Country       *string
	VLans         *string

	// read-only properties (can't be updated in update call)
	ID          ID
	CompanyID   ID
	DimensionID ID
	User        *string
	MACCount    int
	AddrCount   int
	CreatedDate time.Time
	UpdatedDate time.Time
}

// PopulatorRequiredFields is subset of Populator fields required to create a Populator.
type PopulatorRequiredFields struct {
	DimensionID ID
	Value       string
	DeviceName  string
	Direction   PopulatorDirection
}

// NewPopulator creates a Populator with all necessary fields set.
func NewPopulator(p PopulatorRequiredFields) *Populator {
	return &Populator{
		DimensionID: p.DimensionID,
		Value:       p.Value,
		DeviceName:  p.DeviceName,
		Direction:   p.Direction,
	}
}

type PopulatorDirection string

const (
	PopulatorDirectionSrc    PopulatorDirection = "SRC"
	PopulatorDirectionDst    PopulatorDirection = "DST"
	PopulatorDirectionEither PopulatorDirection = "EITHER"
)
