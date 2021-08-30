package resources

import (
	"context"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type QueryAPI struct {
	BaseAPI
}

// NewQueryAPI is constructor.
func NewQueryAPI(transport connection.Transport) *QueryAPI {
	return &QueryAPI{
		BaseAPI{Transport: transport},
	}
}

// SQL query.
func (a *QueryAPI) SQL(ctx context.Context, sql string) (models.QuerySQLResult, error) {
	payload := payloads.QuerySQLRequest{Query: sql}

	var response payloads.QuerySQLResponse
	if err := a.PostAndValidate(ctx, endpoints.QuerySQL(), payload, &response); err != nil {
		return models.QuerySQLResult{}, err
	}

	return response.ToQuerySQLResult(), nil
}

// Data query.
func (a *QueryAPI) Data(ctx context.Context, query models.QueryObject) (models.QueryDataResult, error) {
	payload, err := payloads.QueryObjectToPayload(query)
	if err != nil {
		return models.QueryDataResult{}, err
	}

	var response payloads.QueryDataResponse
	if err := a.PostAndValidate(ctx, endpoints.QueryData(), payload, &response); err != nil {
		return models.QueryDataResult{}, err
	}

	return response.ToQueryDataResult(), nil
}

// Chart query.
func (a *QueryAPI) Chart(ctx context.Context, query models.QueryObject) (models.QueryChartResult, error) {
	payload, err := payloads.QueryObjectToPayload(query)
	if err != nil {
		return models.QueryChartResult{}, err
	}

	var response payloads.QueryChartResponse
	if err := a.PostAndValidate(ctx, endpoints.QueryChart(), payload, &response); err != nil {
		return models.QueryChartResult{}, err
	}

	return response.ToQueryChartResult()
}

// URL query.
func (a *QueryAPI) URL(ctx context.Context, query models.QueryObject) (models.QueryURLResult, error) {
	payload, err := payloads.QueryObjectToPayload(query)
	if err != nil {
		return models.QueryURLResult{}, err
	}

	var response payloads.QueryURLResponse
	if err := a.PostAndValidate(ctx, endpoints.QueryURL(), payload, &response); err != nil {
		return models.QueryURLResult{}, err
	}

	return response.ToQueryURLResult(), nil
}
