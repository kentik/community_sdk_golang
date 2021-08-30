package resources

import (
	"context"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type CustomDimensionsAPI struct {
	BaseAPI
	Populators *populatorsAPI
}

// NewCustomDimensionsAPI is constructor.
func NewCustomDimensionsAPI(transport connection.Transport) *CustomDimensionsAPI {
	return &CustomDimensionsAPI{
		BaseAPI{Transport: transport},
		&populatorsAPI{BaseAPI{Transport: transport}},
	}
}

// GetAll custom dimensions.
func (a *CustomDimensionsAPI) GetAll(ctx context.Context) ([]models.CustomDimension, error) {
	var response payloads.GetAllCustomDimensionsResponse
	if err := a.GetAndValidate(ctx, endpoints.GetAllCustomDimensions(), &response); err != nil {
		return []models.CustomDimension{}, err
	}

	return response.ToCustomDimensions(), nil
}

// Get custom dimension with given ID.
func (a *CustomDimensionsAPI) Get(ctx context.Context, id models.ID) (*models.CustomDimension, error) {
	var response payloads.GetCustomDimensionResponse
	if err := a.GetAndValidate(ctx, endpoints.GetCustomDimension(id), &response); err != nil {
		return nil, err
	}

	result := response.ToCustomDimension()
	return &result, nil
}

// Create new custom dimension.
func (a *CustomDimensionsAPI) Create(ctx context.Context,
	customDimension models.CustomDimension) (*models.CustomDimension, error) {
	payload := payloads.CustomDimensionToPayload(customDimension)

	request := payloads.CreateCustomDimensionRequest(payload)
	var response payloads.CreateCustomDimensionResponse
	if err := a.PostAndValidate(ctx, endpoints.CreateCustomDimension(), request, &response); err != nil {
		return nil, err
	}

	result := response.ToCustomDimension()
	return &result, nil
}

// Update custom dimension.
func (a *CustomDimensionsAPI) Update(ctx context.Context,
	customDimension models.CustomDimension) (*models.CustomDimension, error) {
	payload := payloads.CustomDimensionToPayload(customDimension)

	request := payloads.UpdateCustomDimensionRequest(payload)
	var response payloads.UpdateCustomDimensionResponse
	if err := a.UpdateAndValidate(ctx, endpoints.UpdateCustomDimension(customDimension.ID), request, &response); err != nil {
		return nil, err
	}

	result := response.ToCustomDimension()
	return &result, nil
}

// Delete custom dimension.
func (a *CustomDimensionsAPI) Delete(ctx context.Context, id models.ID) error {
	return a.DeleteAndValidate(ctx, endpoints.DeleteCustomDimension(id), nil)
}

type populatorsAPI struct {
	BaseAPI
}

// Create new populator.
func (a *populatorsAPI) Create(ctx context.Context, populator models.Populator) (*models.Populator, error) {
	payload := payloads.PopulatorToPayload(populator)

	request := payloads.CreatePopulatorRequest{Payload: payload}
	var response payloads.CreatePopulatorResponse
	if err := a.PostAndValidate(ctx, endpoints.CreatePopulator(populator.DimensionID), request, &response); err != nil {
		return nil, err
	}

	result := response.ToPopulator()
	return &result, nil
}

// Update populator.
func (a *populatorsAPI) Update(ctx context.Context, populator models.Populator) (*models.Populator, error) {
	payload := payloads.PopulatorToPayload(populator)

	request := payloads.UpdatePopulatorRequest{Payload: payload}
	var response payloads.UpdatePopulatorResponse
	if err := a.UpdateAndValidate(ctx,
		endpoints.UpdatePopulator(populator.DimensionID, populator.ID),
		request,
		&response); err != nil {
		return nil, err
	}

	result := response.ToPopulator()
	return &result, nil
}

// Delete populator.
func (a *populatorsAPI) Delete(ctx context.Context, dimensionID, populatorID models.ID) error {
	return a.DeleteAndValidate(ctx, endpoints.DeletePopulator(dimensionID, populatorID), nil)
}
