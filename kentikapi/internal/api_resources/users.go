package api_resources

import (
	"context"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type UsersAPI struct {
	BaseAPI
}

func NewUsersAPI(transport api_connection.Transport) *UsersAPI {
	return &UsersAPI{BaseAPI{Transport: transport}}
}

// GetAll lists users.
func (a *UsersAPI) GetAll(ctx context.Context) (_ []models.User, err error) {
	var response models.GetAllUsersResponse
	if err := a.GetAndValidate(ctx, api_endpoints.UsersPath, &response); err != nil {
		return nil, err
	}

	return response.Users, nil
}

// Get shows user with given ID.
func (a *UsersAPI) Get(ctx context.Context, id models.ID) (*models.User, error) {
	var response models.GetUserResponse
	if err := a.GetAndValidate(ctx, api_endpoints.GetUserPath(id), &response); err != nil {
		return nil, err
	}

	return &response.User, nil
}
