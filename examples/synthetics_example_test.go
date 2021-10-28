//+build examples

package examples

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi"
	"github.com/kentik/community_sdk_golang/kentikapi/synthetics"
)

func TestSyntheticsAPIExample(t *testing.T) {
	assert.NoError(t, runAdminServiceExamples())
	assert.NoError(t, runDataServiceExamples())
}

func runAdminServiceExamples() error {
	client := NewClient()
	var err error

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
	client := NewClient()
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	// NOTE: test of id 3336 must exist
	if err = runGetHealthForTests(ctx, client, []string{"3336"}); err != nil {
		fmt.Println(err)
		return err
	}

	// NOTE: test of id 3336 must exist
	if err = runGetTraceForTest(ctx, client, "3336"); err != nil {
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
	hostname.SetTarget("dummy-ht")

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
