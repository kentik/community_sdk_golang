/*
 * Cloud Export Admin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 202101
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package cloudexportstub

type V202101beta1CloudExportType string

// List of V202101beta1CloudExportType
const (
	UNSPECIFIED      V202101beta1CloudExportType = "CLOUD_EXPORT_TYPE_UNSPECIFIED"
	KENTIK_MANAGED   V202101beta1CloudExportType = "CLOUD_EXPORT_TYPE_KENTIK_MANAGED"
	CUSTOMER_MANAGED V202101beta1CloudExportType = "CLOUD_EXPORT_TYPE_CUSTOMER_MANAGED"
)