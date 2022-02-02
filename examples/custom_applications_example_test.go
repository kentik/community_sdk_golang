//go:build examples
// +build examples

//nolint:testpackage,forbidigo
package examples

import (
	"context"
	"fmt"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
)

func TestCustomApplicationsAPIExample(t *testing.T) {
	t.Skipf("Kentik API is broken")

	t.Parallel()
	assert := assert.New(t)
	assert.NoError(runCRUDCustomApplications())
	assert.NoError(runGetAllCustomApplications())
}

func runCRUDCustomApplications() error {
	client, err := NewClient()
	if err != nil {
		return err
	}

	fmt.Println("### CREATE")
	app := models.NewCustomApplication("apitest-customapp-1")
	app.Description = pointer.ToString("Testing custom application api")
	app.IPRange = pointer.ToString("192.168.0.1,192.168.0.2")
	app.Protocol = pointer.ToString("6,17")
	app.Port = pointer.ToString("9001,9002,9003")
	app.ASN = pointer.ToString("asn1,asn2,asn3")
	created, err := client.CustomApplications.Create(context.Background(), *app)
	if err != nil {
		return err
	}
	PrettyPrint(created)
	fmt.Println()

	fmt.Println("### UPDATE")
	created.Name = "apitest-customapp-ONE"
	created.Description = pointer.ToString("Updated description")
	created.Port = pointer.ToString("1023")
	updated, err := client.CustomApplications.Update(context.Background(), *created)
	if err != nil {
		return err
	}
	PrettyPrint(updated)
	fmt.Println()

	fmt.Println("### DELETE")
	err = client.CustomApplications.Delete(context.Background(), created.ID)
	if err != nil {
		return err
	}
	fmt.Println("Success")
	fmt.Println()

	return nil
}

func runGetAllCustomApplications() error {
	client, err := NewClient()
	if err != nil {
		return err
	}

	fmt.Println("### GET ALL")
	applications, err := client.CustomApplications.GetAll(context.Background())
	if err != nil {
		return err
	}
	PrettyPrint(applications)
	fmt.Println()

	return nil
}
