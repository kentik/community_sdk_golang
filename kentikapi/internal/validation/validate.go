package validation

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type Direction string

const (
	DirectionRequest  Direction = "request"
	DirectionResponse Direction = "response"
)

// CheckRequestRequiredFields checks if resource's required fields are not nil.
// Eg. for method=post, it returns error if any of resource's fields marked as `request:"post"`` are set to nil.
func CheckRequestRequiredFields(method string, resource interface{}) error {
	lowercaseMethod := strings.ToLower(method)
	return validateWrapper(lowercaseMethod, DirectionRequest, resource)
}

// CheckResponseRequiredFields checks if resource's required fields are not nil.
// Eg. for method=get, it returns error if any of resource's fields marked as `response:"get"`` are set to nil.
func CheckResponseRequiredFields(method string, resource interface{}) error {
	lowercaseMethod := strings.ToLower(method)
	return validateWrapper(lowercaseMethod, DirectionResponse, resource)
}

// validateWrapper returns error containing the list of required fields that happen to be nil.
func validateWrapper(method string, direction Direction, resource interface{}) error {
	missing := validate(method, string(direction), getTypeName(resource), reflect.ValueOf(resource))
	if len(missing) > 0 {
		return fmt.Errorf("following fields are missing in %s %s: %v", method, direction, missing)
	}
	return nil
}

func getTypeName(i interface{}) string {
	tResource := reflect.TypeOf(i)
	if tResource.Kind() == reflect.Ptr {
		return "*" + tResource.Elem().Name()
	}
	return tResource.Name()
}

//nolint:gocyclo
func validate(method string, direction string, path string, v reflect.Value) []string {
	missing := make([]string, 0)

	switch v.Kind() {
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fieldPath := path + "." + v.Type().Field(i).Name
			if (field.Kind() == reflect.Ptr || field.Kind() == reflect.Interface || field.Kind() == reflect.Slice) &&
				field.IsNil() {
				requiredForMethods := v.Type().Field(i).Tag.Get(direction)
				if strings.Contains(requiredForMethods, method) {
					missing = append(missing, fieldPath)
				}
			} else {
				missing = append(missing, validate(method, direction, fieldPath, field)...)
			}
		}

	case reflect.Slice:
		// slice can hold structs, so - validate
		count := v.Len()
		for i := 0; i < count; i++ {
			item := v.Index(i)
			fieldPath := fmt.Sprintf("%s[%d]", path, i)
			missing = append(missing, validate(method, direction, fieldPath, item)...)
		}

	case reflect.Ptr, reflect.Interface:
		// pointer and interface can hold a struct, so - validate
		// path = path + "." + v.Name
		missing = append(missing, validate(method, direction, path, v.Elem())...)

	default:
		// primitive value has no field tags so nothing to validate
	}
	return missing
}

// ValidateCECreateRequest checks if CloudExport create request contains all required fields.
//nolint:gocyclo
func ValidateCECreateRequest(ce *models.CloudExport) error {
	if ce == nil {
		return nil
	}
	if ce.Name == "" {
		return ceFieldError("Name")
	}
	if ce.PlanID == "" {
		return ceFieldError("PlanID")
	}
	switch ce.CloudProvider {
	case "":
		{
			return ceFieldError("CloudProvider")
		}
	case "aws":
		{
			if ce.GetAWSProperties().Bucket == "" {
				return ceFieldError("Properties.Bucket")
			}
		}
	case "azure":
		{
			if ce.GetAzureProperties().Location == "" {
				return ceFieldError("Properties.Location")
			}
			if ce.GetAzureProperties().ResourceGroup == "" {
				return ceFieldError("Properties.ResourceGroup")
			}
			if ce.GetAzureProperties().StorageAccount == "" {
				return ceFieldError("Properties.StorageAccount")
			}
			if ce.GetAzureProperties().SubscriptionID == "" {
				return ceFieldError("Properties.SubscriptionID")
			}
		}
	case "gce":
		{
			if ce.GetGCEProperties().Project == "" {
				return ceFieldError("Properties.Project")
			}
			if ce.GetGCEProperties().Subscription == "" {
				return ceFieldError("Properties.Subscription")
			}
		}
	case "ibm":
		{
			if ce.GetIBMProperties().Bucket == "" {
				return ceFieldError("Properties.Bucket")
			}
		}
	}
	return nil
}

// ValidateCEUpdateRequest checks if CloudExport update request contains all required fields.
func ValidateCEUpdateRequest(ce *models.CloudExport) error {
	if ce == nil {
		return nil
	}
	if ce.ID == "" {
		return ceFieldError("ID")
	}
	if ce.CloudProvider == "" {
		return ceFieldError("CloudProvider")
	}
	if ce.Properties == nil {
		return ceFieldError("Properties")
	}
	return nil
}

func ceFieldError(field string) error {
	return fmt.Errorf("CloudExport '%s' field is required but not provided", field)
}
