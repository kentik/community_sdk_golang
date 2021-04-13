# SyntheticsDataServiceApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getHealthForTests**](SyntheticsDataServiceApi.md#getHealthForTests) | **POST** /synthetics/v202101beta1/health/tests | Get health status for synthetics test.
[**getTraceForTest**](SyntheticsDataServiceApi.md#getTraceForTest) | **POST** /synthetics/v202101beta1/tests/{id}/results/trace | Get trace route data.


<a name="getHealthForTests"></a>
# **getHealthForTests**
> v202101beta1GetHealthForTestsResponse getHealthForTests(V202101beta1GetHealthForTestsRequest)

Get health status for synthetics test.

    Get synthetics health test for login user. Also returns mesh data on request.

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **V202101beta1GetHealthForTestsRequest** | [**V202101beta1GetHealthForTestsRequest**](../Models/V202101beta1GetHealthForTestsRequest.md)|  |

### Return type

[**v202101beta1GetHealthForTestsResponse**](../Models/v202101beta1GetHealthForTestsResponse.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="getTraceForTest"></a>
# **getTraceForTest**
> v202101beta1GetTraceForTestResponse getTraceForTest(id, V202101beta1GetTraceForTestRequest)

Get trace route data.

    Get trace route data for the specific test id.

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**| Test id | [default to null]
 **V202101beta1GetTraceForTestRequest** | [**V202101beta1GetTraceForTestRequest**](../Models/V202101beta1GetTraceForTestRequest.md)|  |

### Return type

[**v202101beta1GetTraceForTestResponse**](../Models/v202101beta1GetTraceForTestResponse.md)

### Authorization

[email](../README.md#email), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

