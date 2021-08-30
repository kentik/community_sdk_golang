package endpoints

import (
	"fmt"
	"net/url"
	"time"
)

const (
	ManualMitigationPath = "/alerts/manual-mitigate"
	AlertsActivePath     = "/alerts-active/alarms"
	AlertsHistoryPath    = "/alerts-active/alerts-history"
)

func GetActiveAlertsPath(startTime *time.Time, endTime *time.Time, filterBy string, filterVal string,
	showMitigations bool, showAlarms bool, showMatches bool, learningMode bool,
) string {
	timeFormatStr := "2006-01-02T15:04:05"
	v := url.Values{}
	if startTime != nil {
		v.Set("startTime", startTime.Format(timeFormatStr))
	}
	if endTime != nil {
		v.Set("endTime", endTime.Format(timeFormatStr))
	}
	if filterBy != "" {
		v.Set("filterBy", filterBy)
	}
	if filterVal != "" {
		v.Set("filterVal", filterVal)
	}
	v.Set("showMitigations", fmt.Sprint(boolToInt(showMitigations)))
	v.Set("showAlarms", fmt.Sprint(boolToInt(showAlarms)))
	v.Set("showMatches", fmt.Sprint(boolToInt(showMatches)))
	v.Set("learningMode", fmt.Sprint(boolToInt(learningMode)))

	return fmt.Sprintf("%v?%v", AlertsActivePath, v.Encode())
}

func GetAlertsHistoryPath(startTime *time.Time, endTime *time.Time, filterBy string, filterVal string, sortOrder string,
	showMitigations bool, showAlarms bool, showMatches bool, learningMode bool,
) string {
	timeFormatStr := "2006-01-02T15:04:05"
	v := url.Values{}
	if startTime != nil {
		v.Set("startTime", startTime.Format(timeFormatStr))
	}
	if endTime != nil {
		v.Set("endTime", endTime.Format(timeFormatStr))
	}
	if filterBy != "" {
		v.Set("filterBy", filterBy)
	}
	if filterVal != "" {
		v.Set("filterVal", filterVal)
	}
	if sortOrder != "" {
		v.Set("sortOrder", sortOrder)
	}
	v.Set("showMitigations", fmt.Sprint(boolToInt(showMitigations)))
	v.Set("showAlarms", fmt.Sprint(boolToInt(showAlarms)))
	v.Set("showMatches", fmt.Sprint(boolToInt(showMatches)))
	v.Set("learningMode", fmt.Sprint(boolToInt(learningMode)))

	return fmt.Sprintf("%v?%v", AlertsHistoryPath, v.Encode())
}

func boolToInt(val bool) int {
	var intVal int
	if val {
		intVal = 1
	}
	return intVal
}
