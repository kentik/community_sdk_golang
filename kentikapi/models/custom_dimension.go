package models

type CustomDimension struct {
	// read-write properties (can be updated in update call)
	DisplayName string

	// read-only properties (can't be updated in update call)
	Name       string // must start with c_ and be unique even against deleted dimensions (deleted names are retained for 1 year)
	Type       CustomDimensionType
	ID         ID
	CompanyID  ID
	Populators []Populator
}

// CustomDimensionRequiredFields is subset of CustomDimension fields required to create a CustomDimension.
// Note: name must begin with "c_" and be unique even among already deleted custom dimensions as names are retained for 1 year.
type CustomDimensionRequiredFields struct {
	Name        string
	DisplayName string
	Type        CustomDimensionType
}

// NewCustomDimension creates a CustomDimension with all necessary fields set.
func NewCustomDimension(d CustomDimensionRequiredFields) *CustomDimension {
	return &CustomDimension{
		Name:        d.Name,
		DisplayName: d.DisplayName,
		Type:        d.Type,
	}
}

type CustomDimensionType string

const (
	CustomDimensionTypeStr    CustomDimensionType = "string"
	CustomDimensionTypeUint32 CustomDimensionType = "uint32"
)
