package main

import (
	"context"
	"fmt"
	"log"
	"time"

	synthetics "github.com/kentik/api-schema-public/gen/go/kentik/synthetics/v202101beta1"
	"github.com/kentik/community_sdk_golang/examples"
	"github.com/kentik/community_sdk_golang/examples/demos"
	"github.com/kentik/community_sdk_golang/kentikapi"
)

func main() {
	err := showGRPCClient()
	if err != nil {
		log.Fatal(err)
	}
}

func showGRPCClient() error {
	email, token, err := examples.ReadCredentialsFromEnv()
	if err != nil {
		return err
	}

	demos.Step("Create Kentik API gRPC client")

	c, err := kentikapi.NewClient(kentikapi.Config{
		AuthEmail: email,
		AuthToken: token,
	})
	if err != nil {
		log.Fatalln(err)
	}
	ctx, close := context.WithTimeout(context.Background(), 100*time.Second)
	defer close()

	demos.Step("List synthetic agents")
	result, err := c.SyntheticsAdmin.ListAgents(ctx, &synthetics.ListAgentsRequest{})

	fmt.Println("Received result:")
	demos.PrettyPrint(result.GetAgents())

	return err
}
