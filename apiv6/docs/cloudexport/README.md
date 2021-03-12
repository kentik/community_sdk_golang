# Documentation for Cloud Export Admin API

<a name="documentation-for-api-endpoints"></a>
## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CloudExportAdminServiceApi* | [**cloudExportAdminServiceCreateCloudExport**](Apis/CloudExportAdminServiceApi.md#cloudexportadminservicecreatecloudexport) | **POST** /cloud_export/v202101beta1/exports | 
*CloudExportAdminServiceApi* | [**cloudExportAdminServiceDeleteCloudExport**](Apis/CloudExportAdminServiceApi.md#cloudexportadminservicedeletecloudexport) | **DELETE** /cloud_export/v202101beta1/exports/{export.id} | 
*CloudExportAdminServiceApi* | [**cloudExportAdminServiceGetCloudExport**](Apis/CloudExportAdminServiceApi.md#cloudexportadminservicegetcloudexport) | **GET** /cloud_export/v202101beta1/exports/{export.id} | 
*CloudExportAdminServiceApi* | [**cloudExportAdminServiceListCloudExport**](Apis/CloudExportAdminServiceApi.md#cloudexportadminservicelistcloudexport) | **GET** /cloud_export/v202101beta1/exports | 
*CloudExportAdminServiceApi* | [**cloudExportAdminServicePatchCloudExport**](Apis/CloudExportAdminServiceApi.md#cloudexportadminservicepatchcloudexport) | **PATCH** /cloud_export/v202101beta1/exports/{export.id} | 
*CloudExportAdminServiceApi* | [**cloudExportAdminServiceUpdateCloudExport**](Apis/CloudExportAdminServiceApi.md#cloudexportadminserviceupdatecloudexport) | **PUT** /cloud_export/v202101beta1/exports/{export.id} | 


<a name="documentation-for-models"></a>
## Documentation for Models

 - [CloudExportv202101beta1Status](./Models/CloudExportv202101beta1Status.md)
 - [GooglerpcStatus](./Models/GooglerpcStatus.md)
 - [ProtobufAny](./Models/ProtobufAny.md)
 - [V202101beta1AwsProperties](./Models/V202101beta1AwsProperties.md)
 - [V202101beta1AzureProperties](./Models/V202101beta1AzureProperties.md)
 - [V202101beta1BgpProperties](./Models/V202101beta1BgpProperties.md)
 - [V202101beta1CloudExport](./Models/V202101beta1CloudExport.md)
 - [V202101beta1CloudExportType](./Models/V202101beta1CloudExportType.md)
 - [V202101beta1CreateCloudExportRequest](./Models/V202101beta1CreateCloudExportRequest.md)
 - [V202101beta1CreateCloudExportResponse](./Models/V202101beta1CreateCloudExportResponse.md)
 - [V202101beta1GceProperties](./Models/V202101beta1GceProperties.md)
 - [V202101beta1GetCloudExportResponse](./Models/V202101beta1GetCloudExportResponse.md)
 - [V202101beta1IbmProperties](./Models/V202101beta1IbmProperties.md)
 - [V202101beta1ListCloudExportResponse](./Models/V202101beta1ListCloudExportResponse.md)
 - [V202101beta1PatchCloudExportRequest](./Models/V202101beta1PatchCloudExportRequest.md)
 - [V202101beta1PatchCloudExportResponse](./Models/V202101beta1PatchCloudExportResponse.md)
 - [V202101beta1UpdateCloudExportRequest](./Models/V202101beta1UpdateCloudExportRequest.md)
 - [V202101beta1UpdateCloudExportResponse](./Models/V202101beta1UpdateCloudExportResponse.md)


<a name="documentation-for-authorization"></a>
## Documentation for Authorization

<a name="email"></a>
### email

- **Type**: API key
- **API key parameter name**: X-CH-Auth-Email
- **Location**: HTTP header

<a name="token"></a>
### token

- **Type**: API key
- **API key parameter name**: X-CH-Auth-API-Token
- **Location**: HTTP header

