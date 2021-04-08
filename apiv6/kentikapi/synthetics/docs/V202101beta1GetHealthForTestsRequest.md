# V202101beta1GetHealthForTestsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]string** | List of ids of the tests to get health for. | [optional] 
**StartTime** | Pointer to **time.Time** | Start of the time interval for this query. | [optional] 
**EndTime** | Pointer to **time.Time** | End of the time interval for this query. | [optional] 
**AgentIds** | Pointer to **[]string** |  | [optional] 
**TaskIds** | Pointer to **[]string** | Optionally only look at a subset of tasks -- this lets you limit targets. | [optional] 

## Methods

### NewV202101beta1GetHealthForTestsRequest

`func NewV202101beta1GetHealthForTestsRequest() *V202101beta1GetHealthForTestsRequest`

NewV202101beta1GetHealthForTestsRequest instantiates a new V202101beta1GetHealthForTestsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV202101beta1GetHealthForTestsRequestWithDefaults

`func NewV202101beta1GetHealthForTestsRequestWithDefaults() *V202101beta1GetHealthForTestsRequest`

NewV202101beta1GetHealthForTestsRequestWithDefaults instantiates a new V202101beta1GetHealthForTestsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *V202101beta1GetHealthForTestsRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *V202101beta1GetHealthForTestsRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *V202101beta1GetHealthForTestsRequest) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *V202101beta1GetHealthForTestsRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetStartTime

`func (o *V202101beta1GetHealthForTestsRequest) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *V202101beta1GetHealthForTestsRequest) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *V202101beta1GetHealthForTestsRequest) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *V202101beta1GetHealthForTestsRequest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *V202101beta1GetHealthForTestsRequest) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *V202101beta1GetHealthForTestsRequest) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *V202101beta1GetHealthForTestsRequest) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *V202101beta1GetHealthForTestsRequest) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetAgentIds

`func (o *V202101beta1GetHealthForTestsRequest) GetAgentIds() []string`

GetAgentIds returns the AgentIds field if non-nil, zero value otherwise.

### GetAgentIdsOk

`func (o *V202101beta1GetHealthForTestsRequest) GetAgentIdsOk() (*[]string, bool)`

GetAgentIdsOk returns a tuple with the AgentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentIds

`func (o *V202101beta1GetHealthForTestsRequest) SetAgentIds(v []string)`

SetAgentIds sets AgentIds field to given value.

### HasAgentIds

`func (o *V202101beta1GetHealthForTestsRequest) HasAgentIds() bool`

HasAgentIds returns a boolean if a field has been set.

### GetTaskIds

`func (o *V202101beta1GetHealthForTestsRequest) GetTaskIds() []string`

GetTaskIds returns the TaskIds field if non-nil, zero value otherwise.

### GetTaskIdsOk

`func (o *V202101beta1GetHealthForTestsRequest) GetTaskIdsOk() (*[]string, bool)`

GetTaskIdsOk returns a tuple with the TaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskIds

`func (o *V202101beta1GetHealthForTestsRequest) SetTaskIds(v []string)`

SetTaskIds sets TaskIds field to given value.

### HasTaskIds

`func (o *V202101beta1GetHealthForTestsRequest) HasTaskIds() bool`

HasTaskIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


