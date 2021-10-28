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
)

// GooglerpcStatus struct for GooglerpcStatus
type GooglerpcStatus struct {
	Code    *int32         `json:"code,omitempty"`
	Message *string        `json:"message,omitempty"`
	Details *[]ProtobufAny `json:"details,omitempty"`
}

// NewGooglerpcStatus instantiates a new GooglerpcStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGooglerpcStatus() *GooglerpcStatus {
	this := GooglerpcStatus{}
	return &this
}

// NewGooglerpcStatusWithDefaults instantiates a new GooglerpcStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGooglerpcStatusWithDefaults() *GooglerpcStatus {
	this := GooglerpcStatus{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GooglerpcStatus) GetCode() int32 {
	if o == nil || o.Code == nil {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GooglerpcStatus) GetCodeOk() (*int32, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GooglerpcStatus) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *GooglerpcStatus) SetCode(v int32) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GooglerpcStatus) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GooglerpcStatus) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GooglerpcStatus) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GooglerpcStatus) SetMessage(v string) {
	o.Message = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *GooglerpcStatus) GetDetails() []ProtobufAny {
	if o == nil || o.Details == nil {
		var ret []ProtobufAny
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GooglerpcStatus) GetDetailsOk() (*[]ProtobufAny, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *GooglerpcStatus) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []ProtobufAny and assigns it to the Details field.
func (o *GooglerpcStatus) SetDetails(v []ProtobufAny) {
	o.Details = &v
}

func (o GooglerpcStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	return json.Marshal(toSerialize)
}

type NullableGooglerpcStatus struct {
	value *GooglerpcStatus
	isSet bool
}

func (v NullableGooglerpcStatus) Get() *GooglerpcStatus {
	return v.value
}

func (v *NullableGooglerpcStatus) Set(val *GooglerpcStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableGooglerpcStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableGooglerpcStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGooglerpcStatus(val *GooglerpcStatus) *NullableGooglerpcStatus {
	return &NullableGooglerpcStatus{value: val, isSet: true}
}

func (v NullableGooglerpcStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGooglerpcStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}