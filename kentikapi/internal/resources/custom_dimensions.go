package resources

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

// NewCustomDimensionsAPI creates new CustomDimensionsAPI.
func NewCustomDimensionsAPI(transport api_connection.Transport, logPayloads bool) *CustomDimensionsAPI {
	return &CustomDimensionsAPI{
		BaseAPI{Transport: transport, LogPayloads: logPayloads},
		&populatorsAPI{BaseAPI{Transport: transport}},
	}
}

// GetAll custom dimensions.
func (a *CustomDimensionsAPI) GetAll(ctx context.Context) ([]models.CustomDimension, error) {
	var response api_payloads.GetAllCustomDimensionsResponse
	if err := a.GetAndValidate(ctx, api_endpoints.GetAllCustomDimensions(), &response); err != nil {
		return []models.CustomDimension{}, err
	}

	return response.ToCustomDimensions(), nil
}

// Get custom dimension with given ID.
func (a *CustomDimensionsAPI) Get(ctx context.Context, id models.ID) (*models.CustomDimension, error) {
	var response api_payloads.GetCustomDimensionResponse
	if err := a.GetAndValidate(ctx, api_endpoints.GetCustomDimension(id), &response); err != nil {
		return nil, err
	}

	result := response.ToCustomDimension()
	return &result, nil
}

// Create new custom dimension.
func (a *CustomDimensionsAPI) Create(ctx context.Context, cd models.CustomDimension) (*models.CustomDimension, error) {
	payload := api_payloads.CustomDimensionToPayload(cd)

	request := api_payloads.CreateCustomDimensionRequest(payload)
	var response api_payloads.CreateCustomDimensionResponse
	if err := a.PostAndValidate(ctx, api_endpoints.CreateCustomDimension(), request, &response); err != nil {
		return nil, err
	}

	result := response.ToCustomDimension()
	return &result, nil
}

// Update custom dimension.
func (a *CustomDimensionsAPI) Update(ctx context.Context, cd models.CustomDimension) (*models.CustomDimension, error) {
	payload := api_payloads.CustomDimensionToPayload(cd)

	request := api_payloads.UpdateCustomDimensionRequest(payload)
	var response api_payloads.UpdateCustomDimensionResponse
	if err := a.UpdateAndValidate(ctx, api_endpoints.UpdateCustomDimension(cd.ID), request, &response); err != nil {
		return nil, err
	}

	result := response.ToCustomDimension()
	return &result, nil
}

// Delete custom dimension.
func (a *CustomDimensionsAPI) Delete(ctx context.Context, id models.ID) error {
	return a.DeleteAndValidate(ctx, api_endpoints.DeleteCustomDimension(id), nil)
}

type populatorsAPI struct {
	BaseAPI
}

// Create new populator.
func (a *populatorsAPI) Create(ctx context.Context, populator models.Populator) (*models.Populator, error) {
	payload := api_payloads.PopulatorToPayload(populator)

	request := api_payloads.CreatePopulatorRequest{Payload: payload}
	var response api_payloads.CreatePopulatorResponse
	if err := a.PostAndValidate(ctx, api_endpoints.CreatePopulator(populator.DimensionID), request, &response); err != nil {
		return nil, err
	}

	result := response.ToPopulator()
	return &result, nil
}

// Update populator.
func (a *populatorsAPI) Update(ctx context.Context, populator models.Populator) (*models.Populator, error) {
	payload := api_payloads.PopulatorToPayload(populator)

	request := api_payloads.UpdatePopulatorRequest{Payload: payload}
	var response api_payloads.UpdatePopulatorResponse
	if err := a.UpdateAndValidate(
		ctx,
		api_endpoints.UpdatePopulator(populator.DimensionID, populator.ID),
		request,
		&response,
	); err != nil {
		return nil, err
	}

	result := response.ToPopulator()
	return &result, nil
}

// Delete populator.
func (a *populatorsAPI) Delete(ctx context.Context, dimensionID, populatorID models.ID) error {
	return a.DeleteAndValidate(ctx, api_endpoints.DeletePopulator(dimensionID, populatorID), nil)
}
