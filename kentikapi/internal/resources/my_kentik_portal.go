package resources

import (
	"context"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_payloads"
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
	var response api_payloads.GetAllTenantsResponse
	if err := a.GetAndValidate(ctx, api_endpoints.TenantsPath, &response); err != nil {
		return []models.Tenant{}, err
	}

	return response.ToTenants()
}

// Get Tenant Info.
func (a *MyKentikPortalAPI) Get(ctx context.Context, tenantID models.ID) (*models.Tenant, error) {
	var response api_payloads.TenantPayload
	if err := a.GetAndValidate(ctx, api_endpoints.GetTenantPath(tenantID), &response); err != nil {
		return nil, err
	}
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
	var response api_payloads.TenantUserPayload
	if err := a.PostAndValidate(ctx, api_endpoints.CreateTenantUserPath(tenantID), request, &response); err != nil {
		return nil, err
	}

	result, err := response.ToTenantUser()
	return &result, err
}

func (a *MyKentikPortalAPI) DeleteTenantUser(ctx context.Context, tenantID models.ID, userID models.ID) error {
	if err := a.DeleteAndValidate(ctx, api_endpoints.DeleteTenantUserPath(tenantID, userID), nil); err != nil {
		return err
	}

	return nil
}
