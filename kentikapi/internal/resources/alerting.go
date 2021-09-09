package resources

import (
	"context"
	"errors"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type AlertingAPI struct {
	BaseAPI
}

// NewAlertingAPI is constructor.
func NewAlertingAPI(transport api_connection.Transport) *AlertingAPI {
	return &AlertingAPI{
		BaseAPI{Transport: transport},
	}
}

func (a *AlertingAPI) CreateManualMitigation(ctx context.Context, mm models.ManualMitigation) error {
	payload := api_payloads.ManualMitigationToPayload(mm)
	var response api_payloads.CreateManualMitigationResponse

	if err := a.PostAndValidate(ctx, api_endpoints.ManualMitigationPath, payload, &response); err != nil {
		return err
	}

	if response.Response.Result != "OK" {
		return errors.New("creating Manual Mitigation failed")
	}

	return nil
}

func (a *AlertingAPI) GetActiveAlerts(ctx context.Context, params models.AlertsQueryParams) ([]models.Alarm, error) {
	var response api_payloads.GetActiveAlertsResponse

	path := api_endpoints.GetActiveAlertsPath(params.StartTime, params.EndTime, params.FilterBy, params.FilterVal,
		params.ShowMitigations, params.ShowAlarms, params.ShowMatches, params.LearningMode)

	if err := a.GetAndValidate(ctx, path, &response); err != nil {
		return nil, err
	}

	return response.ToAlarms(), nil
}

func (a *AlertingAPI) GetAlertsHistory(ctx context.Context,
	params models.AlertsQueryParams) ([]models.HistoricalAlert, error) {
	var response api_payloads.GetHistoricalAlertsResponse
	path := api_endpoints.GetAlertsHistoryPath(params.StartTime, params.EndTime, params.FilterBy, params.FilterVal,
		params.SortOrder, params.ShowMitigations, params.ShowAlarms, params.ShowMatches, params.LearningMode)

	if err := a.GetAndValidate(ctx, path, &response); err != nil {
		return nil, err
	}

	return response.ToHistoricalAlerts(), nil
}
