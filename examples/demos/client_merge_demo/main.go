//nolint:forbidigo
package main

import (
	"context"
	"fmt"
	"log"

	syntheticspb "github.com/kentik/api-schema-public/gen/go/kentik/synthetics/v202101beta1"
	"github.com/kentik/community_sdk_golang/examples/demos"
	"github.com/kentik/community_sdk_golang/kentikapi"
)

func main() {
	if err := showClientMerge(); err != nil {
		log.Fatal(err)
	}
}

func showClientMerge() error {
	demos.Step("Using Kentik API server")

	email, token, err := kentikapi.ReadCredentialsFromEnv()
	if err != nil {
		return err
	}

	demos.Step("Create Kentik API client")

	c, err := kentikapi.NewClient(kentikapi.Config{
		AuthEmail: email,
		AuthToken: token,
	})
	if err != nil {
		return err
	}

	demos.Step("List users using API v5")
	result, err := c.Users.GetAll(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("Received result:")
	demos.PrettyPrint(result)

	demos.Step("List agents using API v6")
	getResp, err := c.SyntheticsAdmin.ListAgents(context.Background(), &syntheticspb.ListAgentsRequest{})

	fmt.Println("Received result:")
	demos.PrettyPrint(getResp.GetAgents())

	return err
}
