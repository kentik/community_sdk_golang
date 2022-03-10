package resources

import (
	"context"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/utils"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type QueryAPI struct {
	BaseAPI
}

// NewQueryAPI is constructor.
func NewQueryAPI(transport api_connection.Transport, logPayloads bool) *QueryAPI {
	return &QueryAPI{
		BaseAPI{Transport: transport, LogPayloads: logPayloads},
	}
}

// SQL query.
func (a *QueryAPI) SQL(ctx context.Context, sql string) (models.QuerySQLResult, error) {
	payload := api_payloads.QuerySQLRequest{Query: sql}
	utils.LogPayload(a.LogPayloads, "SQL query Kentik API request", payload)

	var response api_payloads.QuerySQLResponse
	if err := a.PostAndValidate(ctx, api_endpoints.QuerySQL(), payload, &response); err != nil {
		return models.QuerySQLResult{}, err
	}
	utils.LogPayload(a.LogPayloads, "SQL query Kentik API response", response)

	return response.ToQuerySQLResult(), nil
}

// Data query.
func (a *QueryAPI) Data(ctx context.Context, query models.QueryObject) (models.QueryDataResult, error) {
	payload, err := api_payloads.QueryObjectToPayload(query)
	if err != nil {
		return models.QueryDataResult{}, err
	}
	utils.LogPayload(a.LogPayloads, "Data query Kentik API request", payload)

	var response api_payloads.QueryDataResponse
	if err := a.PostAndValidate(ctx, api_endpoints.QueryData(), payload, &response); err != nil {
		return models.QueryDataResult{}, err
	}
	utils.LogPayload(a.LogPayloads, "Data query Kentik API response", response)

	return response.ToQueryDataResult(), nil
}

// Chart query.
func (a *QueryAPI) Chart(ctx context.Context, query models.QueryObject) (models.QueryChartResult, error) {
	payload, err := api_payloads.QueryObjectToPayload(query)
	if err != nil {
		return models.QueryChartResult{}, err
	}
	utils.LogPayload(a.LogPayloads, "Chart query Kentik API request", payload)

	var response api_payloads.QueryChartResponse
	if err := a.PostAndValidate(ctx, api_endpoints.QueryChart(), payload, &response); err != nil {
		return models.QueryChartResult{}, err
	}
	utils.LogPayload(a.LogPayloads, "Chart query Kentik API response", "")

	return response.ToQueryChartResult()
}

// URL query.
func (a *QueryAPI) URL(ctx context.Context, query models.QueryObject) (models.QueryURLResult, error) {
	payload, err := api_payloads.QueryObjectToPayload(query)
	if err != nil {
		return models.QueryURLResult{}, err
	}
	utils.LogPayload(a.LogPayloads, "URL query Kentik API request", payload)

	var response api_payloads.QueryURLResponse
	if err := a.PostAndValidate(ctx, api_endpoints.QueryURL(), payload, &response); err != nil {
		return models.QueryURLResult{}, err
	}
	utils.LogPayload(a.LogPayloads, "URL query Kentik API response", response)

	return response.ToQueryURLResult(), nil
}
