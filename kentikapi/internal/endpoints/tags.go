package endpoints

import (
	"fmt"
)

const (
	TagPath  = "/tag"
	TagsPath = "/tags"
)

func GetTagPath(id ResourceID) string {
	return fmt.Sprintf("%v/%v", TagPath, id)
}
