//nolint:forbidigo
package main

import (
	"context"
	"fmt"
	"log"

	synthetics "github.com/kentik/api-schema-public/gen/go/kentik/synthetics/v202101beta1"
	"github.com/kentik/community_sdk_golang/examples/demos"
)

func main() {
	if err := showGRPCClient(); err != nil {
		log.Fatal(err)
	}
}

func showGRPCClient() error {
	demos.Step("Create Kentik API client")
	ctx := context.Background()
	client, err := demos.NewClient()
	if err != nil {
		return err
	}

	demos.Step("List synthetic agents")
	result, err := client.SyntheticsAdmin.ListAgents(ctx, &synthetics.ListAgentsRequest{})

	fmt.Println("Received result:")
	demos.PrettyPrint(result.GetAgents())

	return err
}
