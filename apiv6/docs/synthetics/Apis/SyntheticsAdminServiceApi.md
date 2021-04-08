# SyntheticsAdminServiceApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**syntheticsAdminServiceCreateAgent**](SyntheticsAdminServiceApi.md#syntheticsAdminServiceCreateAgent) | **POST** /synthetics/v202101beta1/agents | 
[**syntheticsAdminServiceCreateTest**](SyntheticsAdminServiceApi.md#syntheticsAdminServiceCreateTest) | **POST** /synthetics/v202101beta1/tests | 
[**syntheticsAdminServiceDeleteAgent**](SyntheticsAdminServiceApi.md#syntheticsAdminServiceDeleteAgent) | **DELETE** /synthetics/v202101beta1/agents/{agent.id} | 
[**syntheticsAdminServiceDeleteTest**](SyntheticsAdminServiceApi.md#syntheticsAdminServiceDeleteTest) | **DELETE** /synthetics/v202101beta1/tests/{id} | 
[**syntheticsAdminServiceGetAgent**](SyntheticsAdminServiceApi.md#syntheticsAdminServiceGetAgent) | **GET** /synthetics/v202101beta1/agents/{agent.id} | 
[**syntheticsAdminServiceGetTest**](SyntheticsAdminServiceApi.md#syntheticsAdminServiceGetTest) | **GET** /synthetics/v202101beta1/tests/{id} | 
[**syntheticsAdminServiceListAgents**](SyntheticsAdminServiceApi.md#syntheticsAdminServiceListAgents) | **GET** /synthetics/v202101beta1/agents | 
[**syntheticsAdminServiceListTests**](SyntheticsAdminServiceApi.md#syntheticsAdminServiceListTests) | **GET** /synthetics/v202101beta1/tests | 
[**syntheticsAdminServicePatchAgent**](SyntheticsAdminServiceApi.md#syntheticsAdminServicePatchAgent) | **PATCH** /synthetics/v202101beta1/agents/{agent.id} | 
[**syntheticsAdminServicePatchTest**](SyntheticsAdminServiceApi.md#syntheticsAdminServicePatchTest) | **PATCH** /synthetics/v202101beta1/tests/{id} | 
[**syntheticsAdminServiceSetTestStatus**](SyntheticsAdminServiceApi.md#syntheticsAdminServiceSetTestStatus) | **PUT** /synthetics/v202101beta1/tests/{id}/status | 


<a name="syntheticsAdminServiceCreateAgent"></a>
# **syntheticsAdminServiceCreateAgent**
> Object syntheticsAdminServiceCreateAgent(body)



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **Object**|  |

### Return type

[**Object**](../Models/object.md)

### Authorization

[X-CH-Auth-API-Token](../README.md#X-CH-Auth-API-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="syntheticsAdminServiceCreateTest"></a>
# **syntheticsAdminServiceCreateTest**
> v202101beta1CreateTestResponse syntheticsAdminServiceCreateTest(V202101beta1CreateTestRequest)



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **V202101beta1CreateTestRequest** | [**V202101beta1CreateTestRequest**](../Models/V202101beta1CreateTestRequest.md)|  |

### Return type

[**v202101beta1CreateTestResponse**](../Models/v202101beta1CreateTestResponse.md)

### Authorization

[X-CH-Auth-API-Token](../README.md#X-CH-Auth-API-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="syntheticsAdminServiceDeleteAgent"></a>
# **syntheticsAdminServiceDeleteAgent**
> Object syntheticsAdminServiceDeleteAgent(agent.id)



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent.id** | **String**|  | [default to null]

### Return type

[**Object**](../Models/object.md)

### Authorization

[X-CH-Auth-API-Token](../README.md#X-CH-Auth-API-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="syntheticsAdminServiceDeleteTest"></a>
# **syntheticsAdminServiceDeleteTest**
> Object syntheticsAdminServiceDeleteTest(id)



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | [default to null]

### Return type

[**Object**](../Models/object.md)

### Authorization

[X-CH-Auth-API-Token](../README.md#X-CH-Auth-API-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="syntheticsAdminServiceGetAgent"></a>
# **syntheticsAdminServiceGetAgent**
> v202101beta1GetAgentResponse syntheticsAdminServiceGetAgent(agent.id)



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent.id** | **String**|  | [default to null]

### Return type

[**v202101beta1GetAgentResponse**](../Models/v202101beta1GetAgentResponse.md)

### Authorization

[X-CH-Auth-API-Token](../README.md#X-CH-Auth-API-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="syntheticsAdminServiceGetTest"></a>
# **syntheticsAdminServiceGetTest**
> v202101beta1GetTestResponse syntheticsAdminServiceGetTest(id)



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | [default to null]

### Return type

[**v202101beta1GetTestResponse**](../Models/v202101beta1GetTestResponse.md)

### Authorization

[X-CH-Auth-API-Token](../README.md#X-CH-Auth-API-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="syntheticsAdminServiceListAgents"></a>
# **syntheticsAdminServiceListAgents**
> v202101beta1ListAgentsResponse syntheticsAdminServiceListAgents()



### Parameters
This endpoint does not need any parameter.

### Return type

[**v202101beta1ListAgentsResponse**](../Models/v202101beta1ListAgentsResponse.md)

### Authorization

[X-CH-Auth-API-Token](../README.md#X-CH-Auth-API-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="syntheticsAdminServiceListTests"></a>
# **syntheticsAdminServiceListTests**
> v202101beta1ListTestsResponse syntheticsAdminServiceListTests(preset)



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **preset** | **Boolean**|  | [optional] [default to null]

### Return type

[**v202101beta1ListTestsResponse**](../Models/v202101beta1ListTestsResponse.md)

### Authorization

[X-CH-Auth-API-Token](../README.md#X-CH-Auth-API-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="syntheticsAdminServicePatchAgent"></a>
# **syntheticsAdminServicePatchAgent**
> v202101beta1PatchAgentResponse syntheticsAdminServicePatchAgent(agent.id, V202101beta1PatchAgentRequest)



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent.id** | **String**|  | [default to null]
 **V202101beta1PatchAgentRequest** | [**V202101beta1PatchAgentRequest**](../Models/V202101beta1PatchAgentRequest.md)|  |

### Return type

[**v202101beta1PatchAgentResponse**](../Models/v202101beta1PatchAgentResponse.md)

### Authorization

[X-CH-Auth-API-Token](../README.md#X-CH-Auth-API-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="syntheticsAdminServicePatchTest"></a>
# **syntheticsAdminServicePatchTest**
> v202101beta1PatchTestResponse syntheticsAdminServicePatchTest(id)



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | [default to null]

### Return type

[**v202101beta1PatchTestResponse**](../Models/v202101beta1PatchTestResponse.md)

### Authorization

[X-CH-Auth-API-Token](../README.md#X-CH-Auth-API-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="syntheticsAdminServiceSetTestStatus"></a>
# **syntheticsAdminServiceSetTestStatus**
> Object syntheticsAdminServiceSetTestStatus(id, V202101beta1SetTestStatusRequest)



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | [default to null]
 **V202101beta1SetTestStatusRequest** | [**V202101beta1SetTestStatusRequest**](../Models/V202101beta1SetTestStatusRequest.md)|  |

### Return type

[**Object**](../Models/object.md)

### Authorization

[X-CH-Auth-API-Token](../README.md#X-CH-Auth-API-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

