package models

import "time"

type GetAllUsersResponse struct {
	Users []User `json:"users"`
}

type GetUserResponse struct {
	User User `json:"user"`
}

type User struct {
	ID           ID         `json:"id,string"`
	Username     string     `json:"username"`
	UserFullName string     `json:"user_full_name"`
	UserEmail    string     `json:"user_email"`
	Role         string     `json:"role"`
	EmailService bool       `json:"email_service"`
	EmailProduct bool       `json:"email_product"`
	LastLogin    *time.Time `json:"last_login"`
	CreatedDate  time.Time  `json:"created_date"`
	UpdatedDate  time.Time  `json:"updated_date"`
	CompanyID    ID         `json:"company_id,string"`
}
