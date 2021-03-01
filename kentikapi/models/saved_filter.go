package models

import "time"

type SavedFilter struct {
	ID                ID
	CompanyID         ID
	FilterName        string
	FilterDescription string
	FilterLevel       string
	CreatedDate       time.Time
	UpdatedDate       time.Time
	Filters           Filters
}

type Filters struct {
	Connector    string
	Custom       *bool
	FilterGroups []FilterGroups
	FilterString *string
}

type FilterGroups struct {
	Connector    string
	FilterString *string
	ID           *ID
	Metric       *string
	Not          bool
	Filters      []Filter
}

type Filter struct {
	FilterField string
	ID          *ID
	FilterValue string
	Operator    string
}
