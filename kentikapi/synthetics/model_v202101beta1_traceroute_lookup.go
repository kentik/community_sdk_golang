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

// V202101beta1TracerouteLookup struct for V202101beta1TracerouteLookup
type V202101beta1TracerouteLookup struct {
	AgentIdByIp  *[]V202101beta1IDByIP `json:"agentIdByIp,omitempty"`
	Agents       *[]V202101beta1Agent  `json:"agents,omitempty"`
	Asns         *[]V202101beta1ASN    `json:"asns,omitempty"`
	DeviceIdByIp *[]V202101beta1IDByIP `json:"deviceIdByIp,omitempty"`
	SiteIdByIp   *[]V202101beta1IDByIP `json:"siteIdByIp,omitempty"`
	Ips          *[]V202101beta1IPInfo `json:"ips,omitempty"`
}

// NewV202101beta1TracerouteLookup instantiates a new V202101beta1TracerouteLookup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV202101beta1TracerouteLookup() *V202101beta1TracerouteLookup {
	this := V202101beta1TracerouteLookup{}
	return &this
}

// NewV202101beta1TracerouteLookupWithDefaults instantiates a new V202101beta1TracerouteLookup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV202101beta1TracerouteLookupWithDefaults() *V202101beta1TracerouteLookup {
	this := V202101beta1TracerouteLookup{}
	return &this
}

// GetAgentIdByIp returns the AgentIdByIp field value if set, zero value otherwise.
func (o *V202101beta1TracerouteLookup) GetAgentIdByIp() []V202101beta1IDByIP {
	if o == nil || o.AgentIdByIp == nil {
		var ret []V202101beta1IDByIP
		return ret
	}
	return *o.AgentIdByIp
}

// GetAgentIdByIpOk returns a tuple with the AgentIdByIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1TracerouteLookup) GetAgentIdByIpOk() (*[]V202101beta1IDByIP, bool) {
	if o == nil || o.AgentIdByIp == nil {
		return nil, false
	}
	return o.AgentIdByIp, true
}

// HasAgentIdByIp returns a boolean if a field has been set.
func (o *V202101beta1TracerouteLookup) HasAgentIdByIp() bool {
	if o != nil && o.AgentIdByIp != nil {
		return true
	}

	return false
}

// SetAgentIdByIp gets a reference to the given []V202101beta1IDByIP and assigns it to the AgentIdByIp field.
func (o *V202101beta1TracerouteLookup) SetAgentIdByIp(v []V202101beta1IDByIP) {
	o.AgentIdByIp = &v
}

// GetAgents returns the Agents field value if set, zero value otherwise.
func (o *V202101beta1TracerouteLookup) GetAgents() []V202101beta1Agent {
	if o == nil || o.Agents == nil {
		var ret []V202101beta1Agent
		return ret
	}
	return *o.Agents
}

// GetAgentsOk returns a tuple with the Agents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1TracerouteLookup) GetAgentsOk() (*[]V202101beta1Agent, bool) {
	if o == nil || o.Agents == nil {
		return nil, false
	}
	return o.Agents, true
}

// HasAgents returns a boolean if a field has been set.
func (o *V202101beta1TracerouteLookup) HasAgents() bool {
	if o != nil && o.Agents != nil {
		return true
	}

	return false
}

// SetAgents gets a reference to the given []V202101beta1Agent and assigns it to the Agents field.
func (o *V202101beta1TracerouteLookup) SetAgents(v []V202101beta1Agent) {
	o.Agents = &v
}

// GetAsns returns the Asns field value if set, zero value otherwise.
func (o *V202101beta1TracerouteLookup) GetAsns() []V202101beta1ASN {
	if o == nil || o.Asns == nil {
		var ret []V202101beta1ASN
		return ret
	}
	return *o.Asns
}

// GetAsnsOk returns a tuple with the Asns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1TracerouteLookup) GetAsnsOk() (*[]V202101beta1ASN, bool) {
	if o == nil || o.Asns == nil {
		return nil, false
	}
	return o.Asns, true
}

// HasAsns returns a boolean if a field has been set.
func (o *V202101beta1TracerouteLookup) HasAsns() bool {
	if o != nil && o.Asns != nil {
		return true
	}

	return false
}

// SetAsns gets a reference to the given []V202101beta1ASN and assigns it to the Asns field.
func (o *V202101beta1TracerouteLookup) SetAsns(v []V202101beta1ASN) {
	o.Asns = &v
}

// GetDeviceIdByIp returns the DeviceIdByIp field value if set, zero value otherwise.
func (o *V202101beta1TracerouteLookup) GetDeviceIdByIp() []V202101beta1IDByIP {
	if o == nil || o.DeviceIdByIp == nil {
		var ret []V202101beta1IDByIP
		return ret
	}
	return *o.DeviceIdByIp
}

