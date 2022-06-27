package synthetics

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"

	syntheticspb "github.com/kentik/api-schema-public/gen/go/kentik/synthetics/v202202"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/convert"
	"github.com/kentik/community_sdk_golang/kentikapi/synthetics"
)

const (
	// testTypeBGPMonitor is hidden from SDK users (not included in public enum)
	testTypeBGPMonitor = "bgp_monitor"
)

type listTestsResponse syntheticspb.ListTestsResponse

func (r *listTestsResponse) ToModel() (*synthetics.GetAllTestsResponse, error) {
	if r == nil {
		return nil, errors.New("response payload is nil")
	}

	tests, err := testsFromPayload(r.Tests)
	if err != nil {
		return nil, err
	}

	return &synthetics.GetAllTestsResponse{
		Tests:             tests,
		InvalidTestsCount: r.InvalidCount,
	}, nil
}

func testsFromPayload(tests []*syntheticspb.Test) ([]synthetics.Test, error) {
	var result []synthetics.Test
	for i, t := range tests {
		if t.Type == testTypeBGPMonitor {
			// silently ignore BGP monitor test, as they are going to be handled in separate BGP monitoring API
			continue
		}

		test, err := testFromPayload(t)
		if err != nil {
			return nil, fmt.Errorf("test with index %v: %w", i, err)
		}
		result = append(result, *test)
	}
	return result, nil
}

// testFromPayload converts synthetics test payload to model.
func testFromPayload(t *syntheticspb.Test) (*synthetics.Test, error) {
	if t == nil {
		return nil, fmt.Errorf("response payload is nil")
	}

	if t.Id == "" {
		return nil, fmt.Errorf("empty test ID in response payload")
	}

	settings, err := testSettingsFromPayload(t)
	if err != nil {
		return nil, fmt.Errorf("convert response payload to test settings: %w", err)
	}

	createdBy := userInfoPtrFromPayload(t.CreatedBy)
	if createdBy == nil {
		return nil, errors.New("createdBy field is nil in response payload")
	}

	return &synthetics.Test{
		Name:          t.Name,
		Status:        synthetics.TestStatus(t.Status.String()),
		Settings:      settings,
		ID:            t.Id,
		Type:          synthetics.TestType(t.Type),
		CreateDate:    t.Cdate.AsTime(),
		UpdateDate:    t.Edate.AsTime(),
		CreatedBy:     *createdBy,
		LastUpdatedBy: userInfoPtrFromPayload(t.LastUpdatedBy),
	}, nil
}

func testSettingsFromPayload(t *syntheticspb.Test) (synthetics.TestSettings, error) {
	if t.Settings == nil {
		return synthetics.TestSettings{}, fmt.Errorf("empty test settings")
	}
	ts := t.Settings

	definition, err := testDefinitionFromPayload(t)
	if err != nil {
		return synthetics.TestSettings{}, err
	}

	hs, err := healthSettingsFromPayload(ts.HealthSettings)
	if err != nil {
		return synthetics.TestSettings{}, err
	}

	return synthetics.TestSettings{
		Definition:           definition,
		AgentIDs:             ts.AgentIds,
		Tasks:                taskTypesFromPayload(ts.Tasks),
		Health:               hs,
		Ping:                 pingSettingsFromPayload(ts.Ping),
		Traceroute:           tracerouteSettingsFromPayload(ts.Trace),
		Period:               time.Duration(ts.Period) * time.Second,
		Family:               synthetics.IPFamily(ts.Family.String()),
		NotificationChannels: ts.NotificationChannels,
	}, nil
}

// nolint: gocyclo
func testDefinitionFromPayload(t *syntheticspb.Test) (synthetics.TestDefinition, error) {
	switch synthetics.TestType(t.Type) {
	case synthetics.TestTypeIP:
		return ipTestDefinitionFromPayload(t.Settings)
	case synthetics.TestTypeNetworkGrid:
		return networkGridTestDefinitionFromPayload(t.Settings)
	case synthetics.TestTypeHostname:
		return hostnameTestDefinitionFromPayload(t.Settings)
	case synthetics.TestTypeAgent:
		return agentTestDefinitionFromPayload(t.Settings)
	case synthetics.TestTypeNetworkMesh:
		return networkMeshTestDefinitionFromPayload(t.Settings)
	case synthetics.TestTypeFlow:
		return flowTestDefinitionFromPayload(t.Settings)
	case synthetics.TestTypeURL:
		return urlTestDefinitionFromPayload(t.Settings)
	case synthetics.TestTypePageLoad:
		return pageLoadTestDefinitionFromPayload(t.Settings)
	case synthetics.TestTypeDNS:
		return dnsTestDefinitionFromPayload(t.Settings)
	case synthetics.TestTypeDNSGrid:
		return dnsGridTestDefinitionFromPayload(t.Settings)
	default:
		return nil, fmt.Errorf("unsupported test type: %v", t.Type)
	}
}

