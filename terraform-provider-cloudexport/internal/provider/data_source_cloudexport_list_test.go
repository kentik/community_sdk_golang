package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// The below tests check if the provider retrieves from server list of all available resources.
// Note: values checked in below tests are provided by localhost_apiserver from CloudExportTestData.json (running in background)
// Note: only check that the expected items are on the returned list, the items detailed check is done in data_source_cloudexport_item_test.go

func TestDataSourceCloudExportList(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudExportDataSourceList,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.kentik-cloudexport_list.exports", "items.0.name", "test_terraform_aws_export"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_list.exports", "items.1.name", "test_terraform_gce_export"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_list.exports", "items.2.name", "test_terraform_ibm_export"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_list.exports", "items.3.name", "test_terraform_azure_export"),
					resource.TestCheckResourceAttr("data.kentik-cloudexport_list.exports", "items.4.name", "test_terraform_bgp_export"),
				),
			},
		},
	})
}

// for specific items attributes, see: CloudExportTestData.json
const testCloudExportDataSourceList = `
provider "kentik-cloudexport" {
	# apiurl = "http://localhost:8080" # KTAPI_URL env variable used instead
}
  
data "kentik-cloudexport_list" "exports" {}
`
