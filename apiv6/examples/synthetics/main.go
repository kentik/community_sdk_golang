package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/kentik/community_sdk_golang/apiv6/examples"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/synthetics"
)

func main() {
	runAdminServiceExamples()
	runDataServiceExamples()
}

func runAdminServiceExamples() {
	client := examples.NewClient()
	var err error

	if err = runCRUDTest(client); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = runListTests(client); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = runCRUDAgent(client); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = runListAgents(client); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runDataServiceExamples() {
	client := examples.NewClient()
	var err error

	// NOTE: test of id 3336 must exist
	if err = runGetHealthForTests(client, []string{"3336"}); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// NOTE: test of id 3336 must exist
	if err = runGetTraceForTest(client, "3336"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runCRUDTest(client *kentikapi.Client) error {
	fmt.Println("### CREATE TEST")
	test := makeExampleTest()
	createReqPayload := *synthetics.NewV202101beta1CreateTestRequest()
	createReqPayload.SetTest(*test)
	createReq := client.SyntheticsAdminServiceApi.TestCreate(context.Background()).V202101beta1CreateTestRequest(createReqPayload)
	createResp, httpResp, err := createReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	examples.PrettyPrint(createResp)
	fmt.Println()
	testID := *createResp.Test.Id

	fmt.Println("### SET TEST STATUS")
	setStatusReqPayload := *synthetics.NewV202101beta1SetTestStatusRequest()
	status := synthetics.V202101BETA1TESTSTATUS_PAUSED
	setStatusReqPayload.Status = &status
	setStatusReqPayload.Id = &testID
	setStatusReq := client.SyntheticsAdminServiceApi.TestStatusUpdate(context.Background(), testID).V202101beta1SetTestStatusRequest(setStatusReqPayload)
	statusResp, httpResp, err := setStatusReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	fmt.Println("Success")
	examples.PrettyPrint(statusResp)
	fmt.Println()

	fmt.Println("### GET TEST")
	getReq := client.SyntheticsAdminServiceApi.TestGet(context.Background(), testID)
	getResp, httpResp, err := getReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	examples.PrettyPrint(getResp)
	fmt.Println()

	test = getResp.Test
	test.Settings.TargetType = nil
	test.Settings.TargetValue = nil

	fmt.Println("### PATCH TEST")
	test.SetName("example-test-1 UPDATED")
	patchReqPayload := *synthetics.NewV202101beta1PatchTestRequest()
	patchReqPayload.SetTest(*test)
	patchReqPayload.SetMask("test.name")
	patchReq := client.SyntheticsAdminServiceApi.TestPatch(context.Background(), *test.Id).V202101beta1PatchTestRequest(patchReqPayload)
	patchResp, httpResp, err := patchReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	examples.PrettyPrint(patchResp)
	fmt.Println()

	fmt.Println("### DELETE TEST")
	deleteReq := client.SyntheticsAdminServiceApi.TestDelete(context.Background(), testID)
	deleteResp, httpResp, err := deleteReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	fmt.Println("Success")
	examples.PrettyPrint(deleteResp)
	fmt.Println()

	return nil
}

func runListTests(client *kentikapi.Client) error {
	fmt.Println("### LIST TESTS")

	getAllReq := client.SyntheticsAdminServiceApi.TestsList(context.Background())
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
		examples.PrettyPrint(tests)
	} else {
		fmt.Println("[no tests received]")
	}
	fmt.Println()

	return nil
}

func runGetHealthForTests(client *kentikapi.Client, testIDs []string) error {
	fmt.Println("### GET HEALTH FOR TESTS")

	healthPayload := *synthetics.NewV202101beta1GetHealthForTestsRequest()
	healthPayload.SetStartTime(time.Now().Add(-time.Hour))
	healthPayload.SetEndTime(time.Now())
	healthPayload.SetIds(testIDs)
	getHealthReq := client.SyntheticsDataServiceApi.GetHealthForTests(context.Background()).V202101beta1GetHealthForTestsRequest(healthPayload)
	getHealthResp, httpResp, err := getHealthReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}

	if getHealthResp.Health != nil {
		healthItems := *getHealthResp.Health
		fmt.Println("Num health items:", len(healthItems))
		examples.PrettyPrint(healthItems)
	} else {
		fmt.Println("[no health items received]")
	}
	fmt.Println()

	return nil
}

func runGetTraceForTest(client *kentikapi.Client, testID string) error {
	fmt.Println("### GET TRACE FOR TEST")

	tracePayload := *synthetics.NewV202101beta1GetTraceForTestRequest()
	tracePayload.SetId(testID)
	tracePayload.SetStartTime(time.Now().Add(-time.Hour))
	tracePayload.SetEndTime(time.Now())
	getTraceReq := client.SyntheticsDataServiceApi.GetTraceForTest(context.Background(), testID).V202101beta1GetTraceForTestRequest(tracePayload)
	getTraceResp, httpResp, err := getTraceReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}

	if getTraceResp.IpInfo != nil {
		ipItems := *getTraceResp.IpInfo
		fmt.Println("Num ip items:", len(ipItems))
		examples.PrettyPrint(ipItems)
	} else {
		fmt.Println("[no ip items received]")
	}

	if getTraceResp.TraceRoutes != nil {
		results := *getTraceResp.TraceRoutes
		fmt.Println("Num trace routes:", len(results))
		examples.PrettyPrint(results)
	} else {
		fmt.Println("[no trace routes received]")
	}

	fmt.Println()

	return nil
}

