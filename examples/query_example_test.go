//go:build examples
// +build examples

//nolint:testpackage,forbidigo
package examples

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
)

func TestQueryAPIExample(t *testing.T) {
	assert := assert.New(t)
	assert.NoError(runQuerySQL())
	assert.NoError(runQueryData())
	assert.NoError(runQueryURL())
	assert.NoError(runQueryChart())
}

func runQuerySQL() error {
	client, err := NewClient()
	if err != nil {
		return err
	}

	fmt.Println("### QUERY SQL")
	// Return kpps and kBps over the last 10 minutes,
	// grouped by minute (the first minute is skipped
	// as it is likely incomplete most of the time)
	sql := `
        SELECT i_start_time, 
        round(sum(in_pkts)/(3600)/1000) AS f_sum_in_pkts, 
        round(sum(in_bytes)/(3600)/1000)*8 AS f_sum_in_bytes 
        FROM all_devices 
        WHERE ctimestamp > 600 
        AND ctimestamp < 60
        GROUP by i_start_time 
        ORDER by i_start_time DESC 
        LIMIT 1000;
    `
	result, err := client.Query.SQL(context.Background(), sql)
	if err != nil {
		return err
	}
	PrettyPrint(result)
	fmt.Println()

	return nil
}

func runQueryData() error {
	client, err := NewClient()
	if err != nil {
		return err
	}

	fmt.Println("### QUERY Data")
	result, err := client.Query.Data(context.Background(), makeQueryObject())
	if err != nil {
		return err
	}
	PrettyPrint(result)
	fmt.Println()

	return nil
}

//nolint:gosec // G204: Subprocess launched with variable - not an issue in this case
func runQueryChart() error {
	client, err := NewClient()
	if err != nil {
		return err
	}

	fmt.Println("### QUERY Chart")
	agg1 := models.NewAggregate("avg_bits_per_sec", "f_sum_both_bytes", models.AggregateFunctionTypeAverage)
	models.SetOptional(&agg1.Raw, true)
	agg2 := models.NewAggregate("p95th_bits_per_sec", "f_sum_both_bytes", models.AggregateFunctionTypePercentile)
	models.SetOptional(&agg2.Rank, 95)
	agg3 := models.NewAggregate("max_bits_per_sec", "f_sum_both_bytes", models.AggregateFunctionTypeMax)

	query := models.NewQuery(
		models.MetricTypeBytes,
		[]models.DimensionType{models.DimensionTypeTraffic},
	)
	query.Aggregates = []models.Aggregate{agg1, agg2, agg3}
	query.LookbackSeconds = 3600
	query.QueryTitle = "Example query"
	query.TopX = 8
	query.Depth = 75
	query.FastData = models.FastDataTypeAuto
	query.TimeFormat = models.TimeFormatLocal
	query.HostnameLookup = true
	models.SetOptional(&query.CIDR, 32)
	models.SetOptional(&query.CIDR6, 128)
	models.SetOptional(&query.Outsort, "avg_bits_per_sec")
	models.SetOptional(&query.AllSelected, true)
	models.SetOptional(&query.VizType, models.ChartViewTypeStackedArea)
	models.SetOptional(&query.ShowOverlay, false)
	models.SetOptional(&query.OverlayDay, -7)
	models.SetOptional(&query.SyncAxes, false)
	models.SetOptional(&query.PPSThreshold, 1)

	queryItem := models.QueryArrayItem{Query: *query, Bucket: "Left +Y Axis"}
	models.SetOptional(&queryItem.IsOverlay, false)

	queryObject := models.QueryObject{Queries: []models.QueryArrayItem{queryItem}}
	models.SetOptional(&queryObject.ImageType, models.ImageTypePNG)

	result, err := client.Query.Chart(context.Background(), queryObject)
	if err != nil {
		return err
	}
	fmt.Printf("Returned chart image type: %s\n", result.ImageType)
	filePath := path.Join(os.TempDir(), "chart.png")
	err = result.SaveImageAs(filePath)
	if err != nil {
		return err
	}
	cmd := exec.Command("firefox", filePath)
	err = cmd.Run()
	if err != nil {
		return err
	}
	fmt.Println()

	return nil
}

func runQueryURL() error {
	client, err := NewClient()
	if err != nil {
		return err
	}

	fmt.Println("### QUERY URL")
	result, err := client.Query.URL(context.Background(), makeQueryObject())
	if err != nil {
		return err
	}
	PrettyPrint(result)
	fmt.Println()

	return nil
}

func makeQueryObject() models.QueryObject {
	agg1 := models.NewAggregate("avg_bits_per_sec", "f_sum_both_bytes", models.AggregateFunctionTypeAverage)
	models.SetOptional(&agg1.Raw, true)
	agg2 := models.NewAggregate("p95th_bits_per_sec", "f_sum_both_bytes", models.AggregateFunctionTypePercentile)
	models.SetOptional(&agg2.Rank, 95)
	agg3 := models.NewAggregate("max_bits_per_sec", "f_sum_both_bytes", models.AggregateFunctionTypeMax)

	query := models.NewQuery(
		models.MetricTypeBytes,
		[]models.DimensionType{models.DimensionTypeTraffic},
	)
	query.Depth = 75
	query.LookbackSeconds = 600 // last 10 minutes
	query.HostnameLookup = true
	query.TopX = 8
	query.Depth = 75
	query.Aggregates = []models.Aggregate{agg1, agg2, agg3}
	models.SetOptional(&query.CIDR, 32)
	models.SetOptional(&query.CIDR6, 128)
	models.SetOptional(&query.Outsort, "avg_bits_per_sec")
	models.SetOptional(&query.AllSelected, true)

	queryItem := models.QueryArrayItem{Query: *query, Bucket: "Left +Y Axis"}

	return models.QueryObject{Queries: []models.QueryArrayItem{queryItem}}
}
