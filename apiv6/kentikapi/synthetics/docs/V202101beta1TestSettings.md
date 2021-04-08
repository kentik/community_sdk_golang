# V202101beta1TestSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to [**V202101beta1HostnameTest**](V202101beta1HostnameTest.md) |  | [optional] 
**Ip** | Pointer to [**V202101beta1IpTest**](V202101beta1IpTest.md) |  | [optional] 
**Agent** | Pointer to [**V202101beta1AgentTest**](V202101beta1AgentTest.md) |  | [optional] 
**Flow** | Pointer to [**V202101beta1FlowTest**](V202101beta1FlowTest.md) |  | [optional] 
**Site** | Pointer to [**V202101beta1SiteTest**](V202101beta1SiteTest.md) |  | [optional] 
**Tag** | Pointer to [**V202101beta1TagTest**](V202101beta1TagTest.md) |  | [optional] 
**Dns** | Pointer to [**V202101beta1DnsTest**](V202101beta1DnsTest.md) |  | [optional] 
**Url** | Pointer to [**V202101beta1UrlTest**](V202101beta1UrlTest.md) |  | [optional] 
**AgentIds** | Pointer to **[]string** |  | [optional] 
**Period** | Pointer to **int64** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 
**Expiry** | Pointer to **int64** |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Tasks** | Pointer to **[]string** |  | [optional] 
**HealthSettings** | Pointer to [**V202101beta1HealthSettings**](V202101beta1HealthSettings.md) |  | [optional] 
**MonitoringSettings** | Pointer to [**V202101beta1TestMonitoringSettings**](V202101beta1TestMonitoringSettings.md) |  | [optional] 
**Ping** | Pointer to [**V202101beta1TestPingSettings**](V202101beta1TestPingSettings.md) |  | [optional] 
**Trace** | Pointer to [**V202101beta1TestTraceSettings**](V202101beta1TestTraceSettings.md) |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**Family** | Pointer to [**V202101beta1IPFamily**](V202101beta1IPFamily.md) |  | [optional] [default to V202101BETA1IPFAMILY_UNSPECIFIED]
**Servers** | Pointer to **[]string** |  | [optional] 
**TargetType** | Pointer to **string** |  | [optional] 
**TargetValue** | Pointer to **string** |  | [optional] 
**UseLocalIp** | Pointer to **bool** |  | [optional] 
**Reciprocal** | Pointer to **bool** |  | [optional] 
**RollupLevel** | Pointer to **int64** |  | [optional] 

## Methods

### NewV202101beta1TestSettings

`func NewV202101beta1TestSettings() *V202101beta1TestSettings`

NewV202101beta1TestSettings instantiates a new V202101beta1TestSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV202101beta1TestSettingsWithDefaults

`func NewV202101beta1TestSettingsWithDefaults() *V202101beta1TestSettings`

NewV202101beta1TestSettingsWithDefaults instantiates a new V202101beta1TestSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *V202101beta1TestSettings) GetHostname() V202101beta1HostnameTest`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *V202101beta1TestSettings) GetHostnameOk() (*V202101beta1HostnameTest, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *V202101beta1TestSettings) SetHostname(v V202101beta1HostnameTest)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *V202101beta1TestSettings) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIp

`func (o *V202101beta1TestSettings) GetIp() V202101beta1IpTest`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *V202101beta1TestSettings) GetIpOk() (*V202101beta1IpTest, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *V202101beta1TestSettings) SetIp(v V202101beta1IpTest)`

SetIp sets Ip field to given value.

### HasIp

`func (o *V202101beta1TestSettings) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetAgent

`func (o *V202101beta1TestSettings) GetAgent() V202101beta1AgentTest`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *V202101beta1TestSettings) GetAgentOk() (*V202101beta1AgentTest, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *V202101beta1TestSettings) SetAgent(v V202101beta1AgentTest)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *V202101beta1TestSettings) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetFlow

`func (o *V202101beta1TestSettings) GetFlow() V202101beta1FlowTest`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *V202101beta1TestSettings) GetFlowOk() (*V202101beta1FlowTest, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *V202101beta1TestSettings) SetFlow(v V202101beta1FlowTest)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *V202101beta1TestSettings) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### GetSite

`func (o *V202101beta1TestSettings) GetSite() V202101beta1SiteTest`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *V202101beta1TestSettings) GetSiteOk() (*V202101beta1SiteTest, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *V202101beta1TestSettings) SetSite(v V202101beta1SiteTest)`

SetSite sets Site field to given value.

### HasSite

