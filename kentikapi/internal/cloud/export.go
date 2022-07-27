package cloud

import (
	"errors"
	"fmt"
	"log"

	"github.com/AlekSi/pointer"
	cloudexportpb "github.com/kentik/api-schema-public/gen/go/kentik/cloud_export/v202101beta1"
	"github.com/kentik/community_sdk_golang/kentikapi/cloud"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	awsProvider   = "aws"
	azureProvider = "azure"
	gceProvider   = "gce"
	ibmProvider   = "ibm"
)

type listExportsResponse cloudexportpb.ListCloudExportResponse

func (r *listExportsResponse) ToModel() (*cloud.GetAllExportsResponse, error) {
	if r == nil {
		return &cloud.GetAllExportsResponse{}, nil
	}

	ces, err := exportsFromPayload(r.Exports)
	if err != nil {
		return nil, err
	}

	return &cloud.GetAllExportsResponse{
		Exports:             ces,
		InvalidExportsCount: r.InvalidExportsCount,
	}, nil
}

func exportsFromPayload(exports []*cloudexportpb.CloudExport) ([]cloud.Export, error) {
	var result []cloud.Export
	for i, e := range exports {
		ce, err := exportFromPayload(e)
		if err != nil {
			return nil, fmt.Errorf("cloud export with index %v: %w", i, err)
		}
		result = append(result, *ce)
	}
	return result, nil
}

// exportFromPayload converts cloud export payload to model.
func exportFromPayload(ce *cloudexportpb.CloudExport) (*cloud.Export, error) {
	if ce == nil {
		return nil, errors.New("response payload is nil")
	}

	if ce.Id == "" {
		return nil, errors.New("empty export ID in response payload")
	}

	properties, err := propertiesFromPayload(ce)
	if err != nil {
		return nil, err
	}

	return &cloud.Export{
		ID:            ce.Id,
		Type:          cloud.ExportType(ce.Type.String()),
		Enabled:       pointer.ToBool(ce.Enabled),
		Name:          ce.Name,
		Description:   ce.Description,
		PlanID:        ce.PlanId,
		Provider:      cloud.Provider(ce.CloudProvider),
		Properties:    properties,
		BGP:           bgpPropertiesFromPayload(ce.GetBgp()),
		CurrentStatus: currentStatusFromPayload(ce.GetCurrentStatus(), ce.Id),
	}, nil
}

func propertiesFromPayload(ce *cloudexportpb.CloudExport) (cloud.ExportProperties, error) {
	switch ce.CloudProvider {
	case awsProvider:
		return awsPropertiesFromPayload(ce.GetAws())
	case azureProvider:
		return azurePropertiesFromPayload(ce.GetAzure())
	case gceProvider:
		return gcePropertiesFromPayload(ce.GetGce())
	case ibmProvider:
		return ibmPropertiesFromPayload(ce.GetIbm())
	default:
		return nil, fmt.Errorf("unsupported cloud provider in response payload: %v", ce.CloudProvider)
	}
}

func awsPropertiesFromPayload(aws *cloudexportpb.AwsProperties) (*cloud.AWSProperties, error) {
	if aws == nil {
		return nil, fmt.Errorf("no AWS properties in response payload")
	}
	return &cloud.AWSProperties{
		Bucket:          aws.GetBucket(),
		IAMRoleARN:      aws.GetIamRoleArn(),
		Region:          aws.GetRegion(),
		DeleteAfterRead: pointer.ToBool(aws.GetDeleteAfterRead()),
		MultipleBuckets: pointer.ToBool(aws.GetMultipleBuckets()),
	}, nil
}

func azurePropertiesFromPayload(azure *cloudexportpb.AzureProperties) (*cloud.AzureProperties, error) {
	if azure == nil {
		return nil, fmt.Errorf("no Azure properties in response payload")
	}

	return &cloud.AzureProperties{
		Location:                 azure.GetLocation(),
		ResourceGroup:            azure.GetResourceGroup(),
		StorageAccount:           azure.GetStorageAccount(),
		SubscriptionID:           azure.GetSubscriptionId(),
		SecurityPrincipalEnabled: pointer.ToBool(azure.GetSecurityPrincipalEnabled()),
	}, nil
}

