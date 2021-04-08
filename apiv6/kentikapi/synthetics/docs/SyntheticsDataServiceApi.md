# \SyntheticsDataServiceApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SyntheticsDataServiceGetHealthForTests**](SyntheticsDataServiceApi.md#SyntheticsDataServiceGetHealthForTests) | **Post** /synthetics/v202101beta1/health/tests | Get health data for a set of tests
[**SyntheticsDataServiceGetTraceForTest**](SyntheticsDataServiceApi.md#SyntheticsDataServiceGetTraceForTest) | **Post** /synthetics/v202101beta1/tests/{id}/results/trace | TODO: Get traces for a single test. Not implemented.



## SyntheticsDataServiceGetHealthForTests

> V202101beta1GetHealthForTestsResponse SyntheticsDataServiceGetHealthForTests(ctx).V202101beta1GetHealthForTestsRequest(v202101beta1GetHealthForTestsRequest).Execute()

Get health data for a set of tests

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
    v202101beta1GetHealthForTestsRequest := *openapiclient.NewV202101beta1GetHealthForTestsRequest() // V202101beta1GetHealthForTestsRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsDataServiceApi.SyntheticsDataServiceGetHealthForTests(context.Background()).V202101beta1GetHealthForTestsRequest(v202101beta1GetHealthForTestsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsDataServiceApi.SyntheticsDataServiceGetHealthForTests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyntheticsDataServiceGetHealthForTests`: V202101beta1GetHealthForTestsResponse
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsDataServiceApi.SyntheticsDataServiceGetHealthForTests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyntheticsDataServiceGetHealthForTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v202101beta1GetHealthForTestsRequest** | [**V202101beta1GetHealthForTestsRequest**](V202101beta1GetHealthForTestsRequest.md) |  | 

### Return type

[**V202101beta1GetHealthForTestsResponse**](V202101beta1GetHealthForTestsResponse.md)

### Authorization

[X-CH-Auth-API-Token](../README.md#X-CH-Auth-API-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyntheticsDataServiceGetTraceForTest

> V202101beta1GetTraceForTestResponse SyntheticsDataServiceGetTraceForTest(ctx, id).V202101beta1GetTraceForTestRequest(v202101beta1GetTraceForTestRequest).Execute()

TODO: Get traces for a single test. Not implemented.

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
    id := "id_example" // string | Test id
    v202101beta1GetTraceForTestRequest := *openapiclient.NewV202101beta1GetTraceForTestRequest() // V202101beta1GetTraceForTestRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsDataServiceApi.SyntheticsDataServiceGetTraceForTest(context.Background(), id).V202101beta1GetTraceForTestRequest(v202101beta1GetTraceForTestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsDataServiceApi.SyntheticsDataServiceGetTraceForTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyntheticsDataServiceGetTraceForTest`: V202101beta1GetTraceForTestResponse
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsDataServiceApi.SyntheticsDataServiceGetTraceForTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyntheticsDataServiceGetTraceForTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v202101beta1GetTraceForTestRequest** | [**V202101beta1GetTraceForTestRequest**](V202101beta1GetTraceForTestRequest.md) |  | 

### Return type

[**V202101beta1GetTraceForTestResponse**](V202101beta1GetTraceForTestResponse.md)

### Authorization

[X-CH-Auth-API-Token](../README.md#X-CH-Auth-API-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

