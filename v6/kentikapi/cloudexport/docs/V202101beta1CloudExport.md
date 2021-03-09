# V202101beta1CloudExport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The internal cloud export identifier. This is Read-only and assigned by Kentik. | [optional] 
**Type** | Pointer to [**V202101beta1CloudExportType**](V202101beta1CloudExportType.md) |  | [optional] [default to UNSPECIFIED]
**Enabled** | Pointer to **bool** | Whether this task is enabled and intended to run, or disabled. | [optional] 
**Name** | Pointer to **string** | A short name for this export. | [optional] 
**Description** | Pointer to **string** | An optional, longer description. | [optional] 
**ApiRoot** | Pointer to **string** |  | [optional] 
**FlowDest** | Pointer to **string** |  | [optional] 
**PlanId** | Pointer to **string** | The identifier of the Kentik plan associated with this task. | [optional] 
**CloudProvider** | Pointer to **string** |  | [optional] 
**Aws** | Pointer to [**V202101beta1AwsProperties**](V202101beta1AwsProperties.md) |  | [optional] 
**Azure** | Pointer to [**V202101beta1AzureProperties**](V202101beta1AzureProperties.md) |  | [optional] 
**Gce** | Pointer to [**V202101beta1GceProperties**](V202101beta1GceProperties.md) |  | [optional] 
**Ibm** | Pointer to [**V202101beta1IbmProperties**](V202101beta1IbmProperties.md) |  | [optional] 
**Bgp** | Pointer to [**V202101beta1BgpProperties**](V202101beta1BgpProperties.md) |  | [optional] 
**CurrentStatus** | Pointer to [**CloudExportv202101beta1Status**](CloudExportv202101beta1Status.md) |  | [optional] 

## Methods

### NewV202101beta1CloudExport

`func NewV202101beta1CloudExport() *V202101beta1CloudExport`

NewV202101beta1CloudExport instantiates a new V202101beta1CloudExport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV202101beta1CloudExportWithDefaults

`func NewV202101beta1CloudExportWithDefaults() *V202101beta1CloudExport`

NewV202101beta1CloudExportWithDefaults instantiates a new V202101beta1CloudExport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V202101beta1CloudExport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V202101beta1CloudExport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V202101beta1CloudExport) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V202101beta1CloudExport) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *V202101beta1CloudExport) GetType() V202101beta1CloudExportType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V202101beta1CloudExport) GetTypeOk() (*V202101beta1CloudExportType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V202101beta1CloudExport) SetType(v V202101beta1CloudExportType)`

SetType sets Type field to given value.

### HasType

`func (o *V202101beta1CloudExport) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEnabled

`func (o *V202101beta1CloudExport) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *V202101beta1CloudExport) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *V202101beta1CloudExport) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *V202101beta1CloudExport) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *V202101beta1CloudExport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V202101beta1CloudExport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V202101beta1CloudExport) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V202101beta1CloudExport) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *V202101beta1CloudExport) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V202101beta1CloudExport) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V202101beta1CloudExport) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V202101beta1CloudExport) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetApiRoot

`func (o *V202101beta1CloudExport) GetApiRoot() string`

GetApiRoot returns the ApiRoot field if non-nil, zero value otherwise.

### GetApiRootOk

`func (o *V202101beta1CloudExport) GetApiRootOk() (*string, bool)`

GetApiRootOk returns a tuple with the ApiRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiRoot

`func (o *V202101beta1CloudExport) SetApiRoot(v string)`

SetApiRoot sets ApiRoot field to given value.

### HasApiRoot

`func (o *V202101beta1CloudExport) HasApiRoot() bool`

HasApiRoot returns a boolean if a field has been set.

### GetFlowDest

`func (o *V202101beta1CloudExport) GetFlowDest() string`

GetFlowDest returns the FlowDest field if non-nil, zero value otherwise.

### GetFlowDestOk

`func (o *V202101beta1CloudExport) GetFlowDestOk() (*string, bool)`

GetFlowDestOk returns a tuple with the FlowDest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDest

`func (o *V202101beta1CloudExport) SetFlowDest(v string)`

SetFlowDest sets FlowDest field to given value.

### HasFlowDest

`func (o *V202101beta1CloudExport) HasFlowDest() bool`

HasFlowDest returns a boolean if a field has been set.

### GetPlanId

`func (o *V202101beta1CloudExport) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *V202101beta1CloudExport) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *V202101beta1CloudExport) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *V202101beta1CloudExport) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetCloudProvider

`func (o *V202101beta1CloudExport) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *V202101beta1CloudExport) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *V202101beta1CloudExport) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *V202101beta1CloudExport) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetAws

`func (o *V202101beta1CloudExport) GetAws() V202101beta1AwsProperties`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *V202101beta1CloudExport) GetAwsOk() (*V202101beta1AwsProperties, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *V202101beta1CloudExport) SetAws(v V202101beta1AwsProperties)`

SetAws sets Aws field to given value.

### HasAws

`func (o *V202101beta1CloudExport) HasAws() bool`

HasAws returns a boolean if a field has been set.

### GetAzure

`func (o *V202101beta1CloudExport) GetAzure() V202101beta1AzureProperties`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *V202101beta1CloudExport) GetAzureOk() (*V202101beta1AzureProperties, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *V202101beta1CloudExport) SetAzure(v V202101beta1AzureProperties)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *V202101beta1CloudExport) HasAzure() bool`

HasAzure returns a boolean if a field has been set.

### GetGce

`func (o *V202101beta1CloudExport) GetGce() V202101beta1GceProperties`

GetGce returns the Gce field if non-nil, zero value otherwise.

### GetGceOk

`func (o *V202101beta1CloudExport) GetGceOk() (*V202101beta1GceProperties, bool)`

GetGceOk returns a tuple with the Gce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGce

`func (o *V202101beta1CloudExport) SetGce(v V202101beta1GceProperties)`

SetGce sets Gce field to given value.

### HasGce

`func (o *V202101beta1CloudExport) HasGce() bool`

HasGce returns a boolean if a field has been set.

### GetIbm

`func (o *V202101beta1CloudExport) GetIbm() V202101beta1IbmProperties`

GetIbm returns the Ibm field if non-nil, zero value otherwise.

### GetIbmOk

`func (o *V202101beta1CloudExport) GetIbmOk() (*V202101beta1IbmProperties, bool)`

GetIbmOk returns a tuple with the Ibm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbm

`func (o *V202101beta1CloudExport) SetIbm(v V202101beta1IbmProperties)`

SetIbm sets Ibm field to given value.

### HasIbm

`func (o *V202101beta1CloudExport) HasIbm() bool`

HasIbm returns a boolean if a field has been set.

### GetBgp

`func (o *V202101beta1CloudExport) GetBgp() V202101beta1BgpProperties`

GetBgp returns the Bgp field if non-nil, zero value otherwise.

### GetBgpOk

`func (o *V202101beta1CloudExport) GetBgpOk() (*V202101beta1BgpProperties, bool)`

GetBgpOk returns a tuple with the Bgp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgp

`func (o *V202101beta1CloudExport) SetBgp(v V202101beta1BgpProperties)`

SetBgp sets Bgp field to given value.

### HasBgp

`func (o *V202101beta1CloudExport) HasBgp() bool`

HasBgp returns a boolean if a field has been set.

### GetCurrentStatus

`func (o *V202101beta1CloudExport) GetCurrentStatus() CloudExportv202101beta1Status`

GetCurrentStatus returns the CurrentStatus field if non-nil, zero value otherwise.

### GetCurrentStatusOk

`func (o *V202101beta1CloudExport) GetCurrentStatusOk() (*CloudExportv202101beta1Status, bool)`

GetCurrentStatusOk returns a tuple with the CurrentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStatus

`func (o *V202101beta1CloudExport) SetCurrentStatus(v CloudExportv202101beta1Status)`

SetCurrentStatus sets CurrentStatus field to given value.

### HasCurrentStatus

`func (o *V202101beta1CloudExport) HasCurrentStatus() bool`

HasCurrentStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


