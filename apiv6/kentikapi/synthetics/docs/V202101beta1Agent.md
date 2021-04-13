# V202101beta1Agent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**V202101beta1AgentStatus**](V202101beta1AgentStatus.md) |  | [optional] [default to V202101BETA1AGENTSTATUS_UNSPECIFIED]
**Alias** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Os** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Lat** | Pointer to **float64** |  | [optional] 
**Long** | Pointer to **float64** |  | [optional] 
**LastAuthed** | Pointer to **time.Time** |  | [optional] 
**Family** | Pointer to [**V202101beta1IPFamily**](V202101beta1IPFamily.md) |  | [optional] [default to V202101BETA1IPFAMILY_UNSPECIFIED]
**Asn** | Pointer to **int64** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Challenge** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**TestIds** | Pointer to **[]string** |  | [optional] 
**LocalIp** | Pointer to **string** |  | [optional] 

## Methods

### NewV202101beta1Agent

`func NewV202101beta1Agent() *V202101beta1Agent`

NewV202101beta1Agent instantiates a new V202101beta1Agent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV202101beta1AgentWithDefaults

`func NewV202101beta1AgentWithDefaults() *V202101beta1Agent`

NewV202101beta1AgentWithDefaults instantiates a new V202101beta1Agent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V202101beta1Agent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V202101beta1Agent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V202101beta1Agent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V202101beta1Agent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *V202101beta1Agent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V202101beta1Agent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V202101beta1Agent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V202101beta1Agent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *V202101beta1Agent) GetStatus() V202101beta1AgentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V202101beta1Agent) GetStatusOk() (*V202101beta1AgentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V202101beta1Agent) SetStatus(v V202101beta1AgentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V202101beta1Agent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAlias

`func (o *V202101beta1Agent) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *V202101beta1Agent) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *V202101beta1Agent) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *V202101beta1Agent) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetType

`func (o *V202101beta1Agent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V202101beta1Agent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V202101beta1Agent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V202101beta1Agent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOs

`func (o *V202101beta1Agent) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *V202101beta1Agent) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *V202101beta1Agent) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *V202101beta1Agent) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetIp

`func (o *V202101beta1Agent) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *V202101beta1Agent) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *V202101beta1Agent) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *V202101beta1Agent) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLat

`func (o *V202101beta1Agent) GetLat() float64`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *V202101beta1Agent) GetLatOk() (*float64, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *V202101beta1Agent) SetLat(v float64)`

SetLat sets Lat field to given value.

### HasLat

`func (o *V202101beta1Agent) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLong

`func (o *V202101beta1Agent) GetLong() float64`

GetLong returns the Long field if non-nil, zero value otherwise.

### GetLongOk

`func (o *V202101beta1Agent) GetLongOk() (*float64, bool)`

GetLongOk returns a tuple with the Long field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLong

`func (o *V202101beta1Agent) SetLong(v float64)`

SetLong sets Long field to given value.

### HasLong

`func (o *V202101beta1Agent) HasLong() bool`

HasLong returns a boolean if a field has been set.

### GetLastAuthed

`func (o *V202101beta1Agent) GetLastAuthed() time.Time`

GetLastAuthed returns the LastAuthed field if non-nil, zero value otherwise.

### GetLastAuthedOk

`func (o *V202101beta1Agent) GetLastAuthedOk() (*time.Time, bool)`

GetLastAuthedOk returns a tuple with the LastAuthed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAuthed

`func (o *V202101beta1Agent) SetLastAuthed(v time.Time)`

SetLastAuthed sets LastAuthed field to given value.

### HasLastAuthed

`func (o *V202101beta1Agent) HasLastAuthed() bool`

HasLastAuthed returns a boolean if a field has been set.

### GetFamily

`func (o *V202101beta1Agent) GetFamily() V202101beta1IPFamily`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *V202101beta1Agent) GetFamilyOk() (*V202101beta1IPFamily, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *V202101beta1Agent) SetFamily(v V202101beta1IPFamily)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *V202101beta1Agent) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetAsn

`func (o *V202101beta1Agent) GetAsn() int64`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *V202101beta1Agent) GetAsnOk() (*int64, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *V202101beta1Agent) SetAsn(v int64)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *V202101beta1Agent) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetSiteId

`func (o *V202101beta1Agent) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *V202101beta1Agent) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *V202101beta1Agent) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *V202101beta1Agent) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetVersion

`func (o *V202101beta1Agent) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *V202101beta1Agent) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *V202101beta1Agent) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *V202101beta1Agent) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetChallenge

`func (o *V202101beta1Agent) GetChallenge() string`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *V202101beta1Agent) GetChallengeOk() (*string, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *V202101beta1Agent) SetChallenge(v string)`

SetChallenge sets Challenge field to given value.

### HasChallenge

`func (o *V202101beta1Agent) HasChallenge() bool`

HasChallenge returns a boolean if a field has been set.

### GetCity

`func (o *V202101beta1Agent) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *V202101beta1Agent) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *V202101beta1Agent) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *V202101beta1Agent) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetRegion

`func (o *V202101beta1Agent) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *V202101beta1Agent) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *V202101beta1Agent) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *V202101beta1Agent) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCountry

`func (o *V202101beta1Agent) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *V202101beta1Agent) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *V202101beta1Agent) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *V202101beta1Agent) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetTestIds

`func (o *V202101beta1Agent) GetTestIds() []string`

GetTestIds returns the TestIds field if non-nil, zero value otherwise.

### GetTestIdsOk

`func (o *V202101beta1Agent) GetTestIdsOk() (*[]string, bool)`

GetTestIdsOk returns a tuple with the TestIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestIds

`func (o *V202101beta1Agent) SetTestIds(v []string)`

SetTestIds sets TestIds field to given value.

### HasTestIds

`func (o *V202101beta1Agent) HasTestIds() bool`

HasTestIds returns a boolean if a field has been set.

### GetLocalIp

`func (o *V202101beta1Agent) GetLocalIp() string`

GetLocalIp returns the LocalIp field if non-nil, zero value otherwise.

### GetLocalIpOk

`func (o *V202101beta1Agent) GetLocalIpOk() (*string, bool)`

GetLocalIpOk returns a tuple with the LocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIp

`func (o *V202101beta1Agent) SetLocalIp(v string)`

SetLocalIp sets LocalIp field to given value.

### HasLocalIp

`func (o *V202101beta1Agent) HasLocalIp() bool`

HasLocalIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


