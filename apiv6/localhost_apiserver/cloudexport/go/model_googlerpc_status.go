/*
 * Cloud Export Admin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 202101
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package cloudexportstub

type GooglerpcStatus struct {
	Code int32 `json:"code,omitempty"`

	Message string `json:"message,omitempty"`

	Details []ProtobufAny `json:"details,omitempty"`
}