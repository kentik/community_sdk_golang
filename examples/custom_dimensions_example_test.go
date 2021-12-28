//go:build examples
// +build examples

//nolint:testpackage,forbidigo
package examples

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
)

func TestCustomDimensionsAPIExample(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	assert.NoError(runGetAllCustomDimensions())
	assert.NoError(runCRUDCustomDimensions())
}

func runGetAllCustomDimensions() error {
	client, err := NewClient()
	if err != nil {
		return err
	}

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
	client, err := NewClient()
	if err != nil {
		return err
	}

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
	populator.InterfaceName = pointer.ToString("interface1,interface2")
	populator.Addr = pointer.ToString("128.0.0.1/32,128.0.0.2/32")
	populator.Port = pointer.ToString("1001,1002")
	populator.TCPFlags = pointer.ToString("160")
	populator.Protocol = pointer.ToString("6,17")
	populator.ASN = pointer.ToString("101,102")
	populator.NextHopASN = pointer.ToString("201,202")
	populator.NextHop = pointer.ToString("128.0.200.1/32,128.0.200.2/32")
	populator.BGPAsPath = pointer.ToString("3001,3002")
	populator.BGPCommunity = pointer.ToString("401:499,501:599")
	populator.DeviceType = pointer.ToString("device-type1")
	populator.Site = pointer.ToString("site1,site2,site3")
	populator.LastHopAsName = pointer.ToString("asn101,asn102")
	populator.NextHopAsName = pointer.ToString("asn201,asn202")
	populator.MAC = pointer.ToString("FF:FF:FF:FF:FF:FA,FF:FF:FF:FF:FF:FF")
	populator.Country = pointer.ToString("NL,SE")
	populator.VLans = pointer.ToString("2001,2002")
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

//nolint: gosec // no need for cryptographically secure RNG here
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
