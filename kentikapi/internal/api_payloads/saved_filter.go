package api_payloads

import (
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/utils"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type GetAllSavedFilterResponse []GetSavedFilterResponse

func (p GetAllSavedFilterResponse) ToSavedFilters() ([]models.SavedFilter, error) {
	var result []models.SavedFilter
	err := utils.ConvertList(p, GetSavedFilterResponse.ToSavedFilter, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type GetSavedFilterResponse struct {
	savedFilterPayload
}

type CreateSavedFilterRequest struct {
	savedFilterPayload
}

type CreateSavedFilterResponse struct {
	savedFilterPayload
}

type UpdateSavedFilterRequest struct {
	savedFilterPayload
}

type UpdateSavedFilterResponse struct {
	savedFilterPayload
}

type savedFilterPayload struct {
	ID                StringAsInt    `json:"id,omitempty"`
	CompanyID         models.ID      `json:"company_id,omitempty"`
	FilterName        string         `json:"filter_name"`
	FilterDescription string         `json:"filter_description"`
	FilterLevel       string         `json:"filter_level,omitempty"`
	CreatedDate       *time.Time     `json:"cdate,omitempty" response:"get,post,put"`
	UpdatedDate       *time.Time     `json:"edate,omitempty" response:"get,post,put"`
	Filters           filtersPayload `json:"filters"`
}

func (p savedFilterPayload) ToSavedFilter() (models.SavedFilter, error) {
	filters, err := p.Filters.ToFilters()
	if err != nil {
		return models.SavedFilter{}, err
	}

	return models.SavedFilter{
		ID:                string(p.ID),
		CompanyID:         p.CompanyID,
		FilterName:        p.FilterName,
		FilterDescription: p.FilterDescription,
		FilterLevel:       p.FilterLevel,
		CreatedDate:       *p.CreatedDate,
		UpdatedDate:       *p.UpdatedDate,
		Filters:           filters,
	}, nil
}

//nolint:revive // savedFilterPayLoad doesn't need to be exported
func SavedFilterToPayload(sf models.SavedFilter) savedFilterPayload {
	return savedFilterPayload{
		ID:                StringAsInt(sf.ID),
		CompanyID:         sf.CompanyID,
		FilterName:        sf.FilterName,
		FilterDescription: sf.FilterDescription,
		FilterLevel:       sf.FilterLevel,
		CreatedDate:       &sf.CreatedDate,
		UpdatedDate:       &sf.UpdatedDate,
		Filters:           filtersToPayload(sf.Filters),
	}
}

func SavedFilterToCreatePayload(sf models.SavedFilter) CreateSavedFilterRequest {
	return CreateSavedFilterRequest{
		savedFilterPayload: SavedFilterToPayload(sf),
	}
}

func SavedFilterToUpdatePayload(sf models.SavedFilter) UpdateSavedFilterRequest {
	return UpdateSavedFilterRequest{
		savedFilterPayload: SavedFilterToPayload(sf),
	}
}

type filtersPayload struct {
	Connector    string                `json:"connector"`
	Custom       *bool                 `json:"custom,omitempty"`
	FilterGroups []filterGroupsPayload `json:"filterGroups"`
	FilterString *string               `json:"filterString,omitempty"`
}

func (p filtersPayload) ToFilters() (models.Filters, error) {
	var filterGroups []models.FilterGroups
	err := utils.ConvertList(p.FilterGroups, filterGroupsPayload.ToFilterGroups, &filterGroups)
	if err != nil {
		return models.Filters{}, err
	}
	return models.Filters{
		Connector:    p.Connector,
		Custom:       p.Custom,
		FilterGroups: filterGroups,
		FilterString: p.FilterString,
	}, nil
}

func filtersToPayload(f models.Filters) filtersPayload {
	var filterGroups []filterGroupsPayload
	for _, fg := range f.FilterGroups {
		filterGroups = append(filterGroups, filterGroupsToPayload(fg))
	}

	return filtersPayload{
		Connector:    f.Connector,
		Custom:       f.Custom,
		FilterGroups: filterGroups,
		FilterString: f.FilterString,
	}
}

type filterGroupsPayload struct {
	Connector    string          `json:"connector"`
	FilterString *string         `json:"filterString,omitempty"`
	ID           *models.ID      `json:"id,omitempty"`
	Metric       *string         `json:"metric,omitempty"`
	Not          bool            `json:"not"`
	Filters      []filterPayload `json:"filters"`
}

func (p filterGroupsPayload) ToFilterGroups() (models.FilterGroups, error) {
	var filters []models.Filter
	err := utils.ConvertList(p.Filters, filterPayload.ToFilter, &filters)
	if err != nil {
		return models.FilterGroups{}, err
	}
	return models.FilterGroups{
		Connector:    p.Connector,
		FilterString: p.FilterString,
		ID:           p.ID,
		Metric:       p.Metric,
		Not:          p.Not,
		Filters:      filters,
	}, nil
}

func filterGroupsToPayload(fg models.FilterGroups) filterGroupsPayload {
	var filters []filterPayload
	for _, f := range fg.Filters {
		filters = append(filters, filterToPayload(f))
	}

	return filterGroupsPayload{
		Connector:    fg.Connector,
		FilterString: fg.FilterString,
		ID:           fg.ID,
		Metric:       fg.Metric,
		Not:          fg.Not,
		Filters:      filters,
	}
}

type filterPayload struct {
	FilterField string     `json:"filterField"`
	ID          *models.ID `json:"id,omitempty"`
	FilterValue string     `json:"filterValue"`
	Operator    string     `json:"operator"`
}

func (p filterPayload) ToFilter() (models.Filter, error) {
	return models.Filter{
		FilterField: p.FilterField,
		ID:          p.ID,
		FilterValue: p.FilterValue,
		Operator:    p.Operator,
	}, nil
}

func filterToPayload(f models.Filter) filterPayload {
	return filterPayload{
		FilterField: f.FilterField,
		ID:          f.ID,
		FilterValue: f.FilterValue,
		Operator:    f.Operator,
	}
}
