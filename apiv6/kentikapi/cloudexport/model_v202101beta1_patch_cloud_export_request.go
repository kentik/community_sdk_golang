/*
 * Cloud Export Admin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 202101
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudexport

import (
	"encoding/json"
)

// V202101beta1PatchCloudExportRequest struct for V202101beta1PatchCloudExportRequest
type V202101beta1PatchCloudExportRequest struct {
	Export     *V202101beta1CloudExport `json:"export,omitempty"`
	UpdateMask *string                  `json:"updateMask,omitempty"`
}

// NewV202101beta1PatchCloudExportRequest instantiates a new V202101beta1PatchCloudExportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV202101beta1PatchCloudExportRequest() *V202101beta1PatchCloudExportRequest {
	this := V202101beta1PatchCloudExportRequest{}
	return &this
}

// NewV202101beta1PatchCloudExportRequestWithDefaults instantiates a new V202101beta1PatchCloudExportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV202101beta1PatchCloudExportRequestWithDefaults() *V202101beta1PatchCloudExportRequest {
	this := V202101beta1PatchCloudExportRequest{}
	return &this
}

// GetExport returns the Export field value if set, zero value otherwise.
func (o *V202101beta1PatchCloudExportRequest) GetExport() V202101beta1CloudExport {
	if o == nil || o.Export == nil {
		var ret V202101beta1CloudExport
		return ret
	}
	return *o.Export
}

// GetExportOk returns a tuple with the Export field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1PatchCloudExportRequest) GetExportOk() (*V202101beta1CloudExport, bool) {
	if o == nil || o.Export == nil {
		return nil, false
	}
	return o.Export, true
}

// HasExport returns a boolean if a field has been set.
func (o *V202101beta1PatchCloudExportRequest) HasExport() bool {
	if o != nil && o.Export != nil {
		return true
	}

	return false
}

// SetExport gets a reference to the given V202101beta1CloudExport and assigns it to the Export field.
func (o *V202101beta1PatchCloudExportRequest) SetExport(v V202101beta1CloudExport) {
	o.Export = &v
}

// GetUpdateMask returns the UpdateMask field value if set, zero value otherwise.
func (o *V202101beta1PatchCloudExportRequest) GetUpdateMask() string {
	if o == nil || o.UpdateMask == nil {
		var ret string
		return ret
	}
	return *o.UpdateMask
}

// GetUpdateMaskOk returns a tuple with the UpdateMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1PatchCloudExportRequest) GetUpdateMaskOk() (*string, bool) {
	if o == nil || o.UpdateMask == nil {
		return nil, false
	}
	return o.UpdateMask, true
}

// HasUpdateMask returns a boolean if a field has been set.
func (o *V202101beta1PatchCloudExportRequest) HasUpdateMask() bool {
	if o != nil && o.UpdateMask != nil {
		return true
	}

	return false
}

// SetUpdateMask gets a reference to the given string and assigns it to the UpdateMask field.
func (o *V202101beta1PatchCloudExportRequest) SetUpdateMask(v string) {
	o.UpdateMask = &v
}

func (o V202101beta1PatchCloudExportRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Export != nil {
		toSerialize["export"] = o.Export
	}
	if o.UpdateMask != nil {
		toSerialize["updateMask"] = o.UpdateMask
	}
	return json.Marshal(toSerialize)
}

type NullableV202101beta1PatchCloudExportRequest struct {
	value *V202101beta1PatchCloudExportRequest
	isSet bool
}

func (v NullableV202101beta1PatchCloudExportRequest) Get() *V202101beta1PatchCloudExportRequest {
	return v.value
}

func (v *NullableV202101beta1PatchCloudExportRequest) Set(val *V202101beta1PatchCloudExportRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV202101beta1PatchCloudExportRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV202101beta1PatchCloudExportRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV202101beta1PatchCloudExportRequest(val *V202101beta1PatchCloudExportRequest) *NullableV202101beta1PatchCloudExportRequest {
	return &NullableV202101beta1PatchCloudExportRequest{value: val, isSet: true}
}

func (v NullableV202101beta1PatchCloudExportRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV202101beta1PatchCloudExportRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
