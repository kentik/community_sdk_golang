package api_payloads

import (
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/utils"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

// Plan only supports GET, so no optional fields except for UpdatedDate which sometimes comes as null
type PlanPayload struct {
	ID            models.ID                     `json:"id" response:"get"`
	CompanyID     models.ID                     `json:"company_id" response:"get"`
	Name          string                  `json:"name" response:"get"`
	Description   string                  `json:"description" response:"get"`
	Active        bool                    `json:"active" response:"get"`
	MaxDevices    int                     `json:"max_devices" response:"get"`
	MaxFPS        int                     `json:"max_fps" response:"get"`
	BGPEnabled    bool                    `json:"bgp_enabled" response:"get"`
	FastRetention int                     `json:"fast_retention" response:"get"`
	FullRetention int                     `json:"full_retention" response:"get"`
	CreatedDate   time.Time                  `json:"cdate" response:"get"`
	UpdatedDate   *time.Time                 `json:"edate"`
	MaxBigdataFPS int                     `json:"max_bigdata_fps" response:"get"`
	DeviceTypes   []planDeviceTypePayload `json:"deviceTypes" response:"get"`
	Devices       []planDevicePayload     `json:"devices" response:"get"`
}

func (p PlanPayload) ToPlan() (models.Plan, error) {
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
		ID            : p.ID,                     
		CompanyID     :p.CompanyID,                     
		Name          : p.Name,                  
		Description   : p.Description,                  
		Active        : p.Active,        
		MaxDevices    : p.MaxDevices,             
		MaxFPS        : p.MaxFPS,         
		BGPEnabled    : p.BGPEnabled,             
		FastRetention : p.FastRetention,         
		FullRetention : p.FullRetention,      
		CreatedDate   : p.CreatedDate,      
		UpdatedDate   : p.UpdatedDate,        
		MaxBigdataFPS : p.MaxBigdataFPS,        
		DeviceTypes   : deviceTypes,
		Devices       : devices  ,
	}, nil
}

type planDeviceTypePayload struct {
	DeviceType string `json:"device_type" response:"get"`
}

func payloadToPlanDeviceType(p planDeviceTypePayload) (models.PlanDeviceType, error) {
	return models.PlanDeviceType{DeviceType: p.DeviceType}, nil
}

type planDevicePayload struct {
	DeviceName string `json:"device_name" response:"get"`
	DeviceType string `json:"device_type" response:"get"`
	ID         models.ID `json:"id,string" response:"get"`
}

func payloadToPlanDevice(p planDevicePayload) (models.PlanDevice, error) {
	return models.PlanDevice{
		DeviceName: p.DeviceName,
		DeviceType: p.DeviceType,
		ID:p.ID,
	}, nil
}