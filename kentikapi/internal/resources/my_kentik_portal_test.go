package resources_test

import (
	"context"
	"testing"
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/resources"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/utils"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTenantsList(t *testing.T) {
	t.Parallel()

	// arrange
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

	companyID := 74333

	expected := []models.Tenant{
		{
			ID:          577,
			CompanyID:   &companyID,
			Name:        "test_tenant",
			Description: "This is test tenant",
			CreatedDate: time.Date(2020, 12, 21, 10, 55, 52, 449e6, time.UTC),
			UpdatedDate: time.Date(2020, 12, 21, 10, 55, 52, 449e6, time.UTC),
			Users: []models.TenantUser{
				{
					ID:        148099,
					CompanyID: 74333,
					Email:     "test@tenant.user",
					TenantID:  577,
				}, {
					ID:        148113,
					CompanyID: 74333,
					Email:     "user@testtenant.com",
					TenantID:  577,
				},
			},
		}, {
			ID:          578,
			Name:        "test_tenant2",
			Description: "",
			CreatedDate: time.Date(2020, 12, 21, 10, 57, 53, 425e6, time.UTC),
			UpdatedDate: time.Date(2020, 12, 21, 10, 57, 53, 425e6, time.UTC),
			Users:       []models.TenantUser{},
		},
	}

	transport := &api_connection.StubTransport{ResponseBody: getAllResponse}
	myKentikPortalAPI := resources.NewMyKentikPortalAPI(transport)

	// act
	tenants, err := myKentikPortalAPI.GetAll(context.Background())

	// assert

	// TODO(lwolanin): validate the request path passed to transport

	require.NoError(t, err)
	assert.Zero(t, transport.RequestBody)

	assert.Equal(t, expected, tenants)
}

func TestGetTenantInfo(t *testing.T) {
	t.Parallel()

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

	companyID := 74333

	expected := models.Tenant{
		ID:          577,
		CompanyID:   &companyID,
		Name:        "test_tenant",
		Description: "This is test tenant",
		CreatedDate: time.Date(2020, 12, 21, 10, 55, 52, 449e6, time.UTC),
		UpdatedDate: time.Date(2020, 12, 21, 10, 55, 52, 449e6, time.UTC),
		Users: []models.TenantUser{
			{
				ID:        148099,
				CompanyID: 74333,
				Email:     "test@tenant.user",
				TenantID:  577,
			}, {
				ID:        148113,
				CompanyID: 74333,
				Email:     "user@testtenant.com",
				TenantID:  577,
			},
		},
	}

	transport := &api_connection.StubTransport{ResponseBody: getTenantInfoResponse}
	myKentikPortalAPI := resources.NewMyKentikPortalAPI(transport)

	// act
	tenant, err := myKentikPortalAPI.Get(context.Background(), 577)

	// assert
	require.NoError(t, err)
	assert.Zero(t, transport.RequestBody)

	// TODO(lwolanin): validate the request path passed to transport

	assert.Equal(t, &expected, tenant)
}

func TestTenantUserCreate(t *testing.T) {
	t.Parallel()

	createTenantUserResponse := `
	{
		"id": "158564",
		"user_email": "test@test.test",
		"last_login": null,
		"tenant_id": "578",
		"company_id": "74333"
	}`

	expected := models.TenantUser{
		ID:        158564,
		Email:     "test@test.test",
		TenantID:  578,
		CompanyID: 74333,
	}

	transport := &api_connection.StubTransport{ResponseBody: createTenantUserResponse}
	myKentikPortalAPI := resources.NewMyKentikPortalAPI(transport)

	// act
	tenantUser, err := myKentikPortalAPI.CreateTenantUser(context.Background(), 577, "test@test.test")

	// assert
	// TODO(lwolanin): Validate the request path passed to transport
	// TODO(lwolanin): Verify that that there is no redundant data sent in request body

	require.NoError(t, err)
	payload := utils.NewJSONPayloadInspector(t, transport.RequestBody)
	assert.NotNil(t, payload.Get("user"))
	assert.Equal(t, "test@test.test", payload.String("user/user_email"))

	assert.Equal(t, &expected, tenantUser)
}

func TestTenantUserDelete(t *testing.T) {
	t.Parallel()

	// arrange
	deleteResponsePayload := ""
	transport := &api_connection.StubTransport{ResponseBody: deleteResponsePayload}
	myKentikPortalAPI := resources.NewMyKentikPortalAPI(transport)

	// act
	tenantID := models.ID(478)
	userID := models.ID(420)
	err := myKentikPortalAPI.DeleteTenantUser(context.Background(), tenantID, userID)

	// assert
	require.NoError(t, err)
	assert.Zero(t, transport.RequestBody)

	// TODO(lwolanin): validate the request path passed to transport
}
