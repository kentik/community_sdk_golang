/*
 * Synthetics Monitoring API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 202101beta1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package syntheticsstub

type V202101beta1MeshResponse struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	LocalIp string `json:"localIp,omitempty"`

	Ip string `json:"ip,omitempty"`

	Alias string `json:"alias,omitempty"`

	Columns []V202101beta1MeshColumn `json:"columns,omitempty"`
}