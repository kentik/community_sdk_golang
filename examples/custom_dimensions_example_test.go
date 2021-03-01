//+build examples

package examples

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
)

func TestCustomDimensionsAPIExample(t *testing.T) {
	assert := assert.New(t)
	assert.NoError(runGetAll())
	assert.NoError(runCRUD())
}

func runGetAll() error {
	client := NewClient()

	fmt.Println("### GET ALL")
	dimensions, err := client.CustomDimensions.GetAll(context.Background())
	if err != nil {
		return err
	}
	PrettyPrint(dimensions)
	fmt.Println()

	return nil
}

func runCRUD() error {
	var err error
	client := NewClient()

	fmt.Println("### CREATE")
	name := "c_testapi_dim_" + randID() // random id as even deleted names are held for 1 year and must be unique
	dimension := models.NewCustomDimension(name, "test_dimension", models.CustomDimensionTypeStr)
	created, err := client.CustomDimensions.Create(context.Background(), *dimension)
	if err != nil {
		return err
	}
	PrettyPrint(created)
	fmt.Println()

	fmt.Println("### UPDATE")
	created.DisplayName = "test_dimension_updated"
	updated, err := client.CustomDimensions.Update(context.Background(), *created)
	if err != nil {
		return err
	}
	PrettyPrint(updated)
	fmt.Println()

	fmt.Println("### GET")
	got, err := client.CustomDimensions.Get(context.Background(), created.ID)
	if err != nil {
		return err
	}
	PrettyPrint(got)
	fmt.Println()

	fmt.Println("### DELETE")
	err = client.CustomDimensions.Delete(context.Background(), created.ID)
	if err != nil {
		return err
	}

	fmt.Println("Success")
	fmt.Println()

	return nil
}

func randID() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	length := 5
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String() // E.g. "ExcbsVQs"
}
