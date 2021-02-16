//+build examples

package examples

import (
	"context"
	"fmt"
	"testing"

	"github.com/kentik/community_sdk_golang/kentikapi"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

func TestDevicesAPIExample(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Fatal(err)
		}
	}()
	runCRUD()
	runGetAll()
}

func runGetAll() {
	fmt.Println("### GET ALL")

	email, token, err := readCredentialsFromEnv()
	panicOnError(err)

	client := kentikapi.NewClient(kentikapi.Config{
		AuthEmail: email,
		AuthToken: token,
	})

	devices, err := client.DevicesAPI.GetAll(context.Background())
	panicOnError(err)

	prettyPrint(devices)
}

func runCRUD() {
	email, token, err := readCredentialsFromEnv()
	panicOnError(err)

	client := kentikapi.NewClient(kentikapi.Config{
		AuthEmail: email,
		AuthToken: token,
	})

	fmt.Println("### GET")
	device, err := client.DevicesAPI.Get(context.Background(), models.ID(79685))
	panicOnError(err)

	prettyPrint(device)
}
