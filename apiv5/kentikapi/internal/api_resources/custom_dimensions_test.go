package api_resources_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/kentik/community_sdk_golang/apiv5/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/apiv5/kentikapi/internal/api_resources"
	"github.com/kentik/community_sdk_golang/apiv5/kentikapi/internal/testutil"
	"github.com/kentik/community_sdk_golang/apiv5/kentikapi/internal/utils"
	"github.com/kentik/community_sdk_golang/apiv5/kentikapi/models"
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
	assert.Equal(http.MethodPost, transport.RequestMethod)
	assert.Equal("/customdimension", transport.RequestPath)
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
	assert.Equal(http.MethodPut, transport.RequestMethod)
	assert.Equal(fmt.Sprintf("/customdimension/%v", dimensionID), transport.RequestPath)
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
	tests := []struct {
		name           string
		transportError error
		responseBody   string
		expectedResult *models.CustomDimension
		expectedError  bool
	}{
		{
			name:           "transport error",
			transportError: assert.AnError,
			expectedError:  true,
		}, {
			name:          "invalid response format",
			responseBody:  "invalid JSON",
			expectedError: true,
		}, {
			name:          "empty response",
			responseBody:  "{}",
			expectedError: true,
		}, {
			name: "custom dimension returned",
			responseBody: `{
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
			}`,
			expectedResult: &models.CustomDimension{
				DisplayName: "dimension_display_name",
				Name:        "c_testapi_dimension_1",
				Type:        models.CustomDimensionTypeStr,
				ID:          42,
				CompanyID:   74333,
				Populators: []models.Populator{{
					Value:         "testapi-dimension-value-1",
					Direction:     models.PopulatorDirectionDst,
					DeviceName:    "128.0.0.100,device1",
					InterfaceName: testutil.StringPtr("interface1,interface2"),
					Addr:          testutil.StringPtr("128.0.0.1/32,128.0.0.2/32"),
					Port:          testutil.StringPtr("1001,1002"),
					TCPFlags:      testutil.StringPtr("160"),
					Protocol:      testutil.StringPtr("6,17"),
					ASN:           testutil.StringPtr("101,102"),
					NextHopASN:    testutil.StringPtr("201,202"),
					NextHop:       testutil.StringPtr("128.0.200.1/32,128.0.200.2/32"),
					BGPAsPath:     testutil.StringPtr("3001,3002"),
					BGPCommunity:  testutil.StringPtr("401:499,501:599"),
					DeviceType:    testutil.StringPtr("device-type1"),
					Site:          testutil.StringPtr("site1,site2,site3"),
					LastHopAsName: testutil.StringPtr("asn101,asn102"),
					NextHopAsName: testutil.StringPtr("asn201,asn202"),
					MAC:           testutil.StringPtr("FF:FF:FF:FF:FF:FA,FF:FF:FF:FF:FF:FF"),
					Country:       testutil.StringPtr("NL,SE"),
					VLans:         testutil.StringPtr("2001,2002"),
					ID:            1510871096,
					CompanyID:     74333,
					DimensionID:   24001,
					User:          testutil.StringPtr("144319"),
					MACCount:      2,
					AddrCount:     2,
					CreatedDate:   time.Date(2020, 12, 15, 8, 32, 19, 503788000, time.UTC),
					UpdatedDate:   time.Date(2020, 12, 15, 8, 32, 19, 503788000, time.UTC),
				}, {
					Value:       "testapi-dimension-value-3",
					Direction:   models.PopulatorDirectionSrc,
					Site:        testutil.StringPtr("site3"),
					ID:          1510862280,
					CompanyID:   74333,
					DimensionID: 24001,
					User:        testutil.StringPtr("144319"),
					CreatedDate: time.Date(2020, 12, 15, 7, 55, 23, 0, time.UTC),
					UpdatedDate: time.Date(2020, 12, 15, 11, 11, 30, 0, time.UTC),
				}},
			},
		}, {
			name: "custom dimension with unknown enums returned",
			responseBody: `{
				"customDimension": {
					"id": 42,
					"name": "c_testapi_dimension_1",
					"display_name": "dimension_display_name",
					"type": "t_teapot",
					"company_id": "74333",
					"populators": [
						{
							"id": 1510862280,
							"dimension_id": 24001,
							"value": "testapi-dimension-value-3",
							"direction": "d_teapot",
							"addr_count": 0,
							"user": "144319",
							"created_date": "2020-12-15T07:55:23.0Z",
							"updated_date": "2020-12-15T11:11:30.0Z",
							"company_id": "74333",
							"mac_count": 0
						}
					]
				}      
			}`,
			expectedResult: &models.CustomDimension{
				DisplayName: "dimension_display_name",
				Name:        "c_testapi_dimension_1",
				Type:        models.CustomDimensionType("t_teapot"),
				ID:          42,
				CompanyID:   74333,
				Populators: []models.Populator{{
					Value:       "testapi-dimension-value-3",
					Direction:   models.PopulatorDirection("d_teapot"),
					ID:          1510862280,
					CompanyID:   74333,
					DimensionID: 24001,
					User:        testutil.StringPtr("144319"),
					CreatedDate: time.Date(2020, 12, 15, 7, 55, 23, 0, time.UTC),
					UpdatedDate: time.Date(2020, 12, 15, 11, 11, 30, 0, time.UTC),
				}},
			},
		}, {
			name: "custom dimension with empty enums returned",
			responseBody: `{
				"customDimension": {
					"id": 42,
					"name": "c_testapi_dimension_1",
					"display_name": "dimension_display_name",
					"type": "",
					"company_id": "74333",
					"populators": [
						{
							"id": 1510862280,
							"dimension_id": 24001,
							"value": "testapi-dimension-value-3",
							"direction": "",
							"addr_count": 0,
							"user": "144319",
							"created_date": "2020-12-15T07:55:23.0Z",
							"updated_date": "2020-12-15T11:11:30.0Z",
							"company_id": "74333",
							"mac_count": 0
						}
					]
				}      
			}`,
			expectedResult: &models.CustomDimension{
				DisplayName: "dimension_display_name",
				Name:        "c_testapi_dimension_1",
				Type:        models.CustomDimensionType(""),
				ID:          42,
				CompanyID:   74333,
				Populators: []models.Populator{{
					Value:       "testapi-dimension-value-3",
					Direction:   models.PopulatorDirection(""),
					ID:          1510862280,
					CompanyID:   74333,
					DimensionID: 24001,
					User:        testutil.StringPtr("144319"),
					CreatedDate: time.Date(2020, 12, 15, 7, 55, 23, 0, time.UTC),
					UpdatedDate: time.Date(2020, 12, 15, 11, 11, 30, 0, time.UTC),
				}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			transport := &api_connection.StubTransport{ResponseBody: tt.responseBody}
			customDimensionsAPI := api_resources.NewCustomDimensionsAPI(transport)
			dimensionID := 42

			// act
			dimension, err := customDimensionsAPI.Get(context.Background(), dimensionID)

			// assert
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, http.MethodGet, transport.RequestMethod)
			assert.Equal(t, fmt.Sprintf("/customdimension/%v", dimensionID), transport.RequestPath)
			assert.Zero(t, transport.RequestBody)

			assert.Equal(t, tt.expectedResult, dimension)
		})
	}
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
	assert.Equal(http.MethodGet, transport.RequestMethod)
	assert.Equal("/customdimensions", transport.RequestPath)
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
	assert.Equal(http.MethodDelete, transport.RequestMethod)
	assert.Equal(fmt.Sprintf("/customdimension/%v", dimensionID), transport.RequestPath)
	assert.Zero(transport.RequestBody)
}

