package resources_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/resources"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/utils"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAll(t *testing.T) {
	t.Parallel()

	getResponsePayload := `
	[
		{
			"id": 42,
			"company_id": "74333",
			"user_id": "144319",
			"name": "apitest-customapp-1",
			"description": "TESTING CUSTOM APPS 1",
			"ip_range": "192.168.0.1,192.168.0.2",
			"protocol": "6,17",
			"port": "9001,9002,9003",
			"asn": "asn1,asn2,asn3",
			"cdate": "2020-12-11T07:07:20.968Z",
			"edate": "2020-12-11T07:07:20.968Z"
		},
		{
			"id": 43,
			"company_id": "74333",
			"user_id": "144319",
			"name": "apitest-customapp-2",
			"description": "TESTING CUSTOM APPS 2",
			"ip_range": "192.168.0.3,192.168.0.4",
			"protocol": "6,17",
			"port": "9011,9012,9013",
			"asn": "asn4,asn5,asn6",
			"cdate": "2020-12-11T07:08:20.968Z",
			"edate": "2020-12-11T07:08:20.968Z"
		}
	]`
	transport := &api_connection.StubTransport{ResponseBody: getResponsePayload}
	applicationsAPI := resources.NewCustomApplicationsAPI(transport)

	// act
	applications, err := applicationsAPI.GetAll(context.Background())

	// assert request properly formed
	assert := assert.New(t)
	require := require.New(t)

	require.NoError(err)
	assert.Equal("/customApplications", transport.RequestPath)
	assert.Zero(transport.RequestBody)

	// and response properly parsed
	assert.Len(applications, 2)

	app1 := applications[1]
	assert.Equal(models.ID(43), app1.ID)
	assert.Equal(models.ID(74333), app1.CompanyID)
	assert.Equal(models.ID(144319), *app1.UserID)
	assert.Equal("apitest-customapp-2", app1.Name)
	assert.Equal("TESTING CUSTOM APPS 2", *app1.Description)
	assert.Equal("192.168.0.3,192.168.0.4", *app1.IPRange)
	assert.Equal("6,17", *app1.Protocol)
	assert.Equal("9011,9012,9013", *app1.Port)
	assert.Equal("asn4,asn5,asn6", *app1.ASN)
	assert.Equal(time.Date(2020, 12, 11, 7, 8, 20, 968*1000000, time.UTC), *app1.CreatedDate)
	assert.Equal(time.Date(2020, 12, 11, 7, 8, 20, 968*1000000, time.UTC), *app1.UpdatedDate)
}

func TestCreateCustomApplication(t *testing.T) {
	t.Parallel()

	createResponsePayload := `
	{
		"name": "apitest-customapp-1",
		"description": "Testing custom application api",
		"ip_range": "192.168.0.1,192.168.0.2",
		"protocol": "6,17",
		"port": "9001,9002,9003",
		"asn": "asn1,asn2,asn3",
		"id": 207,
		"user_id": "144319",
		"company_id": "74333"
	}`
	transport := &api_connection.StubTransport{ResponseBody: createResponsePayload}
	applicationsAPI := resources.NewCustomApplicationsAPI(transport)

	app := models.NewCustomApplication("apitest-customapp-1")
	models.SetOptional(&app.Description, "Testing custom application api")
	models.SetOptional(&app.IPRange, "192.168.0.1,192.168.0.2")
	models.SetOptional(&app.Protocol, "6,17")
	models.SetOptional(&app.Port, "9001,9002,9003")
	models.SetOptional(&app.ASN, "asn1,asn2,asn3")

	// act
	created, err := applicationsAPI.Create(context.Background(), *app)

	// assert request properly formed
	assert := assert.New(t)
	require := require.New(t)

	require.NoError(err)
	assert.Equal("/customApplications", transport.RequestPath)
	payload := utils.NewJSONPayloadInspector(t, transport.RequestBody)
	assert.Equal("apitest-customapp-1", payload.String("name"))
	assert.Equal("Testing custom application api", payload.String("description"))
	assert.Equal("192.168.0.1,192.168.0.2", payload.String("ip_range"))
	assert.Equal("6,17", payload.String("protocol"))
	assert.Equal("9001,9002,9003", payload.String("port"))
	assert.Equal("asn1,asn2,asn3", payload.String("asn"))

	// and response properly parsed
	assert.Equal(models.ID(207), created.ID)
	assert.Equal(models.ID(74333), created.CompanyID)
	assert.Equal(models.ID(144319), *created.UserID)
	assert.Equal("apitest-customapp-1", created.Name)
	assert.Equal("Testing custom application api", *created.Description)
	assert.Equal("192.168.0.1,192.168.0.2", *created.IPRange)
	assert.Equal("6,17", *created.Protocol)
	assert.Equal("9001,9002,9003", *created.Port)
	assert.Equal("asn1,asn2,asn3", *created.ASN)
	assert.Nil(created.CreatedDate)
	assert.Nil(created.UpdatedDate)
}