func ipTestDefinitionFromPayload(ts *syntheticspb.TestSettings) (synthetics.TestDefinition, error) {
	d := ts.GetIp()
	if d == nil {
		return nil, errors.New("IP test definition is nil")
	}

	targets, err := convert.StringsToIPs(d.Targets)
	if err != nil {
		return nil, fmt.Errorf("convert IP targets: %v", err)
	}

	return &synthetics.TestDefinitionIP{
		Targets: targets,
	}, nil
}

func networkGridTestDefinitionFromPayload(ts *syntheticspb.TestSettings) (synthetics.TestDefinition, error) {
	d := ts.GetNetworkGrid()
	if d == nil {
		return nil, errors.New("network grid test definition is nil")
	}

	targets, err := convert.StringsToIPs(d.Targets)
	if err != nil {
		return nil, fmt.Errorf("convert network grid targets: %v", err)
	}

	return &synthetics.TestDefinitionNetworkGrid{
		Targets: targets,
	}, nil
}

func hostnameTestDefinitionFromPayload(ts *syntheticspb.TestSettings) (synthetics.TestDefinition, error) {
	d := ts.GetHostname()
	if d == nil {
		return nil, errors.New("hostname test definition is nil")
	}

	return &synthetics.TestDefinitionHostname{
		Target: d.Target,
	}, nil
}

func agentTestDefinitionFromPayload(ts *syntheticspb.TestSettings) (synthetics.TestDefinition, error) {
	d := ts.GetAgent()
	if d == nil {
		return nil, errors.New("agent test definition is nil")
	}

	return &synthetics.TestDefinitionAgent{
		Target:     d.Target,
		UseLocalIP: d.UseLocalIp,
	}, nil
}

func networkMeshTestDefinitionFromPayload(ts *syntheticspb.TestSettings) (synthetics.TestDefinition, error) {
	d := ts.GetNetworkMesh()
	if d == nil {
		return nil, errors.New("network mesh test definition is nil")
	}

	return &synthetics.TestDefinitionNetworkMesh{
		UseLocalIP: d.UseLocalIp,
	}, nil
}

func flowTestDefinitionFromPayload(ts *syntheticspb.TestSettings) (synthetics.TestDefinition, error) {
	d := ts.GetFlow()
	if d == nil {
		return nil, errors.New("flow test definition is nil")
	}

	return &synthetics.TestDefinitionFlow{
		Type:                  synthetics.FlowTestType(d.Type),
		Target:                d.Target,
		TargetRefreshInterval: time.Duration(d.TargetRefreshIntervalMillis) * time.Millisecond,
		MaxIPTargets:          d.MaxIpTargets,
		MaxProviders:          d.MaxProviders,
		Direction:             synthetics.Direction(d.Direction),
		InetDirection:         synthetics.Direction(d.InetDirection),
	}, nil
}

func urlTestDefinitionFromPayload(ts *syntheticspb.TestSettings) (synthetics.TestDefinition, error) {
	d := ts.GetUrl()
	if d == nil {
		return nil, errors.New("URL test definition is nil")
	}

	target, err := url.Parse(d.Target)
	if err != nil {
		return nil, fmt.Errorf("parse URL test definition target: %v", err)
	}

	return &synthetics.TestDefinitionURL{
		Target:          *target,
		Timeout:         time.Duration(d.Timeout) * time.Millisecond,
		Method:          d.Method,
		Headers:         d.Headers,
		Body:            d.Body,
		IgnoreTLSErrors: d.IgnoreTlsErrors,
	}, nil
}

func pageLoadTestDefinitionFromPayload(ts *syntheticspb.TestSettings) (synthetics.TestDefinition, error) {
	d := ts.GetPageLoad()
	if d == nil {
		return nil, errors.New("page load test definition is nil")
	}

	target, err := url.Parse(d.Target)
	if err != nil {
		return nil, fmt.Errorf("parse page load test definition target: %v", err)
	}

	return &synthetics.TestDefinitionPageLoad{
		Target:          *target,
		Timeout:         time.Duration(d.Timeout) * time.Millisecond,
		Headers:         d.Headers,
		CSSSelectors:    d.CssSelectors,
		IgnoreTLSErrors: d.IgnoreTlsErrors,
	}, nil
}

