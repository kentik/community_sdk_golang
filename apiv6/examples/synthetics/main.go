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

	// if err = runCRUDTest(client); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// if err = runGetAllTests(client); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// if err = runGetTest(client, "3337"); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// if err = runPatchTest(client, "3334"); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// if err = runDeleteTest(client, "3334"); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// if err = runSetTestStatus(client, "3337"); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// private agent ids: 1717 and 1729
	// if err = runCRUDAgent(client); err != nil {
	// 		fmt.Println(err)
	// 		os.Exit(1)
	// }

	// if err = runGetAllAgents(client); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// if err = runGetAgent(client, "1717"); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	if err = runPatchAgent(client, "1717"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// if err = runDeleteAgent(client, "9999s"); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
}

func runDataServiceExamples() {
	client := examples.NewClient()
	var err error

	if err = runGetHealthForTests(client, []string{"3337"}); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// if err = runGetTraceForTest(client, "3337"); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
}

// NOTE: CREATE returns http 500
func runCRUDTest(client *kentikapi.Client) error {
	fmt.Println("### CREATE")
	test := makeExampleTest()
	createReqPayload := *synthetics.NewV202101beta1CreateTestRequest()
	createReqPayload.Test = test
	createReq := client.SyntheticsAdminServiceApi.SyntheticsAdminServiceCreateTest(context.Background()).V202101beta1CreateTestRequest(createReqPayload)
	createResp, httpResp, err := createReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	examples.PrettyPrint(createResp)
	fmt.Println()
	created := *createResp.Test

	// NOTE: PATCH is missing the request payload in swagger spec. Doesnt compile.
	// fmt.Println("### PATCH")
	// created.SetStatus(synthetics.V202101BETA1TESTSTATUS_PAUSED) // change from ACTIVE
	// patchReqPayload := *synthetics.NewV202101beta1PatchTestRequest()
	// patchReqPayload.Test = created
	// patchReq := client.SyntheticsAdminServiceApi.SyntheticsAdminServicePatchTest(context.Background(), *created.Id).V202101beta1PatchTestRequest(patchReqPayload)
	// patchResp, httpResp, err := patchReq.Execute()
	// if err != nil {
	// 	return fmt.Errorf("%v %v", err, httpResp)
	// }
	// examples.PrettyPrint(patchResp)
	// fmt.Println()

	fmt.Println("### GET")
	getReq := client.SyntheticsAdminServiceApi.SyntheticsAdminServiceGetTest(context.Background(), *created.Id)
	getResp, httpResp, err := getReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	examples.PrettyPrint(getResp)
	fmt.Println()

	fmt.Println("### DELETE")
	deleteReq := client.SyntheticsAdminServiceApi.SyntheticsAdminServiceDeleteTest(context.Background(), *created.Id)
	deleteResp, httpResp, err := deleteReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	fmt.Println("Success")
	examples.PrettyPrint(deleteResp)
	fmt.Println()

	return nil
}

// NOTE: GET ALL tests works
func runGetAllTests(client *kentikapi.Client) error {
	fmt.Println("### GET ALL TESTS")

	getAllReq := client.SyntheticsAdminServiceApi.SyntheticsAdminServiceListTests(context.Background())
	getAllResp, httpResp, err := getAllReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	tests := *getAllResp.Tests
	fmt.Println("Num tests:", len(tests))
	fmt.Println("Num invalid tests:", *getAllResp.InvalidTestsCount)
	examples.PrettyPrint(tests)
	fmt.Println()

	return nil
}

// NOTE: GET test works
func runGetTest(client *kentikapi.Client, testID string) error {
	fmt.Println("### GET TEST")

	getReq := client.SyntheticsAdminServiceApi.SyntheticsAdminServiceGetTest(context.Background(), testID)
	getResp, httpResp, err := getReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	examples.PrettyPrint(getResp)
	fmt.Println()

	return nil
}

// NOTE: PATCH test is misformed; missing request payload
func runPatchTest(client *kentikapi.Client, testID string) error {
	fmt.Println("### PATCH TEST")

	// test := synthetics.NewV202101beta1Test()
	// test.SetId("3337")
	// test.SetName("failing codilime ip test - updated")
	// patchReqPayload := synthetics.NewV202101beta1PatchTestRequest()
	// patchReqPayload.SetTest(test)

	// patchReq := client.SyntheticsAdminServiceApi.SyntheticsAdminServicePatchTest(context.Background(), testID).V202101beta1PatchTestRequest(patchReqPayload)
	// patchResp, httpResp, err := patchReq.Execute()
	// if err != nil {
	// 	return fmt.Errorf("%v %v", err, httpResp)
	// }
	// examples.PrettyPrint(patchResp)
	// fmt.Println()

	return nil
}

