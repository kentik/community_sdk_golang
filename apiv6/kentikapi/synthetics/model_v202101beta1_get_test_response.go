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
)

// V202101beta1GetTestResponse struct for V202101beta1GetTestResponse
type V202101beta1GetTestResponse struct {
	Test *V202101beta1Test `json:"test,omitempty"`
}

// NewV202101beta1GetTestResponse instantiates a new V202101beta1GetTestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV202101beta1GetTestResponse() *V202101beta1GetTestResponse {
	this := V202101beta1GetTestResponse{}
	return &this
}

// NewV202101beta1GetTestResponseWithDefaults instantiates a new V202101beta1GetTestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV202101beta1GetTestResponseWithDefaults() *V202101beta1GetTestResponse {
	this := V202101beta1GetTestResponse{}
	return &this
}

// GetTest returns the Test field value if set, zero value otherwise.
func (o *V202101beta1GetTestResponse) GetTest() V202101beta1Test {
	if o == nil || o.Test == nil {
		var ret V202101beta1Test
		return ret
	}
	return *o.Test
}

// GetTestOk returns a tuple with the Test field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1GetTestResponse) GetTestOk() (*V202101beta1Test, bool) {
	if o == nil || o.Test == nil {
		return nil, false
	}
	return o.Test, true
}

// HasTest returns a boolean if a field has been set.
func (o *V202101beta1GetTestResponse) HasTest() bool {
	if o != nil && o.Test != nil {
		return true
	}

	return false
}

// SetTest gets a reference to the given V202101beta1Test and assigns it to the Test field.
func (o *V202101beta1GetTestResponse) SetTest(v V202101beta1Test) {
	o.Test = &v
}

func (o V202101beta1GetTestResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Test != nil {
		toSerialize["test"] = o.Test
	}
	return json.Marshal(toSerialize)
}

type NullableV202101beta1GetTestResponse struct {
	value *V202101beta1GetTestResponse
	isSet bool
}

func (v NullableV202101beta1GetTestResponse) Get() *V202101beta1GetTestResponse {
	return v.value
}

func (v *NullableV202101beta1GetTestResponse) Set(val *V202101beta1GetTestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV202101beta1GetTestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV202101beta1GetTestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV202101beta1GetTestResponse(val *V202101beta1GetTestResponse) *NullableV202101beta1GetTestResponse {
	return &NullableV202101beta1GetTestResponse{value: val, isSet: true}
}

func (v NullableV202101beta1GetTestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV202101beta1GetTestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}