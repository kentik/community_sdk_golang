package models

// NewAWSCloudExport creates a new AWS CloudExport with all required fields set.
func NewAWSCloudExport(obj CloudExportAWSRequiredFields) *CloudExport {
	return &CloudExport{
		Name:          obj.Name,
		PlanID:        obj.PlanID,
		CloudProvider: CloudProviderAWS,
		AWSProperties: &AWSProperties{
			Bucket: obj.AWSProperties.Bucket,
		},
	}
}

// NewAzureCloudExport creates a new AZURE CloudExport with all required fields set.
func NewAzureCloudExport(obj CloudExportAzureRequiredFields) *CloudExport {
	return &CloudExport{
		Name:          obj.Name,
		PlanID:        obj.PlanID,
		CloudProvider: CloudProviderAzure,
		AzureProperties: &AzureProperties{
			Location:       obj.AzureProperties.Location,
			ResourceGroup:  obj.AzureProperties.ResourceGroup,
			StorageAccount: obj.AzureProperties.StorageAccount,
			SubscriptionID: obj.AzureProperties.SubscriptionID,
		},
	}
}

// NewGCECloudExport creates a new GCE CloudExport with all required fields set.
func NewGCECloudExport(obj CloudExportGCERequiredFields) *CloudExport {
	return &CloudExport{
		Name:          obj.Name,
		PlanID:        obj.PlanID,
		CloudProvider: CloudProviderGCE,
		GCEProperties: &GCEProperties{
			Project:      obj.GCEProperties.Project,
			Subscription: obj.GCEProperties.Subscription,
		},
	}
}

// NewIBMCloudExport creates a new IBM CloudExport with all required fields set.
func NewIBMCloudExport(obj CloudExportIBMRequiredFields) *CloudExport {
	return &CloudExport{
		Name:          obj.Name,
		PlanID:        obj.PlanID,
		CloudProvider: CloudProviderIBM,
		IBMProperties: &IBMProperties{
			Bucket: obj.IBMProperties.Bucket,
		},
	}
}

// GetAllCloudExportsResponse model.
type GetAllCloudExportsResponse struct {
	// CloudExports holds all cloud export tasks.
	CloudExports []CloudExport
	// InvalidCloudExportsCount is a number of invalid cloud export tasks.
	InvalidCloudExportsCount uint32
}

// CloudExport defines a cloud export task.
type CloudExport struct {
	// Read-write properties

	// Type of export task.
	Type CloudExportType
	// Enabled specifies whether this task is enabled and intended to run or disabled.
	Enabled *bool
	// Name is a short name for this export task.
	Name string
	// Description is an optional, longer description of this export task.
	Description string
	// PlanID is the identifier of the Kentik plan associated with this export.
	PlanID string
	// CloudProvider is the cloud provider targeted by this export, e.g. AWS, Azure, GCE, IBM.
	CloudProvider CloudProvider
	// AWSProperties are specific to Amazon Web Services VPC flow logs exports.
	AWSProperties *AWSProperties
	// AzureProperties are specific to Azure exports.
	AzureProperties *AzureProperties
	// GCEProperties are specific to Google Cloud export.
	GCEProperties *GCEProperties
	// IBMProperties are specific to IBM Cloud exports.
	IBMProperties *IBMProperties
	// BGPProperties are optional BGP related settings.
	BGP *BGPProperties

	// Read-only properties

	// ID is the internal cloud export identifier. This is read-only and assigned by Kentik.
	ID ID
	// CurrentStatus is the most current status Kentik has about this export. This is read-only and assigned by Kentik.
	CurrentStatus *CloudExportStatus
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

// GCEProperties are specific to Google Cloud export.
type GCEProperties struct {
	// Project is a GCE project name.
	Project string
	// Subscription is a GCE subscription name.
	Subscription string
}

// IBMProperties are specific to IBM Cloud exports.
type IBMProperties struct {
	// Bucket is an IBM bucket.
	Bucket string
}

// BGPProperties are optional BGP related settings.
type BGPProperties struct {
	// ApplyBGP enables applying BGP data discovered via another device to the flow from this export.
	ApplyBGP *bool
	// UseBGPDeviceID specifies which other device to get BGP data from.
	UseBGPDeviceID string
	DeviceBGPType  string
}

// CloudExportStatus is export task status.
type CloudExportStatus struct {
	Status string
	// ErrorMessage holds current error information.
	ErrorMessage string
	// FlowFound informs whether flow logs were found.
	FlowFound            *bool
	APIAccess            *bool
	StorageAccountAccess *bool
}

// CloudExportAWSRequiredFields is subset of fields required to create an AWS CloudExport.
type CloudExportAWSRequiredFields struct {
	Name          string
	PlanID        string
	AWSProperties AWSPropertiesRequiredFields
}

// AWSPropertiesRequiredFields is subset of AWSProperties required to create an AWS CloudExport.
type AWSPropertiesRequiredFields struct {
	Bucket string
}

// CloudExportAzureRequiredFields is subset of fields required to create an Azure CloudExport.
type CloudExportAzureRequiredFields struct {
	Name            string
	PlanID          string
	AzureProperties AzurePropertiesRequiredFields
}

// AzurePropertiesRequiredFields is subset of AzureProperties required to create an Azure CloudExport.
type AzurePropertiesRequiredFields struct {
	Location       string
	ResourceGroup  string
	StorageAccount string
	SubscriptionID string
}

// CloudExportGCERequiredFields is subset of fields required to create a GCE CloudExport.
type CloudExportGCERequiredFields struct {
	Name          string
	PlanID        string
	GCEProperties GCEPropertiesRequiredFields
}

// GCEPropertiesRequiredFields is subset of GCEProperties required to create a GCE CloudExport.
type GCEPropertiesRequiredFields struct {
	Project      string
	Subscription string
}

// CloudExportIBMRequiredFields is subset of fields required to create an IBM CloudExport.
type CloudExportIBMRequiredFields struct {
	Name          string
	PlanID        string
	IBMProperties IBMPropertiesRequiredFields
}

// IBMPropertiesRequiredFields is subset of IBMProperties required to create an IBM CloudExport.
type IBMPropertiesRequiredFields struct {
	Bucket string
}

// CloudExportType is the type of export task.
type CloudExportType string

const (
	// CloudExportTypeUnspecified is invalid or incomplete cloud export.
	CloudExportTypeUnspecified = "CLOUD_EXPORT_TYPE_UNSPECIFIED"
	// CloudExportTypeKentikManaged is for cloud exports that are managed by Kentik.
	CloudExportTypeKentikManaged = "CLOUD_EXPORT_TYPE_KENTIK_MANAGED"
	// CloudExportTypeCustomerManaged is for cloud exports that are managed by Kentik customers,
	// e.g. by running an agent.
	CloudExportTypeCustomerManaged = "CLOUD_EXPORT_TYPE_CUSTOMER_MANAGED"
)

// CloudProvider is the cloud provider name.
type CloudProvider string

const (
	CloudProviderAWS   = "aws"
	CloudProviderAzure = "azure"
	CloudProviderGCE   = "gce"
	CloudProviderIBM   = "ibm"
)
