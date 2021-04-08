# \SyntheticsAdminServiceApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SyntheticsAdminServiceCreateAgent**](SyntheticsAdminServiceApi.md#SyntheticsAdminServiceCreateAgent) | **Post** /synthetics/v202101beta1/agents | 
[**SyntheticsAdminServiceCreateTest**](SyntheticsAdminServiceApi.md#SyntheticsAdminServiceCreateTest) | **Post** /synthetics/v202101beta1/tests | 
[**SyntheticsAdminServiceDeleteAgent**](SyntheticsAdminServiceApi.md#SyntheticsAdminServiceDeleteAgent) | **Delete** /synthetics/v202101beta1/agents/{agent.id} | 
[**SyntheticsAdminServiceDeleteTest**](SyntheticsAdminServiceApi.md#SyntheticsAdminServiceDeleteTest) | **Delete** /synthetics/v202101beta1/tests/{id} | 
[**SyntheticsAdminServiceGetAgent**](SyntheticsAdminServiceApi.md#SyntheticsAdminServiceGetAgent) | **Get** /synthetics/v202101beta1/agents/{agent.id} | 
[**SyntheticsAdminServiceGetTest**](SyntheticsAdminServiceApi.md#SyntheticsAdminServiceGetTest) | **Get** /synthetics/v202101beta1/tests/{id} | 
[**SyntheticsAdminServiceListAgents**](SyntheticsAdminServiceApi.md#SyntheticsAdminServiceListAgents) | **Get** /synthetics/v202101beta1/agents | 
[**SyntheticsAdminServiceListTests**](SyntheticsAdminServiceApi.md#SyntheticsAdminServiceListTests) | **Get** /synthetics/v202101beta1/tests | 
[**SyntheticsAdminServicePatchAgent**](SyntheticsAdminServiceApi.md#SyntheticsAdminServicePatchAgent) | **Patch** /synthetics/v202101beta1/agents/{agent.id} | 
[**SyntheticsAdminServicePatchTest**](SyntheticsAdminServiceApi.md#SyntheticsAdminServicePatchTest) | **Patch** /synthetics/v202101beta1/tests/{id} | 
[**SyntheticsAdminServiceSetTestStatus**](SyntheticsAdminServiceApi.md#SyntheticsAdminServiceSetTestStatus) | **Put** /synthetics/v202101beta1/tests/{id}/status | 



## SyntheticsAdminServiceCreateAgent

> map[string]interface{} SyntheticsAdminServiceCreateAgent(ctx).Body(body).Execute()



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
    body := map[string]interface{}(Object) // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsAdminServiceApi.SyntheticsAdminServiceCreateAgent(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsAdminServiceApi.SyntheticsAdminServiceCreateAgent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyntheticsAdminServiceCreateAgent`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsAdminServiceApi.SyntheticsAdminServiceCreateAgent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyntheticsAdminServiceCreateAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[X-CH-Auth-API-Token](../README.md#X-CH-Auth-API-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyntheticsAdminServiceCreateTest

> V202101beta1CreateTestResponse SyntheticsAdminServiceCreateTest(ctx).V202101beta1CreateTestRequest(v202101beta1CreateTestRequest).Execute()



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
    v202101beta1CreateTestRequest := *openapiclient.NewV202101beta1CreateTestRequest() // V202101beta1CreateTestRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsAdminServiceApi.SyntheticsAdminServiceCreateTest(context.Background()).V202101beta1CreateTestRequest(v202101beta1CreateTestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsAdminServiceApi.SyntheticsAdminServiceCreateTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyntheticsAdminServiceCreateTest`: V202101beta1CreateTestResponse
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsAdminServiceApi.SyntheticsAdminServiceCreateTest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyntheticsAdminServiceCreateTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v202101beta1CreateTestRequest** | [**V202101beta1CreateTestRequest**](V202101beta1CreateTestRequest.md) |  | 

### Return type

[**V202101beta1CreateTestResponse**](V202101beta1CreateTestResponse.md)

### Authorization

[X-CH-Auth-API-Token](../README.md#X-CH-Auth-API-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyntheticsAdminServiceDeleteAgent

> map[string]interface{} SyntheticsAdminServiceDeleteAgent(ctx, agentId).Execute()



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
    agentId := "agentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsAdminServiceApi.SyntheticsAdminServiceDeleteAgent(context.Background(), agentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsAdminServiceApi.SyntheticsAdminServiceDeleteAgent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyntheticsAdminServiceDeleteAgent`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsAdminServiceApi.SyntheticsAdminServiceDeleteAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyntheticsAdminServiceDeleteAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[X-CH-Auth-API-Token](../README.md#X-CH-Auth-API-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyntheticsAdminServiceDeleteTest

> map[string]interface{} SyntheticsAdminServiceDeleteTest(ctx, id).Execute()



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsAdminServiceApi.SyntheticsAdminServiceDeleteTest(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsAdminServiceApi.SyntheticsAdminServiceDeleteTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyntheticsAdminServiceDeleteTest`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsAdminServiceApi.SyntheticsAdminServiceDeleteTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyntheticsAdminServiceDeleteTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[X-CH-Auth-API-Token](../README.md#X-CH-Auth-API-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyntheticsAdminServiceGetAgent

> V202101beta1GetAgentResponse SyntheticsAdminServiceGetAgent(ctx, agentId).Execute()



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
    agentId := "agentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsAdminServiceApi.SyntheticsAdminServiceGetAgent(context.Background(), agentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsAdminServiceApi.SyntheticsAdminServiceGetAgent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyntheticsAdminServiceGetAgent`: V202101beta1GetAgentResponse
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsAdminServiceApi.SyntheticsAdminServiceGetAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyntheticsAdminServiceGetAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V202101beta1GetAgentResponse**](V202101beta1GetAgentResponse.md)

### Authorization

[X-CH-Auth-API-Token](../README.md#X-CH-Auth-API-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyntheticsAdminServiceGetTest

> V202101beta1GetTestResponse SyntheticsAdminServiceGetTest(ctx, id).Execute()



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsAdminServiceApi.SyntheticsAdminServiceGetTest(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsAdminServiceApi.SyntheticsAdminServiceGetTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyntheticsAdminServiceGetTest`: V202101beta1GetTestResponse
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsAdminServiceApi.SyntheticsAdminServiceGetTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyntheticsAdminServiceGetTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V202101beta1GetTestResponse**](V202101beta1GetTestResponse.md)

### Authorization

[X-CH-Auth-API-Token](../README.md#X-CH-Auth-API-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyntheticsAdminServiceListAgents

> V202101beta1ListAgentsResponse SyntheticsAdminServiceListAgents(ctx).Execute()



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
    resp, r, err := api_client.SyntheticsAdminServiceApi.SyntheticsAdminServiceListAgents(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsAdminServiceApi.SyntheticsAdminServiceListAgents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyntheticsAdminServiceListAgents`: V202101beta1ListAgentsResponse
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsAdminServiceApi.SyntheticsAdminServiceListAgents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSyntheticsAdminServiceListAgentsRequest struct via the builder pattern


### Return type

[**V202101beta1ListAgentsResponse**](V202101beta1ListAgentsResponse.md)

### Authorization

[X-CH-Auth-API-Token](../README.md#X-CH-Auth-API-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyntheticsAdminServiceListTests

> V202101beta1ListTestsResponse SyntheticsAdminServiceListTests(ctx).Preset(preset).Execute()



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
    preset := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsAdminServiceApi.SyntheticsAdminServiceListTests(context.Background()).Preset(preset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsAdminServiceApi.SyntheticsAdminServiceListTests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyntheticsAdminServiceListTests`: V202101beta1ListTestsResponse
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsAdminServiceApi.SyntheticsAdminServiceListTests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyntheticsAdminServiceListTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **preset** | **bool** |  | 

### Return type

[**V202101beta1ListTestsResponse**](V202101beta1ListTestsResponse.md)

### Authorization

[X-CH-Auth-API-Token](../README.md#X-CH-Auth-API-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyntheticsAdminServicePatchAgent

> V202101beta1PatchAgentResponse SyntheticsAdminServicePatchAgent(ctx, agentId).V202101beta1PatchAgentRequest(v202101beta1PatchAgentRequest).Execute()



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
    agentId := "agentId_example" // string | 
    v202101beta1PatchAgentRequest := *openapiclient.NewV202101beta1PatchAgentRequest() // V202101beta1PatchAgentRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsAdminServiceApi.SyntheticsAdminServicePatchAgent(context.Background(), agentId).V202101beta1PatchAgentRequest(v202101beta1PatchAgentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsAdminServiceApi.SyntheticsAdminServicePatchAgent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyntheticsAdminServicePatchAgent`: V202101beta1PatchAgentResponse
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsAdminServiceApi.SyntheticsAdminServicePatchAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyntheticsAdminServicePatchAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v202101beta1PatchAgentRequest** | [**V202101beta1PatchAgentRequest**](V202101beta1PatchAgentRequest.md) |  | 

### Return type

[**V202101beta1PatchAgentResponse**](V202101beta1PatchAgentResponse.md)

### Authorization

[X-CH-Auth-API-Token](../README.md#X-CH-Auth-API-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyntheticsAdminServicePatchTest

> V202101beta1PatchTestResponse SyntheticsAdminServicePatchTest(ctx, id).Execute()



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsAdminServiceApi.SyntheticsAdminServicePatchTest(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsAdminServiceApi.SyntheticsAdminServicePatchTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyntheticsAdminServicePatchTest`: V202101beta1PatchTestResponse
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsAdminServiceApi.SyntheticsAdminServicePatchTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyntheticsAdminServicePatchTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V202101beta1PatchTestResponse**](V202101beta1PatchTestResponse.md)

### Authorization

[X-CH-Auth-API-Token](../README.md#X-CH-Auth-API-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyntheticsAdminServiceSetTestStatus

> map[string]interface{} SyntheticsAdminServiceSetTestStatus(ctx, id).V202101beta1SetTestStatusRequest(v202101beta1SetTestStatusRequest).Execute()



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
    id := "id_example" // string | 
    v202101beta1SetTestStatusRequest := *openapiclient.NewV202101beta1SetTestStatusRequest() // V202101beta1SetTestStatusRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsAdminServiceApi.SyntheticsAdminServiceSetTestStatus(context.Background(), id).V202101beta1SetTestStatusRequest(v202101beta1SetTestStatusRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsAdminServiceApi.SyntheticsAdminServiceSetTestStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyntheticsAdminServiceSetTestStatus`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsAdminServiceApi.SyntheticsAdminServiceSetTestStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyntheticsAdminServiceSetTestStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v202101beta1SetTestStatusRequest** | [**V202101beta1SetTestStatusRequest**](V202101beta1SetTestStatusRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[X-CH-Auth-API-Token](../README.md#X-CH-Auth-API-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

