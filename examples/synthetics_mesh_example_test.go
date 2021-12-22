//go:build examples
// +build examples

//nolint:testpackage,forbidigo
package examples

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"text/tabwriter"
	"time"

	syntheticspb "github.com/kentik/api-schema-public/gen/go/kentik/synthetics/v202101beta1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGetMeshTestResultsExample(t *testing.T) {
	assert.NoError(t, runGetMeshTestResults())
}

func runGetMeshTestResults() error {
	testID, err := pickTestID()
	if err != nil {
		return err
	}

	mesh, err := getMeshTestResults(testID)
	if err != nil {
		return err
	}
	if mesh == nil {
		fmt.Println("Empty mesh test result received")
	} else {
		metricsMatrix := newMetricsMatrix(mesh)
		err = printMetricsMatrix(metricsMatrix)
	}

	return err
}

func getMeshTestResults(testID string) ([]*syntheticspb.MeshResponse, error) {
	client, err := NewClient()
	if err != nil {
		return nil, err
	}

	healthPayload := &syntheticspb.GetHealthForTestsRequest{
		Ids:       []string{testID},
		StartTime: timestamppb.New(time.Now().Add(-time.Minute * 5)),
		EndTime:   timestamppb.Now(),
		Augment:   true, // if not set, returned Mesh pointer will be empty
	}

	getHealthResp, err := client.SyntheticsData.GetHealthForTests(context.Background(), healthPayload)
	if err != nil {
		return nil, err
	}

	if getHealthResp.Health != nil {
		healthItems := getHealthResp.GetHealth()
		fmt.Println("Num health items:", len(healthItems))
		if len(healthItems) > 0 {
			return healthItems[0].GetMesh(), nil
		}
		return nil, nil
	}

	fmt.Println("[no health items received]")
	return nil, nil
}

func printMetricsMatrix(matrix metricsMatrix) error {
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
				row += formatLatency(metrics) + "\t"
			} else {
				row += "[X]\t"
			}
		}
		fmt.Fprintln(w, row)
	}

	if err := w.Flush(); err != nil {
		return err
	}
	return nil
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

func formatLatency(metrics *syntheticspb.MeshMetrics) string {
	// latency is returned in thousands of milliseconds, so need to divide by 1000
	return strconv.FormatInt(metrics.GetLatency().Value/1000, 10) + "ms"
}

// metricsMatrix holds "fromAgent" -> "toAgent" connection metrics.
type metricsMatrix struct {
	agents []string
	cells  map[string]map[string]*syntheticspb.MeshMetrics
}

func newMetricsMatrix(mesh []*syntheticspb.MeshResponse) metricsMatrix {
	// fill agents
	agents := []string{}
	for _, agent := range mesh {
		agents = append(agents, agent.GetAlias())
	}

	// fill matrix cells
	cells := make(map[string]map[string]*syntheticspb.MeshMetrics)
	for _, fromAgent := range mesh {
		cells[fromAgent.GetAlias()] = make(map[string]*syntheticspb.MeshMetrics)
		for _, toAgent := range fromAgent.Columns {
			cells[fromAgent.GetAlias()][toAgent.GetAlias()] = toAgent.Metrics
		}
	}
	return metricsMatrix{agents: agents, cells: cells}
}

func (m metricsMatrix) getMetric(fromAgent string, toAgent string) (*syntheticspb.MeshMetrics, bool) {
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

func pickTestID() (string, error) {
	client, err := NewClient()
	if err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	getAllResp, err := client.SyntheticsAdmin.ListTests(ctx, &syntheticspb.ListTestsRequest{})
	if err != nil {
		return "", err
	}

	if getAllResp.Tests != nil {
		for _, test := range getAllResp.GetTests() {
			if test.GetType() == "application_mesh" {
				return test.GetId(), nil
			}
		}
	}
	return "", fmt.Errorf("No tests with type application_mesh for requested Kentik account: %v", err)
}
