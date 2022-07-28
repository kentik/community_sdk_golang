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

func TestDemonstrateTagsCRUD(t *testing.T) {
	t.Parallel()
	err := demonstrateTagsCRUD()
	assert.NoError(t, err)
}

func TestDemonstrateTagsGetAll(t *testing.T) {
	t.Parallel()
	err := demonstrateTagsGetAll()
	assert.NoError(t, err)
}

func demonstrateTagsCRUD() error {
	ctx := context.Background()
	client, err := NewClient()
	if err != nil {
		return err
	}

	fmt.Println("### client.Tags.Create")
	t := models.NewTag("APITEST-TAG-1")
	t.DeviceName = stringPointer("device1,192.168.5.100")
	t.DeviceType = stringPointer("router,switch")
	t.Site = stringPointer("site1,site2")
	t.InterfaceName = stringPointer("interface1,interface2")
	t.Addr = stringPointer("192.168.0.1,192.168.0.2")
	t.Port = stringPointer("9000,9001")
	t.TCPFlags = stringPointer("7")
	t.Protocol = stringPointer("6,17")
	t.ASN = stringPointer("101,102,103")
	t.LastHopAsName = stringPointer("as1,as2,as3")
	t.NextHopASN = stringPointer("51,52,53")
	t.NextHopAsName = stringPointer("as51,as52,as53")
	t.NextHop = stringPointer("192.168.7.1,192.168.7.2")
	t.BGPAsPath = stringPointer("201,202,203")
	t.BGPCommunity = stringPointer("301,302,303")
	t.MAC = stringPointer("FF:FF:FF:FF:FF:FE,FF:FF:FF:FF:FF:FF")
	t.Country = stringPointer("ES,IT")
	t.VLANs = stringPointer("4001,4002,4003")

	tag, err := client.Tags.Create(ctx, *t)
	if err != nil {
		return fmt.Errorf("client.Tags.Create failed: %w", err)
	}

	PrettyPrint(tag)
	fmt.Println()

	fmt.Println("### client.Tags.Get")
	tag, err = client.Tags.Get(ctx, tag.ID)
	if err != nil {
		return fmt.Errorf("client.Tags.Get failed: %w", err)
	}

	PrettyPrint(tag)
	fmt.Println()

	fmt.Println("### client.Tags.Update")
	tag.FlowTag = "APITEST-TAG-ONE"
	t.DeviceType = pointer.ToString("nat")
	t.Country = pointer.ToString("GR")

	tag, err = client.Tags.Update(ctx, *tag)
	if err != nil {
		return fmt.Errorf("client.Tags.Update failed: %w", err)
	}

	PrettyPrint(tag)
	fmt.Println()

	fmt.Println("### client.Tags.Delete")
	err = client.Tags.Delete(context.Background(), tag.ID)
	if err != nil {
		return fmt.Errorf("client.Tags.Delete failed: %w", err)
	}

	return nil
}

func demonstrateTagsGetAll() error {
	ctx := context.Background()
	client, err := NewClient()
	if err != nil {
		return err
	}

	fmt.Println("### client.Tags.GetAll")
	tags, err := client.Tags.GetAll(ctx)
	if err != nil {
		return fmt.Errorf("client.Tags.GetAll failed: %w", err)
	}
	PrettyPrint(tags)

	return nil
}
