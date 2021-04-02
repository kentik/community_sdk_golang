package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/cloudexport"
)

func resourceCloudExport() *schema.Resource {
	return &schema.Resource{
		Description: "Resource representing cloud export item",

		CreateContext: resourceCloudExportCreate,
		ReadContext:   resourceCloudExportRead,
		UpdateContext: resourceCloudExportUpdate,
		DeleteContext: resourceCloudExportDelete,

		Schema: makeCloudExportSchema(CREATE),
	}
}

func resourceCloudExportCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*kentikapi.Client)

	export, err := resourceDataToCloudExport(d)
	if err != nil {
		return diag.FromErr(err)
	}
	createReqPayload := *cloudexport.NewV202101beta1CreateCloudExportRequest()
	createReqPayload.Export = export
	createReq := client.CloudExportAdminServiceApi.CloudExportAdminServiceCreateCloudExport(ctx).V202101beta1CreateCloudExportRequest(createReqPayload)
	createResp, httpResp, err := createReq.Execute()
	if err != nil {
		return diagError("Failed to create cloud export", err, httpResp)
	}

	d.Set("id", *createResp.Export.Id)
	d.SetId(*createResp.Export.Id) // set internal ID, so Terraform knows the resource is now present on server side

	return resourceCloudExportRead(ctx, d, m) // read back the just-created resource to handle the case when server applies modifications to provided data eg strings etc
}

func resourceCloudExportRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*kentikapi.Client)

	exportID := d.Get("id").(string)
	req := client.CloudExportAdminServiceApi.CloudExportAdminServiceGetCloudExport(ctx, exportID)
	getResp, httpResp, err := req.Execute()
	if err != nil {
		return diagError("Failed to read cloud export", err, httpResp)
	}

	mapExport := cloudExportToMap(getResp.Export)
	for k, v := range mapExport {
		if err := d.Set(k, v); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceCloudExportUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// check if any attribute has changed
	if d.HasChange("") {
		export, err := resourceDataToCloudExport(d)
		if err != nil {
			return diag.FromErr(err)
		}
		updateReqPayload := *cloudexport.NewV202101beta1UpdateCloudExportRequest()
		updateReqPayload.Export = export
		client := m.(*kentikapi.Client)
		exportID := d.Get("id").(string)
		updateReq := client.CloudExportAdminServiceApi.CloudExportAdminServiceUpdateCloudExport(ctx, exportID).V202101beta1UpdateCloudExportRequest(updateReqPayload)
		_, httpResp, err := updateReq.Execute()
		if err != nil {
			return diagError("Failed to update cloud export", err, httpResp)
		}
	}

	return resourceCloudExportRead(ctx, d, m) // read back the just-created resource to handle the case when server applies modifications to provided data eg strings etc
}

func resourceCloudExportDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*kentikapi.Client)
	exportID := d.Get("id").(string)

	deleteReq := client.CloudExportAdminServiceApi.CloudExportAdminServiceDeleteCloudExport(ctx, exportID)
	_, httpResp, err := deleteReq.Execute()
	if err != nil {
		return diagError("Failed to delete cloud export", err, httpResp)
	}

	// APIv6 HTTP DELETE used to return 200 OK but not delete the item. Check item was actually deleted
	resourceStillExists := resourceCloudExportRead(ctx, d, m) == nil // no error -> item still exists
	if resourceStillExists {
		return diag.Errorf("API responded with success, but the resource didn't actually get deleted. This is API error")
	}

	return nil
}
