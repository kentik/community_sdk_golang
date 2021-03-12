# V202101beta1PatchCloudExportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Export** | Pointer to [**V202101beta1CloudExport**](V202101beta1CloudExport.md) |  | [optional] 
**UpdateMask** | Pointer to **string** |  | [optional] 

## Methods

### NewV202101beta1PatchCloudExportRequest

`func NewV202101beta1PatchCloudExportRequest() *V202101beta1PatchCloudExportRequest`

NewV202101beta1PatchCloudExportRequest instantiates a new V202101beta1PatchCloudExportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV202101beta1PatchCloudExportRequestWithDefaults

`func NewV202101beta1PatchCloudExportRequestWithDefaults() *V202101beta1PatchCloudExportRequest`

NewV202101beta1PatchCloudExportRequestWithDefaults instantiates a new V202101beta1PatchCloudExportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExport

`func (o *V202101beta1PatchCloudExportRequest) GetExport() V202101beta1CloudExport`

GetExport returns the Export field if non-nil, zero value otherwise.

### GetExportOk

`func (o *V202101beta1PatchCloudExportRequest) GetExportOk() (*V202101beta1CloudExport, bool)`

GetExportOk returns a tuple with the Export field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExport

`func (o *V202101beta1PatchCloudExportRequest) SetExport(v V202101beta1CloudExport)`

SetExport sets Export field to given value.

### HasExport

`func (o *V202101beta1PatchCloudExportRequest) HasExport() bool`

HasExport returns a boolean if a field has been set.

### GetUpdateMask

`func (o *V202101beta1PatchCloudExportRequest) GetUpdateMask() string`

GetUpdateMask returns the UpdateMask field if non-nil, zero value otherwise.

### GetUpdateMaskOk

`func (o *V202101beta1PatchCloudExportRequest) GetUpdateMaskOk() (*string, bool)`

GetUpdateMaskOk returns a tuple with the UpdateMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateMask

`func (o *V202101beta1PatchCloudExportRequest) SetUpdateMask(v string)`

SetUpdateMask sets UpdateMask field to given value.

### HasUpdateMask

`func (o *V202101beta1PatchCloudExportRequest) HasUpdateMask() bool`

HasUpdateMask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


