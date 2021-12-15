//go:build examples
// +build examples

//nolint:testpackage,forbidigo
package examples

import (
	"context"
	"fmt"
	"testing"

	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
)

func TestDevicesAPIExample(t *testing.T) {
	assert := assert.New(t)
	assert.NoError(runCRUDRouter())
	assert.NoError(runCRUDDNS())
	assert.NoError(runGetAllDevices())
	assert.NoError(runGetAllInterfaces())
}

func runGetAllDevices() error {
	client, err := NewClient()
	if err != nil {
		return err
	}

	fmt.Println("### GET ALL")
	devices, err := client.Devices.GetAll(context.Background())
	if err != nil {
		return err
	}
	PrettyPrint(devices)
	fmt.Println()

	return nil
}

func runCRUDRouter() error {
	client, err := NewClient()
	if err != nil {
		return err
	}

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
	if err != nil {
		return err
	}
	PrettyPrint(createdDevice)
	fmt.Println()

	fmt.Println("### UPDATE ROUTER")
	createdDevice.SendingIPS = []string{"128.0.0.15", "128.0.0.16"}
	createdDevice.DeviceSampleRate = 10
	models.SetOptional(&createdDevice.DeviceDescription, "updated description")
	models.SetOptional(&createdDevice.DeviceBGPNeighborASN, "88")
	updatedDevice, err := client.Devices.Update(context.Background(), *createdDevice)
	if err != nil {
		return err
	}
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
	if err != nil {
		return err
	}
	PrettyPrint(createdInterface)
	fmt.Println()

	fmt.Println("### UPDATE INTERFACE")
	createdInterface.SNMPSpeed = 24
	updatedInterface, err := client.Devices.Interfaces.Update(context.Background(), *createdInterface)
	if err != nil {
		return err
	}
	PrettyPrint(updatedInterface)
	fmt.Println()

	fmt.Println("### GET ROUTER")
	gotDevice, err := client.Devices.Get(context.Background(), createdDevice.ID)
	if err != nil {
		return err
	}
	PrettyPrint(gotDevice)
	fmt.Println()

	fmt.Println("### DELETE INTERFACE")
	err = client.Devices.Interfaces.Delete(context.Background(), createdInterface.DeviceID, createdInterface.ID)
	if err != nil {
		return err
	}
	fmt.Println("Success")
	fmt.Println()

	fmt.Println("### DELETE ROUTER")
	err = client.Devices.Delete(context.Background(), createdDevice.ID) // archive
	if err != nil {
		return err
	}
	err = client.Devices.Delete(context.Background(), createdDevice.ID) // delete
	if err != nil {
		return err
	}
	fmt.Println("Success")
	fmt.Println()

	return nil
}

func runCRUDDNS() error {
	client, err := NewClient()
	if err != nil {
		return err
	}

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
	if err != nil {
		return err
	}
	PrettyPrint(createdDevice)
	fmt.Println()

	fmt.Println("### UPDATE")
	createdDevice.DeviceSampleRate = 10
	models.SetOptional(&createdDevice.CDNAttr, models.CDNAttributeNo)
	models.SetOptional(&createdDevice.DeviceDescription, "updated description")
	models.SetOptional(&createdDevice.DeviceBGPFlowSpec, false)
	updatedDevice, err := client.Devices.Update(context.Background(), *createdDevice)
	if err != nil {
		return err
	}
	PrettyPrint(updatedDevice)
	fmt.Println()

	// first make sure the label ids exist!
	// fmt.Println("### APPLY LABELS")
	// labelIDs := []models.ID{models.ID(3011), models.ID( 3012)}
	// labels, err := client.Devices.ApplyLabels(context.Background(),createdDevice.ID, labelIDs)
	// if err != nil {
	// 	return err
	// }
	// PrettyPrint(labels)
	// fmt.Println()

	fmt.Println("### GET")
	gotDevice, err := client.Devices.Get(context.Background(), createdDevice.ID)
	if err != nil {
		return err
	}
	PrettyPrint(gotDevice)
	fmt.Println()

	fmt.Println("### DELETE")
	err = client.Devices.Delete(context.Background(), createdDevice.ID) // archive
	if err != nil {
		return err
	}
	err = client.Devices.Delete(context.Background(), createdDevice.ID) // delete
	if err != nil {
		return err
	}
	fmt.Println("Success")
	fmt.Println()

	return nil
}

func runGetAllInterfaces() error {
	client, err := NewClient()
	if err != nil {
		return err
	}

	fmt.Println("### GET ALL INTERFACES")
	interfaces, err := client.Devices.Interfaces.GetAll(context.Background(), models.ID(80166))
	if err != nil {
		return err
	}
	PrettyPrint(interfaces)
	fmt.Println()

	return nil
}
