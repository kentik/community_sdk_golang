package api_endpoints

import (
	"fmt"

	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

const (
	userPath  = "/user"
	usersPath = "/users"
	devicePath = "/device"
	devicesPath = "/devices"
)

func getUserPath(id models.ID) string {
	return fmt.Sprintf("%v/%v", userPath, id)
}

func getDevicePath(id models.ID) string {
	return fmt.Sprintf("%v/%v", devicePath, id)
}
