# V202101beta1PatchAgentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | Pointer to [**V202101beta1Agent**](V202101beta1Agent.md) |  | [optional] 
**UpdateMask** | Pointer to **string** |  | [optional] 

## Methods

### NewV202101beta1PatchAgentRequest

`func NewV202101beta1PatchAgentRequest() *V202101beta1PatchAgentRequest`

NewV202101beta1PatchAgentRequest instantiates a new V202101beta1PatchAgentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV202101beta1PatchAgentRequestWithDefaults

`func NewV202101beta1PatchAgentRequestWithDefaults() *V202101beta1PatchAgentRequest`

NewV202101beta1PatchAgentRequestWithDefaults instantiates a new V202101beta1PatchAgentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *V202101beta1PatchAgentRequest) GetAgent() V202101beta1Agent`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *V202101beta1PatchAgentRequest) GetAgentOk() (*V202101beta1Agent, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *V202101beta1PatchAgentRequest) SetAgent(v V202101beta1Agent)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *V202101beta1PatchAgentRequest) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetUpdateMask

`func (o *V202101beta1PatchAgentRequest) GetUpdateMask() string`

GetUpdateMask returns the UpdateMask field if non-nil, zero value otherwise.

### GetUpdateMaskOk

`func (o *V202101beta1PatchAgentRequest) GetUpdateMaskOk() (*string, bool)`

GetUpdateMaskOk returns a tuple with the UpdateMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateMask

`func (o *V202101beta1PatchAgentRequest) SetUpdateMask(v string)`

SetUpdateMask sets UpdateMask field to given value.

### HasUpdateMask

`func (o *V202101beta1PatchAgentRequest) HasUpdateMask() bool`

HasUpdateMask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


