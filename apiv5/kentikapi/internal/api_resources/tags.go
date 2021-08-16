package api_resources

import (
	"context"

	"github.com/kentik/community_sdk_golang/apiv5/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/apiv5/kentikapi/internal/api_endpoints"
	"github.com/kentik/community_sdk_golang/apiv5/kentikapi/internal/api_payloads"
	"github.com/kentik/community_sdk_golang/apiv5/kentikapi/models"
)

// TagsAPI aggregates Tags API methods.
type TagsAPI struct {
	BaseAPI
}

// NewTagsAPI creates new TagsAPI.
func NewTagsAPI(transport api_connection.Transport) *TagsAPI {
	return &TagsAPI{BaseAPI{Transport: transport}}
}

// GetAll lists tags.
func (a *TagsAPI) GetAll(ctx context.Context) ([]models.Tag, error) {
	var response api_payloads.GetAllTagsResponse
	if err := a.GetAndValidate(ctx, api_endpoints.TagsPath, &response); err != nil {
		return nil, err
	}

	return response.ToTags(), nil
}

// Get retrieves tag with given ID.
func (a *TagsAPI) Get(ctx context.Context, id models.ID) (*models.Tag, error) {
	var response api_payloads.GetTagResponse
	if err := a.GetAndValidate(ctx, api_endpoints.GetTagPath(id), &response); err != nil {
		return nil, err
	}

	return response.ToTag(), nil
}

// Create creates new tag.
func (a *TagsAPI) Create(ctx context.Context, tag models.Tag) (*models.Tag, error) {
	var response api_payloads.CreateTagResponse
	err := a.PostAndValidate(
		ctx,
		api_endpoints.TagPath,
		api_payloads.CreateTagRequest{Tag: api_payloads.TagToPayload(tag)},
		&response,
	)
	if err != nil {
		return nil, err
	}

	return response.ToTag(), nil
}

// Update updates the tag.
func (a *TagsAPI) Update(ctx context.Context, tag models.Tag) (*models.Tag, error) {
	var response api_payloads.UpdateTagResponse
	err := a.UpdateAndValidate(
		ctx,
		api_endpoints.GetTagPath(tag.ID),
		api_payloads.UpdateTagRequest{Tag: api_payloads.TagToPayload(tag)},
		&response,
	)
	if err != nil {
		return nil, err
	}

	return response.ToTag(), err
}

// Delete removes tag with given ID.
func (a *TagsAPI) Delete(ctx context.Context, id models.ID) error {
	if err := a.DeleteAndValidate(ctx, api_endpoints.GetTagPath(id), nil); err != nil {
		return err
	}

	return nil
}
