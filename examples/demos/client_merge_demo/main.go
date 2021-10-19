package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/kentik/community_sdk_golang/examples/demos"
	"github.com/kentik/community_sdk_golang/kentikapi"
	"log"
	"os"
)

func main() {
	err := showClientMerge()
	if err != nil {
		log.Fatal(err)
	}
}

func showClientMerge() error {
	demos.Step("Using Kentik API server")
	// import "github.com/kentik/community_sdk_golang/apiv6/kentikapi" - synthetics, cloud export
	// import "github.com/kentik/community_sdk_golang/kentikapi" - users, devices, etc.

	email, token, err := ReadCredentialsFromEnv()
	if err != nil {
		return err
	}

	demos.Step("Create Kentik API client")

	c := kentikapi.NewClient(kentikapi.Config{
		AuthEmail: email,
		AuthToken: token,
	})

	demos.Step("List users using API v5")
	result, err := c.Users.GetAll(context.Background())

	fmt.Println("Received result:")
	demos.PrettyPrint(result)

	demos.Step("List agents using API v6")
	getResp, _, err := c.SyntheticsAdminServiceAPI.AgentsList(context.Background()).Execute()

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
