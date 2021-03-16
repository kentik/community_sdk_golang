package api_resources

import (
	"context"
	"fmt"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type DeviceLabelsAPI struct {
	BaseAPI
}

// NewDeviceLabelsAPIis constructor
func NewDeviceLabelsAPI(transport api_connection.Transport) *DeviceLabelsAPI {
	return &DeviceLabelsAPI{
		BaseAPI{Transport: transport},
	}
}

// GetAll labels
func (a *DeviceLabelsAPI) GetAll(ctx context.Context) ([]models.DeviceLabel, error) {
	var response api_payloads.GetAllDeviceLabelsResponse
	if err := a.GetAndValidate(ctx, api_endpoints.GetAllLabels(), &response); err != nil {
		return []models.DeviceLabel{}, err
	}

	return response.ToDeviceLabels()
}

// Get label with given ID
func (a *DeviceLabelsAPI) Get(ctx context.Context, id models.ID) (*models.DeviceLabel, error) {
	var response api_payloads.GetDeviceLabelResponse
	if err := a.GetAndValidate(ctx, api_endpoints.GetLabel(id), &response); err != nil {
		return nil, err
	}

	device, err := response.ToDeviceLabel()
	return &device, err
}

// Create new label
func (a *DeviceLabelsAPI) Create(ctx context.Context, label models.DeviceLabel) (*models.DeviceLabel, error) {
	payload := api_payloads.DeviceLabelToPayload(label)

	request := api_payloads.CreateDeviceLabelRequest(payload)
	var response api_payloads.CreateDeviceLabelResponse
	if err := a.PostAndValidate(ctx, api_endpoints.CreateLabel(), request, &response); err != nil {
		return nil, err
	}

	result, err := response.ToDeviceLabel()
	return &result, err
}

// Update label
func (a *DeviceLabelsAPI) Update(ctx context.Context, label models.DeviceLabel) (*models.DeviceLabel, error) {
	payload := api_payloads.DeviceLabelToPayload(label)

	request := api_payloads.UpdateDeviceLabelRequest(payload)
	var response api_payloads.UpdateDeviceLabelResponse
	if err := a.UpdateAndValidate(ctx, api_endpoints.UpdateLabel(label.ID), request, &response); err != nil {
		return nil, err
	}

	result, err := response.ToDeviceLabel()
	return &result, err
}

// Delete label
func (a *DeviceLabelsAPI) Delete(ctx context.Context, id models.ID) error {
	var response api_payloads.DeleteDeviceLabelResponse
	if err := a.DeleteAndValidate(ctx, api_endpoints.DeleteLabel(id), &response); err != nil {
		return err
	}

	if !response.Success {
		return fmt.Errorf("DeviceLabelsAPI.Delete: API returned success=false for id=%v", id)
	}

	return nil
}
