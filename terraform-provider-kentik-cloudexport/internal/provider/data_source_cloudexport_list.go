package provider

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi"
)

var cloudExportListSchema = map[string]*schema.Schema{
	"items": &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: makeCloudExportSchema(READ_LIST),
		},
	},
}

func dataSourceCloudExportList() *schema.Resource {
	return &schema.Resource{
		Description: "DataSource representing list of cloud exports",
		ReadContext: dataSourceCloudExportListRead,
		Schema:      cloudExportListSchema,
	}
}

func dataSourceCloudExportListRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*kentikapi.Client)
	req := client.CloudExportAdminServiceApi.CloudExportAdminServiceListCloudExport(context.Background())
	listResp, httpResp, err := req.Execute()
	if err != nil {
		return diagError("Failed to read cloud export list", err, httpResp)
	}

	if listResp.Exports != nil {
		numExports := len(*listResp.Exports)
		exports := make([]interface{}, numExports, numExports)
		for i, e := range *listResp.Exports {
			exports[i] = cloudExportToMap(&e)
		}

		if err := d.Set("items", exports); err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10)) // use unixtime as id to force list update every time terraform asks for the list

	return nil
}