func dnsTestDefinitionFromPayload(ts *syntheticspb.TestSettings) (synthetics.TestDefinition, error) {
	d := ts.GetDns()
	if d == nil {
		return nil, errors.New("DNS test definition is nil")
	}

	servers, err := convert.StringsToIPs(d.Servers)
	if err != nil {
		return nil, fmt.Errorf("convert DNS servers: %v", err)
	}

	return &synthetics.TestDefinitionDNS{
		Target:     d.Target,
		Timeout:    time.Duration(d.Timeout) * time.Millisecond,
		RecordType: synthetics.DNSRecord(d.RecordType.String()),
		Servers:    servers,
		Port:       d.Port,
	}, nil
}

func dnsGridTestDefinitionFromPayload(ts *syntheticspb.TestSettings) (synthetics.TestDefinition, error) {
	d := ts.GetDnsGrid()
	if d == nil {
		return nil, errors.New("DNS grid test definition is nil")
	}
	servers, err := convert.StringsToIPs(d.Servers)
	if err != nil {
		return nil, fmt.Errorf("convert DNS grid servers: %v", err)
	}

	return &synthetics.TestDefinitionDNSGrid{
		Target:     d.Target,
		Timeout:    time.Duration(d.Timeout) * time.Millisecond,
		RecordType: synthetics.DNSRecord(d.RecordType.String()),
		Servers:    servers,
		Port:       d.Port,
	}, nil
}

func taskTypesFromPayload(tasks []string) []synthetics.TaskType {
	var result []synthetics.TaskType
	for _, t := range tasks {
		result = append(result, synthetics.TaskType(t))
	}
	return result
}

func healthSettingsFromPayload(hs *syntheticspb.HealthSettings) (synthetics.HealthSettings, error) {
	if hs == nil {
		return synthetics.HealthSettings{}, fmt.Errorf("empty health settings")
	}

	alarmActivation, err := alarmActivationFromPayload(hs.Activation)
	if err != nil {
		return synthetics.HealthSettings{}, err
	}

	return synthetics.HealthSettings{
		LatencyCritical:           convert.MillisecondsF32ToDuration(hs.LatencyCritical),
		LatencyWarning:            convert.MillisecondsF32ToDuration(hs.LatencyWarning),
		LatencyCriticalStdDev:     convert.MillisecondsF32ToDuration(hs.LatencyCriticalStddev),
		LatencyWarningStdDev:      convert.MillisecondsF32ToDuration(hs.LatencyWarningStddev),
		JitterCritical:            convert.MillisecondsF32ToDuration(hs.JitterCritical),
		JitterWarning:             convert.MillisecondsF32ToDuration(hs.JitterWarning),
		JitterCriticalStdDev:      convert.MillisecondsF32ToDuration(hs.JitterCriticalStddev),
		JitterWarningStdDev:       convert.MillisecondsF32ToDuration(hs.JitterWarningStddev),
		PacketLossCritical:        hs.PacketLossCritical,
		PacketLossWarning:         hs.PacketLossWarning,
		HTTPLatencyCritical:       convert.MillisecondsF32ToDuration(hs.HttpLatencyCritical),
		HTTPLatencyWarning:        convert.MillisecondsF32ToDuration(hs.HttpLatencyWarning),
		HTTPLatencyCriticalStdDev: convert.MillisecondsF32ToDuration(hs.HttpLatencyCriticalStddev),
		HTTPLatencyWarningStdDev:  convert.MillisecondsF32ToDuration(hs.HttpLatencyWarningStddev),
		HTTPValidCodes:            hs.HttpValidCodes,
		DNSValidCodes:             hs.DnsValidCodes,
		UnhealthySubtestThreshold: hs.UnhealthySubtestThreshold,
		AlarmActivation:           alarmActivation,
	}, nil
}

func pingSettingsFromPayload(ps *syntheticspb.TestPingSettings) *synthetics.PingSettings {
	if ps == nil {
		return nil
	}

	return &synthetics.PingSettings{
		Timeout:  time.Duration(ps.Timeout) * time.Millisecond,
		Count:    ps.Count,
		Delay:    convert.MillisecondsF32ToDuration(ps.Delay),
		Protocol: synthetics.PingProtocol(ps.Protocol),
		Port:     ps.Port,
	}
}

