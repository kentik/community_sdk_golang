package resources

import (
	"context"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type PlansAPI struct {
	BaseAPI
}

// NewPlansAPI is constructor.
func NewPlansAPI(transport api_connection.Transport) *PlansAPI {
	return &PlansAPI{
		BaseAPI{Transport: transport},
	}
}

// GetAll plans.
func (a *PlansAPI) GetAll(ctx context.Context) ([]models.Plan, error) {
	var response api_payloads.GetAllPlansResponse
	if err := a.GetAndValidate(ctx, api_endpoints.GetAllPlans(), &response); err != nil {
		return []models.Plan{}, err
	}

	return response.ToPlans()
}
