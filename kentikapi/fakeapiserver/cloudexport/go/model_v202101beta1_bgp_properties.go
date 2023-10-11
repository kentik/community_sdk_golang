/*
 * Cloud Export Admin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 202101beta1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package cloudexportstub

// V202101beta1BgpProperties - Optional BGP related settings.
type V202101beta1BgpProperties struct {

	// If true, apply BGP data discovered via another device to the flow from this export.
	ApplyBgp bool `json:"applyBgp,omitempty"`

	UseBgpDeviceId string `json:"useBgpDeviceId,omitempty"`

	DeviceBgpType string `json:"deviceBgpType,omitempty"`
}