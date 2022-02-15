// Package models holds definitions of Kentik API resources
package models

import (
	"time"
)

type ManualMitigation struct {
	IPCidr                string
	Comment               *string
	PlatformID            ID
	MethodID              ID
	MinutesBeforeAutoStop string
}

type Alarm struct {
	AlarmID         ID
	RowType         string
	AlarmState      string
	AlertID         ID
	MitigationID    *ID
	TresholdID      ID
	AlertKey        string
	AlertDimension  string
	AlertMetric     []string
	AlertValue      float32
	AlertValue2nd   float32
	AlertValue3rd   float32
	AlertMatchCount int
	AlertBaseline   int
	AlertSeverity   string
	BaselineUsed    int
	LearningMode    bool
	DebugMode       bool
	AlarmStart      time.Time
	AlarmEnd        *string
	AlarmLastComent *string
	MitAlertID      ID
	MitAlertIP      string
	MitTresholdID   ID
	Args            string
	ID              ID
	PolicyID        ID
	PolicyName      string
	AlertKeyLookup  string
}

type HistoricalAlert struct {
	RowType         string
	OldAlarmState   string
	NewAlarmState   string
	AlertMatchCount string
	AlertSeverity   string
	AlertID         ID
	ThresholdID     ID
	AlarmID         ID
	AlertKey        string
	AlertDimension  string
	AlertMetric     []string
	AlertValue      float32
	AlertValue2nd   int
	AlertValue3rd   int
	AlertBaseline   float32
	BaselineUsed    int
	LearningMode    bool
	DebugMode       bool
	CreationTime    time.Time
	AlarmStartTime  *string
	Comment         *string
	MitigationID    *ID
	MitMethodID     ID
	Args            string
	ID              ID
	PolicyID        ID
	PolicyName      string
	AlertKeyLookup  string
}

type AlertsQueryParams struct {
	StartTime       *time.Time
	EndTime         *time.Time
	FilterBy        string
	FilterVal       string
	SortOrder       string
	ShowMitigations bool
	ShowAlarms      bool
	ShowMatches     bool
	LearningMode    bool
}
