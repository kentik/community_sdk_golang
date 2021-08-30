package resources

import (
	"context"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type SitesAPI struct {
	BaseAPI
}

// NewSitesAPI is constructor.
func NewSitesAPI(transport connection.Transport) *SitesAPI {
	return &SitesAPI{
		BaseAPI{Transport: transport},
	}
}

// GetAll sites.
func (a *SitesAPI) GetAll(ctx context.Context) ([]models.Site, error) {
	var response payloads.GetAllSitesResponse
	if err := a.GetAndValidate(ctx, endpoints.GetAllSites(), &response); err != nil {
		return []models.Site{}, err
	}

	return response.ToSites()
}

// Get site with given ID.
func (a *SitesAPI) Get(ctx context.Context, id models.ID) (*models.Site, error) {
	var response payloads.GetSiteResponse
	if err := a.GetAndValidate(ctx, endpoints.GetSite(id), &response); err != nil {
		return nil, err
	}

	site, err := response.ToSite()
	return &site, err
}

// Create new site.
func (a *SitesAPI) Create(ctx context.Context, site models.Site) (*models.Site, error) {
	payload, err := payloads.SiteToPayload(site)
	if err != nil {
		return nil, err
	}

	request := payloads.CreateSiteRequest{Payload: payload}
	var response payloads.CreateSiteResponse
	if err = a.PostAndValidate(ctx, endpoints.CreateSite(), request, &response); err != nil {
		return nil, err
	}

	result, err := response.ToSite()
	return &result, err
}

// Update site.
func (a *SitesAPI) Update(ctx context.Context, site models.Site) (*models.Site, error) {
	payload, err := payloads.SiteToPayload(site)
	if err != nil {
		return nil, err
	}

	request := payloads.UpdateSiteRequest{Payload: payload}
	var response payloads.UpdateSiteResponse
	if err = a.UpdateAndValidate(ctx, endpoints.UpdateSite(site.ID), request, &response); err != nil {
		return nil, err
	}

	result, err := response.ToSite()
	return &result, err
}

// Delete site.
func (a *SitesAPI) Delete(ctx context.Context, id models.ID) error {
	if err := a.DeleteAndValidate(ctx, endpoints.DeleteSite(id), nil); err != nil {
		return err
	}

	return nil
}
