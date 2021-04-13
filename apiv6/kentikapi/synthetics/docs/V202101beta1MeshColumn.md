# V202101beta1MeshColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 
**Metrics** | Pointer to [**V202101beta1MeshMetrics**](V202101beta1MeshMetrics.md) |  | [optional] 
**Health** | Pointer to [**[]V202101beta1MeshMetrics**](V202101beta1MeshMetrics.md) |  | [optional] 

## Methods

### NewV202101beta1MeshColumn

`func NewV202101beta1MeshColumn() *V202101beta1MeshColumn`

NewV202101beta1MeshColumn instantiates a new V202101beta1MeshColumn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV202101beta1MeshColumnWithDefaults

`func NewV202101beta1MeshColumnWithDefaults() *V202101beta1MeshColumn`

NewV202101beta1MeshColumnWithDefaults instantiates a new V202101beta1MeshColumn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V202101beta1MeshColumn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V202101beta1MeshColumn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V202101beta1MeshColumn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V202101beta1MeshColumn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *V202101beta1MeshColumn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V202101beta1MeshColumn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V202101beta1MeshColumn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V202101beta1MeshColumn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAlias

`func (o *V202101beta1MeshColumn) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *V202101beta1MeshColumn) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *V202101beta1MeshColumn) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *V202101beta1MeshColumn) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetTarget

`func (o *V202101beta1MeshColumn) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *V202101beta1MeshColumn) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *V202101beta1MeshColumn) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *V202101beta1MeshColumn) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetMetrics

`func (o *V202101beta1MeshColumn) GetMetrics() V202101beta1MeshMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *V202101beta1MeshColumn) GetMetricsOk() (*V202101beta1MeshMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *V202101beta1MeshColumn) SetMetrics(v V202101beta1MeshMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *V202101beta1MeshColumn) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetHealth

`func (o *V202101beta1MeshColumn) GetHealth() []V202101beta1MeshMetrics`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *V202101beta1MeshColumn) GetHealthOk() (*[]V202101beta1MeshMetrics, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *V202101beta1MeshColumn) SetHealth(v []V202101beta1MeshMetrics)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *V202101beta1MeshColumn) HasHealth() bool`

HasHealth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


