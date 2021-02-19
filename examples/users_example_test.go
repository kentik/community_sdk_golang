//+build examples

package examples

import (
	"context"
	"fmt"
	"testing"
)

func TestUsersAPIExample(t *testing.T) {
	if err := demonstrateUsersAPI(); err != nil {
		t.Fatal(err)
	}
}

func demonstrateUsersAPI() error {
	ctx := context.Background()
	client := NewClient()

	fmt.Println("### UsersAPI.GetAll")
	users, err := client.UsersAPI.GetAll(ctx)
	if err != nil {
		return fmt.Errorf("UsersAPI.GetAll failed: %s", err)
	}
	PrettyPrint(users)

	fmt.Println("\n### UsersAPI.Get")
	user, err := client.UsersAPI.Get(ctx, 149492)
	if err != nil {
		return fmt.Errorf("UsersAPI.Get failed: %s", err)
	}
	PrettyPrint(user)

	return nil
}
