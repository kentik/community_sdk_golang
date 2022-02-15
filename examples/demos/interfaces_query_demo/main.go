//nolint:forbidigo
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/AlekSi/pointer"
	"github.com/kentik/community_sdk_golang/examples/demos"
	"github.com/kentik/community_sdk_golang/kentikapi"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

func main() {
	demos.Step("Create Kentik API client")
	client, err := demos.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	// Create device and read ID
	demos.Step("Create a device")
	deviceID := createDevice(client)

	// Interfaces
	demos.Step("Create an interface")
	id := createInterface(client, deviceID)

	demos.Step("Get interface")
	getInterface(client, deviceID, id)

	demos.Step("Delete interface")
	deleteInterface(client, deviceID, id)

	demos.Step("Get all interfaces")
	getAllInterfaces(client, deviceID)

	// Delete Device
	demos.Step("Delete device")
	deleteDevice(client, deviceID)

	// Query
	demos.Step("Query for data")
	queryData(client)

	demos.Step("Query for chart")
	queryChart(client)
}

func createDevice(client *kentikapi.Client) models.ID {
	device := models.NewDeviceDNS(
		"interfaces_query_demo_device",
		models.DeviceSubtypeAwsSubnet,
		1,
		"11466",
		models.CDNAttributeYes,
	)
	createdDevice, err := client.Devices.Create(context.Background(), *device)
	demos.ExitOnError(err)
	fmt.Printf("Successfully created device, ID = %v\n", createdDevice.ID)

	return createdDevice.ID
}

//nolint:gomnd
func createInterface(client *kentikapi.Client, deviceID models.ID) models.ID {
	intf := models.NewInterface(
		deviceID,
		models.ID("2"),
		15,
		"testapi-interface-demo",
	)
	createdInterface, err := client.Devices.Interfaces.Create(context.Background(), *intf)
	demos.ExitOnError(err)
	fmt.Printf("Successfully created interface, ID = %v\n", createdInterface.ID)

	return createdInterface.ID
}

func getInterface(client *kentikapi.Client, deviceID, interfaceID models.ID) {
	fmt.Printf("Retrieving interface, deviceID = %v, interfaceID = %v\n", deviceID, interfaceID)
	i, err := client.Devices.Interfaces.Get(context.Background(), deviceID, interfaceID)
	demos.ExitOnError(err)
	demos.PrettyPrint(i)
}

func deleteInterface(client *kentikapi.Client, deviceID, interfaceID models.ID) {
	fmt.Printf("Deleting interface, deviceID = %v, interfaceID = %v\n", deviceID, interfaceID)
	err := client.Devices.Interfaces.Delete(context.Background(), deviceID, interfaceID)
	demos.ExitOnError(err)
	fmt.Println("Successful")
}

func getAllInterfaces(client *kentikapi.Client, deviceID models.ID) {
	fmt.Printf("Listing interfaces for deviceID = %v\n", deviceID)
	interfaces, err := client.Devices.Interfaces.GetAll(context.Background(), deviceID)
	demos.ExitOnError(err)
	demos.PrettyPrint(interfaces)
	fmt.Printf("Total interfaces: %d\n", len(interfaces))
}

func deleteDevice(client *kentikapi.Client, deviceID models.ID) {
	fmt.Printf("Deleting device, deviceID = %v\n", deviceID)
	err := client.Devices.Delete(context.Background(), deviceID)
	demos.ExitOnError(err)
	fmt.Println("Successful")
}

//nolint:gomnd
func queryData(client *kentikapi.Client) {
	// prepare query
	agg1 := models.NewAggregate("avg_bits_per_sec", "f_sum_both_bytes", models.AggregateFunctionTypeAverage)
	agg1.Raw = pointer.ToBool(true)
	agg2 := models.NewAggregate("p95th_bits_per_sec", "f_sum_both_bytes", models.AggregateFunctionTypePercentile)
	agg2.Rank = pointer.ToInt(95)
	agg3 := models.NewAggregate("max_bits_per_sec", "f_sum_both_bytes", models.AggregateFunctionTypeMax)
	query := models.NewQuery(
		models.MetricTypeBytes,
		[]models.DimensionType{models.DimensionTypeTraffic},
	)
	query.Depth = 75
	query.LookbackSeconds = 60 * 30 // last 30 minutes
	query.HostnameLookup = true
	query.TopX = 8
	query.Depth = 75
	query.Aggregates = []models.Aggregate{agg1, agg2, agg3}
	query.CIDR = pointer.ToInt(32)
	query.CIDR6 = pointer.ToInt(128)
	query.Outsort = pointer.ToString("avg_bits_per_sec")
	query.AllSelected = pointer.ToBool(true)
	queryItem := models.QueryArrayItem{Query: *query, Bucket: "Left +Y Axis"}
	queryObject := models.QueryObject{Queries: []models.QueryArrayItem{queryItem}}

	// send query
	fmt.Println("Sending query...")
	result, err := client.Query.Data(context.Background(), queryObject)
	demos.ExitOnError(err)
	fmt.Println("Done.")

	// display result
	err = demos.DisplayQueryDataResult(result)
	demos.ExitOnError(err)
}

//nolint:gomnd
func queryChart(client *kentikapi.Client) {
	// prepare query
	agg1 := models.NewAggregate("avg_bits_per_sec", "f_sum_both_bytes", models.AggregateFunctionTypeAverage)
	agg1.Raw = pointer.ToBool(true)
	agg2 := models.NewAggregate("p95th_bits_per_sec", "f_sum_both_bytes", models.AggregateFunctionTypePercentile)
	agg2.Rank = pointer.ToInt(95)
	agg3 := models.NewAggregate("max_bits_per_sec", "f_sum_both_bytes", models.AggregateFunctionTypeMax)
	query := models.NewQuery(
		models.MetricTypeBytes,
		[]models.DimensionType{models.DimensionTypeTraffic},
	)
	query.Aggregates = []models.Aggregate{agg1, agg2, agg3}
	query.LookbackSeconds = 60 * 60 // last 60 minutes
	query.QueryTitle = "avg_bits_per_sec for last 60 minutes"
	query.TopX = 8
	query.Depth = 75
	query.FastData = models.FastDataTypeAuto
	query.TimeFormat = models.TimeFormatLocal
	query.HostnameLookup = true
	query.CIDR = pointer.ToInt(32)
	query.CIDR6 = pointer.ToInt(128)
	query.Outsort = pointer.ToString("avg_bits_per_sec")
	query.AllSelected = pointer.ToBool(true)
	query.VizType = models.ChartViewTypePtr(models.ChartViewTypeStackedArea)
	query.ShowOverlay = pointer.ToBool(false)
	query.OverlayDay = pointer.ToInt(-7)
	query.SyncAxes = pointer.ToBool(false)
	query.PPSThreshold = pointer.ToInt(1)
	queryItem := models.QueryArrayItem{Query: *query, Bucket: "Left +Y Axis"}
	queryItem.IsOverlay = pointer.ToBool(false)
	queryObject := models.QueryObject{Queries: []models.QueryArrayItem{queryItem}}
	queryObject.ImageType = models.ImageTypePtr(models.ImageTypePNG)

	// send query
	fmt.Println("Sending query...")
	result, err := client.Query.Chart(context.Background(), queryObject)
	demos.ExitOnError(err)
	fmt.Println("Done.")

	// display result
	err = demos.DisplayQueryChartResult(result)
	demos.ExitOnError(err)
}
