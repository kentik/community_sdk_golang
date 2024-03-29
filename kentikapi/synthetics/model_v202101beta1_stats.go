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

// V202101beta1Stats struct for V202101beta1Stats
type V202101beta1Stats struct {
	Average *int32 `json:"average,omitempty"`
	Max     *int32 `json:"max,omitempty"`
	Total   *int32 `json:"total,omitempty"`
}

// NewV202101beta1Stats instantiates a new V202101beta1Stats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV202101beta1Stats() *V202101beta1Stats {
	this := V202101beta1Stats{}
	return &this
}

// NewV202101beta1StatsWithDefaults instantiates a new V202101beta1Stats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV202101beta1StatsWithDefaults() *V202101beta1Stats {
	this := V202101beta1Stats{}
	return &this
}

// GetAverage returns the Average field value if set, zero value otherwise.
func (o *V202101beta1Stats) GetAverage() int32 {
	if o == nil || o.Average == nil {
		var ret int32
		return ret
	}
	return *o.Average
}

// GetAverageOk returns a tuple with the Average field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1Stats) GetAverageOk() (*int32, bool) {
	if o == nil || o.Average == nil {
		return nil, false
	}
	return o.Average, true
}

// HasAverage returns a boolean if a field has been set.
func (o *V202101beta1Stats) HasAverage() bool {
	if o != nil && o.Average != nil {
		return true
	}

	return false
}

// SetAverage gets a reference to the given int32 and assigns it to the Average field.
func (o *V202101beta1Stats) SetAverage(v int32) {
	o.Average = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *V202101beta1Stats) GetMax() int32 {
	if o == nil || o.Max == nil {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1Stats) GetMaxOk() (*int32, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *V202101beta1Stats) HasMax() bool {
	if o != nil && o.Max != nil {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *V202101beta1Stats) SetMax(v int32) {
	o.Max = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *V202101beta1Stats) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1Stats) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *V202101beta1Stats) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *V202101beta1Stats) SetTotal(v int32) {
	o.Total = &v
}

func (o V202101beta1Stats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Average != nil {
		toSerialize["average"] = o.Average
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableV202101beta1Stats struct {
	value *V202101beta1Stats
	isSet bool
}

func (v NullableV202101beta1Stats) Get() *V202101beta1Stats {
	return v.value
}

func (v *NullableV202101beta1Stats) Set(val *V202101beta1Stats) {
	v.value = val
	v.isSet = true
}

func (v NullableV202101beta1Stats) IsSet() bool {
	return v.isSet
}

func (v *NullableV202101beta1Stats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV202101beta1Stats(val *V202101beta1Stats) *NullableV202101beta1Stats {
	return &NullableV202101beta1Stats{value: val, isSet: true}
}

func (v NullableV202101beta1Stats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV202101beta1Stats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
