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

type DeviceItem struct {
	ID            ID
	DeviceName    string
	DeviceSubtype string
	DeviceType    *string
}
