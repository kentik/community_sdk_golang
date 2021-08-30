//nolint:dupl
package resources

import (
	"context"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

// TagsAPI aggregates Tags API methods.
type TagsAPI struct {
	BaseAPI
}

// NewTagsAPI creates new TagsAPI.
func NewTagsAPI(transport connection.Transport) *TagsAPI {
	return &TagsAPI{BaseAPI{Transport: transport}}
}

// GetAll lists tags.
func (a *TagsAPI) GetAll(ctx context.Context) ([]models.Tag, error) {
	var response payloads.GetAllTagsResponse
	if err := a.GetAndValidate(ctx, endpoints.TagsPath, &response); err != nil {
		return nil, err
	}

	return response.ToTags(), nil
}

// Get retrieves tag with given ID.
func (a *TagsAPI) Get(ctx context.Context, id models.ID) (*models.Tag, error) {
	var response payloads.GetTagResponse
	if err := a.GetAndValidate(ctx, endpoints.GetTagPath(id), &response); err != nil {
		return nil, err
	}

	return response.ToTag(), nil
}

// Create creates new tag.
func (a *TagsAPI) Create(ctx context.Context, tag models.Tag) (*models.Tag, error) {
	var response payloads.CreateTagResponse
	err := a.PostAndValidate(
		ctx,
		endpoints.TagPath,
		payloads.CreateTagRequest{Tag: payloads.TagToPayload(tag)},
		&response,
	)
	if err != nil {
		return nil, err
	}

	return response.ToTag(), nil
}

// Update updates the tag.
func (a *TagsAPI) Update(ctx context.Context, tag models.Tag) (*models.Tag, error) {
	var response payloads.UpdateTagResponse
	err := a.UpdateAndValidate(
		ctx,
		endpoints.GetTagPath(tag.ID),
		payloads.UpdateTagRequest{Tag: payloads.TagToPayload(tag)},
		&response,
	)
	if err != nil {
		return nil, err
	}

	return response.ToTag(), err
}

// Delete removes tag with given ID.
func (a *TagsAPI) Delete(ctx context.Context, id models.ID) error {
	if err := a.DeleteAndValidate(ctx, endpoints.GetTagPath(id), nil); err != nil {
		return err
	}

	return nil
}
