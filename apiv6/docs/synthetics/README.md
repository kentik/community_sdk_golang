# Documentation for Synthetics Monitoring API

<a name="documentation-for-api-endpoints"></a>
## Documentation for API Endpoints

All URIs are relative to *http://localhost:8080*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*SyntheticsAdminServiceApi* | [**agentCreate**](Apis/SyntheticsAdminServiceApi.md#agentcreate) | **POST** /synthetics/v202101beta1/agents | Create Agent.
*SyntheticsAdminServiceApi* | [**agentDelete**](Apis/SyntheticsAdminServiceApi.md#agentdelete) | **DELETE** /synthetics/v202101beta1/agents/{agent.id} | Delete an agent.
*SyntheticsAdminServiceApi* | [**agentGet**](Apis/SyntheticsAdminServiceApi.md#agentget) | **GET** /synthetics/v202101beta1/agents/{agent.id} | Get information about an agent.
*SyntheticsAdminServiceApi* | [**agentPatch**](Apis/SyntheticsAdminServiceApi.md#agentpatch) | **PATCH** /synthetics/v202101beta1/agents/{agent.id} | Patch an agent.
*SyntheticsAdminServiceApi* | [**agentsList**](Apis/SyntheticsAdminServiceApi.md#agentslist) | **GET** /synthetics/v202101beta1/agents | List Agents.
*SyntheticsAdminServiceApi* | [**exportPatch**](Apis/SyntheticsAdminServiceApi.md#exportpatch) | **PATCH** /synthetics/v202101beta1/tests/{id} | Patch a Synthetics Test.
*SyntheticsAdminServiceApi* | [**testCreate**](Apis/SyntheticsAdminServiceApi.md#testcreate) | **POST** /synthetics/v202101beta1/tests | Create Synthetics Test.
*SyntheticsAdminServiceApi* | [**testDelete**](Apis/SyntheticsAdminServiceApi.md#testdelete) | **DELETE** /synthetics/v202101beta1/tests/{id} | Delete an Synthetics Test.
*SyntheticsAdminServiceApi* | [**testGet**](Apis/SyntheticsAdminServiceApi.md#testget) | **GET** /synthetics/v202101beta1/tests/{id} | Get information about Synthetics Test.
*SyntheticsAdminServiceApi* | [**testStatusUpdate**](Apis/SyntheticsAdminServiceApi.md#teststatusupdate) | **PUT** /synthetics/v202101beta1/tests/{id}/status | Update a test status.
*SyntheticsAdminServiceApi* | [**testsList**](Apis/SyntheticsAdminServiceApi.md#testslist) | **GET** /synthetics/v202101beta1/tests | List Synthetics Tests.
*SyntheticsDataServiceApi* | [**getHealthForTests**](Apis/SyntheticsDataServiceApi.md#gethealthfortests) | **POST** /synthetics/v202101beta1/health/tests | Get health status for synthetics test.
*SyntheticsDataServiceApi* | [**getTraceForTest**](Apis/SyntheticsDataServiceApi.md#gettracefortest) | **POST** /synthetics/v202101beta1/tests/{id}/results/trace | Get trace route data.


<a name="documentation-for-models"></a>
## Documentation for Models

 - [ProtobufAny](./Models/ProtobufAny.md)
 - [RpcStatus](./Models/RpcStatus.md)
 - [V202101beta1ASN](./Models/V202101beta1ASN.md)
 - [V202101beta1Agent](./Models/V202101beta1Agent.md)
 - [V202101beta1AgentHealth](./Models/V202101beta1AgentHealth.md)
 - [V202101beta1AgentStatus](./Models/V202101beta1AgentStatus.md)
 - [V202101beta1AgentTaskConfig](./Models/V202101beta1AgentTaskConfig.md)
 - [V202101beta1AgentTest](./Models/V202101beta1AgentTest.md)
 - [V202101beta1City](./Models/V202101beta1City.md)
 - [V202101beta1Country](./Models/V202101beta1Country.md)
 - [V202101beta1CreateTestRequest](./Models/V202101beta1CreateTestRequest.md)
 - [V202101beta1CreateTestResponse](./Models/V202101beta1CreateTestResponse.md)
 - [V202101beta1DNS](./Models/V202101beta1DNS.md)
 - [V202101beta1DNSTaskDefinition](./Models/V202101beta1DNSTaskDefinition.md)
 - [V202101beta1DnsTest](./Models/V202101beta1DnsTest.md)
 - [V202101beta1FlowTest](./Models/V202101beta1FlowTest.md)
 - [V202101beta1Geo](./Models/V202101beta1Geo.md)
 - [V202101beta1GetAgentResponse](./Models/V202101beta1GetAgentResponse.md)
 - [V202101beta1GetHealthForTestsRequest](./Models/V202101beta1GetHealthForTestsRequest.md)
 - [V202101beta1GetHealthForTestsResponse](./Models/V202101beta1GetHealthForTestsResponse.md)
 - [V202101beta1GetTestResponse](./Models/V202101beta1GetTestResponse.md)
 - [V202101beta1GetTraceForTestRequest](./Models/V202101beta1GetTraceForTestRequest.md)
 - [V202101beta1GetTraceForTestResponse](./Models/V202101beta1GetTraceForTestResponse.md)
 - [V202101beta1HTTPTaskDefinition](./Models/V202101beta1HTTPTaskDefinition.md)
 - [V202101beta1Health](./Models/V202101beta1Health.md)
 - [V202101beta1HealthMoment](./Models/V202101beta1HealthMoment.md)
 - [V202101beta1HealthSettings](./Models/V202101beta1HealthSettings.md)
 - [V202101beta1HostnameTest](./Models/V202101beta1HostnameTest.md)
 - [V202101beta1IPFamily](./Models/V202101beta1IPFamily.md)
 - [V202101beta1IPInfo](./Models/V202101beta1IPInfo.md)
 - [V202101beta1IpTest](./Models/V202101beta1IpTest.md)
 - [V202101beta1KnockTaskDefinition](./Models/V202101beta1KnockTaskDefinition.md)
 - [V202101beta1ListAgentsResponse](./Models/V202101beta1ListAgentsResponse.md)
 - [V202101beta1ListTestsResponse](./Models/V202101beta1ListTestsResponse.md)
 - [V202101beta1MeshColumn](./Models/V202101beta1MeshColumn.md)
 - [V202101beta1MeshMetric](./Models/V202101beta1MeshMetric.md)
 - [V202101beta1MeshMetrics](./Models/V202101beta1MeshMetrics.md)
 - [V202101beta1MeshResponse](./Models/V202101beta1MeshResponse.md)
 - [V202101beta1PatchAgentRequest](./Models/V202101beta1PatchAgentRequest.md)
 - [V202101beta1PatchAgentResponse](./Models/V202101beta1PatchAgentResponse.md)
 - [V202101beta1PatchTestResponse](./Models/V202101beta1PatchTestResponse.md)
 - [V202101beta1PingTaskDefinition](./Models/V202101beta1PingTaskDefinition.md)
 - [V202101beta1Region](./Models/V202101beta1Region.md)
 - [V202101beta1SetTestStatusRequest](./Models/V202101beta1SetTestStatusRequest.md)
 - [V202101beta1ShakeTaskDefinition](./Models/V202101beta1ShakeTaskDefinition.md)
 - [V202101beta1SiteTest](./Models/V202101beta1SiteTest.md)
 - [V202101beta1TagTest](./Models/V202101beta1TagTest.md)
 - [V202101beta1Task](./Models/V202101beta1Task.md)
 - [V202101beta1TaskHealth](./Models/V202101beta1TaskHealth.md)
 - [V202101beta1TaskState](./Models/V202101beta1TaskState.md)
 - [V202101beta1Test](./Models/V202101beta1Test.md)
 - [V202101beta1TestHealth](./Models/V202101beta1TestHealth.md)
 - [V202101beta1TestMonitoringSettings](./Models/V202101beta1TestMonitoringSettings.md)
 - [V202101beta1TestPingSettings](./Models/V202101beta1TestPingSettings.md)
 - [V202101beta1TestSettings](./Models/V202101beta1TestSettings.md)
 - [V202101beta1TestStatus](./Models/V202101beta1TestStatus.md)
 - [V202101beta1TestTraceSettings](./Models/V202101beta1TestTraceSettings.md)
 - [V202101beta1Trace](./Models/V202101beta1Trace.md)
 - [V202101beta1TraceHop](./Models/V202101beta1TraceHop.md)
 - [V202101beta1TraceTaskDefinition](./Models/V202101beta1TraceTaskDefinition.md)
 - [V202101beta1TracerouteResult](./Models/V202101beta1TracerouteResult.md)
 - [V202101beta1UrlTest](./Models/V202101beta1UrlTest.md)
 - [V202101beta1UserInfo](./Models/V202101beta1UserInfo.md)


<a name="documentation-for-authorization"></a>
## Documentation for Authorization

<a name="email"></a>
### email

- **Type**: API key
- **API key parameter name**: X-CH-Auth-Email
- **Location**: HTTP header

<a name="token"></a>
### token

- **Type**: API key
- **API key parameter name**: X-CH-Auth-API-Token
- **Location**: HTTP header

