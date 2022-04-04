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

	syntheticspb "github.com/kentik/api-schema-public/gen/go/kentik/synthetics/v202202"
	"github.com/kentik/community_sdk_golang/kentikapi"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestDemonstrateSyntheticsNetworkMeshTestResults(t *testing.T) {
	t.Parallel()
	err := demonstrateSyntheticsNetworkMeshTestResults()
	assert.NoError(t, err)
}

func demonstrateSyntheticsNetworkMeshTestResults() error {
	ctx := context.Background()
	client, err := NewClient()
	if err != nil {
		return err
	}

	testID, err := pickNetworkMeshTestID(ctx, client)
	if err != nil {
		return err
	}

	trs, err := getLastTenTestResults(ctx, client, testID)
	if err != nil {
		return err
	}
	if len(trs) == 0 {
		fmt.Println("No mesh test results received - exiting example")
		return nil
	}

	getAllAgentsResp, err := client.SyntheticsAdmin.ListAgents(ctx, &syntheticspb.ListAgentsRequest{})
	if err != nil {
		return fmt.Errorf("client.SyntheticsAdmin.ListAgents: %w", err)
	}

	m, err := newMetricsMatrix(trs, getAllAgentsResp.GetAgents())
	if err != nil {
		return fmt.Errorf("new metrics matrix: %w", err)
	}

	return printMetricsMatrix(m)
}

