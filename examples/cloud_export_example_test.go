//go:build examples
// +build examples

package examples

import (
	"context"
	"fmt"
	"testing"
	"time"

	cloudexportpb "github.com/kentik/api-schema-public/gen/go/kentik/cloud_export/v202101beta1"
	"github.com/kentik/community_sdk_golang/kentikapi/cloudexport"
	"github.com/stretchr/testify/assert"
)

func TestCloudExportAPIExample(t *testing.T) {
	assert.NoError(t, runCRUDCloudExport())
	assert.NoError(t, runGetAllCloudExports())
}

func runCRUDCloudExport() error {
	client, err := NewClient()
	if err != nil {
		return err
	}

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
	client, err := NewClient()
	if err != nil {
		return err
	}

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

func TestCloudExportGRPCAPIExample(t *testing.T) {
	assert.NoError(t, runGRPCCRUDCloudExport())
	assert.NoError(t, runGRPCGetAllCloudExports())
}

func runGRPCCRUDCloudExport() error {
	client, err := NewClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	fmt.Println("### CREATE")
	gce := &cloudexportpb.GceProperties{
		Project:      "test gce project",
		Subscription: "test gce subsribtion",
	}
	export := &cloudexportpb.CloudExport{
		Type:          cloudexportpb.CloudExportType_CLOUD_EXPORT_TYPE_KENTIK_MANAGED,
		Name:          "test_gce_export",
		PlanId:        "11467",
		CloudProvider: "gce",
		Properties:    &cloudexportpb.CloudExport_Gce{gce},
	}
	createReqPayload := &cloudexportpb.CreateCloudExportRequest{Export: export}

	createResp, err := client.CloudExportAdmin.CreateCloudExport(ctx, createReqPayload)
	if err != nil {
		return err
	}
	PrettyPrint(createResp.GetExport())
	fmt.Println()

	created := createResp.GetExport()

	fmt.Println("### UPDATE")
	created.Description = "Updated description"
	updateReqPayload := &cloudexportpb.UpdateCloudExportRequest{
		Export: created,
	}

	updateResp, err := client.CloudExportAdmin.UpdateCloudExport(ctx, updateReqPayload)
	if err != nil {
		return err
	}
	PrettyPrint(updateResp.GetExport())
	fmt.Println()

	fmt.Println("### GET")
	getReqPayLoad := &cloudexportpb.GetCloudExportRequest{Id: created.GetId()}
	getResp, err := client.CloudExportAdmin.GetCloudExport(ctx, getReqPayLoad)
	if err != nil {
		return err
	}
	PrettyPrint(getResp.GetExport())
	fmt.Println()

	fmt.Println("### DELETE")
	deleteReqPayLoad := &cloudexportpb.DeleteCloudExportRequest{Id: created.Id}
	_, err = client.CloudExportAdmin.DeleteCloudExport(ctx, deleteReqPayLoad)
	if err != nil {
		return err
	}
	fmt.Println("Success")
	fmt.Println()

	return nil
}

func runGRPCGetAllCloudExports() error {
	client, err := NewClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	fmt.Println("### GET ALL")
	getAllReqPayLoad := &cloudexportpb.ListCloudExportRequest{}
	getAllResp, err := client.CloudExportAdmin.ListCloudExport(ctx, getAllReqPayLoad)
	if err != nil {
		return err
	}
	exports := getAllResp.GetExports()
	fmt.Println("Num exports:", len(exports))
	PrettyPrint(exports)
	fmt.Println()

	return nil
}
