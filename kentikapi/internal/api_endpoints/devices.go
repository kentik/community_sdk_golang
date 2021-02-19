package api_endpoints

import (
	"fmt"
)

const (
	DevicePath  = "/device"
	DevicesPath = "/devices"
)

func GetDevicePath(id ResourceID) string {
	return fmt.Sprintf("%v/%v", DevicePath, id)
}

func UpdateDevicePath(id ResourceID) string {
	return fmt.Sprintf("%v/%v", DevicePath, id)
}
