package resources

import (
	"context"
	"fmt"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type DeviceLabelsAPI struct {
	BaseAPI
}

// NewDeviceLabelsAPI is constructor.
func NewDeviceLabelsAPI(transport connection.Transport) *DeviceLabelsAPI {
	return &DeviceLabelsAPI{
		BaseAPI{Transport: transport},
	}
}

// GetAll labels.
func (a *DeviceLabelsAPI) GetAll(ctx context.Context) ([]models.DeviceLabel, error) {
	var response payloads.GetAllDeviceLabelsResponse
	if err := a.GetAndValidate(ctx, endpoints.GetAllLabels(), &response); err != nil {
		return []models.DeviceLabel{}, err
	}

	return response.ToDeviceLabels()
}

// Get label with given ID.
func (a *DeviceLabelsAPI) Get(ctx context.Context, id models.ID) (*models.DeviceLabel, error) {
	var response payloads.GetDeviceLabelResponse
	if err := a.GetAndValidate(ctx, endpoints.GetLabel(id), &response); err != nil {
		return nil, err
	}

	device, err := response.ToDeviceLabel()
	return &device, err
}

// Create new label.
func (a *DeviceLabelsAPI) Create(ctx context.Context, label models.DeviceLabel) (*models.DeviceLabel, error) {
	payload := payloads.DeviceLabelToPayload(label)

	request := payloads.CreateDeviceLabelRequest(payload)
	var response payloads.CreateDeviceLabelResponse
	if err := a.PostAndValidate(ctx, endpoints.CreateLabel(), request, &response); err != nil {
		return nil, err
	}

	result, err := response.ToDeviceLabel()
	return &result, err
}

// Update label.
func (a *DeviceLabelsAPI) Update(ctx context.Context, label models.DeviceLabel) (*models.DeviceLabel, error) {
	payload := payloads.DeviceLabelToPayload(label)

	request := payloads.UpdateDeviceLabelRequest(payload)
	var response payloads.UpdateDeviceLabelResponse
	if err := a.UpdateAndValidate(ctx, endpoints.UpdateLabel(label.ID), request, &response); err != nil {
		return nil, err
	}

	result, err := response.ToDeviceLabel()
	return &result, err
}

// Delete label.
func (a *DeviceLabelsAPI) Delete(ctx context.Context, id models.ID) error {
	var response payloads.DeleteDeviceLabelResponse
	if err := a.DeleteAndValidate(ctx, endpoints.DeleteLabel(id), &response); err != nil {
		return err
	}

	if !response.Success {
		return fmt.Errorf("DeviceLabelsAPI.Delete: API returned success=false for id=%v", id)
	}

	return nil
}
