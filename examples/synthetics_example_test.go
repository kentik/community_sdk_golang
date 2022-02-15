//go:build examples
// +build examples

//nolint:testpackage,forbidigo
package examples

import (
	"context"
	"fmt"
	"testing"
	"time"

	syntheticspb "github.com/kentik/api-schema-public/gen/go/kentik/synthetics/v202101beta1"
	"github.com/kentik/community_sdk_golang/kentikapi"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestSyntheticsAPIExample(t *testing.T) {
	t.Parallel()
	assert.NoError(t, runAdminServiceExamples())
	assert.NoError(t, runDataServiceExamples())
}

func runAdminServiceExamples() error {
	ctx := context.Background()
	client, err := NewClient()
	if err != nil {
		return err
	}

	if err = runCRUDTest(ctx, client); err != nil {
		fmt.Println(err)
		return err
	}

	if err = runListTests(ctx, client); err != nil {
		fmt.Println(err)
		return err
	}

	if err = runCRUDAgent(ctx, client); err != nil {
		fmt.Println(err)
		return err
	}

	if err = runListAgents(ctx, client); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func runDataServiceExamples() error {
	ctx := context.Background()

	client, err := NewClient()
	if err != nil {
		return err
	}

	testID, err := pickTestID(ctx)
	if err != nil {
		return err
	}

	if err = runGetHealthForTest(ctx, client, testID); err != nil {
		fmt.Println(err)
		return err
	}

	if err = runGetTraceForTest(ctx, client, testID); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func runCRUDTest(ctx context.Context, client *kentikapi.Client) error {
	fmt.Println("### CREATE TEST")
	test := makeExampleTest()
	createReqPayload := &syntheticspb.CreateTestRequest{Test: test}
	createResp, err := client.SyntheticsAdmin.CreateTest(ctx, createReqPayload)
	if err != nil {
		return err
	}

	PrettyPrint(createResp.Test)
	fmt.Println("Created test")
	testID := createResp.Test.Id

	fmt.Println("### SET TEST STATUS")
	setStatusReqPayload := &syntheticspb.SetTestStatusRequest{
		Id:     testID,
		Status: syntheticspb.TestStatus_TEST_STATUS_PAUSED,
	}

	_, err = client.SyntheticsAdmin.SetTestStatus(ctx, setStatusReqPayload)
	if err != nil {
		return err
	}
	fmt.Println("Success")
	fmt.Println()

	fmt.Println("### GET TEST")
	getReqPayLoad := &syntheticspb.GetTestRequest{Id: testID}
	getResp, err := client.SyntheticsAdmin.GetTest(ctx, getReqPayLoad)
	if err != nil {
		return err
	}
	PrettyPrint(getResp.GetTest())
	fmt.Println()

	test = getResp.Test

	fmt.Println("### PATCH TEST")
	test.Name = "example-test-1 UPDATED"
	patchReqPayload := &syntheticspb.PatchTestRequest{
		Test: test,
		Mask: &fieldmaskpb.FieldMask{Paths: []string{"test.name"}},
	}

	patchResp, err := client.SyntheticsAdmin.PatchTest(ctx, patchReqPayload)
	if err != nil {
		return err
	}
	PrettyPrint(patchResp.GetTest())
	fmt.Println()

	fmt.Println("### DELETE TEST")
	deleteReqPayLoad := &syntheticspb.DeleteTestRequest{Id: testID}
	_, err = client.SyntheticsAdmin.DeleteTest(ctx, deleteReqPayLoad)
	if err != nil {
		return err
	}
	fmt.Println("Success")
	fmt.Println()

	return nil
}

func runListTests(ctx context.Context, client *kentikapi.Client) error {
	fmt.Println("### LIST TESTS")

	getAllResp, err := client.SyntheticsAdmin.ListTests(ctx, &syntheticspb.ListTestsRequest{})
	if err != nil {
		return err
	}

	if getAllResp.Tests != nil {
		tests := getAllResp.GetTests()
		fmt.Println("Num tests:", len(tests))
		if getAllResp.InvalidTestsCount != 0 {
			fmt.Println("Num invalid tests:", getAllResp.InvalidTestsCount)
		}
		PrettyPrint(tests)
	} else {
		fmt.Println("[no tests received]")
	}
	fmt.Println()

	return nil
}

func runGetHealthForTest(ctx context.Context, client *kentikapi.Client, testID string) error {
	fmt.Println("### GET HEALTH FOR TESTS")

	healthPayLoad := &syntheticspb.GetHealthForTestsRequest{
		Ids:       []string{testID},
		StartTime: timestamppb.New(time.Now().Add(-time.Hour)),
		EndTime:   timestamppb.Now(),
	}

	getHealthResp, err := client.SyntheticsData.GetHealthForTests(ctx, healthPayLoad)
	if err != nil {
		return err
	}

	if getHealthResp.Health != nil {
		healthItems := getHealthResp.GetHealth()
		fmt.Println("Num health items:", len(healthItems))
		PrettyPrint(healthItems)
	} else {
		fmt.Println("[no health items received]")
	}
	fmt.Println()

	return nil
}

func runGetTraceForTest(ctx context.Context, client *kentikapi.Client, testID string) error {
	fmt.Println("### GET TRACE FOR TEST")

	tracePayLoad := &syntheticspb.GetTraceForTestRequest{
		Id:        testID,
		StartTime: timestamppb.New(time.Now().Add(-time.Hour)),
		EndTime:   timestamppb.Now(),
	}

	getTraceResp, err := client.SyntheticsData.GetTraceForTest(ctx, tracePayLoad)
	if err != nil {
		return err
	}

	if getTraceResp.Lookups != nil {
		lookups := getTraceResp.Lookups
		fmt.Println("Agents id by ip:")
		PrettyPrint(lookups)
	} else {
		fmt.Println("[no agents received]")
	}

	if getTraceResp.TraceRoutes != nil {
		results := getTraceResp.TraceRoutes
		fmt.Println("Num trace routes:", len(results))
		PrettyPrint(results)
	} else {
		fmt.Println("[no trace routes received]")
	}

	if getTraceResp.TraceRoutesInfo != nil {
		results := getTraceResp.TraceRoutesInfo
		fmt.Println("Trace routes info:")
		PrettyPrint(results)
	} else {
		fmt.Println("[no trace routes info received]")
	}

	fmt.Println()

	return nil
}

func runCRUDAgent(ctx context.Context, client *kentikapi.Client) error {
	// NOTE: no CREATE method exists for agents in the API, thus no example for CREATE
	// NOTE: agent of id 1717 must exist
	agentID, err := pickAgentID()
	if err != nil {
		return err
	}

	fmt.Println("### GET AGENT")
	getReqPayLoad := &syntheticspb.GetAgentRequest{Id: agentID}
	getResp, err := client.SyntheticsAdmin.GetAgent(ctx, getReqPayLoad)
	if err != nil {
		return err
	}
	PrettyPrint(getResp.GetAgent())
	fmt.Println()

	fmt.Println("### PATCH AGENT")
	agent := getResp.GetAgent()
	if agent.GetStatus() == syntheticspb.AgentStatus_AGENT_STATUS_OK {
		agent.Status = syntheticspb.AgentStatus_AGENT_STATUS_WAIT
	} else {
		agent.Status = syntheticspb.AgentStatus_AGENT_STATUS_OK
	}
	agent.Name = ""
	patchReqPayload := &syntheticspb.PatchAgentRequest{
		Agent: agent,
		Mask:  &fieldmaskpb.FieldMask{Paths: []string{"agent.status"}},
	}
	patchResp, err := client.SyntheticsAdmin.PatchAgent(ctx, patchReqPayload)
	if err != nil {
		return err
	}
	PrettyPrint(patchResp.GetAgent())
	fmt.Println()

	// NOTE: as we can't create agents through the API - let's not delete them
	// fmt.Println("### DELETE AGENT")
	// deleteReqPayLoad := &syntheticspb.DeleteAgentRequest{Id: agentID}
	// _, err = client.SyntheticsAdmin.DeleteAgent(ctx, deleteReqPayLoad)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println("Success")
	// fmt.Println()

	return nil
}

func runListAgents(ctx context.Context, client *kentikapi.Client) error {
	fmt.Println("### LIST AGENTS")

	getAllReq := &syntheticspb.ListAgentsRequest{}
	getAllResp, err := client.SyntheticsAdmin.ListAgents(ctx, getAllReq)
	if err != nil {
		return err
	}

	if getAllResp.Agents != nil {
		agents := getAllResp.GetAgents()
		fmt.Println("Num agents:", len(agents))
		if getAllResp.InvalidAgentsCount != 0 {
			fmt.Println("Num invalid agents:", getAllResp.InvalidAgentsCount)
		}
		PrettyPrint(agents)
	} else {
		fmt.Println("[no agents received]")
	}
	fmt.Println()

	return nil
}

func makeExampleTest() *syntheticspb.Test {
	healthSettings := &syntheticspb.HealthSettings{
		LatencyCritical:     0,
		LatencyWarning:      0,
		PacketLossCritical:  0,
		PacketLossWarning:   0,
		JitterCritical:      0,
		JitterWarning:       0,
		HttpLatencyCritical: 0,
		HttpLatencyWarning:  0,
		HttpValidCodes:      []uint32{},
		DnsValidCodes:       []uint32{},
	}

	monitorSettings := &syntheticspb.TestMonitoringSettings{
		ActivationGracePeriod: "2",
		ActivationTimeUnit:    "m",
		ActivationTimeWindow:  "5",
		ActivationTimes:       "3",
		NotificationChannels:  []string{},
	}

	ping := &syntheticspb.TestPingSettings{
		Period: 60,
		Count:  5,
		Expiry: 3000,
	}

	trace := &syntheticspb.TestTraceSettings{
		Period:   60,
		Protocol: "udp",
		Port:     33434,
		Expiry:   22500,
		Limit:    30,
	}

	settings := &syntheticspb.TestSettings{
		Definition: &syntheticspb.TestSettings_Hostname{
			Hostname: &syntheticspb.HostnameTest{Target: "example.com"},
		},
		AgentIds: []string{
			"890",
		},
		Period: 0,
		Count:  0,
		Expiry: 0,
		Limit:  0,
		Tasks: []string{
			"ping",
			"traceroute",
		},
		HealthSettings:     healthSettings,
		MonitoringSettings: monitorSettings,
		Ping:               ping,
		Trace:              trace,
		Port:               443,
		Protocol:           "icmp",
		Family:             syntheticspb.IPFamily_IP_FAMILY_DUAL,
		Servers:            []string{},
		UseLocalIp:         false,
		Reciprocal:         false,
		RollupLevel:        1,
	}

	test := &syntheticspb.Test{
		Name:     "example-test-1",
		Type:     "hostname",
		DeviceId: "0",
		Status:   syntheticspb.TestStatus_TEST_STATUS_ACTIVE,
		Settings: settings,
	}

	return test
}

func pickAgentID() (string, error) {
	ctx := context.Background()
	client, err := NewClient()
	if err != nil {
		return "", err
	}

	getAllResp, err := client.SyntheticsAdmin.ListAgents(ctx, &syntheticspb.ListAgentsRequest{})
	if err != nil {
		return "", err
	}

	if getAllResp.GetAgents() != nil {
		for _, agent := range getAllResp.GetAgents() {
			if agent.GetType() == "private" {
				return agent.GetId(), nil
			}
		}
	}
	return "", fmt.Errorf("No private agent found: %v", err)
}
