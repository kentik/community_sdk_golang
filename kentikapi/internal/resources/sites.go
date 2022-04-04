package resources

import (
	"context"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type SitesAPI struct {
	BaseAPI
}

// NewSitesAPI creates new SitesAPI.
func NewSitesAPI(transport api_connection.Transport, logPayloads bool) *SitesAPI {
	return &SitesAPI{
		BaseAPI{Transport: transport, LogPayloads: logPayloads},
	}
}

// GetAll sites.
func (a *SitesAPI) GetAll(ctx context.Context) ([]models.Site, error) {
	var response api_payloads.GetAllSitesResponse
	if err := a.GetAndValidate(ctx, api_endpoints.GetAllSites(), &response); err != nil {
		return []models.Site{}, err
	}

	return response.ToSites()
}

// Get site with given ID.
func (a *SitesAPI) Get(ctx context.Context, id models.ID) (*models.Site, error) {
	var response api_payloads.GetSiteResponse
	if err := a.GetAndValidate(ctx, api_endpoints.GetSite(id), &response); err != nil {
		return nil, err
	}

	site, err := response.ToSite()
	return &site, err
}

// Create a new site.
func (a *SitesAPI) Create(ctx context.Context, site models.Site) (*models.Site, error) {
	payload, err := api_payloads.SiteToPayload(site)
	if err != nil {
		return nil, err
	}

	request := api_payloads.CreateSiteRequest{Payload: payload}
	var response api_payloads.CreateSiteResponse
	if err = a.PostAndValidate(ctx, api_endpoints.CreateSite(), request, &response); err != nil {
		return nil, err
	}

	result, err := response.ToSite()
	return &result, err
}

// Update site.
func (a *SitesAPI) Update(ctx context.Context, site models.Site) (*models.Site, error) {
	payload, err := api_payloads.SiteToPayload(site)
	if err != nil {
		return nil, err
	}

	request := api_payloads.UpdateSiteRequest{Payload: payload}
	var response api_payloads.UpdateSiteResponse
	if err = a.UpdateAndValidate(ctx, api_endpoints.UpdateSite(site.ID), request, &response); err != nil {
		return nil, err
	}

	result, err := response.ToSite()
	return &result, err
}

// Delete site.
func (a *SitesAPI) Delete(ctx context.Context, id models.ID) error {
	if err := a.DeleteAndValidate(ctx, api_endpoints.DeleteSite(id), nil); err != nil {
		return err
	}

	return nil
}
