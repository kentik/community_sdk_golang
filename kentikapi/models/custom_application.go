package models

import "time"

type CustomApplication struct {
	// Read-write properties

	Name        string
	Description *string
	IPRange     *string
	Protocol    *string
	Port        *string
	ASN         *string

	// Read-only properties

	ID          ID
	CompanyID   ID
	UserID      *ID
	CreatedDate *time.Time
	UpdatedDate *time.Time
}

// NewCustomApplication crates a CustomApplication with all required fields set.
func NewCustomApplication(name string) *CustomApplication {
	return &CustomApplication{Name: name}
}
