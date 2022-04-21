//go:build examples

//nolint:testpackage,forbidigo
package examples

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	syntheticspb "github.com/kentik/api-schema-public/gen/go/kentik/synthetics/v202202"
	"github.com/kentik/community_sdk_golang/kentikapi"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestDemonstrateSyntheticsAgentsAPI(t *testing.T) {
	t.Parallel()
	err := demonstrateSyntheticsAgentAPI()
	assert.NoError(t, err)
}

func TestDemonstrateSyntheticsTestsAPI(t *testing.T) {
	t.Parallel()
	err := demonstrateSyntheticsTestsAPI()
	assert.NoError(t, err)
}

func TestDemonstrateSyntheticsDataServiceAPI(t *testing.T) {
	t.Parallel()
	err := demonstrateSyntheticsDataServiceAPI()
	assert.NoError(t, err)
}

// demonstrateSyntheticsAgentAPI demonstrates available methods of Synthetics Agent API.
// Note that there is no create method in the API.
// Delete method exists but is omitted here, because of lack of create method.
func demonstrateSyntheticsAgentAPI() error {
	ctx := context.Background()
	client, err := NewClient()
	if err != nil {
		return err
	}

	fmt.Println("### Getting all synthetics agents")
	getAllResp, err := client.Synthetics.GetAllAgents(ctx)
	if err != nil {
		return fmt.Errorf("client.Synthetics.GetAll: %w", err)
	}

	fmt.Printf("Got all agents: %v\n", getAllResp)
	fmt.Println("Number of agents:", len(getAllResp.Agents))
	fmt.Println("Number of invalid agents:", getAllResp.InvalidAgentsCount)

	agentID, err := pickPrivateAgentID(ctx)
	if err != nil {
		return fmt.Errorf("pick agent ID: %w", err)
	}

	fmt.Println("### Getting synthetics agent with ID", agentID)
	agent, err := client.Synthetics.GetAgent(ctx, agentID)
	if err != nil {
		return fmt.Errorf("client.Synthetics.Get: %w", err)
	}

	fmt.Println("Got agent:")
	PrettyPrint(agent)

	fmt.Println("### Updating synthetic agent")
	originalAlias := agent.Alias
	agent.Alias = "go-sdk-example-updated-alias"

	agent, err = client.Synthetics.UpdateAgent(ctx, agent)
	if err != nil {
		return fmt.Errorf("client.SyntheticsAdmin.UpdateAgent: %w", err)
	}

	fmt.Println("Updated agent:", agent)
	PrettyPrint(agent)

	fmt.Println("### Activating the synthetics agent")
	originalStatus := agent.Status
	agent, err = client.Synthetics.ActivateAgent(ctx, agentID)
	if err != nil {
		return fmt.Errorf("client.Synthetics.Activate: %w", err)
	}

	fmt.Println("Activated agent:")
	PrettyPrint(agent)

	fmt.Println("### Deactivating the synthetics agent")
	agent, err = client.Synthetics.DeactivateAgent(ctx, agentID)
	if err != nil {
		return fmt.Errorf("client.Synthetics.Deactivate: %w", err)
	}

	fmt.Println("Deactivated agent:")
	PrettyPrint(agent)

	fmt.Println("### Updating synthetic agent to revert changes")
	agent.Alias = originalAlias
	agent.Status = originalStatus

	agent, err = client.Synthetics.UpdateAgent(ctx, agent)
	if err != nil {
		return fmt.Errorf("client.SyntheticsAdmin.UpdateAgent (revert): %w", err)
	}
	fmt.Println("Updated agent:")
	PrettyPrint(agent)

	return nil
}

func pickPrivateAgentID(ctx context.Context) (string, error) {
	client, err := NewClient()
	if err != nil {
		return "", err
	}

	getAllResp, err := client.SyntheticsAdmin.ListAgents(ctx, &syntheticspb.ListAgentsRequest{})
	if err != nil {
		return "", fmt.Errorf("client.SyntheticsAdmin.ListAgents: %w", err)
	}

	if getAllResp.GetAgents() != nil {
		for _, agent := range getAllResp.GetAgents() {
			if agent.GetType() == "private" {
				return agent.GetId(), nil
			}
		}
	}
	return "", fmt.Errorf("no private agent found: %w", err)
}

