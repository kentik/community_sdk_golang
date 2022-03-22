package api_payloads

import (
	"fmt"
	"log"

	"github.com/AlekSi/pointer"
	cloudexportpb "github.com/kentik/api-schema-public/gen/go/kentik/cloud_export/v202101beta1"
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
		return nil, fmt.Errorf("cloud export response payload is nil")
	}

	return &models.CloudExport{
		ID:              ce.Id,
		Type:            models.CloudExportType(ce.Type.String()),
		Enabled:         pointer.ToBool(ce.Enabled),
		Name:            ce.Name,
		Description:     ce.Description,
		PlanID:          ce.PlanId,
		CloudProvider:   models.CloudProvider(ce.CloudProvider),
		AWSProperties:   awsPropertiesFromPayload(ce.GetAws()),
		AzureProperties: azurePropertiesFromPayload(ce.GetAzure()),
		GCEProperties:   gcePropertiesFromPayload(ce.GetGce()),
		IBMProperties:   ibmPropertiesFromPayload(ce.GetIbm()),
		BGP:             bgpPropertiesFromPayload(ce.GetBgp()),
		CurrentStatus:   currentStatusFromPayload(ce.GetCurrentStatus(), ce.Id),
	}, nil
}

func awsPropertiesFromPayload(aws *cloudexportpb.AwsProperties) *models.AWSProperties {
	if aws == nil {
		return nil
	}
	return &models.AWSProperties{
		Bucket:          aws.GetBucket(),
		IAMRoleARN:      aws.GetIamRoleArn(),
		Region:          aws.GetRegion(),
		DeleteAfterRead: pointer.ToBool(aws.GetDeleteAfterRead()),
		MultipleBuckets: pointer.ToBool(aws.GetMultipleBuckets()),
	}
}

func azurePropertiesFromPayload(azure *cloudexportpb.AzureProperties) *models.AzureProperties {
	if azure == nil {
		return nil
	}

	return &models.AzureProperties{
		Location:                 azure.GetLocation(),
		ResourceGroup:            azure.GetResourceGroup(),
		StorageAccount:           azure.GetStorageAccount(),
		SubscriptionID:           azure.GetSubscriptionId(),
		SecurityPrincipalEnabled: pointer.ToBool(azure.GetSecurityPrincipalEnabled()),
	}
}

func gcePropertiesFromPayload(gce *cloudexportpb.GceProperties) *models.GCEProperties {
	if gce == nil {
		return nil
	}

	return &models.GCEProperties{
		Project:      gce.GetProject(),
		Subscription: gce.GetSubscription(),
	}
}

func ibmPropertiesFromPayload(ibm *cloudexportpb.IbmProperties) *models.IBMProperties {
	if ibm == nil {
		return nil
	}

	return &models.IBMProperties{
		Bucket: ibm.GetBucket(),
	}
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

// CloudExportToPayload converts cloud export from model to payload.
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
		return nil, fmt.Errorf("invalid cloud provider: %v", ce.CloudProvider)
	}
	return payload, nil
}

func awsPropertiesToPayload(ce *models.CloudExport) *cloudexportpb.CloudExport_Aws {
	return &cloudexportpb.CloudExport_Aws{
		Aws: &cloudexportpb.AwsProperties{
			Bucket:          ce.AWSProperties.Bucket,
			IamRoleArn:      ce.AWSProperties.IAMRoleARN,
			Region:          ce.AWSProperties.Region,
			DeleteAfterRead: pointer.GetBool(ce.AWSProperties.DeleteAfterRead),
			MultipleBuckets: pointer.GetBool(ce.AWSProperties.MultipleBuckets),
		},
	}
}

func azurePropertiesToPayload(ce *models.CloudExport) *cloudexportpb.CloudExport_Azure {
	return &cloudexportpb.CloudExport_Azure{
		Azure: &cloudexportpb.AzureProperties{
			Location:                 ce.AzureProperties.Location,
			ResourceGroup:            ce.AzureProperties.ResourceGroup,
			StorageAccount:           ce.AzureProperties.StorageAccount,
			SubscriptionId:           ce.AzureProperties.SubscriptionID,
			SecurityPrincipalEnabled: pointer.GetBool(ce.AzureProperties.SecurityPrincipalEnabled),
		},
	}
}

func gcePropertiesToPayload(ce *models.CloudExport) *cloudexportpb.CloudExport_Gce {
	return &cloudexportpb.CloudExport_Gce{
		Gce: &cloudexportpb.GceProperties{
			Project:      ce.GCEProperties.Project,
			Subscription: ce.GCEProperties.Subscription,
		},
	}
}

func ibmPropertiesToPayload(ce *models.CloudExport) *cloudexportpb.CloudExport_Ibm {
	return &cloudexportpb.CloudExport_Ibm{
		Ibm: &cloudexportpb.IbmProperties{
			Bucket: ce.IBMProperties.Bucket,
		},
	}
}

func boolProtoPtrToBoolPtr(v *wrapperspb.BoolValue) *bool {
	if v == nil {
		return nil
	}
	return pointer.ToBool(v.GetValue())
}
