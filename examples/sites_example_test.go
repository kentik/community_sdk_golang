//+build examples

package examples

import (
	"context"
	"fmt"
	"runtime/debug"
	"testing"

	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

func TestSitesAPIExample(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			t.Fatal(err)
		}
	}()

	runCRUD()
	runGetAll()
}

func runCRUD() {
	var err error
	client := NewClient()

	fmt.Println("### CREATE")
	site := models.NewSite("apitest-site-1")
	models.SetOptional(&site.Longitude, 12.0)
	created, err := client.Sites.Create(context.Background(), *site)
	PanicOnError(err)
	PrettyPrint(created)
	fmt.Println()

	fmt.Println("### UPDATE")
	created.SiteName = "apitest-site-one"
	models.SetOptional(&created.Latitude, 49.0)
	updated, err := client.Sites.Update(context.Background(), *created)
	PanicOnError(err)
	PrettyPrint(updated)
	fmt.Println()

	fmt.Println("### GET")
	gotSite, err := client.Sites.Get(context.Background(), created.ID)
	PanicOnError(err)
	PrettyPrint(gotSite)
	fmt.Println()

	fmt.Println("### DELETE")
	err = client.Sites.Delete(context.Background(), created.ID)
	PanicOnError(err)
	fmt.Println("Success")
	fmt.Println()
}

func runGetAll() {
	client := NewClient()

	fmt.Println("### GET ALL")
	sites, err := client.Sites.GetAll(context.Background())
	PanicOnError(err)
	PrettyPrint(sites)
	fmt.Println()
}
