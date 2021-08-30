package resources

import (
	"context"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type PlansAPI struct {
	BaseAPI
}

// NewPlansAPI is constructor.
func NewPlansAPI(transport connection.Transport) *PlansAPI {
	return &PlansAPI{
		BaseAPI{Transport: transport},
	}
}

// GetAll plans.
func (a *PlansAPI) GetAll(ctx context.Context) ([]models.Plan, error) {
	var response payloads.GetAllPlansResponse
	if err := a.GetAndValidate(ctx, endpoints.GetAllPlans(), &response); err != nil {
		return []models.Plan{}, err
	}

	return response.ToPlans()
}
