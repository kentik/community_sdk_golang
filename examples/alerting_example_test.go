//+build examples

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
	client := NewClient()

	fmt.Println("GET ALERTS")
	startTime := time.Date(2020, time.January, 1, 12, 0, 0, 0, time.Local)
	endTime := time.Date(2021, time.March, 8, 15, 30, 0, 0, time.Local)
	alerts, err := client.Alerting.GetActiveAlerts(context.Background(), startTime, endTime, "", "", true, true, false, false)
	if err != nil {
		return err
	}
	PrettyPrint(alerts)
	fmt.Println()

	fmt.Println("GET ALERTS HISTORY")
	history, err := client.Alerting.GetAlertsHistory(context.Background(), startTime, endTime, "", "", "", true, true, false, false)
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
