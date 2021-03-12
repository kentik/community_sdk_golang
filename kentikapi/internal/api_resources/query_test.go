package api_resources_test

import (
	"context"
	"encoding/base64"
	"testing"
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_resources"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/utils"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQuerySQL(t *testing.T) {
	// arrange
	querySQL := `
		SELECT i_start_time,
		round(sum(in_pkts)/(3600)/1000) AS f_sum_in_pkts,
		round(sum(in_bytes)/(3600)/1000)*8 AS f_sum_in_bytes
		FROM all_devices
		WHERE ctimestamp > 3660 AND ctimestamp < 60
		GROUP by i_start_time
		ORDER by i_start_time DESC
		LIMIT 1000;
	`

	queryResponsePayload := `
	{
		"rows": [
			{
				"f_sum_in_bytes": 10,
				"f_sum_in_pkts": 20,
				"i_start_time": "2021-01-25T11:39:00Z"
			},
			{
				"f_sum_in_bytes": 50,
				"f_sum_in_pkts": 60,
				"i_start_time": "2021-01-25T11:38:00Z"
			},
			{
				"f_sum_in_bytes": 80,
				"f_sum_in_pkts": 90,
				"i_start_time": "2021-01-25T11:37:00Z"
			}
		]
	}`
	transport := &api_connection.StubTransport{ResponseBody: queryResponsePayload}
	queryAPI := api_resources.NewQueryAPI(transport)

	// act
	result, err := queryAPI.SQL(context.TODO(), querySQL)

	// assert request properly formed
	assert := assert.New(t)
	require := require.New(t)

	require.NoError(err)
	assert.Equal("/query/sql", transport.RequestPath)
	payload := utils.NewJSONPayloadInspector(t, transport.RequestBody)
	assert.Equal(querySQL, payload.String("query"))

	// and response properly parsed
	require.Equal(3, len(result.Rows))
	row0, ok := result.Rows[0].(map[string]interface{})
	require.True(ok)
	assert.Equal(10.0, row0["f_sum_in_bytes"])
	assert.Equal(20.0, row0["f_sum_in_pkts"])
	assert.Equal("2021-01-25T11:39:00Z", row0["i_start_time"])

	row1, ok := result.Rows[1].(map[string]interface{})
	require.True(ok)
	assert.Equal(50.0, row1["f_sum_in_bytes"])
	assert.Equal(60.0, row1["f_sum_in_pkts"])
	assert.Equal("2021-01-25T11:38:00Z", row1["i_start_time"])

	row2, ok := result.Rows[2].(map[string]interface{})
	require.True(ok)
	assert.Equal(80.0, row2["f_sum_in_bytes"])
	assert.Equal(90.0, row2["f_sum_in_pkts"])
	assert.Equal("2021-01-25T11:37:00Z", row2["i_start_time"])
}

func TestQueryData(t *testing.T) {
	// arrange
	queryResponsePayload := `
	{
		"results": [
			{
				"bucket": "Left +Y Axis",
				"data": [
					{
						"key": "Total",
						"avg_bits_per_sec": 19738.220765027323,
						"p95th_bits_per_sec": 22745.466666666667,
						"max_bits_per_sec": 25902.533333333333,
						"name": "Total",
						"timeSeries": {
							"both_bits_per_sec": {
								"flow": [
									[
										1608538980000,
										20751.333333333332,
										60
									],
									[
										1608539040000,
										16364.133333333333,
										60
									],
									[
										1608539100000,
										19316.933333333334,
										60
									]
								]
							}
						}
					}
				]
			}
		]
	}`
	transport := &api_connection.StubTransport{ResponseBody: queryResponsePayload}
	queryAPI := api_resources.NewQueryAPI(transport)

	agg1 := models.Aggregate{Name: "avg_bits_per_sec", Column: "f_sum_both_bytes", Fn: models.AggregateFunctionTypeAverage}
	models.SetOptional(&agg1.Raw, true)
	agg2 := models.Aggregate{Name: "p95th_bits_per_sec", Column: "f_sum_both_bytes", Fn: models.AggregateFunctionTypePercentile}
	models.SetOptional(&agg2.Rank, 95)
	agg3 := models.Aggregate{Name: "max_bits_per_sec", Column: "f_sum_both_bytes", Fn: models.AggregateFunctionTypeMax}
	query := models.NewQuery(
		models.MetricTypeBytes,
		[]models.DimensionType{models.DimensionTypeTraffic},
	)
	query.Depth = 75
	query.HostnameLookup = true
	query.Aggregates = []models.Aggregate{agg1, agg2, agg3}
	models.SetOptional(&query.StartingTime, time.Date(2001, 1, 1, 7, 45, 12, 234, time.UTC))
	models.SetOptional(&query.EndingTime, time.Date(2001, 11, 23, 14, 17, 43, 458, time.UTC))
	models.SetOptional(&query.CIDR, 32)
	models.SetOptional(&query.CIDR6, 128)
	models.SetOptional(&query.Outsort, "avg_bits_per_sec")
	models.SetOptional(&query.AllSelected, true)
	queryItem := models.QueryArrayItem{Query: *query, Bucket: "Left +Y Axis"}
	queryObject := models.QueryObject{Queries: []models.QueryArrayItem{queryItem}}

	// act
	result, err := queryAPI.Data(context.TODO(), queryObject)

	// assert request properly formed
	assert := assert.New(t)
	require := require.New(t)

	require.NoError(err)
	assert.Equal("/query/topXdata", transport.RequestPath)
	payload := utils.NewJSONPayloadInspector(t, transport.RequestBody)
	assert.Equal(1, payload.Count("queries/*"))
	assert.Equal("Left +Y Axis", payload.String("queries/*[1]/bucket"))
	assert.Equal(1, payload.Count("queries/*[1]/query/dimension/*"))
	assert.Equal("Traffic", payload.String("queries/*[1]/query/dimension/*[1]"))
	assert.Equal(32, payload.Int("queries/*[1]/query/cidr"))
	assert.Equal(128, payload.Int("queries/*[1]/query/cidr6"))
	assert.Equal("bytes", payload.String("queries/*[1]/query/metric"))
	assert.Equal(8, payload.Int("queries/*[1]/query/topx"))
	assert.Equal(75, payload.Int("queries/*[1]/query/depth"))
	assert.Equal("Auto", payload.String("queries/*[1]/query/fastData"))
	assert.Equal("avg_bits_per_sec", payload.String("queries/*[1]/query/outsort"))
	assert.Equal("2001-01-01 07:45:00", payload.String("queries/*[1]/query/starting_time"))
	assert.Equal("2001-11-23 14:17:00", payload.String("queries/*[1]/query/ending_time"))
	assert.True(payload.Bool("queries/*[1]/query/hostname_lookup"))
	assert.Equal(0, payload.Count("queries/*[1]/query/device_name/*"))
	assert.True(payload.Bool("queries/*[1]/query/all_selected"))
	assert.Equal("", payload.String("queries/*[1]/query/descriptor"))
	assert.Equal(3, payload.Count("queries/*[1]/query/aggregates/*"))
	// agg1
	assert.Equal("avg_bits_per_sec", payload.String("queries/*[1]/query/aggregates/*[1]/name"))
	assert.Equal("f_sum_both_bytes", payload.String("queries/*[1]/query/aggregates/*[1]/column"))
	assert.Equal("average", payload.String("queries/*[1]/query/aggregates/*[1]/fn"))
	assert.True(payload.Bool("queries/*[1]/query/aggregates/*[1]/raw"))
	// agg2
	assert.Equal("p95th_bits_per_sec", payload.String("queries/*[1]/query/aggregates/*[2]/name"))
	assert.Equal("f_sum_both_bytes", payload.String("queries/*[1]/query/aggregates/*[2]/column"))
	assert.Equal("percentile", payload.String("queries/*[1]/query/aggregates/*[2]/fn"))
	assert.Equal(95, payload.Int("queries/*[1]/query/aggregates/*[2]/rank"))
	// agg3
	assert.Equal("max_bits_per_sec", payload.String("queries/*[1]/query/aggregates/*[3]/name"))
	assert.Equal("f_sum_both_bytes", payload.String("queries/*[1]/query/aggregates/*[3]/column"))
	assert.Equal("max", payload.String("queries/*[1]/query/aggregates/*[3]/fn"))

	// and response properly parsed
	require.Len(result.Results, 1)
	row1, ok := result.Results[0].(map[string]interface{})
	require.True(ok)
	assert.Equal("Left +Y Axis", row1["bucket"])
	assert.NotZero(row1["data"])
}

func TestQueryChart(t *testing.T) {
	// arrange
	data := "ImageDataEncodedBase64=="
	queryResponsePayload := `{"dataUri": "data:image/png;base64,ImageDataEncodedBase64=="}`
	transport := &api_connection.StubTransport{ResponseBody: queryResponsePayload}
	queryAPI := api_resources.NewQueryAPI(transport)

	agg1 := models.Aggregate{Name: "avg_bits_per_sec", Column: "f_sum_both_bytes", Fn: models.AggregateFunctionTypeAverage}
	models.SetOptional(&agg1.Raw, true)
	agg2 := models.Aggregate{Name: "p95th_bits_per_sec", Column: "f_sum_both_bytes", Fn: models.AggregateFunctionTypePercentile}
	models.SetOptional(&agg2.Rank, 95)
	agg3 := models.Aggregate{Name: "max_bits_per_sec", Column: "f_sum_both_bytes", Fn: models.AggregateFunctionTypeMax}

	query := models.NewQuery(
		models.MetricTypeBytes,
		[]models.DimensionType{models.DimensionTypeTraffic},
	)
	// filter_ = Filter(filterField="dst_as", operator="=", filterValue="") // SavedFilters is dependency here
	// filter_group = FilterGroups(connector="All", not_=False, filters=[filter_]) // SavedFilters is dependency here
	filters := models.Filters{Connector: "", FilterGroups: nil} // filterGroups=[filter_group]) // SavedFilters is dependency here
	query.FiltersObj = &filters
	query.Aggregates = []models.Aggregate{agg1, agg2, agg3}
	query.LookbackSeconds = 0
	query.QueryTitle = "title"
	query.TopX = 8
	query.Depth = 75
	query.FastData = models.FastDataTypeAuto
	query.TimeFormat = models.TimeFormatLocal
	query.HostnameLookup = false
	query.DeviceName = []string{"dev1", "dev2"}
	query.MatrixBy = []string{models.DimensionTypeSrcGeoCity.String(), models.DimensionTypeDstGeoCity.String()}
	query.Descriptor = "descriptor"
	models.SetOptional(&query.StartingTime, time.Date(2001, 1, 1, 7, 45, 12, 234, time.UTC))
	models.SetOptional(&query.EndingTime, time.Date(2001, 11, 23, 14, 17, 43, 458, time.UTC))
	models.SetOptional(&query.CIDR, 32)
	models.SetOptional(&query.CIDR6, 128)
	models.SetOptional(&query.Outsort, "avg_bits_per_sec")
	models.SetOptional(&query.AllSelected, false)
	models.SetOptional(&query.VizType, models.ChartViewTypeStackedArea)
	models.SetOptional(&query.ShowOverlay, false)
	models.SetOptional(&query.OverlayDay, -7)
	models.SetOptional(&query.SyncAxes, false)
	models.SetOptional(&query.PPSThreshold, 1)

	queryItem := models.QueryArrayItem{Query: *query, Bucket: "Left +Y Axis"}
	models.SetOptional(&queryItem.IsOverlay, false)
	queryObject := models.QueryObject{Queries: []models.QueryArrayItem{queryItem}}
	models.SetOptional(&queryObject.ImageType, models.ImageTypePNG)

	// act
	result, err := queryAPI.Chart(context.TODO(), queryObject)

	// assert request properly formed
	assert := assert.New(t)
	require := require.New(t)

	require.NoError(err)
	assert.Equal("/query/topXchart", transport.RequestPath)
	payload := utils.NewJSONPayloadInspector(t, transport.RequestBody)
	assert.Equal(1, payload.Count("queries/*"))
	assert.Equal("Left +Y Axis", payload.String("queries/*[1]/bucket"))
	assert.Equal(1, payload.Count("queries/*[1]/query/dimension/*"))
	assert.Equal("Traffic", payload.String("queries/*[1]/query/dimension/*[1]"))
	assert.Equal(32, payload.Int("queries/*[1]/query/cidr"))
	assert.Equal(128, payload.Int("queries/*[1]/query/cidr6"))
	assert.Equal("bytes", payload.String("queries/*[1]/query/metric"))
	assert.Equal(8, payload.Int("queries/*[1]/query/topx"))
	assert.Equal(75, payload.Int("queries/*[1]/query/depth"))
	assert.Equal("Auto", payload.String("queries/*[1]/query/fastData"))
	assert.Equal("avg_bits_per_sec", payload.String("queries/*[1]/query/outsort"))
	assert.Equal(0, payload.Int("queries/*[1]/query/lookback_seconds"))
	assert.Equal("2001-01-01 07:45:00", payload.String("queries/*[1]/query/starting_time"))
	assert.Equal("2001-11-23 14:17:00", payload.String("queries/*[1]/query/ending_time"))
	assert.False(payload.Bool("queries/*[1]/query/hostname_lookup"))
	assert.Equal(2, payload.Count("queries/*[1]/query/device_name/*"))
	assert.Equal("dev1", payload.String("queries/*[1]/query/device_name/*[1]"))
	assert.Equal("dev2", payload.String("queries/*[1]/query/device_name/*[2]"))
	assert.False(payload.Bool("queries/*[1]/query/all_selected"))
	assert.Equal("descriptor", payload.String("queries/*[1]/query/descriptor"))
	assert.Equal("stackedArea", payload.String("queries/*[1]/query/viz_type"))
	assert.False(payload.Bool("queries/*[1]/query/show_overlay"))
	assert.Equal(-7, payload.Int("queries/*[1]/query/overlay_day"))
	assert.False(payload.Bool("queries/*[1]/query/sync_axes"))
	assert.Equal("title", payload.String("queries/*[1]/query/query_title"))
	assert.Equal(2, payload.Count("queries/*[1]/query/matrixBy/*"))
	assert.Equal("src_geo_city", payload.String("queries/*[1]/query/matrixBy/*[1]"))
	assert.Equal("dst_geo_city", payload.String("queries/*[1]/query/matrixBy/*[2]"))
	assert.Equal(0, payload.Count("queries/*[1]/query/saved_filters/*"))
	assert.Equal(1, payload.Int("queries/*[1]/query/pps_threshold"))
	assert.Equal("Local", payload.String("queries/*[1]/query/time_format"))
	// assert.Equal(1, payload.Count("queries/*[1]/query/filters_obj/filterGroups"))
	// assert.Equal("All", payload.String("queries/*[1]/query/filters_obj/filterGroups/*[1]/connector"))
	// assert.False(payload.Bool("queries/*[1]/query/filters_obj/filterGroups/*[1]/not"))
	// assert.Equal(1, payload.Count("queries/*[1]/query/filters_obj/filterGroups/*[1]/filters"))
	// assert.Equal("dst_as", payload.String("queries/*[1]/query/filters_obj/filterGroups/*[1]/filters/*[1]/filterField"))
	// assert.Equal("=", payload.String("queries/*[1]/query/filters_obj/filterGroups/*[1]/filters/*[1]/operator"))
	// assert.Equal("", payload.String("queries/*[1]/query/filters_obj/filterGroups/*[1]/filters/*[1]/filterValue"))

	assert.Equal(3, payload.Count("queries/*[1]/query/aggregates/*"))
	// agg1
	assert.Equal("avg_bits_per_sec", payload.String("queries/*[1]/query/aggregates/*[1]/name"))
	assert.Equal("f_sum_both_bytes", payload.String("queries/*[1]/query/aggregates/*[1]/column"))
	assert.Equal("average", payload.String("queries/*[1]/query/aggregates/*[1]/fn"))
	assert.True(payload.Bool("queries/*[1]/query/aggregates/*[1]/raw"))
	// agg2
	assert.Equal("p95th_bits_per_sec", payload.String("queries/*[1]/query/aggregates/*[2]/name"))
	assert.Equal("f_sum_both_bytes", payload.String("queries/*[1]/query/aggregates/*[2]/column"))
	assert.Equal("percentile", payload.String("queries/*[1]/query/aggregates/*[2]/fn"))
	assert.Equal(95, payload.Int("queries/*[1]/query/aggregates/*[2]/rank"))
	// agg3
	assert.Equal("max_bits_per_sec", payload.String("queries/*[1]/query/aggregates/*[3]/name"))
	assert.Equal("f_sum_both_bytes", payload.String("queries/*[1]/query/aggregates/*[3]/column"))
	assert.Equal("max", payload.String("queries/*[1]/query/aggregates/*[3]/fn"))

	// and response properly parsed
	assert.Equal(models.ImageTypePNG, result.ImageType)
	decodedData, err := base64.StdEncoding.DecodeString(data)
	require.NoError(err)
	assert.Equal(decodedData, result.ImageData)
}

func TestQueryURL(t *testing.T) {
	// arrange
	unquotedResponse := "https://portal.kentik.com/portal/#Charts/shortUrl/e0d24b3cc8dfe41f9093668e531cbe96"
	queryResponsePayload := `"` + unquotedResponse + `"` // actual response is url in quotation marks
	transport := &api_connection.StubTransport{ResponseBody: queryResponsePayload}
	queryAPI := api_resources.NewQueryAPI(transport)

	agg1 := models.Aggregate{Name: "avg_bits_per_sec", Column: "f_sum_both_bytes", Fn: models.AggregateFunctionTypeAverage}
	models.SetOptional(&agg1.Raw, true)
	agg2 := models.Aggregate{Name: "p95th_bits_per_sec", Column: "f_sum_both_bytes", Fn: models.AggregateFunctionTypePercentile}
	models.SetOptional(&agg2.Rank, 95)
	agg3 := models.Aggregate{Name: "max_bits_per_sec", Column: "f_sum_both_bytes", Fn: models.AggregateFunctionTypeMax}
	query := models.NewQuery(
		models.MetricTypeBytes,
		[]models.DimensionType{models.DimensionTypeTraffic},
	)
	query.Depth = 75
	query.HostnameLookup = true
	query.Aggregates = []models.Aggregate{agg1, agg2, agg3}
	models.SetOptional(&query.StartingTime, time.Date(2001, 1, 1, 7, 45, 12, 234, time.UTC))
	models.SetOptional(&query.EndingTime, time.Date(2001, 11, 23, 14, 17, 43, 458, time.UTC))
	models.SetOptional(&query.CIDR, 32)
	models.SetOptional(&query.CIDR6, 128)
	models.SetOptional(&query.Outsort, "avg_bits_per_sec")
	models.SetOptional(&query.AllSelected, true)
	queryItem := models.QueryArrayItem{Query: *query, Bucket: "Left +Y Axis"}
	queryObject := models.QueryObject{Queries: []models.QueryArrayItem{queryItem}}

	// act
	result, err := queryAPI.URL(context.TODO(), queryObject)

	// assert request properly formed
	assert := assert.New(t)
	require := require.New(t)

	require.NoError(err)
	assert.Equal("/query/url", transport.RequestPath)
	payload := utils.NewJSONPayloadInspector(t, transport.RequestBody)
	assert.Equal(1, payload.Count("queries/*"))
	assert.Equal("Left +Y Axis", payload.String("queries/*[1]/bucket"))
	assert.Equal(1, payload.Count("queries/*[1]/query/dimension/*"))
	assert.Equal("Traffic", payload.String("queries/*[1]/query/dimension/*[1]"))
	assert.Equal(32, payload.Int("queries/*[1]/query/cidr"))
	assert.Equal(128, payload.Int("queries/*[1]/query/cidr6"))
	assert.Equal("bytes", payload.String("queries/*[1]/query/metric"))
	assert.Equal(8, payload.Int("queries/*[1]/query/topx"))
	assert.Equal(75, payload.Int("queries/*[1]/query/depth"))
	assert.Equal("Auto", payload.String("queries/*[1]/query/fastData"))
	assert.Equal("avg_bits_per_sec", payload.String("queries/*[1]/query/outsort"))
	assert.Equal("2001-01-01 07:45:00", payload.String("queries/*[1]/query/starting_time"))
	assert.Equal("2001-11-23 14:17:00", payload.String("queries/*[1]/query/ending_time"))
	assert.True(payload.Bool("queries/*[1]/query/hostname_lookup"))
	assert.Equal(0, payload.Count("queries/*[1]/query/device_name/*"))
	assert.True(payload.Bool("queries/*[1]/query/all_selected"))
	assert.Equal("", payload.String("queries/*[1]/query/descriptor"))
	assert.Equal(3, payload.Count("queries/*[1]/query/aggregates/*"))
	// agg1
	assert.Equal("avg_bits_per_sec", payload.String("queries/*[1]/query/aggregates/*[1]/name"))
	assert.Equal("f_sum_both_bytes", payload.String("queries/*[1]/query/aggregates/*[1]/column"))
	assert.Equal("average", payload.String("queries/*[1]/query/aggregates/*[1]/fn"))
	assert.True(payload.Bool("queries/*[1]/query/aggregates/*[1]/raw"))
	// agg2
	assert.Equal("p95th_bits_per_sec", payload.String("queries/*[1]/query/aggregates/*[2]/name"))
	assert.Equal("f_sum_both_bytes", payload.String("queries/*[1]/query/aggregates/*[2]/column"))
	assert.Equal("percentile", payload.String("queries/*[1]/query/aggregates/*[2]/fn"))
	assert.Equal(95, payload.Int("queries/*[1]/query/aggregates/*[2]/rank"))
	// agg3
	assert.Equal("max_bits_per_sec", payload.String("queries/*[1]/query/aggregates/*[3]/name"))
	assert.Equal("f_sum_both_bytes", payload.String("queries/*[1]/query/aggregates/*[3]/column"))
	assert.Equal("max", payload.String("queries/*[1]/query/aggregates/*[3]/fn"))

	// and response properly parsed
	assert.Equal(unquotedResponse, result.URL)
}
