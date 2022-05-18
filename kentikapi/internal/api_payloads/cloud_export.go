package api_payloads

import (
	"fmt"
	"log"

	"github.com/AlekSi/pointer"
	cloudexportpb "github.com/kentik/api-schema-public/gen/go/kentik/cloud_export/v202101beta1"
	kentikErrors "github.com/kentik/community_sdk_golang/kentikapi/internal/errors"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type ListCloudExportsResponse cloudexportpb.ListCloudExportResponse

func (r *ListCloudExportsResponse) ToModel() (*models.GetAllCloudExportsResponse, error) {
	if r == nil {
		return nil, nil
	}

	ces, err := cloudExportsFromPayload(r.Exports)
	if err != nil {
		return nil, err
	}

	return &models.GetAllCloudExportsResponse{
		CloudExports:             ces,
		InvalidCloudExportsCount: r.InvalidExportsCount,
	}, nil
}

func cloudExportsFromPayload(exports []*cloudexportpb.CloudExport) ([]models.CloudExport, error) {
	var result []models.CloudExport
	for i, e := range exports {
		ce, err := CloudExportFromPayload(e)
		if err != nil {
			return nil, fmt.Errorf("cloud export with index %v: %w", i, err)
		}
		result = append(result, *ce)
	}
	return result, nil
}

// CloudExportFromPayload converts cloud export payload to model.
func CloudExportFromPayload(ce *cloudexportpb.CloudExport) (*models.CloudExport, error) {
	if ce == nil {
		return nil, kentikErrors.New(kentikErrors.InvalidResponse, "cloud export response payload is nil")
	}

	properties, err := propertiesFromPayload(ce)
	if err != nil {
		return nil, kentikErrors.New(kentikErrors.InvalidResponse, err.Error())
	}

	return &models.CloudExport{
		ID:            ce.Id,
		Type:          models.CloudExportType(ce.Type.String()),
		Enabled:       pointer.ToBool(ce.Enabled),
		Name:          ce.Name,
		Description:   ce.Description,
		PlanID:        ce.PlanId,
		CloudProvider: models.CloudProvider(ce.CloudProvider),
		Properties:    properties,
		BGP:           bgpPropertiesFromPayload(ce.GetBgp()),
		CurrentStatus: currentStatusFromPayload(ce.GetCurrentStatus(), ce.Id),
	}, nil
}

func propertiesFromPayload(ce *cloudexportpb.CloudExport) (models.CloudExportProperties, error) {
	switch ce.CloudProvider {
	case "aws":
		return awsPropertiesFromPayload(ce.GetAws())
	case "azure":
		return azurePropertiesFromPayload(ce.GetAzure())
	case "gce":
		return gcePropertiesFromPayload(ce.GetGce())
	case "ibm":
		return ibmPropertiesFromPayload(ce.GetIbm())
	default:
		return nil, fmt.Errorf("invalid cloud provider in response payload: %v", ce.CloudProvider)
	}
}

func awsPropertiesFromPayload(aws *cloudexportpb.AwsProperties) (*models.AWSProperties, error) {
	if aws == nil {
		return nil, fmt.Errorf("no AWS properties in response payload")
	}
	return &models.AWSProperties{
		Bucket:          aws.GetBucket(),
		IAMRoleARN:      aws.GetIamRoleArn(),
		Region:          aws.GetRegion(),
		DeleteAfterRead: pointer.ToBool(aws.GetDeleteAfterRead()),
		MultipleBuckets: pointer.ToBool(aws.GetMultipleBuckets()),
	}, nil
}

func azurePropertiesFromPayload(azure *cloudexportpb.AzureProperties) (*models.AzureProperties, error) {
	if azure == nil {
		return nil, fmt.Errorf("no Azure properties in response payload")
	}

	return &models.AzureProperties{
		Location:                 azure.GetLocation(),
		ResourceGroup:            azure.GetResourceGroup(),
		StorageAccount:           azure.GetStorageAccount(),
		SubscriptionID:           azure.GetSubscriptionId(),
		SecurityPrincipalEnabled: pointer.ToBool(azure.GetSecurityPrincipalEnabled()),
	}, nil
}

func gcePropertiesFromPayload(gce *cloudexportpb.GceProperties) (*models.GCEProperties, error) {
	if gce == nil {
		return nil, fmt.Errorf("no GCE properties in response payload")
	}

	return &models.GCEProperties{
		Project:      gce.GetProject(),
		Subscription: gce.GetSubscription(),
	}, nil
}

func ibmPropertiesFromPayload(ibm *cloudexportpb.IbmProperties) (*models.IBMProperties, error) {
	if ibm == nil {
		return nil, fmt.Errorf("no IBM properties in response payload")
	}

	return &models.IBMProperties{
		Bucket: ibm.GetBucket(),
	}, nil
}

func bgpPropertiesFromPayload(bgp *cloudexportpb.BgpProperties) *models.BGPProperties {
	if bgp == nil {
		return nil
	}

	return &models.BGPProperties{
		ApplyBGP:       pointer.ToBool(bgp.GetApplyBgp()),
		UseBGPDeviceID: bgp.GetUseBgpDeviceId(),
		DeviceBGPType:  bgp.GetDeviceBgpType(),
	}
}

func currentStatusFromPayload(cs *cloudexportpb.Status, id string) *models.CloudExportStatus {
	if cs == nil {
		log.Printf("Warning: currentStatusFromPayload: CloudExport.CurrentStatus is nil; resource ID: %v\n", id)
		return nil
	}

	return &models.CloudExportStatus{
		Status:               cs.GetStatus(),
		ErrorMessage:         cs.GetErrorMessage(),
		FlowFound:            boolProtoPtrToBoolPtr(cs.GetFlowFound()),
		APIAccess:            boolProtoPtrToBoolPtr(cs.GetApiAccess()),
		StorageAccountAccess: boolProtoPtrToBoolPtr(cs.GetStorageAccountAccess()),
	}
}

// CloudExportToPayload converts cloud export from model to payload. It sets only ID and read-write fields.
func CloudExportToPayload(ce *models.CloudExport) (*cloudexportpb.CloudExport, error) {
	if ce == nil {
		return nil, nil
	}

	payload := &cloudexportpb.CloudExport{
		Id:            ce.ID,
		Type:          cloudexportpb.CloudExportType(cloudexportpb.CloudExportType_value[string(ce.Type)]),
		Enabled:       pointer.GetBool(ce.Enabled),
		Name:          ce.Name,
		Description:   ce.Description,
		PlanId:        ce.PlanID,
		CloudProvider: string(ce.CloudProvider),
		Bgp:           bgpPropertiesToPayload(ce.BGP),
		CurrentStatus: nil, // read-only
	}

	return cePayloadWithProperties(payload, ce)
}

func bgpPropertiesToPayload(bgp *models.BGPProperties) *cloudexportpb.BgpProperties {
	if bgp == nil {
		return nil
	}

	return &cloudexportpb.BgpProperties{
		ApplyBgp:       pointer.GetBool(bgp.ApplyBGP),
		UseBgpDeviceId: bgp.UseBGPDeviceID,
		DeviceBgpType:  bgp.DeviceBGPType,
	}
}

func cePayloadWithProperties(payload *cloudexportpb.CloudExport, ce *models.CloudExport) (*cloudexportpb.CloudExport, error) {
	switch ce.CloudProvider {
	case "aws":
		payload.Properties = awsPropertiesToPayload(ce)
	case "azure":
		payload.Properties = azurePropertiesToPayload(ce)
	case "gce":
		payload.Properties = gcePropertiesToPayload(ce)
	case "ibm":
		payload.Properties = ibmPropertiesToPayload(ce)
	default:
		return nil, kentikErrors.New(kentikErrors.InvalidRequest, fmt.Sprintf("invalid cloud provider: %v", ce.CloudProvider))
	}
	return payload, nil
}

func awsPropertiesToPayload(ce *models.CloudExport) *cloudexportpb.CloudExport_Aws {
	return &cloudexportpb.CloudExport_Aws{
		Aws: &cloudexportpb.AwsProperties{
			Bucket:          ce.GetAWSProperties().Bucket,
			IamRoleArn:      ce.GetAWSProperties().IAMRoleARN,
			Region:          ce.GetAWSProperties().Region,
			DeleteAfterRead: pointer.GetBool(ce.GetAWSProperties().DeleteAfterRead),
			MultipleBuckets: pointer.GetBool(ce.GetAWSProperties().MultipleBuckets),
		},
	}
}

func azurePropertiesToPayload(ce *models.CloudExport) *cloudexportpb.CloudExport_Azure {
	return &cloudexportpb.CloudExport_Azure{
		Azure: &cloudexportpb.AzureProperties{
			Location:                 ce.GetAzureProperties().Location,
			ResourceGroup:            ce.GetAzureProperties().ResourceGroup,
			StorageAccount:           ce.GetAzureProperties().StorageAccount,
			SubscriptionId:           ce.GetAzureProperties().SubscriptionID,
			SecurityPrincipalEnabled: pointer.GetBool(ce.GetAzureProperties().SecurityPrincipalEnabled),
		},
	}
}

func gcePropertiesToPayload(ce *models.CloudExport) *cloudexportpb.CloudExport_Gce {
	return &cloudexportpb.CloudExport_Gce{
		Gce: &cloudexportpb.GceProperties{
			Project:      ce.GetGCEProperties().Project,
			Subscription: ce.GetGCEProperties().Subscription,
		},
	}
}

func ibmPropertiesToPayload(ce *models.CloudExport) *cloudexportpb.CloudExport_Ibm {
	return &cloudexportpb.CloudExport_Ibm{
		Ibm: &cloudexportpb.IbmProperties{
			Bucket: ce.GetIBMProperties().Bucket,
		},
	}
}

func boolProtoPtrToBoolPtr(v *wrapperspb.BoolValue) *bool {
	if v == nil {
		return nil
	}
	return pointer.ToBool(v.GetValue())
}
