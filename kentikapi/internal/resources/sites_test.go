package resources_test

import (
	"context"
	"testing"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/resources"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/utils"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSite(t *testing.T) {
	t.Parallel()

	// arrange
	getResponsePayload := `
	{
		"site": {
			"id": 42,
			"site_name": "apitest-site-1",
			"lat": 54.349276,
			"lon": 18.659577,
			"company_id": 3250
		}
	}`
	transport := &connection.StubTransport{ResponseBody: getResponsePayload}
	sitesAPI := resources.NewSitesAPI(transport)
	siteID := models.ID(42)

	// act
	site, err := sitesAPI.Get(context.Background(), siteID)

	// assert request properly formed
	assert := assert.New(t)
	require := require.New(t)

	require.NoError(err)
	assert.Zero(transport.RequestBody)

	// and response properly parsed
	assert.Equal(models.ID(42), site.ID)
	assert.Equal("apitest-site-1", site.SiteName)
	assert.Equal(54.349276, *site.Latitude)
	assert.Equal(18.659577, *site.Longitude)
	assert.Equal(models.ID(3250), site.CompanyID)
}

func TestGetAllSites(t *testing.T) {
	t.Parallel()

	// arrange
	getResponsePayload := `
	{
		"sites": [
			{
				"id": 7758,
				"site_name": "AWS us-east-1",
				"lat": null,
				"lon": null,
				"company_id": 74333
			},
			{
				"id": 8483,
				"site_name": "marina gdańsk",
				"lat": 54.348972,
				"lon": 18.659791,
				"company_id": 74333
			},
			{
				"id": 8592,
				"site_name": "mysite",
				"lat": 1,
				"lon": 2,
				"company_id": 74333
			}
		]
	}`
	transport := &connection.StubTransport{ResponseBody: getResponsePayload}
	sitesAPI := resources.NewSitesAPI(transport)

	// act
	sites, err := sitesAPI.GetAll(context.Background())

	// assert request properly formed
	assert := assert.New(t)
	require := require.New(t)

	require.NoError(err)
	assert.Zero(transport.RequestBody)

	// and response properly parsed
	require.Equal(3, len(sites))
	// site 0
	site := sites[0]
	assert.Equal(models.ID(7758), site.ID)
	assert.Equal("AWS us-east-1", site.SiteName)
	assert.Nil(site.Latitude)
	assert.Nil(site.Longitude)
	assert.Equal(models.ID(74333), site.CompanyID)
	// site 1
	site = sites[1]
	assert.Equal(models.ID(8483), site.ID)
	assert.Equal("marina gdańsk", site.SiteName)
	assert.Equal(54.348972, *site.Latitude)
	assert.Equal(18.659791, *site.Longitude)
	assert.Equal(models.ID(74333), site.CompanyID)
	// site 2
	site = sites[2]
	assert.Equal(models.ID(8592), site.ID)
	assert.Equal("mysite", site.SiteName)
	assert.Equal(1.0, *site.Latitude)
	assert.Equal(2.0, *site.Longitude)
	assert.Equal(models.ID(74333), site.CompanyID)
}

func TestCreateSite(t *testing.T) {
	t.Parallel()

	// arrange
	createResponsePayload := `
	{
		"site": {
			"id": 42,
			"site_name": "apitest-site-1",
			"lat": 54.349276,
			"lon": 18.659577,
			"company_id": "3250"
		}     
	}`
	transport := &connection.StubTransport{ResponseBody: createResponsePayload}
	sitesAPI := resources.NewSitesAPI(transport)

	// act
	site := models.NewSite("apitest-site-1")
	models.SetOptional(&site.Longitude, 18.659577)
	models.SetOptional(&site.Latitude, 54.349276)
	created, err := sitesAPI.Create(context.Background(), *site)

	// assert request properly formed
	assert := assert.New(t)
	require := require.New(t)

	require.NoError(err)
	payload := utils.NewJSONPayloadInspector(t, transport.RequestBody)
	require.NotNil(payload.Get("site"))
	assert.Equal("apitest-site-1", payload.String("site/site_name"))
	assert.Equal(18.659577, payload.Float("site/lon"))
	assert.Equal(54.349276, payload.Float("site/lat"))

	// and response properly parsed
	assert.Equal(models.ID(42), created.ID)
	assert.Equal("apitest-site-1", created.SiteName)
	assert.Equal(54.349276, *created.Latitude)
	assert.Equal(18.659577, *created.Longitude)
	assert.Equal(models.ID(3250), created.CompanyID)
}

func TestUpdateSite(t *testing.T) {
	t.Parallel()

	// arrange
	updateResponsePayload := `
	{
		"site": {
			"id": "42",
			"site_name": "new-site",
			"lat": -15.0,
			"lon": -45.0,
			"company_id": "3250"
		}
	}`
	transport := &connection.StubTransport{ResponseBody: updateResponsePayload}
	sitesAPI := resources.NewSitesAPI(transport)

	// act
	siteID := models.ID(42)
	site := models.Site{ID: siteID, SiteName: "new-site"}
	models.SetOptional(&site.Longitude, -45.0)
	updated, err := sitesAPI.Update(context.Background(), site)

	// assert request properly formed
	assert := assert.New(t)
	require := require.New(t)

	require.NoError(err)
	payload := utils.NewJSONPayloadInspector(t, transport.RequestBody)
	require.NotNil(payload.Get("site"))
	assert.Nil(payload.Get("site/lat"))
	assert.Equal("new-site", payload.String("site/site_name"))
	assert.Equal(-45.0, payload.Float("site/lon"))

	// # and response properly parsed
	assert.Equal(models.ID(42), updated.ID)
	assert.Equal("new-site", updated.SiteName)
	assert.Equal(-15.0, *updated.Latitude)
	assert.Equal(-45.0, *updated.Longitude)
	assert.Equal(models.ID(3250), updated.CompanyID)
}

func TestDeleteSite(t *testing.T) {
	t.Parallel()

	// arrange
	deleteResponsePayload := "" // deleting site responds with empty body
	transport := &connection.StubTransport{ResponseBody: deleteResponsePayload}
	sitesAPI := resources.NewSitesAPI(transport)

	// act
	siteID := models.ID(42)
	err := sitesAPI.Delete(context.Background(), siteID)

	// assert
	require.NoError(t, err)
	assert.Zero(t, transport.RequestBody)
}
