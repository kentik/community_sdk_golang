# CloudExportAdminServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**cloudExportAdminServiceCreateCloudExport**](CloudExportAdminServiceApi.md#cloudExportAdminServiceCreateCloudExport) | **POST** /cloud_export/v202101beta1/exports | 
[**cloudExportAdminServiceDeleteCloudExport**](CloudExportAdminServiceApi.md#cloudExportAdminServiceDeleteCloudExport) | **DELETE** /cloud_export/v202101beta1/exports/{export.id} | 
[**cloudExportAdminServiceGetCloudExport**](CloudExportAdminServiceApi.md#cloudExportAdminServiceGetCloudExport) | **GET** /cloud_export/v202101beta1/exports/{export.id} | 
[**cloudExportAdminServiceListCloudExport**](CloudExportAdminServiceApi.md#cloudExportAdminServiceListCloudExport) | **GET** /cloud_export/v202101beta1/exports | 
[**cloudExportAdminServicePatchCloudExport**](CloudExportAdminServiceApi.md#cloudExportAdminServicePatchCloudExport) | **PATCH** /cloud_export/v202101beta1/exports/{export.id} | 
[**cloudExportAdminServiceUpdateCloudExport**](CloudExportAdminServiceApi.md#cloudExportAdminServiceUpdateCloudExport) | **PUT** /cloud_export/v202101beta1/exports/{export.id} | 


<a name="cloudExportAdminServiceCreateCloudExport"></a>
# **cloudExportAdminServiceCreateCloudExport**
> v202101beta1CreateCloudExportResponse cloudExportAdminServiceCreateCloudExport(V202101beta1CreateCloudExportRequest)



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **V202101beta1CreateCloudExportRequest** | [**V202101beta1CreateCloudExportRequest**](../Models/V202101beta1CreateCloudExportRequest.md)|  |

### Return type

[**v202101beta1CreateCloudExportResponse**](../Models/v202101beta1CreateCloudExportResponse.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="cloudExportAdminServiceDeleteCloudExport"></a>
# **cloudExportAdminServiceDeleteCloudExport**
> Object cloudExportAdminServiceDeleteCloudExport(export.id)



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **export.id** | **String**|  | [default to null]

### Return type

[**Object**](../Models/object.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="cloudExportAdminServiceGetCloudExport"></a>
# **cloudExportAdminServiceGetCloudExport**
> v202101beta1GetCloudExportResponse cloudExportAdminServiceGetCloudExport(export.id)



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **export.id** | **String**|  | [default to null]

### Return type

[**v202101beta1GetCloudExportResponse**](../Models/v202101beta1GetCloudExportResponse.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="cloudExportAdminServiceListCloudExport"></a>
# **cloudExportAdminServiceListCloudExport**
> v202101beta1ListCloudExportResponse cloudExportAdminServiceListCloudExport()



### Parameters
This endpoint does not need any parameter.

### Return type

[**v202101beta1ListCloudExportResponse**](../Models/v202101beta1ListCloudExportResponse.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="cloudExportAdminServicePatchCloudExport"></a>
# **cloudExportAdminServicePatchCloudExport**
> v202101beta1PatchCloudExportResponse cloudExportAdminServicePatchCloudExport(export.id, V202101beta1PatchCloudExportRequest)



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **export.id** | **String**| The internal cloud export identifier. This is Read-only and assigned by Kentik. | [default to null]
 **V202101beta1PatchCloudExportRequest** | [**V202101beta1PatchCloudExportRequest**](../Models/V202101beta1PatchCloudExportRequest.md)|  |

### Return type

[**v202101beta1PatchCloudExportResponse**](../Models/v202101beta1PatchCloudExportResponse.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="cloudExportAdminServiceUpdateCloudExport"></a>
# **cloudExportAdminServiceUpdateCloudExport**
> v202101beta1UpdateCloudExportResponse cloudExportAdminServiceUpdateCloudExport(export.id, V202101beta1UpdateCloudExportRequest)



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **export.id** | **String**| The internal cloud export identifier. This is Read-only and assigned by Kentik. | [default to null]
 **V202101beta1UpdateCloudExportRequest** | [**V202101beta1UpdateCloudExportRequest**](../Models/V202101beta1UpdateCloudExportRequest.md)|  |

### Return type

[**v202101beta1UpdateCloudExportResponse**](../Models/v202101beta1UpdateCloudExportResponse.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

