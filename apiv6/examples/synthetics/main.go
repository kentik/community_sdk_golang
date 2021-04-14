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

	// (16.04.2021 - API call for CREATE returns http 500)
	if err = runCRUDTest(client); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = runListTests(client); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = runSetTestStatus(client, "3337"); err != nil {
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

	if err = runGetHealthForTests(client, []string{"3337"}); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = runGetTraceForTest(client, "3337"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runCRUDTest(client *kentikapi.Client) error {
	// NOTE: CREATE doesn't work yet (16.04.2021)
	// fmt.Println("### CREATE TEST")
	// test := makeExampleTest()
	// createReqPayload := *synthetics.NewV202101beta1CreateTestRequest()
	// createReqPayload.SetTest(*test)
	// createReq := client.SyntheticsAdminServiceApi.TestCreate(context.Background()).V202101beta1CreateTestRequest(createReqPayload)
	// createResp, httpResp, err := createReq.Execute()
	// if err != nil {
	// 	return fmt.Errorf("%v %v", err, httpResp)
	// }
	// examples.PrettyPrint(createResp)
	// fmt.Println()
	// created := createResp.Test

	testID := "3337" // existing test id; temporary solution until CREATE works

	fmt.Println("### GET TEST")
	getReq := client.SyntheticsAdminServiceApi.TestGet(context.Background(), testID)
	getResp, httpResp, err := getReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	examples.PrettyPrint(getResp)
	fmt.Println()

	test := getResp.Test
	test.Settings.TargetType = nil
	test.Settings.TargetValue = nil

	fmt.Println("### PATCH TEST")
	test.SetName("[please don't delete me!]")
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

	// NOTE: uncomment once CREATE works
	// fmt.Println("### DELETE TEST")
	// deleteReq := client.SyntheticsAdminServiceApi.TestDelete(context.Background(), *test.Id)
	// deleteResp, httpResp, err := deleteReq.Execute()
	// if err != nil {
	// 	return fmt.Errorf("%v %v", err, httpResp)
	// }
	// fmt.Println("Success")
	// examples.PrettyPrint(deleteResp)
	// fmt.Println()

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
		fmt.Println("Num invalid tests:", *getAllResp.InvalidTestsCount)
		examples.PrettyPrint(tests)
	} else {
		fmt.Println("[no tests received]")
	}
	fmt.Println()

	return nil
}

func runSetTestStatus(client *kentikapi.Client, testID string) error {
	fmt.Println("### SET TEST STATUS")

	setStatusReqPayload := *synthetics.NewV202101beta1SetTestStatusRequest()
	setStatusReqPayload.Id = &testID
	status := synthetics.V202101BETA1TESTSTATUS_ACTIVE
	setStatusReqPayload.Status = &status
	setStatusReq := client.SyntheticsAdminServiceApi.TestStatusUpdate(context.Background(), testID).V202101beta1SetTestStatusRequest(setStatusReqPayload)

	getResp, httpResp, err := setStatusReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	fmt.Println("Success")
	examples.PrettyPrint(getResp)
	fmt.Println()

	return nil
}

func runGetHealthForTests(client *kentikapi.Client, testIDs []string) error {
	fmt.Println("### GET HEALTH FOR TESTS")

	healthPayload := *synthetics.NewV202101beta1GetHealthForTestsRequest()
	healthPayload.SetEndTime(time.Now())
	healthPayload.SetStartTime(time.Now().Add(-time.Hour * 24))
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
	tracePayload.SetEndTime(time.Now())
	tracePayload.SetStartTime(time.Now().Add(-time.Hour * 1))

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

// NOTE: no CREATE method exists for agents in the API - they register themselves on their own
func runCRUDAgent(client *kentikapi.Client) error {
	agentID := "1717" // private, existing agent

	fmt.Println("### GET AGENT")
	getReq := client.SyntheticsAdminServiceApi.AgentGet(context.Background(), agentID)
	getResp, httpResp, err := getReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	examples.PrettyPrint(getResp)
	fmt.Println()

	fmt.Println("### PATCH AGENT")
	agent := *synthetics.NewV202101beta1Agent()
	if getResp.Agent.GetFamily() == synthetics.V202101BETA1IPFAMILY_V6 {
		agent.SetFamily("IP_FAMILY_V4")
	} else {
		agent.SetFamily("IP_FAMILY_V6")
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
	ping.SetCount(3)
	ping.SetExpiry(3000)
	ping.SetPeriod(60)

	trace := *synthetics.NewV202101beta1TestTraceSettings()
	trace.SetCount(3)
	trace.SetExpiry(22500)
	trace.SetLimit(30)
	trace.SetPeriod(60)
	trace.SetPort(33434)
	trace.SetProtocol("udp")

	monitoring := *synthetics.NewV202101beta1TestMonitoringSettings()
	monitoring.SetActivationGracePeriod("2")
	monitoring.SetActivationTimeUnit("m")
	monitoring.SetActivationTimeWindow("5")
	monitoring.SetActivationTimes("3")
	monitoring.SetNotificationChannels([]string{})

	health := *synthetics.NewV202101beta1HealthSettings()
	health.SetDnsValidCodes([]int64{5})
	health.SetHttpLatencyCritical(500.0)
	health.SetHttpLatencyWarning(300)
	health.SetHttpValidCodes([]int64{5})
	health.SetJitterCritical(100)
	health.SetJitterWarning(50)
	health.SetLatencyCritical(500)
	health.SetLatencyWarning(300)
	health.SetPacketLossCritical(50)
	health.SetPacketLossWarning(30)

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
	settings.SetFamily(synthetics.V202101BETA1IPFAMILY_V4)
	settings.SetServers([]string{"https://fala.com"})
	// settings.SetTargetType("")
	// settings.SetTargetValue("")
	settings.SetUseLocalIp(false)
	settings.SetReciprocal(false)
	settings.SetRollupLevel(1)
	settings.SetRollupLevel(1)

	user := *synthetics.NewV202101beta1UserInfo()
	user.SetId("144319")
	user.SetFullName("Mateusz Midor")
	user.SetEmail("mateusz.midor@codilime.com")

	test := synthetics.NewV202101beta1Test()
	test.SetName("exampletest1")
	test.SetType("ip-address")
	test.SetDeviceId("1000")
	test.SetStatus(synthetics.V202101BETA1TESTSTATUS_ACTIVE)
	test.SetSettings(settings)
	test.SetExpiresOn(time.Now().Add(time.Hour * 6))
	test.SetCdate(time.Now())
	// test.SetEdate(time.Now())
	// test.SetCreatedBy(user)
	// test.SetLastUpdatedBy(user)

	return test
}
