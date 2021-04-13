# V202101beta1MeshMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **time.Time** |  | [optional] 
**Latency** | Pointer to [**V202101beta1MeshMetric**](V202101beta1MeshMetric.md) |  | [optional] 
**PacketLoss** | Pointer to [**V202101beta1MeshMetric**](V202101beta1MeshMetric.md) |  | [optional] 
**Jitter** | Pointer to [**V202101beta1MeshMetric**](V202101beta1MeshMetric.md) |  | [optional] 

## Methods

### NewV202101beta1MeshMetrics

`func NewV202101beta1MeshMetrics() *V202101beta1MeshMetrics`

NewV202101beta1MeshMetrics instantiates a new V202101beta1MeshMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV202101beta1MeshMetricsWithDefaults

`func NewV202101beta1MeshMetricsWithDefaults() *V202101beta1MeshMetrics`

NewV202101beta1MeshMetricsWithDefaults instantiates a new V202101beta1MeshMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *V202101beta1MeshMetrics) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *V202101beta1MeshMetrics) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *V202101beta1MeshMetrics) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *V202101beta1MeshMetrics) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetLatency

`func (o *V202101beta1MeshMetrics) GetLatency() V202101beta1MeshMetric`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *V202101beta1MeshMetrics) GetLatencyOk() (*V202101beta1MeshMetric, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *V202101beta1MeshMetrics) SetLatency(v V202101beta1MeshMetric)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *V202101beta1MeshMetrics) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetPacketLoss

`func (o *V202101beta1MeshMetrics) GetPacketLoss() V202101beta1MeshMetric`

GetPacketLoss returns the PacketLoss field if non-nil, zero value otherwise.

### GetPacketLossOk

`func (o *V202101beta1MeshMetrics) GetPacketLossOk() (*V202101beta1MeshMetric, bool)`

GetPacketLossOk returns a tuple with the PacketLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketLoss

`func (o *V202101beta1MeshMetrics) SetPacketLoss(v V202101beta1MeshMetric)`

SetPacketLoss sets PacketLoss field to given value.

### HasPacketLoss

`func (o *V202101beta1MeshMetrics) HasPacketLoss() bool`

HasPacketLoss returns a boolean if a field has been set.

### GetJitter

`func (o *V202101beta1MeshMetrics) GetJitter() V202101beta1MeshMetric`

GetJitter returns the Jitter field if non-nil, zero value otherwise.

### GetJitterOk

`func (o *V202101beta1MeshMetrics) GetJitterOk() (*V202101beta1MeshMetric, bool)`

GetJitterOk returns a tuple with the Jitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitter

`func (o *V202101beta1MeshMetrics) SetJitter(v V202101beta1MeshMetric)`

SetJitter sets Jitter field to given value.

### HasJitter

`func (o *V202101beta1MeshMetrics) HasJitter() bool`

HasJitter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


