package resources

import (
	"context"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type CustomApplicationsAPI struct {
	BaseAPI
}

// NewCustomApplicationsAPI is constructor.
func NewCustomApplicationsAPI(transport api_connection.Transport) *CustomApplicationsAPI {
	return &CustomApplicationsAPI{
		BaseAPI{Transport: transport},
	}
}

// GetAll custom applications.
func (a *CustomApplicationsAPI) GetAll(ctx context.Context) ([]models.CustomApplication, error) {
	var response api_payloads.GetAllCustomApplicationsResponse
	if err := a.GetAndValidate(ctx, api_endpoints.GetAllCustomApplications(), &response); err != nil {
		return []models.CustomApplication{}, err
	}

	return response.ToCustomApplications()
}

// Create new custom application.
func (a *CustomApplicationsAPI) Create(ctx context.Context, customApplication models.CustomApplication,
) (*models.CustomApplication, error) {
	payload := api_payloads.CustomApplicationToPayload(customApplication)
	request := api_payloads.CreateCustomApplicationRequest(payload)
	var response api_payloads.CreateCustomApplicationResponse
	if err := a.PostAndValidate(ctx, api_endpoints.CreateCustomApplication(), request, &response); err != nil {
		return nil, err
	}

	result, err := response.ToCustomApplication()
	return &result, err
}

// Update custom application.
func (a *CustomApplicationsAPI) Update(ctx context.Context, customApplication models.CustomApplication,
) (*models.CustomApplication, error) {
	payload := api_payloads.CustomApplicationToPayload(customApplication)
	request := api_payloads.UpdateCustomApplicationRequest(payload)
	var response api_payloads.UpdateCustomApplicationResponse
	if err := a.UpdateAndValidate(
		ctx,
		api_endpoints.UpdateCustomApplication(customApplication.ID),
		request,
		&response,
	); err != nil {
		return nil, err
	}

	result, err := response.ToCustomApplication()
	return &result, err
}

// Delete custom application.
func (a *CustomApplicationsAPI) Delete(ctx context.Context, id models.ID) error {
	return a.DeleteAndValidate(ctx, api_endpoints.DeleteCustomApplication(id), nil)
}
