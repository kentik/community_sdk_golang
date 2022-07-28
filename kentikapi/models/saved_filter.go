package models

import "time"

type SavedFilter struct {
	// Read-write properties

	FilterName        string
	FilterDescription string
	FilterLevel       string
	Filters           Filters

	// Read-only properties

	ID          ID
	CompanyID   ID
	CreatedDate time.Time
	UpdatedDate time.Time
}

// SavedFilterRequiredFields is a subset of SavedFilter fields required to create a SavedFilter.
type SavedFilterRequiredFields struct {
	FilterName        string
	FilterDescription string
	Filters           Filters
}

// NewSavedFilter creates a new SavedFilter with all required fields set.
func NewSavedFilter(f SavedFilterRequiredFields) *SavedFilter {
	return &SavedFilter{
		FilterName:        f.FilterName,
		FilterDescription: f.FilterDescription,
		Filters:           f.Filters,
	}
}

type Filters struct {
	// Read-write properties

	Connector    string
	FilterGroups []FilterGroups

	// Read-only properties (reserved for internal use)

	Custom       *bool
	FilterString *string
}

type FilterGroups struct {
	// Read-write properties

	Connector string
	Not       bool
	Filters   []Filter

	// Read-only properties

	FilterString *string
	ID           *ID
	Metric       *string
}

type Filter struct {
	// Read-write properties

	FilterField string
	FilterValue string
	Operator    string

	// Read-only properties

	ID *ID
}
