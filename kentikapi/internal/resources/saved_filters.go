package resources

import (
	"context"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/utils"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type SavedFiltersAPI struct {
	BaseAPI
}

func NewSavedFiltersAPI(transport api_connection.Transport, logPayloads bool) *SavedFiltersAPI {
	return &SavedFiltersAPI{BaseAPI{Transport: transport, LogPayloads: logPayloads}}
}

func (a *SavedFiltersAPI) GetAll(ctx context.Context) ([]models.SavedFilter, error) {
	utils.LogPayload(a.LogPayloads, "GetAll saved filters Kentik API request", "")
	var response api_payloads.GetAllSavedFilterResponse
	if err := a.GetAndValidate(ctx, api_endpoints.SavedFiltersPath, &response); err != nil {
		return nil, err
	}
	utils.LogPayload(a.LogPayloads, "GetAll saved filters Kentik API response", response)

	return response.ToSavedFilters()
}

func (a *SavedFiltersAPI) Get(ctx context.Context, filterID models.ID) (*models.SavedFilter, error) {
	utils.LogPayload(a.LogPayloads, "Get saved filter Kentik API request ID", filterID)
	var response api_payloads.GetSavedFilterResponse
	err := a.GetAndValidate(ctx, api_endpoints.GetSavedFilter(filterID), &response)
	if err != nil {
		return nil, err
	}
	utils.LogPayload(a.LogPayloads, "Get saved filter Kentik API response", response)

	savedFilter, err := response.ToSavedFilter()
	return &savedFilter, err
}

func (a *SavedFiltersAPI) Create(ctx context.Context, savedFilter models.SavedFilter) (*models.SavedFilter, error) {
	payload := api_payloads.SavedFilterToCreatePayload(savedFilter)
	utils.LogPayload(a.LogPayloads, "Create saved filter Kentik API request", payload)

	var response api_payloads.CreateSavedFilterResponse
	if err := a.PostAndValidate(ctx, api_endpoints.SavedFilterPath, payload, &response); err != nil {
		return nil, err
	}
	utils.LogPayload(a.LogPayloads, "Create saved filter Kentik API response", response)

	result, err := response.ToSavedFilter()
	return &result, err
}

func (a *SavedFiltersAPI) Update(ctx context.Context, savedFilter models.SavedFilter) (*models.SavedFilter, error) {
	payload := api_payloads.SavedFilterToUpdatePayload(savedFilter)
	utils.LogPayload(a.LogPayloads, "Update saved filter Kentik API request", payload)

	var response api_payloads.UpdateSavedFilterResponse
	if err := a.UpdateAndValidate(ctx, api_endpoints.GetSavedFilter(savedFilter.ID), payload, &response); err != nil {
		return nil, err
	}
	utils.LogPayload(a.LogPayloads, "Update saved filter Kentik API response", response)

	result, err := response.ToSavedFilter()
	return &result, err
}

func (a *SavedFiltersAPI) Detete(ctx context.Context, id models.ID) error {
	utils.LogPayload(a.LogPayloads, "Detete saved filter Kentik API request ID", id)
	return a.DeleteAndValidate(ctx, api_endpoints.GetSavedFilter(id), nil)
}
