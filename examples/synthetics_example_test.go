//go:build examples
// +build examples

package examples

import (
	"context"
	"fmt"
	syntheticspb "github.com/kentik/api-schema-public/gen/go/kentik/synthetics/v202101beta1"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/kentik/community_sdk_golang/kentikapi"
	"github.com/kentik/community_sdk_golang/kentikapi/synthetics"
)

func TestSyntheticsAPIExample(t *testing.T) {
	assert.NoError(t, runAdminServiceExamples())
	assert.NoError(t, runDataServiceExamples())
}

func runAdminServiceExamples() error {
	client, err := NewClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

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
	client, err := NewClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	// NOTE: test of id 3541 must exist
	if err = runGetHealthForTests(ctx, client, []string{"3541"}); err != nil {
		fmt.Println(err)
		return err
	}

	// NOTE: test of id 3541 must exist
	if err = runGetTraceForTest(ctx, client, "3541"); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func runCRUDTest(ctx context.Context, client *kentikapi.Client) error {
	fmt.Println("### CREATE TEST")
	test := makeExampleTest()
	createReqPayload := *synthetics.NewV202101beta1CreateTestRequest()
	createReqPayload.SetTest(*test)

	createResp, httpResp, err := client.SyntheticsAdminServiceAPI.
		TestCreate(ctx).
		Body(createReqPayload).
		Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}

	PrettyPrint(createResp)
	fmt.Println()
	testID := *createResp.Test.Id

	fmt.Println("### SET TEST STATUS")
	setStatusReqPayload := *synthetics.NewV202101beta1SetTestStatusRequest()
	status := synthetics.V202101BETA1TESTSTATUS_PAUSED
	setStatusReqPayload.Status = &status
	setStatusReqPayload.Id = &testID

	statusResp, httpResp, err := client.SyntheticsAdminServiceAPI.
		TestStatusUpdate(ctx, testID).
		Body(setStatusReqPayload).
		Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	fmt.Println("Success")
	PrettyPrint(statusResp)
	fmt.Println()

	fmt.Println("### GET TEST")
	getReq := client.SyntheticsAdminServiceAPI.TestGet(ctx, testID)
	getResp, httpResp, err := getReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	PrettyPrint(getResp)
	fmt.Println()

	test = getResp.Test

	fmt.Println("### PATCH TEST")
	test.SetName("example-test-1 UPDATED")
	patchReqPayload := *synthetics.NewV202101beta1PatchTestRequest()
	patchReqPayload.SetTest(*test)
	patchReqPayload.SetMask("test.name")

	patchResp, httpResp, err := client.SyntheticsAdminServiceAPI.
		TestPatch(ctx, *test.Id).
		Body(patchReqPayload).
		Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	PrettyPrint(patchResp)
	fmt.Println()

	fmt.Println("### DELETE TEST")
	deleteReq := client.SyntheticsAdminServiceAPI.TestDelete(ctx, testID)
	deleteResp, httpResp, err := deleteReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	fmt.Println("Success")
	PrettyPrint(deleteResp)
	fmt.Println()

	return nil
}

func runListTests(ctx context.Context, client *kentikapi.Client) error {
	fmt.Println("### LIST TESTS")

	getAllReq := client.SyntheticsAdminServiceAPI.TestsList(ctx)
	getAllResp, httpResp, err := getAllReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}

	if getAllResp.Tests != nil {
		tests := *getAllResp.Tests
		fmt.Println("Num tests:", len(tests))
		if getAllResp.InvalidTestsCount != nil {
			fmt.Println("Num invalid tests:", *getAllResp.InvalidTestsCount)
		}
		PrettyPrint(tests)
	} else {
		fmt.Println("[no tests received]")
	}
	fmt.Println()

	return nil
}

