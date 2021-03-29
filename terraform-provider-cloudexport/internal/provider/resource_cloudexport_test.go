package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// The below tests check if the provider creates, updates, retrieves and deletes requested resources.
// Note: resource create/update/get/delete requests that result from below tests are processed by localhost_apiserver (running in background)
// Note: we only check the user-provided values as we don't control the server-provided ones

func TestResourceCloudExportAWS(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceCloudExportCreateAWS,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("kentik-cloudexport_item.test_aws", "id"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_aws", "type", "CLOUD_EXPORT_TYPE_KENTIK_MANAGED"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_aws", "enabled", "true"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_aws", "name", "resource_test_terraform_aws_export"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_aws", "description", "resource test aws export"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_aws", "plan_id", "9948"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_aws", "cloud_provider", "aws"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_aws", "aws.0.bucket", "resource-terraform-aws-bucket"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_aws", "aws.0.iam_role_arn", "arn:aws:iam::003740049406:role/trafficTerraformIngestRole"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_aws", "aws.0.region", "eu-central-1"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_aws", "aws.0.delete_after_read", "true"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_aws", "aws.0.multiple_buckets", "true"),
				),
			},
			{
				Config: testAccResourceCloudExportUpdateAWS,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("kentik-cloudexport_item.test_aws", "id"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_aws", "type", "CLOUD_EXPORT_TYPE_CUSTOMER_MANAGED"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_aws", "enabled", "false"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_aws", "name", "resource_test_terraform_aws_export_updated"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_aws", "description", "resource test aws export updated"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_aws", "plan_id", "3333"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_aws", "cloud_provider", "aws"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_aws", "aws.0.bucket", "resource-terraform-aws-bucket-updated"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_aws", "aws.0.iam_role_arn", "arn:aws:iam::003740049406:role/trafficTerraformIngestRole_updated"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_aws", "aws.0.region", "eu-central-1-updated"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_aws", "aws.0.delete_after_read", "false"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_aws", "aws.0.multiple_buckets", "false"),
				),
			},
			{
				Config: testAccResourceCloudExportDestroy,
				Check: resource.ComposeTestCheckFunc(
					testResourceDoesntExists("kentik-cloudexport_item.test_aws"),
				),
			},
		},
	})
}

func TestResourceCloudExportGCE(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceCloudExportCreateGCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("kentik-cloudexport_item.test_gce", "id"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_gce", "type", "CLOUD_EXPORT_TYPE_KENTIK_MANAGED"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_gce", "enabled", "true"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_gce", "name", "resource_test_terraform_gce_export"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_gce", "description", "resource test gce export"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_gce", "plan_id", "9948"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_gce", "cloud_provider", "gce"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_gce", "gce.0.project", "gce project"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_gce", "gce.0.subscription", "gce subscription"),
				),
			},
			{
				Config: testAccResourceCloudExportUpdateGCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("kentik-cloudexport_item.test_gce", "id"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_gce", "type", "CLOUD_EXPORT_TYPE_CUSTOMER_MANAGED"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_gce", "enabled", "false"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_gce", "name", "resource_test_terraform_gce_export_updated"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_gce", "description", "resource test gce export updated"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_gce", "plan_id", "3333"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_gce", "cloud_provider", "gce"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_gce", "gce.0.project", "gce project updated"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_gce", "gce.0.subscription", "gce subscription updated"),
				),
			},
			{
				Config: testAccResourceCloudExportDestroy,
				Check: resource.ComposeTestCheckFunc(
					testResourceDoesntExists("kentik-cloudexport_item.test_gce"),
				),
			},
		},
	})
}

func TestResourceCloudExportIBM(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceCloudExportCreateIBM,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("kentik-cloudexport_item.test_ibm", "id"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_ibm", "type", "CLOUD_EXPORT_TYPE_KENTIK_MANAGED"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_ibm", "enabled", "true"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_ibm", "name", "resource_test_terraform_ibm_export"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_ibm", "description", "resource test ibm export"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_ibm", "plan_id", "9948"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_ibm", "cloud_provider", "ibm"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_ibm", "ibm.0.bucket", "ibm-bucket"),
				),
			},
			{
				Config: testAccResourceCloudExportUpdateIBM,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("kentik-cloudexport_item.test_ibm", "id"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_ibm", "type", "CLOUD_EXPORT_TYPE_CUSTOMER_MANAGED"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_ibm", "enabled", "false"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_ibm", "name", "resource_test_terraform_ibm_export_updated"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_ibm", "description", "resource test ibm export updated"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_ibm", "plan_id", "3333"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_ibm", "cloud_provider", "ibm"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_ibm", "ibm.0.bucket", "ibm-bucket-updated"),
				),
			},
			{
				Config: testAccResourceCloudExportDestroy,
				Check: resource.ComposeTestCheckFunc(
					testResourceDoesntExists("kentik-cloudexport_item.test_ibm"),
				),
			},
		},
	})
}

