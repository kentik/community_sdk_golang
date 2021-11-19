//go:build examples
// +build examples

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
	assert.NoError(runGetAllCustomDimensions())
	assert.NoError(runCRUDCustomDimensions())
}

func runGetAllCustomDimensions() error {
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

func runCRUDCustomDimensions() error {
	var err error
	client := NewClient()

	fmt.Println("### CREATE DIMENSION")
	name := "c_testapi_dim_" + randID() // random id as even deleted names are held for 1 year and must be unique
	dimension := models.NewCustomDimension(name, "test_dimension", models.CustomDimensionTypeStr)
	created, err := client.CustomDimensions.Create(context.Background(), *dimension)
	if err != nil {
		return err
	}
	PrettyPrint(created)
	fmt.Println()

	fmt.Println("### UPDATE DIMENSION")
	created.DisplayName = "test_dimension_updated"
	updated, err := client.CustomDimensions.Update(context.Background(), *created)
	if err != nil {
		return err
	}
	PrettyPrint(updated)
	fmt.Println()

	fmt.Println("### CREATE POPULATOR")
	populator := models.NewPopulator(created.ID, "testapi-dimension-value-1", "device1,128.0.0.100", models.PopulatorDirectionDst)
	models.SetOptional(&populator.InterfaceName, "interface1,interface2")
	models.SetOptional(&populator.Addr, "128.0.0.1/32,128.0.0.2/32")
	models.SetOptional(&populator.Port, "1001,1002")
	models.SetOptional(&populator.TCPFlags, "160")
	models.SetOptional(&populator.Protocol, "6,17")
	models.SetOptional(&populator.ASN, "101,102")
	models.SetOptional(&populator.NextHopASN, "201,202")
	models.SetOptional(&populator.NextHop, "128.0.200.1/32,128.0.200.2/32")
	models.SetOptional(&populator.BGPAsPath, "3001,3002")
	models.SetOptional(&populator.BGPCommunity, "401:499,501:599")
	models.SetOptional(&populator.DeviceType, "device-type1")
	models.SetOptional(&populator.Site, "site1,site2,site3")
	models.SetOptional(&populator.LastHopAsName, "asn101,asn102")
	models.SetOptional(&populator.NextHopAsName, "asn201,asn202")
	models.SetOptional(&populator.MAC, "FF:FF:FF:FF:FF:FA,FF:FF:FF:FF:FF:FF")
	models.SetOptional(&populator.Country, "NL,SE")
	models.SetOptional(&populator.VLans, "2001,2002")
	createdPopulator, err := client.CustomDimensions.Populators.Create(context.Background(), *populator)
	if err != nil {
		return err
	}
	PrettyPrint(createdPopulator)
	fmt.Println()

	fmt.Println("### UPDATE POPULATOR")
	createdPopulator.Value = "testapi-dimension-value-updated"
	createdPopulator.Direction = models.PopulatorDirectionEither
	updatedPopulator, err := client.CustomDimensions.Populators.Update(context.Background(), *createdPopulator)
	if err != nil {
		return err
	}
	PrettyPrint(updatedPopulator)
	fmt.Println()

	fmt.Println("### GET DIMENSION")
	got, err := client.CustomDimensions.Get(context.Background(), created.ID)
	if err != nil {
		return err
	}
	PrettyPrint(got)
	fmt.Println()

	fmt.Println("### DELETE POPULATOR")
	err = client.CustomDimensions.Populators.Delete(context.Background(), createdPopulator.DimensionID, createdPopulator.ID)
	if err != nil {
		return err
	}
	fmt.Println("Success")
	fmt.Println()

	fmt.Println("### DELETE DIMENSION")
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
