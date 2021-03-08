package api_resources_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_resources"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/utils"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateDeviceLabel(t *testing.T) {
	// arrange
	createResponsePayload := `
	{
		"id": 42,
		"name": "apitest-device_label-1",
		"color": "#00FF00",
		"user_id": "52",
		"company_id": "72",
		"order": 0,
		"devices": [],
		"created_date": "2018-05-16T20:21:10.406Z",
		"updated_date": "2018-05-16T20:21:10.406Z"
	}`
	transport := &api_connection.StubTransport{ResponseBody: createResponsePayload}
	labelsAPI := api_resources.NewDeviceLabelsAPI(transport)
	label := models.NewDeviceLabel("apitest-device_label-1", "#00FF00")

	// act
	label, err := labelsAPI.Create(context.Background(), *label)

	// assert request properly formed
	assert := assert.New(t)
	require := require.New(t)

	require.NoError(err)
	assert.Equal("/deviceLabels", transport.RequestPath)
	payload := utils.NewJSONPayloadInspector(t, transport.RequestBody)
	assert.Equal("apitest-device_label-1", payload.String("name"))
	assert.Equal("#00FF00", payload.String("color"))

	// and response properly parsed
	assert.Equal(models.ID(42), label.ID)
	assert.Equal("apitest-device_label-1", label.Name)
	assert.Equal("#00FF00", label.Color)
	assert.Equal(models.ID(52), *label.UserID)
	assert.Equal(models.ID(72), label.CompanyID)
	assert.Equal(time.Date(2018, 5, 16, 20, 21, 10, 406*1000000, time.UTC), label.CreatedDate)
	assert.Equal(time.Date(2018, 5, 16, 20, 21, 10, 406*1000000, time.UTC), label.UpdatedDate)
	assert.Len(label.Devices, 0)
}

func TestUpdateDeviceLabel(t *testing.T) {
	// arrange
	updateResponsePayload := `
	{
		"id": 42,
		"name": "apitest-device_label-one",
		"color": "#AA00FF",
		"user_id": "52",
		"company_id": "72",
		"devices": [],
		"created_date": "2018-05-16T20:21:10.406Z",
		"updated_date": "2018-06-16T20:21:10.406Z"
	}`
	transport := &api_connection.StubTransport{ResponseBody: updateResponsePayload}
	labelsAPI := api_resources.NewDeviceLabelsAPI(transport)
	label := models.DeviceLabel{Name: "apitest-device_label-one", Color: "#AA00FF"}
	label.ID = models.ID(42)

	// act
	updated, err := labelsAPI.Update(context.Background(), label)

	// assert request properly formed
	assert := assert.New(t)
	require := require.New(t)

	require.NoError(err)
	assert.Equal(fmt.Sprintf("/deviceLabels/%v", label.ID), transport.RequestPath)
	payload := utils.NewJSONPayloadInspector(t, transport.RequestBody)
	assert.Equal("apitest-device_label-one", payload.String("name"))
	assert.Equal("#AA00FF", payload.String("color"))

	// and response properly parsed
	assert.Equal(models.ID(42), updated.ID)
	assert.Equal("apitest-device_label-one", updated.Name)
	assert.Equal("#AA00FF", updated.Color)
	assert.Equal(models.ID(52), *updated.UserID)
	assert.Equal(models.ID(72), updated.CompanyID)
	assert.Equal(time.Date(2018, 5, 16, 20, 21, 10, 406*1000000, time.UTC), updated.CreatedDate)
	assert.Equal(time.Date(2018, 6, 16, 20, 21, 10, 406*1000000, time.UTC), updated.UpdatedDate)
	assert.Len(updated.Devices, 0)
}

