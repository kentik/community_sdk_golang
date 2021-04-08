# V202101beta1GetTraceForTestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **time.Time** | Start of the time interval for this query. | [optional] 
**EndTime** | Pointer to **time.Time** | End of the time interval for this query. | [optional] 
**AgentIds** | Pointer to **[]string** |  | [optional] 
**TargetIps** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV202101beta1GetTraceForTestRequest

`func NewV202101beta1GetTraceForTestRequest() *V202101beta1GetTraceForTestRequest`

NewV202101beta1GetTraceForTestRequest instantiates a new V202101beta1GetTraceForTestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV202101beta1GetTraceForTestRequestWithDefaults

`func NewV202101beta1GetTraceForTestRequestWithDefaults() *V202101beta1GetTraceForTestRequest`

NewV202101beta1GetTraceForTestRequestWithDefaults instantiates a new V202101beta1GetTraceForTestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V202101beta1GetTraceForTestRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V202101beta1GetTraceForTestRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V202101beta1GetTraceForTestRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V202101beta1GetTraceForTestRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStartTime

`func (o *V202101beta1GetTraceForTestRequest) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *V202101beta1GetTraceForTestRequest) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *V202101beta1GetTraceForTestRequest) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *V202101beta1GetTraceForTestRequest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *V202101beta1GetTraceForTestRequest) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *V202101beta1GetTraceForTestRequest) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *V202101beta1GetTraceForTestRequest) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *V202101beta1GetTraceForTestRequest) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetAgentIds

`func (o *V202101beta1GetTraceForTestRequest) GetAgentIds() []string`

GetAgentIds returns the AgentIds field if non-nil, zero value otherwise.

### GetAgentIdsOk

`func (o *V202101beta1GetTraceForTestRequest) GetAgentIdsOk() (*[]string, bool)`

GetAgentIdsOk returns a tuple with the AgentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentIds

`func (o *V202101beta1GetTraceForTestRequest) SetAgentIds(v []string)`

SetAgentIds sets AgentIds field to given value.

### HasAgentIds

`func (o *V202101beta1GetTraceForTestRequest) HasAgentIds() bool`

HasAgentIds returns a boolean if a field has been set.

### GetTargetIps

`func (o *V202101beta1GetTraceForTestRequest) GetTargetIps() []string`

GetTargetIps returns the TargetIps field if non-nil, zero value otherwise.

### GetTargetIpsOk

`func (o *V202101beta1GetTraceForTestRequest) GetTargetIpsOk() (*[]string, bool)`

GetTargetIpsOk returns a tuple with the TargetIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetIps

`func (o *V202101beta1GetTraceForTestRequest) SetTargetIps(v []string)`

SetTargetIps sets TargetIps field to given value.

### HasTargetIps

`func (o *V202101beta1GetTraceForTestRequest) HasTargetIps() bool`

HasTargetIps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


