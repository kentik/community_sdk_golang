package cloud

import "github.com/kentik/community_sdk_golang/kentikapi/models"

// NewAWSExport creates a new AWS Export with all required fields set.
func NewAWSExport(obj AWSExportRequiredFields) *Export {
	return &Export{
		Name:     obj.Name,
		PlanID:   obj.PlanID,
		Provider: ProviderAWS,
		Properties: &AWSProperties{
			Bucket: obj.AWSProperties.Bucket,
		},
	}
}

// NewAzureExport creates a new Azure Export with all required fields set.
func NewAzureExport(obj AzureExportRequiredFields) *Export {
	return &Export{
		Name:     obj.Name,
		PlanID:   obj.PlanID,
		Provider: ProviderAzure,
		Properties: &AzureProperties{
			Location:       obj.AzureProperties.Location,
			ResourceGroup:  obj.AzureProperties.ResourceGroup,
			StorageAccount: obj.AzureProperties.StorageAccount,
			SubscriptionID: obj.AzureProperties.SubscriptionID,
		},
	}
}

// NewGCEExport creates a new GCE Export with all required fields set.
func NewGCEExport(obj GCEExportRequiredFields) *Export {
	return &Export{
		Name:     obj.Name,
		PlanID:   obj.PlanID,
		Provider: ProviderGCE,
		Properties: &GCEProperties{
			Project:      obj.GCEProperties.Project,
			Subscription: obj.GCEProperties.Subscription,
		},
	}
}

// NewIBMExport creates a new IBM Export with all required fields set.
func NewIBMExport(obj IBMExportRequiredFields) *Export {
	return &Export{
		Name:     obj.Name,
		PlanID:   obj.PlanID,
		Provider: ProviderIBM,
		Properties: &IBMProperties{
			Bucket: obj.IBMProperties.Bucket,
		},
	}
}

// GetAllExportsResponse model.
type GetAllExportsResponse struct {
	// Exports holds all cloud export tasks.
	Exports []Export
	// InvalidExportsCount is a number of invalid cloud export tasks.
	InvalidExportsCount uint32
}

// Export defines a cloud export task.
type Export struct {
	// Read-write properties

	// Type of export task.
	Type ExportType
	// Enabled specifies whether this task is enabled and intended to run or disabled.
	Enabled *bool
	// Name is a short name for this export task.
	Name string
	// Description is an optional, longer description of this export task.
	Description string
	// PlanID is the identifier of the Kentik plan associated with this export.
	PlanID string
	// Provider is the cloud provider targeted by this export, e.g. AWS, Azure, GCE, IBM.
	Provider Provider
	// / Properties specific to the cloud provider (AWS, Azure, GCE, IBM).
	Properties ExportProperties
	// BGPProperties are optional BGP related settings.
	BGP *BGPProperties

	// Read-only properties

	// ID is unique cloud export identification. It is read-only.
	ID models.ID
	// CurrentStatus is the most current status Kentik has about this export. It is read-only.
	CurrentStatus *ExportStatus
}

func (ce *Export) GetAWSProperties() *AWSProperties {
	p, _ := ce.Properties.(*AWSProperties) //nolint:errcheck // user can check the pointer
	return p
}

func (ce *Export) GetAzureProperties() *AzureProperties {
	p, _ := ce.Properties.(*AzureProperties) //nolint:errcheck // user can check the pointer
	return p
}

func (ce *Export) GetGCEProperties() *GCEProperties {
	p, _ := ce.Properties.(*GCEProperties) //nolint:errcheck // user can check the pointer
	return p
}

func (ce *Export) GetIBMProperties() *IBMProperties {
	p, _ := ce.Properties.(*IBMProperties) //nolint:errcheck // user can check the pointer
	return p
}

// ExportProperties emulates a union of AWSProperties, AzureProperties, GCEProperties and IBMProperties.
type ExportProperties interface {
	isExportProperties()
}

// AWSProperties are specific to Amazon Web Services VPC flow logs exports.
type AWSProperties struct {
	// Bucket is source S3 bucket to fetch VPC flow logs from.
	Bucket string
	// IAMRoleARN is an ARN for the IAM role to assume when fetching data or making AWS calls for this export.
	IAMRoleARN string
	// Region is an AWS region where the source S3 bucket resides.
	Region string
	// DeleteAfterRead enables the deletion of VPC flow log chunks from S3 after they have been read.
	DeleteAfterRead *bool
	// MultipleBuckets enables using multiple source S3 buckets.
	MultipleBuckets *bool
}

