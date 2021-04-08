# V202101beta1SetTestStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**V202101beta1TestStatus**](V202101beta1TestStatus.md) |  | [optional] [default to V202101BETA1TESTSTATUS_UNSPECIFIED]

## Methods

### NewV202101beta1SetTestStatusRequest

`func NewV202101beta1SetTestStatusRequest() *V202101beta1SetTestStatusRequest`

NewV202101beta1SetTestStatusRequest instantiates a new V202101beta1SetTestStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV202101beta1SetTestStatusRequestWithDefaults

`func NewV202101beta1SetTestStatusRequestWithDefaults() *V202101beta1SetTestStatusRequest`

NewV202101beta1SetTestStatusRequestWithDefaults instantiates a new V202101beta1SetTestStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V202101beta1SetTestStatusRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V202101beta1SetTestStatusRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V202101beta1SetTestStatusRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V202101beta1SetTestStatusRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *V202101beta1SetTestStatusRequest) GetStatus() V202101beta1TestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V202101beta1SetTestStatusRequest) GetStatusOk() (*V202101beta1TestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V202101beta1SetTestStatusRequest) SetStatus(v V202101beta1TestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V202101beta1SetTestStatusRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


