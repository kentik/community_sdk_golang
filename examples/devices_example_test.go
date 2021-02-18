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
	runCRUDRouter()
	runCRUDDNS()
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

func runCRUDRouter() {
	var err error
	email, token, err := readCredentialsFromEnv()
	panicOnError(err)

	client := kentikapi.NewClient(kentikapi.Config{
		AuthEmail: email,
		AuthToken: token,
	})

	fmt.Println("### CREATE ROUTER")
	snmpv3conf := models.NewSNMPv3Conf("John")
	snmpv3conf = snmpv3conf.WithAuthentication(models.AuthenticationProtocolMD5, "Auth_Pass")
	snmpv3conf = snmpv3conf.WithPrivacy(models.PrivacyProtocolDES, "Priv_Pass")
	device := models.NewDeviceRouter(
		"testapi_router_router_full",
		models.DeviceSubtypeRouter,
		1,
		models.ID(11466),
		[]string{"128.0.0.10"},
		false,
	).WithBGPTypeDevice("77")
	models.SetOptional(&device.DeviceDescription, "testapi device type router subrype router with full config")
	models.SetOptional(&device.DeviceSNMNPIP, "127.0.0.1")
	models.SetOptional(&device.DeviceSNMPv3Conf, *snmpv3conf)
	models.SetOptional(&device.DeviceBGPNeighborIP, "127.0.0.2")
	models.SetOptional(&device.DeviceBGPPassword, "bgp-optional-password")
	models.SetOptional(&device.SiteID, 8483)
	models.SetOptional(&device.DeviceBGPFlowSpec, true)
	models.SetOptional(&device.DeviceBGPNeighborIP, "127.0.0.42")
	models.SetOptional(&device.DeviceBGPPassword, "bgp-optional-password")

	createdDevice, err := client.DevicesAPI.Create(context.Background(), *device)
	panicOnError(err)
	prettyPrint(createdDevice)

	fmt.Println()
	fmt.Println("### GET")
	gotDevice, err := client.DevicesAPI.Get(context.Background(), createdDevice.ID)
	panicOnError(err)
	prettyPrint(gotDevice)

	// fmt.Println()
	// fmt.Println("### UPDATE")

	// fmt.Println()
	// fmt.Println("### DELETE")
}

func runCRUDDNS() {
	var err error
	email, token, err := readCredentialsFromEnv()
	panicOnError(err)

	client := kentikapi.NewClient(kentikapi.Config{
		AuthEmail: email,
		AuthToken: token,
	})

	fmt.Println("### CREATE DNS")
	device := models.NewDeviceDNS(
		"testapi_dns_awssubnet",
		models.DeviceSubtypeAwsSubnet,
		1,
		models.ID(11466),
		models.CDNAttributeYes,
	)
	models.SetOptional(&device.SiteID, 8483)
	models.SetOptional(&device.DeviceBGPFlowSpec, true)

	createdDevice, err := client.DevicesAPI.Create(context.Background(), *device)
	panicOnError(err)
	prettyPrint(createdDevice)

	fmt.Println()
	fmt.Println("### GET")
	gotDevice, err := client.DevicesAPI.Get(context.Background(), createdDevice.ID)
	panicOnError(err)
	prettyPrint(gotDevice)

	// fmt.Println()
	// fmt.Println("### UPDATE")

	// fmt.Println()
	// fmt.Println("### DELETE")
}
