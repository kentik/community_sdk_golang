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

// V202101beta1TestPingSettings struct for V202101beta1TestPingSettings
type V202101beta1TestPingSettings struct {
	Period *float32 `json:"period,omitempty"`
	Count  *float32 `json:"count,omitempty"`
	Expiry *float32 `json:"expiry,omitempty"`
	Delay  *float32 `json:"delay,omitempty"`
}

// NewV202101beta1TestPingSettings instantiates a new V202101beta1TestPingSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV202101beta1TestPingSettings() *V202101beta1TestPingSettings {
	this := V202101beta1TestPingSettings{}
	return &this
}

// NewV202101beta1TestPingSettingsWithDefaults instantiates a new V202101beta1TestPingSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV202101beta1TestPingSettingsWithDefaults() *V202101beta1TestPingSettings {
	this := V202101beta1TestPingSettings{}
	return &this
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *V202101beta1TestPingSettings) GetPeriod() float32 {
	if o == nil || o.Period == nil {
		var ret float32
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1TestPingSettings) GetPeriodOk() (*float32, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *V202101beta1TestPingSettings) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given float32 and assigns it to the Period field.
func (o *V202101beta1TestPingSettings) SetPeriod(v float32) {
	o.Period = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *V202101beta1TestPingSettings) GetCount() float32 {
	if o == nil || o.Count == nil {
		var ret float32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1TestPingSettings) GetCountOk() (*float32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *V202101beta1TestPingSettings) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given float32 and assigns it to the Count field.
func (o *V202101beta1TestPingSettings) SetCount(v float32) {
	o.Count = &v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *V202101beta1TestPingSettings) GetExpiry() float32 {
	if o == nil || o.Expiry == nil {
		var ret float32
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1TestPingSettings) GetExpiryOk() (*float32, bool) {
	if o == nil || o.Expiry == nil {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *V202101beta1TestPingSettings) HasExpiry() bool {
	if o != nil && o.Expiry != nil {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given float32 and assigns it to the Expiry field.
func (o *V202101beta1TestPingSettings) SetExpiry(v float32) {
	o.Expiry = &v
}

// GetDelay returns the Delay field value if set, zero value otherwise.
func (o *V202101beta1TestPingSettings) GetDelay() float32 {
	if o == nil || o.Delay == nil {
		var ret float32
		return ret
	}
	return *o.Delay
}

// GetDelayOk returns a tuple with the Delay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1TestPingSettings) GetDelayOk() (*float32, bool) {
	if o == nil || o.Delay == nil {
		return nil, false
	}
	return o.Delay, true
}

// HasDelay returns a boolean if a field has been set.
func (o *V202101beta1TestPingSettings) HasDelay() bool {
	if o != nil && o.Delay != nil {
		return true
	}

	return false
}

// SetDelay gets a reference to the given float32 and assigns it to the Delay field.
func (o *V202101beta1TestPingSettings) SetDelay(v float32) {
	o.Delay = &v
}

func (o V202101beta1TestPingSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Period != nil {
		toSerialize["period"] = o.Period
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Expiry != nil {
		toSerialize["expiry"] = o.Expiry
	}
	if o.Delay != nil {
		toSerialize["delay"] = o.Delay
	}
	return json.Marshal(toSerialize)
}

type NullableV202101beta1TestPingSettings struct {
	value *V202101beta1TestPingSettings
	isSet bool
}

func (v NullableV202101beta1TestPingSettings) Get() *V202101beta1TestPingSettings {
	return v.value
}

func (v *NullableV202101beta1TestPingSettings) Set(val *V202101beta1TestPingSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableV202101beta1TestPingSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableV202101beta1TestPingSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV202101beta1TestPingSettings(val *V202101beta1TestPingSettings) *NullableV202101beta1TestPingSettings {
	return &NullableV202101beta1TestPingSettings{value: val, isSet: true}
}

func (v NullableV202101beta1TestPingSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV202101beta1TestPingSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
