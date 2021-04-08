# V202101beta1Trace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hops** | Pointer to [**[]V202101beta1TraceHop**](V202101beta1TraceHop.md) |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 
**Ips** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV202101beta1Trace

`func NewV202101beta1Trace() *V202101beta1Trace`

NewV202101beta1Trace instantiates a new V202101beta1Trace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV202101beta1TraceWithDefaults

`func NewV202101beta1TraceWithDefaults() *V202101beta1Trace`

NewV202101beta1TraceWithDefaults instantiates a new V202101beta1Trace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHops

`func (o *V202101beta1Trace) GetHops() []V202101beta1TraceHop`

GetHops returns the Hops field if non-nil, zero value otherwise.

### GetHopsOk

`func (o *V202101beta1Trace) GetHopsOk() (*[]V202101beta1TraceHop, bool)`

GetHopsOk returns a tuple with the Hops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHops

`func (o *V202101beta1Trace) SetHops(v []V202101beta1TraceHop)`

SetHops sets Hops field to given value.

### HasHops

`func (o *V202101beta1Trace) HasHops() bool`

HasHops returns a boolean if a field has been set.

### GetTarget

`func (o *V202101beta1Trace) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *V202101beta1Trace) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *V202101beta1Trace) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *V202101beta1Trace) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetIps

`func (o *V202101beta1Trace) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *V202101beta1Trace) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *V202101beta1Trace) SetIps(v []string)`

SetIps sets Ips field to given value.

### HasIps

`func (o *V202101beta1Trace) HasIps() bool`

HasIps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


