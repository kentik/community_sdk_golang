//nolint:dupl
package resources

import (
	"context"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/utils"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

// TagsAPI aggregates Tags API methods.
type TagsAPI struct {
	BaseAPI
}

// NewTagsAPI creates new TagsAPI.
func NewTagsAPI(transport api_connection.Transport, logPayloads bool) *TagsAPI {
	return &TagsAPI{BaseAPI{Transport: transport, LogPayloads: logPayloads}}
}

// GetAll lists tags.
func (a *TagsAPI) GetAll(ctx context.Context) ([]models.Tag, error) {
	utils.LogPayload(a.LogPayloads, "GetAll tags Kentik API request", "")
	var response api_payloads.GetAllTagsResponse
	if err := a.GetAndValidate(ctx, api_endpoints.TagsPath, &response); err != nil {
		return nil, err
	}
	utils.LogPayload(a.LogPayloads, "GetAll tags Kentik API response", response)

	return response.ToTags(), nil
}

// Get retrieves tag with given ID.
func (a *TagsAPI) Get(ctx context.Context, id models.ID) (*models.Tag, error) {
	utils.LogPayload(a.LogPayloads, "Get tag Kentik API request ID", id)
	var response api_payloads.GetTagResponse
	if err := a.GetAndValidate(ctx, api_endpoints.GetTagPath(id), &response); err != nil {
		return nil, err
	}
	utils.LogPayload(a.LogPayloads, "Get tag Kentik API response", response)

	return response.ToTag(), nil
}

// Create creates new tag.
func (a *TagsAPI) Create(ctx context.Context, tag models.Tag) (*models.Tag, error) {
	payload := api_payloads.CreateTagRequest{Tag: api_payloads.TagToPayload(tag)}
	utils.LogPayload(a.LogPayloads, "Create tag Kentik API request", payload)

	var response api_payloads.CreateTagResponse
	err := a.PostAndValidate(
		ctx,
		api_endpoints.TagPath,
		payload,
		&response,
	)
	utils.LogPayload(a.LogPayloads, "Create tag Kentik API response", response)
	if err != nil {
		return nil, err
	}

	return response.ToTag(), nil
}

// Update updates the tag.
func (a *TagsAPI) Update(ctx context.Context, tag models.Tag) (*models.Tag, error) {
	request := api_payloads.UpdateTagRequest{Tag: api_payloads.TagToPayload(tag)}
	utils.LogPayload(a.LogPayloads, "Update tag Kentik API request", request)

	var response api_payloads.UpdateTagResponse
	err := a.UpdateAndValidate(
		ctx,
		api_endpoints.GetTagPath(tag.ID),
		request,
		&response,
	)
	utils.LogPayload(a.LogPayloads, "Update tag Kentik API response", response)
	if err != nil {
		return nil, err
	}

	return response.ToTag(), err
}

// Delete removes tag with given ID.
func (a *TagsAPI) Delete(ctx context.Context, id models.ID) error {
	utils.LogPayload(a.LogPayloads, "Delete tag Kentik API request ID", id)
	if err := a.DeleteAndValidate(ctx, api_endpoints.GetTagPath(id), nil); err != nil {
		return err
	}

	return nil
}
