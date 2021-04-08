# V202101beta1HealthMoment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **time.Time** |  | [optional] 
**SrcIp** | Pointer to **string** |  | [optional] 
**DstIp** | Pointer to **string** |  | [optional] 
**PacketLoss** | Pointer to **int64** |  | [optional] 
**AvgLatency** | Pointer to **int64** |  | [optional] 
**AvgWeightedLatency** | Pointer to **int64** |  | [optional] 
**RollingAvgLatency** | Pointer to **int64** |  | [optional] 
**RollingStddevLatency** | Pointer to **int64** |  | [optional] 
**RollingAvgWeightedLatency** | Pointer to **int64** |  | [optional] 
**LatencyHealth** | Pointer to **string** |  | [optional] 
**PacketLossHealth** | Pointer to **string** |  | [optional] 
**OverallHealth** | Pointer to [**V202101beta1Health**](V202101beta1Health.md) |  | [optional] 
**AvgJitter** | Pointer to **int64** |  | [optional] 
**RollingAvgJitter** | Pointer to **int64** |  | [optional] 
**RollingStdJitter** | Pointer to **int64** |  | [optional] 
**JitterHealth** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**TaskType** | Pointer to **string** |  | [optional] 

## Methods

### NewV202101beta1HealthMoment

`func NewV202101beta1HealthMoment() *V202101beta1HealthMoment`

NewV202101beta1HealthMoment instantiates a new V202101beta1HealthMoment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV202101beta1HealthMomentWithDefaults

`func NewV202101beta1HealthMomentWithDefaults() *V202101beta1HealthMoment`

NewV202101beta1HealthMomentWithDefaults instantiates a new V202101beta1HealthMoment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *V202101beta1HealthMoment) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *V202101beta1HealthMoment) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *V202101beta1HealthMoment) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *V202101beta1HealthMoment) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetSrcIp

`func (o *V202101beta1HealthMoment) GetSrcIp() string`

GetSrcIp returns the SrcIp field if non-nil, zero value otherwise.

### GetSrcIpOk

`func (o *V202101beta1HealthMoment) GetSrcIpOk() (*string, bool)`

GetSrcIpOk returns a tuple with the SrcIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcIp

`func (o *V202101beta1HealthMoment) SetSrcIp(v string)`

SetSrcIp sets SrcIp field to given value.

### HasSrcIp

`func (o *V202101beta1HealthMoment) HasSrcIp() bool`

HasSrcIp returns a boolean if a field has been set.

### GetDstIp

`func (o *V202101beta1HealthMoment) GetDstIp() string`

GetDstIp returns the DstIp field if non-nil, zero value otherwise.

### GetDstIpOk

`func (o *V202101beta1HealthMoment) GetDstIpOk() (*string, bool)`

GetDstIpOk returns a tuple with the DstIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstIp

`func (o *V202101beta1HealthMoment) SetDstIp(v string)`

SetDstIp sets DstIp field to given value.

### HasDstIp

`func (o *V202101beta1HealthMoment) HasDstIp() bool`

HasDstIp returns a boolean if a field has been set.

### GetPacketLoss

`func (o *V202101beta1HealthMoment) GetPacketLoss() int64`

GetPacketLoss returns the PacketLoss field if non-nil, zero value otherwise.

### GetPacketLossOk

`func (o *V202101beta1HealthMoment) GetPacketLossOk() (*int64, bool)`

GetPacketLossOk returns a tuple with the PacketLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketLoss

`func (o *V202101beta1HealthMoment) SetPacketLoss(v int64)`

SetPacketLoss sets PacketLoss field to given value.

### HasPacketLoss

`func (o *V202101beta1HealthMoment) HasPacketLoss() bool`

HasPacketLoss returns a boolean if a field has been set.

### GetAvgLatency

`func (o *V202101beta1HealthMoment) GetAvgLatency() int64`

GetAvgLatency returns the AvgLatency field if non-nil, zero value otherwise.

### GetAvgLatencyOk

`func (o *V202101beta1HealthMoment) GetAvgLatencyOk() (*int64, bool)`

GetAvgLatencyOk returns a tuple with the AvgLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgLatency

`func (o *V202101beta1HealthMoment) SetAvgLatency(v int64)`

SetAvgLatency sets AvgLatency field to given value.

### HasAvgLatency

`func (o *V202101beta1HealthMoment) HasAvgLatency() bool`

HasAvgLatency returns a boolean if a field has been set.

### GetAvgWeightedLatency

`func (o *V202101beta1HealthMoment) GetAvgWeightedLatency() int64`

