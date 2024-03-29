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

// V202101beta1DnsGridTest struct for V202101beta1DnsGridTest
type V202101beta1DnsGridTest struct {
	Targets *[]string              `json:"targets,omitempty"`
	Type    *V202101beta1DNSRecord `json:"type,omitempty"`
}

// NewV202101beta1DnsGridTest instantiates a new V202101beta1DnsGridTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV202101beta1DnsGridTest() *V202101beta1DnsGridTest {
	this := V202101beta1DnsGridTest{}
	var type_ V202101beta1DNSRecord = V202101BETA1DNSRECORD_UNSPECIFIED
	this.Type = &type_
	return &this
}

// NewV202101beta1DnsGridTestWithDefaults instantiates a new V202101beta1DnsGridTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV202101beta1DnsGridTestWithDefaults() *V202101beta1DnsGridTest {
	this := V202101beta1DnsGridTest{}
	var type_ V202101beta1DNSRecord = V202101BETA1DNSRECORD_UNSPECIFIED
	this.Type = &type_
	return &this
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *V202101beta1DnsGridTest) GetTargets() []string {
	if o == nil || o.Targets == nil {
		var ret []string
		return ret
	}
	return *o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1DnsGridTest) GetTargetsOk() (*[]string, bool) {
	if o == nil || o.Targets == nil {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *V202101beta1DnsGridTest) HasTargets() bool {
	if o != nil && o.Targets != nil {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []string and assigns it to the Targets field.
func (o *V202101beta1DnsGridTest) SetTargets(v []string) {
	o.Targets = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V202101beta1DnsGridTest) GetType() V202101beta1DNSRecord {
	if o == nil || o.Type == nil {
		var ret V202101beta1DNSRecord
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1DnsGridTest) GetTypeOk() (*V202101beta1DNSRecord, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V202101beta1DnsGridTest) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given V202101beta1DNSRecord and assigns it to the Type field.
func (o *V202101beta1DnsGridTest) SetType(v V202101beta1DNSRecord) {
	o.Type = &v
}

func (o V202101beta1DnsGridTest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Targets != nil {
		toSerialize["targets"] = o.Targets
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableV202101beta1DnsGridTest struct {
	value *V202101beta1DnsGridTest
	isSet bool
}

func (v NullableV202101beta1DnsGridTest) Get() *V202101beta1DnsGridTest {
	return v.value
}

func (v *NullableV202101beta1DnsGridTest) Set(val *V202101beta1DnsGridTest) {
	v.value = val
	v.isSet = true
}

func (v NullableV202101beta1DnsGridTest) IsSet() bool {
	return v.isSet
}

func (v *NullableV202101beta1DnsGridTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV202101beta1DnsGridTest(val *V202101beta1DnsGridTest) *NullableV202101beta1DnsGridTest {
	return &NullableV202101beta1DnsGridTest{value: val, isSet: true}
}

func (v NullableV202101beta1DnsGridTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV202101beta1DnsGridTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
