package resources_test

import (
	"context"
	"testing"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/resources"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAllPlans(t *testing.T) {
	t.Parallel()

	// arrange
	getResponsePayload := `
	{
		"plans":[
			{
				"active":true,
				"bgp_enabled":true,
				"cdate":"2020-09-03T08:41:57.489Z",
				"company_id":74333,
				"description":"Your Free Trial includes 6 devices at a maximum of 1000 fps each. Please contact...",
				"deviceTypes":[
					{
						"device_type":"router"
					},
					{
						"device_type":"host-nprobe-dns-www"
					}
				],
				"devices":[
					{
						"id":"77714",
						"device_name":"testapi_router_minimal_1",
						"device_type":"router"
					},
					{
						"id":"77720",
						"device_name":"testapi_dns_minimal_1",
						"device_type":"host-nprobe-dns-www"
					},
					{ 
						"id":"77724",
						"device_name":"testapi_router_minimal_postman",
						"device_type":"router"
					},
					{
						"id":"77715",
						"device_name":"testapi_router_full_1",
						"device_type":"router"
					}
				],
				"edate":"2020-09-03T08:41:57.489Z",
				"fast_retention":30,
				"full_retention":30,
				"id":11466,
				"max_bigdata_fps":30,
				"max_devices":6,
				"max_fps":1000,
				"name":"Free Trial Plan",
				"metadata":{
				}
			}
		]
	}`
	transport := &connection.StubTransport{ResponseBody: getResponsePayload}
	plansAPI := resources.NewPlansAPI(transport)

	// act
	plans, err := plansAPI.GetAll(context.Background())

	// assert request properly formed
	assert := assert.New(t)
	require := require.New(t)

	require.NoError(err)
	assert.Zero(transport.RequestBody)
	assert.Equal("/plans", transport.RequestPath)

	// and response properly parsed
	require.Equal(1, len(plans))
	// plan 0
	plan := plans[0]
	assert.True(plan.Active)
	assert.Equal(models.ID(74333), plan.CompanyID)
	assert.Equal(2, len(plan.DeviceTypes))
	assert.Equal("router", plan.DeviceTypes[0].DeviceType)
	assert.Equal("host-nprobe-dns-www", plan.DeviceTypes[1].DeviceType)
	assert.Equal(4, len(plan.Devices))
	assert.Equal(models.ID(77714), plan.Devices[0].ID)
	assert.Equal("testapi_router_minimal_1", plan.Devices[0].DeviceName)
	assert.Equal("router", plan.Devices[0].DeviceType)
	assert.Equal(models.ID(77720), plan.Devices[1].ID)
	assert.Equal("testapi_dns_minimal_1", plan.Devices[1].DeviceName)
	assert.Equal("host-nprobe-dns-www", plan.Devices[1].DeviceType)
}