func demonstrateSyntheticsTestsAPI() error {
	ctx := context.Background()
	client, err := NewClient()
	if err != nil {
		return err
	}

	fmt.Println("### Getting all synthetic tests")
	getAllResp, err := client.SyntheticsAdmin.ListTests(ctx, &syntheticspb.ListTestsRequest{})
	if err != nil {
		return fmt.Errorf("client.SyntheticsAdmin.ListTests: %w", err)
	}

	fmt.Println("Got all tests:", getAllResp)
	fmt.Println("Number of tests:", len(getAllResp.GetTests()))
	fmt.Println("Number of invalid tests:", getAllResp.InvalidCount)
	fmt.Println()

	fmt.Println("### Creating hostname synthetic test")
	createResp, err := client.SyntheticsAdmin.CreateTest(ctx, &syntheticspb.CreateTestRequest{Test: makeHostnameTest()})
	if err != nil {
		return fmt.Errorf("client.SyntheticsAdmin.CreateTest: %w", err)
	}

	fmt.Println("Created test:", createResp.String())
	fmt.Println()

	fmt.Println("### Setting synthetic test status to paused")
	_, err = client.SyntheticsAdmin.SetTestStatus(ctx, &syntheticspb.SetTestStatusRequest{
		Id:     createResp.Test.Id,
		Status: syntheticspb.TestStatus_TEST_STATUS_PAUSED,
	})
	if err != nil {
		return fmt.Errorf("client.SyntheticsAdmin.SetTestStatus: %w", err)
	}
	fmt.Println("Set synthetic test status successfully")
	fmt.Println()

	fmt.Println("### Getting created synthetic test")
	getReqPayLoad := &syntheticspb.GetTestRequest{Id: createResp.Test.Id}
	getResp, err := client.SyntheticsAdmin.GetTest(ctx, getReqPayLoad)
	if err != nil {
		return fmt.Errorf("client.SyntheticsAdmin.GetTest: %w", err)
	}
	fmt.Println("Got test:", getResp)
	fmt.Println()

	fmt.Println("### Updating synthetic test")
	test := getResp.Test
	test.Name = "go-sdk-updated-hostname-test"

	updateResp, err := client.SyntheticsAdmin.UpdateTest(ctx, &syntheticspb.UpdateTestRequest{Test: test})
	if err != nil {
		return fmt.Errorf("client.SyntheticsAdmin.UpdateTest: %w", err)
	}
	fmt.Println("Updated test:", updateResp)
	fmt.Println()

	fmt.Println("### Deleting synthetic test")
	_, err = client.SyntheticsAdmin.DeleteTest(ctx, &syntheticspb.DeleteTestRequest{Id: test.Id})
	if err != nil {
		return fmt.Errorf("client.SyntheticsAdmin.DeleteTest: %w", err)
	}
	fmt.Println("Deleted synthetic test successfully")
	fmt.Println()

	return nil
}

func makeHostnameTest() *syntheticspb.Test {
	return &syntheticspb.Test{
		Name:   "go-sdk-example-hostname-test",
		Type:   "hostname",
		Status: syntheticspb.TestStatus_TEST_STATUS_ACTIVE,
		Settings: &syntheticspb.TestSettings{
			Definition: &syntheticspb.TestSettings_Hostname{
				Hostname: &syntheticspb.HostnameTest{Target: "www.example.com"},
			},
			AgentIds: []string{"890"},
			Tasks: []string{
				"ping",
				"traceroute",
			},
			HealthSettings: &syntheticspb.HealthSettings{
				LatencyCritical:           1,
				LatencyWarning:            2,
				PacketLossCritical:        3,
				PacketLossWarning:         4,
				JitterCritical:            5,
				JitterWarning:             6,
				HttpLatencyCritical:       7,
				HttpLatencyWarning:        8,
				HttpValidCodes:            []uint32{200, 201},
				DnsValidCodes:             []uint32{1, 2, 3},
				LatencyCriticalStddev:     9,
				LatencyWarningStddev:      10,
				JitterCriticalStddev:      11,
				JitterWarningStddev:       12,
				HttpLatencyCriticalStddev: 13,
				HttpLatencyWarningStddev:  14,
				UnhealthySubtestThreshold: 15,
				Activation: &syntheticspb.ActivationSettings{
					GracePeriod: "2",
					TimeUnit:    "m",
					TimeWindow:  "5",
					Times:       "3",
				},
			},
			Ping: &syntheticspb.TestPingSettings{
				Count:    10,
				Protocol: "icmp",
				Port:     0,
				Timeout:  10000,
				Delay:    100,
			},
			Trace: &syntheticspb.TestTraceSettings{
				Count:    5,
				Protocol: "tcp",
				Port:     443,
				Timeout:  59999,
				Limit:    255,
				Delay:    100,
			},
			Period: 60,
			Family: syntheticspb.IPFamily_IP_FAMILY_DUAL,
		},
	}
}

