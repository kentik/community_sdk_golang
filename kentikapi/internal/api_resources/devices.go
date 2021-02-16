package api_resources

import (
	"context"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type DevicesAPI struct {
	BaseAPI
}

func NewDevicesAPI(transport api_connection.Transport) *DevicesAPI {
	return &DevicesAPI{BaseAPI{Transport: transport}}
}

// GetAll lists devices
func (a *DevicesAPI) GetAll(ctx context.Context) ([]models.Device, error) {
	var response api_payloads.GetAllDevicesResponse
	if err := a.GetAndValidate(ctx, api_endpoints.DevicesPath, &response); err != nil {
		return []models.Device{}, err
	}

	return response.ToDevices()
}

// Get retrieves device with given ID
func (a *DevicesAPI) Get(ctx context.Context, id models.ID) (*models.Device, error) {
	var response api_payloads.GetDeviceResponse
	if err := a.GetAndValidate(ctx, api_endpoints.GetDevicePath(id), &response); err != nil {
		return nil, err
	}

	device, err := response.ToDevice()
	return &device, err
}
