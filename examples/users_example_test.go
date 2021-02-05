//+build examples

package examples

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/kentik/community_sdk_golang/kentikapi"
)

func TestUsersAPIExample(t *testing.T) {
	if err := demonstrateUsersAPI(); err != nil {
		t.Fatal(err)
	}
}

func demonstrateUsersAPI() error {
	ctx := context.Background()
	authEmail, authToken, err := readCredentialsFromEnv()
	if err != nil {
		return fmt.Errorf("reading credentials failed: %v", err)
	}

	client := kentikapi.NewClient(kentikapi.Config{
		AuthEmail: authEmail,
		AuthToken: authToken,
	})

	fmt.Println("### UsersAPI.GetAll")
	users, err := client.UsersAPI.GetAll(ctx)
	if err != nil {
		return fmt.Errorf("UsersAPI.GetAll failed: %s", err)
	}
	for _, u := range users {
		fmt.Printf("%+v\n", u)
	}

	fmt.Println("\n### UsersAPI.Get")
	user, err := client.UsersAPI.Get(ctx, 149492)
	if err != nil {
		return fmt.Errorf("UsersAPI.Get failed: %s", err)
	}
	fmt.Printf("%+v\n", user)

	return nil
}

func readCredentialsFromEnv() (authEmail, authToken string, _ error) {
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
