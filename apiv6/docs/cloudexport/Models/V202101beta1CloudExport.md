# V202101beta1CloudExport
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | [**String**](string.md) | The internal cloud export identifier. This is Read-only and assigned by Kentik. | [optional] [default to null]
**type** | [**v202101beta1CloudExportType**](v202101beta1CloudExportType.md) |  | [optional] [default to null]
**enabled** | [**Boolean**](boolean.md) | Whether this task is enabled and intended to run, or disabled. | [optional] [default to null]
**name** | [**String**](string.md) | A short name for this export. | [optional] [default to null]
**description** | [**String**](string.md) | An optional, longer description. | [optional] [default to null]
**apiRoot** | [**String**](string.md) |  | [optional] [default to null]
**flowDest** | [**String**](string.md) |  | [optional] [default to null]
**planId** | [**String**](string.md) | The identifier of the Kentik plan associated with this task. | [optional] [default to null]
**cloudProvider** | [**String**](string.md) |  | [optional] [default to null]
**aws** | [**v202101beta1AwsProperties**](v202101beta1AwsProperties.md) |  | [optional] [default to null]
**azure** | [**v202101beta1AzureProperties**](v202101beta1AzureProperties.md) |  | [optional] [default to null]
**gce** | [**v202101beta1GceProperties**](v202101beta1GceProperties.md) |  | [optional] [default to null]
**ibm** | [**v202101beta1IbmProperties**](v202101beta1IbmProperties.md) |  | [optional] [default to null]
**bgp** | [**v202101beta1BgpProperties**](v202101beta1BgpProperties.md) |  | [optional] [default to null]
**currentStatus** | [**cloud_exportv202101beta1Status**](cloud_exportv202101beta1Status.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

