package resources

import (
	"context"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/utils"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type CustomApplicationsAPI struct {
	BaseAPI
}

// NewCustomApplicationsAPI is constructor.
func NewCustomApplicationsAPI(transport api_connection.Transport, logPayloads bool) *CustomApplicationsAPI {
	return &CustomApplicationsAPI{
		BaseAPI{Transport: transport, LogPayloads: logPayloads},
	}
}

// GetAll custom applications.
func (a *CustomApplicationsAPI) GetAll(ctx context.Context) ([]models.CustomApplication, error) {
	utils.LogPayload(a.LogPayloads, "GetAll custom applications Kentik API request", "")
	var response api_payloads.GetAllCustomApplicationsResponse
	if err := a.GetAndValidate(ctx, api_endpoints.GetAllCustomApplications(), &response); err != nil {
		return []models.CustomApplication{}, err
	}
	utils.LogPayload(a.LogPayloads, "GetAll custom applications Kentik API response", response)

	return response.ToCustomApplications()
}

// Create new custom application.
func (a *CustomApplicationsAPI) Create(ctx context.Context, customApplication models.CustomApplication,
) (*models.CustomApplication, error) {
	payload := api_payloads.CustomApplicationToPayload(customApplication)
	request := api_payloads.CreateCustomApplicationRequest(payload)
	utils.LogPayload(a.LogPayloads, "Create custom application Kentik API request", request)
	var response api_payloads.CreateCustomApplicationResponse
	if err := a.PostAndValidate(ctx, api_endpoints.CreateCustomApplication(), request, &response); err != nil {
		return nil, err
	}
	utils.LogPayload(a.LogPayloads, "Create custom application Kentik API response", response)

	result, err := response.ToCustomApplication()
	return &result, err
}

// Update custom application.
func (a *CustomApplicationsAPI) Update(ctx context.Context, customApplication models.CustomApplication,
) (*models.CustomApplication, error) {
	payload := api_payloads.CustomApplicationToPayload(customApplication)
	request := api_payloads.UpdateCustomApplicationRequest(payload)
	utils.LogPayload(a.LogPayloads, "Update custom application Kentik API request", request)
	var response api_payloads.UpdateCustomApplicationResponse
	if err := a.UpdateAndValidate(
		ctx,
		api_endpoints.UpdateCustomApplication(customApplication.ID),
		request,
		&response,
	); err != nil {
		return nil, err
	}
	utils.LogPayload(a.LogPayloads, "Update custom application Kentik API response", response)

	result, err := response.ToCustomApplication()
	return &result, err
}

// Delete custom application.
func (a *CustomApplicationsAPI) Delete(ctx context.Context, id models.ID) error {
	utils.LogPayload(a.LogPayloads, "Delete custom application Kentik API request ID", id)
	return a.DeleteAndValidate(ctx, api_endpoints.DeleteCustomApplication(id), nil)
}