`func (o *V202101beta1TestSettings) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetTag

`func (o *V202101beta1TestSettings) GetTag() V202101beta1TagTest`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *V202101beta1TestSettings) GetTagOk() (*V202101beta1TagTest, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *V202101beta1TestSettings) SetTag(v V202101beta1TagTest)`

SetTag sets Tag field to given value.

### HasTag

`func (o *V202101beta1TestSettings) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetDns

`func (o *V202101beta1TestSettings) GetDns() V202101beta1DnsTest`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *V202101beta1TestSettings) GetDnsOk() (*V202101beta1DnsTest, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *V202101beta1TestSettings) SetDns(v V202101beta1DnsTest)`

SetDns sets Dns field to given value.

### HasDns

`func (o *V202101beta1TestSettings) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetUrl

`func (o *V202101beta1TestSettings) GetUrl() V202101beta1UrlTest`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V202101beta1TestSettings) GetUrlOk() (*V202101beta1UrlTest, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V202101beta1TestSettings) SetUrl(v V202101beta1UrlTest)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *V202101beta1TestSettings) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetAgentIds

`func (o *V202101beta1TestSettings) GetAgentIds() []string`

GetAgentIds returns the AgentIds field if non-nil, zero value otherwise.

### GetAgentIdsOk

`func (o *V202101beta1TestSettings) GetAgentIdsOk() (*[]string, bool)`

GetAgentIdsOk returns a tuple with the AgentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentIds

`func (o *V202101beta1TestSettings) SetAgentIds(v []string)`

SetAgentIds sets AgentIds field to given value.

### HasAgentIds

`func (o *V202101beta1TestSettings) HasAgentIds() bool`

HasAgentIds returns a boolean if a field has been set.

### GetPeriod

`func (o *V202101beta1TestSettings) GetPeriod() int64`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *V202101beta1TestSettings) GetPeriodOk() (*int64, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *V202101beta1TestSettings) SetPeriod(v int64)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *V202101beta1TestSettings) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetCount

`func (o *V202101beta1TestSettings) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *V202101beta1TestSettings) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *V202101beta1TestSettings) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *V202101beta1TestSettings) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetExpiry

`func (o *V202101beta1TestSettings) GetExpiry() int64`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *V202101beta1TestSettings) GetExpiryOk() (*int64, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *V202101beta1TestSettings) SetExpiry(v int64)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *V202101beta1TestSettings) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetLimit

`func (o *V202101beta1TestSettings) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *V202101beta1TestSettings) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *V202101beta1TestSettings) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *V202101beta1TestSettings) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetTasks

`func (o *V202101beta1TestSettings) GetTasks() []string`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *V202101beta1TestSettings) GetTasksOk() (*[]string, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *V202101beta1TestSettings) SetTasks(v []string)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *V202101beta1TestSettings) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetHealthSettings

`func (o *V202101beta1TestSettings) GetHealthSettings() V202101beta1HealthSettings`

GetHealthSettings returns the HealthSettings field if non-nil, zero value otherwise.

### GetHealthSettingsOk

`func (o *V202101beta1TestSettings) GetHealthSettingsOk() (*V202101beta1HealthSettings, bool)`

GetHealthSettingsOk returns a tuple with the HealthSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthSettings

`func (o *V202101beta1TestSettings) SetHealthSettings(v V202101beta1HealthSettings)`

SetHealthSettings sets HealthSettings field to given value.

### HasHealthSettings

`func (o *V202101beta1TestSettings) HasHealthSettings() bool`

HasHealthSettings returns a boolean if a field has been set.

### GetMonitoringSettings

`func (o *V202101beta1TestSettings) GetMonitoringSettings() V202101beta1TestMonitoringSettings`

GetMonitoringSettings returns the MonitoringSettings field if non-nil, zero value otherwise.

### GetMonitoringSettingsOk

`func (o *V202101beta1TestSettings) GetMonitoringSettingsOk() (*V202101beta1TestMonitoringSettings, bool)`

GetMonitoringSettingsOk returns a tuple with the MonitoringSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringSettings

`func (o *V202101beta1TestSettings) SetMonitoringSettings(v V202101beta1TestMonitoringSettings)`

SetMonitoringSettings sets MonitoringSettings field to given value.

### HasMonitoringSettings

`func (o *V202101beta1TestSettings) HasMonitoringSettings() bool`

HasMonitoringSettings returns a boolean if a field has been set.

### GetPing

`func (o *V202101beta1TestSettings) GetPing() V202101beta1TestPingSettings`

GetPing returns the Ping field if non-nil, zero value otherwise.

### GetPingOk

`func (o *V202101beta1TestSettings) GetPingOk() (*V202101beta1TestPingSettings, bool)`

GetPingOk returns a tuple with the Ping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPing

`func (o *V202101beta1TestSettings) SetPing(v V202101beta1TestPingSettings)`

SetPing sets Ping field to given value.

### HasPing

`func (o *V202101beta1TestSettings) HasPing() bool`

HasPing returns a boolean if a field has been set.

### GetTrace

`func (o *V202101beta1TestSettings) GetTrace() V202101beta1TestTraceSettings`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *V202101beta1TestSettings) GetTraceOk() (*V202101beta1TestTraceSettings, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *V202101beta1TestSettings) SetTrace(v V202101beta1TestTraceSettings)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *V202101beta1TestSettings) HasTrace() bool`

