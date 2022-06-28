//go:build examples

//nolint:testpackage,forbidigo
package examples

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	syntheticspb "github.com/kentik/api-schema-public/gen/go/kentik/synthetics/v202202"
	"github.com/kentik/community_sdk_golang/kentikapi"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/kentik/community_sdk_golang/kentikapi/synthetics"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestDemonstrateSyntheticsAgentsAPI(t *testing.T) {
	t.Parallel()
	err := demonstrateSyntheticsAgentsAPI()
	assert.NoError(t, err)
}

func TestDemonstrateSyntheticsTestsAPI(t *testing.T) {
	t.Parallel()
	err := demonstrateSyntheticsTestsAPI()
	assert.NoError(t, err)
}

// TestDemonstrateSyntheticsTestsAPI_CreateMinimalTests demonstrates creating synthetics tests with only
// required fields set.
func TestDemonstrateSyntheticsTestsAPI_CreateMinimalTests(t *testing.T) {
	t.Parallel()
	err := createMinimalTests()
	assert.NoError(t, err)
}

func TestDemonstrateSyntheticsDataServiceAPI(t *testing.T) {
	t.Parallel()
	err := demonstrateSyntheticsDataServiceAPI()
	assert.NoError(t, err)
}

// demonstrateSyntheticsAgentsAPI demonstrates available methods of Synthetics Agent API.
// Delete method exists but is omitted here, because of lack of create method in the API.
// If you have no private agent at your environment, you can replace pickPrivateAgentID function call with
// pickDualIPRustAgentID. However, it is not possible to modify (update/activate/deactivate) global agent,
// so those pieces of code need to be commented out in such case.
func demonstrateSyntheticsAgentsAPI() error {
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
	if getAllResp.InvalidAgentsCount == 0 {
		fmt.Printf(
			"Kentik API returned %v invalid agents. Please, contact Kentik support.\n",
			getAllResp.InvalidAgentsCount,
		)
	}

	// Pick a private agent, so it is possible to modify it
	agentID, err := pickPrivateAgentID(getAllResp.Agents)
	if err != nil {
		return fmt.Errorf("pick private agent ID: %w", err)
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

func demonstrateSyntheticsTestsAPI() error {
	ctx := context.Background()
	client, err := NewClient()
	if err != nil {
		return err
	}

	fmt.Println("### Getting all synthetic tests")
	getAllResp, err := client.Synthetics.GetAllTests(ctx)
	if err != nil {
		return fmt.Errorf("client.Synthetics.GetAllTests: %w", err)
	}

	fmt.Println("Got all tests:", getAllResp)
	fmt.Println("Number of tests:", len(getAllResp.Tests))
	fmt.Println("Number of invalid tests:", getAllResp.InvalidTestsCount)

	fmt.Println("### Creating hostname synthetic test")
	test, err := newHostnameTest(ctx, client)
	if err != nil {
		return fmt.Errorf("new hostname test: %w", err)
	}

	test, err = client.Synthetics.CreateTest(ctx, test)
	if err != nil {
		return fmt.Errorf("client.SyntheticsAdmin.CreateTest: %w", err)
	}

	fmt.Println("Created test:")
	PrettyPrint(test)

	fmt.Println("### Getting created synthetic test")
	test, err = client.Synthetics.GetTest(ctx, test.ID)
	if err != nil {
		return fmt.Errorf("client.Synthetics.GetTest: %w", err)
	}

	fmt.Println("Got test:")
	PrettyPrint(test)
	fmt.Println("Test's target hostname:", test.Settings.GetHostnameDefinition().Target)

	fmt.Println("### Updating synthetic test")
	test.Name = "go-sdk-example-updated-hostname-test"
	test.Settings.Period = time.Second
	test.Settings.Ping.Timeout = time.Millisecond
	test.Settings.Traceroute.Limit = 1

	test, err = client.Synthetics.UpdateTest(ctx, test)
	if err != nil {
		return fmt.Errorf("client.SyntheticsAdmin.UpdateTest: %w", err)
	}

	fmt.Println("Updated test:")
	PrettyPrint(test)

	fmt.Println("### Setting synthetic test status to paused")
	err = client.Synthetics.SetTestStatus(ctx, test.ID, synthetics.TestStatusPaused)
	if err != nil {
		return fmt.Errorf("client.Synthetics.SetTestStatus: %w", err)
	}
	fmt.Println("Set synthetic test status successfully")

	fmt.Println("### Deleting synthetic test")
	err = client.Synthetics.DeleteTest(ctx, test.ID)
	if err != nil {
		return fmt.Errorf("client.Synthetics.DeleteTest: %w", err)
	}
	fmt.Println("Deleted synthetic test successfully")

	return nil
}

func newHostnameTest(ctx context.Context, client *kentikapi.Client) (*synthetics.Test, error) {
	getAllResp, err := client.Synthetics.GetAllAgents(ctx)
	if err != nil {
		return nil, fmt.Errorf("client.Synthetics.GetAllAgents: %w", err)
	}

	agentID, err := pickDualIPRustAgentID(getAllResp.Agents)
	if err != nil {
		return nil, fmt.Errorf("pick agent ID for hostname test: %w", err)
	}

	test := synthetics.NewHostnameTest(synthetics.HostnameTestRequiredFields{
		BasePingTraceTestRequiredFields: synthetics.BasePingTraceTestRequiredFields{
			BaseTestRequiredFields: synthetics.BaseTestRequiredFields{
				Name:     "go-sdk-example-hostname-test",
				AgentIDs: []string{agentID},
			},
			Ping: synthetics.PingSettingsRequiredFields{
				Timeout:  10 * time.Second,
				Count:    10,
				Protocol: synthetics.PingProtocolTCP,
				Port:     65535,
			},
			Traceroute: synthetics.TracerouteSettingsRequiredFields{
				Timeout:  59999 * time.Millisecond,
				Count:    5,
				Delay:    100 * time.Millisecond,
				Protocol: synthetics.TracerouteProtocolUDP,
				Limit:    255,
			},
		},
		Definition: synthetics.TestDefinitionHostnameRequiredFields{
			Target: "www.example.com",
		},
	})

	test.Settings.Period = 15 * time.Minute
	test.Settings.Family = synthetics.IPFamilyV4
	test.Settings.NotificationChannels = []string{} // must contain IDs of existing notification channels
	test.Settings.Health = synthetics.HealthSettings{
		LatencyCritical:           50 * time.Millisecond,
		LatencyWarning:            20 * time.Millisecond,
		LatencyCriticalStdDev:     100 * time.Millisecond,
		LatencyWarningStdDev:      100 * time.Millisecond,
		JitterCritical:            50 * time.Millisecond,
		JitterWarning:             20 * time.Millisecond,
		JitterCriticalStdDev:      100 * time.Millisecond,
		JitterWarningStdDev:       100 * time.Millisecond,
		PacketLossCritical:        100,
		PacketLossWarning:         100,
		HTTPLatencyCritical:       50 * time.Millisecond,
		HTTPLatencyWarning:        20 * time.Millisecond,
		HTTPLatencyCriticalStdDev: 100 * time.Millisecond,
		HTTPLatencyWarningStdDev:  100 * time.Millisecond,
		HTTPValidCodes:            []uint32{http.StatusOK, http.StatusCreated},
		DNSValidCodes:             []uint32{1, 2, 3},
		UnhealthySubtestThreshold: 2,
		AlarmActivation: &synthetics.AlarmActivationSettings{
			TimeWindow:  75 * time.Minute,
			Times:       4,
			GracePeriod: 3,
		},
	}
	test.Settings.Ping.Delay = 100 * time.Millisecond
	test.Settings.Ping.Port = 65535
	test.Settings.Traceroute.Port = 1

	return test, nil
}

func createMinimalTests() error {
	ctx := context.Background()
	client, err := NewClient()
	if err != nil {
		return err
	}

	getAllResp, err := client.Synthetics.GetAllAgents(ctx)
	if err != nil {
		return fmt.Errorf("client.Synthetics.GetAllAgents: %w", err)
	}

	dualIPRustAgentID, err := pickDualIPRustAgentID(getAllResp.Agents)
	if err != nil {
		return err
	}

	ipV4RustAgentID, err := pickIPV4RustAgentID(getAllResp.Agents)
	if err != nil {
		return err
	}

	nodeAgentID, err := pickNodeAgentID(getAllResp.Agents)
	if err != nil {
		return err
	}

	for _, t := range []*synthetics.Test{
		newMinimalIPTest([]models.ID{dualIPRustAgentID}),
		newMinimalNetworkGridTest([]models.ID{dualIPRustAgentID}),
		newMinimalHostnameTest([]models.ID{dualIPRustAgentID}),
		newMinimalAgentTest([]models.ID{dualIPRustAgentID}),
		newMinimalNetworkMeshTest([]models.ID{dualIPRustAgentID, ipV4RustAgentID}), // multiple agents required
		newMinimalFlowTest([]models.ID{dualIPRustAgentID}),
		newMinimalURLTest([]models.ID{dualIPRustAgentID}),
		newMinimalPageLoadTest([]models.ID{nodeAgentID}), // agent with implementation type Node required
		newMinimalDNSTest([]models.ID{dualIPRustAgentID}),
		newMinimalDNSGridTest([]models.ID{dualIPRustAgentID}),
	} {
		err = createAndDeleteTest(ctx, client, t)
		if err != nil {
			return err
		}
	}

	return nil
}

func createAndDeleteTest(ctx context.Context, client *kentikapi.Client, test *synthetics.Test) error {
	fmt.Println("### Creating synthetic test", test.Name)
	tName := test.Name // test object is nil on error
	test, err := client.Synthetics.CreateTest(ctx, test)
	if err != nil {
		return fmt.Errorf("create test %q: %w", tName, err)
	}

	fmt.Println("Created synthetic test:")
	PrettyPrint(test)

	fmt.Println("### Deleting synthetic test", test.Name)
	err = client.Synthetics.DeleteTest(ctx, test.ID)
	if err != nil {
		return fmt.Errorf("delete test %q: %w", test.Name, err)
	}
	fmt.Printf("Deleted synthetic test %q successfully\n", test.Name)
	return nil
}

func newMinimalIPTest(agentIDs []models.ID) *synthetics.Test {
	return synthetics.NewIPTest(synthetics.IPTestRequiredFields{
		BasePingTraceTestRequiredFields: synthetics.BasePingTraceTestRequiredFields{
			BaseTestRequiredFields: synthetics.BaseTestRequiredFields{
				Name:     "go-sdk-example-minimal-ip-test",
				AgentIDs: agentIDs,
			},
			Ping: synthetics.PingSettingsRequiredFields{
				Timeout:  10 * time.Second,
				Count:    10,
				Protocol: synthetics.PingProtocolTCP,
				Port:     65535,
			},
			Traceroute: synthetics.TracerouteSettingsRequiredFields{
				Timeout:  59999 * time.Millisecond,
				Count:    5,
				Delay:    100 * time.Millisecond,
				Protocol: synthetics.TracerouteProtocolUDP,
				Limit:    255,
			},
		},
		Definition: synthetics.TestDefinitionIPRequiredFields{
			Targets: []net.IP{net.ParseIP("192.0.2.213"), net.ParseIP("2001:db8:dead:beef:dead:beef:dead:beef")},
		},
	})
}

func newMinimalNetworkGridTest(agentIDs []models.ID) *synthetics.Test {
	return synthetics.NewNetworkGridTest(synthetics.NetworkGridTestRequiredFields{
		BasePingTraceTestRequiredFields: synthetics.BasePingTraceTestRequiredFields{
			BaseTestRequiredFields: synthetics.BaseTestRequiredFields{
				Name:     "go-sdk-example-minimal-network-grid-test",
				AgentIDs: agentIDs,
			},
			Ping: synthetics.PingSettingsRequiredFields{
				Timeout:  10 * time.Second,
				Count:    10,
				Protocol: synthetics.PingProtocolTCP,
				Port:     65535,
			},
			Traceroute: synthetics.TracerouteSettingsRequiredFields{
				Timeout:  59999 * time.Millisecond,
				Count:    5,
				Delay:    100 * time.Millisecond,
				Protocol: synthetics.TracerouteProtocolUDP,
				Limit:    255,
			},
		},
		Definition: synthetics.TestDefinitionNetworkGridRequiredFields{
			Targets: []net.IP{net.ParseIP("192.0.2.213"), net.ParseIP("2001:db8:dead:beef:dead:beef:dead:beef")},
		},
	})
}

func newMinimalHostnameTest(agentIDs []models.ID) *synthetics.Test {
	return synthetics.NewHostnameTest(synthetics.HostnameTestRequiredFields{
		BasePingTraceTestRequiredFields: synthetics.BasePingTraceTestRequiredFields{
			BaseTestRequiredFields: synthetics.BaseTestRequiredFields{
				Name:     "go-sdk-example-minimal-hostname-test",
				AgentIDs: agentIDs,
			},
			Ping: synthetics.PingSettingsRequiredFields{
				Timeout:  10 * time.Second,
				Count:    10,
				Protocol: synthetics.PingProtocolTCP,
				Port:     65535,
			},
			Traceroute: synthetics.TracerouteSettingsRequiredFields{
				Timeout:  59999 * time.Millisecond,
				Count:    5,
				Delay:    100 * time.Millisecond,
				Protocol: synthetics.TracerouteProtocolUDP,
				Limit:    255,
			},
		},
		Definition: synthetics.TestDefinitionHostnameRequiredFields{
			Target: "www.example.com",
		},
	})
}

func newMinimalAgentTest(agentIDs []models.ID) *synthetics.Test {
	return synthetics.NewAgentTest(synthetics.AgentTestRequiredFields{
		BasePingTraceTestRequiredFields: synthetics.BasePingTraceTestRequiredFields{
			BaseTestRequiredFields: synthetics.BaseTestRequiredFields{
				Name:     "go-sdk-example-minimal-agent-test",
				AgentIDs: agentIDs,
			},
			Ping: synthetics.PingSettingsRequiredFields{
				Timeout:  10 * time.Second,
				Count:    10,
				Protocol: synthetics.PingProtocolICMP,
			},
			Traceroute: synthetics.TracerouteSettingsRequiredFields{
				Timeout:  59999 * time.Millisecond,
				Count:    5,
				Delay:    100 * time.Millisecond,
				Protocol: synthetics.TracerouteProtocolUDP,
				Limit:    255,
			},
		},
		Definition: synthetics.TestDefinitionAgentRequiredFields{
			Target: agentIDs[0],
		},
	})
}

func newMinimalNetworkMeshTest(agentIDs []models.ID) *synthetics.Test {
	return synthetics.NewNetworkMeshTest(synthetics.NetworkMeshTestRequiredFields{
		BasePingTraceTestRequiredFields: synthetics.BasePingTraceTestRequiredFields{
			BaseTestRequiredFields: synthetics.BaseTestRequiredFields{
				Name:     "go-sdk-example-minimal-network-mesh-test",
				AgentIDs: agentIDs,
			},
			Ping: synthetics.PingSettingsRequiredFields{
				Timeout:  10 * time.Second,
				Count:    10,
				Protocol: synthetics.PingProtocolICMP,
			},
			Traceroute: synthetics.TracerouteSettingsRequiredFields{
				Timeout:  59999 * time.Millisecond,
				Count:    5,
				Delay:    100 * time.Millisecond,
				Protocol: synthetics.TracerouteProtocolUDP,
				Limit:    255,
			},
		},
	})
}

func newMinimalFlowTest(agentIDs []models.ID) *synthetics.Test {
	return synthetics.NewFlowTest(synthetics.FlowTestRequiredFields{
		BasePingTraceTestRequiredFields: synthetics.BasePingTraceTestRequiredFields{
			BaseTestRequiredFields: synthetics.BaseTestRequiredFields{
				Name:     "go-sdk-example-minimal-flow-test",
				AgentIDs: agentIDs,
			},
			Ping: synthetics.PingSettingsRequiredFields{
				Timeout:  10 * time.Second,
				Count:    10,
				Protocol: synthetics.PingProtocolICMP,
			},
			Traceroute: synthetics.TracerouteSettingsRequiredFields{
				Timeout:  59999 * time.Millisecond,
				Count:    5,
				Delay:    100 * time.Millisecond,
				Protocol: synthetics.TracerouteProtocolUDP,
				Limit:    255,
			},
		},
		Definition: synthetics.TestDefinitionFlowRequiredFields{
			Type:          synthetics.FlowTestTypeCity,
			Target:        "Warsaw",
			Direction:     synthetics.DirectionSrc,
			InetDirection: synthetics.DirectionDst,
		},
	})
}

func newMinimalURLTest(agentIDs []models.ID) *synthetics.Test {
	return synthetics.NewURLTest(synthetics.URLTestRequiredFields{
		BaseTestRequiredFields: synthetics.BaseTestRequiredFields{
			Name:     "go-sdk-example-minimal-url-test",
			AgentIDs: agentIDs,
		},
		Definition: synthetics.TestDefinitionURLRequiredFields{
			Target: url.URL{
				Scheme:   "https",
				Host:     "www.example.com:443",
				RawQuery: "dummy=query",
			},
			Timeout: time.Minute,
		},
	})
}

func newMinimalPageLoadTest(agentIDs []models.ID) *synthetics.Test {
	return synthetics.NewPageLoadTest(synthetics.PageLoadTestRequiredFields{
		BaseTestRequiredFields: synthetics.BaseTestRequiredFields{
			Name:     "go-sdk-example-minimal-page-load-test",
			AgentIDs: agentIDs,
		},
		Definition: synthetics.TestDefinitionPageLoadRequiredFields{
			Target: url.URL{
				Scheme:   "https",
				Host:     "www.example.com:443",
				RawQuery: "dummy=query",
			},
			Timeout: time.Minute,
		},
	})
}

func newMinimalDNSTest(agentIDs []models.ID) *synthetics.Test {
	return synthetics.NewDNSTest(synthetics.DNSTestRequiredFields{
		BaseTestRequiredFields: synthetics.BaseTestRequiredFields{
			Name:     "go-sdk-example-minimal-page-load-test",
			AgentIDs: agentIDs,
		},
		Definition: synthetics.TestDefinitionDNSRequiredFields{
			Target:     "www.example.com",
			Timeout:    time.Minute,
			RecordType: synthetics.DNSRecordAAAA,
			Servers:    []net.IP{net.ParseIP("192.0.2.213"), net.ParseIP("2001:db8:dead:beef:dead:beef:dead:beef")},
			Port:       53,
		},
	})
}

func newMinimalDNSGridTest(agentIDs []models.ID) *synthetics.Test {
	return synthetics.NewDNSGridTest(synthetics.DNSGridTestRequiredFields{
		BaseTestRequiredFields: synthetics.BaseTestRequiredFields{
			Name:     "go-sdk-example-minimal-page-load-test",
			AgentIDs: agentIDs,
		},
		Definition: synthetics.TestDefinitionDNSGridRequiredFields{
			Target:     "www.example.com",
			Timeout:    time.Minute,
			RecordType: synthetics.DNSRecordAAAA,
			Servers:    []net.IP{net.ParseIP("192.0.2.213"), net.ParseIP("2001:db8:dead:beef:dead:beef:dead:beef")},
			Port:       53,
		},
	})
}

func pickPrivateAgentID(agents []synthetics.Agent) (models.ID, error) {
	var matchedIDs []models.ID
	for _, a := range agents {
		if a.Type == synthetics.AgentTypePrivate {
			matchedIDs = append(matchedIDs, a.ID)
		}
	}

	if len(matchedIDs) == 0 {
		return "", fmt.Errorf("no agent meeting criteria (AgentTypePrivate) found")
	}

	// Randomizing picked ID helps to verify that criteria above are sufficient
	rand.Seed(time.Now().UnixNano())
	agentID := matchedIDs[rand.Intn(len(matchedIDs))] // nolint: gosec // no security concerns here
	log.Printf(
		"Found %v agents meeting criteria (AgentTypePrivate), picked agent with ID %v\n",
		len(matchedIDs), agentID,
	)
	return agentID, nil
}

func pickDualIPRustAgentID(agents []synthetics.Agent) (models.ID, error) {
	var matchedIDs []models.ID
	for _, a := range agents {
		if a.IPFamily == synthetics.IPFamilyDual && a.ImplementationType == synthetics.AgentImplementationTypeRust {
			matchedIDs = append(matchedIDs, a.ID)
		}
	}

	if len(matchedIDs) == 0 {
		return "", fmt.Errorf("no agent meeting criteria (IPFamilyDual, AgentImplementationTypeRust) found")
	}

	// Randomizing picked ID helps to verify that criteria above are sufficient
	rand.Seed(time.Now().UnixNano())
	agentID := matchedIDs[rand.Intn(len(matchedIDs))] // nolint: gosec // no security concerns here
	log.Printf(
		"Found %v agents meeting criteria (IPFamilyDual, AgentImplementationTypeRust), picked agent with ID %v\n",
		len(matchedIDs), agentID,
	)
	return agentID, nil
}

func pickIPV4RustAgentID(agents []synthetics.Agent) (models.ID, error) {
	var matchedIDs []models.ID
	for _, a := range agents {
		if a.IPFamily == synthetics.IPFamilyV4 && a.ImplementationType == synthetics.AgentImplementationTypeRust {
			matchedIDs = append(matchedIDs, a.ID)
		}
	}

	if len(matchedIDs) == 0 {
		return "", fmt.Errorf("no agent meeting criteria (IPFamilyV4, AgentImplementationTypeRust) found")
	}

	// Randomizing picked ID helps to verify that criteria above are sufficient
	rand.Seed(time.Now().UnixNano())
	agentID := matchedIDs[rand.Intn(len(matchedIDs))] // nolint: gosec // no security concerns here
	log.Printf(
		"Found %v agents meeting criteria (IPFamilyV4, AgentImplementationTypeRust), picked agent with ID %v\n",
		len(matchedIDs), agentID,
	)
	return agentID, nil
}

func pickNodeAgentID(agents []synthetics.Agent) (models.ID, error) {
	var matchedIDs []models.ID
	for _, a := range agents {
		if a.ImplementationType == synthetics.AgentImplementationTypeNode {
			matchedIDs = append(matchedIDs, a.ID)
		}
	}

	if len(matchedIDs) == 0 {
		return "", fmt.Errorf("no agent meeting criteria (AgentImplementationTypeNode) found")
	}

	// Randomizing picked ID helps to verify that criteria above are sufficient
	rand.Seed(time.Now().UnixNano())
	agentID := matchedIDs[rand.Intn(len(matchedIDs))] // nolint: gosec // no security concerns here
	log.Printf(
		"Found %v agents meeting criteria (AgentImplementationTypeNode), picked agent with ID %v\n",
		len(matchedIDs), agentID,
	)
	return agentID, nil
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