func TestGetLabel(t *testing.T) {
	// arrange
	getResponsePayload := `
	{
		"id": 32,
		"name": "ISP",
		"color": "#f1d5b9",
		"user_id": "52",
		"company_id": "72",
		"order": 0,
		"devices": [
			{
				"id": "42",
				"device_name": "my_device_1",
				"device_subtype": "router"
			}
		],
		"created_date": "2018-05-16T20:21:10.406Z",
		"updated_date": "2018-05-16T20:21:10.406Z"
	}`
	transport := &api_connection.StubTransport{ResponseBody: getResponsePayload}
	labelsAPI := api_resources.NewDeviceLabelsAPI(transport)
	labelID := models.ID(32)

	// act
	label, err := labelsAPI.Get(context.Background(), labelID)

	// assert request properly formed
	assert := assert.New(t)
	require := require.New(t)

	require.NoError(err)
	assert.Equal(fmt.Sprintf("/deviceLabels/%v", label.ID), transport.RequestPath)
	assert.Zero(transport.RequestBody)

	// and response properly parsed
	assert.Equal(models.ID(32), label.ID)
	assert.Equal("ISP", label.Name)
	assert.Equal("#f1d5b9", label.Color)
	assert.Equal(models.ID(52), *label.UserID)
	assert.Equal(models.ID(72), label.CompanyID)
	assert.Equal(time.Date(2018, 5, 16, 20, 21, 10, 406*1000000, time.UTC), label.CreatedDate)
	assert.Equal(time.Date(2018, 5, 16, 20, 21, 10, 406*1000000, time.UTC), label.UpdatedDate)
	assert.Len(label.Devices, 1)
	assert.Equal(models.ID(42), label.Devices[0].ID)
	assert.Equal("my_device_1", label.Devices[0].DeviceName)
	assert.Equal("router", label.Devices[0].DeviceSubtype)
	assert.Nil(label.Devices[0].DeviceType)
}

func TestGetAllLabels(t *testing.T) {
	// arrange
	getResponsePayload := `
    [
        {
            "id": 41,
            "name": "device_labels_1",
            "color": "#5289D9",
            "user_id": null,
            "company_id": "74333",
            "devices": [],
            "created_date": "2020-11-20T12:54:49.575Z",
            "updated_date": "2020-11-20T12:54:49.575Z"
        },
        {
            "id": 42,
            "name": "device_labels_2",
            "color": "#3F4EA0",
            "user_id": "136885",
            "company_id": "74333",
            "devices": [
                {
                    "id": "1",
                    "device_name": "device1",
                    "device_type": "type1",
                    "device_subtype": "subtype1"
                },
                {
                    "id": "2",
                    "device_name": "device2",
                    "device_type": "type2",
                    "device_subtype": "subtype2"
                }
            ],
            "created_date": "2020-11-20T13:45:27.430Z",
            "updated_date": "2020-11-20T13:45:27.430Z"
        }
    ]`
	transport := &api_connection.StubTransport{ResponseBody: getResponsePayload}
	labelsAPI := api_resources.NewDeviceLabelsAPI(transport)

	// act
	labels, err := labelsAPI.GetAll(context.Background())

	// assert request properly formed
	assert := assert.New(t)
	require := require.New(t)

	require.NoError(err)
	assert.Equal("/deviceLabels", transport.RequestPath)
	assert.Zero(transport.RequestBody)

	// and response properly parsed
	require.Len(labels, 2)

	assert.Equal(models.ID(41), labels[0].ID)
	assert.Equal("device_labels_1", labels[0].Name)
	assert.Equal("#5289D9", labels[0].Color)
	assert.Nil(labels[0].UserID)
	assert.Equal(models.ID(74333), labels[0].CompanyID)
	assert.Equal(time.Date(2020, 11, 20, 12, 54, 49, 575*1000000, time.UTC), labels[0].CreatedDate)
	assert.Equal(time.Date(2020, 11, 20, 12, 54, 49, 575*1000000, time.UTC), labels[0].UpdatedDate)
	assert.Len(labels[0].Devices, 0)

	assert.Equal(models.ID(42), labels[1].ID)
	assert.Equal("device_labels_2", labels[1].Name)
	assert.Equal("#3F4EA0", labels[1].Color)
	assert.Equal(models.ID(136885), *labels[1].UserID)
	assert.Equal(models.ID(74333), labels[1].CompanyID)
	assert.Equal(time.Date(2020, 11, 20, 13, 45, 27, 430*1000000, time.UTC), labels[1].CreatedDate)
	assert.Equal(time.Date(2020, 11, 20, 13, 45, 27, 430*1000000, time.UTC), labels[1].UpdatedDate)
	assert.Len(labels[1].Devices, 2)

	assert.Equal(models.ID(2), labels[1].Devices[1].ID)
	assert.Equal("device2", labels[1].Devices[1].DeviceName)
	assert.Equal("subtype2", labels[1].Devices[1].DeviceSubtype)
	assert.Equal("type2", *labels[1].Devices[1].DeviceType)
}

func TestDeleteDeviceLabel(t *testing.T) {
	// arrange
	deleteResponsePayload := `
	{
		"success": true
	}`
	transport := &api_connection.StubTransport{ResponseBody: deleteResponsePayload}
	labelsAPI := api_resources.NewDeviceLabelsAPI(transport)

	// act
	labelID := models.ID(42)
	err := labelsAPI.Delete(context.Background(), labelID)

	// assert
	require.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("/deviceLabels/%v", labelID), transport.RequestPath)
	assert.Zero(t, transport.RequestBody)
}
