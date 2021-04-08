# V202101beta1ListTestsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tests** | Pointer to [**[]V202101beta1Test**](V202101beta1Test.md) |  | [optional] 
**InvalidTestsCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewV202101beta1ListTestsResponse

`func NewV202101beta1ListTestsResponse() *V202101beta1ListTestsResponse`

NewV202101beta1ListTestsResponse instantiates a new V202101beta1ListTestsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV202101beta1ListTestsResponseWithDefaults

`func NewV202101beta1ListTestsResponseWithDefaults() *V202101beta1ListTestsResponse`

NewV202101beta1ListTestsResponseWithDefaults instantiates a new V202101beta1ListTestsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTests

`func (o *V202101beta1ListTestsResponse) GetTests() []V202101beta1Test`

GetTests returns the Tests field if non-nil, zero value otherwise.

### GetTestsOk

`func (o *V202101beta1ListTestsResponse) GetTestsOk() (*[]V202101beta1Test, bool)`

GetTestsOk returns a tuple with the Tests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTests

`func (o *V202101beta1ListTestsResponse) SetTests(v []V202101beta1Test)`

SetTests sets Tests field to given value.

### HasTests

`func (o *V202101beta1ListTestsResponse) HasTests() bool`

HasTests returns a boolean if a field has been set.

### GetInvalidTestsCount

`func (o *V202101beta1ListTestsResponse) GetInvalidTestsCount() int64`

GetInvalidTestsCount returns the InvalidTestsCount field if non-nil, zero value otherwise.

### GetInvalidTestsCountOk

`func (o *V202101beta1ListTestsResponse) GetInvalidTestsCountOk() (*int64, bool)`

GetInvalidTestsCountOk returns a tuple with the InvalidTestsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidTestsCount

`func (o *V202101beta1ListTestsResponse) SetInvalidTestsCount(v int64)`

SetInvalidTestsCount sets InvalidTestsCount field to given value.

### HasInvalidTestsCount

`func (o *V202101beta1ListTestsResponse) HasInvalidTestsCount() bool`

HasInvalidTestsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