func runGetHealthForTests(ctx context.Context, client *kentikapi.Client, testIDs []string) error {
	fmt.Println("### GET HEALTH FOR TESTS")

	healthPayload := *synthetics.NewV202101beta1GetHealthForTestsRequest()
	healthPayload.SetStartTime(time.Now().Add(-time.Hour))
	healthPayload.SetEndTime(time.Now())
	healthPayload.SetIds(testIDs)

	getHealthResp, httpResp, err := client.SyntheticsDataServiceAPI.
		GetHealthForTests(ctx).
		Body(healthPayload).
		Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}

	if getHealthResp.Health != nil {
		healthItems := *getHealthResp.Health
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

	tracePayload := *synthetics.NewV202101beta1GetTraceForTestRequest()
	tracePayload.SetId(testID)
	tracePayload.SetStartTime(time.Now().Add(-time.Hour))
	tracePayload.SetEndTime(time.Now())

	getTraceResp, httpResp, err := client.SyntheticsDataServiceAPI.
		GetTraceForTest(ctx, testID).
		Body(tracePayload).
		Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}

	if getTraceResp.Lookups != nil {
		lookups := *getTraceResp.Lookups
		fmt.Println("Agents id by ip:")
		PrettyPrint(lookups)
	} else {
		fmt.Println("[no agents received]")
	}

	if getTraceResp.TraceRoutes != nil {
		results := *getTraceResp.TraceRoutes
		fmt.Println("Num trace routes:", len(results))
		PrettyPrint(results)
	} else {
		fmt.Println("[no trace routes received]")
	}

	if getTraceResp.TraceRoutesInfo != nil {
		results := *getTraceResp.TraceRoutesInfo
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
	agentID := "1717"

	fmt.Println("### GET AGENT")
	getReq := client.SyntheticsAdminServiceAPI.AgentGet(ctx, agentID)
	getResp, httpResp, err := getReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	PrettyPrint(getResp)
	fmt.Println()

	//TODO: PATCH AGENT is not working properly. request.agent.name should be ignored given mask but isn't
	//fmt.Println("### PATCH AGENT")
	//agent := *getResp.Agent
	//if getResp.Agent.GetFamily() == synthetics.V202101BETA1IPFAMILY_V6 {
	//	agent.SetFamily(synthetics.V202101BETA1IPFAMILY_V4)
	//} else {
	//	agent.SetFamily(synthetics.V202101BETA1IPFAMILY_V6)
	//}
	//patchReqPayload := *synthetics.NewV202101beta1PatchAgentRequest()
	//patchReqPayload.SetAgent(agent)
	//patchReqPayload.SetMask("agent.family")
	//
	//patchResp, httpResp, err := client.SyntheticsAdminServiceAPI.
	//	AgentPatch(context.Background(), agentID).
	//	Body(patchReqPayload).
	//	Execute()
	//if err != nil {
	//	return fmt.Errorf("%v %v", err, httpResp)
	//}
	//PrettyPrint(patchResp)
	//fmt.Println()

	// NOTE: as we can't create agents through the API - let's not delete them
	// fmt.Println("### DELETE AGENT")
	// deleteReq := client.SyntheticsAdminServiceAPI.AgentDelete(context.Background(), agentID)
	// deleteResp, httpResp, err := deleteReq.Execute()
	// if err != nil {
	// 	return fmt.Errorf("%v %v", err, httpResp)
	// }
	// fmt.Println("Success")
	// examples.PrettyPrint(deleteResp)
	// fmt.Println()

	return nil
}

func runListAgents(ctx context.Context, client *kentikapi.Client) error {
	fmt.Println("### LIST AGENTS")

	getAllReq := client.SyntheticsAdminServiceAPI.AgentsList(ctx)
	getAllResp, httpResp, err := getAllReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}

	if getAllResp.Agents != nil {
		agents := *getAllResp.Agents
		fmt.Println("Num agents:", len(agents))
		if getAllResp.InvalidAgentsCount != nil {
			fmt.Println("Num invalid agents:", *getAllResp.InvalidAgentsCount)
		}
		PrettyPrint(agents)
	} else {
		fmt.Println("[no agents received]")
	}
	fmt.Println()

	return nil
}

