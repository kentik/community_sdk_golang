package resources

import (
	"context"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/utils"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type DevicesAPI struct {
	BaseAPI
	Interfaces *interfacesAPI
}

// NewDevicesAPI is constructor.
func NewDevicesAPI(transport api_connection.Transport, logPayloads bool) *DevicesAPI {
	return &DevicesAPI{
		BaseAPI{Transport: transport, LogPayloads: logPayloads},
		&interfacesAPI{BaseAPI{Transport: transport}},
	}
}

// GetAll devices.
func (a *DevicesAPI) GetAll(ctx context.Context) ([]models.Device, error) {
	utils.LogPayload(a.LogPayloads, "GetAll devices Kentik API request", "")
	var response api_payloads.GetAllDevicesResponse
	if err := a.GetAndValidate(ctx, api_endpoints.DevicesPath, &response); err != nil {
		return []models.Device{}, err
	}
	utils.LogPayload(a.LogPayloads, "GetAll devices Kentik API response", response)

	return response.ToDevices()
}

// Get device with given ID.
func (a *DevicesAPI) Get(ctx context.Context, id models.ID) (*models.Device, error) {
	utils.LogPayload(a.LogPayloads, "Get device Kentik API request ID", id)
	var response api_payloads.GetDeviceResponse
	if err := a.GetAndValidate(ctx, api_endpoints.GetDevice(id), &response); err != nil {
		return nil, err
	}
	utils.LogPayload(a.LogPayloads, "Get device Kentik API response", response)

	device, err := response.ToDevice()
	return &device, err
}

// Create new device.
func (a *DevicesAPI) Create(ctx context.Context, device models.Device) (*models.Device, error) {
	request := api_payloads.CreateDeviceRequest{Payload: api_payloads.DeviceToPayload(device)}
	utils.LogPayload(a.LogPayloads, "Create device Kentik API request", request)
	var response api_payloads.CreateDeviceResponse
	if err := a.PostAndValidate(ctx, api_endpoints.DevicePath, request, &response); err != nil {
		return nil, err
	}
	utils.LogPayload(a.LogPayloads, "Create device Kentik API response", response)

	result, err := response.ToDevice()
	return &result, err
}

// Update device.
func (a *DevicesAPI) Update(ctx context.Context, device models.Device) (*models.Device, error) {
	request := api_payloads.UpdateDeviceRequest{Payload: api_payloads.DeviceToPayload(device)}
	utils.LogPayload(a.LogPayloads, "Update device Kentik API request", request)
	var response api_payloads.UpdateDeviceResponse
	if err := a.UpdateAndValidate(ctx, api_endpoints.UpdateDevice(device.ID), request, &response); err != nil {
		return nil, err
	}
	utils.LogPayload(a.LogPayloads, "Update device Kentik API response", response)

	result, err := response.ToDevice()
	return &result, err
}

// Delete device
// Note: KentikAPI requires sending delete request twice to actually delete the device.
// This is a safety measure preventing deletion by mistake.
func (a *DevicesAPI) Delete(ctx context.Context, id models.ID) error {
	utils.LogPayload(a.LogPayloads, "Delete device Kentik API request ID", id)
	return a.DeleteAndValidate(ctx, api_endpoints.GetDevice(id), nil)
}

// ApplyLabels assigns labels to given device.
func (a *DevicesAPI) ApplyLabels(ctx context.Context, deviceID models.ID, labels []models.ID) (models.AppliedLabels, error) {
	payload, err := api_payloads.LabelIDsToPayload(labels)
	if err != nil {
		return models.AppliedLabels{}, err
	}
	request := api_payloads.ApplyLabelsRequest{Labels: payload}
	utils.LogPayload(a.LogPayloads, "ApplyLabels to given device Kentik API request", request)
	var response api_payloads.ApplyLabelsResponse
	if err := a.UpdateAndValidate(ctx, api_endpoints.ApplyDeviceLabels(deviceID), request, &response); err != nil {
		return models.AppliedLabels{}, err
	}
	utils.LogPayload(a.LogPayloads, "ApplyLabels to given device Kentik API response", response)

	return response.ToAppliedLabels()
}

type interfacesAPI struct {
	BaseAPI
}

// GetAll interfaces of given device.
func (a *interfacesAPI) GetAll(ctx context.Context, deviceID models.ID) ([]models.Interface, error) {
	utils.LogPayload(a.LogPayloads, "GetAll interfaces of given device Kentik API request device ID", deviceID)
	var response api_payloads.GetAllInterfacesResponse
	if err := a.GetAndValidate(ctx, api_endpoints.GetAllInterfaces(deviceID), &response); err != nil {
		return nil, err
	}
	utils.LogPayload(a.LogPayloads, "GetAll interfaces of given device Kentik API response", response)

	return response.ToInterfaces()
}

// Get interface of given device with given ID.
func (a *interfacesAPI) Get(ctx context.Context, deviceID, interfaceID models.ID) (*models.Interface, error) {
	utils.LogPayload(a.LogPayloads, "Get interface of given device Kentik API request device ID", deviceID)
	var response api_payloads.GetInterfaceResponse
	if err := a.GetAndValidate(ctx, api_endpoints.GetInterface(deviceID, interfaceID), &response); err != nil {
		return nil, err
	}
	utils.LogPayload(a.LogPayloads, "Get interface of given device Kentik API response", response)

	intf, err := response.ToInterface()
	return &intf, err
}

// Create new interface under given device.
func (a *interfacesAPI) Create(ctx context.Context, intf models.Interface) (*models.Interface, error) {
	payload, err := api_payloads.InterfaceToPayload(intf)
	if err != nil {
		return nil, err
	}
	request := api_payloads.CreateInterfaceRequest(payload)
	utils.LogPayload(a.LogPayloads, "Create new interface under given device Kentik API request", request)
	var response api_payloads.CreateInterfaceResponse
	if err = a.PostAndValidate(ctx, api_endpoints.CreateInterface(intf.DeviceID), request, &response); err != nil {
		return nil, err
	}
	utils.LogPayload(a.LogPayloads, "Create new interface under given device Kentik API response", response)

	result, err := response.ToInterface()
	return &result, err
}

// Delete interface.
func (a *interfacesAPI) Delete(ctx context.Context, deviceID, interfaceID models.ID) error {
	utils.LogPayload(a.LogPayloads, "Delete interface Kentik API request ID", interfaceID)
	return a.DeleteAndValidate(ctx, api_endpoints.DeleteInterface(deviceID, interfaceID), nil)
}

// Update interface.
func (a *interfacesAPI) Update(ctx context.Context, intf models.Interface) (*models.Interface, error) {
	payload, err := api_payloads.InterfaceToPayload(intf)
	if err != nil {
		return nil, err
	}
	request := api_payloads.UpdateInterfaceRequest(payload)
	utils.LogPayload(a.LogPayloads, "Update interface Kentik API request", request)
	var response api_payloads.UpdateInterfaceResponse
	if err = a.UpdateAndValidate(ctx, api_endpoints.UpdateInterface(intf.DeviceID, intf.ID), request, &response); err != nil {
		return nil, err
	}
	utils.LogPayload(a.LogPayloads, "Update interface Kentik API response", response)

	result, err := response.ToInterface()
	return &result, err
}
