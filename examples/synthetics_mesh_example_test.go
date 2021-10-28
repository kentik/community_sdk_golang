//+build examples

package examples

import (
	"context"
	"flag"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
	"text/tabwriter"
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/synthetics"
)

var testID = flag.String("testid", "3541", "id of mesh test to display the result matrix for")

func TestGetMeshTestResultsExample(t *testing.T) {
	assert.NoError(t, runGetMeshTestResults())
}

func runGetMeshTestResults() error {
	flag.Parse()

	mesh, err := getMeshTestResults(*testID)
	if mesh == nil {
		fmt.Println("Empty mesh test result received")
	} else {
		metricsMatrix := newMetricsMatrix(*mesh)
		printMetricsMatrix(metricsMatrix)
	}

	return err
}

func getMeshTestResults(testID string) (*[]synthetics.V202101beta1MeshResponse, error) {
	client := NewClient()

	healthPayload := *synthetics.NewV202101beta1GetHealthForTestsRequest()
	healthPayload.SetStartTime(time.Now().Add(-time.Minute * 5))
	healthPayload.SetEndTime(time.Now())
	healthPayload.SetIds([]string{testID})
	healthPayload.SetAugment(true) // if not set, returned Mesh pointer will be empty

	getHealthResp, httpResp, err := client.SyntheticsDataServiceAPI.
		GetHealthForTests(context.Background()).
		Body(healthPayload).
		Execute()
	if err != nil {
		fmt.Printf("%v %v", err, httpResp)
		return nil, err
	}

	if getHealthResp.Health != nil {
		healthItems := *getHealthResp.Health
		fmt.Println("Num health items:", len(healthItems))
		if len(healthItems) > 0 {
			return healthItems[0].Mesh, nil
		} else {
			return nil, nil
		}
	} else {
		fmt.Println("[no health items received]")
		return nil, nil
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
	latency, err := strconv.ParseInt(*metrics.GetLatency().Value, 10, 64)
	if err != nil {
		return "error"
	}

	return strconv.FormatInt(latency/1000, 10) + "ms" // latency is returned in thousands of milliseconds, so need to divide by 1000
}

// metricsMatrix holds "fromAgent" -> "toAgent" connection metrics
type metricsMatrix struct {
	agents []string
	cells  map[string]map[string]*synthetics.V202101beta1MeshMetrics
}

func newMetricsMatrix(mesh []synthetics.V202101beta1MeshResponse) metricsMatrix {
	// fill agents
	agents := []string{}
	for _, agent := range mesh {
		agents = append(agents, agent.GetAlias())
	}

	// fill matrix cells
	cells := make(map[string]map[string]*synthetics.V202101beta1MeshMetrics)
	for _, fromAgent := range mesh {
		cells[fromAgent.GetAlias()] = make(map[string]*synthetics.V202101beta1MeshMetrics)
		for _, toAgent := range *fromAgent.Columns {
			cells[fromAgent.GetAlias()][toAgent.GetAlias()] = toAgent.Metrics
		}
	}
	return metricsMatrix{agents: agents, cells: cells}
}

func (m metricsMatrix) getMetric(fromAgent string, toAgent string) (*synthetics.V202101beta1MeshMetrics, bool) {
	toAgents, ok := m.cells[fromAgent]
	if !ok {
		return nil, false
	}

	metric, ok := toAgents[toAgent]
	if !ok {
		return nil, false
	}

	return metric, true
}
