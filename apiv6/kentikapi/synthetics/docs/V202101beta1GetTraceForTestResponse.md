# V202101beta1GetTraceForTestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpInfo** | Pointer to [**[]V202101beta1IPInfo**](V202101beta1IPInfo.md) |  | [optional] 
**Results** | Pointer to [**[]V202101beta1TracerouteResult**](V202101beta1TracerouteResult.md) |  | [optional] 

## Methods

### NewV202101beta1GetTraceForTestResponse

`func NewV202101beta1GetTraceForTestResponse() *V202101beta1GetTraceForTestResponse`

NewV202101beta1GetTraceForTestResponse instantiates a new V202101beta1GetTraceForTestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV202101beta1GetTraceForTestResponseWithDefaults

`func NewV202101beta1GetTraceForTestResponseWithDefaults() *V202101beta1GetTraceForTestResponse`

NewV202101beta1GetTraceForTestResponseWithDefaults instantiates a new V202101beta1GetTraceForTestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpInfo

`func (o *V202101beta1GetTraceForTestResponse) GetIpInfo() []V202101beta1IPInfo`

GetIpInfo returns the IpInfo field if non-nil, zero value otherwise.

### GetIpInfoOk

`func (o *V202101beta1GetTraceForTestResponse) GetIpInfoOk() (*[]V202101beta1IPInfo, bool)`

GetIpInfoOk returns a tuple with the IpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpInfo

`func (o *V202101beta1GetTraceForTestResponse) SetIpInfo(v []V202101beta1IPInfo)`

SetIpInfo sets IpInfo field to given value.

### HasIpInfo

`func (o *V202101beta1GetTraceForTestResponse) HasIpInfo() bool`

HasIpInfo returns a boolean if a field has been set.

### GetResults

`func (o *V202101beta1GetTraceForTestResponse) GetResults() []V202101beta1TracerouteResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V202101beta1GetTraceForTestResponse) GetResultsOk() (*[]V202101beta1TracerouteResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V202101beta1GetTraceForTestResponse) SetResults(v []V202101beta1TracerouteResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *V202101beta1GetTraceForTestResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


