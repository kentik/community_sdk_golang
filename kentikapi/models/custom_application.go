package models

import "time"

type CustomApplication struct {
	// read-write properties (can be updated in update call)
	Name        string
	Description *string
	IPRange     *string
	Protocol    *string
	Port        *string
	ASN         *string

	// read-only properties (can't be updated in update call)
	ID          ID
	CompanyID   ID
	UserID      *ID
	CreatedDate *time.Time
	UpdatedDate *time.Time
}

// NewCustomApplication crates a CustomApplication with all required fields set
func NewCustomApplication(name string) *CustomApplication {
	return &CustomApplication{Name: name}
}