GetAvgWeightedLatency returns the AvgWeightedLatency field if non-nil, zero value otherwise.

### GetAvgWeightedLatencyOk

`func (o *V202101beta1HealthMoment) GetAvgWeightedLatencyOk() (*int64, bool)`

GetAvgWeightedLatencyOk returns a tuple with the AvgWeightedLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgWeightedLatency

`func (o *V202101beta1HealthMoment) SetAvgWeightedLatency(v int64)`

SetAvgWeightedLatency sets AvgWeightedLatency field to given value.

### HasAvgWeightedLatency

`func (o *V202101beta1HealthMoment) HasAvgWeightedLatency() bool`

HasAvgWeightedLatency returns a boolean if a field has been set.

### GetRollingAvgLatency

`func (o *V202101beta1HealthMoment) GetRollingAvgLatency() int64`

GetRollingAvgLatency returns the RollingAvgLatency field if non-nil, zero value otherwise.

### GetRollingAvgLatencyOk

`func (o *V202101beta1HealthMoment) GetRollingAvgLatencyOk() (*int64, bool)`

GetRollingAvgLatencyOk returns a tuple with the RollingAvgLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollingAvgLatency

`func (o *V202101beta1HealthMoment) SetRollingAvgLatency(v int64)`

SetRollingAvgLatency sets RollingAvgLatency field to given value.

### HasRollingAvgLatency

`func (o *V202101beta1HealthMoment) HasRollingAvgLatency() bool`

HasRollingAvgLatency returns a boolean if a field has been set.

### GetRollingStddevLatency

`func (o *V202101beta1HealthMoment) GetRollingStddevLatency() int64`

GetRollingStddevLatency returns the RollingStddevLatency field if non-nil, zero value otherwise.

### GetRollingStddevLatencyOk

`func (o *V202101beta1HealthMoment) GetRollingStddevLatencyOk() (*int64, bool)`

GetRollingStddevLatencyOk returns a tuple with the RollingStddevLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollingStddevLatency

`func (o *V202101beta1HealthMoment) SetRollingStddevLatency(v int64)`

SetRollingStddevLatency sets RollingStddevLatency field to given value.

### HasRollingStddevLatency

`func (o *V202101beta1HealthMoment) HasRollingStddevLatency() bool`

HasRollingStddevLatency returns a boolean if a field has been set.

### GetRollingAvgWeightedLatency

`func (o *V202101beta1HealthMoment) GetRollingAvgWeightedLatency() int64`

GetRollingAvgWeightedLatency returns the RollingAvgWeightedLatency field if non-nil, zero value otherwise.

### GetRollingAvgWeightedLatencyOk

`func (o *V202101beta1HealthMoment) GetRollingAvgWeightedLatencyOk() (*int64, bool)`

GetRollingAvgWeightedLatencyOk returns a tuple with the RollingAvgWeightedLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollingAvgWeightedLatency

`func (o *V202101beta1HealthMoment) SetRollingAvgWeightedLatency(v int64)`

SetRollingAvgWeightedLatency sets RollingAvgWeightedLatency field to given value.

### HasRollingAvgWeightedLatency

`func (o *V202101beta1HealthMoment) HasRollingAvgWeightedLatency() bool`

HasRollingAvgWeightedLatency returns a boolean if a field has been set.

### GetLatencyHealth

`func (o *V202101beta1HealthMoment) GetLatencyHealth() string`

GetLatencyHealth returns the LatencyHealth field if non-nil, zero value otherwise.

### GetLatencyHealthOk

`func (o *V202101beta1HealthMoment) GetLatencyHealthOk() (*string, bool)`

GetLatencyHealthOk returns a tuple with the LatencyHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyHealth

`func (o *V202101beta1HealthMoment) SetLatencyHealth(v string)`

SetLatencyHealth sets LatencyHealth field to given value.

### HasLatencyHealth

`func (o *V202101beta1HealthMoment) HasLatencyHealth() bool`

HasLatencyHealth returns a boolean if a field has been set.

### GetPacketLossHealth

`func (o *V202101beta1HealthMoment) GetPacketLossHealth() string`

GetPacketLossHealth returns the PacketLossHealth field if non-nil, zero value otherwise.

### GetPacketLossHealthOk

`func (o *V202101beta1HealthMoment) GetPacketLossHealthOk() (*string, bool)`

GetPacketLossHealthOk returns a tuple with the PacketLossHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketLossHealth

`func (o *V202101beta1HealthMoment) SetPacketLossHealth(v string)`

SetPacketLossHealth sets PacketLossHealth field to given value.

