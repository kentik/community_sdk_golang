package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi"
)

func dataSourceCloudExportItem() *schema.Resource {
	return &schema.Resource{
		Description: "DataSource representing single cloud export item",
		ReadContext: dataSourceCloudExportItemRead,
		Schema:      makeCloudExportSchema(READ_SINGLE),
	}
}

func dataSourceCloudExportItemRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*kentikapi.Client)
	exportID := d.Get("id").(string)
	req := client.CloudExportAdminServiceApi.CloudExportAdminServiceGetCloudExport(context.Background(), exportID)
	getResp, httpResp, err := req.Execute()
	if err != nil {
		return diagError("Failed to read cloud export item", err, httpResp)
	}

	mapExport := cloudExportToMap(getResp.Export)
	for k, v := range mapExport {
		if err := d.Set(k, v); err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(*getResp.Export.Id)

	return nil
}
