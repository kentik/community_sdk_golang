//+build examples

package examples

import (
	"context"
	"fmt"
	"testing"

	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
)

func TestMyKentikPortalAPIExample(t *testing.T) {
	assert := assert.New(t)
	assert.NoError(runCRUDExample())
	assert.NoError(getAllTenants())
}

func runCRUDExample() error {
	client, err := NewClient()
	if err != nil {
		return err
	}
	var tenant_id models.ID = 577

	fmt.Println("### GET")
	tenant, err := client.MyKentikPortal.Get(context.Background(), tenant_id)
	if err != nil {
		return err
	}

	PrettyPrint(tenant)

	fmt.Println("### CREATE USER")
	user, err := client.MyKentikPortal.CreateTenantUser(context.Background(), tenant_id, "test1@user.com")
	if err != nil {
		return err
	}

	PrettyPrint(user)

	fmt.Println("### DELETE USER")
	err = client.MyKentikPortal.DeleteTenantUser(context.Background(), tenant_id, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func getAllTenants() error {
	client, err := NewClient()
	if err != nil {
		return err
	}

	fmt.Println("### GET ALL")
	tenants, err := client.MyKentikPortal.GetAll(context.Background())
	if err != nil {
		return err
	}

	PrettyPrint(tenants)
	return nil
}
