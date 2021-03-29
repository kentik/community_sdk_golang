package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// The below tests check if the provider retrieves from server and deserializes properly the requested resources.
// Note: values checked in below tests are provided by localhost_apiserver from CloudExportTestData.json (running in background)

func TestDataSourceCloudExportItemAWS(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudExportDataSourceItems,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.aws", "id", "1"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.aws", "type", "CLOUD_EXPORT_TYPE_KENTIK_MANAGED"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.aws", "enabled", "true"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.aws", "name", "test_terraform_aws_export"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.aws", "description", "terraform aws cloud export"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.aws", "api_root", "http://localhost:8080"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.aws", "flow_dest", "http://localhost:8080/flow"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.aws", "plan_id", "11467"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.aws", "cloud_provider", "aws"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.aws", "current_status.0.status", "OK"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.aws", "current_status.0.error_message", "No errors"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.aws", "current_status.0.flow_found", "true"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.aws", "current_status.0.api_access", "true"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.aws", "current_status.0.storage_account_access", "true"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.aws", "aws.0.bucket", "terraform-aws-bucket"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.aws", "aws.0.iam_role_arn", "arn:aws:iam::003740049406:role/trafficTerraformIngestRole"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.aws", "aws.0.region", "us-east-2"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.aws", "aws.0.delete_after_read", "false"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.aws", "aws.0.multiple_buckets", "false"),
				),
			},
		},
	})
}

func TestDataSourceCloudExportItemGCE(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudExportDataSourceItems,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.gce", "id", "2"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.gce", "type", "CLOUD_EXPORT_TYPE_CUSTOMER_MANAGED"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.gce", "enabled", "true"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.gce", "name", "test_terraform_gce_export"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.gce", "description", "terraform gce cloud export"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.gce", "api_root", "http://localhost:8080"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.gce", "flow_dest", "http://localhost:8080/flow"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.gce", "plan_id", "21600"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.gce", "cloud_provider", "gce"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.gce", "current_status.0.status", "NOK"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.gce", "current_status.0.error_message", "Timeout"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.gce", "current_status.0.flow_found", "false"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.gce", "current_status.0.api_access", "false"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.gce", "current_status.0.storage_account_access", "false"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.gce", "gce.0.project", "project gce"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.gce", "gce.0.subscription", "subscription gce"),
				),
			},
		},
	})
}

func TestDataSourceCloudExportItemIBM(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudExportDataSourceItems,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.ibm", "id", "3"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.ibm", "type", "CLOUD_EXPORT_TYPE_KENTIK_MANAGED"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.ibm", "enabled", "false"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.ibm", "name", "test_terraform_ibm_export"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.ibm", "description", "terraform ibm cloud export"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.ibm", "api_root", "http://localhost:8080"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.ibm", "flow_dest", "http://localhost:8080/flow"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.ibm", "plan_id", "11467"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.ibm", "cloud_provider", "ibm"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.ibm", "current_status.0.status", "OK"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.ibm", "current_status.0.error_message", "No errors"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.ibm", "current_status.0.flow_found", "false"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.ibm", "current_status.0.api_access", "false"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.ibm", "current_status.0.storage_account_access", "false"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.ibm", "ibm.0.bucket", "terraform-ibm-bucket"),
				),
			},
		},
	})
}

func TestDataSourceCloudExportItemAzure(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudExportDataSourceItems,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.azure", "id", "4"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.azure", "type", "CLOUD_EXPORT_TYPE_KENTIK_MANAGED"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.azure", "enabled", "true"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.azure", "name", "test_terraform_azure_export"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.azure", "description", "terraform azure cloud export"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.azure", "api_root", "http://localhost:8080"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.azure", "flow_dest", "http://localhost:8080/flow"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.azure", "plan_id", "11467"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.azure", "cloud_provider", "azure"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.azure", "current_status.0.status", "OK"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.azure", "current_status.0.error_message", "No errors"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.azure", "current_status.0.flow_found", "false"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.azure", "current_status.0.api_access", "false"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.azure", "current_status.0.storage_account_access", "false"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.azure", "azure.0.location", "centralus"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.azure", "azure.0.resource_group", "traffic-generator"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.azure", "azure.0.storage_account", "kentikstorage"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.azure", "azure.0.subscription_id", "784bd5ec-122b-41b7-9719-22f23d5b49c8"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.azure", "azure.0.security_principal_enabled", "true"),
				),
			},
		},
	})
}

func TestDataSourceCloudExportItemBGP(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudExportDataSourceItems,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.bgp", "id", "5"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.bgp", "type", "CLOUD_EXPORT_TYPE_KENTIK_MANAGED"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.bgp", "enabled", "true"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.bgp", "name", "test_terraform_bgp_export"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.bgp", "description", "terraform bgp cloud export"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.bgp", "api_root", "http://localhost:8080"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.bgp", "flow_dest", "http://localhost:8080/flow"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.bgp", "plan_id", "11467"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.bgp", "cloud_provider", "bgp"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.bgp", "current_status.0.status", "OK"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.bgp", "current_status.0.error_message", "No errors"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.bgp", "current_status.0.flow_found", "false"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.bgp", "current_status.0.api_access", "false"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.bgp", "current_status.0.storage_account_access", "false"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.bgp", "bgp.0.apply_bgp", "true"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.bgp", "bgp.0.use_bgp_device_id", "1324"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_item.bgp", "bgp.0.device_bgp_type", "router"),
				),
			},
		},
	})
}

// for specific items attributes, see: CloudExportTestData.json
const testCloudExportDataSourceItems = `
provider "kentik-cloudexport" {
	# apiurl = "http://localhost:8080" # KTAPI_URL env variable used instead
}
  
data "kentik-cloudexport_item" "aws" {
	id = "1"
}

data "kentik-cloudexport_item" "gce" {
	id = "2"
}

data "kentik-cloudexport_item" "ibm" {
	id = "3"
}

data "kentik-cloudexport_item" "azure" {
	id = "4"
}

data "kentik-cloudexport_item" "bgp" {
	id = "5"
}
`
