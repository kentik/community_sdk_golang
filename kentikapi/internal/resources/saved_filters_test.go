package resources_test

import (
	"context"
	"testing"
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/resources"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
)

func TestSavedFiltersList(t *testing.T) {
	t.Parallel()

	getAllresponsePayload := `
	[
	    {
	        "id": 8162,
	        "company_id": "74333",
	        "filters": {
	            "connector": "All",
	            "filterGroups": [
	                {
	                    "name": "",
	                    "named": false,
	                    "connector": "Any",
	                    "not": false,
	                    "autoAdded": "",
	                    "filters": [
	                        {
	                            "filterField": "inet_src_addr",
	                            "operator": "ILIKE",
	                            "filterValue": "1.2.3.4,172.217.169.46"
	                        }
	                    ],
	                    "saved_filters": [],
	                    "filterGroups": []
	                }
	            ]
	        },
	        "filter_name": "test-filter-dev",
	        "filter_description": "Test filter used for Kentik API Python library development purposes",
	        "cdate": "2020-12-18T16:57:00.475Z",
	        "edate": "2020-12-18T18:13:29.955Z",
	        "filter_level": "company"
	    },
	    {
	        "id": 8275,
	        "company_id": "74333",
	        "filters": {
	            "connector": "All",
	            "filterGroups": [
	                {
	                    "connector": "Any",
	                    "filters": [
	                        {
	                            "filterField": "dst_as",
	                            "operator": "=",
	                            "filterValue": "5"
	                        }
	                    ],
	                    "not": false
	                }
	            ]
	        },
	        "filter_name": "test_filter",
	        "filter_description": "This is test filter description1",
	        "cdate": "2021-02-08T11:59:22.272Z",
	        "edate": "2021-03-04T08:46:12.322Z",
	        "filter_level": "company"
	    }
	]`

	expected := []models.SavedFilter{
		{
			ID:        8162,
			CompanyID: 74333,
			Filters: models.Filters{
				Connector: "All",
				FilterGroups: []models.FilterGroups{
					{
						Connector: "Any",
						Not:       false,
						Filters: []models.Filter{
							{
								FilterField: "inet_src_addr",
								Operator:    "ILIKE",
								FilterValue: "1.2.3.4,172.217.169.46",
							},
						},
					},
				},
			},
			FilterName:        "test-filter-dev",
			FilterDescription: "Test filter used for Kentik API Python library development purposes",
			CreatedDate:       time.Date(2020, 12, 18, 16, 57, 0, 475e6, time.UTC),
			UpdatedDate:       time.Date(2020, 12, 18, 18, 13, 29, 955e6, time.UTC),
			FilterLevel:       "company",
		},
		{
			ID:        8275,
			CompanyID: 74333,
			Filters: models.Filters{
				Connector: "All",
				FilterGroups: []models.FilterGroups{
					{
						Connector: "Any",
						Filters: []models.Filter{
							{
								FilterField: "dst_as",
								Operator:    "=",
								FilterValue: "5",
							},
						},
						Not: false,
					},
				},
			},
			FilterName:        "test_filter",
			FilterDescription: "This is test filter description1",
			CreatedDate:       time.Date(2021, 2, 8, 11, 59, 22, 272e6, time.UTC),
			UpdatedDate:       time.Date(2021, 3, 4, 8, 46, 12, 322e6, time.UTC),
			FilterLevel:       "company",
		},
	}

	transport := &connection.StubTransport{ResponseBody: getAllresponsePayload}
	savedFiltersAPI := resources.NewSavedFiltersAPI(transport)

	savedFilters, err := savedFiltersAPI.GetAll(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, expected, savedFilters)
	assert.Empty(t, transport.RequestBody)
	assert.Equal(t, "/saved-filters/custom", transport.RequestPath)
}

func TestGetSavedFilterInfo(t *testing.T) {
	t.Parallel()

	getResponsePayload := `
	{
		"id": 8275,
		"company_id": "74333",
		"filters": {
			"connector": "All",
			"filterGroups": [
				{
					"connector": "Any",
					"filters": [
						{
							"filterField": "dst_as",
							"operator": "=",
							"filterValue": "5"
						}
					],
					"not": false
				}
			]
		},
		"filter_name": "test_filter",
		"filter_description": "This is test filter description1",
		"cdate": "2021-02-08T11:59:22.272Z",
		"edate": "2021-03-04T08:46:12.322Z",
		"filter_level": "company"
	}`
	expected := models.SavedFilter{
		ID:        8275,
		CompanyID: 74333,
		Filters: models.Filters{
			Connector: "All",
			FilterGroups: []models.FilterGroups{
				{
					Connector: "Any",
					Filters: []models.Filter{
						{
							FilterField: "dst_as",
							Operator:    "=",
							FilterValue: "5",
						},
					},
					Not: false,
				},
			},
		},
		FilterName:        "test_filter",
		FilterDescription: "This is test filter description1",
		CreatedDate:       time.Date(2021, 2, 8, 11, 59, 22, 272e6, time.UTC),
		UpdatedDate:       time.Date(2021, 3, 4, 8, 46, 12, 322e6, time.UTC),
		FilterLevel:       "company",
	}

	transport := &connection.StubTransport{ResponseBody: getResponsePayload}
	savedFiltersAPI := resources.NewSavedFiltersAPI(transport)

	savedFilter, err := savedFiltersAPI.Get(context.Background(), 8275)

	assert.NoError(t, err)
	assert.Equal(t, &expected, savedFilter)
	assert.Empty(t, transport.RequestBody)
	assert.Equal(t, "/saved-filter/custom/8275", transport.RequestPath)
}

