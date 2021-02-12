package kentikapi

import (
	"fmt"

	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

const (
	userPath  = "/user"
	usersPath = "/users"
)

func getUserPath(id models.ID) string {
	return fmt.Sprintf("%v/%v", userPath, id)
}
