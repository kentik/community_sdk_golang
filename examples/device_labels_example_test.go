//+build examples

package examples

import (
	"context"
	"fmt"
	"testing"

	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
)

func TestDeviceLabelsAPIExample(t *testing.T) {
	assert.NoError(t, runGetAllDeviceLabels())
	assert.NoError(t, runCRUDDeviceLabels())
}

func runGetAllDeviceLabels() error {
	client, err := NewClient()
	if err != nil {
		return err
	}

	fmt.Println("### GET ALL")
	labels, err := client.DeviceLabels.GetAll(context.Background())
	if err != nil {
		return err
	}
	PrettyPrint(labels)
	fmt.Println()

	return nil
}

func runCRUDDeviceLabels() error {
	client, err := NewClient()
	if err != nil {
		return err
	}

	fmt.Println("### CREATE")
	label := models.NewDeviceLabel("apitest-device_label-1", "#00FF00")

	created, err := client.DeviceLabels.Create(context.Background(), *label)
	if err != nil {
		return err
	}
	PrettyPrint(created)
	fmt.Println()

	fmt.Println("### UPDATE")
	created.Name = "apitest-device_label-one"
	created.Color = "#AA00CC"
	updated, err := client.DeviceLabels.Update(context.Background(), *created)
	if err != nil {
		return err
	}
	PrettyPrint(updated)
	fmt.Println()

	fmt.Println("### GET")
	got, err := client.DeviceLabels.Get(context.Background(), created.ID)
	if err != nil {
		return err
	}
	PrettyPrint(got)
	fmt.Println()

	fmt.Println("### DELETE")
	err = client.DeviceLabels.Delete(context.Background(), created.ID)
	if err != nil {
		return err
	}
	fmt.Println("Success")
	fmt.Println()

	return nil
}
