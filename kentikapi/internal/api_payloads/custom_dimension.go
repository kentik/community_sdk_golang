package api_payloads

import (
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

// GetAllCustomDimensionsResponse represents CustomDimensionsAPI GetAll JSON response
type GetAllCustomDimensionsResponse struct {
	Payload []CustomDimensionPayload `json:"customDimensions"`
}

func (r GetAllCustomDimensionsResponse) ToCustomDimensions() []models.CustomDimension {
	result := make([]models.CustomDimension, 0, len(r.Payload))
	for _, p := range r.Payload {
		result = append(result, payloadToCustomDimension(p))
	}
	return result
}

// GetCustomDimensionResponse represents CustomDimensionsAPI Get JSON response
type GetCustomDimensionResponse struct {
	Payload CustomDimensionPayload `json:"customDimension"`
}

func (r GetCustomDimensionResponse) ToCustomDimension() models.CustomDimension {
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
func payloadToCustomDimension(p CustomDimensionPayload) models.CustomDimension {
	return models.CustomDimension{
		Name:        *p.Name,
		DisplayName: p.DisplayName,
		Type:        models.CustomDimensionType(*p.Type),
		Populators:  payloadToPopulators(p.Populators),
		ID:          *p.ID,
		CompanyID:   *p.CompanyID,
	}
}

// CustomDimensionToPayload prepares POST/PUT request payload: fill only the user-provided fields
func CustomDimensionToPayload(cd models.CustomDimension) CustomDimensionPayload {
	cdType := string(cd.Type)
	return CustomDimensionPayload{
		Name:        &cd.Name,
		DisplayName: cd.DisplayName,
		Type:        &cdType,
	}
}