// prepare a Test for sending in CREATE request
func makeExampleTest() *synthetics.V202101beta1Test {
	hostname := synthetics.NewV202101beta1HostnameTest()
	hostname.SetTarget("dummy-ht.com")

	ping := *synthetics.NewV202101beta1TestPingSettings()
	ping.SetPeriod(60)
	ping.SetCount(5)
	ping.SetExpiry(3000)

	trace := *synthetics.NewV202101beta1TestTraceSettings()
	trace.SetPeriod(60)
	trace.SetCount(3)
	trace.SetProtocol("udp")
	trace.SetPort(33434)
	trace.SetExpiry(22500)
	trace.SetLimit(30)

	monitoring := *synthetics.NewV202101beta1TestMonitoringSettings()
	monitoring.SetActivationGracePeriod("2")
	monitoring.SetActivationTimeUnit("m")
	monitoring.SetActivationTimeWindow("5")
	monitoring.SetActivationTimes("3")
	monitoring.SetNotificationChannels([]string{})

	health := *synthetics.NewV202101beta1HealthSettings()
	health.SetDnsValidCodes([]int64{})
	health.SetHttpLatencyCritical(0)
	health.SetHttpLatencyWarning(0)
	health.SetHttpValidCodes([]int64{})
	health.SetJitterCritical(0)
	health.SetJitterWarning(0)
	health.SetLatencyCritical(0)
	health.SetLatencyWarning(0)
	health.SetPacketLossCritical(0)
	health.SetPacketLossWarning(0)

	settings := *synthetics.NewV202101beta1TestSettingsWithDefaults()
	settings.SetHostname(*hostname)
	settings.SetAgentIds([]string{"890"})
	settings.SetPeriod(0)
	settings.SetCount(0)
	settings.SetExpiry(0)
	settings.SetLimit(0)
	settings.SetTasks([]string{"ping", "traceroute"})
	settings.SetHealthSettings(health)
	settings.SetMonitoringSettings(monitoring)
	settings.SetPing(ping)
	settings.SetTrace(trace)
	settings.SetPort(443)
	settings.SetProtocol("icmp")
	settings.SetFamily(synthetics.V202101BETA1IPFAMILY_DUAL)
	settings.SetServers([]string{})
	settings.SetUseLocalIp(false)
	settings.SetReciprocal(false)
	settings.SetRollupLevel(1)

	user := *synthetics.NewV202101beta1UserInfo()
	user.SetId("144566")
	user.SetFullName("John Doe")
	user.SetEmail("john.doe@acme.com")

	test := synthetics.NewV202101beta1Test()
	test.SetName("example-test-1")
	test.SetType("hostname")
	test.SetDeviceId("1000")
	test.SetStatus(synthetics.V202101BETA1TESTSTATUS_ACTIVE)
	test.SetSettings(settings)
	test.SetExpiresOn(time.Now().Add(time.Hour * 6))
	test.SetCdate(time.Now())
	test.SetCreatedBy(user)

	return test
}

func TestSyntheticsAPIGRPCExample(t *testing.T) {
	assert.NoError(t, runGRPCAdminServiceExamples())
	assert.NoError(t, runGRPCDataServiceExamples())
}

