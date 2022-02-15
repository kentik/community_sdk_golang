package api_payloads

import (
	"strconv"
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/utils"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

// GetAllCustomApplicationsResponse represents CustomApplicationsAPI GetAll JSON response.
type GetAllCustomApplicationsResponse []CustomApplicationPayload

func (r GetAllCustomApplicationsResponse) ToCustomApplications() (result []models.CustomApplication, err error) {
	err = utils.ConvertList(r, payloadToCustomApplication, &result)
	return result, err
}

// CreateCustomApplicationRequest represents CustomApplicationsAPI Create JSON request.
type CreateCustomApplicationRequest CustomApplicationPayload

// CreateCustomApplicationResponse represents CustomApplicationsAPI Create JSON response.
type CreateCustomApplicationResponse CustomApplicationPayload

func (r CreateCustomApplicationResponse) ToCustomApplication() (models.CustomApplication, error) {
	return payloadToCustomApplication(CustomApplicationPayload(r))
}

// UpdateCustomApplicationRequest represents CustomApplicationsAPI Update JSON request.
type UpdateCustomApplicationRequest = CreateCustomApplicationRequest

// UpdateCustomApplicationResponse represents CustomApplicationsAPI Update JSON response.
type UpdateCustomApplicationResponse = CreateCustomApplicationResponse

type CustomApplicationPayload struct {
	// following fields can appear in request: post/put, response: get/post/put
	Name        *string `json:"name" request:"post" response:"get,post,put"`
	Description *string `json:"description,omitempty"`
	IPRange     *string `json:"ip_range,omitempty"`
	Protocol    *string `json:"protocol,omitempty"`
	Port        *string `json:"port,omitempty"`
	ASN         *string `json:"asn,omitempty"`

	// following fields can appear in request: none, response: get/post/put
	ID          *int       `json:"id,omitempty" response:"get,post,put"`
	CompanyID   *string    `json:"company_id,omitempty" response:"get,post,put"`
	UserID      *string    `json:"user_id,omitempty"`                  // user_id happens to be returned as null
	CreatedDate *time.Time `json:"cdate,omitempty" response:"get,put"` // POST doesn't return cdate
	UpdatedDate *time.Time `json:"edate,omitempty" response:"get,put"` // POST doesn't return edate
}

// payloadToCustomApplication transforms GET/POST/PUT response payload into CustomApplication.
func payloadToCustomApplication(p CustomApplicationPayload) (models.CustomApplication, error) {
	return models.CustomApplication{
		Name:        *p.Name,
		Description: p.Description,
		IPRange:     p.IPRange,
		Protocol:    p.Protocol,
		Port:        p.Port,
		ASN:         p.ASN,
		ID:          strconv.Itoa(*p.ID),
		CompanyID:   *p.CompanyID,
		UserID:      p.UserID,
		CreatedDate: p.CreatedDate,
		UpdatedDate: p.UpdatedDate,
	}, nil
}

// CustomApplicationToPayload prepares POST/PUT request payload: fill only the user-provided fields.
func CustomApplicationToPayload(ca models.CustomApplication) CustomApplicationPayload {
	return CustomApplicationPayload{
		Name:        &ca.Name,
		Description: ca.Description,
		IPRange:     ca.IPRange,
		Protocol:    ca.Protocol,
		Port:        ca.Port,
		ASN:         ca.ASN,
	}
}