func runCRUDAgent(client *kentikapi.Client) error {
	// NOTE: no CREATE method exists for agents in the API, thus no example for CREATE
	// NOTE: agent of id 1717 must exist
	agentID := "1717"

	fmt.Println("### GET AGENT")
	getReq := client.SyntheticsAdminServiceApi.AgentGet(context.Background(), agentID)
	getResp, httpResp, err := getReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	examples.PrettyPrint(getResp)
	fmt.Println()

	fmt.Println("### PATCH AGENT")
	agent := *getResp.Agent
	if getResp.Agent.GetFamily() == synthetics.V202101BETA1IPFAMILY_V6 {
		agent.SetFamily(synthetics.V202101BETA1IPFAMILY_V4)
	} else {
		agent.SetFamily(synthetics.V202101BETA1IPFAMILY_V6)
	}
	patchReqPayload := *synthetics.NewV202101beta1PatchAgentRequest()
	patchReqPayload.SetAgent(agent)
	patchReqPayload.SetMask("agent.family")
	patchReq := client.SyntheticsAdminServiceApi.AgentPatch(context.Background(), agentID).V202101beta1PatchAgentRequest(patchReqPayload)
	patchResp, httpResp, err := patchReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	examples.PrettyPrint(patchResp)
	fmt.Println()

	// NOTE: as we can't create agents through the API - let's not delete them
	// fmt.Println("### DELETE AGENT")
	// deleteReq := client.SyntheticsAdminServiceApi.AgentDelete(context.Background(), agentID)
	// deleteResp, httpResp, err := deleteReq.Execute()
	// if err != nil {
	// 	return fmt.Errorf("%v %v", err, httpResp)
	// }
	// fmt.Println("Success")
	// examples.PrettyPrint(deleteResp)
	// fmt.Println()

	return nil
}

func runListAgents(client *kentikapi.Client) error {
	fmt.Println("### LIST AGENTS")

	getAllReq := client.SyntheticsAdminServiceApi.AgentsList(context.Background())
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
		examples.PrettyPrint(agents)
	} else {
		fmt.Println("[no agents received]")
	}
	fmt.Println()

	return nil
}

// prepare a Test for sending in CREATE request
func makeExampleTest() *synthetics.V202101beta1Test {
	ipSetting := *synthetics.NewV202101beta1IpTest()
	ipSetting.SetTargets([]string{"127.0.0.1"})

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
	settings.SetIp(ipSetting)
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
	test.SetType("ip-address")
	test.SetDeviceId("1000")
	test.SetStatus(synthetics.V202101BETA1TESTSTATUS_ACTIVE)
	test.SetSettings(settings)
	test.SetExpiresOn(time.Now().Add(time.Hour * 6))
	test.SetCdate(time.Now())
	test.SetCreatedBy(user)

	return test
}
