package api_payloads

import (
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/utils"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

// GetAllPlansResponse represents PlansAPI GetAll JSON response.
type GetAllPlansResponse struct {
	Plans []PlanPayload `json:"plans"`
}

func (r GetAllPlansResponse) ToPlans() (result []models.Plan, err error) {
	err = utils.ConvertList(r.Plans, payloadToPlan, &result)
	return result, err
}

// PlanPayload represents JSON Plan payload as it is transmitted from KentikAPI.
type PlanPayload struct {
	// following fields can appear in request: none, response: get
	ID            models.ID               `json:"id"`
	CompanyID     models.ID               `json:"company_id"`
	Name          string                  `json:"name"`
	Description   string                  `json:"description"`
	Active        bool                    `json:"active"`
	MaxDevices    int                     `json:"max_devices"`
	MaxFPS        int                     `json:"max_fps"`
	BGPEnabled    bool                    `json:"bgp_enabled"`
	FastRetention int                     `json:"fast_retention"`
	FullRetention int                     `json:"full_retention"`
	CreatedDate   time.Time               `json:"cdate"`
	UpdatedDate   *time.Time              `json:"edate"` // the only optional field
	MaxBigdataFPS int                     `json:"max_bigdata_fps"`
	DeviceTypes   []planDeviceTypePayload `json:"deviceTypes"`
	Devices       []planDevicePayload     `json:"devices"`
}

//nolint:dupl
func payloadToPlan(p PlanPayload) (models.Plan, error) {
	var deviceTypes []models.PlanDeviceType
	err := utils.ConvertList(p.DeviceTypes, payloadToPlanDeviceType, &deviceTypes)
	if err != nil {
		return models.Plan{}, err
	}

	var devices []models.PlanDevice
	err = utils.ConvertList(p.Devices, payloadToPlanDevice, &devices)
	if err != nil {
		return models.Plan{}, err
	}

	return models.Plan{
		ID:            p.ID,
		CompanyID:     p.CompanyID,
		Name:          p.Name,
		Description:   p.Description,
		Active:        p.Active,
		MaxDevices:    p.MaxDevices,
		MaxFPS:        p.MaxFPS,
		BGPEnabled:    p.BGPEnabled,
		FastRetention: p.FastRetention,
		FullRetention: p.FullRetention,
		CreatedDate:   p.CreatedDate,
		UpdatedDate:   p.UpdatedDate,
		MaxBigdataFPS: p.MaxBigdataFPS,
		DeviceTypes:   deviceTypes,
		Devices:       devices,
	}, nil
}

// planDeviceTypePayload represents JSON Plan.DeviceTypes payload as it is transmitted from KentikAPI.
type planDeviceTypePayload struct {
	DeviceType string `json:"device_type"`
}

func payloadToPlanDeviceType(p planDeviceTypePayload) (models.PlanDeviceType, error) {
	return models.PlanDeviceType{DeviceType: p.DeviceType}, nil
}

// planDevicePayload represents JSON Plan.Devices payload as it is transmitted from KentikAPI.
type planDevicePayload struct {
	DeviceName string    `json:"device_name"`
	DeviceType string    `json:"device_type"`
	ID         models.ID `json:"id,string"`
}

func payloadToPlanDevice(p planDevicePayload) (models.PlanDevice, error) {
	return models.PlanDevice{
		DeviceName: p.DeviceName,
		DeviceType: p.DeviceType,
		ID:         p.ID,
	}, nil
}
