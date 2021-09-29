/*
 * Synthetics Monitoring API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 202101beta1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package syntheticsstub

type V202101beta1TraceProbe struct {
	AsPath []int32 `json:"asPath,omitempty"`

	Completed bool `json:"completed,omitempty"`

	HopCount int32 `json:"hopCount,omitempty"`

	RegionPath []string `json:"regionPath,omitempty"`

	SitePath []int32 `json:"sitePath,omitempty"`

	Hops []V202101beta1TraceHop `json:"hops,omitempty"`
}
