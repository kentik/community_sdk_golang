# \CloudExportAdminServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudExportAdminServiceCreateCloudExport**](CloudExportAdminServiceApi.md#CloudExportAdminServiceCreateCloudExport) | **Post** /cloud_export/v202101beta1/exports | 
[**CloudExportAdminServiceDeleteCloudExport**](CloudExportAdminServiceApi.md#CloudExportAdminServiceDeleteCloudExport) | **Delete** /cloud_export/v202101beta1/exports/{export.id} | 
[**CloudExportAdminServiceGetCloudExport**](CloudExportAdminServiceApi.md#CloudExportAdminServiceGetCloudExport) | **Get** /cloud_export/v202101beta1/exports/{export.id} | 
[**CloudExportAdminServiceListCloudExport**](CloudExportAdminServiceApi.md#CloudExportAdminServiceListCloudExport) | **Get** /cloud_export/v202101beta1/exports | 
[**CloudExportAdminServicePatchCloudExport**](CloudExportAdminServiceApi.md#CloudExportAdminServicePatchCloudExport) | **Patch** /cloud_export/v202101beta1/exports/{export.id} | 
[**CloudExportAdminServiceUpdateCloudExport**](CloudExportAdminServiceApi.md#CloudExportAdminServiceUpdateCloudExport) | **Put** /cloud_export/v202101beta1/exports/{export.id} | 



## CloudExportAdminServiceCreateCloudExport

> V202101beta1CreateCloudExportResponse CloudExportAdminServiceCreateCloudExport(ctx).V202101beta1CreateCloudExportRequest(v202101beta1CreateCloudExportRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    v202101beta1CreateCloudExportRequest := *openapiclient.NewV202101beta1CreateCloudExportRequest() // V202101beta1CreateCloudExportRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudExportAdminServiceApi.CloudExportAdminServiceCreateCloudExport(context.Background()).V202101beta1CreateCloudExportRequest(v202101beta1CreateCloudExportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudExportAdminServiceApi.CloudExportAdminServiceCreateCloudExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudExportAdminServiceCreateCloudExport`: V202101beta1CreateCloudExportResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudExportAdminServiceApi.CloudExportAdminServiceCreateCloudExport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudExportAdminServiceCreateCloudExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v202101beta1CreateCloudExportRequest** | [**V202101beta1CreateCloudExportRequest**](V202101beta1CreateCloudExportRequest.md) |  | 

### Return type

[**V202101beta1CreateCloudExportResponse**](V202101beta1CreateCloudExportResponse.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudExportAdminServiceDeleteCloudExport

> map[string]interface{} CloudExportAdminServiceDeleteCloudExport(ctx, exportId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    exportId := "exportId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudExportAdminServiceApi.CloudExportAdminServiceDeleteCloudExport(context.Background(), exportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudExportAdminServiceApi.CloudExportAdminServiceDeleteCloudExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudExportAdminServiceDeleteCloudExport`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CloudExportAdminServiceApi.CloudExportAdminServiceDeleteCloudExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudExportAdminServiceDeleteCloudExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudExportAdminServiceGetCloudExport

> V202101beta1GetCloudExportResponse CloudExportAdminServiceGetCloudExport(ctx, exportId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    exportId := "exportId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudExportAdminServiceApi.CloudExportAdminServiceGetCloudExport(context.Background(), exportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudExportAdminServiceApi.CloudExportAdminServiceGetCloudExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudExportAdminServiceGetCloudExport`: V202101beta1GetCloudExportResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudExportAdminServiceApi.CloudExportAdminServiceGetCloudExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudExportAdminServiceGetCloudExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V202101beta1GetCloudExportResponse**](V202101beta1GetCloudExportResponse.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudExportAdminServiceListCloudExport

> V202101beta1ListCloudExportResponse CloudExportAdminServiceListCloudExport(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudExportAdminServiceApi.CloudExportAdminServiceListCloudExport(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudExportAdminServiceApi.CloudExportAdminServiceListCloudExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudExportAdminServiceListCloudExport`: V202101beta1ListCloudExportResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudExportAdminServiceApi.CloudExportAdminServiceListCloudExport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudExportAdminServiceListCloudExportRequest struct via the builder pattern


### Return type

[**V202101beta1ListCloudExportResponse**](V202101beta1ListCloudExportResponse.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudExportAdminServicePatchCloudExport

> V202101beta1PatchCloudExportResponse CloudExportAdminServicePatchCloudExport(ctx, exportId).V202101beta1PatchCloudExportRequest(v202101beta1PatchCloudExportRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    exportId := "exportId_example" // string | The internal cloud export identifier. This is Read-only and assigned by Kentik.
    v202101beta1PatchCloudExportRequest := *openapiclient.NewV202101beta1PatchCloudExportRequest() // V202101beta1PatchCloudExportRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudExportAdminServiceApi.CloudExportAdminServicePatchCloudExport(context.Background(), exportId).V202101beta1PatchCloudExportRequest(v202101beta1PatchCloudExportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudExportAdminServiceApi.CloudExportAdminServicePatchCloudExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudExportAdminServicePatchCloudExport`: V202101beta1PatchCloudExportResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudExportAdminServiceApi.CloudExportAdminServicePatchCloudExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exportId** | **string** | The internal cloud export identifier. This is Read-only and assigned by Kentik. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudExportAdminServicePatchCloudExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v202101beta1PatchCloudExportRequest** | [**V202101beta1PatchCloudExportRequest**](V202101beta1PatchCloudExportRequest.md) |  | 

### Return type

[**V202101beta1PatchCloudExportResponse**](V202101beta1PatchCloudExportResponse.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudExportAdminServiceUpdateCloudExport

> V202101beta1UpdateCloudExportResponse CloudExportAdminServiceUpdateCloudExport(ctx, exportId).V202101beta1UpdateCloudExportRequest(v202101beta1UpdateCloudExportRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    exportId := "exportId_example" // string | The internal cloud export identifier. This is Read-only and assigned by Kentik.
    v202101beta1UpdateCloudExportRequest := *openapiclient.NewV202101beta1UpdateCloudExportRequest() // V202101beta1UpdateCloudExportRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudExportAdminServiceApi.CloudExportAdminServiceUpdateCloudExport(context.Background(), exportId).V202101beta1UpdateCloudExportRequest(v202101beta1UpdateCloudExportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudExportAdminServiceApi.CloudExportAdminServiceUpdateCloudExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudExportAdminServiceUpdateCloudExport`: V202101beta1UpdateCloudExportResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudExportAdminServiceApi.CloudExportAdminServiceUpdateCloudExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exportId** | **string** | The internal cloud export identifier. This is Read-only and assigned by Kentik. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudExportAdminServiceUpdateCloudExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v202101beta1UpdateCloudExportRequest** | [**V202101beta1UpdateCloudExportRequest**](V202101beta1UpdateCloudExportRequest.md) |  | 

### Return type

[**V202101beta1UpdateCloudExportResponse**](V202101beta1UpdateCloudExportResponse.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

