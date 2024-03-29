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

// V202101beta1AgentStatus the model 'V202101beta1AgentStatus'
type V202101beta1AgentStatus string

// List of v202101beta1AgentStatus
const (
	V202101BETA1AGENTSTATUS_UNSPECIFIED V202101beta1AgentStatus = "AGENT_STATUS_UNSPECIFIED"
	V202101BETA1AGENTSTATUS_OK          V202101beta1AgentStatus = "AGENT_STATUS_OK"
	V202101BETA1AGENTSTATUS_WAIT        V202101beta1AgentStatus = "AGENT_STATUS_WAIT"
	V202101BETA1AGENTSTATUS_DELETED     V202101beta1AgentStatus = "AGENT_STATUS_DELETED"
)

var allowedV202101beta1AgentStatusEnumValues = []V202101beta1AgentStatus{
	"AGENT_STATUS_UNSPECIFIED",
	"AGENT_STATUS_OK",
	"AGENT_STATUS_WAIT",
	"AGENT_STATUS_DELETED",
}

func (v *V202101beta1AgentStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V202101beta1AgentStatus(value)
	for _, existing := range allowedV202101beta1AgentStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V202101beta1AgentStatus", value)
}

// NewV202101beta1AgentStatusFromValue returns a pointer to a valid V202101beta1AgentStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV202101beta1AgentStatusFromValue(v string) (*V202101beta1AgentStatus, error) {
	ev := V202101beta1AgentStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V202101beta1AgentStatus: valid values are %v", v, allowedV202101beta1AgentStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V202101beta1AgentStatus) IsValid() bool {
	for _, existing := range allowedV202101beta1AgentStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v202101beta1AgentStatus value
func (v V202101beta1AgentStatus) Ptr() *V202101beta1AgentStatus {
	return &v
}

type NullableV202101beta1AgentStatus struct {
	value *V202101beta1AgentStatus
	isSet bool
}

func (v NullableV202101beta1AgentStatus) Get() *V202101beta1AgentStatus {
	return v.value
}

func (v *NullableV202101beta1AgentStatus) Set(val *V202101beta1AgentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableV202101beta1AgentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableV202101beta1AgentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV202101beta1AgentStatus(val *V202101beta1AgentStatus) *NullableV202101beta1AgentStatus {
	return &NullableV202101beta1AgentStatus{value: val, isSet: true}
}

func (v NullableV202101beta1AgentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV202101beta1AgentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
