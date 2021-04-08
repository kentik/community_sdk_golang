# V202101beta1TestHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestId** | Pointer to **string** |  | [optional] 
**Tasks** | Pointer to [**[]V202101beta1TaskHealth**](V202101beta1TaskHealth.md) |  | [optional] 
**OverallHealth** | Pointer to [**V202101beta1Health**](V202101beta1Health.md) |  | [optional] 
**HealthTs** | Pointer to [**[]V202101beta1Health**](V202101beta1Health.md) |  | [optional] 

## Methods

### NewV202101beta1TestHealth

`func NewV202101beta1TestHealth() *V202101beta1TestHealth`

NewV202101beta1TestHealth instantiates a new V202101beta1TestHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV202101beta1TestHealthWithDefaults

`func NewV202101beta1TestHealthWithDefaults() *V202101beta1TestHealth`

NewV202101beta1TestHealthWithDefaults instantiates a new V202101beta1TestHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestId

`func (o *V202101beta1TestHealth) GetTestId() string`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *V202101beta1TestHealth) GetTestIdOk() (*string, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *V202101beta1TestHealth) SetTestId(v string)`

SetTestId sets TestId field to given value.

### HasTestId

`func (o *V202101beta1TestHealth) HasTestId() bool`

HasTestId returns a boolean if a field has been set.

### GetTasks

`func (o *V202101beta1TestHealth) GetTasks() []V202101beta1TaskHealth`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *V202101beta1TestHealth) GetTasksOk() (*[]V202101beta1TaskHealth, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *V202101beta1TestHealth) SetTasks(v []V202101beta1TaskHealth)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *V202101beta1TestHealth) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetOverallHealth

`func (o *V202101beta1TestHealth) GetOverallHealth() V202101beta1Health`

GetOverallHealth returns the OverallHealth field if non-nil, zero value otherwise.

### GetOverallHealthOk

`func (o *V202101beta1TestHealth) GetOverallHealthOk() (*V202101beta1Health, bool)`

GetOverallHealthOk returns a tuple with the OverallHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallHealth

`func (o *V202101beta1TestHealth) SetOverallHealth(v V202101beta1Health)`

SetOverallHealth sets OverallHealth field to given value.

### HasOverallHealth

`func (o *V202101beta1TestHealth) HasOverallHealth() bool`

HasOverallHealth returns a boolean if a field has been set.

### GetHealthTs

`func (o *V202101beta1TestHealth) GetHealthTs() []V202101beta1Health`

GetHealthTs returns the HealthTs field if non-nil, zero value otherwise.

### GetHealthTsOk

`func (o *V202101beta1TestHealth) GetHealthTsOk() (*[]V202101beta1Health, bool)`

GetHealthTsOk returns a tuple with the HealthTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthTs

`func (o *V202101beta1TestHealth) SetHealthTs(v []V202101beta1Health)`

SetHealthTs sets HealthTs field to given value.

### HasHealthTs

`func (o *V202101beta1TestHealth) HasHealthTs() bool`

HasHealthTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


