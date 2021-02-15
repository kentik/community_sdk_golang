package api_endpoints

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type UsersAPI struct {
	transport api_connection.Transport
}

func NewUsersAPI(transport api_connection.Transport) *UsersAPI {
	return &UsersAPI{transport: transport}
}

// GetAll lists users.
func (a *UsersAPI) GetAll(ctx context.Context) (_ []models.User, err error) {
	responseBody, err := a.transport.Get(ctx, usersPath)
	if err != nil {
		return nil, err
	}

	var data models.GetAllUsersResponse
	if err = json.Unmarshal(responseBody, &data); err != nil {
		return nil, fmt.Errorf("unmarshal response body: %v", err)
	}

	return data.Users, nil
}

// Get shows user with given ID.
func (a *UsersAPI) Get(ctx context.Context, id models.ID) (*models.User, error) {
	responseBody, err := a.transport.Get(ctx, getUserPath(id))
	if err != nil {
		return nil, err
	}

	var data models.GetUserResponse
	if err = json.Unmarshal(responseBody, &data); err != nil {
		return nil, fmt.Errorf("unmarshal response body: %v", err)
	}

	return &data.User, nil
}
