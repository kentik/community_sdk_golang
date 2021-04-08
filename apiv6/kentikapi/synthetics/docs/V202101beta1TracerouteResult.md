# V202101beta1TracerouteResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **time.Time** |  | [optional] 
**Traces** | Pointer to [**[]V202101beta1Trace**](V202101beta1Trace.md) |  | [optional] 

## Methods

### NewV202101beta1TracerouteResult

`func NewV202101beta1TracerouteResult() *V202101beta1TracerouteResult`

NewV202101beta1TracerouteResult instantiates a new V202101beta1TracerouteResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV202101beta1TracerouteResultWithDefaults

`func NewV202101beta1TracerouteResultWithDefaults() *V202101beta1TracerouteResult`

NewV202101beta1TracerouteResultWithDefaults instantiates a new V202101beta1TracerouteResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *V202101beta1TracerouteResult) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *V202101beta1TracerouteResult) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *V202101beta1TracerouteResult) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *V202101beta1TracerouteResult) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTraces

`func (o *V202101beta1TracerouteResult) GetTraces() []V202101beta1Trace`

GetTraces returns the Traces field if non-nil, zero value otherwise.

### GetTracesOk

`func (o *V202101beta1TracerouteResult) GetTracesOk() (*[]V202101beta1Trace, bool)`

GetTracesOk returns a tuple with the Traces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraces

`func (o *V202101beta1TracerouteResult) SetTraces(v []V202101beta1Trace)`

SetTraces sets Traces field to given value.

### HasTraces

`func (o *V202101beta1TracerouteResult) HasTraces() bool`

HasTraces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


