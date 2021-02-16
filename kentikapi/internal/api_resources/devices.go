package api_resources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/validation"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type DevicesAPI struct {
	transport api_connection.Transport
}

func NewDevicesAPI(transport api_connection.Transport) *DevicesAPI {
	return &DevicesAPI{transport: transport}
}

func (a *DevicesAPI) GetAll(ctx context.Context) ([]models.Device, error) {
	responseBody, err := a.transport.Get(ctx, api_endpoints.DevicesPath)
	if err != nil {
		return nil, err
	}
	var response api_payloads.DeviceGetAllResponse
	if err = json.Unmarshal(responseBody, &response); err != nil {
		return nil, fmt.Errorf("unmarshal response body: %v", err)
	}

	if err = validation.ValidateResponse("get", response); err != nil {
		return []models.Device{}, err
	}

	return response.ToDevices()
}

func (a *DevicesAPI) Get(ctx context.Context, id models.ID) (*models.Device, error) {
	responseBody, err := a.transport.Get(ctx, api_endpoints.GetDevicePath(id))
	if err != nil {
		return nil, err
	}

	var response api_payloads.DeviceGetResponse
	if err = json.Unmarshal(responseBody, &response); err != nil {
		return nil, fmt.Errorf("unmarshal response body: %v", err)
	}

	if err = validation.ValidateResponse("get", response); err != nil {
		return nil, err
	}

	device, err := response.ToDevice()
	return &device, err
}
