# V202101beta1Test

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**DeviceId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**V202101beta1TestStatus**](V202101beta1TestStatus.md) |  | [optional] [default to V202101BETA1TESTSTATUS_UNSPECIFIED]
**Settings** | Pointer to [**V202101beta1TestSettings**](V202101beta1TestSettings.md) |  | [optional] 
**ExpiresOn** | Pointer to **time.Time** |  | [optional] 
**Cdate** | Pointer to **time.Time** |  | [optional] 
**Edate** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to [**V202101beta1UserInfo**](V202101beta1UserInfo.md) |  | [optional] 
**LastUpdatedBy** | Pointer to [**V202101beta1UserInfo**](V202101beta1UserInfo.md) |  | [optional] 

## Methods

### NewV202101beta1Test

`func NewV202101beta1Test() *V202101beta1Test`

NewV202101beta1Test instantiates a new V202101beta1Test object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV202101beta1TestWithDefaults

`func NewV202101beta1TestWithDefaults() *V202101beta1Test`

NewV202101beta1TestWithDefaults instantiates a new V202101beta1Test object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V202101beta1Test) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V202101beta1Test) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V202101beta1Test) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V202101beta1Test) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *V202101beta1Test) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V202101beta1Test) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V202101beta1Test) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V202101beta1Test) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *V202101beta1Test) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V202101beta1Test) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V202101beta1Test) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V202101beta1Test) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDeviceId

`func (o *V202101beta1Test) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *V202101beta1Test) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *V202101beta1Test) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *V202101beta1Test) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetStatus

`func (o *V202101beta1Test) GetStatus() V202101beta1TestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V202101beta1Test) GetStatusOk() (*V202101beta1TestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V202101beta1Test) SetStatus(v V202101beta1TestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V202101beta1Test) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSettings

`func (o *V202101beta1Test) GetSettings() V202101beta1TestSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *V202101beta1Test) GetSettingsOk() (*V202101beta1TestSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *V202101beta1Test) SetSettings(v V202101beta1TestSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *V202101beta1Test) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetExpiresOn

`func (o *V202101beta1Test) GetExpiresOn() time.Time`

GetExpiresOn returns the ExpiresOn field if non-nil, zero value otherwise.

### GetExpiresOnOk

`func (o *V202101beta1Test) GetExpiresOnOk() (*time.Time, bool)`

GetExpiresOnOk returns a tuple with the ExpiresOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresOn

`func (o *V202101beta1Test) SetExpiresOn(v time.Time)`

SetExpiresOn sets ExpiresOn field to given value.

### HasExpiresOn

`func (o *V202101beta1Test) HasExpiresOn() bool`

HasExpiresOn returns a boolean if a field has been set.

### GetCdate

`func (o *V202101beta1Test) GetCdate() time.Time`

GetCdate returns the Cdate field if non-nil, zero value otherwise.

### GetCdateOk

`func (o *V202101beta1Test) GetCdateOk() (*time.Time, bool)`

GetCdateOk returns a tuple with the Cdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdate

`func (o *V202101beta1Test) SetCdate(v time.Time)`

SetCdate sets Cdate field to given value.

### HasCdate

`func (o *V202101beta1Test) HasCdate() bool`

HasCdate returns a boolean if a field has been set.

### GetEdate

`func (o *V202101beta1Test) GetEdate() time.Time`

GetEdate returns the Edate field if non-nil, zero value otherwise.

### GetEdateOk

`func (o *V202101beta1Test) GetEdateOk() (*time.Time, bool)`

GetEdateOk returns a tuple with the Edate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdate

`func (o *V202101beta1Test) SetEdate(v time.Time)`

SetEdate sets Edate field to given value.

### HasEdate

`func (o *V202101beta1Test) HasEdate() bool`

HasEdate returns a boolean if a field has been set.

### GetCreatedBy

`func (o *V202101beta1Test) GetCreatedBy() V202101beta1UserInfo`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *V202101beta1Test) GetCreatedByOk() (*V202101beta1UserInfo, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *V202101beta1Test) SetCreatedBy(v V202101beta1UserInfo)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *V202101beta1Test) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *V202101beta1Test) GetLastUpdatedBy() V202101beta1UserInfo`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *V202101beta1Test) GetLastUpdatedByOk() (*V202101beta1UserInfo, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *V202101beta1Test) SetLastUpdatedBy(v V202101beta1UserInfo)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *V202101beta1Test) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