func tracerouteSettingsFromPayload(ts *syntheticspb.TestTraceSettings) *synthetics.TracerouteSettings {
	if ts == nil {
		return nil
	}

	return &synthetics.TracerouteSettings{
		Timeout:  time.Duration(ts.Timeout) * time.Millisecond,
		Count:    ts.Count,
		Delay:    convert.MillisecondsF32ToDuration(ts.Delay),
		Protocol: synthetics.TracerouteProtocol(ts.Protocol),
		Port:     ts.Port,
		Limit:    ts.Limit,
	}
}

func alarmActivationFromPayload(as *syntheticspb.ActivationSettings) (*synthetics.AlarmActivationSettings, error) {
	if as == nil {
		return nil, nil
	}

	timeWindow, err := time.ParseDuration(as.TimeWindow + as.TimeUnit)
	if err != nil {
		return nil, fmt.Errorf("parse alarm activation time window %q: %v", as.TimeWindow+as.TimeUnit, err)
	}

	times, err := strconv.ParseUint(as.Times, 10, 0)
	if err != nil {
		return nil, fmt.Errorf("parse alarm activation times %q: %v", as.Times, err)
	}

	gracePeriod, err := strconv.ParseUint(as.GracePeriod, 10, 0)
	if err != nil {
		return nil, fmt.Errorf("parse alarm activation grace period %q: %v", as.GracePeriod, err)
	}

	return &synthetics.AlarmActivationSettings{
		TimeWindow:  timeWindow,
		Times:       uint(times),
		GracePeriod: uint(gracePeriod),
	}, nil
}

func userInfoPtrFromPayload(ui *syntheticspb.UserInfo) *synthetics.UserInfo {
	if ui == nil {
		return nil
	}
	return &synthetics.UserInfo{
		ID:       ui.Id,
		Email:    ui.Email,
		FullName: ui.FullName,
	}
}

// testToPayload converts synthetics test from model to payload. It sets only ID and read-write fields.
func testToPayload(t *synthetics.Test) (*syntheticspb.Test, error) {
	if t == nil {
		return nil, errors.New("test object is nil")
	}

	ts, err := testSettingsToPayload(t.Settings, t.Type)
	if err != nil {
		return nil, err
	}

	return &syntheticspb.Test{
		Id:            t.ID,
		Name:          t.Name,
		Type:          string(t.Type),
		Status:        syntheticspb.TestStatus(syntheticspb.TestStatus_value[string(t.Status)]),
		Settings:      ts,
		Cdate:         nil, // read-only
		Edate:         nil, // read-only
		CreatedBy:     nil, // read-only
		LastUpdatedBy: nil, // read-only
	}, nil
}

func testSettingsToPayload(ts synthetics.TestSettings, testType synthetics.TestType) (*syntheticspb.TestSettings, error) {
	tsPayload := &syntheticspb.TestSettings{
		AgentIds:             ts.AgentIDs,
		Tasks:                taskTypesToPayload(ts.Tasks),
		HealthSettings:       healthSettingsToPayload(ts.Health),
		Ping:                 pingSettingsToPayload(ts.Ping),
		Trace:                tracerouteSettingsToPayload(ts.Traceroute),
		Period:               uint32(ts.Period / time.Second),
		Family:               syntheticspb.IPFamily(syntheticspb.IPFamily_value[string(ts.Family)]),
		NotificationChannels: ts.NotificationChannels,
	}

	return testSettingsPayloadWithDefinition(tsPayload, ts, testType)
}

func taskTypesToPayload(tasks []synthetics.TaskType) []string {
	var result []string
	for _, t := range tasks {
		result = append(result, string(t))
	}
	return result
}

