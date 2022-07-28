package models

type CustomDimension struct {
	// Read-write properties

	DisplayName string

	// Read-only properties

	// Name must start with c_ and be unique even against deleted dimensions (deleted names are retained for 1 year).
	Name       string
	Type       CustomDimensionType
	ID         ID
	CompanyID  ID
	Populators []Populator
}

// CustomDimensionRequiredFields is a subset of CustomDimension fields required to create a CustomDimension.
type CustomDimensionRequiredFields struct {
	// Name must begin with "c_" and be unique even among already deleted custom dimensions as names are retained for 1 year.
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
