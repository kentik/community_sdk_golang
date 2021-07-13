/*
 * Synthetics Monitoring API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 202101beta1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package syntheticsstub

import (
	"context"
	"net/http"
)

// SyntheticsDataServiceApiService is a service that implements the logic for the SyntheticsDataServiceApiServicer
// This service should implement the business logic for every endpoint for the SyntheticsDataServiceApi API.
// Include any external packages or services that will be required by this service.
type SyntheticsDataServiceApiService struct {
	repo *SyntheticsRepo
}

// NewSyntheticsDataServiceApiService creates a default api service
func NewSyntheticsDataServiceApiService(repo *SyntheticsRepo) SyntheticsDataServiceApiServicer {
	return &SyntheticsDataServiceApiService{
		repo: repo,
	}
}

// GetHealthForTests - Get health status for synthetics test.
func (s *SyntheticsDataServiceApiService) GetHealthForTests(_ context.Context, body V202101beta1GetHealthForTestsRequest) (ImplResponse, error) {
	return Response(http.StatusOK, &V202101beta1GetHealthForTestsResponse{
		Health: s.repo.GetHealthForTests(),
	}), nil
}

// GetTraceForTest - Get trace route data.
func (s *SyntheticsDataServiceApiService) GetTraceForTest(_ context.Context, id string, body V202101beta1GetTraceForTestRequest) (ImplResponse, error) {
	return Response(http.StatusOK, &V202101beta1GetTraceForTestResponse{
		IpInfo:      s.repo.GetIpInfo(),
		TraceRoutes: s.repo.GetTraceRoutes(),
	}), nil
}
