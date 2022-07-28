package models

import "time"

type DeviceLabel struct {
	// Read-write properties

	Name  string
	Color string

	// Read-only properties

	Devices     []DeviceItem
	ID          ID
	UserID      *ID
	CompanyID   ID
	CreatedDate time.Time
	UpdatedDate time.Time
}

// DeviceLabelRequiredFields is a subset of DeviceLabel fields required to create a DeviceLabel.
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
