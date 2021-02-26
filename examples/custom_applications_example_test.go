//+build examples

package examples

import (
	"context"
	"fmt"
	"testing"

	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
)

func TestCustomApplicationsAPIExample(t *testing.T) {
	assert := assert.New(t)
	assert.NoError(runCRUD())
	assert.NoError(runGetAll())
}

func runCRUD() error {
	var err error
	client := NewClient()

	fmt.Println("### CREATE")
	app := models.NewCustomApplication("apitest-customapp-1")
	models.SetOptional(&app.Description, "Testing custom application api")
	models.SetOptional(&app.IPRange, "192.168.0.1,192.168.0.2")
	models.SetOptional(&app.Protocol, "6,17")
	models.SetOptional(&app.Port, "9001,9002,9003")
	models.SetOptional(&app.ASN, "asn1,asn2,asn3")
	created, err := client.CustomApplications.Create(context.Background(), *app)
	if err != nil {
		return err
	}
	PrettyPrint(created)
	fmt.Println()

	fmt.Println("### UPDATE")
	created.Name = "apitest-customapp-ONE"
	models.SetOptional(&created.Description, "Updated description")
	models.SetOptional(&created.Port, "1023")
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

func runGetAll() error {
	client := NewClient()

	fmt.Println("### GET ALL")
	applications, err := client.CustomApplications.GetAll(context.Background())
	if err != nil {
		return err
	}
	PrettyPrint(applications)
	fmt.Println()

	return nil
}
