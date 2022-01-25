//go:build examples
// +build examples

//nolint:testpackage,forbidigo
package examples

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
)

func TestMyKentikPortalAPIExample(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	assert.NoError(runCRUDExample())
	assert.NoError(getAllTenants())
}

func runCRUDExample() error {
	client, err := NewClient()
	if err != nil {
		return err
	}
	tenantID, err := pickTenantID()
	if err != nil {
		return err
	}

	fmt.Println("### GET")
	tenant, err := client.MyKentikPortal.Get(context.Background(), tenantID)
	if err != nil {
		return err
	}

	PrettyPrint(tenant)

	fmt.Println("### CREATE USER")
	user, err := client.MyKentikPortal.CreateTenantUser(context.Background(), tenantID, "test1@user.com")
	if err != nil {
		return err
	}

	PrettyPrint(user)

	fmt.Println("### DELETE USER")
	err = client.MyKentikPortal.DeleteTenantUser(context.Background(), tenantID, user.ID)
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

func pickTenantID() (models.ID, error) {
	client, err := NewClient()
	if err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	tenants, err := client.MyKentikPortal.GetAll(ctx)
	if err != nil {
		return "", err
	}

	if tenants != nil {
		return tenants[0].ID, nil
	}
	return "", fmt.Errorf("No tenants in requested Kentik account: %v", err)
}
