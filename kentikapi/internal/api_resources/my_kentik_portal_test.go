package api_resources_test

import (
	"testing"
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_resources"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/utils"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTenantsList(t *testing.T) {
	//arrange
	getAllResponse := `
	[
		{
			"id": 577,
			"name": "test_tenant",
			"description": "This is test tenant",
			"users": [
				{
					"id": "148099",
					"user_email": "test@tenant.user",
					"last_login": null,
					"tenant_id": "577",
					"company_id": "74333"
				},
				{
					"id": "148113",
					"user_email": "user@testtenant.com",
					"last_login": null,
					"tenant_id": "577",
					"company_id": "74333"
				}
			],
			"created_date": "2020-12-21T10:55:52.449Z",
			"updated_date": "2020-12-21T10:55:52.449Z"
		},
		{
			"id": 578,
			"name": "test_tenant2",
			"description": "",
			"users": [],
			"created_date": "2020-12-21T10:57:53.425Z",
			"updated_date": "2020-12-21T10:57:53.425Z"
		}
	]`

	transport := &api_connection.StubTransport{ResponseBody: getAllResponse}
	myKentikPortalAPI := api_resources.NewMyKentikPortalAPI(transport)

	//act
	tenants, err := myKentikPortalAPI.GetAll(nil)

	//assert
	assert := assert.New(t)
	require := require.New(t)

	require.NoError(err)
	assert.Zero(transport.RequestBody)

	assert.Len(tenants, 2)
	assert.Equal(models.ID(577), tenants[0].ID)
	assert.Equal("test_tenant", tenants[0].Name)
	assert.Equal("This is test tenant", tenants[0].Description)
	assert.Len(tenants[0].Users, 2)
	assert.Equal(2020, tenants[0].CreatedDate.Year())
	assert.Equal(time.Month(12), tenants[0].CreatedDate.Month())
	assert.Equal(21, tenants[0].CreatedDate.Day())
	assert.Equal(10, tenants[0].CreatedDate.Hour())
	assert.Equal(55, tenants[0].CreatedDate.Minute())
	assert.Equal(52, tenants[0].CreatedDate.Second())
	assert.Equal(2020, tenants[0].UpdatedDate.Year())
	assert.Equal(time.Month(12), tenants[0].UpdatedDate.Month())
	assert.Equal(21, tenants[0].UpdatedDate.Day())
	assert.Equal(10, tenants[0].UpdatedDate.Hour())
	assert.Equal(55, tenants[0].UpdatedDate.Minute())
	assert.Equal(52, tenants[0].UpdatedDate.Second())
	assert.Equal(models.ID(148099), tenants[0].Users[0].ID)
	assert.Equal("test@tenant.user", tenants[0].Users[0].Email)
	assert.Nil(tenants[0].Users[0].LastLogin)
	assert.Equal(models.ID(577), tenants[0].Users[0].TenantID)
	assert.Equal(models.ID(74333), tenants[0].Users[0].CompanyID)
	assert.Equal(models.ID(148113), tenants[0].Users[1].ID)
	assert.Equal("user@testtenant.com", tenants[0].Users[1].Email)
	assert.Nil(tenants[0].Users[1].LastLogin)
	assert.Equal(models.ID(577), tenants[0].Users[1].TenantID)
	assert.Equal(models.ID(74333), tenants[0].Users[1].CompanyID)

	assert.Equal(models.ID(578), tenants[1].ID)
	assert.Equal("test_tenant2", tenants[1].Name)
	assert.Equal("", tenants[1].Description)
	assert.Empty(tenants[1].Users)
	assert.Equal(2020, tenants[1].CreatedDate.Year())
	assert.Equal(time.Month(12), tenants[1].CreatedDate.Month())
	assert.Equal(21, tenants[1].CreatedDate.Day())
	assert.Equal(10, tenants[1].CreatedDate.Hour())
	assert.Equal(57, tenants[1].CreatedDate.Minute())
	assert.Equal(53, tenants[1].CreatedDate.Second())
	assert.Equal(2020, tenants[1].UpdatedDate.Year())
	assert.Equal(time.Month(12), tenants[1].UpdatedDate.Month())
	assert.Equal(21, tenants[1].UpdatedDate.Day())
	assert.Equal(10, tenants[1].UpdatedDate.Hour())
	assert.Equal(57, tenants[1].UpdatedDate.Minute())
	assert.Equal(53, tenants[1].UpdatedDate.Second())

}

func TestGetTenantInfo(t *testing.T) {
	getTenantInfoResponse := `
	{
		"id": 577,
		"name": "test_tenant",
		"description": "This is test tenant",
		"users": [
			{
				"id": "148099",
				"user_email": "test@tenant.user",
				"last_login": null,
				"tenant_id": "577",
				"company_id": "74333"
			},
			{
				"id": "148113",
				"user_email": "user@testtenant.com",
				"last_login": null,
				"tenant_id": "577",
				"company_id": "74333"
			}
		],
		"created_date": "2020-12-21T10:55:52.449Z",
		"updated_date": "2020-12-21T10:55:52.449Z"
	}`

	transport := &api_connection.StubTransport{ResponseBody: getTenantInfoResponse}
	myKentikPortalAPI := api_resources.NewMyKentikPortalAPI(transport)

	//act
	tenant, err := myKentikPortalAPI.Get(nil, 577)

	//assert
	assert := assert.New(t)
	require := require.New(t)

	require.NoError(err)
	assert.Zero(transport.RequestBody)

	assert.Equal(models.ID(577), tenant.ID)
	assert.Equal("test_tenant", tenant.Name)
	assert.Equal("This is test tenant", tenant.Description)
	assert.Len(tenant.Users, 2)
	assert.Equal(2020, tenant.CreatedDate.Year())
	assert.Equal(time.Month(12), tenant.CreatedDate.Month())
	assert.Equal(21, tenant.CreatedDate.Day())
	assert.Equal(10, tenant.CreatedDate.Hour())
	assert.Equal(55, tenant.CreatedDate.Minute())
	assert.Equal(52, tenant.CreatedDate.Second())
	assert.Equal(2020, tenant.UpdatedDate.Year())
	assert.Equal(time.Month(12), tenant.UpdatedDate.Month())
	assert.Equal(21, tenant.UpdatedDate.Day())
	assert.Equal(10, tenant.UpdatedDate.Hour())
	assert.Equal(55, tenant.UpdatedDate.Minute())
	assert.Equal(52, tenant.UpdatedDate.Second())
	assert.Equal(models.ID(148099), tenant.Users[0].ID)
	assert.Equal("test@tenant.user", tenant.Users[0].Email)
	assert.Nil(tenant.Users[0].LastLogin)
	assert.Equal(models.ID(577), tenant.Users[0].TenantID)
	assert.Equal(models.ID(74333), tenant.Users[0].CompanyID)
	assert.Equal(models.ID(148113), tenant.Users[1].ID)
	assert.Equal("user@testtenant.com", tenant.Users[1].Email)
	assert.Nil(tenant.Users[1].LastLogin)
	assert.Equal(models.ID(577), tenant.Users[1].TenantID)
	assert.Equal(models.ID(74333), tenant.Users[1].CompanyID)
}

func TestTenantUserCreate(t *testing.T) {
	createTenantUserResponse := `
	{
		"id": "158564",
		"user_email": "test@test.test",
		"last_login": null,
		"tenant_id": "578",
		"company_id": "74333"
	}`

	transport := &api_connection.StubTransport{ResponseBody: createTenantUserResponse}
	myKentikPortalAPI := api_resources.NewMyKentikPortalAPI(transport)

	//act
	tenantUser, err := myKentikPortalAPI.CreateTenantUser(nil, 577, "test@test.test")

	//assert
	assert := assert.New(t)
	require := require.New(t)

	require.NoError(err)
	payload := utils.NewJSONPayloadInspector(t, transport.RequestBody)
	assert.NotNil(payload.Get("user"))
	assert.Equal("test@test.test", payload.String("user/user_email"))

	assert.Equal(models.ID(158564), tenantUser.ID)
	assert.Equal("test@test.test", tenantUser.Email)
	assert.Nil(tenantUser.LastLogin)
	assert.Equal(models.ID(578), tenantUser.TenantID)
	assert.Equal(models.ID(74333), tenantUser.CompanyID)

}

func TestTenantUserDelete(t *testing.T) {
	// arrange
	deleteResponsePayload := ""
	transport := &api_connection.StubTransport{ResponseBody: deleteResponsePayload}
	myKentikPortalAPI := api_resources.NewMyKentikPortalAPI(transport)

	// act
	tenantID := models.ID(478)
	userID := models.ID(420)
	err := myKentikPortalAPI.DeleteTenantUser(nil, tenantID, userID)

	// assert
	assert := assert.New(t)
	require := require.New(t)
	require.NoError(err)
	assert.Zero(transport.RequestBody)
}
