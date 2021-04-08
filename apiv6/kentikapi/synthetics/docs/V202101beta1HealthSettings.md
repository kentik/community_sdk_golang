# V202101beta1HealthSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LatencyCritical** | Pointer to **float32** |  | [optional] 
**LatencyWarning** | Pointer to **float32** |  | [optional] 
**PacketLossCritical** | Pointer to **float32** |  | [optional] 
**PacketLossWarning** | Pointer to **float32** |  | [optional] 
**JitterCritical** | Pointer to **float32** |  | [optional] 
**JitterWarning** | Pointer to **float32** |  | [optional] 
**HttpLatencyCritical** | Pointer to **float32** |  | [optional] 
**HttpLatencyWarning** | Pointer to **float32** |  | [optional] 
**HttpValidCodes** | Pointer to **[]int64** |  | [optional] 
**DnsValidCodes** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewV202101beta1HealthSettings

`func NewV202101beta1HealthSettings() *V202101beta1HealthSettings`

NewV202101beta1HealthSettings instantiates a new V202101beta1HealthSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV202101beta1HealthSettingsWithDefaults

`func NewV202101beta1HealthSettingsWithDefaults() *V202101beta1HealthSettings`

NewV202101beta1HealthSettingsWithDefaults instantiates a new V202101beta1HealthSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatencyCritical

`func (o *V202101beta1HealthSettings) GetLatencyCritical() float32`

GetLatencyCritical returns the LatencyCritical field if non-nil, zero value otherwise.

### GetLatencyCriticalOk

`func (o *V202101beta1HealthSettings) GetLatencyCriticalOk() (*float32, bool)`

GetLatencyCriticalOk returns a tuple with the LatencyCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyCritical

`func (o *V202101beta1HealthSettings) SetLatencyCritical(v float32)`

SetLatencyCritical sets LatencyCritical field to given value.

### HasLatencyCritical

`func (o *V202101beta1HealthSettings) HasLatencyCritical() bool`

HasLatencyCritical returns a boolean if a field has been set.

### GetLatencyWarning

`func (o *V202101beta1HealthSettings) GetLatencyWarning() float32`

GetLatencyWarning returns the LatencyWarning field if non-nil, zero value otherwise.

### GetLatencyWarningOk

`func (o *V202101beta1HealthSettings) GetLatencyWarningOk() (*float32, bool)`

GetLatencyWarningOk returns a tuple with the LatencyWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyWarning

`func (o *V202101beta1HealthSettings) SetLatencyWarning(v float32)`

SetLatencyWarning sets LatencyWarning field to given value.

### HasLatencyWarning

`func (o *V202101beta1HealthSettings) HasLatencyWarning() bool`

HasLatencyWarning returns a boolean if a field has been set.

### GetPacketLossCritical

`func (o *V202101beta1HealthSettings) GetPacketLossCritical() float32`

GetPacketLossCritical returns the PacketLossCritical field if non-nil, zero value otherwise.

### GetPacketLossCriticalOk

`func (o *V202101beta1HealthSettings) GetPacketLossCriticalOk() (*float32, bool)`

GetPacketLossCriticalOk returns a tuple with the PacketLossCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketLossCritical

`func (o *V202101beta1HealthSettings) SetPacketLossCritical(v float32)`

SetPacketLossCritical sets PacketLossCritical field to given value.

### HasPacketLossCritical

`func (o *V202101beta1HealthSettings) HasPacketLossCritical() bool`

HasPacketLossCritical returns a boolean if a field has been set.

### GetPacketLossWarning

`func (o *V202101beta1HealthSettings) GetPacketLossWarning() float32`

GetPacketLossWarning returns the PacketLossWarning field if non-nil, zero value otherwise.

### GetPacketLossWarningOk

`func (o *V202101beta1HealthSettings) GetPacketLossWarningOk() (*float32, bool)`

GetPacketLossWarningOk returns a tuple with the PacketLossWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketLossWarning

`func (o *V202101beta1HealthSettings) SetPacketLossWarning(v float32)`

SetPacketLossWarning sets PacketLossWarning field to given value.

### HasPacketLossWarning

`func (o *V202101beta1HealthSettings) HasPacketLossWarning() bool`

HasPacketLossWarning returns a boolean if a field has been set.

### GetJitterCritical

`func (o *V202101beta1HealthSettings) GetJitterCritical() float32`

