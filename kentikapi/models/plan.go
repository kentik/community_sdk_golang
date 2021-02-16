package models

import "time"

type Plan struct {
	ID            ID
	CompanyID     ID
	Name          string
	Description   string
	Active        bool
	MaxDevices    int
	MaxFPS        int
	BGPEnabled    bool
	FastRetention int
	FullRetention int
	CreatedDate   time.Time
	UpdatedDate   *time.Time
	MaxBigdataFPS int
	DeviceTypes   []PlanDeviceType
	Devices       []PlanDevice
}

type PlanDeviceType struct {
	DeviceType string
}

type PlanDevice struct {
	DeviceName string
	DeviceType string
	ID         ID
}
