package models

import "time"

// User model.
type User struct {
	// read-write properties
	Username     string
	UserFullName string
	UserEmail    string
	// Role of the user, e.g. "Member", "Administrator", "Super Administrator".
	Role         string
	EmailService bool
	EmailProduct bool

	// read-only properties
	ID           ID
	LastLogin    *time.Time
	CreatedDate  time.Time
	UpdatedDate  time.Time
	CompanyID    ID
	UserAPIToken *string
}

// UserRequiredFields is subset of User fields required to create a User.
type UserRequiredFields struct {
	Username     string
	UserFullName string
	UserEmail    string
	// Role of the user, e.g. "Member", "Administrator", "Super Administrator".
	Role         string
	EmailService bool
	EmailProduct bool
}

// NewUser creates a new User with all required fields set.
func NewUser(u UserRequiredFields) *User {
	return &User{
		Username:     u.Username,
		UserFullName: u.UserFullName,
		UserEmail:    u.UserEmail,
		Role:         u.Role,
		EmailService: u.EmailService,
		EmailProduct: u.EmailProduct,
	}
}