func runGRPCAdminServiceExamples() error {
	client, err := NewClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	if err = runGRPCCRUDTest(ctx, client); err != nil {
		fmt.Println(err)
		return err
	}

	if err = runGRPCListTests(ctx, client); err != nil {
		fmt.Println(err)
		return err
	}

	if err = runGRPCCRUDAgent(ctx, client); err != nil {
		fmt.Println(err)
		return err
	}

	if err = runGRPCListAgents(ctx, client); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func runGRPCDataServiceExamples() error {
	client, err := NewClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	// NOTE: test of id 3541 must exist
	if err = runGRPCGetHealthForTests(ctx, client, []string{"3541"}); err != nil {
		fmt.Println(err)
		return err
	}

	// NOTE: test of id 3541 must exist
	if err = runGRPCGetTraceForTest(ctx, client, "3541"); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func runGRPCCRUDTest(ctx context.Context, client *kentikapi.Client) error {
	fmt.Println("### CREATE TEST")
	test := makeGRPCExampleTest()
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

func runGRPCListTests(ctx context.Context, client *kentikapi.Client) error {
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

func runGRPCGetHealthForTests(ctx context.Context, client *kentikapi.Client, testIDs []string) error {
	fmt.Println("### GET HEALTH FOR TESTS")

	healthPayLoad := &syntheticspb.GetHealthForTestsRequest{
		Ids: testIDs,
		StartTime: timestamppb.New(time.Now().Add(-time.Hour)),
		EndTime: timestamppb.Now(),
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

func runGRPCGetTraceForTest(ctx context.Context, client *kentikapi.Client, testID string) error {
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

func runGRPCCRUDAgent(ctx context.Context, client *kentikapi.Client) error {
	// NOTE: no CREATE method exists for agents in the API, thus no example for CREATE
	// NOTE: agent of id 1717 must exist
	agentID := "1717"

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
	if agent.GetFamily() == syntheticspb.IPFamily_IP_FAMILY_V6 {
		agent.Family = syntheticspb.IPFamily_IP_FAMILY_V4
	} else {
		agent.Family = syntheticspb.IPFamily_IP_FAMILY_V6
	}
	agent.Name = ""
	patchReqPayload := &syntheticspb.PatchAgentRequest{
		Agent: agent,
		Mask:  &fieldmaskpb.FieldMask{Paths: []string{"agent.family"}},
	}
	patchResp, err := client.SyntheticsAdmin.PatchAgent(ctx, patchReqPayload)
	if err != nil {
		return err
	}
	PrettyPrint(patchResp.GetAgent())
	fmt.Println()

	// NOTE: as we can't create agents through the API - let's not delete them
	//fmt.Println("### DELETE AGENT")
	//deleteReqPayLoad := &syntheticspb.DeleteAgentRequest{Id: agentID}
	//_, err = client.SyntheticsAdmin.DeleteAgent(ctx, deleteReqPayLoad)
	//if err != nil {
	//	return err
	//}
	//fmt.Println("Success")
	//fmt.Println()

	return nil
}

func runGRPCListAgents(ctx context.Context, client *kentikapi.Client) error {
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

func makeGRPCExampleTest() *syntheticspb.Test {
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
		Count:    3,
		Protocol: "udp",
		Port:     33434,
		Expiry:   22500,
		Limit:    30,
	}

	settings := &syntheticspb.TestSettings{
		Definition: &syntheticspb.TestSettings_Hostname{
			Hostname: &syntheticspb.HostnameTest{Target: "dummy-ht.com"},
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
		Ping: ping,
		Trace: trace,
		Port:        443,
		Protocol:    "icmp",
		Family:      syntheticspb.IPFamily_IP_FAMILY_DUAL,
		Servers:     []string{},
		UseLocalIp:  false,
		Reciprocal:  false,
		RollupLevel: 1,
		Http:        nil,
	}

	userInfo := &syntheticspb.UserInfo{
		Id:       "144566",
		Email:    "John Doe",
		FullName: "john.doe@acme.com",
	}

	test := &syntheticspb.Test{
		Name:          "example-test-1",
		Type:          "hostname",
		DeviceId:      "1000",
		Status:        syntheticspb.TestStatus_TEST_STATUS_ACTIVE,
		Settings:      settings,
		ExpiresOn:     timestamppb.New(time.Now().Add(time.Hour*6)),
		Cdate:         timestamppb.Now(),
		CreatedBy:     userInfo,
	}

	return test
}
