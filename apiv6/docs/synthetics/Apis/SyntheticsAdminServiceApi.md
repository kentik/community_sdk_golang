# SyntheticsAdminServiceApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**agentCreate**](SyntheticsAdminServiceApi.md#agentCreate) | **POST** /synthetics/v202101beta1/agents | Create Agent.
[**agentDelete**](SyntheticsAdminServiceApi.md#agentDelete) | **DELETE** /synthetics/v202101beta1/agents/{agent.id} | Delete an agent.
[**agentGet**](SyntheticsAdminServiceApi.md#agentGet) | **GET** /synthetics/v202101beta1/agents/{agent.id} | Get information about an agent.
[**agentPatch**](SyntheticsAdminServiceApi.md#agentPatch) | **PATCH** /synthetics/v202101beta1/agents/{agent.id} | Patch an agent.
[**agentsList**](SyntheticsAdminServiceApi.md#agentsList) | **GET** /synthetics/v202101beta1/agents | List Agents.
[**exportPatch**](SyntheticsAdminServiceApi.md#exportPatch) | **PATCH** /synthetics/v202101beta1/tests/{id} | Patch a Synthetics Test.
[**testCreate**](SyntheticsAdminServiceApi.md#testCreate) | **POST** /synthetics/v202101beta1/tests | Create Synthetics Test.
[**testDelete**](SyntheticsAdminServiceApi.md#testDelete) | **DELETE** /synthetics/v202101beta1/tests/{id} | Delete an Synthetics Test.
[**testGet**](SyntheticsAdminServiceApi.md#testGet) | **GET** /synthetics/v202101beta1/tests/{id} | Get information about Synthetics Test.
[**testStatusUpdate**](SyntheticsAdminServiceApi.md#testStatusUpdate) | **PUT** /synthetics/v202101beta1/tests/{id}/status | Update a test status.
[**testsList**](SyntheticsAdminServiceApi.md#testsList) | **GET** /synthetics/v202101beta1/tests | List Synthetics Tests.


<a name="agentCreate"></a>
# **agentCreate**
> Object agentCreate(body)

Create Agent.

    Create agent from request. Returns created agent.

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **Object**|  |

### Return type

[**Object**](../Models/object.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="agentDelete"></a>
# **agentDelete**
> Object agentDelete(agent.id)

Delete an agent.

    Deletes the agent specified with id.

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent.id** | **String**|  | [default to null]

### Return type

[**Object**](../Models/object.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="agentGet"></a>
# **agentGet**
> v202101beta1GetAgentResponse agentGet(agent.id)

Get information about an agent.

    Returns information about export specified with export ID.

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent.id** | **String**|  | [default to null]

### Return type

[**v202101beta1GetAgentResponse**](../Models/v202101beta1GetAgentResponse.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="agentPatch"></a>
# **agentPatch**
> v202101beta1PatchAgentResponse agentPatch(agent.id, V202101beta1PatchAgentRequest)

Patch an agent.

    Partially Updates the attributes of agent specified with id and update_mask fields.

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent.id** | **String**|  | [default to null]
 **V202101beta1PatchAgentRequest** | [**V202101beta1PatchAgentRequest**](../Models/V202101beta1PatchAgentRequest.md)|  |

### Return type

[**v202101beta1PatchAgentResponse**](../Models/v202101beta1PatchAgentResponse.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="agentsList"></a>
# **agentsList**
> v202101beta1ListAgentsResponse agentsList()

List Agents.

    Returns a list of agents.

### Parameters
This endpoint does not need any parameter.

### Return type

[**v202101beta1ListAgentsResponse**](../Models/v202101beta1ListAgentsResponse.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="exportPatch"></a>
# **exportPatch**
> v202101beta1PatchTestResponse exportPatch(id)

Patch a Synthetics Test.

    Partially Updates the attributes of synthetics test specified with id and update_mask fields.

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | [default to null]

### Return type

[**v202101beta1PatchTestResponse**](../Models/v202101beta1PatchTestResponse.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="testCreate"></a>
# **testCreate**
> v202101beta1CreateTestResponse testCreate(V202101beta1CreateTestRequest)

Create Synthetics Test.

    Create synthetics test from request. Returns created test.

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **V202101beta1CreateTestRequest** | [**V202101beta1CreateTestRequest**](../Models/V202101beta1CreateTestRequest.md)|  |

### Return type

[**v202101beta1CreateTestResponse**](../Models/v202101beta1CreateTestResponse.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="testDelete"></a>
# **testDelete**
> Object testDelete(id)

Delete an Synthetics Test.

    Deletes the synthetics test specified with id.

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | [default to null]

### Return type

[**Object**](../Models/object.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="testGet"></a>
# **testGet**
> v202101beta1GetTestResponse testGet(id)

Get information about Synthetics Test.

    Returns information about synthetics test specified with test ID.

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | [default to null]

### Return type

[**v202101beta1GetTestResponse**](../Models/v202101beta1GetTestResponse.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="testStatusUpdate"></a>
# **testStatusUpdate**
> Object testStatusUpdate(id, V202101beta1SetTestStatusRequest)

Update a test status.

    Update the status of a test.

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | [default to null]
 **V202101beta1SetTestStatusRequest** | [**V202101beta1SetTestStatusRequest**](../Models/V202101beta1SetTestStatusRequest.md)|  |

### Return type

[**Object**](../Models/object.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="testsList"></a>
# **testsList**
> v202101beta1ListTestsResponse testsList(preset)

List Synthetics Tests.

    Returns a list of syntehtics tests.

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **preset** | **Boolean**|  | [optional] [default to null]

### Return type

[**v202101beta1ListTestsResponse**](../Models/v202101beta1ListTestsResponse.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