HasTrace returns a boolean if a field has been set.

### GetPort

`func (o *V202101beta1TestSettings) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *V202101beta1TestSettings) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *V202101beta1TestSettings) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *V202101beta1TestSettings) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *V202101beta1TestSettings) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *V202101beta1TestSettings) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *V202101beta1TestSettings) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *V202101beta1TestSettings) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetFamily

`func (o *V202101beta1TestSettings) GetFamily() V202101beta1IPFamily`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *V202101beta1TestSettings) GetFamilyOk() (*V202101beta1IPFamily, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *V202101beta1TestSettings) SetFamily(v V202101beta1IPFamily)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *V202101beta1TestSettings) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetServers

`func (o *V202101beta1TestSettings) GetServers() []string`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *V202101beta1TestSettings) GetServersOk() (*[]string, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *V202101beta1TestSettings) SetServers(v []string)`

SetServers sets Servers field to given value.

### HasServers

`func (o *V202101beta1TestSettings) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetTargetType

`func (o *V202101beta1TestSettings) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *V202101beta1TestSettings) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *V202101beta1TestSettings) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *V202101beta1TestSettings) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetTargetValue

`func (o *V202101beta1TestSettings) GetTargetValue() string`

GetTargetValue returns the TargetValue field if non-nil, zero value otherwise.

### GetTargetValueOk

`func (o *V202101beta1TestSettings) GetTargetValueOk() (*string, bool)`

GetTargetValueOk returns a tuple with the TargetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetValue

`func (o *V202101beta1TestSettings) SetTargetValue(v string)`

SetTargetValue sets TargetValue field to given value.

### HasTargetValue

`func (o *V202101beta1TestSettings) HasTargetValue() bool`

HasTargetValue returns a boolean if a field has been set.

### GetUseLocalIp

`func (o *V202101beta1TestSettings) GetUseLocalIp() bool`

GetUseLocalIp returns the UseLocalIp field if non-nil, zero value otherwise.

### GetUseLocalIpOk

`func (o *V202101beta1TestSettings) GetUseLocalIpOk() (*bool, bool)`

GetUseLocalIpOk returns a tuple with the UseLocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLocalIp

`func (o *V202101beta1TestSettings) SetUseLocalIp(v bool)`

SetUseLocalIp sets UseLocalIp field to given value.

### HasUseLocalIp

`func (o *V202101beta1TestSettings) HasUseLocalIp() bool`

HasUseLocalIp returns a boolean if a field has been set.

### GetReciprocal

`func (o *V202101beta1TestSettings) GetReciprocal() bool`

GetReciprocal returns the Reciprocal field if non-nil, zero value otherwise.

### GetReciprocalOk

`func (o *V202101beta1TestSettings) GetReciprocalOk() (*bool, bool)`

GetReciprocalOk returns a tuple with the Reciprocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReciprocal

`func (o *V202101beta1TestSettings) SetReciprocal(v bool)`

SetReciprocal sets Reciprocal field to given value.

### HasReciprocal

`func (o *V202101beta1TestSettings) HasReciprocal() bool`

HasReciprocal returns a boolean if a field has been set.

### GetRollupLevel

`func (o *V202101beta1TestSettings) GetRollupLevel() int64`

GetRollupLevel returns the RollupLevel field if non-nil, zero value otherwise.

### GetRollupLevelOk

`func (o *V202101beta1TestSettings) GetRollupLevelOk() (*int64, bool)`

GetRollupLevelOk returns a tuple with the RollupLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollupLevel

`func (o *V202101beta1TestSettings) SetRollupLevel(v int64)`

SetRollupLevel sets RollupLevel field to given value.

### HasRollupLevel

`func (o *V202101beta1TestSettings) HasRollupLevel() bool`

HasRollupLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


