//go:build examples

//nolint:testpackage,forbidigo
package examples

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"text/tabwriter"
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/kentik/community_sdk_golang/kentikapi/synthetics"
	"github.com/stretchr/testify/assert"
)

func TestDemonstrateSyntheticsTestResultsAPI(t *testing.T) {
	t.Parallel()
	err := demonstrateSyntheticsTestResultsAPI()
	assert.NoError(t, err)
}

func TestDemonstrateSyntheticsTestResultsAPI_NetworkMesh(t *testing.T) {
	t.Parallel()
	err := demonstrateSyntheticsNetworkMeshTestResults()
	assert.NoError(t, err)
}

func demonstrateSyntheticsTestResultsAPI() error {
	ctx := context.Background()
	client, err := NewClient()
	if err != nil {
		return err
	}

	test, err := pickNetworkMeshTest(ctx, client)
	if err != nil {
		return err
	}

	fmt.Println("### Getting results for test with ID", test.ID)
	results, err := client.Synthetics.GetResultsForTests(ctx, synthetics.GetResultsForTestsRequest{
		TestIDs:   []string{test.ID},
		StartTime: time.Now().Add(-time.Hour * 240), // last 10 days
		EndTime:   time.Now(),
	})
	if err != nil {
		return fmt.Errorf("client.Synthetics.GetResultsForTests: %w", err)
	}

	resultsJSON, _ := json.Marshal(results) // nolint:errcheck
	fmt.Println("Got test results (JSON-formatted):", string(resultsJSON))
	fmt.Println("Got test results (simplified):", formatTestResultsSlice(results))
	fmt.Println("Number of test results:", len(results))

	fmt.Println("### Getting traceroute results for test with ID", test.ID)
	traceResp, err := client.Synthetics.GetTraceForTest(ctx, synthetics.GetTraceForTestRequest{
		TestID:    test.ID,
		StartTime: time.Now().Add(-time.Hour * 240), // last 10 days
		EndTime:   time.Now(),
	})
	if err != nil {
		return fmt.Errorf("client.Synthetics.GetTraceForTest: %w", err)
	}

	traceRespJSON, _ := json.Marshal(traceResp) // nolint:errcheck
	fmt.Println("Got traceroute results for test (in JSON format):", string(traceRespJSON))
	fmt.Println("Number of nodes:", len(traceResp.Nodes))
	fmt.Println("Number of paths:", len(traceResp.Paths))
	return nil
}

func pickNetworkMeshTest(ctx context.Context, c *kentikapi.Client) (synthetics.Test, error) {
	getAllResp, err := c.Synthetics.GetAllTests(ctx)
	if err != nil {
		return synthetics.Test{}, fmt.Errorf("c.Synthetics.GetAllTests: %w", err)
	}

	for _, test := range getAllResp.Tests {
		if test.Type == synthetics.TestTypeNetworkMesh {
			fmt.Printf("Picked network_mesh test named %q with ID %v\n", test.Name, test.ID)
			return test, nil
		}
	}
	return synthetics.Test{}, errors.New("no network_mesh tests found")
}

func formatTestResultsSlice(trs []synthetics.TestResults) string {
	var s []string
	for _, tr := range trs {
		s = append(s, formatTestResults(tr))
	}
	return fmt.Sprintf("{\n%v\n}", strings.Join(s, ", "))
}

func formatTestResults(tr synthetics.TestResults) string {
	return fmt.Sprintf(
		"{\n  test_id=%v\n  time=%v\n  health=%v\n  len(agents)=%v\n  agents=%v\n}",
		tr.TestID, tr.Time, tr.Health, len(tr.Agents), formatAgentsResults(tr.Agents),
	)
}

func formatAgentsResults(ars []synthetics.AgentResults) string {
	var s []string
	for _, ar := range ars {
		s = append(s, fmt.Sprintf("{agent_id=%v health=%v len(tasks)=%v}", ar.AgentID, ar.Health, len(ar.Tasks)))
	}
	return fmt.Sprintf("{%v}", strings.Join(s, ", "))
}

