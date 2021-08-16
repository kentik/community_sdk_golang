package api_endpoints

import (
	"fmt"
)

func GetAllLabels() string {
	return "/deviceLabels"
}

func GetLabel(id ResourceID) string {
	return fmt.Sprintf("/deviceLabels/%v", id)
}

func CreateLabel() string {
	return "/deviceLabels"
}

func UpdateLabel(id ResourceID) string {
	return fmt.Sprintf("/deviceLabels/%v", id)
}

func DeleteLabel(id ResourceID) string {
	return fmt.Sprintf("/deviceLabels/%v", id)
}
