package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/kentik/community_sdk_golang/apiv6/examples"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/synthetics"
)

var testID = flag.String("testid", "3541", "id of mesh test to display the result matrix for")

func main() {
	flag.Parse()

	mesh := getMeshTestResults(*testID)
	if mesh == nil {
		fmt.Println("Empty mesh test result received")
		os.Exit(1)
	}

	metricsMatrix := newMetricsMatrix(*mesh)
	printMetricsMatrix(metricsMatrix)
}

func getMeshTestResults(testID string) *[]synthetics.V202101beta1MeshResponse {
	client := examples.NewClient()

	healthPayload := *synthetics.NewV202101beta1GetHealthForTestsRequest()
	healthPayload.SetStartTime(time.Now().Add(-time.Minute * 5))
	healthPayload.SetEndTime(time.Now())
	healthPayload.SetIds([]string{testID})
	healthPayload.SetAugment(true) // if not set, returned Mesh pointer will be empty
	getHealthReq := client.SyntheticsDataServiceApi.GetHealthForTests(context.Background()).V202101beta1GetHealthForTestsRequest(healthPayload)
	getHealthResp, httpResp, err := getHealthReq.Execute()
	if err != nil {
		fmt.Printf("%v %v", err, httpResp)
		return nil
	}
	if getHealthResp.Health != nil {
		healthItems := *getHealthResp.Health
		fmt.Println("Num health items:", len(healthItems))
		if len(healthItems) > 0 {
			return healthItems[0].Mesh
		} else {
			return nil
		}
	} else {
		fmt.Println("[no health items received]")
		return nil
	}
}

func printMetricsMatrix(matrix metricsMatrix) {
	w := makeTabWriter()

	// print table header
	header := ".\t"
	for _, x := range matrix.agents {
		header = header + x + "\t"
	}
	fmt.Fprintln(w, header)

	// print table rows
	for _, fromAgent := range matrix.agents {
		row := fromAgent + "\t"
		for _, toAgent := range matrix.agents {
			if metrics, ok := matrix.getMetric(fromAgent, toAgent); ok {
				row = row + formatLatency(metrics) + "\t"
			} else {
				row = row + "[X]\t"
			}
		}
		fmt.Fprintln(w, row)
	}

	w.Flush()
}

func makeTabWriter() *tabwriter.Writer {
	const minWidth = 0  // minimal cell width including any padding
	const tabWidth = 2  // width of tab characters (equivalent number of spaces)
	const padding = 4   // distance between cells
	const padchar = ' ' // ASCII char used for padding
	const flags = 0     // formatting control
	w := tabwriter.NewWriter(os.Stdout, minWidth, tabWidth, padding, padchar, flags)
	return w
}

func formatLatency(metrics *synthetics.V202101beta1MeshMetrics) string {
	latencyMS, err := strconv.ParseInt(*metrics.GetLatency().Value, 10, 64)
	if err != nil {
		return "error"
	}

	return strconv.FormatInt(latencyMS/1000, 10) + "ms" // latency is returned in thousands of milliseconds, so need to divide by 1000
}
