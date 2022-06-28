//go:build examples
// +build examples

//nolint:forbidigo,testpackage
package examples

import (
	"context"
	"fmt"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/kentik/community_sdk_golang/kentikapi/cloud"
	"github.com/stretchr/testify/assert"
)

func TestDemonstrateCloudAPIWithAWSExport(t *testing.T) {
	t.Parallel()
	err := demonstrateCloudAPIWithAWSExport()
	assert.NoError(t, err)
}

func TestDemonstrateCloudAPIWithAzureExport(t *testing.T) {
	t.Parallel()
	err := demonstrateCloudAPIWithAzureExport()
	assert.NoError(t, err)
}

func TestDemonstrateCloudAPIWithGCEExport(t *testing.T) {
	t.Parallel()
	err := demonstrateCloudAPIWithGCEExport()
	assert.NoError(t, err)
}

func TestDemonstrateCloudAPIWithIBMExport(t *testing.T) {
	t.Parallel()
	err := demonstrateCloudAPIWithIBMExport()
	assert.NoError(t, err)
}

func demonstrateCloudAPIWithAWSExport() error {
	ctx := context.Background()
	client, err := NewClient()
	if err != nil {
		return err
	}

	fmt.Println("### Getting all cloud exports")
	getAllResp, err := client.Cloud.GetAllExports(ctx)
	if err != nil {
		return fmt.Errorf("client.Cloud.GetAll: %w", err)
	}

	fmt.Println("Got all cloud exports:")
	PrettyPrint(getAllResp.Exports)
	fmt.Println("Number of cloud exports:", len(getAllResp.Exports))
	if getAllResp.InvalidExportsCount > 0 {
		fmt.Printf(
			"Kentik API returned %v invalid cloud exports. Please, contact Kentik support.\n",
			getAllResp.InvalidExportsCount,
		)
	}

	fmt.Println("### Creating AWS cloud export")
	ce := cloud.NewAWSExport(cloud.AWSExportRequiredFields{
		Name:   "go-sdk-example-aws-export",
		PlanID: "11467",
		AWSProperties: cloud.AWSPropertiesRequiredFields{
			Bucket: "dummy-bucket",
		},
	})
	ce.Type = cloud.ExportTypeKentikManaged
	ce.Enabled = pointer.ToBool(true)
	ce.Description = "Dummy AWS description"
	ce.GetAWSProperties().IAMRoleARN = "dummy-iam-role-arn"
	ce.GetAWSProperties().Region = "dummy-region"
	ce.GetAWSProperties().DeleteAfterRead = pointer.ToBool(true)
	ce.GetAWSProperties().MultipleBuckets = pointer.ToBool(true)
	ce.BGP = &cloud.BGPProperties{
		ApplyBGP:       pointer.ToBool(true),
		UseBGPDeviceID: "dummy-device-id",
		DeviceBGPType:  "dummy-device-bgp-type",
	}

	ce, err = client.Cloud.CreateExport(ctx, ce)
	if err != nil {
		return fmt.Errorf("client.Cloud.Create: %w", err)
	}

	fmt.Println("Created AWS cloud export:")
	PrettyPrint(ce)

	fmt.Println("### Getting AWS cloud export")
	ce, err = client.Cloud.GetExport(ctx, ce.ID)
	if err != nil {
		return fmt.Errorf("client.Cloud.Get: %w", err)
	}

	fmt.Println("Got AWS cloud export:")
	PrettyPrint(ce)

	fmt.Println("### Updating AWS cloud export")
	ce.Name = "go-sdk-example-updated-aws-export"
	ce.Description = "Updated description"
	ce.GetAWSProperties().Bucket = "updated-bucket"
	ce.BGP.UseBGPDeviceID = "updated-bgp-device-id"
	ce, err = client.Cloud.UpdateExport(ctx, ce)
	if err != nil {
		return fmt.Errorf("client.Cloud.Update: %w", err)
	}

	fmt.Println("Updated cloud export:")
	PrettyPrint(ce)

	fmt.Println("### Deleting AWS cloud export")
	err = client.Cloud.DeleteExport(ctx, ce.ID)
	if err != nil {
		return fmt.Errorf("client.Cloud.Delete: %w", err)
	}

	fmt.Println("Deleted AWS cloud export")
	return nil
}

