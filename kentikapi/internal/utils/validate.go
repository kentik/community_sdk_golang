package utils

import (
	"fmt"
	"reflect"
	"strings"
)

type Direction string

const DirectionRequest Direction = "request"
const DirectionResponse Direction = "response"

func ValidateRequest(method string, resource interface{}) error {
	lowercaseMethod := strings.ToLower(method)
	return validateWrapper(lowercaseMethod, DirectionRequest, resource)
}

func ValidateResponse(method string, resource interface{}) error {
	lowercaseMethod := strings.ToLower(method)
	return validateWrapper(lowercaseMethod, DirectionResponse, resource)
}

func validateWrapper(method string, direction Direction, resource interface{}) error {
	missing := validate(method, string(direction), "", reflect.ValueOf(resource))
	if len(missing) > 0 {
		return fmt.Errorf("following fields are missing in %s %s: %v", method, direction, missing)
	}
	return nil
}

func validate(method string, direction string, path string, v reflect.Value) []string {
	missing := make([]string, 0)

	switch v.Kind() {

	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fieldPath := path + "." + v.Type().Field(i).Name
			if (field.Kind() == reflect.Ptr || field.Kind() == reflect.Interface) && field.IsNil() {
				requiredForMethods := v.Type().Field(i).Tag.Get(direction)
				if strings.Contains(requiredForMethods, method) {
					missing = append(missing, fieldPath)
				}
			} else {
				missing = append(missing, validate(method, direction, fieldPath+field.Type().Name(), field)...)
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
		missing = append(missing, validate(method, direction, path, v.Elem())...)

	default:
		// primitive value has no field tags so nothing to validate

	}
	return missing
}
