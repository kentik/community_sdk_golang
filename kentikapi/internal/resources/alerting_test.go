package resources_test

import (
	"context"
	"testing"
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/resources"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
)

//nolint:gochecknoglobals
var (
	time1 = time.Date(2020, time.January, 19, 13, 50, 0, 0, time.Local)
	time2 = time.Date(2021, time.March, 19, 13, 50, 0, 0, time.Local)
)

func TestCrerateManualMitigation(t *testing.T) {
	createResponsePayload := `
	{
		"response": {
			"result": "OK"
		}
    }`
	expectedRequestBody := `{"ipCidr":"192.168.0.0/24","platformID":"1234","methodID":"12345","minutesBeforeAutoStop":"20"}`

	transport := &api_connection.StubTransport{ResponseBody: createResponsePayload}
	alertingAPI := resources.NewAlertingAPI(transport, true)
	mm := models.ManualMitigation{
		IPCidr:                "192.168.0.0/24",
		PlatformID:            "1234",
		MethodID:              "12345",
		Comment:               nil,
		MinutesBeforeAutoStop: "20",
	}

	assert.NoError(t, alertingAPI.CreateManualMitigation(context.Background(), mm))
	assert.Equal(t, "/alerts/manual-mitigate", transport.RequestPath)
	assert.Equal(t, expectedRequestBody, transport.RequestBody)
}

func TestGetActiveAlerts(t *testing.T) {
	getResponsePayload := `
	[
        {
            "alarm_id": 82867908,
            "row_type": "Alarm",
            "alarm_state": "ALARM",
            "alert_id": 15094,
            "mitigation_id": null,
            "threshold_id": 76518,
            "alert_key": "443",
            "alert_dimension": "Port_dst",
            "alert_metric": [
                "bits"
            ],
            "alert_value": 2270.4,
            "alert_value2nd": 0,
            "alert_value3rd": 0,
            "alert_match_count": 5,
            "alert_baseline": 769,
            "alert_severity": "minor",
            "baseline_used": 15,
            "learning_mode": 0,
            "debug_mode": 0,
            "alarm_start": "2021-01-19T13:50:00.000Z",
            "alarm_end": "0000-00-00 00:00:00",
            "alarm_last_comment": null,
            "mit_alert_id": 0,
            "mit_alert_ip": "",
            "mit_threshold_id": 0,
            "mit_method_id": 0,
            "args": "",
            "id": 0,
            "policy_id": 15094,
            "policy_name": "test_policy1",
            "alert_key_lookup": "443"
        }
    ]`
	transport := &api_connection.StubTransport{ResponseBody: getResponsePayload}
	alertingAPI := resources.NewAlertingAPI(transport, true)

	alarmEndStr := "0000-00-00 00:00:00"
	expected := []models.Alarm{
		{
			AlarmID:         "82867908",
			RowType:         "Alarm",
			AlarmState:      "ALARM",
			AlertID:         "15094",
			MitigationID:    nil,
			TresholdID:      "76518",
			AlertKey:        "443",
			AlertDimension:  "Port_dst",
			AlertMetric:     []string{"bits"},
			AlertValue:      2270.4,
			AlertValue2nd:   0,
			AlertValue3rd:   0,
			AlertMatchCount: 5,
			AlertBaseline:   769,
			AlertSeverity:   "minor",
			BaselineUsed:    15,
			LearningMode:    false,
			DebugMode:       false,
			AlarmStart:      time.Date(2021, time.January, 19, 13, 50, 0, 0, time.UTC),
			AlarmEnd:        &alarmEndStr,
			MitAlertID:      "0",
			MitAlertIP:      "",
			MitTresholdID:   "0",
			Args:            "",
			ID:              "0",
			PolicyID:        "15094",
			PolicyName:      "test_policy1",
			AlertKeyLookup:  "443",
		},
	}

	params := models.AlertsQueryParams{
		StartTime:       &time1,
		EndTime:         &time2,
		FilterBy:        "",
		FilterVal:       "",
		ShowMitigations: true,
		ShowAlarms:      true,
		ShowMatches:     false,
		LearningMode:    false,
	}

	alerts, err := alertingAPI.GetActiveAlerts(context.Background(), params)

	assert.Equal(t, expected, alerts)
	assert.NoError(t, err)
	assert.Equal(
		t,
		"/alerts-active/alarms?endTime=2021-03-19T13%3A50%3A00&learningMode=0&showAlarms=1&"+
			"showMatches=0&showMitigations=1&startTime=2020-01-19T13%3A50%3A00",
		transport.RequestPath,
	)
}

func TestGetAlertsHistory(t *testing.T) {
	getResponsePayload := `
	[
        {
            "row_type": "Alarm",
            "old_alarm_state": "CLEAR",
            "new_alarm_state": "ALARM",
            "alert_match_count": "1",
            "alert_severity": "minor",
            "alert_id": 15094,
            "threshold_id": 76518,
            "alarm_id": 82867908,
            "alert_key": "443",
            "alert_dimension": "Port_dst",
            "alert_metric": [
                "bits"
            ],
            "alert_value": 2270.4,
            "alert_value2nd": 0,
            "alert_value3rd": 0,
            "alert_baseline": 769,
            "baseline_used": 15,
            "learning_mode": 0,
            "debug_mode": 0,
            "ctime": "2021-01-19T13:50:00.000Z",
            "alarm_start_time": "2021-01-19 13:50:00",
            "comment": null,
            "mitigation_id": null,
            "mit_method_id": 0,
            "args": "",
            "id": 0,
            "policy_id": 15094,
            "policy_name": "test_policy1",
            "alert_key_lookup": "443"
        }
    ]`
	transport := &api_connection.StubTransport{ResponseBody: getResponsePayload}
	alertingAPI := resources.NewAlertingAPI(transport, true)

	dateStr := "2021-01-19 13:50:00"
	expected := []models.HistoricalAlert{
		{
			RowType:         "Alarm",
			OldAlarmState:   "CLEAR",
			NewAlarmState:   "ALARM",
			AlertMatchCount: "1",
			AlertSeverity:   "minor",
			AlertID:         "15094",
			ThresholdID:     "76518",
			AlarmID:         "82867908",
			AlertKey:        "443",
			AlertDimension:  "Port_dst",
			AlertMetric:     []string{"bits"},
			AlertValue:      2270.4,
			AlertValue2nd:   0,
			AlertValue3rd:   0,
			AlertBaseline:   769,
			BaselineUsed:    15,
			LearningMode:    false,
			DebugMode:       false,
			CreationTime:    time.Date(2021, time.January, 19, 13, 50, 0, 0, time.UTC),
			AlarmStartTime:  &dateStr,
			MitMethodID:     "0",
			ID:              "0",
			PolicyID:        "15094",
			PolicyName:      "test_policy1",
			AlertKeyLookup:  "443",
		},
	}

	params := models.AlertsQueryParams{
		StartTime:       &time1,
		EndTime:         &time2,
		FilterBy:        "",
		FilterVal:       "",
		SortOrder:       "",
		ShowMitigations: true,
		ShowAlarms:      true,
		ShowMatches:     false,
		LearningMode:    false,
	}

	alerts, err := alertingAPI.GetAlertsHistory(context.Background(), params)

	assert.Equal(t, expected, alerts)
	assert.NoError(t, err)
	assert.Equal(
		t, "/alerts-active/alerts-history?endTime=2021-03-19T13%3A50%3A00&learningMode=0&"+
			"showAlarms=1&showMatches=0&showMitigations=1&startTime=2020-01-19T13%3A50%3A00",
		transport.RequestPath,
	)
}
