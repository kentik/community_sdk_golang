//+build examples

package examples

import (
	"context"
	"fmt"
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
	prettyPrint(users)

	fmt.Println("\n### UsersAPI.Get")
	user, err := client.UsersAPI.Get(ctx, 149492)
	if err != nil {
		return fmt.Errorf("UsersAPI.Get failed: %s", err)
	}
	prettyPrint(user)

	return nil
}


