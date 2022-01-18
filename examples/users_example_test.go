//go:build examples
// +build examples

//nolint:testpackage,forbidigo
package examples

import (
	"context"
	"fmt"
	"testing"

	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
)

func TestDemonstrateUsersAPI(t *testing.T) {
	t.Parallel()
	err := demonstrateUsersAPI()
	assert.NoError(t, err)
}

func demonstrateUsersAPI() error {
	ctx := context.Background()
	client, err := NewClient()
	if err != nil {
		return err
	}

	fmt.Println("Invoking client.Users.GetAll")
	users, err := client.Users.GetAll(ctx)
	if err != nil {
		return fmt.Errorf("client.Users.GetAll: %w", err)
	}
	PrettyPrint(users)

	fmt.Println("Invoking client.Users.Create")
	user, err := client.Users.Create(ctx, *models.NewUser(models.UserRequiredFields{
		Username:     "test-user",
		UserFullName: "Test User",
		UserEmail:    "test@user.example",
		Role:         "Member",
		EmailService: true,
		EmailProduct: true,
	}))
	if err != nil {
		return fmt.Errorf("client.Users.Create: %w", err)
	}

	PrettyPrint(user)
	fmt.Println()

	fmt.Println("Invoking client.Users.Get")
	user, err = client.Users.Get(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("client.Users.Get: %w", err)
	}

	PrettyPrint(user)
	fmt.Println()

	fmt.Println("Invoking client.Users.Update")
	user.UserFullName = "Updated User"
	user.EmailProduct = false
	user, err = client.Users.Update(ctx, *user)
	if err != nil {
		return fmt.Errorf("client.Users.Update: %w", err)
	}

	PrettyPrint(user)
	fmt.Println()

	fmt.Println("Invoking client.Users.Delete")
	err = client.Users.Delete(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("client.Users.Delete: %w", err)
	}

	fmt.Println("client.Users.Delete succeeded")
	return nil
}