func TestCreatePopulator(t *testing.T) {
	// arrange
	createResponsePayload := `
	{
		"populator": {
			"dimension_id": 24001,
			"value": "testapi-dimension-value-1",
			"direction": "DST",
			"device_name": "128.0.0.100,device1",
			"interface_name": "interface1,interface2",
			"addr": "128.0.0.1/32,128.0.0.2/32",
			"port": "1001,1002",
			"tcp_flags": "160",
			"protocol": "6,17",
			"asn": "101,102",
			"nexthop_asn": "201,202",
			"nexthop": "128.0.200.1/32,128.0.200.2/32",
			"bgp_aspath": "3001,3002",
			"bgp_community": "401:499,501:599",
			"device_type": "device-type1",
			"site": "site1,site2,site3",
			"lasthop_as_name": "asn101,asn102",
			"nexthop_as_name": "asn201,asn202",
			"mac": "FF:FF:FF:FF:FF:FA,FF:FF:FF:FF:FF:FF",
			"country": "NL,SE",
			"vlans": "2001,2002",
			"id": 1510862280,
			"company_id": "74333",
			"user": "144319",
			"mac_count": 2,
			"addr_count": 2,
			"created_date": "2020-12-15T07:55:23.0Z",
			"updated_date": "2020-12-15T07:55:23.0Z"
		}
	}`
	transport := &api_connection.StubTransport{ResponseBody: createResponsePayload}
	customDimensionsAPI := api_resources.NewCustomDimensionsAPI(transport)
	dimensionID := models.ID(24001)
	populator := models.NewPopulator(dimensionID, "testapi-dimension-value-1", "device1,128.0.0.100", models.PopulatorDirectionDst)
	models.SetOptional(&populator.InterfaceName, "interface1,interface2")
	models.SetOptional(&populator.Addr, "128.0.0.1/32,128.0.0.2/32")
	models.SetOptional(&populator.Port, "1001,1002")
	models.SetOptional(&populator.TCPFlags, "160")
	models.SetOptional(&populator.Protocol, "6,17")
	models.SetOptional(&populator.ASN, "101,102")
	models.SetOptional(&populator.NextHopASN, "201,202")
	models.SetOptional(&populator.NextHop, "128.0.200.1/32,128.0.200.2/32")
	models.SetOptional(&populator.BGPAsPath, "3001,3002")
	models.SetOptional(&populator.BGPCommunity, "401:499,501:599")
	models.SetOptional(&populator.DeviceType, "device-type1")
	models.SetOptional(&populator.Site, "site1,site2,site3")
	models.SetOptional(&populator.LastHopAsName, "asn101,asn102")
	models.SetOptional(&populator.NextHopAsName, "asn201,asn202")
	models.SetOptional(&populator.MAC, "FF:FF:FF:FF:FF:FA,FF:FF:FF:FF:FF:FF")
	models.SetOptional(&populator.Country, "NL,SE")
	models.SetOptional(&populator.VLans, "2001,2002")

	// act
	created, err := customDimensionsAPI.Populators.Create(context.Background(), *populator)

	// assert request properly formed
	assert := assert.New(t)
	require := require.New(t)

	require.NoError(err)
	assert.Equal(http.MethodPost, transport.RequestMethod)
	assert.Equal(fmt.Sprintf("/customdimension/%v/populator", dimensionID), transport.RequestPath)
	payload := utils.NewJSONPayloadInspector(t, transport.RequestBody)
	require.NotNil(payload.Get("populator"))
	assert.Equal("testapi-dimension-value-1", payload.String("populator/value"))
	assert.Equal("DST", payload.String("populator/direction"))
	assert.Equal("device1,128.0.0.100", payload.String("populator/device_name"))
	assert.Equal("device-type1", payload.String("populator/device_type"))
	assert.Equal("site1,site2,site3", payload.String("populator/site"))
	assert.Equal("interface1,interface2", payload.String("populator/interface_name"))
	assert.Equal("128.0.0.1/32,128.0.0.2/32", payload.String("populator/addr"))
	assert.Equal("1001,1002", payload.String("populator/port"))
	assert.Equal("160", payload.String("populator/tcp_flags"))
	assert.Equal("6,17", payload.String("populator/protocol"))
	assert.Equal("101,102", payload.String("populator/asn"))
	assert.Equal("asn101,asn102", payload.String("populator/lasthop_as_name"))
	assert.Equal("201,202", payload.String("populator/nexthop_asn"))
	assert.Equal("asn201,asn202", payload.String("populator/nexthop_as_name"))
	assert.Equal("128.0.200.1/32,128.0.200.2/32", payload.String("populator/nexthop"))
	assert.Equal("3001,3002", payload.String("populator/bgp_aspath"))
	assert.Equal("401:499,501:599", payload.String("populator/bgp_community"))
	assert.Equal("FF:FF:FF:FF:FF:FA,FF:FF:FF:FF:FF:FF", payload.String("populator/mac"))
	assert.Equal("NL,SE", payload.String("populator/country"))
	assert.Equal("2001,2002", payload.String("populator/vlans"))

	// and response properly parsed
	assert.Equal(models.ID(1510862280), created.ID)
	assert.Equal(models.ID(24001), created.DimensionID)
	assert.Equal("testapi-dimension-value-1", created.Value)
	assert.Equal(models.PopulatorDirectionDst, created.Direction)
	assert.Equal("128.0.0.100,device1", created.DeviceName)
	assert.Equal("interface1,interface2", *created.InterfaceName)
	assert.Equal("128.0.0.1/32,128.0.0.2/32", *created.Addr)
	assert.Equal("1001,1002", *created.Port)
	assert.Equal("160", *created.TCPFlags)
	assert.Equal("6,17", *created.Protocol)
	assert.Equal("101,102", *created.ASN)
	assert.Equal("201,202", *created.NextHopASN)
	assert.Equal("128.0.200.1/32,128.0.200.2/32", *created.NextHop)
	assert.Equal("3001,3002", *created.BGPAsPath)
	assert.Equal("401:499,501:599", *created.BGPCommunity)
	assert.Equal("device-type1", *created.DeviceType)
	assert.Equal("site1,site2,site3", *created.Site)
	assert.Equal("asn101,asn102", *created.LastHopAsName)
	assert.Equal("asn201,asn202", *created.NextHopAsName)
	assert.Equal("FF:FF:FF:FF:FF:FA,FF:FF:FF:FF:FF:FF", *created.MAC)
	assert.Equal("NL,SE", *created.Country)
	assert.Equal("2001,2002", *created.VLans)
	assert.Equal(models.ID(74333), created.CompanyID)
	assert.Equal("144319", *created.User)
	assert.Equal(2, created.MACCount)
	assert.Equal(2, created.AddrCount)
	assert.Equal(time.Date(2020, 12, 15, 7, 55, 23, 0, time.UTC), created.CreatedDate)
	assert.Equal(time.Date(2020, 12, 15, 7, 55, 23, 0, time.UTC), created.UpdatedDate)
}