func demonstrateSyntheticsNetworkMeshTestResults() error {
	ctx := context.Background()
	client, err := NewClient()
	if err != nil {
		return err
	}

	test, err := pickNetworkMeshTest(ctx, client)
	if err != nil {
		return err
	}

	trs, err := getLastTenTestResults(ctx, client, test.ID)
	if err != nil {
		return err
	}
	if len(trs) == 0 {
		fmt.Println("No mesh test results received - exiting example")
		return nil
	}

	agents, err := prepareAgents(ctx, client, test)
	if err != nil {
		return fmt.Errorf("prepare agents: %w", err)
	}
	fmt.Println("Test agents: ", test.Settings.AgentIDs)
	fmt.Println("Prepared agents: ", agents)

	m, err := newMetricsMatrix(trs, agents)
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

func getLastTenTestResults(ctx context.Context, c *kentikapi.Client, testID string) ([]synthetics.TestResults, error) {
	trs, err := c.Synthetics.GetResultsForTests(ctx, synthetics.GetResultsForTestsRequest{
		TestIDs: []models.ID{testID},
		// Last 12 hours; should provide sufficient number of results (assuming 1 hour max test period)
		// It could be optimized by picking 10-12 test periods instead to lower the number of retrieved results
		StartTime: time.Now().Add(-time.Hour * 12),
		EndTime:   time.Now(),
	})
	if err != nil {
		return nil, fmt.Errorf("GetResultsForTests: %w", err)
	}

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

// metricsMatrix holds network mesh test results for ping task.
type metricsMatrix struct {
	// agents hold agents that are involved in the synthetic test.
	agents []synthetics.Agent
	// cells hold all "fromAgent" -> "toAgent" connection metrics.
	cells map[string]map[string]*synthetics.PingResults
}

func newMetricsMatrix(trs []synthetics.TestResults, agents []synthetics.Agent) (metricsMatrix, error) {
	agentIPToAgentMap := makeAgentIPToAgentMap(agents)
	cells := make(map[string]map[string]*synthetics.PingResults)
	for _, tr := range trs {
		for _, agentResults := range tr.Agents {
			fromAgentID := agentResults.AgentID
			if cells[fromAgentID] == nil {
				cells[fromAgentID] = make(map[string]*synthetics.PingResults)
			}

			for _, taskResult := range agentResults.Tasks {
				if taskResult.TaskType != synthetics.TaskTypePing {
					continue
				}

				ping := taskResult.GetPingResults()
				if ping == nil {
					return metricsMatrix{}, errors.New("ping results is nil and should not be")
				}

				toAgent, ok := agentIPToAgentMap[ping.Target]
				if !ok {
					fmt.Printf(
						"Ignoring ping results to target IP %v - no such agent IP in test configuration\n",
						ping.Target,
					)
					continue
				}

				if cells[fromAgentID][toAgent.ID] == nil {
					cells[fromAgentID][toAgent.ID] = ping
				}
			}
		}
	}

	fmt.Println("Latest test results time:", trs[0].Time)
	return metricsMatrix{agents: agents, cells: cells}, nil
}

// prepareAgents returns agents involved in given synthetic test.
func prepareAgents(ctx context.Context, c *kentikapi.Client, test synthetics.Test) ([]synthetics.Agent, error) {
	aResp, err := c.Synthetics.GetAllAgents(ctx)
	if err != nil {
		return nil, err
	}

	agentIDToAgentMap := makeAgentIDToAgentMap(aResp.Agents)
	var agents []synthetics.Agent
	for _, agentID := range test.Settings.AgentIDs {
		a, ok := agentIDToAgentMap[agentID]
		if !ok {
			return nil, fmt.Errorf("agent with ID %v not found in listed agents", agentID)
		}
		agents = append(agents, a)
	}

	return agents, nil
}

func makeAgentIDToAgentMap(agents []synthetics.Agent) map[string]synthetics.Agent {
	m := make(map[string]synthetics.Agent, len(agents))
	for _, a := range agents {
		m[a.ID] = a
	}
	return m
}

func makeAgentIPToAgentMap(agents []synthetics.Agent) map[string]synthetics.Agent {
	m := make(map[string]synthetics.Agent, len(agents))
	for _, a := range agents {
		m[a.IP] = a
	}
	return m
}

func (m metricsMatrix) GetPingResults(fromAgentID string, toAgentID string) *synthetics.PingResults {
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
		header = header + x.Alias + "\t"
	}

	if _, err := fmt.Fprintln(w, header); err != nil {
		fmt.Printf("Warn: failed to print header: %v\n", err)
	}
}

type formatCellFunc = func(*synthetics.PingResults) string

func printMatrixRows(matrix metricsMatrix, w *tabwriter.Writer, formatCell formatCellFunc) {
	for _, fromAgent := range matrix.agents {
		row := fromAgent.Alias + "\t"
		for _, toAgent := range matrix.agents {
			pr := matrix.GetPingResults(fromAgent.ID, toAgent.ID)
			row += formatCell(pr)
		}

		_, err := fmt.Fprintln(w, row)
		if err != nil {
			fmt.Printf("Warn: failed to print row: %v\n", err)
		}
	}
}

const noResult = "[X]\t"

func formatPingLatency(pr *synthetics.PingResults) string {
	if pr == nil {
		return noResult
	}

	return formatMetricData(pr.Latency, isCurrentMeasurementValid(*pr))
}

func formatPingJitter(pr *synthetics.PingResults) string {
	if pr == nil {
		return noResult
	}

	return formatMetricData(pr.Jitter, isCurrentMeasurementValid(*pr))
}

func formatPingPacketLoss(pr *synthetics.PingResults) string {
	if pr == nil {
		return noResult
	}

	return fmt.Sprintf("%.2f %% / %v\t", toPercent(pr.PacketLoss.Current), pr.PacketLoss.Health)
}

func toPercent(v float64) float64 {
	return v * 100
}

func formatMetricData(md synthetics.MetricData, isCurrentMeasurementValid bool) string {
	return fmt.Sprintf(
		"%v / %v / %v / %v\t",
		formatCurrentMetricValue(md.Current, isCurrentMeasurementValid),
		formatRollingMetricValue(md.RollingAvg),
		formatRollingMetricValue(md.RollingStdDev),
		md.Health,
	)
}

func formatCurrentMetricValue(metricValue time.Duration, isCurrentMeasurementValid bool) string {
	if !isCurrentMeasurementValid {
		return "[X]"
	}
	return formatMetricValue(metricValue)
}

func formatRollingMetricValue(metricValue time.Duration) string {
	return formatMetricValue(metricValue)
}

// formatMetric formats the value of metric to milliseconds.
func formatMetricValue(metricValue time.Duration) string {
	return strconv.FormatInt(metricValue.Milliseconds(), 10)
}

// isCurrentMeasurementValid returns true if current ping packet loss is less than 100%.
func isCurrentMeasurementValid(pr synthetics.PingResults) bool {
	return pr.PacketLoss.Current < 1
}