func healthSettingsToPayload(hs synthetics.HealthSettings) *syntheticspb.HealthSettings {
	return &syntheticspb.HealthSettings{
		LatencyCritical:           float32(hs.LatencyCritical / time.Millisecond),
		LatencyWarning:            float32(hs.LatencyWarning / time.Millisecond),
		PacketLossCritical:        hs.PacketLossCritical,
		PacketLossWarning:         hs.PacketLossWarning,
		JitterCritical:            float32(hs.JitterCritical / time.Millisecond),
		JitterWarning:             float32(hs.JitterWarning / time.Millisecond),
		HttpLatencyCritical:       float32(hs.HTTPLatencyCritical / time.Millisecond),
		HttpLatencyWarning:        float32(hs.HTTPLatencyWarning / time.Millisecond),
		HttpValidCodes:            hs.HTTPValidCodes,
		DnsValidCodes:             hs.DNSValidCodes,
		LatencyCriticalStddev:     float32(hs.LatencyCriticalStdDev / time.Millisecond),
		LatencyWarningStddev:      float32(hs.LatencyWarningStdDev / time.Millisecond),
		JitterCriticalStddev:      float32(hs.JitterCriticalStdDev / time.Millisecond),
		JitterWarningStddev:       float32(hs.JitterWarningStdDev / time.Millisecond),
		HttpLatencyCriticalStddev: float32(hs.HTTPLatencyCriticalStdDev / time.Millisecond),
		HttpLatencyWarningStddev:  float32(hs.HTTPLatencyWarningStdDev / time.Millisecond),
		UnhealthySubtestThreshold: hs.UnhealthySubtestThreshold,
		Activation:                alarmActivationToPayload(hs.AlarmActivation),
	}
}

func alarmActivationToPayload(as *synthetics.AlarmActivationSettings) *syntheticspb.ActivationSettings {
	if as == nil {
		return nil
	}

	return &syntheticspb.ActivationSettings{
		GracePeriod: strconv.FormatUint(uint64(as.GracePeriod), 10),
		TimeUnit:    "m",
		TimeWindow:  strconv.FormatInt(int64(as.TimeWindow/time.Minute), 10),
		Times:       strconv.FormatUint(uint64(as.Times), 10),
	}
}

func pingSettingsToPayload(ps *synthetics.PingSettings) *syntheticspb.TestPingSettings {
	if ps == nil {
		return nil
	}

	return &syntheticspb.TestPingSettings{
		Count:    ps.Count,
		Protocol: string(ps.Protocol),
		Port:     ps.Port,
		Timeout:  uint32(ps.Timeout / time.Millisecond),
		Delay:    float32(ps.Delay / time.Millisecond),
	}
}

func tracerouteSettingsToPayload(ts *synthetics.TracerouteSettings) *syntheticspb.TestTraceSettings {
	if ts == nil {
		return nil
	}

	return &syntheticspb.TestTraceSettings{
		Count:    ts.Count,
		Protocol: string(ts.Protocol),
		Port:     ts.Port,
		Timeout:  uint32(ts.Timeout / time.Millisecond),
		Limit:    ts.Limit,
		Delay:    float32(ts.Delay / time.Millisecond),
	}
}

// nolint: gocyclo
func testSettingsPayloadWithDefinition(
	tsPayload *syntheticspb.TestSettings, ts synthetics.TestSettings, testType synthetics.TestType,
) (*syntheticspb.TestSettings, error) {
	switch testType {
	case synthetics.TestTypeIP:
		tsPayload.Definition = ipTestDefinitionToPayload(ts)
	case synthetics.TestTypeNetworkGrid:
		tsPayload.Definition = networkGridTestDefinitionToPayload(ts)
	case synthetics.TestTypeHostname:
		tsPayload.Definition = hostnameTestDefinitionToPayload(ts)
	case synthetics.TestTypeAgent:
		tsPayload.Definition = agentTestDefinitionToPayload(ts)
	case synthetics.TestTypeNetworkMesh:
		tsPayload.Definition = networkMeshTestDefinitionToPayload(ts)
	case synthetics.TestTypeFlow:
		tsPayload.Definition = flowTestDefinitionToPayload(ts)
	case synthetics.TestTypeURL:
		tsPayload.Definition = urlTestDefinitionToPayload(ts)
	case synthetics.TestTypePageLoad:
		tsPayload.Definition = pageLoadTestDefinitionToPayload(ts)
	case synthetics.TestTypeDNS:
		tsPayload.Definition = dnsTestDefinitionToPayload(ts)
	case synthetics.TestTypeDNSGrid:
		tsPayload.Definition = dnsGridTestDefinitionToPayload(ts)
	default:
		return nil, fmt.Errorf("unsupported test type: %v", testType)
	}

	return tsPayload, nil
}

func ipTestDefinitionToPayload(ts synthetics.TestSettings) *syntheticspb.TestSettings_Ip {
	return &syntheticspb.TestSettings_Ip{
		Ip: &syntheticspb.IpTest{
			Targets: convert.IPsToStrings(ts.GetIPDefinition().Targets),
		},
	}
}

