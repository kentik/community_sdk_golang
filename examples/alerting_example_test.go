//go:build examples
// +build examples

package examples

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
)

func TestAlertingAPIExample(t *testing.T) {
	assert.NoError(t, runCreateManualMitigation())
}

func runCreateManualMitigation() error {
	client, err := NewClient()
	if err != nil {
		return err
	}

	time1 := time.Date(2020, time.January, 1, 12, 0, 0, 0, time.Local)
	time2 := time.Date(2021, time.March, 8, 15, 30, 0, 0, time.Local)

	fmt.Println("GET ALERTS")
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
	alerts, err := client.Alerting.GetActiveAlerts(context.Background(), params)
	if err != nil {
		return err
	}
	PrettyPrint(alerts)
	fmt.Println()

	fmt.Println("GET ALERTS HISTORY")
	params.SortOrder = ""
	history, err := client.Alerting.GetAlertsHistory(context.Background(), params)
	if err != nil {
		return err
	}
	PrettyPrint(history)
	fmt.Println()

	fmt.Println("### CREATE MANUAL MITIGATION")
	// This data is invalid and saved filter will not be created.
	mm := models.ManualMitigation{
		IPCidr:                "192.168.0.0/24",
		PlatformID:            1234,
		MethodID:              12345,
		Comment:               nil,
		MinutesBeforeAutoStop: "20",
	}
	err = client.Alerting.CreateManualMitigation(context.Background(), mm)
	if err != nil {
		fmt.Println("Saved Filter not created")
	}
	fmt.Println()

	return nil
}
