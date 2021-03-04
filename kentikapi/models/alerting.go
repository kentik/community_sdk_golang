package models

import "time"

type ManualMitigation struct {
	IPCidr                string
	Comment               *string
	PlatformID            ID
	MethodID              ID
	MinutesBeforeAutoStop string
}

type Alarm struct {
	AlarmID         ID        `json:"alarm_id"`
	RowType         string    `json:"row_type"`
	AlarmState      string    `json:"alarm_state"`
	AlertID         ID        `json:"alert_id"`
	MitigationID    ID        `json:"mitigation_id"`
	TresholdID      ID        `json:"threshold_id"`
	AlertKey        string    `json:"alert_key"`
	AlertDimension  string    `json:"alert_dimension"`
	AlertMetric     []string  `json:"alert_metric"`
	AlertValue      float32   `json:"alert_value"`
	AlertValue2nd   float32   `json:"alert_value2nd"`
	AlertValue3rd   float32   `json:"alert_value3rd"`
	AlertMatchCount int       `json:"alert_match_count"`
	AlertBaseline   int       `json:"alert_baseline"`
	AlertSeverity   string    `json:"alert_severity"`
	BaselineUsed    int       `json:"baseline_used"`
	LearningMode    Bool      `json:"learning_mode"`
	DebugMode       Bool      `json:"debug_mode"`
	AlarmStart      time.Time `json:"alarm_start"`
	AlarmEnd        *string   `json:"alarm_end"`
	AlarmLastComent *string   `json:"alarm_last_comment"`
	MitAlertID      ID        `json:"mit_alert_id"`
	MitAlertIP      string    `json:"mit_alert_ip"`
	MitTresholdID   int       `json:"mit_treshold_id"`
	Args            string    `json:"args"`
	ID              ID        `json:"id"`
	PolicyID        ID        `json:"policy_id"`
	PolicyName      string    `json:"policy_name"`
	AlertKeyLookup  string    `json:"alert_key_lookup"`
}

type HistoricalAlert struct {
	RowType         string    `json:"row_type"`
	OldAlarmState   string    `json:"old_alarm_state"`
	NewAlarmState   string    `json:"new_alarm_state"`
	AlertMatchCount string    `json:"alert_match_count"`
	AlertSeverity   string    `json:"alert_severity"`
	AlertID         ID        `json:"alert_id"`
	ThresholdID     ID        `json:"threshold_id"`
	AlarmID         ID        `json:"alarm_id"`
	AlertKey        string    `json:"alert_key"`
	AlertDimension  string    `json:"alert_dimension"`
	AlertMetric     []string  `json:"alert_metric"`
	AlertValue      float32   `json:"alert_value"`
	AlertValue2nd   int       `json:"alert_value2nd"`
	AlertValue3rd   int       `json:"alert_value3rd"`
	AlertBaseline   int       `json:"alert_baseline"`
	BaselineUsed    int       `json:"baseline_used"`
	LearningMode    Bool      `json:"learning_mode"`
	DebugMode       Bool      `json:"debug_mode"`
	CreationTime    time.Time `json:"ctime"`
	AlarmStartTime  *string   `json:"alarm_start_time"`
	Comment         *string   `json:"comment"`
	MitigationID    *int      `json:"mitigation_id"`
	MitMethodID     int       `json:"mit_method_id"`
	Args            string    `json:"args"`
	ID              ID        `json:"id"`
	PolicyID        ID        `json:"policy_id"`
	PolicyName      string    `json:"policy_name"`
	AlertKeyLookup  string    `json:"alert_key_lookup"`
}
