package models

import "time"

type SavedFilter struct {
	// read-write properties (can be updated in update call)
	FilterName        string
	FilterDescription string
	FilterLevel       string
	Filters           Filters

	// read-only properties (can't be updated in update call)
	ID          ID
	CompanyID   ID
	CreatedDate time.Time
	UpdatedDate time.Time
}

// NewSavedFilter creates a new SavedFilter with all required fields set.
func NewSavedFilter(name string, description string, level string, filters Filters) SavedFilter {
	return SavedFilter{
		FilterName:        name,
		FilterDescription: description,
		FilterLevel:       level,
		Filters:           filters,
	}
}

type Filters struct {
	// read-write properties
	Connector    string
	FilterGroups []FilterGroups

	// read-only properties (reserved for internal use)
	Custom       *bool
	FilterString *string
}

type FilterGroups struct {
	// read-write properties
	Connector string
	Not       bool
	Filters   []Filter

	// read-only properties
	FilterString *string
	ID           *ID
	Metric       *string
}

type Filter struct {
	// read-write properties
	FilterField string
	FilterValue string
	Operator    string

	// read-only properties
	ID *ID
}
