package api_resources

import (
	"context"
	"errors"
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type AlertingAPI struct {
	BaseAPI
}

// NewAlertingAPI is constructor
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
		return errors.New("Creating Manual Mitigation Failed")
	}

	return nil
}

func (a *AlertingAPI) GetActiveAlerts(ctx context.Context, startTime time.Time, endTime time.Time, filterBy string, filterVal string, showMitigations bool,
	showAlarms bool, showMatches bool, learningMode bool) ([]models.Alarm, error) {
	var response api_payloads.GetActiveAlertsResponse
	timeFormatStr := "2006-01-02T15:04:05"
	path := api_endpoints.GetActiveAlertsPath(startTime.Format(timeFormatStr), endTime.Format(timeFormatStr), filterBy, filterVal, boolToInt(showMitigations),
		boolToInt(showAlarms), boolToInt(showMatches), boolToInt(learningMode))

	if err := a.GetAndValidate(ctx, path, &response); err != nil {
		return nil, err
	}

	return response.ToAlarms()
}

func (a *AlertingAPI) GetAlertsHistory(ctx context.Context, startTime time.Time, endTime time.Time, filterBy string, filterVal string, sortOrder string,
	showMitigations bool, showAlarms bool, showMatches bool, learningMode bool) ([]models.HistoricalAlert, error) {
	var response api_payloads.GetHistoricalAlertsResponse
	timeFormatStr := "2006-01-02T15:04:05"
	path := api_endpoints.GetAlertsHistoryPath(startTime.Format(timeFormatStr), endTime.Format(timeFormatStr), filterBy, filterVal, sortOrder,
		boolToInt(showMitigations), boolToInt(showAlarms), boolToInt(showMatches), boolToInt(learningMode))

	if err := a.GetAndValidate(ctx, path, &response); err != nil {
		return nil, err
	}

	return response.ToHistoricalAlerts()
}

func boolToInt(val bool) int {
	var intVal int
	if val {
		intVal = 1
	}
	return intVal
}
