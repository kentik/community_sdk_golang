# \SyntheticsAdminServiceApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentCreate**](SyntheticsAdminServiceApi.md#AgentCreate) | **Post** /synthetics/v202101beta1/agents | Create Agent.
[**AgentDelete**](SyntheticsAdminServiceApi.md#AgentDelete) | **Delete** /synthetics/v202101beta1/agents/{agent.id} | Delete an agent.
[**AgentGet**](SyntheticsAdminServiceApi.md#AgentGet) | **Get** /synthetics/v202101beta1/agents/{agent.id} | Get information about an agent.
[**AgentPatch**](SyntheticsAdminServiceApi.md#AgentPatch) | **Patch** /synthetics/v202101beta1/agents/{agent.id} | Patch an agent.
[**AgentsList**](SyntheticsAdminServiceApi.md#AgentsList) | **Get** /synthetics/v202101beta1/agents | List Agents.
[**ExportPatch**](SyntheticsAdminServiceApi.md#ExportPatch) | **Patch** /synthetics/v202101beta1/tests/{id} | Patch a Synthetics Test.
[**TestCreate**](SyntheticsAdminServiceApi.md#TestCreate) | **Post** /synthetics/v202101beta1/tests | Create Synthetics Test.
[**TestDelete**](SyntheticsAdminServiceApi.md#TestDelete) | **Delete** /synthetics/v202101beta1/tests/{id} | Delete an Synthetics Test.
[**TestGet**](SyntheticsAdminServiceApi.md#TestGet) | **Get** /synthetics/v202101beta1/tests/{id} | Get information about Synthetics Test.
[**TestStatusUpdate**](SyntheticsAdminServiceApi.md#TestStatusUpdate) | **Put** /synthetics/v202101beta1/tests/{id}/status | Update a test status.
[**TestsList**](SyntheticsAdminServiceApi.md#TestsList) | **Get** /synthetics/v202101beta1/tests | List Synthetics Tests.



## AgentCreate

> map[string]interface{} AgentCreate(ctx).Body(body).Execute()

Create Agent.



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
    resp, r, err := api_client.SyntheticsAdminServiceApi.AgentCreate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsAdminServiceApi.AgentCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentCreate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsAdminServiceApi.AgentCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentDelete

> map[string]interface{} AgentDelete(ctx, agentId).Execute()

Delete an agent.



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
    resp, r, err := api_client.SyntheticsAdminServiceApi.AgentDelete(context.Background(), agentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsAdminServiceApi.AgentDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsAdminServiceApi.AgentDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentDeleteRequest struct via the builder pattern


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


## AgentGet

> V202101beta1GetAgentResponse AgentGet(ctx, agentId).Execute()

Get information about an agent.



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
    resp, r, err := api_client.SyntheticsAdminServiceApi.AgentGet(context.Background(), agentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsAdminServiceApi.AgentGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentGet`: V202101beta1GetAgentResponse
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsAdminServiceApi.AgentGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V202101beta1GetAgentResponse**](V202101beta1GetAgentResponse.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentPatch

> V202101beta1PatchAgentResponse AgentPatch(ctx, agentId).V202101beta1PatchAgentRequest(v202101beta1PatchAgentRequest).Execute()

Patch an agent.



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
    resp, r, err := api_client.SyntheticsAdminServiceApi.AgentPatch(context.Background(), agentId).V202101beta1PatchAgentRequest(v202101beta1PatchAgentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsAdminServiceApi.AgentPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentPatch`: V202101beta1PatchAgentResponse
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsAdminServiceApi.AgentPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v202101beta1PatchAgentRequest** | [**V202101beta1PatchAgentRequest**](V202101beta1PatchAgentRequest.md) |  | 

### Return type

[**V202101beta1PatchAgentResponse**](V202101beta1PatchAgentResponse.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentsList

> V202101beta1ListAgentsResponse AgentsList(ctx).Execute()

List Agents.



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
    resp, r, err := api_client.SyntheticsAdminServiceApi.AgentsList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsAdminServiceApi.AgentsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentsList`: V202101beta1ListAgentsResponse
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsAdminServiceApi.AgentsList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAgentsListRequest struct via the builder pattern


### Return type

[**V202101beta1ListAgentsResponse**](V202101beta1ListAgentsResponse.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportPatch

> V202101beta1PatchTestResponse ExportPatch(ctx, id).Execute()

Patch a Synthetics Test.



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
    resp, r, err := api_client.SyntheticsAdminServiceApi.ExportPatch(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsAdminServiceApi.ExportPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportPatch`: V202101beta1PatchTestResponse
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsAdminServiceApi.ExportPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V202101beta1PatchTestResponse**](V202101beta1PatchTestResponse.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestCreate

> V202101beta1CreateTestResponse TestCreate(ctx).V202101beta1CreateTestRequest(v202101beta1CreateTestRequest).Execute()

Create Synthetics Test.



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
    resp, r, err := api_client.SyntheticsAdminServiceApi.TestCreate(context.Background()).V202101beta1CreateTestRequest(v202101beta1CreateTestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsAdminServiceApi.TestCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestCreate`: V202101beta1CreateTestResponse
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsAdminServiceApi.TestCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v202101beta1CreateTestRequest** | [**V202101beta1CreateTestRequest**](V202101beta1CreateTestRequest.md) |  | 

### Return type

[**V202101beta1CreateTestResponse**](V202101beta1CreateTestResponse.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestDelete

> map[string]interface{} TestDelete(ctx, id).Execute()

Delete an Synthetics Test.



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
    resp, r, err := api_client.SyntheticsAdminServiceApi.TestDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsAdminServiceApi.TestDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsAdminServiceApi.TestDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestDeleteRequest struct via the builder pattern


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


## TestGet

> V202101beta1GetTestResponse TestGet(ctx, id).Execute()

Get information about Synthetics Test.



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
    resp, r, err := api_client.SyntheticsAdminServiceApi.TestGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsAdminServiceApi.TestGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestGet`: V202101beta1GetTestResponse
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsAdminServiceApi.TestGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V202101beta1GetTestResponse**](V202101beta1GetTestResponse.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestStatusUpdate

> map[string]interface{} TestStatusUpdate(ctx, id).V202101beta1SetTestStatusRequest(v202101beta1SetTestStatusRequest).Execute()

Update a test status.



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
    resp, r, err := api_client.SyntheticsAdminServiceApi.TestStatusUpdate(context.Background(), id).V202101beta1SetTestStatusRequest(v202101beta1SetTestStatusRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsAdminServiceApi.TestStatusUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestStatusUpdate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsAdminServiceApi.TestStatusUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestStatusUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v202101beta1SetTestStatusRequest** | [**V202101beta1SetTestStatusRequest**](V202101beta1SetTestStatusRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestsList

> V202101beta1ListTestsResponse TestsList(ctx).Preset(preset).Execute()

List Synthetics Tests.



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
    resp, r, err := api_client.SyntheticsAdminServiceApi.TestsList(context.Background()).Preset(preset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsAdminServiceApi.TestsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestsList`: V202101beta1ListTestsResponse
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsAdminServiceApi.TestsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **preset** | **bool** |  | 

### Return type

[**V202101beta1ListTestsResponse**](V202101beta1ListTestsResponse.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

