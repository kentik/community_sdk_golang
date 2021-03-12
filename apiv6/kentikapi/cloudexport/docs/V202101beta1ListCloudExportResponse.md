# V202101beta1ListCloudExportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exports** | Pointer to [**[]V202101beta1CloudExport**](V202101beta1CloudExport.md) |  | [optional] 
**InvalidExportsCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewV202101beta1ListCloudExportResponse

`func NewV202101beta1ListCloudExportResponse() *V202101beta1ListCloudExportResponse`

NewV202101beta1ListCloudExportResponse instantiates a new V202101beta1ListCloudExportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV202101beta1ListCloudExportResponseWithDefaults

`func NewV202101beta1ListCloudExportResponseWithDefaults() *V202101beta1ListCloudExportResponse`

NewV202101beta1ListCloudExportResponseWithDefaults instantiates a new V202101beta1ListCloudExportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExports

`func (o *V202101beta1ListCloudExportResponse) GetExports() []V202101beta1CloudExport`

GetExports returns the Exports field if non-nil, zero value otherwise.

### GetExportsOk

`func (o *V202101beta1ListCloudExportResponse) GetExportsOk() (*[]V202101beta1CloudExport, bool)`

GetExportsOk returns a tuple with the Exports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExports

`func (o *V202101beta1ListCloudExportResponse) SetExports(v []V202101beta1CloudExport)`

SetExports sets Exports field to given value.

### HasExports

`func (o *V202101beta1ListCloudExportResponse) HasExports() bool`

HasExports returns a boolean if a field has been set.

### GetInvalidExportsCount

`func (o *V202101beta1ListCloudExportResponse) GetInvalidExportsCount() int64`

GetInvalidExportsCount returns the InvalidExportsCount field if non-nil, zero value otherwise.

### GetInvalidExportsCountOk

`func (o *V202101beta1ListCloudExportResponse) GetInvalidExportsCountOk() (*int64, bool)`

GetInvalidExportsCountOk returns a tuple with the InvalidExportsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidExportsCount

`func (o *V202101beta1ListCloudExportResponse) SetInvalidExportsCount(v int64)`

SetInvalidExportsCount sets InvalidExportsCount field to given value.

### HasInvalidExportsCount

`func (o *V202101beta1ListCloudExportResponse) HasInvalidExportsCount() bool`

HasInvalidExportsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


