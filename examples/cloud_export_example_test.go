//+build examples

package examples

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/cloudexport"
	"github.com/stretchr/testify/assert"
)

func TestCloudExportAPIExample(t *testing.T) {
	assert.NoError(t, runCRUDCloudExport())
	assert.NoError(t, runGetAllCloudExports())
}

func runCRUDCloudExport() error {
	client := NewClient()
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	export := cloudexport.NewV202101beta1CloudExport()

	fmt.Println("### CREATE")
	gce := *cloudexport.NewV202101beta1GceProperties()
	gce.SetProject("test gce project")
	gce.SetSubscription("test gce subscription")
	export.SetGce(gce)
	export.SetCloudProvider("gce")
	export.SetName("test_gce_export")
	export.SetPlanId("11467")
	export.SetType(cloudexport.V202101BETA1CLOUDEXPORTTYPE_KENTIK_MANAGED)
	createReqPayload := *cloudexport.NewV202101beta1CreateCloudExportRequest()
	createReqPayload.Export = export

	createResp, httpResp, err := client.CloudExportAdminServiceAPI.
		ExportCreate(ctx).
		Body(createReqPayload).
		Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	PrettyPrint(createResp)
	fmt.Println()

	created := createResp.Export

	fmt.Println("### UPDATE")
	created.SetDescription("Updated description")
	updateReqPayload := *cloudexport.NewV202101beta1UpdateCloudExportRequest()
	updateReqPayload.Export = created

	updateResp, httpResp, err := client.CloudExportAdminServiceAPI.
		ExportUpdate(ctx, *created.Id).
		Body(updateReqPayload).
		Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	PrettyPrint(updateResp)
	fmt.Println()

	fmt.Println("### GET")
	getResp, httpResp, err := client.CloudExportAdminServiceAPI.
		ExportGet(ctx, *created.Id).
		Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	PrettyPrint(getResp)
	fmt.Println()

	fmt.Println("### DELETE")
	deleteResp, httpResp, err := client.CloudExportAdminServiceAPI.
		ExportDelete(ctx, *created.Id).
		Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	fmt.Println("Success")
	PrettyPrint(deleteResp)
	fmt.Println()

	return nil
}

func runGetAllCloudExports() error {
	client := NewClient()
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	fmt.Println("### GET ALL")
	getAllResp, httpResp, err := client.CloudExportAdminServiceAPI.
		ExportList(ctx).
		Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	exports := *getAllResp.Exports
	fmt.Println("Num exports:", len(exports))
	PrettyPrint(exports)
	fmt.Println()

	return nil
}
