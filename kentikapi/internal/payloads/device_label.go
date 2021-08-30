package payloads

import (
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/utils"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

// GetAllDeviceLabelsResponse represents DeviceLabelsAPI GetAll JSON response.
type GetAllDeviceLabelsResponse []DeviceLabelPayload

func (r GetAllDeviceLabelsResponse) ToDeviceLabels() (result []models.DeviceLabel, err error) {
	err = utils.ConvertList(r, PayloadToDeviceLabel, &result)
	return result, err
}

// GetDeviceLabelResponse represents DeviceLabelsAPI Get JSON response.
type GetDeviceLabelResponse DeviceLabelPayload

func (r GetDeviceLabelResponse) ToDeviceLabel() (models.DeviceLabel, error) {
	return PayloadToDeviceLabel(DeviceLabelPayload(r))
}

// CreateDeviceLabelRequest represents DeviceLabelsAPI Create JSON request.
type CreateDeviceLabelRequest DeviceLabelPayload

// CreateDeviceLabelResponse represents DeviceLabelsAPI Create JSON response.
type CreateDeviceLabelResponse = GetDeviceLabelResponse

// UpdateDeviceLabelRequest represents DeviceLabelsAPI Update JSON request.
type UpdateDeviceLabelRequest DeviceLabelPayload

// UpdateDeviceLabelResponse represents DeviceLabelsAPI Update JSON response.
type UpdateDeviceLabelResponse = GetDeviceLabelResponse

// DeleteDeviceLabelResponse represents DeviceLabelsAPI Delete JSON response. Yes delete returns an object.
type DeleteDeviceLabelResponse struct {
	Success bool `json:"success"`
}

// DeviceLabelPayload represents JSON DeviceLabel payload as it is transmitted from KentikAPI.
type DeviceLabelPayload struct {
	// following fields can appear in request: post/put, response: get/post/put
	Name  string  `json:"name"` // name is always required
	Color *string `json:"color,omitempty" request:"post" response:"get,post,put"`

	// following fields can appear in request: none, response: get/post/put
	ID          *models.ID          `json:"id,omitempty" response:"get,post,put"`
	UserID      *models.ID          `json:"user_id,string,omitempty"` // user_id is not always returned
	CompanyID   *models.ID          `json:"company_id,string,omitempty" response:"get,post,put"`
	Devices     []deviceItemPayload `json:"devices,omitempty"`
	CreatedDate *time.Time          `json:"created_date,omitempty" response:"get,post,put"`
	UpdatedDate *time.Time          `json:"updated_date,omitempty" response:"get,post,put"`
}

type deviceItemPayload struct {
	// following fields can appear in request: none, response: get, put.
	// Not in post as newly created label is not assigned to any device
	ID            models.ID `json:"id,string"`
	DeviceName    string    `json:"device_name"`
	DeviceSubtype string    `json:"device_subtype"`
	DeviceType    *string   `json:"device_type"` // device_type is not always returned
}

// PayloadToDeviceLabel transforms GET/POST/PUT response payload into DeviceLabel.
func PayloadToDeviceLabel(p DeviceLabelPayload) (models.DeviceLabel, error) {
	var devices []models.DeviceItem
	err := utils.ConvertList(p.Devices, payloadToDeviceItem, &devices)
	if err != nil {
		return models.DeviceLabel{}, err
	}

	return models.DeviceLabel{
		Name:        p.Name,
		Color:       *p.Color,
		ID:          *p.ID,
		UserID:      p.UserID,
		CompanyID:   *p.CompanyID,
		Devices:     devices,
		CreatedDate: *p.CreatedDate,
		UpdatedDate: *p.UpdatedDate,
	}, nil
}

func payloadToDeviceItem(p deviceItemPayload) (models.DeviceItem, error) {
	return models.DeviceItem{
		ID:            p.ID,
		DeviceName:    p.DeviceName,
		DeviceType:    p.DeviceType,
		DeviceSubtype: p.DeviceSubtype,
	}, nil
}

// DeviceLabelToPayload prepares POST/PUT request payload: fill only the user-provided fields.
func DeviceLabelToPayload(l models.DeviceLabel) DeviceLabelPayload {
	return DeviceLabelPayload{
		Name:  l.Name,
		Color: &l.Color,
	}
}
