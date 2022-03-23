package models

import "time"

type DeviceLabel struct {
	// read-write properties (can be updated in update call)
	Name  string
	Color string

	// read-only properties (can't be updated in update call)
	Devices     []DeviceItem
	ID          ID
	UserID      *ID
	CompanyID   ID
	CreatedDate time.Time
	UpdatedDate time.Time
}

// DeviceLabelRequiredFields is subset of DeviceLabel fields required to create a DeviceLabel.
type DeviceLabelRequiredFields struct {
	Name  string
	Color string
}

// NewDeviceLabel creates a DeviceLabel with all necessary fields set.
func NewDeviceLabel(d DeviceLabelRequiredFields) *DeviceLabel {
	return &DeviceLabel{
		Name:  d.Name,
		Color: d.Color,
	}
}

type DeviceItem struct {
	ID            ID
	DeviceName    string
	DeviceSubtype string
	DeviceType    *string
}
