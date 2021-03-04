package api_payloads

import (
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type CreateManualMitigationRequest struct {
	IPCidr                string    `json:"ipCidr"`
	Comment               *string   `json:"comment"`
	PlatformID            models.ID `json:"platformID"`
	MethodID              models.ID `json:"methodID"`
	MinutesBeforeAutoStop string    `json:"minutesBeforeAutoStop"`
}

func ManualMitigationToPayload(mm models.ManualMitigation) CreateManualMitigationRequest {
	return CreateManualMitigationRequest{
		IPCidr:                mm.IPCidr,
		Comment:               mm.Comment,
		PlatformID:            mm.PlatformID,
		MethodID:              mm.MethodID,
		MinutesBeforeAutoStop: mm.MinutesBeforeAutoStop,
	}
}

type CreateManualMitigationResponse struct {
	Response Response `json:"response"`
}

type Response struct {
	Result string `json:"result"`
}

type GetActiveAlertsResponse []models.Alarm

func (r GetActiveAlertsResponse) ToAlarms() ([]models.Alarm, error) {
	return r, nil
}

type GetHistoricalAlertsResponse []models.HistoricalAlert

func (r GetHistoricalAlertsResponse) ToHistoricalAlerts() ([]models.HistoricalAlert, error) {
	return r, nil
}
