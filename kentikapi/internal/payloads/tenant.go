package payloads

import (
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/utils"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type GetAllTenantsResponse []TenantPayload

func (p GetAllTenantsResponse) ToTenants() (result []models.Tenant, err error) {
	convertFunc := func(d TenantPayload) (models.Tenant, error) {
		return d.ToTenant()
	}
	err = utils.ConvertList(p, convertFunc, &result)
	return result, err
}

type TenantPayload struct {
	ID          models.ID           `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Users       []TenantUserPayload `json:"users"`
	CreatedDate time.Time           `json:"created_date"`
	UpdatedDate time.Time           `json:"updated_date"`
}

type TenantUserPayload struct {
	ID        models.ID  `json:"id,string"`
	Email     string     `json:"user_email"`
	Name      *string    `json:"user_name,omitempty"`
	Fullname  *string    `json:"user_full_name,omitempty"`
	LastLogin *time.Time `json:"last_login,omitempty"`
	TenantID  models.ID  `json:"tenant_id,string"`
	CompanyID models.ID  `json:"company_id,string"`
}

func (p TenantPayload) ToTenant() (models.Tenant, error) {
	var users []models.TenantUser
	err := utils.ConvertList(p.Users, TenantUserPayload.ToTenantUser, &users)
	if err != nil {
		return models.Tenant{}, err
	}
	var companyID *models.ID
	if len(users) != 0 {
		companyID = &p.Users[0].CompanyID
	}
	return models.Tenant{
		ID:          p.ID,
		CompanyID:   companyID,
		Name:        p.Name,
		Description: p.Description,
		CreatedDate: p.CreatedDate,
		UpdatedDate: p.UpdatedDate,
		Users:       users,
	}, nil
}

func (p TenantUserPayload) ToTenantUser() (models.TenantUser, error) {
	return models.TenantUser{
		ID:        p.ID,
		Email:     p.Email,
		Name:      p.Name,
		Fullname:  p.Fullname,
		LastLogin: p.LastLogin,
		TenantID:  p.TenantID,
		CompanyID: p.CompanyID,
	}, nil
}

type CreateTenantUserRequest struct {
	User CreateTenantUserPayload `json:"user"`
}

type CreateTenantUserPayload struct {
	Email string `json:"user_email"`
}
