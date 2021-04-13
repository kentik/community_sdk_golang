# V202101beta1TaskHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Task** | Pointer to [**V202101beta1Task**](V202101beta1Task.md) |  | [optional] 
**Agents** | Pointer to [**[]V202101beta1AgentHealth**](V202101beta1AgentHealth.md) |  | [optional] 
**OverallHealth** | Pointer to [**V202101beta1Health**](V202101beta1Health.md) |  | [optional] 
**TargetAgent** | Pointer to [**V202101beta1Agent**](V202101beta1Agent.md) |  | [optional] 

## Methods

### NewV202101beta1TaskHealth

`func NewV202101beta1TaskHealth() *V202101beta1TaskHealth`

NewV202101beta1TaskHealth instantiates a new V202101beta1TaskHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV202101beta1TaskHealthWithDefaults

`func NewV202101beta1TaskHealthWithDefaults() *V202101beta1TaskHealth`

NewV202101beta1TaskHealthWithDefaults instantiates a new V202101beta1TaskHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTask

`func (o *V202101beta1TaskHealth) GetTask() V202101beta1Task`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *V202101beta1TaskHealth) GetTaskOk() (*V202101beta1Task, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *V202101beta1TaskHealth) SetTask(v V202101beta1Task)`

SetTask sets Task field to given value.

### HasTask

`func (o *V202101beta1TaskHealth) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetAgents

`func (o *V202101beta1TaskHealth) GetAgents() []V202101beta1AgentHealth`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *V202101beta1TaskHealth) GetAgentsOk() (*[]V202101beta1AgentHealth, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *V202101beta1TaskHealth) SetAgents(v []V202101beta1AgentHealth)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *V202101beta1TaskHealth) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### GetOverallHealth

`func (o *V202101beta1TaskHealth) GetOverallHealth() V202101beta1Health`

GetOverallHealth returns the OverallHealth field if non-nil, zero value otherwise.

### GetOverallHealthOk

`func (o *V202101beta1TaskHealth) GetOverallHealthOk() (*V202101beta1Health, bool)`

GetOverallHealthOk returns a tuple with the OverallHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallHealth

`func (o *V202101beta1TaskHealth) SetOverallHealth(v V202101beta1Health)`

SetOverallHealth sets OverallHealth field to given value.

### HasOverallHealth

`func (o *V202101beta1TaskHealth) HasOverallHealth() bool`

HasOverallHealth returns a boolean if a field has been set.

### GetTargetAgent

`func (o *V202101beta1TaskHealth) GetTargetAgent() V202101beta1Agent`

GetTargetAgent returns the TargetAgent field if non-nil, zero value otherwise.

### GetTargetAgentOk

`func (o *V202101beta1TaskHealth) GetTargetAgentOk() (*V202101beta1Agent, bool)`

GetTargetAgentOk returns a tuple with the TargetAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAgent

`func (o *V202101beta1TaskHealth) SetTargetAgent(v V202101beta1Agent)`

SetTargetAgent sets TargetAgent field to given value.

### HasTargetAgent

`func (o *V202101beta1TaskHealth) HasTargetAgent() bool`

HasTargetAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