func TestCreateSavedFilter(t *testing.T) {
	t.Parallel()

	postResponsePayload := `
	{
        "filter_name":"test_filter1",
        "filter_description":"This is test filter description",
        "filters": {
            "connector":"All",
            "filterGroups": [
                {
                    "connector":"All",
                    "filters": [
                        {
                            "filterField":"dst_as",
                            "filterValue":"81",
                            "operator":"="
                        }],
                    "not":false
                }]
            },
        "company_id":"74333",
        "filter_level":"company",
        "edate":"2020-12-26T10:46:13.095Z",
        "cdate":"2020-12-16T10:46:13.095Z",
        "id":8152
    }`
	expected := models.SavedFilter{
		FilterName:        "test_filter1",
		FilterDescription: "This is test filter description",
		Filters: models.Filters{
			Connector: "All",
			FilterGroups: []models.FilterGroups{
				{
					Connector: "All",
					Filters: []models.Filter{
						{
							FilterField: "dst_as",
							FilterValue: "81",
							Operator:    "=",
						},
					},
					Not: false,
				},
			},
		},
		CompanyID:   74333,
		FilterLevel: "company",
		CreatedDate: time.Date(2020, 12, 16, 10, 46, 13, 95e6, time.UTC),
		UpdatedDate: time.Date(2020, 12, 26, 10, 46, 13, 95e6, time.UTC),
		ID:          8152,
	}
	// TODO(lwolanin): To test request payloads use JSONPayloadInspector like in most of tests
	expectedRequestPayload := "{\"filter_name\":\"test_filter1\"," +
		"\"filter_description\":\"This is test filter description\",\"cdate\":\"0001-01-01T00:00:00Z\"," +
		"\"edate\":\"0001-01-01T00:00:00Z\",\"filters\":{\"connector\":\"All\"," +
		"\"filterGroups\":[{\"connector\":\"All\",\"not\":false,\"filters\":[{\"filterField\":\"dst_as\"," +
		"\"filterValue\":\"81\",\"operator\":\"=\"}]}]}}"

	transport := &connection.StubTransport{ResponseBody: postResponsePayload}
	savedFiltersAPI := resources.NewSavedFiltersAPI(transport)

	newSavedFilter := models.SavedFilter{
		FilterName:        "test_filter1",
		FilterDescription: "This is test filter description",
		Filters: models.Filters{
			Connector: "All",
			FilterGroups: []models.FilterGroups{
				{
					Connector: "All",
					Not:       false,
					Filters: []models.Filter{
						{
							FilterField: "dst_as",
							FilterValue: "81",
							Operator:    "=",
						},
					},
				},
			},
		},
	}
	savedFilter, err := savedFiltersAPI.Create(context.Background(), newSavedFilter)

	assert.NoError(t, err)
	assert.Equal(t, &expected, savedFilter)
	assert.Equal(t, "/saved-filter/custom", transport.RequestPath)
	assert.Equal(t, expectedRequestPayload, transport.RequestBody)
}

func TestUpdateSavedFilter(t *testing.T) {
	t.Parallel()

	updateResponsePayload := `
	{
		"id":8153,
		"company_id":"74333",
		"filters":{
				"connector":"All",
				"filterGroups":[
					{
						"connector":"All",
						"filters":[{
								"filterField":"dst_as",
								"filterValue":"81",
								"operator":"="
							}],
						"not":false
					}]
			},
		"filter_name":"test_filter1",
		"filter_description":"Updated Saved Filter description",
		"cdate":"2020-12-16T11:26:18.578Z",
		"edate":"2020-12-16T11:26:19.187Z",
		"filter_level":"company"
	}`
	expectedRequestPayload := "{\"id\":8153,\"filter_name\":\"test_filter1\"," +
		"\"filter_description\":\"Updated Saved Filter description\",\"cdate\":\"0001-01-01T00:00:00Z\"," +
		"\"edate\":\"0001-01-01T00:00:00Z\",\"filters\":{\"connector\":\"All\"," +
		"\"filterGroups\":[{\"connector\":\"All\",\"not\":false,\"filters\":[{\"filterField\":\"dst_as\"," +
		"\"filterValue\":\"81\",\"operator\":\"=\"}]}]}}"

	transport := &connection.StubTransport{ResponseBody: updateResponsePayload}
	savedFiltersAPI := resources.NewSavedFiltersAPI(transport)

	filterID := 8153
	toUpdate := models.SavedFilter{
		FilterName: "test_filter1",
		Filters: models.Filters{
			Connector: "All",
			FilterGroups: []models.FilterGroups{
				{
					Connector: "All",
					Not:       false,
					Filters: []models.Filter{
						{
							FilterField: "dst_as",
							FilterValue: "81",
							Operator:    "=",
						},
					},
				},
			},
		},
		ID:                filterID,
		FilterDescription: "Updated Saved Filter description",
	}
	updated, err := savedFiltersAPI.Update(context.Background(), toUpdate)

	assert.NoError(t, err)
	assert.Equal(t, "Updated Saved Filter description", updated.FilterDescription)
	assert.Equal(t, expectedRequestPayload, transport.RequestBody)
	assert.Equal(t, "/saved-filter/custom/8153", transport.RequestPath)
}

func TestDeleteSavedFilter(t *testing.T) {
	t.Parallel()

	deleteResponsePayload := ""

	transport := &connection.StubTransport{ResponseBody: deleteResponsePayload}
	savedFiltersAPI := resources.NewSavedFiltersAPI(transport)

	filterID := 8153
	err := savedFiltersAPI.Detete(context.Background(), filterID)

	assert.NoError(t, err)
	assert.Empty(t, transport.RequestBody)
	assert.Equal(t, "/saved-filter/custom/8153", transport.RequestPath)
}