// NOTE: DELETE test works
func runDeleteTest(client *kentikapi.Client, testID string) error {
	fmt.Println("### DELETE TEST")

	deleteReq := client.SyntheticsAdminServiceApi.SyntheticsAdminServiceDeleteTest(context.Background(), testID)
	deleteResp, httpResp, err := deleteReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	fmt.Println("Success")
	examples.PrettyPrint(deleteResp)
	fmt.Println()

	return nil
}

// NOTE: SET TEST STATUS works
func runSetTestStatus(client *kentikapi.Client, testID string) error {
	fmt.Println("### SET TEST STATUS")

	setStatusReqPayload := *synthetics.NewV202101beta1SetTestStatusRequest()
	setStatusReqPayload.Id = &testID
	status := synthetics.V202101BETA1TESTSTATUS_ACTIVE
	setStatusReqPayload.Status = &status
	setStatusReq := client.SyntheticsAdminServiceApi.SyntheticsAdminServiceSetTestStatus(context.Background(), testID).V202101beta1SetTestStatusRequest(setStatusReqPayload)

	getResp, httpResp, err := setStatusReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	fmt.Println("Success")
	examples.PrettyPrint(getResp)
	fmt.Println()

	return nil
}

// NOTE: GET HEALTH FOR TESTS WORKS
func runGetHealthForTests(client *kentikapi.Client, testIDs []string) error {
	fmt.Println("### GET HEALTH FOR TESTS")

	healthPayload := *synthetics.NewV202101beta1GetHealthForTestsRequest()
	healthPayload.SetEndTime(time.Now())
	healthPayload.SetStartTime(time.Now().Add(-time.Hour * 24 * 1))
	healthPayload.SetIds(testIDs)

	getHealthReq := client.SyntheticsDataServiceApi.SyntheticsDataServiceGetHealthForTests(context.Background()).V202101beta1GetHealthForTestsRequest(healthPayload)
	getHealthResp, httpResp, err := getHealthReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	healthItems := *getHealthResp.Health
	fmt.Println("Num health items:", len(healthItems))
	examples.PrettyPrint(healthItems)
	fmt.Println()

	return nil
}

// NOTE: GET TRACE FOR TEST works
func runGetTraceForTest(client *kentikapi.Client, testID string) error {
	fmt.Println("### GET TRACE FOR TEST")

	tracePayload := *synthetics.NewV202101beta1GetTraceForTestRequest()
	tracePayload.SetId(testID)
	tracePayload.SetEndTime(time.Now())
	tracePayload.SetStartTime(time.Now().Add(-time.Hour * 24 * 1))

	getTraceReq := client.SyntheticsDataServiceApi.SyntheticsDataServiceGetTraceForTest(context.Background(), "3337").V202101beta1GetTraceForTestRequest(tracePayload)
	getTraceResp, httpResp, err := getTraceReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}

	if getTraceResp.IpInfo != nil {
		ipItems := *getTraceResp.IpInfo
		fmt.Println("Num ip items:", len(ipItems))
		examples.PrettyPrint(ipItems)
	} else {
		fmt.Println("[no ip items]")
	}

	if getTraceResp.Results != nil {
		results := *getTraceResp.Results
		fmt.Println("Num results:", len(results))
		examples.PrettyPrint(results)
	} else {
		fmt.Println("[no results]")
	}

	fmt.Println()

	return nil
}

