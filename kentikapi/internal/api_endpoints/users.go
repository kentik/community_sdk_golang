package api_endpoints

import (
	"fmt"
)

const (
	UserPath  = "/user"
	UsersPath = "/users"
)

func GetUser(id ResourceID) string {
	return fmt.Sprintf("%v/%v", UserPath, id)
}