func demonstrateCloudAPIWithAzureExport() error {
	ctx := context.Background()
	client, err := NewClient()
	if err != nil {
		return err
	}

	fmt.Println("### Creating Azure cloud export")
	ce := cloud.NewAzureExport(cloud.AzureExportRequiredFields{
		Name:   "go-sdk-example-azure-export",
		PlanID: "11467",
		AzureProperties: cloud.AzurePropertiesRequiredFields{
			Location:       "dummy-location",
			ResourceGroup:  "dummy-rg",
			StorageAccount: "dummy-sa",
			SubscriptionID: "dummy-sid",
		},
	})
	ce.Type = cloud.ExportTypeKentikManaged
	ce.Enabled = pointer.ToBool(true)
	ce.Description = "Dummy Azure description"
	ce.GetAzureProperties().SecurityPrincipalEnabled = pointer.ToBool(true)
	ce.BGP = &cloud.BGPProperties{
		ApplyBGP:       pointer.ToBool(true),
		UseBGPDeviceID: "dummy-device-id",
		DeviceBGPType:  "dummy-device-bgp-type",
	}
	ce, err = client.Cloud.CreateExport(ctx, ce)
	if err != nil {
		return fmt.Errorf("client.Cloud.Create: %w", err)
	}

	fmt.Println("Created Azure cloud export:")
	PrettyPrint(ce)

	fmt.Println("### Deleting Azure cloud export")
	err = client.Cloud.DeleteExport(ctx, ce.ID)
	if err != nil {
		return fmt.Errorf("client.Cloud.Delete: %w", err)
	}

	fmt.Println("Deleted Azure cloud export")
	return nil
}

func demonstrateCloudAPIWithGCEExport() error {
	ctx := context.Background()
	client, err := NewClient()
	if err != nil {
		return err
	}

	fmt.Println("### Creating GCE cloud export")
	ce := cloud.NewGCEExport(cloud.GCEExportRequiredFields{
		Name:   "go-sdk-example-gce-export",
		PlanID: "11467",
		GCEProperties: cloud.GCEPropertiesRequiredFields{
			Project:      "dummy-project",
			Subscription: "dummy-subscription",
		},
	})
	ce.Type = cloud.ExportTypeKentikManaged
	ce.Enabled = pointer.ToBool(true)
	ce.Description = "Dummy GCE description"
	ce.BGP = &cloud.BGPProperties{
		ApplyBGP:       pointer.ToBool(true),
		UseBGPDeviceID: "dummy-device-id",
		DeviceBGPType:  "dummy-device-bgp-type",
	}

	ce, err = client.Cloud.CreateExport(ctx, ce)
	if err != nil {
		return fmt.Errorf("client.Cloud.Create: %w", err)
	}

	fmt.Println("Created GCE cloud export:")
	PrettyPrint(ce)

	fmt.Println("### Deleting GCE cloud export")
	err = client.Cloud.DeleteExport(ctx, ce.ID)
	if err != nil {
		return fmt.Errorf("client.Cloud.Delete: %w", err)
	}

	fmt.Println("Deleted GCE cloud export")
	return nil
}

func demonstrateCloudAPIWithIBMExport() error {
	ctx := context.Background()
	client, err := NewClient()
	if err != nil {
		return err
	}

	fmt.Println("### Creating IBM cloud export")
	ce := cloud.NewIBMExport(cloud.IBMExportRequiredFields{
		Name:   "go-sdk-example-ibm-export",
		PlanID: "11467",
		IBMProperties: cloud.IBMPropertiesRequiredFields{
			Bucket: "dummy-bucket",
		},
	})
	ce.Type = cloud.ExportTypeKentikManaged
	ce.Enabled = pointer.ToBool(true)
	ce.Description = "Dummy IBM description"
	ce.BGP = &cloud.BGPProperties{
		ApplyBGP:       pointer.ToBool(true),
		UseBGPDeviceID: "dummy-device-id",
		DeviceBGPType:  "dummy-device-bgp-type",
	}

	ce, err = client.Cloud.CreateExport(ctx, ce)
	if err != nil {
		return fmt.Errorf("client.Cloud.Create: %w", err)
	}

	fmt.Println("Created IBM cloud export:")
	PrettyPrint(ce)

	fmt.Println("### Deleting IBM cloud export")
	err = client.Cloud.DeleteExport(ctx, ce.ID)
	if err != nil {
		return fmt.Errorf("client.Cloud.Delete: %w", err)
	}

	fmt.Println("Deleted IBM cloud export")
	return nil
}