func gcePropertiesFromPayload(gce *cloudexportpb.GceProperties) (*cloud.GCEProperties, error) {
	if gce == nil {
		return nil, fmt.Errorf("no GCE properties in response payload")
	}

	return &cloud.GCEProperties{
		Project:      gce.GetProject(),
		Subscription: gce.GetSubscription(),
	}, nil
}

func ibmPropertiesFromPayload(ibm *cloudexportpb.IbmProperties) (*cloud.IBMProperties, error) {
	if ibm == nil {
		return nil, fmt.Errorf("no IBM properties in response payload")
	}

	return &cloud.IBMProperties{
		Bucket: ibm.GetBucket(),
	}, nil
}

func bgpPropertiesFromPayload(bgp *cloudexportpb.BgpProperties) *cloud.BGPProperties {
	if bgp == nil {
		return nil
	}

	return &cloud.BGPProperties{
		ApplyBGP:       pointer.ToBool(bgp.GetApplyBgp()),
		UseBGPDeviceID: bgp.GetUseBgpDeviceId(),
		DeviceBGPType:  bgp.GetDeviceBgpType(),
	}
}

func currentStatusFromPayload(cs *cloudexportpb.Status, id string) *cloud.ExportStatus {
	if cs == nil {
		log.Printf("Warning: currentStatusFromPayload: Export.CurrentStatus is nil; resource ID: %v\n", id)
		return nil
	}

	return &cloud.ExportStatus{
		Status:               cs.GetStatus(),
		ErrorMessage:         cs.GetErrorMessage(),
		FlowFound:            boolProtoPtrToBoolPtr(cs.GetFlowFound()),
		APIAccess:            boolProtoPtrToBoolPtr(cs.GetApiAccess()),
		StorageAccountAccess: boolProtoPtrToBoolPtr(cs.GetStorageAccountAccess()),
	}
}

// exportToPayload converts cloud export from model to payload. It sets only ID and read-write fields.
func exportToPayload(ce *cloud.Export) (*cloudexportpb.CloudExport, error) {
	if ce == nil {
		return nil, fmt.Errorf("cloud export object is nil")
	}

	payload := &cloudexportpb.CloudExport{
		Id:            ce.ID,
		Type:          cloudexportpb.CloudExportType(cloudexportpb.CloudExportType_value[string(ce.Type)]),
		Enabled:       pointer.GetBool(ce.Enabled),
		Name:          ce.Name,
		Description:   ce.Description,
		PlanId:        ce.PlanID,
		CloudProvider: string(ce.Provider),
		Bgp:           bgpPropertiesToPayload(ce.BGP),
		CurrentStatus: nil, // read-only
	}

	return cePayloadWithProperties(payload, ce)
}

func bgpPropertiesToPayload(bgp *cloud.BGPProperties) *cloudexportpb.BgpProperties {
	if bgp == nil {
		return nil
	}

	return &cloudexportpb.BgpProperties{
		ApplyBgp:       pointer.GetBool(bgp.ApplyBGP),
		UseBgpDeviceId: bgp.UseBGPDeviceID,
		DeviceBgpType:  bgp.DeviceBGPType,
	}
}

func cePayloadWithProperties(payload *cloudexportpb.CloudExport, ce *cloud.Export) (*cloudexportpb.CloudExport, error) {
	switch ce.Provider {
	case awsProvider:
		payload.Properties = awsPropertiesToPayload(ce)
	case azureProvider:
		payload.Properties = azurePropertiesToPayload(ce)
	case gceProvider:
		payload.Properties = gcePropertiesToPayload(ce)
	case ibmProvider:
		payload.Properties = ibmPropertiesToPayload(ce)
	default:
		return nil, fmt.Errorf("unsupported cloud provider: %v", ce.Provider)
	}
	return payload, nil
}

