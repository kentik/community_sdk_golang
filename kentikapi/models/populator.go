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

// NewPopulator creates a Populator with all necessary fields set
func NewPopulator(value, deviceName string, direction PopulatorDirection) *Populator {
	return &Populator{
		Value:      value,
		DeviceName: deviceName,
		Direction:  direction,
	}
}

type PopulatorDirection int

const (
	PopulatorDirectionSrc    PopulatorDirection = iota // "SRC"
	PopulatorDirectionDst                              // "DST"
	PopulatorDirectionEither                           // "EITHER"
)
