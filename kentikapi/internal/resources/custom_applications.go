package resources

import (
	"context"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type CustomApplicationsAPI struct {
	BaseAPI
}

// NewCustomApplicationsAPI is constructor.
func NewCustomApplicationsAPI(transport connection.Transport) *CustomApplicationsAPI {
	return &CustomApplicationsAPI{
		BaseAPI{Transport: transport},
	}
}

// GetAll custom applications.
func (a *CustomApplicationsAPI) GetAll(ctx context.Context) ([]models.CustomApplication, error) {
	var response payloads.GetAllCustomApplicationsResponse
	if err := a.GetAndValidate(ctx, endpoints.GetAllCustomApplications(), &response); err != nil {
		return []models.CustomApplication{}, err
	}

	return response.ToCustomApplications()
}

// Create new custom application.
func (a *CustomApplicationsAPI) Create(ctx context.Context,
	customApplication models.CustomApplication) (*models.CustomApplication, error) {
	payload := payloads.CustomApplicationToPayload(customApplication)
	request := payloads.CreateCustomApplicationRequest(payload)
	var response payloads.CreateCustomApplicationResponse
	if err := a.PostAndValidate(ctx, endpoints.CreateCustomApplication(), request, &response); err != nil {
		return nil, err
	}

	result, err := response.ToCustomApplication()
	return &result, err
}

// Update custom application.
func (a *CustomApplicationsAPI) Update(ctx context.Context,
	customApplication models.CustomApplication) (*models.CustomApplication, error) {
	payload := payloads.CustomApplicationToPayload(customApplication)
	request := payloads.UpdateCustomApplicationRequest(payload)
	var response payloads.UpdateCustomApplicationResponse
	if err := a.UpdateAndValidate(ctx, endpoints.UpdateCustomApplication(customApplication.ID), request, &response); err != nil {
		return nil, err
	}

	result, err := response.ToCustomApplication()
	return &result, err
}

// Delete custom application.
func (a *CustomApplicationsAPI) Delete(ctx context.Context, id models.ID) error {
	return a.DeleteAndValidate(ctx, endpoints.DeleteCustomApplication(id), nil)
}
