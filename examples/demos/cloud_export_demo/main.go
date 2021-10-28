package main

import (
	"context"
	"fmt"

	"github.com/kentik/community_sdk_golang/kentikapi/cloudexport"

	"github.com/kentik/community_sdk_golang/examples/demos"
	"github.com/kentik/community_sdk_golang/kentikapi"
)

func main() {
	client := demos.NewClient()

	demos.Step("Create a cloud export")
	exportId := createCloudExport(client)

	demos.Step("Get cloud export")
	export := getCloudExport(client, exportId)

	demos.Step("Update cloud export")
	updateCloudExport(client, export)

	demos.Step("Delete cloud export")
	deleteCloudExport(client, exportId)

	demos.Step("Get all cloud exports")
	getAllCloudExports(client)

	demos.Step("Finished!")
}

func createCloudExport(client *kentikapi.Client) string {
	gce := *cloudexport.NewV202101beta1GceProperties()
	gce.SetProject("demo gce project")
	gce.SetSubscription("demo gce subscription")
	export := cloudexport.NewV202101beta1CloudExport()
	export.SetGce(gce)
	export.SetCloudProvider("gce")
	export.SetName("demo_gce_export_3")
	export.SetPlanId("11467")
	export.SetType(cloudexport.V202101BETA1CLOUDEXPORTTYPE_KENTIK_MANAGED)
	createReqPayload := *cloudexport.NewV202101beta1CreateCloudExportRequest()
	createReqPayload.Export = export

	createResp, _, err := client.CloudExportAdminServiceAPI.
		ExportCreate(context.Background()).
		Body(createReqPayload).
		Execute()
	demos.ExitOnError(err)

	fmt.Printf("Successfuly created cloud export, ID = %s\n", *createResp.Export.Id)
	return *createResp.Export.Id
}

func getCloudExport(client *kentikapi.Client, id string) *cloudexport.V202101beta1CloudExport {
	fmt.Printf("Retrieving cloud export of ID = %s\n", id)

	getResp, _, err := client.CloudExportAdminServiceAPI.
		ExportGet(context.Background(), id).
		Execute()
	demos.ExitOnError(err)

	demos.PrettyPrint(getResp)
	return getResp.Export
}

func updateCloudExport(client *kentikapi.Client, export *cloudexport.V202101beta1CloudExport) {
	export.SetDescription("Updated description") // update description
	updateReqPayload := *cloudexport.NewV202101beta1UpdateCloudExportRequest()
	updateReqPayload.Export = export

	updateResp, _, err := client.CloudExportAdminServiceAPI.
		ExportUpdate(context.Background(), *export.Id).
		Body(updateReqPayload).
		Execute()
	demos.ExitOnError(err)

	demos.PrettyPrint(updateResp)
}

func deleteCloudExport(client *kentikapi.Client, id string) {
	fmt.Printf("Deleting cloud export of ID = %s\n", id)
	deleteResp, _, err := client.CloudExportAdminServiceAPI.
		ExportDelete(context.Background(), id).
		Execute()
	demos.ExitOnError(err)

	demos.PrettyPrint(deleteResp)
	fmt.Println("Successful")
}

func getAllCloudExports(client *kentikapi.Client) {
	getAllResp, _, err := client.CloudExportAdminServiceAPI.
		ExportList(context.Background()).
		Execute()
	demos.ExitOnError(err)

	exports := *getAllResp.Exports
	fmt.Println("Num exports:", len(exports))
	demos.PrettyPrint(exports)
	fmt.Println()
}
