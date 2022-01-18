//go:build examples
// +build examples

//nolint:testpackage,forbidigo
package examples

import (
	"context"
	"fmt"
	"testing"

	cloudexportpb "github.com/kentik/api-schema-public/gen/go/kentik/cloud_export/v202101beta1"
	"github.com/stretchr/testify/assert"
)

func TestDemonstrateCloudExportAPI(t *testing.T) {
	t.Parallel()
	err := demonstrateCloudExportAPI()
	assert.NoError(t, err)
}

func demonstrateCloudExportAPI() error {
	ctx := context.Background()
	client, err := NewClient()
	if err != nil {
		return err
	}

	fmt.Println("Invoking client.CloudExportAdmin.ListCloudExport")
	getAllResp, err := client.CloudExportAdmin.ListCloudExport(ctx, &cloudexportpb.ListCloudExportRequest{})
	if err != nil {
		return fmt.Errorf("client.CloudExportAdmin.ListCloudExport: %w", err)
	}

	fmt.Println("Number of exports:", len(getAllResp.GetExports()))
	fmt.Println("Invalid exports count:", getAllResp.GetInvalidExportsCount())
	PrettyPrint(getAllResp.GetExports())
	fmt.Println()

	fmt.Println("Invoking client.CloudExportAdmin.CreateCloudExport")
	createResp, err := client.CloudExportAdmin.CreateCloudExport(ctx, &cloudexportpb.CreateCloudExportRequest{
		Export: &cloudexportpb.CloudExport{
			Type:          cloudexportpb.CloudExportType_CLOUD_EXPORT_TYPE_KENTIK_MANAGED,
			Name:          "test_gce_export",
			PlanId:        "11467",
			CloudProvider: "gce",
			Properties: &cloudexportpb.CloudExport_Gce{
				Gce: &cloudexportpb.GceProperties{
					Project:      "test gce project",
					Subscription: "test gce subscription",
				},
			},
		},
	})
	if err != nil {
		return fmt.Errorf("client.CloudExportAdmin.CreateCloudExport: %w", err)
	}

	PrettyPrint(createResp.GetExport())
	fmt.Println()

	fmt.Println("Invoking client.CloudExportAdmin.UpdateCloudExport")
	export := createResp.GetExport()
	export.Description = "Updated description"
	updateResp, err := client.CloudExportAdmin.UpdateCloudExport(ctx, &cloudexportpb.UpdateCloudExportRequest{
		Export: export,
	})
	if err != nil {
		return fmt.Errorf("client.CloudExportAdmin.UpdateCloudExport: %w", err)
	}

	PrettyPrint(updateResp.GetExport())
	fmt.Println()

	fmt.Println("Invoking client.CloudExportAdmin.DeleteCloudExport")
	_, err = client.CloudExportAdmin.DeleteCloudExport(ctx, &cloudexportpb.DeleteCloudExportRequest{
		Id: export.Id,
	})
	if err != nil {
		return fmt.Errorf("client.CloudExportAdmin.DeleteCloudExport: %w", err)
	}

	fmt.Println("client.CloudExportAdmin.DeleteCloudExport succeeded")
	return nil
}
