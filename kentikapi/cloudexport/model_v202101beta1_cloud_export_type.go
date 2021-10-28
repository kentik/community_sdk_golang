/*
 * Cloud Export Admin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 202101beta1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudexport

import (
	"encoding/json"
	"fmt"
)

// V202101beta1CloudExportType the model 'V202101beta1CloudExportType'
type V202101beta1CloudExportType string

// List of v202101beta1CloudExportType
const (
	V202101BETA1CLOUDEXPORTTYPE_UNSPECIFIED      V202101beta1CloudExportType = "CLOUD_EXPORT_TYPE_UNSPECIFIED"
	V202101BETA1CLOUDEXPORTTYPE_KENTIK_MANAGED   V202101beta1CloudExportType = "CLOUD_EXPORT_TYPE_KENTIK_MANAGED"
	V202101BETA1CLOUDEXPORTTYPE_CUSTOMER_MANAGED V202101beta1CloudExportType = "CLOUD_EXPORT_TYPE_CUSTOMER_MANAGED"
)

var allowedV202101beta1CloudExportTypeEnumValues = []V202101beta1CloudExportType{
	"CLOUD_EXPORT_TYPE_UNSPECIFIED",
	"CLOUD_EXPORT_TYPE_KENTIK_MANAGED",
	"CLOUD_EXPORT_TYPE_CUSTOMER_MANAGED",
}

func (v *V202101beta1CloudExportType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V202101beta1CloudExportType(value)
	for _, existing := range allowedV202101beta1CloudExportTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V202101beta1CloudExportType", value)
}

// NewV202101beta1CloudExportTypeFromValue returns a pointer to a valid V202101beta1CloudExportType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV202101beta1CloudExportTypeFromValue(v string) (*V202101beta1CloudExportType, error) {
	ev := V202101beta1CloudExportType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V202101beta1CloudExportType: valid values are %v", v, allowedV202101beta1CloudExportTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V202101beta1CloudExportType) IsValid() bool {
	for _, existing := range allowedV202101beta1CloudExportTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v202101beta1CloudExportType value
func (v V202101beta1CloudExportType) Ptr() *V202101beta1CloudExportType {
	return &v
}

type NullableV202101beta1CloudExportType struct {
	value *V202101beta1CloudExportType
	isSet bool
}

func (v NullableV202101beta1CloudExportType) Get() *V202101beta1CloudExportType {
	return v.value
}

func (v *NullableV202101beta1CloudExportType) Set(val *V202101beta1CloudExportType) {
	v.value = val
	v.isSet = true
}

func (v NullableV202101beta1CloudExportType) IsSet() bool {
	return v.isSet
}

func (v *NullableV202101beta1CloudExportType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV202101beta1CloudExportType(val *V202101beta1CloudExportType) *NullableV202101beta1CloudExportType {
	return &NullableV202101beta1CloudExportType{value: val, isSet: true}
}

func (v NullableV202101beta1CloudExportType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV202101beta1CloudExportType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}