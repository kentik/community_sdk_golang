package api_payloads

import (
	"strconv"
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type CreateManualMitigationRequest struct {
	IPCidr                string    `json:"ipCidr"`
	Comment               *string   `json:"comment,omitempty"`
	PlatformID            models.ID `json:"platformID"`
	MethodID              models.ID `json:"methodID"`
	MinutesBeforeAutoStop string    `json:"minutesBeforeAutoStop"`
}

type AlarmPayload struct {
	AlarmID          int               `json:"alarm_id"`
	RowType          string            `json:"row_type"`
	AlarmState       string            `json:"alarm_state"`
	AlertID          int               `json:"alert_id"`
	MitigationID     *StringAsInt      `json:"mitigation_id"`
	ThresholdID      int               `json:"threshold_id"`
	AlertKey         string            `json:"alert_key"`
	AlertDimension   string            `json:"alert_dimension"`
	AlertMetric      []string          `json:"alert_metric"`
	AlertValue       float32           `json:"alert_value"`
	AlertValue2nd    float32           `json:"alert_value2nd"`
	AlertValue3rd    float32           `json:"alert_value3rd"`
	AlertMatchCount  int               `json:"alert_match_count"`
	AlertBaseline    int               `json:"alert_baseline"`
	AlertSeverity    string            `json:"alert_severity"`
	BaselineUsed     int               `json:"baseline_used"`
	LearningMode     BoolAsStringOrInt `json:"learning_mode"`
	DebugMode        BoolAsStringOrInt `json:"debug_mode"`
	AlarmStart       time.Time         `json:"alarm_start"`
	AlarmEnd         *string           `json:"alarm_end,omitempty"`
	AlarmLastComment *string           `json:"alarm_last_comment,omitempty"`
	MitAlertID       int               `json:"mit_alert_id"`
	MitAlertIP       string            `json:"mit_alert_ip"`
	MitThresholdID   int               `json:"mit_treshold_id"`
	Args             string            `json:"args"`
	ID               int               `json:"id"`
	PolicyID         int               `json:"policy_id"`
	PolicyName       string            `json:"policy_name"`
	AlertKeyLookup   string            `json:"alert_key_lookup"`
}

func (p AlarmPayload) ToAlarm() models.Alarm {
	return models.Alarm{
		AlarmID:          strconv.Itoa(p.AlarmID),
		RowType:          p.RowType,
		AlarmState:       p.AlarmState,
		AlertID:          strconv.Itoa(p.AlertID),
		MitigationID:     (*models.ID)(p.MitigationID),
		ThresholdID:      strconv.Itoa(p.ThresholdID),
		AlertKey:         p.AlertKey,
		AlertDimension:   p.AlertDimension,
		AlertMetric:      p.AlertMetric,
		AlertValue:       p.AlertValue,
		AlertValue2nd:    p.AlertValue2nd,
		AlertValue3rd:    p.AlertValue3rd,
		AlertMatchCount:  p.AlertMatchCount,
		AlertBaseline:    p.AlertBaseline,
		AlertSeverity:    p.AlertSeverity,
		BaselineUsed:     p.BaselineUsed,
		LearningMode:     bool(p.LearningMode),
		DebugMode:        bool(p.DebugMode),
		AlarmStart:       p.AlarmStart,
		AlarmEnd:         p.AlarmEnd,
		AlarmLastComment: p.AlarmLastComment,
		MitAlertID:       strconv.Itoa(p.MitAlertID),
		MitAlertIP:       p.MitAlertIP,
		MitThresholdID:   strconv.Itoa(p.MitThresholdID),
		Args:             p.Args,
		ID:               strconv.Itoa(p.ID),
		PolicyID:         strconv.Itoa(p.PolicyID),
		PolicyName:       p.PolicyName,
		AlertKeyLookup:   p.AlertKeyLookup,
	}
}

type HistoricalAlert struct {
	RowType         string            `json:"row_type"`
	OldAlarmState   string            `json:"old_alarm_state"`
	NewAlarmState   string            `json:"new_alarm_state"`
	AlertMatchCount string            `json:"alert_match_count"`
	AlertSeverity   string            `json:"alert_severity"`
	AlertID         int               `json:"alert_id"`
	ThresholdID     int               `json:"threshold_id"`
	AlarmID         int               `json:"alarm_id"`
	AlertKey        string            `json:"alert_key"`
	AlertDimension  string            `json:"alert_dimension"`
	AlertMetric     []string          `json:"alert_metric"`
	AlertValue      float32           `json:"alert_value"`
	AlertValue2nd   int               `json:"alert_value2nd"`
	AlertValue3rd   int               `json:"alert_value3rd"`
	AlertBaseline   float32           `json:"alert_baseline"`
	BaselineUsed    int               `json:"baseline_used"`
	LearningMode    BoolAsStringOrInt `json:"learning_mode"`
	DebugMode       BoolAsStringOrInt `json:"debug_mode"`
	CreationTime    time.Time         `json:"ctime"`
	AlarmStartTime  *string           `json:"alarm_start_time,omitempty"`
	Comment         *string           `json:"comment,omitempty"`
	MitigationID    *StringAsInt      `json:"mitigation_id,omitempty"`
	MitMethodID     int               `json:"mit_method_id"`
	Args            string            `json:"args"`
	ID              int               `json:"id"`
	PolicyID        int               `json:"policy_id"`
	PolicyName      string            `json:"policy_name"`
	AlertKeyLookup  string            `json:"alert_key_lookup"`
}

func (p HistoricalAlert) ToHistoricalAlert() models.HistoricalAlert {
	return models.HistoricalAlert{
		RowType:         p.RowType,
		OldAlarmState:   p.OldAlarmState,
		NewAlarmState:   p.NewAlarmState,
		AlertMatchCount: p.AlertMatchCount,
		AlertSeverity:   p.AlertSeverity,
		AlertID:         strconv.Itoa(p.AlertID),
		ThresholdID:     strconv.Itoa(p.ThresholdID),
		AlarmID:         strconv.Itoa(p.AlarmID),
		AlertKey:        p.AlertKey,
		AlertDimension:  p.AlertDimension,
		AlertMetric:     p.AlertMetric,
		AlertValue:      p.AlertValue,
		AlertValue2nd:   p.AlertValue2nd,
		AlertValue3rd:   p.AlertValue3rd,
		AlertBaseline:   p.AlertBaseline,
		BaselineUsed:    p.BaselineUsed,
		LearningMode:    bool(p.LearningMode),
		DebugMode:       bool(p.DebugMode),
		CreationTime:    p.CreationTime,
		AlarmStartTime:  p.AlarmStartTime,
		Comment:         p.Comment,
		MitigationID:    (*models.ID)(p.MitigationID),
		MitMethodID:     strconv.Itoa(p.MitMethodID),
		Args:            p.Args,
		ID:              strconv.Itoa(p.ID),
		PolicyID:        strconv.Itoa(p.PolicyID),
		PolicyName:      p.PolicyName,
		AlertKeyLookup:  p.AlertKeyLookup,
	}
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

type GetActiveAlertsResponse []AlarmPayload

func (r GetActiveAlertsResponse) ToAlarms() []models.Alarm {
	var alarms []models.Alarm

	for _, a := range r {
		alarms = append(alarms, a.ToAlarm())
	}

	return alarms
}

type GetHistoricalAlertsResponse []HistoricalAlert

func (r GetHistoricalAlertsResponse) ToHistoricalAlerts() []models.HistoricalAlert {
	var alerts []models.HistoricalAlert

	for _, a := range r {
		alerts = append(alerts, a.ToHistoricalAlert())
	}

	return alerts
}
