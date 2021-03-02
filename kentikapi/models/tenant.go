package models

import "time"

type Tenant struct {
	ID          ID
	CompanyID   *ID
	Name        string
	Description string
	CreatedDate time.Time
	UpdatedDate time.Time
	Users       []TenantUser
}

type TenantUser struct {
	ID        ID
	CompanyID ID
	Email     string
	Name      *string
	Fullname  *string
	TenantID  ID
	LastLogin *time.Time
}
