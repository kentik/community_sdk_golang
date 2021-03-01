package api_resources

import (
	"context"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type CustomDimensionsAPI struct {
	BaseAPI
	Populators *populatorsAPI
}

// NewCustomDimensionsAPI is constructor
func NewCustomDimensionsAPI(transport api_connection.Transport) *CustomDimensionsAPI {
	return &CustomDimensionsAPI{
		BaseAPI{Transport: transport},
		&populatorsAPI{BaseAPI{Transport: transport}},
	}
}

// GetAll custom dimensions
func (a *CustomDimensionsAPI) GetAll(ctx context.Context) ([]models.CustomDimension, error) {
	var response api_payloads.GetAllCustomDimensionsResponse
	if err := a.GetAndValidate(ctx, api_endpoints.GetAllCustomDimensions(), &response); err != nil {
		return []models.CustomDimension{}, err
	}

	return response.ToCustomDimensions()
}

// Get custom dimension with given ID
func (a *CustomDimensionsAPI) Get(ctx context.Context, id models.ID) (*models.CustomDimension, error) {
	var response api_payloads.GetCustomDimensionResponse
	if err := a.GetAndValidate(ctx, api_endpoints.GetCustomDimension(id), &response); err != nil {
		return nil, err
	}

	customDimension, err := response.ToCustomDimension()
	return &customDimension, err
}

// Create new custom dimension
func (a *CustomDimensionsAPI) Create(ctx context.Context, customDimension models.CustomDimension) (*models.CustomDimension, error) {
	payload := api_payloads.CustomDimensionToPayload(customDimension)

	request := api_payloads.CreateCustomDimensionRequest(payload)
	var response api_payloads.CreateCustomDimensionResponse
	if err := a.PostAndValidate(ctx, api_endpoints.CreateCustomDimension(), request, &response); err != nil {
		return nil, err
	}

	result, err := response.ToCustomDimension()
	return &result, err
}

// Update custom dimension
func (a *CustomDimensionsAPI) Update(ctx context.Context, customDimension models.CustomDimension) (*models.CustomDimension, error) {
	payload := api_payloads.CustomDimensionToPayload(customDimension)

	request := api_payloads.UpdateCustomDimensionRequest(payload)
	var response api_payloads.UpdateCustomDimensionResponse
	if err := a.UpdateAndValidate(ctx, api_endpoints.UpdateCustomDimension(customDimension.ID), request, &response); err != nil {
		return nil, err
	}

	result, err := response.ToCustomDimension()
	return &result, err
}

// Delete custom dimension
func (a *CustomDimensionsAPI) Delete(ctx context.Context, id models.ID) error {
	if err := a.DeleteAndValidate(ctx, api_endpoints.GetCustomDimension(id), nil); err != nil {
		return err
	}

	return nil
}

type populatorsAPI struct {
	BaseAPI
}
