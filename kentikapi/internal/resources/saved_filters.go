package resources

import (
	"context"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type SavedFiltersAPI struct {
	BaseAPI
}

func NewSavedFiltersAPI(transport connection.Transport) *SavedFiltersAPI {
	return &SavedFiltersAPI{BaseAPI{Transport: transport}}
}

func (a *SavedFiltersAPI) GetAll(ctx context.Context) ([]models.SavedFilter, error) {
	var response payloads.GetAllSavedFilterResponse
	if err := a.GetAndValidate(ctx, endpoints.SavedFiltersPath, &response); err != nil {
		return nil, err
	}

	return response.ToSavedFilters()
}

func (a *SavedFiltersAPI) Get(ctx context.Context, filterID models.ID) (*models.SavedFilter, error) {
	var response payloads.GetSavedFilterResponse
	err := a.GetAndValidate(ctx, endpoints.GetSavedFilter(filterID), &response)
	if err != nil {
		return nil, err
	}

	savedFilter, err := response.ToSavedFilter()

	return &savedFilter, err
}

func (a *SavedFiltersAPI) Create(ctx context.Context, savedFilter models.SavedFilter) (*models.SavedFilter, error) {
	payload := payloads.SavedFilterToCreatePayload(savedFilter)

	var response payloads.CreateSavedFilterResponse
	if err := a.PostAndValidate(ctx, endpoints.SavedFilterPath, payload, &response); err != nil {
		return nil, err
	}

	result, err := response.ToSavedFilter()
	return &result, err
}

func (a *SavedFiltersAPI) Update(ctx context.Context, savedFilter models.SavedFilter) (*models.SavedFilter, error) {
	payload := payloads.SavedFilterToUpdatePayload(savedFilter)

	var response payloads.UpdateSavedFilterResponse
	if err := a.UpdateAndValidate(ctx, endpoints.GetSavedFilter(savedFilter.ID), payload, &response); err != nil {
		return nil, err
	}

	result, err := response.ToSavedFilter()
	return &result, err
}

func (a *SavedFiltersAPI) Detete(ctx context.Context, id models.ID) error {
	return a.DeleteAndValidate(ctx, endpoints.GetSavedFilter(id), nil)
}
