package main

import (
	"context"
	"fmt"
	"os"

	"github.com/kentik/community_sdk_golang/apiv6/examples"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/cloudexport"
)

func main() {
	client := examples.NewClient()
	var err error

	if err = runCRUD(client); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = runGetAll(client); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runCRUD(client *kentikapi.Client) error {
	export := cloudexport.NewV202101beta1CloudExport()

	fmt.Println("### CREATE")
	gce := *cloudexport.NewV202101beta1GceProperties()
	gce.SetProject("test gce project")
	gce.SetSubscription("test gce subscription")
	export.SetGce(gce)
	export.SetCloudProvider("gce")
	export.SetName("test_gce_export")
	export.SetPlanId("11467")
	export.SetType(cloudexport.KENTIK_MANAGED)
	createReqPayload := *cloudexport.NewV202101beta1CreateCloudExportRequest()
	createReqPayload.Export = export
	createReq := client.CloudExportAdminServiceApi.CloudExportAdminServiceCreateCloudExport(context.Background()).V202101beta1CreateCloudExportRequest(createReqPayload)
	createResp, httpResp, err := createReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	examples.PrettyPrint(createResp)
	fmt.Println()

	fmt.Println("### UPDATE")
	export.SetDescription("Updated description")
	updateReqPayload := *cloudexport.NewV202101beta1UpdateCloudExportRequest()
	updateReqPayload.Export = export
	updateReq := client.CloudExportAdminServiceApi.CloudExportAdminServiceUpdateCloudExport(context.Background(), *createResp.Export.Id).V202101beta1UpdateCloudExportRequest(updateReqPayload)
	updateResp, httpResp, err := updateReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	examples.PrettyPrint(updateResp)
	fmt.Println()

	fmt.Println("### GET")
	getReq := client.CloudExportAdminServiceApi.CloudExportAdminServiceGetCloudExport(context.Background(), *createResp.Export.Id)
	getResp, httpResp, err := getReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	examples.PrettyPrint(getResp)
	fmt.Println()

	fmt.Println("### DELETE")
	deleteReq := client.CloudExportAdminServiceApi.CloudExportAdminServiceDeleteCloudExport(context.Background(), *createResp.Export.Id)
	deleteResp, httpResp, err := deleteReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	fmt.Println("Success")
	examples.PrettyPrint(deleteResp)
	fmt.Println()

	return nil
}

func runGetAll(client *kentikapi.Client) error {
	fmt.Println("### GET ALL")
	getAllReq := client.CloudExportAdminServiceApi.CloudExportAdminServiceListCloudExport(context.Background())
	getAllResp, httpResp, err := getAllReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	exports := *getAllResp.Exports
	fmt.Println("Num exports:", len(exports))
	examples.PrettyPrint(exports)
	fmt.Println()

	return nil
}
