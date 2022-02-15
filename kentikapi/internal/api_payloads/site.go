package api_payloads

import (
	"github.com/kentik/community_sdk_golang/kentikapi/internal/utils"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

// GetSiteResponse represents SitesAPI Get JSON response.
type GetSiteResponse struct {
	Payload SitePayload `json:"site"`
}

func (r GetSiteResponse) ToSite() (result models.Site, err error) {
	return payloadToSite(r.Payload)
}

// GetAllSitesResponse represents SitesAPI GetAll JSON response.
type GetAllSitesResponse struct {
	Payload []SitePayload `json:"sites"`
}

func (r GetAllSitesResponse) ToSites() (result []models.Site, err error) {
	err = utils.ConvertList(r.Payload, payloadToSite, &result)
	return result, err
}

// CreateSiteRequest represents SitesAPI Create JSON request.
type CreateSiteRequest struct {
	Payload SitePayload `json:"site"`
}

// CreateSiteResponse represents SitesAPI Create JSON Response.
type CreateSiteResponse = GetSiteResponse

// UpdateSiteRequest represents SitesAPI Update JSON request.
type UpdateSiteRequest = CreateSiteRequest

// UpdateSiteResponse represents SitesAPI Update JSON response.
type UpdateSiteResponse = CreateSiteResponse

// SitePayload represents JSON Plan payload as it is transmitted to and from KentikAPI.
type SitePayload struct {
	ID        StringAsInt `json:"id"`        // caveat, POST and GET return id as int but PUT as string
	SiteName  string      `json:"site_name"` // site_name is required always, also in PUT
	Latitude  *float64    `json:"lat,omitempty"`
	Longitude *float64    `json:"lon,omitempty"`
	CompanyID StringAsInt `json:"company_id"` // caveat, GET returns company_id as int but POST and PUT as string
}

func payloadToSite(p SitePayload) (models.Site, error) {
	return models.Site{
		ID:        models.ID(p.ID),
		SiteName:  p.SiteName,
		Latitude:  p.Latitude,
		Longitude: p.Longitude,
		CompanyID: models.ID(p.CompanyID),
	}, nil
}

func SiteToPayload(site models.Site) (SitePayload, error) {
	return SitePayload{
		ID:        StringAsInt(site.ID),
		SiteName:  site.SiteName,
		Latitude:  site.Latitude,
		Longitude: site.Longitude,
		CompanyID: StringAsInt(site.CompanyID),
	}, nil
}
