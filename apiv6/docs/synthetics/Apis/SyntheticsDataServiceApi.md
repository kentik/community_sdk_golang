# SyntheticsDataServiceApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**syntheticsDataServiceGetHealthForTests**](SyntheticsDataServiceApi.md#syntheticsDataServiceGetHealthForTests) | **POST** /synthetics/v202101beta1/health/tests | Get health data for a set of tests
[**syntheticsDataServiceGetTraceForTest**](SyntheticsDataServiceApi.md#syntheticsDataServiceGetTraceForTest) | **POST** /synthetics/v202101beta1/tests/{id}/results/trace | TODO: Get traces for a single test. Not implemented.


<a name="syntheticsDataServiceGetHealthForTests"></a>
# **syntheticsDataServiceGetHealthForTests**
> v202101beta1GetHealthForTestsResponse syntheticsDataServiceGetHealthForTests(V202101beta1GetHealthForTestsRequest)

Get health data for a set of tests

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **V202101beta1GetHealthForTestsRequest** | [**V202101beta1GetHealthForTestsRequest**](../Models/V202101beta1GetHealthForTestsRequest.md)|  |

### Return type

[**v202101beta1GetHealthForTestsResponse**](../Models/v202101beta1GetHealthForTestsResponse.md)

### Authorization

[X-CH-Auth-API-Token](../README.md#X-CH-Auth-API-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="syntheticsDataServiceGetTraceForTest"></a>
# **syntheticsDataServiceGetTraceForTest**
> v202101beta1GetTraceForTestResponse syntheticsDataServiceGetTraceForTest(id, V202101beta1GetTraceForTestRequest)

TODO: Get traces for a single test. Not implemented.

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**| Test id | [default to null]
 **V202101beta1GetTraceForTestRequest** | [**V202101beta1GetTraceForTestRequest**](../Models/V202101beta1GetTraceForTestRequest.md)|  |

### Return type

[**v202101beta1GetTraceForTestResponse**](../Models/v202101beta1GetTraceForTestResponse.md)

### Authorization

[X-CH-Auth-API-Token](../README.md#X-CH-Auth-API-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