// GetDeviceIdByIpOk returns a tuple with the DeviceIdByIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1TracerouteLookup) GetDeviceIdByIpOk() (*[]V202101beta1IDByIP, bool) {
	if o == nil || o.DeviceIdByIp == nil {
		return nil, false
	}
	return o.DeviceIdByIp, true
}

// HasDeviceIdByIp returns a boolean if a field has been set.
func (o *V202101beta1TracerouteLookup) HasDeviceIdByIp() bool {
	if o != nil && o.DeviceIdByIp != nil {
		return true
	}

	return false
}

// SetDeviceIdByIp gets a reference to the given []V202101beta1IDByIP and assigns it to the DeviceIdByIp field.
func (o *V202101beta1TracerouteLookup) SetDeviceIdByIp(v []V202101beta1IDByIP) {
	o.DeviceIdByIp = &v
}

// GetSiteIdByIp returns the SiteIdByIp field value if set, zero value otherwise.
func (o *V202101beta1TracerouteLookup) GetSiteIdByIp() []V202101beta1IDByIP {
	if o == nil || o.SiteIdByIp == nil {
		var ret []V202101beta1IDByIP
		return ret
	}
	return *o.SiteIdByIp
}

// GetSiteIdByIpOk returns a tuple with the SiteIdByIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1TracerouteLookup) GetSiteIdByIpOk() (*[]V202101beta1IDByIP, bool) {
	if o == nil || o.SiteIdByIp == nil {
		return nil, false
	}
	return o.SiteIdByIp, true
}

// HasSiteIdByIp returns a boolean if a field has been set.
func (o *V202101beta1TracerouteLookup) HasSiteIdByIp() bool {
	if o != nil && o.SiteIdByIp != nil {
		return true
	}

	return false
}

// SetSiteIdByIp gets a reference to the given []V202101beta1IDByIP and assigns it to the SiteIdByIp field.
func (o *V202101beta1TracerouteLookup) SetSiteIdByIp(v []V202101beta1IDByIP) {
	o.SiteIdByIp = &v
}

// GetIps returns the Ips field value if set, zero value otherwise.
func (o *V202101beta1TracerouteLookup) GetIps() []V202101beta1IPInfo {
	if o == nil || o.Ips == nil {
		var ret []V202101beta1IPInfo
		return ret
	}
	return *o.Ips
}

// GetIpsOk returns a tuple with the Ips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1TracerouteLookup) GetIpsOk() (*[]V202101beta1IPInfo, bool) {
	if o == nil || o.Ips == nil {
		return nil, false
	}
	return o.Ips, true
}

// HasIps returns a boolean if a field has been set.
func (o *V202101beta1TracerouteLookup) HasIps() bool {
	if o != nil && o.Ips != nil {
		return true
	}

	return false
}

// SetIps gets a reference to the given []V202101beta1IPInfo and assigns it to the Ips field.
func (o *V202101beta1TracerouteLookup) SetIps(v []V202101beta1IPInfo) {
	o.Ips = &v
}

func (o V202101beta1TracerouteLookup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AgentIdByIp != nil {
		toSerialize["agentIdByIp"] = o.AgentIdByIp
	}
	if o.Agents != nil {
		toSerialize["agents"] = o.Agents
	}
	if o.Asns != nil {
		toSerialize["asns"] = o.Asns
	}
	if o.DeviceIdByIp != nil {
		toSerialize["deviceIdByIp"] = o.DeviceIdByIp
	}
	if o.SiteIdByIp != nil {
		toSerialize["siteIdByIp"] = o.SiteIdByIp
	}
	if o.Ips != nil {
		toSerialize["ips"] = o.Ips
	}
	return json.Marshal(toSerialize)
}

type NullableV202101beta1TracerouteLookup struct {
	value *V202101beta1TracerouteLookup
	isSet bool
}

func (v NullableV202101beta1TracerouteLookup) Get() *V202101beta1TracerouteLookup {
	return v.value
}

func (v *NullableV202101beta1TracerouteLookup) Set(val *V202101beta1TracerouteLookup) {
	v.value = val
	v.isSet = true
}

func (v NullableV202101beta1TracerouteLookup) IsSet() bool {
	return v.isSet
}

func (v *NullableV202101beta1TracerouteLookup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV202101beta1TracerouteLookup(val *V202101beta1TracerouteLookup) *NullableV202101beta1TracerouteLookup {
	return &NullableV202101beta1TracerouteLookup{value: val, isSet: true}
}

func (v NullableV202101beta1TracerouteLookup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV202101beta1TracerouteLookup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
