//go:build examples
// +build examples

package examples

import (
	"context"
	"fmt"
	"testing"

	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
)

func TestSitesAPIExample(t *testing.T) {
	assert.NoError(t, runCRUD())
	assert.NoError(t, runGetAll())
}

func runCRUD() error {
	var err error
	client := NewClient()

	fmt.Println("### CREATE")
	site := models.NewSite("apitest-site-1")
	models.SetOptional(&site.Longitude, 12.0)
	created, err := client.Sites.Create(context.Background(), *site)
	if err != nil {
		return err
	}
	PrettyPrint(created)
	fmt.Println()

	fmt.Println("### UPDATE")
	created.SiteName = "apitest-site-one"
	models.SetOptional(&created.Latitude, 49.0)
	updated, err := client.Sites.Update(context.Background(), *created)
	if err != nil {
		return err
	}
	PrettyPrint(updated)
	fmt.Println()

	fmt.Println("### GET")
	gotSite, err := client.Sites.Get(context.Background(), created.ID)
	if err != nil {
		return err
	}
	PrettyPrint(gotSite)
	fmt.Println()

	fmt.Println("### DELETE")
	err = client.Sites.Delete(context.Background(), created.ID)
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
	sites, err := client.Sites.GetAll(context.Background())
	if err != nil {
		return err
	}
	PrettyPrint(sites)
	fmt.Println()

	return nil
}
