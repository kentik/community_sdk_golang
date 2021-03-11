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
	SavedFilterPayload
}

type CreateSavedFilterRequest struct {
	SavedFilterPayload
}

type CreateSavedFilterResponse struct {
	SavedFilterPayload
}

type UpdateSavedFilterRequest struct {
	SavedFilterPayload
}

type UpdateSavedFilterResponse struct {
	SavedFilterPayload
}

type SavedFilterPayload struct {
	ID                IntAsString    `json:"id,omitempty"`
	CompanyID         models.ID      `json:"company_id,string,omitempty"`
	FilterName        string         `json:"filter_name"`
	FilterDescription string         `json:"filter_description"`
	FilterLevel       string         `json:"filter_level,omitempty"`
	CreatedDate       *time.Time     `json:"cdate,omitempty" response:"get,post,put"`
	UpdatedDate       *time.Time     `json:"edate,omitempty" response:"get,post,put"`
	Filters           FiltersPayload `json:"filters"`
}

func (p SavedFilterPayload) ToSavedFilter() (models.SavedFilter, error) {
	filters, err := p.Filters.ToFilters()
	if err != nil {
		return models.SavedFilter{}, err
	}

	return models.SavedFilter{
		ID:                int(p.ID),
		CompanyID:         p.CompanyID,
		FilterName:        p.FilterName,
		FilterDescription: p.FilterDescription,
		FilterLevel:       p.FilterLevel,
		CreatedDate:       *p.CreatedDate,
		UpdatedDate:       *p.UpdatedDate,
		Filters:           filters,
	}, nil
}

func SavedFilterToPayload(sf models.SavedFilter) SavedFilterPayload {
	return SavedFilterPayload{
		ID:                IntAsString(sf.ID),
		CompanyID:         sf.CompanyID,
		FilterName:        sf.FilterName,
		FilterDescription: sf.FilterDescription,
		FilterLevel:       sf.FilterLevel,
		CreatedDate:       &sf.CreatedDate,
		UpdatedDate:       &sf.UpdatedDate,
		Filters:           FiltersToPayload(sf.Filters),
	}
}

type FiltersPayload struct {
	Connector    string                `json:"connector"`
	Custom       *bool                 `json:"custom,omitempty"`
	FilterGroups []FilterGroupsPayload `json:"filterGroups" response:"get,post,put"`
	FilterString *string               `json:"filterString,omitempty"`
}

func (p FiltersPayload) ToFilters() (models.Filters, error) {
	var filterGroups []models.FilterGroups
	err := utils.ConvertList(p.FilterGroups, FilterGroupsPayload.ToFilterGroups, &filterGroups)
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

func FiltersToPayload(f models.Filters) FiltersPayload {
	var filterGroups []FilterGroupsPayload
	for _, fg := range f.FilterGroups {
		filterGroups = append(filterGroups, FilterGroupsToPayload(fg))
	}

	return FiltersPayload{
		Connector:    f.Connector,
		Custom:       f.Custom,
		FilterGroups: filterGroups,
		FilterString: f.FilterString,
	}
}

type FilterGroupsPayload struct {
	Connector    string          `json:"connector"`
	FilterString *string         `json:"filterString,omitempty"`
	ID           *models.ID      `json:"id,omitempty"`
	Metric       *string         `json:"metric,omitempty"`
	Not          bool            `json:"not"`
	Filters      []FilterPayload `json:"filters"`
}

func (p FilterGroupsPayload) ToFilterGroups() (models.FilterGroups, error) {
	var filters []models.Filter
	err := utils.ConvertList(p.Filters, FilterPayload.ToFilter, &filters)
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

func FilterGroupsToPayload(fg models.FilterGroups) FilterGroupsPayload {
	var filters []FilterPayload
	for _, f := range fg.Filters {
		filters = append(filters, FilterToPayload(f))
	}

	return FilterGroupsPayload{
		Connector:    fg.Connector,
		FilterString: fg.FilterString,
		ID:           fg.ID,
		Metric:       fg.Metric,
		Not:          fg.Not,
		Filters:      filters,
	}
}

type FilterPayload struct {
	FilterField string     `json:"filterField"`
	ID          *models.ID `json:"id,omitempty"`
	FilterValue string     `json:"filterValue"`
	Operator    string     `json:"operator"`
}

func (p FilterPayload) ToFilter() (models.Filter, error) {
	return models.Filter{
		FilterField: p.FilterField,
		ID:          p.ID,
		FilterValue: p.FilterValue,
		Operator:    p.Operator,
	}, nil
}

func FilterToPayload(f models.Filter) FilterPayload {
	return FilterPayload{
		FilterField: f.FilterField,
		ID:          f.ID,
		FilterValue: f.FilterValue,
		Operator:    f.Operator,
	}
}
