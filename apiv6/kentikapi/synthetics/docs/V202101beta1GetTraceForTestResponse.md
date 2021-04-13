# V202101beta1GetTraceForTestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpInfo** | Pointer to [**[]V202101beta1IPInfo**](V202101beta1IPInfo.md) |  | [optional] 
**TraceRoutes** | Pointer to [**[]V202101beta1TracerouteResult**](V202101beta1TracerouteResult.md) |  | [optional] 

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

### GetTraceRoutes

`func (o *V202101beta1GetTraceForTestResponse) GetTraceRoutes() []V202101beta1TracerouteResult`

GetTraceRoutes returns the TraceRoutes field if non-nil, zero value otherwise.

### GetTraceRoutesOk

`func (o *V202101beta1GetTraceForTestResponse) GetTraceRoutesOk() (*[]V202101beta1TracerouteResult, bool)`

GetTraceRoutesOk returns a tuple with the TraceRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceRoutes

`func (o *V202101beta1GetTraceForTestResponse) SetTraceRoutes(v []V202101beta1TracerouteResult)`

SetTraceRoutes sets TraceRoutes field to given value.

### HasTraceRoutes

`func (o *V202101beta1GetTraceForTestResponse) HasTraceRoutes() bool`

HasTraceRoutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


