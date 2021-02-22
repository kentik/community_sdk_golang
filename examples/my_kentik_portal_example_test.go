//+build examples

package examples

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyKentikPortalAPIExample(t *testing.T) {
	assert := assert.New(t)
	assert.NoError(runCRUDexample())
	assert.NoError(getAllTenants())
}

func runCRUDexample() error {
	client := NewClient()

	fmt.Println("### GET")
	tenant, err := client.MyKentikPortal.Get(context.Background(), 577)
	if err != nil {
		return err
	}

	PrettyPrint(tenant)

	fmt.Println("### CREATE USER")
	user, err := client.MyKentikPortal.CreateTenantUser(context.Background(), 577, "test1@user.com")
	if err != nil {
		return err
	}

	PrettyPrint(user)

	fmt.Println("### DELETE USER")
	err = client.MyKentikPortal.DeleteTenantUser(context.Background(), 577, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func getAllTenants() error {
	client := NewClient()

	fmt.Println("### GET ALL")
	tenants, err := client.MyKentikPortal.GetAll(context.Background())
	if err != nil {
		return err
	}

	PrettyPrint(tenants)
	return nil
}
