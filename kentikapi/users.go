package kentikapi

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type usersAPI struct {
	transport transport
}

// GetAll lists users.
func (a *usersAPI) GetAll(ctx context.Context) (_ []models.User, err error) {
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
func (a *usersAPI) Get(ctx context.Context, id models.ID) (*models.User, error) {
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
