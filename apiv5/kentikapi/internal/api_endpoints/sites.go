package api_endpoints

import "fmt"

func GetAllSites() string {
	return "/sites"
}

func GetSite(id ResourceID) string {
	return fmt.Sprintf("/site/%v", id)
}

func CreateSite() string {
	return "/site"
}

func UpdateSite(id ResourceID) string {
	return fmt.Sprintf("/site/%v", id)
}

func DeleteSite(id ResourceID) string {
	return fmt.Sprintf("/site/%v", id)
}