GetJitterCritical returns the JitterCritical field if non-nil, zero value otherwise.

### GetJitterCriticalOk

`func (o *V202101beta1HealthSettings) GetJitterCriticalOk() (*float32, bool)`

GetJitterCriticalOk returns a tuple with the JitterCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitterCritical

`func (o *V202101beta1HealthSettings) SetJitterCritical(v float32)`

SetJitterCritical sets JitterCritical field to given value.

### HasJitterCritical

`func (o *V202101beta1HealthSettings) HasJitterCritical() bool`

HasJitterCritical returns a boolean if a field has been set.

### GetJitterWarning

`func (o *V202101beta1HealthSettings) GetJitterWarning() float32`

GetJitterWarning returns the JitterWarning field if non-nil, zero value otherwise.

### GetJitterWarningOk

`func (o *V202101beta1HealthSettings) GetJitterWarningOk() (*float32, bool)`

GetJitterWarningOk returns a tuple with the JitterWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitterWarning

`func (o *V202101beta1HealthSettings) SetJitterWarning(v float32)`

SetJitterWarning sets JitterWarning field to given value.

### HasJitterWarning

`func (o *V202101beta1HealthSettings) HasJitterWarning() bool`

HasJitterWarning returns a boolean if a field has been set.

### GetHttpLatencyCritical

`func (o *V202101beta1HealthSettings) GetHttpLatencyCritical() float32`

GetHttpLatencyCritical returns the HttpLatencyCritical field if non-nil, zero value otherwise.

### GetHttpLatencyCriticalOk

`func (o *V202101beta1HealthSettings) GetHttpLatencyCriticalOk() (*float32, bool)`

GetHttpLatencyCriticalOk returns a tuple with the HttpLatencyCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpLatencyCritical

`func (o *V202101beta1HealthSettings) SetHttpLatencyCritical(v float32)`

SetHttpLatencyCritical sets HttpLatencyCritical field to given value.

### HasHttpLatencyCritical

`func (o *V202101beta1HealthSettings) HasHttpLatencyCritical() bool`

HasHttpLatencyCritical returns a boolean if a field has been set.

### GetHttpLatencyWarning

`func (o *V202101beta1HealthSettings) GetHttpLatencyWarning() float32`

GetHttpLatencyWarning returns the HttpLatencyWarning field if non-nil, zero value otherwise.

### GetHttpLatencyWarningOk

`func (o *V202101beta1HealthSettings) GetHttpLatencyWarningOk() (*float32, bool)`

GetHttpLatencyWarningOk returns a tuple with the HttpLatencyWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpLatencyWarning

`func (o *V202101beta1HealthSettings) SetHttpLatencyWarning(v float32)`

SetHttpLatencyWarning sets HttpLatencyWarning field to given value.

### HasHttpLatencyWarning

`func (o *V202101beta1HealthSettings) HasHttpLatencyWarning() bool`

HasHttpLatencyWarning returns a boolean if a field has been set.

### GetHttpValidCodes

`func (o *V202101beta1HealthSettings) GetHttpValidCodes() []int64`

GetHttpValidCodes returns the HttpValidCodes field if non-nil, zero value otherwise.

### GetHttpValidCodesOk

`func (o *V202101beta1HealthSettings) GetHttpValidCodesOk() (*[]int64, bool)`

GetHttpValidCodesOk returns a tuple with the HttpValidCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpValidCodes

`func (o *V202101beta1HealthSettings) SetHttpValidCodes(v []int64)`

SetHttpValidCodes sets HttpValidCodes field to given value.

### HasHttpValidCodes

`func (o *V202101beta1HealthSettings) HasHttpValidCodes() bool`

HasHttpValidCodes returns a boolean if a field has been set.

### GetDnsValidCodes

`func (o *V202101beta1HealthSettings) GetDnsValidCodes() []int64`

GetDnsValidCodes returns the DnsValidCodes field if non-nil, zero value otherwise.

### GetDnsValidCodesOk

`func (o *V202101beta1HealthSettings) GetDnsValidCodesOk() (*[]int64, bool)`

GetDnsValidCodesOk returns a tuple with the DnsValidCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsValidCodes

`func (o *V202101beta1HealthSettings) SetDnsValidCodes(v []int64)`

SetDnsValidCodes sets DnsValidCodes field to given value.

### HasDnsValidCodes

`func (o *V202101beta1HealthSettings) HasDnsValidCodes() bool`

HasDnsValidCodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


