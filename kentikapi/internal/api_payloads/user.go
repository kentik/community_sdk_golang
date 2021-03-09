package api_payloads

import (
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type GetAllUsersResponse struct {
	Users []userPayload `json:"users"`
}

func (r GetAllUsersResponse) ToUsers() []models.User {
	var users []models.User
	for _, up := range r.Users {
		users = append(users, *up.ToUser())
	}
	return users
}

type GetUserResponse struct {
	User userPayload `json:"user"`
}

func (r GetUserResponse) ToUser() *models.User {
	return r.User.ToUser()
}

type CreateUserRequest struct {
	User userPayload `json:"user"`
}

type CreateUserResponse = GetUserResponse

type userPayload struct {
	ID           IntAsString  `json:"id,omitempty"`
	Username     string       `json:"username"`
	UserFullName string       `json:"user_full_name"`
	UserEmail    string       `json:"user_email"`
	Role         string       `json:"role"`
	EmailService boolAsString `json:"email_service"`
	EmailProduct boolAsString `json:"email_product"`
	LastLogin    *time.Time   `json:"last_login,omitempty"`
	CreatedDate  *time.Time   `json:"created_date,omitempty" response:"get,post,put"`
	UpdatedDate  *time.Time   `json:"updated_date,omitempty" response:"get,post,put"`
	CompanyID    IntAsString  `json:"company_id,omitempty"`
	UserAPIToken string       `json:"user_api_token,omitempty"`
}

func (p userPayload) ToUser() *models.User {
	return &models.User{
		ID:           models.ID(p.ID),
		Username:     p.Username,
		UserFullName: p.UserFullName,
		UserEmail:    p.UserEmail,
		Role:         p.Role,
		EmailService: bool(p.EmailService),
		EmailProduct: bool(p.EmailProduct),
		LastLogin:    p.LastLogin,
		CreatedDate:  *p.CreatedDate,
		UpdatedDate:  *p.UpdatedDate,
		CompanyID:    models.ID(p.CompanyID),
		UserAPIToken: p.UserAPIToken,
	}
}

// UserToPayload prepares POST/PUT request payload: fill only the user-provided fields.
func UserToPayload(u models.User) userPayload {
	return userPayload{
		Username:     u.Username,
		UserFullName: u.UserFullName,
		UserEmail:    u.UserEmail,
		Role:         u.Role,
		EmailService: boolAsString(u.EmailService),
		EmailProduct: boolAsString(u.EmailProduct),
	}
}