func demonstrateSyntheticsDataServiceAPI() error {
	ctx := context.Background()

	client, err := NewClient()
	if err != nil {
		return err
	}

	testID, err := pickNetworkMeshTestID(ctx, client)
	if err != nil {
		return err
	}

	if err = getResultsForTest(ctx, client, testID); err != nil {
		return err
	}

	if err = getTraceForTest(ctx, client, testID); err != nil {
		return err
	}

	return nil
}

func pickNetworkMeshTestID(ctx context.Context, c *kentikapi.Client) (string, error) {
	getAllResp, err := c.SyntheticsAdmin.ListTests(ctx, &syntheticspb.ListTestsRequest{})
	if err != nil {
		return "", fmt.Errorf("c.SyntheticsAdmin.ListTests: %w", err)
	}

	if getAllResp.Tests != nil {
		for _, test := range getAllResp.GetTests() {
			if test.GetType() == "network_mesh" {
				fmt.Printf("Picked network_mesh test named %q with ID %v\n", test.GetName(), test.GetId())
				return test.GetId(), nil
			}
		}
	}
	return "", errors.New("no network_mesh tests found")
}

func getResultsForTest(ctx context.Context, client *kentikapi.Client, testID string) error {
	fmt.Println("### Getting results for test with ID", testID)
	resp, err := client.SyntheticsData.GetResultsForTests(ctx, &syntheticspb.GetResultsForTestsRequest{
		Ids:       []string{testID},
		StartTime: timestamppb.New(time.Now().Add(-time.Hour * 24000)), // last 1000 days
		EndTime:   timestamppb.Now(),
	})
	if err != nil {
		return fmt.Errorf("client.SyntheticsData.GetResultsForTests: %w", err)
	}

	fmt.Println("Got test results:", formatTestResultsSlice(resp.GetResults()))
	fmt.Println("Number of test results:", len(resp.GetResults()))
	fmt.Println()

	return nil
}

func formatTestResultsSlice(trs []*syntheticspb.TestResults) string {
	var s []string
	for _, tr := range trs {
		s = append(s, formatTestResults(tr))
	}
	return fmt.Sprintf("{\n%v\n}", strings.Join(s, ", "))
}

func formatTestResults(tr *syntheticspb.TestResults) string {
	return fmt.Sprintf(
		"{\n  test_id=%v\n  time=%v\n  health=%v\n  len(agents)=%v\n  agents=%v\n}",
		tr.GetTestId(), tr.GetTime().AsTime(), tr.GetHealth(), len(tr.GetAgents()), formatAgentsResults(tr.GetAgents()),
	)
}

func formatAgentsResults(ars []*syntheticspb.AgentResults) string {
	var s []string
	for _, ar := range ars {
		s = append(s, fmt.Sprintf("{agent_id=%v health=%v len(tasks)=%v}", ar.GetAgentId(), ar.GetHealth(), len(ar.GetTasks())))
	}
	return fmt.Sprintf("{%v}", strings.Join(s, ", "))
}

func getTraceForTest(ctx context.Context, client *kentikapi.Client, testID string) error {
	fmt.Println("### Getting trace for test with ID", testID)
	resp, err := client.SyntheticsData.GetTraceForTest(ctx, &syntheticspb.GetTraceForTestRequest{
		Id:        testID,
		StartTime: timestamppb.New(time.Now().Add(-time.Hour * 24000)), // last 1000 days
		EndTime:   timestamppb.Now(),
	})
	if err != nil {
		return fmt.Errorf("client.SyntheticsData.GetTraceForTest: %w", err)
	}

	fmt.Println("Got trace for test")
	fmt.Println("Number of nodes:", len(resp.GetNodes()))
	fmt.Println("Number of paths:", len(resp.GetPaths()))
	fmt.Println()

	return nil
}
