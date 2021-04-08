# V202101beta1Health

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Health** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewV202101beta1Health

`func NewV202101beta1Health() *V202101beta1Health`

NewV202101beta1Health instantiates a new V202101beta1Health object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV202101beta1HealthWithDefaults

`func NewV202101beta1HealthWithDefaults() *V202101beta1Health`

NewV202101beta1HealthWithDefaults instantiates a new V202101beta1Health object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealth

`func (o *V202101beta1Health) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *V202101beta1Health) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *V202101beta1Health) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *V202101beta1Health) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetTime

`func (o *V202101beta1Health) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *V202101beta1Health) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *V202101beta1Health) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *V202101beta1Health) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