// NOTE: CREATE not implemented; Agent API seems to be ongoing work
func runCRUDAgent(client *kentikapi.Client) error {
	// fmt.Println("### CREATE") // create not really supported in openapi spec V202101beta1
	// agent := *synthetics.NewV202101beta1Agent()
	// agent.SetName("test-agent-name")
	// createReqPayload := *synthetics.NewV202101beta1CreateAgentRequest() // no such function: NewV202101beta1CreateAgentRequest
	// createReqPayload.Agent = agent
	// createReq := client.SyntheticsAdminServiceApi.SyntheticsAdminServiceCreateAgent(context.Background()).V202101beta1CreateAgentRequest(createReqPayload)
	// createResp, httpResp, err := createReq.Execute()
	// if err != nil {
	// 	return fmt.Errorf("%v %v", err, httpResp)
	// }
	// examples.PrettyPrint(createResp)
	// fmt.Println()

	// fmt.Println("### UPDATE") // createResp is generic object so no way to extract agent ID to use in UPDATE request
	// agent.SetName("test-agent-name-updated")
	// updateReq := client.SyntheticsAdminServiceApi.SyntheticsAdminServicePatchAgent(context.Background(), *createResp.Agent.Id)
	// updateResp, httpResp, err := updateReq.Execute()
	// if err != nil {
	// 	return fmt.Errorf("%v %v", err, httpResp)
	// }
	// examples.PrettyPrint(updateResp)
	// fmt.Println()

	// fmt.Println("### GET") // createResp is generic object so no way to extract agent ID to use in GET request
	// getReq := client.SyntheticsAdminServiceApi.SyntheticsAdminServiceGetAgent(context.Background(), *createResp.Agent.Id)
	// getResp, httpResp, err := getReq.Execute()
	// if err != nil {
	// 	return fmt.Errorf("%v %v", err, httpResp)
	// }
	// examples.PrettyPrint(getResp)
	// fmt.Println()

	// fmt.Println("### DELETE") // createResp is generic object so no way to extract agent ID to use in DELETE request
	// deleteReq := client.SyntheticsAdminServiceApi.SyntheticsAdminServiceDeleteAgent(context.Background(), *createResp.Agent.Id)
	// deleteResp, httpResp, err := deleteReq.Execute()
	// if err != nil {
	// 	return fmt.Errorf("%v %v", err, httpResp)
	// }
	// fmt.Println("Success")
	// examples.PrettyPrint(deleteResp)
	// fmt.Println()

	return nil
}

// NOTE: GET ALL agents works
func runGetAllAgents(client *kentikapi.Client) error {
	fmt.Println("### GET ALL AGENTS")

	getAllReq := client.SyntheticsAdminServiceApi.SyntheticsAdminServiceListAgents(context.Background())
	getAllResp, httpResp, err := getAllReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	agents := *getAllResp.Agents
	fmt.Println("Num agents:", len(agents))
	examples.PrettyPrint(agents)
	fmt.Println()

	return nil
}

// NOTE: GET agent works
func runGetAgent(client *kentikapi.Client, agentID string) error {
	fmt.Println("### GET AGENT")

	getReq := client.SyntheticsAdminServiceApi.SyntheticsAdminServiceGetAgent(context.Background(), agentID)
	getResp, httpResp, err := getReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	examples.PrettyPrint(getResp)
	fmt.Println()

	return nil
}

// NOTE: PATCH agent returns http 500
func runPatchAgent(client *kentikapi.Client, agentID string) error {
	fmt.Println("### PATCH AGENT")

	agent := *synthetics.NewV202101beta1Agent()
	agent.SetId(agentID)
	agent.SetCity("Wonderland")

	patchReqPayload := *synthetics.NewV202101beta1PatchAgentRequest()
	patchReqPayload.SetAgent(agent)
	patchReqPayload.SetUpdateMask("")

	patchReq := client.SyntheticsAdminServiceApi.SyntheticsAdminServicePatchAgent(context.Background(), agentID).V202101beta1PatchAgentRequest(patchReqPayload)
	patchResp, httpResp, err := patchReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	examples.PrettyPrint(patchResp)
	fmt.Println()

	return nil
}

// NOTE: DELETE agent works
func runDeleteAgent(client *kentikapi.Client, agentID string) error {
	fmt.Println("### DELETE AGENT")
	deleteReq := client.SyntheticsAdminServiceApi.SyntheticsAdminServiceDeleteAgent(context.Background(), agentID)
	deleteResp, httpResp, err := deleteReq.Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	fmt.Println("Success")
	examples.PrettyPrint(deleteResp)
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
	settings.SetServers([]string{})
	settings.SetTargetType("")
	settings.SetTargetValue("")
	settings.SetUseLocalIp(false)
	settings.SetReciprocal(false)
	settings.SetRollupLevel(1)
	settings.SetRollupLevel(1)

	user := *synthetics.NewV202101beta1UserInfo()
	user.SetId("144319")
	user.SetFullName("Mateusz Midor")
	user.SetEmail("mateusz.midor@codilime.com")

	test := synthetics.NewV202101beta1Test()
	test.SetName("example-test-1")
	test.SetType("ip-address")
	test.SetDeviceId("1000")
	test.SetStatus(synthetics.V202101BETA1TESTSTATUS_UNSPECIFIED)
	test.SetSettings(settings)
	test.SetExpiresOn(time.Now().Add(time.Hour * 6))
	test.SetCdate(time.Now())
	test.SetEdate(time.Now())
	test.SetCreatedBy(user)
	test.SetLastUpdatedBy(user)

	return test
}