func TestResourceCloudExportAzure(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceCloudExportCreateAzure,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("kentik-cloudexport_item.test_azure", "id"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_azure", "type", "CLOUD_EXPORT_TYPE_KENTIK_MANAGED"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_azure", "enabled", "true"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_azure", "name", "resource_test_terraform_azure_export"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_azure", "description", "resource test azure export"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_azure", "plan_id", "9948"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_azure", "cloud_provider", "azure"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_azure", "azure.0.location", "centralus"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_azure", "azure.0.resource_group", "traffic-generator"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_azure", "azure.0.storage_account", "kentikstorage"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_azure", "azure.0.subscription_id", "7777"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_azure", "azure.0.security_principal_enabled", "true"),
				),
			},
			{
				Config: testAccResourceCloudExportUpdateAzure,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("kentik-cloudexport_item.test_azure", "id"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_azure", "type", "CLOUD_EXPORT_TYPE_CUSTOMER_MANAGED"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_azure", "enabled", "false"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_azure", "name", "resource_test_terraform_azure_export_updated"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_azure", "description", "resource test azure export updated"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_azure", "plan_id", "3333"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_azure", "cloud_provider", "azure"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_azure", "azure.0.location", "centralus-updated"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_azure", "azure.0.resource_group", "traffic-generator-updated"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_azure", "azure.0.storage_account", "kentikstorage-updated"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_azure", "azure.0.subscription_id", "8888"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_azure", "azure.0.security_principal_enabled", "false"),
				),
			},
			{
				Config: testAccResourceCloudExportDestroy,
				Check: resource.ComposeTestCheckFunc(
					testResourceDoesntExists("kentik-cloudexport_item.test_azure"),
				),
			},
		},
	})
}

func TestResourceCloudExportBGP(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceCloudExportCreateBGP,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("kentik-cloudexport_item.test_bgp", "id"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_bgp", "type", "CLOUD_EXPORT_TYPE_KENTIK_MANAGED"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_bgp", "enabled", "true"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_bgp", "name", "resource_test_terraform_bgp_export"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_bgp", "description", "resource test bgp export"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_bgp", "plan_id", "9948"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_bgp", "cloud_provider", "bgp"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_bgp", "bgp.0.apply_bgp", "true"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_bgp", "bgp.0.use_bgp_device_id", "1324"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_bgp", "bgp.0.device_bgp_type", "router"),
				),
			},
			{
				Config: testAccResourceCloudExportUpdateBGP,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("kentik-cloudexport_item.test_bgp", "id"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_bgp", "type", "CLOUD_EXPORT_TYPE_CUSTOMER_MANAGED"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_bgp", "enabled", "false"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_bgp", "name", "resource_test_terraform_bgp_export_updated"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_bgp", "description", "resource test bgp export updated"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_bgp", "plan_id", "3333"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_bgp", "cloud_provider", "bgp"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_bgp", "bgp.0.apply_bgp", "false"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_bgp", "bgp.0.use_bgp_device_id", "4444"),
					resource.TestCheckResourceAttr("kentik-cloudexport_item.test_bgp", "bgp.0.device_bgp_type", "dns"),
				),
			},
			{
				Config: testAccResourceCloudExportDestroy,
				Check: resource.ComposeTestCheckFunc(
					testResourceDoesntExists("kentik-cloudexport_item.test_bgp"),
				),
			},
		},
	})
}

func testResourceDoesntExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		_, exists := s.RootModule().Resources[name]
		if exists {
			return fmt.Errorf("Resource %q found when not expected", name)
		}

		return nil
	}
}

const testAccResourceCloudExportCreateAWS = `
provider "kentik-cloudexport" {
	# apiurl = "http://localhost:8080" # KTAPI_URL env variable used instead
}

resource "kentik-cloudexport_item" "test_aws" {
	name= "resource_test_terraform_aws_export"
	type= "CLOUD_EXPORT_TYPE_KENTIK_MANAGED"
	enabled=true
	description= "resource test aws export"
	plan_id= "9948"
	cloud_provider= "aws"
	aws {
		bucket= "resource-terraform-aws-bucket"
		iam_role_arn= "arn:aws:iam::003740049406:role/trafficTerraformIngestRole"
		region= "eu-central-1"
		delete_after_read= true
		multiple_buckets= true
	}
  }
`
const testAccResourceCloudExportUpdateAWS = `
provider "kentik-cloudexport" {
	# apiurl = "http://localhost:8080" # KTAPI_URL env variable used instead
}

resource "kentik-cloudexport_item" "test_aws" {
	name= "resource_test_terraform_aws_export_updated"
	type= "CLOUD_EXPORT_TYPE_CUSTOMER_MANAGED"
	enabled=false
	description= "resource test aws export updated"
	plan_id= "3333"
	cloud_provider= "aws"
	aws {
		bucket= "resource-terraform-aws-bucket-updated"
		iam_role_arn= "arn:aws:iam::003740049406:role/trafficTerraformIngestRole_updated"
		region= "eu-central-1-updated"
		delete_after_read= false
		multiple_buckets= false
	}
  }
`

