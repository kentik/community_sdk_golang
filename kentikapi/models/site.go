package models

type Site struct {
	ID        ID
	SiteName  string
	Latitude  *float64
	Longitude *float64
	CompanyID ID
}

// NewSite creates a new Site with all necessary fields set
// Optional fields that can be set for Site include:
// - Longitude
// - Latitude.
func NewSite(name string) *Site {
	return &Site{
		SiteName: name,
	}
}