func TestUpdatePopulator(t *testing.T) {
	// arrange
	updateResponsePayload := `
	{
        "populator": {
            "id": 1510862280,
            "dimension_id": 24001,
            "value": "testapi-dimension-value-3",
            "direction": "SRC",
            "interface_name": "interface3",
            "addr_count": 0,
            "tcp_flags": "12",
            "protocol": "17",
            "user": "144319",
            "created_date": "2020-12-15T07:55:23.0Z",
            "updated_date": "2020-12-15T10:50:22.0Z",
            "company_id": "74333",
            "device_type": "device-type3",
            "site": "site3",
            "mac_count": 0
        }
	}`

	transport := &api_connection.StubTransport{ResponseBody: updateResponsePayload}
	customDimensionsAPI := api_resources.NewCustomDimensionsAPI(transport)
	dimensionID := models.ID(24001)
	populatorID := models.ID(1510862280)

	populator := models.Populator{
		DimensionID: dimensionID,
		ID:          populatorID,
		Value:       "testapi-dimension-value-3",
		Direction:   models.PopulatorDirectionSrc,
	}
	models.SetOptional(&populator.InterfaceName, "interface3")
	models.SetOptional(&populator.TCPFlags, "12")
	models.SetOptional(&populator.Protocol, "17")
	models.SetOptional(&populator.DeviceType, "device-type3")
	models.SetOptional(&populator.Site, "site3")

	// act
	updated, err := customDimensionsAPI.Populators.Update(context.Background(), populator)

	// assert request properly formed
	assert := assert.New(t)
	require := require.New(t)

	require.NoError(err)
	assert.Equal(http.MethodPut, transport.RequestMethod)
	assert.Equal(fmt.Sprintf("/customdimension/%v/populator/%v", dimensionID, populatorID), transport.RequestPath)
	payload := utils.NewJSONPayloadInspector(t, transport.RequestBody)
	require.NotNil(payload.Get("populator"))
	assert.Equal("testapi-dimension-value-3", payload.String("populator/value"))
	assert.Equal("SRC", payload.String("populator/direction"))
	assert.Equal("device-type3", payload.String("populator/device_type"))
	assert.Equal("site3", payload.String("populator/site"))
	assert.Equal("interface3", payload.String("populator/interface_name"))
	assert.Equal("12", payload.String("populator/tcp_flags"))
	assert.Equal("17", payload.String("populator/protocol"))

	// and response properly parsed
	assert.Equal(models.ID(1510862280), updated.ID)
	assert.Equal(models.ID(24001), updated.DimensionID)
	assert.Equal("testapi-dimension-value-3", updated.Value)
	assert.Equal(models.PopulatorDirectionSrc, updated.Direction)
	assert.Equal("interface3", *updated.InterfaceName)
	assert.Equal("12", *updated.TCPFlags)
	assert.Equal("17", *updated.Protocol)
	assert.Equal("device-type3", *updated.DeviceType)
	assert.Equal("site3", *updated.Site)
	assert.Equal(models.ID(74333), updated.CompanyID)
	assert.Equal("144319", *updated.User)
	assert.Equal(0, updated.MACCount)
	assert.Equal(0, updated.AddrCount)
	assert.Equal(time.Date(2020, 12, 15, 7, 55, 23, 0, time.UTC), updated.CreatedDate)
	assert.Equal(time.Date(2020, 12, 15, 10, 50, 22, 0, time.UTC), updated.UpdatedDate)
}

func TestDeletePopulator(t *testing.T) {
	// arrange
	deleteResponsePayload := "" // deleting device responds with empty body
	transport := &api_connection.StubTransport{ResponseBody: deleteResponsePayload}
	customDimensionsAPI := api_resources.NewCustomDimensionsAPI(transport)

	// act
	dimensionID := models.ID(42)
	populatorID := models.ID(5012)
	err := customDimensionsAPI.Populators.Delete(context.Background(), dimensionID, populatorID)

	// assert
	assert := assert.New(t)
	require := require.New(t)
	require.NoError(err)
	assert.Equal(http.MethodDelete, transport.RequestMethod)
	assert.Equal(fmt.Sprintf("/customdimension/%v/populator/%v", dimensionID, populatorID), transport.RequestPath)
	assert.Zero(transport.RequestBody)
}