func networkGridTestDefinitionToPayload(ts synthetics.TestSettings) *syntheticspb.TestSettings_NetworkGrid {
	return &syntheticspb.TestSettings_NetworkGrid{
		NetworkGrid: &syntheticspb.IpTest{
			Targets: convert.IPsToStrings(ts.GetNetworkGridDefinition().Targets),
		},
	}
}

func hostnameTestDefinitionToPayload(ts synthetics.TestSettings) *syntheticspb.TestSettings_Hostname {
	return &syntheticspb.TestSettings_Hostname{
		Hostname: &syntheticspb.HostnameTest{
			Target: ts.GetHostnameDefinition().Target,
		},
	}
}

func agentTestDefinitionToPayload(ts synthetics.TestSettings) *syntheticspb.TestSettings_Agent {
	return &syntheticspb.TestSettings_Agent{
		Agent: &syntheticspb.AgentTest{
			Target:     ts.GetAgentDefinition().Target,
			UseLocalIp: ts.GetAgentDefinition().UseLocalIP,
		},
	}
}

func networkMeshTestDefinitionToPayload(ts synthetics.TestSettings) *syntheticspb.TestSettings_NetworkMesh {
	return &syntheticspb.TestSettings_NetworkMesh{
		NetworkMesh: &syntheticspb.NetworkMeshTest{
			UseLocalIp: ts.GetNetworkMeshDefinition().UseLocalIP,
		},
	}
}

func flowTestDefinitionToPayload(ts synthetics.TestSettings) *syntheticspb.TestSettings_Flow {
	d := ts.GetFlowDefinition()
	return &syntheticspb.TestSettings_Flow{
		Flow: &syntheticspb.FlowTest{
			Target:                      d.Target,
			TargetRefreshIntervalMillis: uint32(d.TargetRefreshInterval / time.Millisecond),
			MaxProviders:                d.MaxProviders,
			MaxIpTargets:                d.MaxIPTargets,
			Type:                        string(d.Type),
			InetDirection:               string(d.InetDirection),
			Direction:                   string(d.Direction),
		},
	}
}

func urlTestDefinitionToPayload(ts synthetics.TestSettings) *syntheticspb.TestSettings_Url {
	d := ts.GetURLDefinition()
	return &syntheticspb.TestSettings_Url{
		Url: &syntheticspb.UrlTest{
			Target:          d.Target.String(),
			Timeout:         uint32(d.Timeout / time.Millisecond),
			Method:          d.Method,
			Headers:         d.Headers,
			Body:            d.Body,
			IgnoreTlsErrors: d.IgnoreTLSErrors,
		},
	}
}

func pageLoadTestDefinitionToPayload(ts synthetics.TestSettings) *syntheticspb.TestSettings_PageLoad {
	d := ts.GetPageLoadDefinition()
	return &syntheticspb.TestSettings_PageLoad{
		PageLoad: &syntheticspb.PageLoadTest{
			Target:          d.Target.String(),
			Timeout:         uint32(d.Timeout / time.Millisecond),
			Headers:         d.Headers,
			IgnoreTlsErrors: d.IgnoreTLSErrors,
			CssSelectors:    d.CSSSelectors,
		},
	}
}

func dnsTestDefinitionToPayload(ts synthetics.TestSettings) *syntheticspb.TestSettings_Dns {
	return &syntheticspb.TestSettings_Dns{
		Dns: dnsTestSubDefinitionToPayload(ts.GetDNSDefinition()),
	}
}

func dnsGridTestDefinitionToPayload(ts synthetics.TestSettings) *syntheticspb.TestSettings_DnsGrid {
	return &syntheticspb.TestSettings_DnsGrid{
		DnsGrid: dnsTestSubDefinitionToPayload(ts.GetDNSGridDefinition()),
	}
}

func dnsTestSubDefinitionToPayload(d *synthetics.TestDefinitionDNS) *syntheticspb.DnsTest {
	return &syntheticspb.DnsTest{
		Target:     d.Target,
		Timeout:    uint32(d.Timeout / time.Millisecond),
		RecordType: syntheticspb.DNSRecord(syntheticspb.DNSRecord_value[string(d.RecordType)]),
		Servers:    convert.IPsToStrings(d.Servers),
		Port:       d.Port,
	}
}
