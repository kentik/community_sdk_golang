# \SyntheticsDataServiceApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHealthForTests**](SyntheticsDataServiceApi.md#GetHealthForTests) | **Post** /synthetics/v202101beta1/health/tests | Get health status for synthetics test.
[**GetTraceForTest**](SyntheticsDataServiceApi.md#GetTraceForTest) | **Post** /synthetics/v202101beta1/tests/{id}/results/trace | Get trace route data.



## GetHealthForTests

> V202101beta1GetHealthForTestsResponse GetHealthForTests(ctx).V202101beta1GetHealthForTestsRequest(v202101beta1GetHealthForTestsRequest).Execute()

Get health status for synthetics test.



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
    resp, r, err := api_client.SyntheticsDataServiceApi.GetHealthForTests(context.Background()).V202101beta1GetHealthForTestsRequest(v202101beta1GetHealthForTestsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsDataServiceApi.GetHealthForTests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHealthForTests`: V202101beta1GetHealthForTestsResponse
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsDataServiceApi.GetHealthForTests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHealthForTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v202101beta1GetHealthForTestsRequest** | [**V202101beta1GetHealthForTestsRequest**](V202101beta1GetHealthForTestsRequest.md) |  | 

### Return type

[**V202101beta1GetHealthForTestsResponse**](V202101beta1GetHealthForTestsResponse.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTraceForTest

> V202101beta1GetTraceForTestResponse GetTraceForTest(ctx, id).V202101beta1GetTraceForTestRequest(v202101beta1GetTraceForTestRequest).Execute()

Get trace route data.



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
    resp, r, err := api_client.SyntheticsDataServiceApi.GetTraceForTest(context.Background(), id).V202101beta1GetTraceForTestRequest(v202101beta1GetTraceForTestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsDataServiceApi.GetTraceForTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTraceForTest`: V202101beta1GetTraceForTestResponse
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsDataServiceApi.GetTraceForTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTraceForTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v202101beta1GetTraceForTestRequest** | [**V202101beta1GetTraceForTestRequest**](V202101beta1GetTraceForTestRequest.md) |  | 

### Return type

[**V202101beta1GetTraceForTestResponse**](V202101beta1GetTraceForTestResponse.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

