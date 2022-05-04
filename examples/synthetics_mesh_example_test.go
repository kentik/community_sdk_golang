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

	err = printPingLatencyMatrix(m)
	if err != nil {
		return fmt.Errorf("print ping latency matrix: %w", err)
	}

	err = printPingJitterMatrix(m)
	if err != nil {
		return fmt.Errorf("print ping jitter matrix: %w", err)
	}

	err = printPingPacketLossMatrix(m)
	if err != nil {
		return fmt.Errorf("print ping packet loss matrix: %w", err)
	}

	return nil
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

// metricsMatrix holds mesh test results for ping task.
type metricsMatrix struct {
	// agents hold agents data in the same order as cells.
	agents []*syntheticspb.Agent
	// cells hold all "fromAgent" -> "toAgent" connection metrics.
	cells map[string]map[string]*syntheticspb.PingResults
}

func newMetricsMatrix(trs []*syntheticspb.TestResults, allAgents []*syntheticspb.Agent) (metricsMatrix, error) {
	agents, err := prepareAgents(trs[0], allAgents)
	if err != nil {
		return metricsMatrix{}, fmt.Errorf("prepare agents: %w", err)
	}
	agentIPToAgentMap := makeAgentIPToAgentMap(agents)

	cells := make(map[string]map[string]*syntheticspb.PingResults)
	for _, tr := range trs {
		for _, agentResults := range tr.GetAgents() {
			fromAgentID := agentResults.GetAgentId()
			if cells[fromAgentID] == nil {
				cells[fromAgentID] = make(map[string]*syntheticspb.PingResults)
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

				if cells[fromAgentID][toAgent.GetId()] == nil {
					cells[fromAgentID][toAgent.GetId()] = ping
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

func (m metricsMatrix) GetPingResults(fromAgentID string, toAgentID string) *syntheticspb.PingResults {
	toAgents, ok := m.cells[fromAgentID]
	if !ok {
		return nil
	}

	pingResults, ok := toAgents[toAgentID]
	if !ok {
		return nil
	}

	return pingResults
}

func printPingLatencyMatrix(matrix metricsMatrix) error {
	w := makeTabWriter()

	fmt.Println(
		"Table cells contain ping latency and its health in format: " +
			"\"current [ms] / rolling avg. [ms] / rolling stddev. [ms] / health\"",
	)
	printMatrixHeader(matrix, w)
	printMatrixRows(matrix, w, formatPingLatency)
	_, _ = fmt.Fprintln(w)
	return w.Flush()
}

func printPingJitterMatrix(matrix metricsMatrix) error {
	w := makeTabWriter()

	fmt.Println(
		"Table cells contain ping jitter and its health in format: " +
			"\"current [ms] / rolling avg. [ms] / rolling stddev. [ms] / health\"",
	)
	printMatrixHeader(matrix, w)
	printMatrixRows(matrix, w, formatPingJitter)
	_, _ = fmt.Fprintln(w)
	return w.Flush()
}

func printPingPacketLossMatrix(matrix metricsMatrix) error {
	w := makeTabWriter()

	fmt.Println(
		"Table cells contain ping packet loss and its health in format: " +
			"\"current [%] / health\"",
	)
	printMatrixHeader(matrix, w)
	printMatrixRows(matrix, w, formatPingPacketLoss)
	_, _ = fmt.Fprintln(w)
	return w.Flush()
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

type formatCellFunc = func(*syntheticspb.PingResults) string

func printMatrixRows(matrix metricsMatrix, w *tabwriter.Writer, formatCell formatCellFunc) {
	for _, fromAgent := range matrix.agents {
		row := fromAgent.GetAlias() + "\t"
		for _, toAgent := range matrix.agents {
			pr := matrix.GetPingResults(fromAgent.GetId(), toAgent.GetId())
			row += formatCell(pr)
		}

		_, err := fmt.Fprintln(w, row)
		if err != nil {
			fmt.Printf("Warn: failed to print row: %v\n", err)
		}
	}
}

func formatPingLatency(pr *syntheticspb.PingResults) string {
	return formatMetricData(pr.GetLatency(), isCurrentMeasurementValid(pr))
}

func formatPingJitter(pr *syntheticspb.PingResults) string {
	return formatMetricData(pr.GetJitter(), isCurrentMeasurementValid(pr))
}

func formatPingPacketLoss(pr *syntheticspb.PingResults) string {
	pl := pr.GetPacketLoss()
	if pl == nil {
		return "[X]\t"
	}

	return fmt.Sprintf("%.2f %% / %v\t", toPercent(pl.GetCurrent()), pl.GetHealth())
}

func toPercent(v float64) float64 {
	return v * 100
}

func formatMetricData(md *syntheticspb.MetricData, isCurrentMeasurementValid bool) string {
	if md == nil {
		return "[X]\t"
	}

	return fmt.Sprintf(
		"%v / %v / %v / %v\t",
		formatCurrentMetricValue(md.GetCurrent(), isCurrentMeasurementValid),
		formatRollingMetricValue(md.GetRollingAvg()),
		formatRollingMetricValue(md.GetRollingStddev()),
		md.GetHealth(),
	)
}

func formatCurrentMetricValue(metricValue uint32, isCurrentMeasurementValid bool) string {
	if !isCurrentMeasurementValid {
		return "[X]"
	}
	return formatMetricValue(metricValue)
}

func formatRollingMetricValue(metricValue uint32) string {
	return formatMetricValue(metricValue)
}

// formatMetric formats the value of metric given in nanoseconds to milliseconds.
func formatMetricValue(metricValue uint32) string {
	return strconv.Itoa(int(metricValue) / 1000)
}

// isCurrentMeasurementValid returns true if current ping packet loss is less than 100%.
func isCurrentMeasurementValid(pr *syntheticspb.PingResults) bool {
	return pr.GetPacketLoss().GetCurrent() < 1
}
