package resources

import (
	"context"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

// UsersAPI aggregates Users API methods.
type UsersAPI struct {
	BaseAPI
}

// NewUsersAPI creates new UsersAPI.
func NewUsersAPI(transport api_connection.Transport) *UsersAPI {
	return &UsersAPI{BaseAPI{Transport: transport}}
}

// GetAll lists users.
func (a *UsersAPI) GetAll(ctx context.Context) ([]models.User, error) {
	var response api_payloads.GetAllUsersResponse
	if err := a.GetAndValidate(ctx, api_endpoints.UsersPath, &response); err != nil {
		return nil, err
	}

	return response.ToUsers(), nil
}

// Get retrieves user with given ID.
func (a *UsersAPI) Get(ctx context.Context, id models.ID) (*models.User, error) {
	var response api_payloads.GetUserResponse
	if err := a.GetAndValidate(ctx, api_endpoints.GetUserPath(id), &response); err != nil {
		return nil, err
	}

	return response.ToUser(), nil
}

// Create creates new user.
func (a *UsersAPI) Create(ctx context.Context, user models.User) (*models.User, error) {
	var response api_payloads.CreateUserResponse
	err := a.PostAndValidate(
		ctx,
		api_endpoints.UserPath,
		api_payloads.CreateUserRequest{User: api_payloads.UserToPayload(user)},
		&response,
	)
	if err != nil {
		return nil, err
	}

	return response.ToUser(), nil
}

// Update updates the user.
func (a *UsersAPI) Update(ctx context.Context, user models.User) (*models.User, error) {
	var response api_payloads.UpdateUserResponse
	err := a.UpdateAndValidate(
		ctx,
		api_endpoints.GetUserPath(user.ID),
		api_payloads.UpdateUserRequest{User: api_payloads.UserToPayload(user)},
		&response,
	)
	if err != nil {
		return nil, err
	}

	return response.ToUser(), err
}

// Delete removes user with given ID.
func (a *UsersAPI) Delete(ctx context.Context, id models.ID) error {
	if err := a.DeleteAndValidate(ctx, api_endpoints.GetUserPath(id), nil); err != nil {
		return err
	}

	return nil
}
