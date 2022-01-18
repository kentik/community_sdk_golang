//nolint:forbidigo
package main

import (
	"context"
	"fmt"
	"log"

	syntheticspb "github.com/kentik/api-schema-public/gen/go/kentik/synthetics/v202101beta1"
	"github.com/kentik/community_sdk_golang/examples/demos"
)

func main() {
	if err := showClientMerge(); err != nil {
		log.Fatal(err)
	}
}

func showClientMerge() error {
	demos.Step("Create Kentik API client")
	ctx := context.Background()
	client, err := demos.NewClient()
	if err != nil {
		return err
	}

	demos.Step("List users using API v5")
	result, err := client.Users.GetAll(ctx)
	if err != nil {
		return err
	}

	fmt.Println("Received result:")
	demos.PrettyPrint(result)

	demos.Step("List agents using API v6")
	getResp, err := client.SyntheticsAdmin.ListAgents(ctx, &syntheticspb.ListAgentsRequest{})

	fmt.Println("Received result:")
	demos.PrettyPrint(getResp.GetAgents())

	return err
}
