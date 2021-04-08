# V202101beta1Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**TestId** | Pointer to **string** |  | [optional] 
**DeviceId** | Pointer to **string** |  | [optional] 
**State** | Pointer to [**V202101beta1TaskState**](V202101beta1TaskState.md) |  | [optional] [default to V202101BETA1TASKSTATE_UNSPECIFIED]
**Status** | Pointer to **string** |  | [optional] 
**Family** | Pointer to [**V202101beta1IPFamily**](V202101beta1IPFamily.md) |  | [optional] [default to V202101BETA1IPFAMILY_UNSPECIFIED]
**Ping** | Pointer to [**V202101beta1PingTaskDefinition**](V202101beta1PingTaskDefinition.md) |  | [optional] 
**Traceroute** | Pointer to [**V202101beta1TraceTaskDefinition**](V202101beta1TraceTaskDefinition.md) |  | [optional] 
**Http** | Pointer to [**V202101beta1HTTPTaskDefinition**](V202101beta1HTTPTaskDefinition.md) |  | [optional] 
**Knock** | Pointer to [**V202101beta1KnockTaskDefinition**](V202101beta1KnockTaskDefinition.md) |  | [optional] 
**Dns** | Pointer to [**V202101beta1DNSTaskDefinition**](V202101beta1DNSTaskDefinition.md) |  | [optional] 
**Shake** | Pointer to [**V202101beta1ShakeTaskDefinition**](V202101beta1ShakeTaskDefinition.md) |  | [optional] 

## Methods

### NewV202101beta1Task

`func NewV202101beta1Task() *V202101beta1Task`

NewV202101beta1Task instantiates a new V202101beta1Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV202101beta1TaskWithDefaults

`func NewV202101beta1TaskWithDefaults() *V202101beta1Task`

NewV202101beta1TaskWithDefaults instantiates a new V202101beta1Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V202101beta1Task) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V202101beta1Task) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V202101beta1Task) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V202101beta1Task) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTestId

`func (o *V202101beta1Task) GetTestId() string`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *V202101beta1Task) GetTestIdOk() (*string, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *V202101beta1Task) SetTestId(v string)`

SetTestId sets TestId field to given value.

### HasTestId

`func (o *V202101beta1Task) HasTestId() bool`

HasTestId returns a boolean if a field has been set.

### GetDeviceId

`func (o *V202101beta1Task) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *V202101beta1Task) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *V202101beta1Task) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *V202101beta1Task) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetState

`func (o *V202101beta1Task) GetState() V202101beta1TaskState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V202101beta1Task) GetStateOk() (*V202101beta1TaskState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V202101beta1Task) SetState(v V202101beta1TaskState)`

SetState sets State field to given value.

### HasState

`func (o *V202101beta1Task) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatus

`func (o *V202101beta1Task) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V202101beta1Task) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V202101beta1Task) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V202101beta1Task) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFamily

`func (o *V202101beta1Task) GetFamily() V202101beta1IPFamily`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *V202101beta1Task) GetFamilyOk() (*V202101beta1IPFamily, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *V202101beta1Task) SetFamily(v V202101beta1IPFamily)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *V202101beta1Task) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetPing

`func (o *V202101beta1Task) GetPing() V202101beta1PingTaskDefinition`

GetPing returns the Ping field if non-nil, zero value otherwise.

### GetPingOk

`func (o *V202101beta1Task) GetPingOk() (*V202101beta1PingTaskDefinition, bool)`

GetPingOk returns a tuple with the Ping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPing

`func (o *V202101beta1Task) SetPing(v V202101beta1PingTaskDefinition)`

SetPing sets Ping field to given value.

### HasPing

`func (o *V202101beta1Task) HasPing() bool`

HasPing returns a boolean if a field has been set.

### GetTraceroute

`func (o *V202101beta1Task) GetTraceroute() V202101beta1TraceTaskDefinition`

GetTraceroute returns the Traceroute field if non-nil, zero value otherwise.

### GetTracerouteOk

`func (o *V202101beta1Task) GetTracerouteOk() (*V202101beta1TraceTaskDefinition, bool)`

GetTracerouteOk returns a tuple with the Traceroute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceroute

`func (o *V202101beta1Task) SetTraceroute(v V202101beta1TraceTaskDefinition)`

SetTraceroute sets Traceroute field to given value.

### HasTraceroute

`func (o *V202101beta1Task) HasTraceroute() bool`

HasTraceroute returns a boolean if a field has been set.

### GetHttp

`func (o *V202101beta1Task) GetHttp() V202101beta1HTTPTaskDefinition`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *V202101beta1Task) GetHttpOk() (*V202101beta1HTTPTaskDefinition, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *V202101beta1Task) SetHttp(v V202101beta1HTTPTaskDefinition)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *V202101beta1Task) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetKnock

`func (o *V202101beta1Task) GetKnock() V202101beta1KnockTaskDefinition`

GetKnock returns the Knock field if non-nil, zero value otherwise.

### GetKnockOk

`func (o *V202101beta1Task) GetKnockOk() (*V202101beta1KnockTaskDefinition, bool)`

GetKnockOk returns a tuple with the Knock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnock

`func (o *V202101beta1Task) SetKnock(v V202101beta1KnockTaskDefinition)`

SetKnock sets Knock field to given value.

### HasKnock

`func (o *V202101beta1Task) HasKnock() bool`

HasKnock returns a boolean if a field has been set.

### GetDns

`func (o *V202101beta1Task) GetDns() V202101beta1DNSTaskDefinition`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *V202101beta1Task) GetDnsOk() (*V202101beta1DNSTaskDefinition, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *V202101beta1Task) SetDns(v V202101beta1DNSTaskDefinition)`

SetDns sets Dns field to given value.

### HasDns

`func (o *V202101beta1Task) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetShake

`func (o *V202101beta1Task) GetShake() V202101beta1ShakeTaskDefinition`

GetShake returns the Shake field if non-nil, zero value otherwise.

### GetShakeOk

`func (o *V202101beta1Task) GetShakeOk() (*V202101beta1ShakeTaskDefinition, bool)`

GetShakeOk returns a tuple with the Shake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShake

`func (o *V202101beta1Task) SetShake(v V202101beta1ShakeTaskDefinition)`

SetShake sets Shake field to given value.

### HasShake

`func (o *V202101beta1Task) HasShake() bool`

HasShake returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


