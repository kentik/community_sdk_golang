# V202101beta1ListAgentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agents** | Pointer to [**[]V202101beta1Agent**](V202101beta1Agent.md) |  | [optional] 
**InvalidAgentsCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewV202101beta1ListAgentsResponse

`func NewV202101beta1ListAgentsResponse() *V202101beta1ListAgentsResponse`

NewV202101beta1ListAgentsResponse instantiates a new V202101beta1ListAgentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV202101beta1ListAgentsResponseWithDefaults

`func NewV202101beta1ListAgentsResponseWithDefaults() *V202101beta1ListAgentsResponse`

NewV202101beta1ListAgentsResponseWithDefaults instantiates a new V202101beta1ListAgentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgents

`func (o *V202101beta1ListAgentsResponse) GetAgents() []V202101beta1Agent`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *V202101beta1ListAgentsResponse) GetAgentsOk() (*[]V202101beta1Agent, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *V202101beta1ListAgentsResponse) SetAgents(v []V202101beta1Agent)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *V202101beta1ListAgentsResponse) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### GetInvalidAgentsCount

`func (o *V202101beta1ListAgentsResponse) GetInvalidAgentsCount() int64`

GetInvalidAgentsCount returns the InvalidAgentsCount field if non-nil, zero value otherwise.

### GetInvalidAgentsCountOk

`func (o *V202101beta1ListAgentsResponse) GetInvalidAgentsCountOk() (*int64, bool)`

GetInvalidAgentsCountOk returns a tuple with the InvalidAgentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidAgentsCount

`func (o *V202101beta1ListAgentsResponse) SetInvalidAgentsCount(v int64)`

SetInvalidAgentsCount sets InvalidAgentsCount field to given value.

### HasInvalidAgentsCount

`func (o *V202101beta1ListAgentsResponse) HasInvalidAgentsCount() bool`

HasInvalidAgentsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


