package api_payloads

import "github.com/kentik/community_sdk_golang/kentikapi/models"

type SitePayload struct {
	ID        models.ID `json:"id" response:"get,post,put"`
	SiteName  string    `json:"site_name" request:"post" response:"get,post,put"`
	Latitude  *float64  `json:"lat"`
	Longitude *float64  `json:"lon"`
	CompanyID models.ID `json:"company_id" request:"get" response:"get,post,put"`
}
