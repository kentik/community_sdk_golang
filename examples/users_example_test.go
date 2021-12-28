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

func TestDemonstrateUsersCRUD(t *testing.T) {
	t.Parallel()
	err := demonstrateUsersCRUD()
	assert.NoError(t, err)
}

func TestDemonstrateUsersGetAll(t *testing.T) {
	t.Parallel()
	err := demonstrateUsersGetAll()
	assert.NoError(t, err)
}

func demonstrateUsersCRUD() error {
	ctx := context.Background()
	client, err := NewClient()
	if err != nil {
		return err
	}

	fmt.Println("### client.Users.Create")
	user, err := client.Users.Create(ctx, *models.NewUser(models.UserRequiredFields{
		Username:     "testuser",
		UserFullName: "Test User",
		UserEmail:    "test@user.example",
		Role:         "Member",
		EmailService: true,
		EmailProduct: true,
	}))
	if err != nil {
		return fmt.Errorf("client.Users.Create failed: %w", err)
	}

	PrettyPrint(user)
	fmt.Println()

	fmt.Println("### client.Users.Get")
	user, err = client.Users.Get(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("client.Users.Get failed: %w", err)
	}

	PrettyPrint(user)
	fmt.Println()

	fmt.Println("### client.Users.Update")
	user.UserFullName = "Updated User"
	user.EmailProduct = false
	user, err = client.Users.Update(ctx, *user)
	if err != nil {
		return fmt.Errorf("client.Users.Update failed: %w", err)
	}

	PrettyPrint(user)
	fmt.Println()

	fmt.Println("### client.Users.Delete")
	err = client.Users.Delete(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("client.Users.Delete failed: %w", err)
	}

	return nil
}

func demonstrateUsersGetAll() error {
	ctx := context.Background()
	client, err := NewClient()
	if err != nil {
		return err
	}

	fmt.Println("### client.Users.GetAll")
	users, err := client.Users.GetAll(ctx)
	if err != nil {
		return fmt.Errorf("client.Users.GetAll failed: %w", err)
	}
	PrettyPrint(users)

	return nil
}