const testAccResourceCloudExportCreateGCE = `
provider "kentik-cloudexport" {
	# apiurl = "http://localhost:8080" # KTAPI_URL env variable used instead
}

resource "kentik-cloudexport_item" "test_gce" {
	name= "resource_test_terraform_gce_export"
	type= "CLOUD_EXPORT_TYPE_KENTIK_MANAGED"
	enabled=true
	description= "resource test gce export"
	plan_id= "9948"
	cloud_provider= "gce"
	gce {
		project= "gce project"
		subscription= "gce subscription"
	}
  }
`
const testAccResourceCloudExportUpdateGCE = `
provider "kentik-cloudexport" {
	# apiurl = "http://localhost:8080" # KTAPI_URL env variable used instead
}

resource "kentik-cloudexport_item" "test_gce" {
	name= "resource_test_terraform_gce_export_updated"
	type= "CLOUD_EXPORT_TYPE_CUSTOMER_MANAGED"
	enabled=false
	description= "resource test gce export updated"
	plan_id= "3333"
	cloud_provider= "gce"
	gce {
		project= "gce project updated"
		subscription= "gce subscription updated"
	}
  }
`

const testAccResourceCloudExportCreateIBM = `
provider "kentik-cloudexport" {
	# apiurl = "http://localhost:8080" # KTAPI_URL env variable used instead
}

resource "kentik-cloudexport_item" "test_ibm" {
	name= "resource_test_terraform_ibm_export"
	type= "CLOUD_EXPORT_TYPE_KENTIK_MANAGED"
	enabled=true
	description= "resource test ibm export"
	plan_id= "9948"
	cloud_provider= "ibm"
	ibm {
		bucket= "ibm-bucket"
	}
  }
`
const testAccResourceCloudExportUpdateIBM = `
provider "kentik-cloudexport" {
	# apiurl = "http://localhost:8080" # KTAPI_URL env variable used instead
}

resource "kentik-cloudexport_item" "test_ibm" {
	name= "resource_test_terraform_ibm_export_updated"
	type= "CLOUD_EXPORT_TYPE_CUSTOMER_MANAGED"
	enabled=false
	description= "resource test ibm export updated"
	plan_id= "3333"
	cloud_provider= "ibm"
	ibm {
		bucket= "ibm-bucket-updated"
	}
  }
`

const testAccResourceCloudExportCreateAzure = `
provider "kentik-cloudexport" {
	# apiurl = "http://localhost:8080" # KTAPI_URL env variable used instead
}

resource "kentik-cloudexport_item" "test_azure" {
	name= "resource_test_terraform_azure_export"
	type= "CLOUD_EXPORT_TYPE_KENTIK_MANAGED"
	enabled=true
	description= "resource test azure export"
	plan_id= "9948"
	cloud_provider= "azure"
	azure {
		location= "centralus"
		resource_group= "traffic-generator"
		storage_account= "kentikstorage"
		subscription_id= "7777"
		security_principal_enabled=true
	}
  }
`
const testAccResourceCloudExportUpdateAzure = `
provider "kentik-cloudexport" {
	# apiurl = "http://localhost:8080" # KTAPI_URL env variable used instead
}

resource "kentik-cloudexport_item" "test_azure" {
	name= "resource_test_terraform_azure_export_updated"
	type= "CLOUD_EXPORT_TYPE_CUSTOMER_MANAGED"
	enabled=false
	description= "resource test azure export updated"
	plan_id= "3333"
	cloud_provider= "azure"
	azure {
		location= "centralus-updated"
		resource_group= "traffic-generator-updated"
		storage_account= "kentikstorage-updated"
		subscription_id= "8888"
		security_principal_enabled=false
	}
  }
`

const testAccResourceCloudExportCreateBGP = `
provider "kentik-cloudexport" {
	# apiurl = "http://localhost:8080" # KTAPI_URL env variable used instead
}

resource "kentik-cloudexport_item" "test_bgp" {
	name= "resource_test_terraform_bgp_export"
	type= "CLOUD_EXPORT_TYPE_KENTIK_MANAGED"
	enabled=true
	description= "resource test bgp export"
	plan_id= "9948"
	cloud_provider= "bgp"
	bgp {
		apply_bgp= true
		use_bgp_device_id= "1324"
		device_bgp_type= "router"
	}
  }
`
const testAccResourceCloudExportUpdateBGP = `
provider "kentik-cloudexport" {
	# apiurl = "http://localhost:8080" # KTAPI_URL env variable used instead
}

resource "kentik-cloudexport_item" "test_bgp" {
	name= "resource_test_terraform_bgp_export_updated"
	type= "CLOUD_EXPORT_TYPE_CUSTOMER_MANAGED"
	enabled=false
	description= "resource test bgp export updated"
	plan_id= "3333"
	cloud_provider= "bgp"
	bgp {
		apply_bgp= false
		use_bgp_device_id= "4444"
		device_bgp_type= "dns"
	}
  }
`

const testAccResourceCloudExportDestroy = `
provider "kentik-cloudexport" {
	# apiurl = "http://localhost:8080" # KTAPI_URL env variable used instead
}
`
