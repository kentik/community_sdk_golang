package endpoints

import (
	"fmt"
)

const (
	SavedFilterPath  = "/saved-filter/custom"
	SavedFiltersPath = "/saved-filters/custom"
)

func GetSavedFilter(id ResourceID) string {
	return fmt.Sprintf("%v/%v", SavedFilterPath, id)
}
