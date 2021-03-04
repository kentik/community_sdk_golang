package api_endpoints

import "fmt"

const (
	ManualMitigationPath = "/alerts/manual-mitigate"
	AlertsActivePath     = "/alerts-active/alarms"
	AlertsHistoryPath    = "/alerts-active/alerts-history"
)

func GetActiveAlertsPath(startTime string, endTime string, filterBy string, filterVal string, showMitigations int, showAlarms int, showMatches int, learningMode int) string {
	return fmt.Sprintf("%v?startTime=%v&endTime=%v&filterBy=%v&filterVal=%v&showMitigations=%v&showAlarms=%v&showMatches=%v&learningMode=%v",
		AlertsActivePath, startTime, endTime, filterBy, filterVal, showMitigations, showAlarms, showMatches, learningMode)
}

func GetAlertsHistoryPath(startTime string, endTime string, filterBy string, filterVal string, sortOrder string,
	showMitigations int, showAlarms int, showMatches int, learningMode int) string {
	return fmt.Sprintf("%v?startTime=%v&endTime=%v&filterBy=%v&filterVal=%v&sortOrder=%v&showMitigations=%v&showAlarms=%v&showMatches=%v&learningMode=%v",
		AlertsActivePath, startTime, endTime, filterBy, filterVal, sortOrder, showMitigations, showAlarms, showMatches, learningMode)
}