func TestUpdateCustomApplication(t *testing.T) {
	t.Parallel()

	updateResponsePayload := `
	{
		"id": 207,
		"company_id": "74333",
		"user_id": "144319",
		"name": "apitest-customapp-one",
		"description": "TESTING CUSTOM APPS",
		"ip_range": "192.168.5.1,192.168.5.2",
		"protocol": "6,17",
		"port": "9011,9012,9013",
		"asn": "asn1,asn2,asn3",
		"cdate": "2020-12-11T07:07:20.968Z",
		"edate": "2020-12-11T07:07:20.968Z"
	}`
	transport := &api_connection.StubTransport{ResponseBody: updateResponsePayload}
	applicationsAPI := resources.NewCustomApplicationsAPI(transport)

	appID := models.ID(207)
	app := models.CustomApplication{
		ID:   appID,
		Name: "apitest-customapp-one",
	}
	models.SetOptional(&app.Description, "TESTING CUSTOM APPS")
	models.SetOptional(&app.IPRange, "192.168.5.1,192.168.5.2")
	models.SetOptional(&app.Port, "9011,9012,9013")

	// act
	updated, err := applicationsAPI.Update(context.Background(), app)

	// assert request properly formed
	assert := assert.New(t)
	require := require.New(t)

	require.NoError(err)
	assert.Equal(fmt.Sprintf("/customApplications/%v", appID), transport.RequestPath)
	payload := utils.NewJSONPayloadInspector(t, transport.RequestBody)
	assert.Equal("apitest-customapp-one", payload.String("name"))
	assert.Equal("TESTING CUSTOM APPS", payload.String("description"))
	assert.Equal("192.168.5.1,192.168.5.2", payload.String("ip_range"))
	assert.Nil(payload.Get("protocol"))
	assert.Equal("9011,9012,9013", payload.String("port"))
	assert.Nil(payload.Get("asn"))

	// and response properly parsed
	assert.Equal(models.ID(207), updated.ID)
	assert.Equal(models.ID(74333), updated.CompanyID)
	assert.Equal(models.ID(144319), *updated.UserID)
	assert.Equal("apitest-customapp-one", updated.Name)
	assert.Equal("TESTING CUSTOM APPS", *updated.Description)
	assert.Equal("192.168.5.1,192.168.5.2", *updated.IPRange)
	assert.Equal("6,17", *updated.Protocol)
	assert.Equal("9011,9012,9013", *updated.Port)
	assert.Equal("asn1,asn2,asn3", *updated.ASN)
	assert.Equal(time.Date(2020, 12, 11, 7, 7, 20, 968*1000000, time.UTC), *updated.CreatedDate)
	assert.Equal(time.Date(2020, 12, 11, 7, 7, 20, 968*1000000, time.UTC), *updated.UpdatedDate)
}

func TestDeleteCustomApplication(t *testing.T) {
	t.Parallel()

	// arrange
	deleteResponsePayload := "" // deleting custom application responds with empty body
	transport := &api_connection.StubTransport{ResponseBody: deleteResponsePayload}
	appliationsAPI := resources.NewCustomApplicationsAPI(transport)

	// act
	appID := models.ID(42)
	err := appliationsAPI.Delete(context.Background(), appID)

	// assert
	assert := assert.New(t)
	require := require.New(t)
	require.NoError(err)
	assert.Equal(fmt.Sprintf("/customApplications/%v", appID), transport.RequestPath)
	assert.Zero(transport.RequestBody)
}
