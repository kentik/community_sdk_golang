package endpoints

import "fmt"

func GetAllCustomApplications() string {
	return "/customApplications"
}

func CreateCustomApplication() string {
	return "/customApplications"
}

func UpdateCustomApplication(id ResourceID) string {
	return fmt.Sprintf("/customApplications/%v", id)
}

func DeleteCustomApplication(id ResourceID) string {
	return fmt.Sprintf("/customApplications/%v", id)
}
