package api_payloads

import "github.com/kentik/community_sdk_golang/kentikapi/models"

// SitePayload represents JSON Plan payload as it is transmitted to and from KentikAPI
type SitePayload struct {
	ID        models.ID `json:"id"`
	SiteName  string    `json:"site_name"`
	Latitude  *float64  `json:"lat"`
	Longitude *float64  `json:"lon"`
	CompanyID models.ID `json:"company_id" response:"get,post,put"`
}
