package main

import (
	"context"
	"errors"
	"fmt"
	syntheticspb "github.com/kentik/api-schema-public/gen/go/kentik/synthetics/v202101beta1"
	"log"
	"os"

	"github.com/kentik/community_sdk_golang/examples/demos"
	"github.com/kentik/community_sdk_golang/kentikapi"
)

func main() {
	err := showClientMerge()
	if err != nil {
		log.Fatal(err)
	}
}

func showClientMerge() error {
	demos.Step("Using Kentik API server")

	email, token, err := ReadCredentialsFromEnv()
	if err != nil {
		return err
	}

	demos.Step("Create Kentik API client")

	c, err := kentikapi.NewClient(kentikapi.Config{
		AuthEmail: email,
		AuthToken: token,
	})
	if err != nil {
		log.Fatal(err)
	}

	demos.Step("List users using API v5")
	result, err := c.Users.GetAll(context.Background())

	fmt.Println("Received result:")
	demos.PrettyPrint(result)

	demos.Step("List agents using API v6")
	getResp, err := c.SyntheticsAdmin.ListAgents(context.Background(), &syntheticspb.ListAgentsRequest{})

	fmt.Println("Received result:")
	demos.PrettyPrint(getResp)

	return err
}

func ReadCredentialsFromEnv() (authEmail, authToken string, _ error) {
	authEmail, ok := os.LookupEnv("KTAPI_AUTH_EMAIL")
	if !ok || authEmail == "" {
		return "", "", errors.New("KTAPI_AUTH_EMAIL environment variable needs to be set")
	}

	authToken, ok = os.LookupEnv("KTAPI_AUTH_TOKEN")
	if !ok || authToken == "" {
		return "", "", errors.New("KTAPI_AUTH_TOKEN environment variable needs to be set")
	}

	return authEmail, authToken, nil
}
