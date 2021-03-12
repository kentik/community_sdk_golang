package main

import (
	"context"
	"fmt"

	"github.com/kentik/community_sdk_golang/apiv6/kentikapi"
)

func main() {
	client := NewClient()
	getAllExports(client)
}

func getAllExports(client *kentikapi.Client) {
	req := client.CloudExportAdminServiceApi.CloudExportAdminServiceListCloudExport(context.Background())
	domainResponse, httpResponse, err := client.CloudExportAdminServiceApi.CloudExportAdminServiceListCloudExportExecute(req)

	if err != nil {
		fmt.Println(httpResponse)
		fmt.Println(err)
		return
	}

	exports := *domainResponse.Exports
	fmt.Println("Num exports:", len(exports))
	PrettyPrint(exports)
}
