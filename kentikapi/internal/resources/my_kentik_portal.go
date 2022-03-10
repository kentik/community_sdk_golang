package resources

import (
	"context"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/utils"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type MyKentikPortalAPI struct {
	BaseAPI
}

func NewMyKentikPortalAPI(transport api_connection.Transport, logPayloads bool) *MyKentikPortalAPI {
	return &MyKentikPortalAPI{BaseAPI{Transport: transport, LogPayloads: logPayloads}}
}

// GetAll lists all tenants.
func (a *MyKentikPortalAPI) GetAll(ctx context.Context) ([]models.Tenant, error) {
	utils.LogPayload(a.LogPayloads, "GetAll tenats Kentik API request", "")
	var response api_payloads.GetAllTenantsResponse
	if err := a.GetAndValidate(ctx, api_endpoints.TenantsPath, &response); err != nil {
		return []models.Tenant{}, err
	}
	utils.LogPayload(a.LogPayloads, "GetAll tenats Kentik API response", response)

	return response.ToTenants()
}

// Get Tenant Info.
func (a *MyKentikPortalAPI) Get(ctx context.Context, tenantID models.ID) (*models.Tenant, error) {
	utils.LogPayload(a.LogPayloads, "Get tenant Kentik API request ID", tenantID)
	var response api_payloads.TenantPayload
	if err := a.GetAndValidate(ctx, api_endpoints.GetTenantPath(tenantID), &response); err != nil {
		return nil, err
	}
	utils.LogPayload(a.LogPayloads, "Get tenant Kentik API response", response)
	tenant, err := response.ToTenant()
	return &tenant, err
}

func (a *MyKentikPortalAPI) CreateTenantUser(ctx context.Context, tenantID models.ID, userEmail string,
) (*models.TenantUser, error) {
	request := api_payloads.CreateTenantUserRequest{
		User: api_payloads.CreateTenantUserPayload{
			Email: userEmail,
		},
	}
	utils.LogPayload(a.LogPayloads, "CreateTenantUser Kentik API request", request)
	var response api_payloads.TenantUserPayload
	if err := a.PostAndValidate(ctx, api_endpoints.CreateTenantUserPath(tenantID), request, &response); err != nil {
		return nil, err
	}
	utils.LogPayload(a.LogPayloads, "CreateTenantUser Kentik API response", response)

	result, err := response.ToTenantUser()
	return &result, err
}

func (a *MyKentikPortalAPI) DeleteTenantUser(ctx context.Context, tenantID models.ID, userID models.ID) error {
	utils.LogPayload(a.LogPayloads, "DeleteTenantUser Kentik API request ID", tenantID)
	if err := a.DeleteAndValidate(ctx, api_endpoints.DeleteTenantUserPath(tenantID, userID), nil); err != nil {
		return err
	}

	return nil
}
