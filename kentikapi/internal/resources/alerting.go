package resources

import (
	"context"
	"errors"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type AlertingAPI struct {
	BaseAPI
}

// NewAlertingAPI is constructor.
func NewAlertingAPI(transport connection.Transport) *AlertingAPI {
	return &AlertingAPI{
		BaseAPI{Transport: transport},
	}
}

func (a *AlertingAPI) CreateManualMitigation(ctx context.Context, mm models.ManualMitigation) error {
	payload := payloads.ManualMitigationToPayload(mm)
	var response payloads.CreateManualMitigationResponse

	if err := a.PostAndValidate(ctx, endpoints.ManualMitigationPath, payload, &response); err != nil {
		return err
	}

	if response.Response.Result != "OK" {
		return errors.New("creating Manual Mitigation failed")
	}

	return nil
}

func (a *AlertingAPI) GetActiveAlerts(ctx context.Context, params models.AlertsQueryParams) ([]models.Alarm, error) {
	var response payloads.GetActiveAlertsResponse

	path := endpoints.GetActiveAlertsPath(params.StartTime, params.EndTime, params.FilterBy, params.FilterVal,
		params.ShowMitigations, params.ShowAlarms, params.ShowMatches, params.LearningMode)

	if err := a.GetAndValidate(ctx, path, &response); err != nil {
		return nil, err
	}

	return response.ToAlarms(), nil
}

func (a *AlertingAPI) GetAlertsHistory(ctx context.Context,
	params models.AlertsQueryParams) ([]models.HistoricalAlert, error) {
	var response payloads.GetHistoricalAlertsResponse
	path := endpoints.GetAlertsHistoryPath(params.StartTime, params.EndTime, params.FilterBy, params.FilterVal,
		params.SortOrder, params.ShowMitigations, params.ShowAlarms, params.ShowMatches, params.LearningMode)

	if err := a.GetAndValidate(ctx, path, &response); err != nil {
		return nil, err
	}

	return response.ToHistoricalAlerts(), nil
}
