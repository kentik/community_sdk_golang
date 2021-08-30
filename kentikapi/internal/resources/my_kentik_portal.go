package resources

import (
	"context"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type MyKentikPortalAPI struct {
	BaseAPI
}

func NewMyKentikPortalAPI(transport connection.Transport) *MyKentikPortalAPI {
	return &MyKentikPortalAPI{BaseAPI{Transport: transport}}
}

// GetAll lists all tenants.
func (a *MyKentikPortalAPI) GetAll(ctx context.Context) ([]models.Tenant, error) {
	var response payloads.GetAllTenantsResponse
	if err := a.GetAndValidate(ctx, endpoints.TenantsPath, &response); err != nil {
		return []models.Tenant{}, err
	}

	return response.ToTenants()
}

// Get Tenant Info.
func (a *MyKentikPortalAPI) Get(ctx context.Context, tenantID models.ID) (*models.Tenant, error) {
	var response payloads.TenantPayload
	if err := a.GetAndValidate(ctx, endpoints.GetTenantPath(tenantID), &response); err != nil {
		return nil, err
	}
	tenant, err := response.ToTenant()
	return &tenant, err
}

func (a *MyKentikPortalAPI) CreateTenantUser(ctx context.Context,
	tenantID models.ID, userEmail string) (*models.TenantUser, error) {
	request := payloads.CreateTenantUserRequest{
		User: payloads.CreateTenantUserPayload{
			Email: userEmail,
		},
	}
	var response payloads.TenantUserPayload
	if err := a.PostAndValidate(ctx, endpoints.CreateTenantUserPath(tenantID), request, &response); err != nil {
		return nil, err
	}

	result, err := response.ToTenantUser()
	return &result, err
}

func (a *MyKentikPortalAPI) DeleteTenantUser(ctx context.Context, tenantID models.ID, userID models.ID) error {
	if err := a.DeleteAndValidate(ctx, endpoints.DeleteTenantUserPath(tenantID, userID), nil); err != nil {
		return err
	}

	return nil
}
