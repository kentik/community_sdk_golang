# V202101beta1IPInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** |  | [optional] 
**Asn** | Pointer to [**V202101beta1ASN**](V202101beta1ASN.md) |  | [optional] 
**Geo** | Pointer to [**V202101beta1Geo**](V202101beta1Geo.md) |  | [optional] 
**Dns** | Pointer to [**V202101beta1DNS**](V202101beta1DNS.md) |  | [optional] 

## Methods

### NewV202101beta1IPInfo

`func NewV202101beta1IPInfo() *V202101beta1IPInfo`

NewV202101beta1IPInfo instantiates a new V202101beta1IPInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV202101beta1IPInfoWithDefaults

`func NewV202101beta1IPInfoWithDefaults() *V202101beta1IPInfo`

NewV202101beta1IPInfoWithDefaults instantiates a new V202101beta1IPInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *V202101beta1IPInfo) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *V202101beta1IPInfo) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *V202101beta1IPInfo) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *V202101beta1IPInfo) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetAsn

`func (o *V202101beta1IPInfo) GetAsn() V202101beta1ASN`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *V202101beta1IPInfo) GetAsnOk() (*V202101beta1ASN, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *V202101beta1IPInfo) SetAsn(v V202101beta1ASN)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *V202101beta1IPInfo) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetGeo

`func (o *V202101beta1IPInfo) GetGeo() V202101beta1Geo`

GetGeo returns the Geo field if non-nil, zero value otherwise.

### GetGeoOk

`func (o *V202101beta1IPInfo) GetGeoOk() (*V202101beta1Geo, bool)`

GetGeoOk returns a tuple with the Geo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeo

`func (o *V202101beta1IPInfo) SetGeo(v V202101beta1Geo)`

SetGeo sets Geo field to given value.

### HasGeo

`func (o *V202101beta1IPInfo) HasGeo() bool`

HasGeo returns a boolean if a field has been set.

### GetDns

`func (o *V202101beta1IPInfo) GetDns() V202101beta1DNS`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *V202101beta1IPInfo) GetDnsOk() (*V202101beta1DNS, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *V202101beta1IPInfo) SetDns(v V202101beta1DNS)`

SetDns sets Dns field to given value.

### HasDns

`func (o *V202101beta1IPInfo) HasDns() bool`

HasDns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