func awsPropertiesToPayload(ce *cloud.Export) *cloudexportpb.CloudExport_Aws {
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

func azurePropertiesToPayload(ce *cloud.Export) *cloudexportpb.CloudExport_Azure {
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

func gcePropertiesToPayload(ce *cloud.Export) *cloudexportpb.CloudExport_Gce {
	return &cloudexportpb.CloudExport_Gce{
		Gce: &cloudexportpb.GceProperties{
			Project:      ce.GetGCEProperties().Project,
			Subscription: ce.GetGCEProperties().Subscription,
		},
	}
}

func ibmPropertiesToPayload(ce *cloud.Export) *cloudexportpb.CloudExport_Ibm {
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

// validateCreateExportRequest checks if Export create request contains all required fields.
func validateCreateExportRequest(ce *cloud.Export) error {
	if ce == nil {
		return errors.New("cloud export object is nil")
	}
	if ce.Name == "" {
		return missingExportFieldError("Name")
	}
	if ce.PlanID == "" {
		return missingExportFieldError("PlanID")
	}
	return validateCEProvider(ce)
}

// validateExportUpdateRequest checks if Export update request contains all required fields.
func validateExportUpdateRequest(ce *cloud.Export) error {
	if ce == nil {
		return errors.New("cloud export object is nil")
	}
	if ce.ID == "" {
		return missingExportFieldError("ID")
	}
	if ce.Name == "" {
		return missingExportFieldError("Name")
	}
	if ce.PlanID == "" {
		return missingExportFieldError("PlanID")
	}
	return validateCEProvider(ce)
}

func validateCEProvider(ce *cloud.Export) error {
	switch ce.Provider {
	case "":
		return missingExportFieldError("Provider")
	case awsProvider:
		return validateAWSProvider(ce)
	case azureProvider:
		return validateAzureProvider(ce)
	case gceProvider:
		return validateGCEProvider(ce)
	case ibmProvider:
		return validateIBMProvider(ce)
	default:
		return fmt.Errorf("cloud provider '%s' is not supported", ce.Provider)
	}
}

func validateAWSProvider(ce *cloud.Export) error {
	if ce.GetAWSProperties() == nil {
		return missingExportFieldError("Properties")
	}
	if ce.GetAWSProperties().Bucket == "" {
		return missingExportFieldError("Properties.Bucket")
	}
	return nil
}

func validateAzureProvider(ce *cloud.Export) error {
	if ce.GetAzureProperties() == nil {
		return missingExportFieldError("Properties")
	}
	if ce.GetAzureProperties().Location == "" {
		return missingExportFieldError("Properties.Location")
	}
	if ce.GetAzureProperties().ResourceGroup == "" {
		return missingExportFieldError("Properties.ResourceGroup")
	}
	if ce.GetAzureProperties().StorageAccount == "" {
		return missingExportFieldError("Properties.StorageAccount")
	}
	if ce.GetAzureProperties().SubscriptionID == "" {
		return missingExportFieldError("Properties.SubscriptionID")
	}
	return nil
}

func validateGCEProvider(ce *cloud.Export) error {
	if ce.GetGCEProperties() == nil {
		return missingExportFieldError("Properties")
	}
	if ce.GetGCEProperties().Project == "" {
		return missingExportFieldError("Properties.Project")
	}
	if ce.GetGCEProperties().Subscription == "" {
		return missingExportFieldError("Properties.Subscription")
	}
	return nil
}

func validateIBMProvider(ce *cloud.Export) error {
	if ce.GetIBMProperties() == nil {
		return missingExportFieldError("Properties")
	}
	if ce.GetIBMProperties().Bucket == "" {
		return missingExportFieldError("Properties.Bucket")
	}
	return nil
}

func missingExportFieldError(field string) error {
	return fmt.Errorf("export's %q field is missing", field)
}
