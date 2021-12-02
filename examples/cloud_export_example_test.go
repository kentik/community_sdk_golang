//go:build examples
// +build examples

package examples

import (
	"context"
	"fmt"
	"testing"
	"time"

	cloudexportpb "github.com/kentik/api-schema-public/gen/go/kentik/cloud_export/v202101beta1"
	"github.com/stretchr/testify/assert"
)

func TestCloudExportAPIExample(t *testing.T) {
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
