package api_resources_test

import (
	"context"
	"testing"
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_resources"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/utils"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateCustomDimension(t *testing.T) {
	// arrange
	createResponsePayload := `
	{
		"customDimension": {
			"id": 42,
			"name": "c_testapi_dimension_1",
			"display_name": "dimension_display_name",
			"type": "string",
			"company_id": "74333",
			"populators": []
		}
	}`
	transport := &api_connection.StubTransport{ResponseBody: createResponsePayload}
	customDimensionsAPI := api_resources.NewCustomDimensionsAPI(transport)
	dimension := models.NewCustomDimension("c_testapi_dimension_1", "dimension_display_name", models.CustomDimensionTypeStr)

	// act
	created, err := customDimensionsAPI.Create(context.Background(), *dimension)

	// assert request properly formed
	assert := assert.New(t)
	require := require.New(t)

	require.NoError(err)
	payload := utils.NewJSONPayloadInspector(t, transport.RequestBody)
	assert.Equal("c_testapi_dimension_1", payload.String("name"))
	assert.Equal("dimension_display_name", payload.String("display_name"))
	assert.Equal("string", payload.String("type"))

	// and response properly parsed
	assert.Equal(models.ID(42), created.ID)
	assert.Equal("c_testapi_dimension_1", created.Name)
	assert.Equal("dimension_display_name", created.DisplayName)
	assert.Equal(models.CustomDimensionTypeStr, created.Type)
	assert.Equal(models.ID(74333), created.CompanyID)
	assert.Len(created.Populators, 0)
}

func TestUpdateCustomDimension(t *testing.T) {
	// arrange
	updateResponsePayload := `
	{
		"customDimension": {
			"id": 42,
			"name": "c_testapi_dimension_1",
			"display_name": "dimension_display_name2",
			"type": "string",
			"company_id": "74333",
			"populators": []
		}
	}`
	transport := &api_connection.StubTransport{ResponseBody: updateResponsePayload}
	customDimensionsAPI := api_resources.NewCustomDimensionsAPI(transport)
	dimensionID := models.ID(42)
	dimension := models.CustomDimension{ID: dimensionID, DisplayName: "dimension_display_name2"}

	// act
	updated, err := customDimensionsAPI.Update(context.Background(), dimension)

	// assert request properly formed
	assert := assert.New(t)
	require := require.New(t)

	require.NoError(err)
	payload := utils.NewJSONPayloadInspector(t, transport.RequestBody)
	assert.Equal("dimension_display_name2", payload.String("display_name"))

	// # and response properly parsed
	assert.Equal(models.ID(42), updated.ID)
	assert.Equal("c_testapi_dimension_1", updated.Name)
	assert.Equal("dimension_display_name2", updated.DisplayName)
	assert.Equal(models.CustomDimensionTypeStr, updated.Type)
	assert.Equal(models.ID(74333), updated.CompanyID)
	assert.Len(updated.Populators, 0)
}

func TestGetCustomDimension(t *testing.T) {
	// arrange
	getResponsePayload := `
	{
		"customDimension": {
			"id": 42,
			"name": "c_testapi_dimension_1",
			"display_name": "dimension_display_name",
			"type": "string",
			"company_id": "74333",
			"populators": [
				{
					"id": 1510871096,
					"dimension_id": 24001,
					"value": "testapi-dimension-value-1",
					"direction": "DST",
					"device_name": "128.0.0.100,device1",
					"interface_name": "interface1,interface2",
					"addr": "128.0.0.1/32,128.0.0.2/32",
					"addr_count": 2,
					"port": "1001,1002",
					"tcp_flags": "160",
					"protocol": "6,17",
					"asn": "101,102",
					"nexthop_asn": "201,202",
					"nexthop": "128.0.200.1/32,128.0.200.2/32",
					"bgp_aspath": "3001,3002",
					"bgp_community": "401:499,501:599",
					"user": "144319",
					"created_date": "2020-12-15T08:32:19.503788Z",
					"updated_date": "2020-12-15T08:32:19.503788Z",
					"company_id": "74333",
					"device_type": "device-type1",
					"site": "site1,site2,site3",
					"lasthop_as_name": "asn101,asn102",
					"nexthop_as_name": "asn201,asn202",
					"mac": "FF:FF:FF:FF:FF:FA,FF:FF:FF:FF:FF:FF",
					"mac_count": 2,
					"country": "NL,SE",
					"vlans": "2001,2002"
				},
				{
					"id": 1510862280,
					"dimension_id": 24001,
					"value": "testapi-dimension-value-3",
					"direction": "SRC",
					"addr_count": 0,
					"user": "144319",
					"created_date": "2020-12-15T07:55:23.0Z",
					"updated_date": "2020-12-15T11:11:30.0Z",
					"company_id": "74333",
					"site": "site3",
					"mac_count": 0
				}
			]
		}      
	}`
	transport := &api_connection.StubTransport{ResponseBody: getResponsePayload}
	customDimensionsAPI := api_resources.NewCustomDimensionsAPI(transport)
	dimensionID := models.ID(42)

	// act
	dimension, err := customDimensionsAPI.Get(context.Background(), dimensionID)

	// assert request properly formed
	assert := assert.New(t)
	require := require.New(t)

	require.NoError(err)
	assert.Zero(transport.RequestBody)

	// and response properly parsed
	assert.Equal(models.ID(42), dimension.ID)
	assert.Equal("c_testapi_dimension_1", dimension.Name)
	assert.Equal("dimension_display_name", dimension.DisplayName)
	assert.Equal(models.CustomDimensionTypeStr, dimension.Type)
	assert.Equal(models.ID(74333), dimension.CompanyID)
	assert.Len(dimension.Populators, 2)
	assert.Equal(models.ID(1510862280), dimension.Populators[1].ID)
	assert.Equal(models.ID(24001), dimension.Populators[1].DimensionID)
	assert.Equal("testapi-dimension-value-3", dimension.Populators[1].Value)
	assert.Equal(models.PopulatorDirectionSrc, dimension.Populators[1].Direction)
	assert.Equal(0, dimension.Populators[1].AddrCount)
	assert.Equal("144319", *dimension.Populators[1].User)
	assert.Equal(time.Date(2020, 12, 15, 7, 55, 23, 0, time.UTC), dimension.Populators[1].CreatedDate)
	assert.Equal(time.Date(2020, 12, 15, 11, 11, 30, 0, time.UTC), dimension.Populators[1].UpdatedDate)
	assert.Equal(models.ID(74333), dimension.Populators[1].CompanyID)
	assert.Equal("site3", *dimension.Populators[1].Site)
	assert.Equal(0, dimension.Populators[1].MACCount)
}

func TestGetAllCustomDimensions(t *testing.T) {
	// arrange
	getResponsePayload := `
	{
		"customDimensions": [
			{
				"id": 42,
				"name": "c_testapi_dimension_1",
				"display_name": "dimension_display_name1",
				"type": "string",
				"populators": [],
				"company_id": "74333"
			},
			{
				"id": 43,
				"name": "c_testapi_dimension_2",
				"display_name": "dimension_display_name2",
				"type": "uint32",
				"company_id": "74334",
				"populators": [
					{
						"id": 1510862280,
						"dimension_id": 24001,
						"value": "testapi-dimension-value-3",
						"device_type": "device-type3",
						"direction": "SRC",
						"interface_name": "interface3",
						"addr_count": 0,
						"user": "144319",
						"created_date": "2020-12-15T07:55:23.0Z",
						"updated_date": "2020-12-15T10:50:22.0Z",
						"company_id": "74333",
						"site": "site3",
						"mac_count": 0
					}
				]
			}
		]
	}`
	transport := &api_connection.StubTransport{ResponseBody: getResponsePayload}
	customDimensionsAPI := api_resources.NewCustomDimensionsAPI(transport)

	// act
	dimensions, err := customDimensionsAPI.GetAll(context.Background())

	// assert request properly formed
	assert := assert.New(t)
	require := require.New(t)

	require.NoError(err)
	assert.Zero(transport.RequestBody)

	// and response properly parsed
	require.Equal(2, len(dimensions))

	// dimension 0
	assert.Equal(models.ID(42), dimensions[0].ID)
	assert.Equal("c_testapi_dimension_1", dimensions[0].Name)
	assert.Equal("dimension_display_name1", dimensions[0].DisplayName)
	assert.Equal(models.CustomDimensionTypeStr, dimensions[0].Type)

	// dimension 1
	assert.Equal(models.ID(43), dimensions[1].ID)
	assert.Equal("c_testapi_dimension_2", dimensions[1].Name)
	assert.Equal("dimension_display_name2", dimensions[1].DisplayName)
	assert.Equal(models.CustomDimensionTypeUint32, dimensions[1].Type)
	assert.Equal(models.ID(74334), dimensions[1].CompanyID)
	assert.Len(dimensions[1].Populators, 1)
	assert.Equal(models.ID(1510862280), dimensions[1].Populators[0].ID)
	assert.Equal(models.ID(24001), dimensions[1].Populators[0].DimensionID)
	assert.Equal("testapi-dimension-value-3", dimensions[1].Populators[0].Value)
	assert.Equal(models.PopulatorDirectionSrc, dimensions[1].Populators[0].Direction)
	assert.Equal(0, dimensions[1].Populators[0].AddrCount)
	assert.Equal("144319", *dimensions[1].Populators[0].User)
	assert.Equal(time.Date(2020, 12, 15, 7, 55, 23, 0, time.UTC), dimensions[1].Populators[0].CreatedDate)
	assert.Equal(time.Date(2020, 12, 15, 10, 50, 22, 0, time.UTC), dimensions[1].Populators[0].UpdatedDate)
	assert.Equal(models.ID(74333), dimensions[1].Populators[0].CompanyID)
	assert.Equal("site3", *dimensions[1].Populators[0].Site)
	assert.Equal(0, dimensions[1].Populators[0].MACCount)
}

func TestDeleteCustomDimension(t *testing.T) {
	// arrange
	deleteResponsePayload := "" // deleting device responds with empty body
	transport := &api_connection.StubTransport{ResponseBody: deleteResponsePayload}
	customDimensionsAPI := api_resources.NewCustomDimensionsAPI(transport)

	// act
	dimensionID := models.ID(42)
	err := customDimensionsAPI.Delete(context.Background(), dimensionID)

	// assert
	assert := assert.New(t)
	require := require.New(t)
	require.NoError(err)
	assert.Zero(transport.RequestBody)
}
