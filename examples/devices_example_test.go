//+build examples

package examples

import (
	"context"
	"fmt"
	"runtime/debug"
	"testing"

	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

func TestDevicesAPIExample(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			t.Fatal(err)
		}
	}()

	runCRUDRouter()
	runCRUDDNS()
	runGetAllDevices()
	runGetInterface()
	runGetAllInterfaces()
}

func runGetAllDevices() {
	client := NewClient()

	fmt.Println("### GET ALL")
	devices, err := client.Devices.GetAll(context.Background())
	PanicOnError(err)
	PrettyPrint(devices)
	fmt.Println()
}

func runCRUDRouter() {
	var err error
	client := NewClient()

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
	createdDevice, err := client.Devices.Create(context.Background(), *device)
	PanicOnError(err)
	PrettyPrint(createdDevice)
	fmt.Println()

	fmt.Println("### UPDATE ROUTER")
	createdDevice.SendingIPS = []string{"128.0.0.15", "128.0.0.16"}
	createdDevice.DeviceSampleRate = 10
	models.SetOptional(&createdDevice.DeviceDescription, "updated description")
	models.SetOptional(&createdDevice.DeviceBGPNeighborASN, "88")
	updatedDevice, err := client.Devices.Update(context.Background(), *createdDevice)
	PanicOnError(err)
	PrettyPrint(updatedDevice)
	fmt.Println()

	fmt.Println("### CREATE INTERFACE")
	intf := models.NewInterface(
		createdDevice.ID,
		models.ID(2),
		15,
		"testapi-interface",
	)
	createdInterface, err := client.Devices.Interfaces.Create(context.Background(), *intf)
	PanicOnError(err)
	PrettyPrint(createdInterface)
	fmt.Println()

	fmt.Println("### GET ROUTER")
	gotDevice, err := client.Devices.Get(context.Background(), createdDevice.ID)
	PanicOnError(err)
	PrettyPrint(gotDevice)
	fmt.Println()

	fmt.Println("### DELETE INTERFACE")
	err = client.Devices.Interfaces.Delete(context.Background(), createdInterface.DeviceID, createdInterface.ID)
	fmt.Println("Success")
	fmt.Println()

	fmt.Println("### DELETE ROUTER")
	err = client.Devices.Delete(context.Background(), createdDevice.ID) // archive
	PanicOnError(err)
	err = client.Devices.Delete(context.Background(), createdDevice.ID) // delete
	PanicOnError(err)
	fmt.Println("Success")
	fmt.Println()
}

func runCRUDDNS() {
	var err error
	client := NewClient()

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

	createdDevice, err := client.Devices.Create(context.Background(), *device)
	PanicOnError(err)
	PrettyPrint(createdDevice)
	fmt.Println()

	fmt.Println("### UPDATE")
	createdDevice.DeviceSampleRate = 10
	models.SetOptional(&createdDevice.CDNAttr, models.CDNAttributeNo)
	models.SetOptional(&createdDevice.DeviceDescription, "updated description")
	models.SetOptional(&createdDevice.DeviceBGPFlowSpec, false)
	updatedDevice, err := client.Devices.Update(context.Background(), *createdDevice)
	PanicOnError(err)
	PrettyPrint(updatedDevice)
	fmt.Println()

	// first make sure the label ids exist!
	// fmt.Println("### APPLY LABELS")
	// labelIDs := []models.ID{models.ID(3011), models.ID( 3012)}
	// labels, err := client.Devices.ApplyLabels(context.Background(),createdDevice.ID, labelIDs)
	// PanicOnError(err)
	// PrettyPrint(labels)
	// fmt.Println()

	fmt.Println("### GET")
	gotDevice, err := client.Devices.Get(context.Background(), createdDevice.ID)
	PanicOnError(err)
	PrettyPrint(gotDevice)
	fmt.Println()

	fmt.Println("### DELETE")
	err = client.Devices.Delete(context.Background(), createdDevice.ID) // archive
	PanicOnError(err)
	err = client.Devices.Delete(context.Background(), createdDevice.ID) // delete
	PanicOnError(err)
	fmt.Println("Success")
	fmt.Println()
}

func runGetInterface() {
	client := NewClient()

	fmt.Println("### GET INTERFACE")
	deviceID := models.ID(80166)
	interfaceID := models.ID(9385804334)
	intf, err := client.Devices.Interfaces.Get(context.Background(), deviceID, interfaceID)
	PanicOnError(err)
	PrettyPrint(intf)
	fmt.Println()
}

func runGetAllInterfaces() {
	client := NewClient()

	fmt.Println("### GET ALL INTERFACES")
	interfaces, err := client.Devices.Interfaces.GetAll(context.Background(), models.ID(80166))
	PanicOnError(err)
	PrettyPrint(interfaces)
	fmt.Println()
}
