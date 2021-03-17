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

// NewCustomDimension creates a CustomDimension with all necessary fields set
// Note: name must begin with "c_" and be unique even among already deleted custom dimensions as names are retained for 1 year
func NewCustomDimension(name, displayName string, dimensionType CustomDimensionType) *CustomDimension {
	return &CustomDimension{
		Name:        name,
		DisplayName: displayName,
		Type:        dimensionType,
	}
}

type CustomDimensionType int

const (
	CustomDimensionTypeStr    CustomDimensionType = iota // "string"
	CustomDimensionTypeUint32                            // "uint32"
)
