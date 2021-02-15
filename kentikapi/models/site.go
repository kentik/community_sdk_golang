package models

type Site struct {
	ID        ID
	SiteName  string
	Latitude  *float64
	Longitude *float64
	CompanyID ID
}