### HasPacketLossHealth

`func (o *V202101beta1HealthMoment) HasPacketLossHealth() bool`

HasPacketLossHealth returns a boolean if a field has been set.

### GetOverallHealth

`func (o *V202101beta1HealthMoment) GetOverallHealth() V202101beta1Health`

GetOverallHealth returns the OverallHealth field if non-nil, zero value otherwise.

### GetOverallHealthOk

`func (o *V202101beta1HealthMoment) GetOverallHealthOk() (*V202101beta1Health, bool)`

GetOverallHealthOk returns a tuple with the OverallHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallHealth

`func (o *V202101beta1HealthMoment) SetOverallHealth(v V202101beta1Health)`

SetOverallHealth sets OverallHealth field to given value.

### HasOverallHealth

`func (o *V202101beta1HealthMoment) HasOverallHealth() bool`

HasOverallHealth returns a boolean if a field has been set.

### GetAvgJitter

`func (o *V202101beta1HealthMoment) GetAvgJitter() int64`

GetAvgJitter returns the AvgJitter field if non-nil, zero value otherwise.

### GetAvgJitterOk

`func (o *V202101beta1HealthMoment) GetAvgJitterOk() (*int64, bool)`

GetAvgJitterOk returns a tuple with the AvgJitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgJitter

`func (o *V202101beta1HealthMoment) SetAvgJitter(v int64)`

SetAvgJitter sets AvgJitter field to given value.

### HasAvgJitter

`func (o *V202101beta1HealthMoment) HasAvgJitter() bool`

HasAvgJitter returns a boolean if a field has been set.

### GetRollingAvgJitter

`func (o *V202101beta1HealthMoment) GetRollingAvgJitter() int64`

GetRollingAvgJitter returns the RollingAvgJitter field if non-nil, zero value otherwise.

### GetRollingAvgJitterOk

`func (o *V202101beta1HealthMoment) GetRollingAvgJitterOk() (*int64, bool)`

GetRollingAvgJitterOk returns a tuple with the RollingAvgJitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollingAvgJitter

`func (o *V202101beta1HealthMoment) SetRollingAvgJitter(v int64)`

SetRollingAvgJitter sets RollingAvgJitter field to given value.

### HasRollingAvgJitter

`func (o *V202101beta1HealthMoment) HasRollingAvgJitter() bool`

HasRollingAvgJitter returns a boolean if a field has been set.

### GetRollingStdJitter

`func (o *V202101beta1HealthMoment) GetRollingStdJitter() int64`

GetRollingStdJitter returns the RollingStdJitter field if non-nil, zero value otherwise.

### GetRollingStdJitterOk

`func (o *V202101beta1HealthMoment) GetRollingStdJitterOk() (*int64, bool)`

GetRollingStdJitterOk returns a tuple with the RollingStdJitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollingStdJitter

`func (o *V202101beta1HealthMoment) SetRollingStdJitter(v int64)`

SetRollingStdJitter sets RollingStdJitter field to given value.

### HasRollingStdJitter

`func (o *V202101beta1HealthMoment) HasRollingStdJitter() bool`

HasRollingStdJitter returns a boolean if a field has been set.

### GetJitterHealth

`func (o *V202101beta1HealthMoment) GetJitterHealth() string`

GetJitterHealth returns the JitterHealth field if non-nil, zero value otherwise.

### GetJitterHealthOk

`func (o *V202101beta1HealthMoment) GetJitterHealthOk() (*string, bool)`

GetJitterHealthOk returns a tuple with the JitterHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitterHealth

`func (o *V202101beta1HealthMoment) SetJitterHealth(v string)`

SetJitterHealth sets JitterHealth field to given value.

### HasJitterHealth

`func (o *V202101beta1HealthMoment) HasJitterHealth() bool`

HasJitterHealth returns a boolean if a field has been set.

### GetData

`func (o *V202101beta1HealthMoment) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V202101beta1HealthMoment) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V202101beta1HealthMoment) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *V202101beta1HealthMoment) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSize

`func (o *V202101beta1HealthMoment) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *V202101beta1HealthMoment) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *V202101beta1HealthMoment) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *V202101beta1HealthMoment) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStatus

`func (o *V202101beta1HealthMoment) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V202101beta1HealthMoment) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V202101beta1HealthMoment) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V202101beta1HealthMoment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaskType

`func (o *V202101beta1HealthMoment) GetTaskType() string`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *V202101beta1HealthMoment) GetTaskTypeOk() (*string, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *V202101beta1HealthMoment) SetTaskType(v string)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *V202101beta1HealthMoment) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


