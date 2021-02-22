package api_endpoints

import (
	"fmt"
)

const (
	DevicePath  = "/device"
	DevicesPath = "/devices"
)

func GetDevice(id ResourceID) string {
	return fmt.Sprintf("%v/%v", DevicePath, id)
}

func UpdateDevice(id ResourceID) string {
	return fmt.Sprintf("%v/%v", DevicePath, id)
}

func ApplyDeviceLabels(id ResourceID) string {
	return fmt.Sprintf("/devices/%v/labels", id)
}

func GetInterface(deviceID, interfaceID ResourceID) string {
	return fmt.Sprintf("/device/%v/interface/%v", deviceID, interfaceID)
}

func GetAllInterfaces(deviceID ResourceID) string {
	return fmt.Sprintf("/device/%v/interfaces", deviceID)
}