func (p *AWSProperties) isExportProperties() {}

// AzureProperties are specific to Azure exports.
type AzureProperties struct {
	// Location is an Azure location.
	Location string
	// ResourceGroup is the name of resource group from which to collect flow logs.
	ResourceGroup string
	// StorageAccount is the name of storage account for storing flow logs.  where flow logs will be collected.
	StorageAccount string
	// SubscriptionID is the Azure subscription ID.
	SubscriptionID string
	// SecurityPrincipalEnabled enables security principal.
	SecurityPrincipalEnabled *bool
}

func (p *AzureProperties) isExportProperties() {}

// GCEProperties are specific to Google Cloud export.
type GCEProperties struct {
	// Project is a GCE project name.
	Project string
	// Subscription is a GCE subscription name.
	Subscription string
}

func (p *GCEProperties) isExportProperties() {}

// IBMProperties are specific to IBM Cloud exports.
type IBMProperties struct {
	// Bucket is an IBM bucket.
	Bucket string
}

func (p *IBMProperties) isExportProperties() {}

// BGPProperties are optional BGP related settings.
type BGPProperties struct {
	// ApplyBGP enables applying BGP data discovered via another device to the flow from this export.
	ApplyBGP *bool
	// UseBGPDeviceID specifies which other device to get BGP data from.
	UseBGPDeviceID string
	DeviceBGPType  string
}

// ExportStatus is export task status.
type ExportStatus struct {
	Status string
	// ErrorMessage holds current error information.
	ErrorMessage string
	// FlowFound informs whether flow logs were found.
	FlowFound            *bool
	APIAccess            *bool
	StorageAccountAccess *bool
}

// AWSExportRequiredFields is a subset of fields required to create an AWS Export.
type AWSExportRequiredFields struct {
	Name          string
	PlanID        string
	AWSProperties AWSPropertiesRequiredFields
}

// AWSPropertiesRequiredFields is a subset of AWSProperties required to create an AWS Export.
type AWSPropertiesRequiredFields struct {
	Bucket string
}

// AzureExportRequiredFields is a subset of fields required to create an Azure Export.
type AzureExportRequiredFields struct {
	Name            string
	PlanID          string
	AzureProperties AzurePropertiesRequiredFields
}

// AzurePropertiesRequiredFields is a subset of AzureProperties required to create an Azure Export.
type AzurePropertiesRequiredFields struct {
	Location       string
	ResourceGroup  string
	StorageAccount string
	SubscriptionID string
}

// GCEExportRequiredFields is a subset of fields required to create a GCE Export.
type GCEExportRequiredFields struct {
	Name          string
	PlanID        string
	GCEProperties GCEPropertiesRequiredFields
}

// GCEPropertiesRequiredFields is a subset of GCEProperties required to create a GCE Export.
type GCEPropertiesRequiredFields struct {
	Project      string
	Subscription string
}

// IBMExportRequiredFields is a subset of fields required to create an IBM Export.
type IBMExportRequiredFields struct {
	Name          string
	PlanID        string
	IBMProperties IBMPropertiesRequiredFields
}

// IBMPropertiesRequiredFields is a subset of IBMProperties required to create an IBM Export.
type IBMPropertiesRequiredFields struct {
	Bucket string
}

// ExportType is the type of export task.
type ExportType string

const (
	// ExportTypeUnspecified is invalid or incomplete cloud export.
	ExportTypeUnspecified ExportType = "CLOUD_EXPORT_TYPE_UNSPECIFIED"
	// ExportTypeKentikManaged is for cloud exports that are managed by Kentik.
	ExportTypeKentikManaged ExportType = "CLOUD_EXPORT_TYPE_KENTIK_MANAGED"
	// ExportTypeCustomerManaged is for cloud exports that are managed by Kentik customers,
	// e.g. by running an agent.
	ExportTypeCustomerManaged ExportType = "CLOUD_EXPORT_TYPE_CUSTOMER_MANAGED"
)

// Provider is the name of cloud provider.
type Provider string

const (
	ProviderAWS   Provider = "aws"
	ProviderAzure Provider = "azure"
	ProviderGCE   Provider = "gce" // gcp value in Agents API
	ProviderIBM   Provider = "ibm"
)
