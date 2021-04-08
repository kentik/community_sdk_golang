# V202101beta1AgentHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | Pointer to [**V202101beta1Agent**](V202101beta1Agent.md) |  | [optional] 
**Health** | Pointer to [**[]V202101beta1HealthMoment**](V202101beta1HealthMoment.md) |  | [optional] 
**OverallHealth** | Pointer to [**V202101beta1Health**](V202101beta1Health.md) |  | [optional] 

## Methods

### NewV202101beta1AgentHealth

`func NewV202101beta1AgentHealth() *V202101beta1AgentHealth`

NewV202101beta1AgentHealth instantiates a new V202101beta1AgentHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV202101beta1AgentHealthWithDefaults

`func NewV202101beta1AgentHealthWithDefaults() *V202101beta1AgentHealth`

NewV202101beta1AgentHealthWithDefaults instantiates a new V202101beta1AgentHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *V202101beta1AgentHealth) GetAgent() V202101beta1Agent`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *V202101beta1AgentHealth) GetAgentOk() (*V202101beta1Agent, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *V202101beta1AgentHealth) SetAgent(v V202101beta1Agent)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *V202101beta1AgentHealth) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetHealth

`func (o *V202101beta1AgentHealth) GetHealth() []V202101beta1HealthMoment`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *V202101beta1AgentHealth) GetHealthOk() (*[]V202101beta1HealthMoment, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *V202101beta1AgentHealth) SetHealth(v []V202101beta1HealthMoment)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *V202101beta1AgentHealth) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetOverallHealth

`func (o *V202101beta1AgentHealth) GetOverallHealth() V202101beta1Health`

GetOverallHealth returns the OverallHealth field if non-nil, zero value otherwise.

### GetOverallHealthOk

`func (o *V202101beta1AgentHealth) GetOverallHealthOk() (*V202101beta1Health, bool)`

GetOverallHealthOk returns a tuple with the OverallHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallHealth

`func (o *V202101beta1AgentHealth) SetOverallHealth(v V202101beta1Health)`

SetOverallHealth sets OverallHealth field to given value.

### HasOverallHealth

`func (o *V202101beta1AgentHealth) HasOverallHealth() bool`

HasOverallHealth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