func getLastTenTestResults(ctx context.Context, c *kentikapi.Client, testID string) ([]*syntheticspb.TestResults, error) {
	resp, err := c.SyntheticsData.GetResultsForTests(ctx, &syntheticspb.GetResultsForTestsRequest{
		Ids:       []string{testID},
		StartTime: timestamppb.New(time.Now().Add(-time.Hour * 240000)), // 1000 days
		EndTime:   timestamppb.Now(),
	})
	if err != nil {
		return nil, fmt.Errorf("GetResultsForTests: %w", err)
	}

	trs := resp.GetResults()
	if len(trs) == 0 {
		return nil, nil
	}

	fmt.Println("Number of test results:", len(trs))
	// latest test trs are returned first in the array
	return trs[0:min(10, len(trs))], nil
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// metricsMatrix holds mesh test result data for single point of time.
type metricsMatrix struct {
	// agents hold agents data in the same order as cells.
	agents []*syntheticspb.Agent
	// cells hold all "fromAgent" -> "toAgent" connection metrics.
	cells map[string]map[string]*syntheticspb.MetricData
}

func newMetricsMatrix(trs []*syntheticspb.TestResults, allAgents []*syntheticspb.Agent) (metricsMatrix, error) {
	agents, err := prepareAgents(trs[0], allAgents)
	if err != nil {
		return metricsMatrix{}, fmt.Errorf("prepare agents: %w", err)
	}
	agentIPToAgentMap := makeAgentIPToAgentMap(agents)

	cells := make(map[string]map[string]*syntheticspb.MetricData)
	for _, tr := range trs {
		for _, agentResults := range tr.GetAgents() {
			fromAgentID := agentResults.GetAgentId()
			if cells[fromAgentID] == nil {
				cells[fromAgentID] = make(map[string]*syntheticspb.MetricData)
			}

			for _, taskResult := range agentResults.GetTasks() {
				ping := taskResult.GetPing()
				if ping == nil {
					continue
				}

				toAgent, ok := agentIPToAgentMap[ping.Target]
				if !ok {
					return metricsMatrix{}, fmt.Errorf("agent with IP %q not found", ping.Target)
				}

				if cells[fromAgentID][toAgent.GetId()] == nil && ping.Latency != nil {
					cells[fromAgentID][toAgent.GetId()] = ping.Latency
				}
			}
		}
	}

	fmt.Println("Latest test results time:", trs[0].Time.AsTime())
	return metricsMatrix{agents: agents, cells: cells}, nil
}

// prepareAgents prepares agents that are involved in test results.
func prepareAgents(tr *syntheticspb.TestResults, allAgents []*syntheticspb.Agent) ([]*syntheticspb.Agent, error) {
	agentIDToAgentMap := makeAgentIDToAgentMap(allAgents)

	var agents []*syntheticspb.Agent
	for _, ars := range tr.GetAgents() {
		a, ok := agentIDToAgentMap[ars.GetAgentId()]
		if !ok {
			return nil, fmt.Errorf("agent with ID %v not found in listed agents", ars.GetAgentId())
		}
		agents = append(agents, a)
	}
	return agents, nil
}

func makeAgentIDToAgentMap(agents []*syntheticspb.Agent) map[string]*syntheticspb.Agent {
	m := make(map[string]*syntheticspb.Agent)
	for _, a := range agents {
		m[a.GetId()] = a
	}
	return m
}

func makeAgentIPToAgentMap(agents []*syntheticspb.Agent) map[string]*syntheticspb.Agent {
	m := make(map[string]*syntheticspb.Agent)
	for _, a := range agents {
		m[a.GetIp()] = a
	}
	return m
}

func printMetricsMatrix(matrix metricsMatrix) error {
	w := makeTabWriter()

	fmt.Println(
		"Table cells contain ping latency and connection health in format: " +
			"\"current [ms] / rolling avg. [ms] / rolling stddev. [ms] / health\"",
	)
	printMatrixHeader(matrix, w)
	printMatrixRows(matrix, w)

	if err := w.Flush(); err != nil {
		return fmt.Errorf("flush tab writer: %w", err)
	}
	return nil
}

func makeTabWriter() *tabwriter.Writer {
	const minWidth = 0  // minimal cell width including any padding
	const tabWidth = 2  // width of tab characters (equivalent number of spaces)
	const padding = 4   // distance between cells
	const padChar = ' ' // ASCII char used for padding
	const flags = 0     // formatting control
	return tabwriter.NewWriter(os.Stdout, minWidth, tabWidth, padding, padChar, flags)
}

func printMatrixHeader(matrix metricsMatrix, w *tabwriter.Writer) {
	header := ".\t"
	for _, x := range matrix.agents {
		header = header + x.GetAlias() + "\t"
	}

	if _, err := fmt.Fprintln(w, header); err != nil {
		fmt.Printf("Warn: failed to print header: %v\n", err)
	}
}

func printMatrixRows(matrix metricsMatrix, w *tabwriter.Writer) {
	for _, fromAgent := range matrix.agents {
		row := fromAgent.GetAlias() + "\t"
		for _, toAgent := range matrix.agents {
			row += formatCell(matrix.getMetric(fromAgent.GetId(), toAgent.GetId()))
		}

		_, err := fmt.Fprintln(w, row)
		if err != nil {
			fmt.Printf("Warn: failed to print row: %v\n", err)
		}
	}
}

func formatCell(metrics *syntheticspb.MetricData) string {
	if metrics == nil {
		return "[X]\t"
	}

	return fmt.Sprintf(
		"%v / %v / %v / %v\t",
		formatMetricValue(metrics.GetCurrent()),
		formatMetricValue(metrics.GetRollingAvg()),
		formatMetricValue(metrics.GetRollingStddev()),
		metrics.GetHealth(),
	)
}

// formatCell formats the value of metric given in nanosecond to millisecond value.
func formatMetricValue(metricValue uint32) string {
	if metricValue == 0 {
		return "[X]"
	}
	return strconv.Itoa(int(metricValue) / 1000)
}

func (m metricsMatrix) getMetric(fromAgentID string, toAgentID string) *syntheticspb.MetricData {
	toAgents, ok := m.cells[fromAgentID]
	if !ok {
		return nil
	}

	metric, ok := toAgents[toAgentID]
	if !ok {
		return nil
	}

	return metric
}
