package api_payloads

import (
	"github.com/kentik/community_sdk_golang/kentikapi/internal/utils"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

// GetAllCustomDimensionsResponse represents CustomDimensionsAPI GetAll JSON response
type GetAllCustomDimensionsResponse struct {
	Payload []CustomDimensionPayload `json:"customDimensions"`
}

func (r GetAllCustomDimensionsResponse) ToCustomDimensions() (result []models.CustomDimension, err error) {
	err = utils.ConvertList(r.Payload, payloadToCustomDimension, &result)
	return result, err
}

// GetCustomDimensionResponse represents CustomDimensionsAPI Get JSON response
type GetCustomDimensionResponse struct {
	Payload CustomDimensionPayload `json:"customDimension"`
}

func (r GetCustomDimensionResponse) ToCustomDimension() (models.CustomDimension, error) {
	return payloadToCustomDimension(r.Payload)
}

// CreateCustomDimensionRequest represents CustomDimensionsAPI Create JSON request
type CreateCustomDimensionRequest CustomDimensionPayload

// CreateCustomDimensionResponse represents CustomDimensionsAPI Create JSON response
type CreateCustomDimensionResponse = GetCustomDimensionResponse

// UpdateCustomDimensionRequest represents CustomDimensionsAPI Update JSON request
type UpdateCustomDimensionRequest = CreateCustomDimensionRequest

// UpdateCustomDimensionResponse represents CustomDimensionsAPI Update JSON response
type UpdateCustomDimensionResponse = GetCustomDimensionResponse

// CustomDimensionPayload represents JSON CustomDimension payload as it is transmitted to and from KentikAPI
type CustomDimensionPayload struct {
	// following fields can appear in request: post/put, response: get/post/put
	DisplayName string `json:"display_name"` // display_name is always required

	// following fields can appear in request: post, response: get/post/put
	Name *string `json:"name" request:"post" response:"get,post,put"`
	Type *string `json:"type" request:"post" response:"get,post,put"`

	// following fields can appear in request: none, response: get/post/put
	Populators []PopulatorPayload `json:"populators" response:"get,post,put"`
	ID         *models.ID         `json:"id" response:"get,post,put"`
	CompanyID  *models.ID         `json:"company_id,string" response:"get,post,put"`
}

// payloadToCustomDimension transforms GET/POST/PUT response payload into CustomDimension
func payloadToCustomDimension(p CustomDimensionPayload) (models.CustomDimension, error) {
	cdType, err := models.CustomDimensionTypeString(*p.Type)
	if err != nil {
		return models.CustomDimension{}, err
	}

	var populators []models.Populator
	err = utils.ConvertList(p.Populators, payloadToPopulator, &populators)
	if err != nil {
		return models.CustomDimension{}, err
	}

	return models.CustomDimension{
		Name:        *p.Name,
		DisplayName: p.DisplayName,
		Type:        cdType,
		Populators:  populators,
		ID:          *p.ID,
		CompanyID:   *p.CompanyID,
	}, nil
}

// CustomDimensionToPayload prepares POST/PUT request payload: fill only the user-provided fields
func CustomDimensionToPayload(cd models.CustomDimension) CustomDimensionPayload {
	cdType := cd.Type.String()
	return CustomDimensionPayload{
		Name:        &cd.Name,
		DisplayName: cd.DisplayName,
		Type:        &cdType,
	}
}
