package resources

import (
	"context"
	"fmt"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/utils"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type DeviceLabelsAPI struct {
	BaseAPI
}

// NewDeviceLabelsAPI is constructor.
func NewDeviceLabelsAPI(transport api_connection.Transport, logPayloads bool) *DeviceLabelsAPI {
	return &DeviceLabelsAPI{
		BaseAPI{Transport: transport, LogPayloads: logPayloads},
	}
}

// GetAll labels.
func (a *DeviceLabelsAPI) GetAll(ctx context.Context) ([]models.DeviceLabel, error) {
	utils.LogPayload(a.LogPayloads, "GetAll labels Kentik API request", "")
	var response api_payloads.GetAllDeviceLabelsResponse
	if err := a.GetAndValidate(ctx, api_endpoints.GetAllLabels(), &response); err != nil {
		return []models.DeviceLabel{}, err
	}
	utils.LogPayload(a.LogPayloads, "GetAll labels Kentik API response", response)

	return response.ToDeviceLabels()
}

// Get label with given ID.
func (a *DeviceLabelsAPI) Get(ctx context.Context, id models.ID) (*models.DeviceLabel, error) {
	utils.LogPayload(a.LogPayloads, "Get label Kentik API request ID", id)
	var response api_payloads.GetDeviceLabelResponse
	if err := a.GetAndValidate(ctx, api_endpoints.GetLabel(id), &response); err != nil {
		return nil, err
	}
	utils.LogPayload(a.LogPayloads, "Get label Kentik API response", response)

	device, err := response.ToDeviceLabel()
	return &device, err
}

// Create new label.
func (a *DeviceLabelsAPI) Create(ctx context.Context, label models.DeviceLabel) (*models.DeviceLabel, error) {
	payload := api_payloads.DeviceLabelToPayload(label)

	request := api_payloads.CreateDeviceLabelRequest(payload)
	utils.LogPayload(a.LogPayloads, "Create label Kentik API request", request)
	var response api_payloads.CreateDeviceLabelResponse
	if err := a.PostAndValidate(ctx, api_endpoints.CreateLabel(), request, &response); err != nil {
		return nil, err
	}
	utils.LogPayload(a.LogPayloads, "Create label Kentik API response", response)

	result, err := response.ToDeviceLabel()
	return &result, err
}

// Update label.
func (a *DeviceLabelsAPI) Update(ctx context.Context, label models.DeviceLabel) (*models.DeviceLabel, error) {
	payload := api_payloads.DeviceLabelToPayload(label)

	request := api_payloads.UpdateDeviceLabelRequest(payload)
	utils.LogPayload(a.LogPayloads, "Update label Kentik API request", request)
	var response api_payloads.UpdateDeviceLabelResponse
	if err := a.UpdateAndValidate(ctx, api_endpoints.UpdateLabel(label.ID), request, &response); err != nil {
		return nil, err
	}
	utils.LogPayload(a.LogPayloads, "Update label Kentik API response", response)

	result, err := response.ToDeviceLabel()
	return &result, err
}

// Delete label.
func (a *DeviceLabelsAPI) Delete(ctx context.Context, id models.ID) error {
	utils.LogPayload(a.LogPayloads, "Delete label Kentik API request ID", id)
	var response api_payloads.DeleteDeviceLabelResponse
	if err := a.DeleteAndValidate(ctx, api_endpoints.DeleteLabel(id), &response); err != nil {
		return err
	}

	if !response.Success {
		return fmt.Errorf("DeviceLabelsAPI.Delete: API returned success=false for id=%v", id)
	}

	return nil
}
