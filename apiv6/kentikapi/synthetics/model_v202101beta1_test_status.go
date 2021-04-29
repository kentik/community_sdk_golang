/*
 * Synthetics Monitoring API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 202101beta1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synthetics

import (
	"encoding/json"
	"fmt"
)

// V202101beta1TestStatus the model 'V202101beta1TestStatus'
type V202101beta1TestStatus string

// List of v202101beta1TestStatus
const (
	V202101BETA1TESTSTATUS_UNSPECIFIED V202101beta1TestStatus = "TEST_STATUS_UNSPECIFIED"
	V202101BETA1TESTSTATUS_ACTIVE      V202101beta1TestStatus = "TEST_STATUS_ACTIVE"
	V202101BETA1TESTSTATUS_PAUSED      V202101beta1TestStatus = "TEST_STATUS_PAUSED"
	V202101BETA1TESTSTATUS_DELETED     V202101beta1TestStatus = "TEST_STATUS_DELETED"
)

func (v *V202101beta1TestStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V202101beta1TestStatus(value)
	for _, existing := range []V202101beta1TestStatus{"TEST_STATUS_UNSPECIFIED", "TEST_STATUS_ACTIVE", "TEST_STATUS_PAUSED", "TEST_STATUS_DELETED"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V202101beta1TestStatus", value)
}

// Ptr returns reference to v202101beta1TestStatus value
func (v V202101beta1TestStatus) Ptr() *V202101beta1TestStatus {
	return &v
}

type NullableV202101beta1TestStatus struct {
	value *V202101beta1TestStatus
	isSet bool
}

func (v NullableV202101beta1TestStatus) Get() *V202101beta1TestStatus {
	return v.value
}

func (v *NullableV202101beta1TestStatus) Set(val *V202101beta1TestStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableV202101beta1TestStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableV202101beta1TestStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV202101beta1TestStatus(val *V202101beta1TestStatus) *NullableV202101beta1TestStatus {
	return &NullableV202101beta1TestStatus{value: val, isSet: true}
}

func (v NullableV202101beta1TestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV202101beta1TestStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}