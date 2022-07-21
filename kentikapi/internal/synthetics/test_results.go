package synthetics

import (
	"fmt"
	"net"
	"net/url"
	"time"

	syntheticspb "github.com/kentik/api-schema-public/gen/go/kentik/synthetics/v202202"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/convert"
	"github.com/kentik/community_sdk_golang/kentikapi/synthetics"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func getResultsForTestsRequestToPayload(
	req synthetics.GetResultsForTestsRequest,
) *syntheticspb.GetResultsForTestsRequest {
	return &syntheticspb.GetResultsForTestsRequest{
		Ids:       req.TestIDs,
		StartTime: timestamppb.New(req.StartTime),
		EndTime:   timestamppb.New(req.EndTime),
		AgentIds:  req.AgentIDs,
		Targets:   convert.IPsToStrings(req.Targets),
	}
}

type getResultsForTestsResponse syntheticspb.GetResultsForTestsResponse

func (r *getResultsForTestsResponse) ToModel() ([]synthetics.TestResults, error) {
	if r == nil {
		return nil, nil
	}

	var trSlice []synthetics.TestResults
	for i, trPayload := range r.Results {
		tr, err := testResultsFromPayload(trPayload)
		if err != nil {
			return nil, fmt.Errorf("test results with index %v: %w", i, err)
		}
		trSlice = append(trSlice, tr)
	}

	return trSlice, nil
}

func testResultsFromPayload(trPayload *syntheticspb.TestResults) (synthetics.TestResults, error) {
	if trPayload == nil {
		return synthetics.TestResults{}, fmt.Errorf("test results are nil")
	}

	agentResults, err := agentResultsSliceFromPayload(trPayload.Agents)
	if err != nil {
		return synthetics.TestResults{}, err
	}

	return synthetics.TestResults{
		TestID: trPayload.TestId,
		Time:   trPayload.Time.AsTime(),
		Health: synthetics.Health(trPayload.Health),
		Agents: agentResults,
	}, nil
}

func agentResultsSliceFromPayload(arsPayload []*syntheticspb.AgentResults) ([]synthetics.AgentResults, error) {
	var ars []synthetics.AgentResults
	for i, arPayload := range arsPayload {
		ar, err := agentResultsFromPayload(arPayload)
		if err != nil {
			return nil, fmt.Errorf("agent results with index %v: %w", i, err)
		}
		ars = append(ars, ar)
	}

	return ars, nil
}

func agentResultsFromPayload(arPayload *syntheticspb.AgentResults) (synthetics.AgentResults, error) {
	if arPayload == nil {
		return synthetics.AgentResults{}, fmt.Errorf("agent results are nil")
	}

	tasks, err := taskResultsSliceFromPayload(arPayload.Tasks)
	if err != nil {
		return synthetics.AgentResults{}, err
	}

	return synthetics.AgentResults{
		AgentID: arPayload.AgentId,
		Health:  synthetics.Health(arPayload.Health),
		Tasks:   tasks,
	}, nil
}

func taskResultsSliceFromPayload(trsPayload []*syntheticspb.TaskResults) ([]synthetics.TaskResults, error) {
	var trs []synthetics.TaskResults
	for i, trPayload := range trsPayload {
		tr, err := taskResultsFromPayload(trPayload)
		if err != nil {
			return nil, fmt.Errorf("task results with index %v: %w", i, err)
		}
		trs = append(trs, tr)
	}

	return trs, nil
}

func taskResultsFromPayload(trPayload *syntheticspb.TaskResults) (synthetics.TaskResults, error) {
	if trPayload == nil {
		return synthetics.TaskResults{}, fmt.Errorf("task results are nil")
	}

	task, taskType, err := taskSpecificResultsFromPayload(trPayload)
	if err != nil {
		return synthetics.TaskResults{}, err
	}

	return synthetics.TaskResults{
		Health:   synthetics.Health(trPayload.Health),
		TaskType: taskType,
		Task:     task,
	}, nil
}

func taskSpecificResultsFromPayload(
	trPayload *syntheticspb.TaskResults,
) (synthetics.TaskSpecificResults, synthetics.TaskType, error) {
	if trPayload.TaskType == nil {
		return nil, "", fmt.Errorf("task specific results are nil")
	}

	switch tt := trPayload.TaskType.(type) {
	case *syntheticspb.TaskResults_Ping:
		r, err := pingTaskResultsFromPayload(tt)
		return r, synthetics.TaskTypePing, err
	case *syntheticspb.TaskResults_Http:
		r, err := httpTaskResultsFromPayload(tt)
		return r, synthetics.TaskTypeHTTP, err
	case *syntheticspb.TaskResults_Dns:
		r, err := dnsTaskResultsFromPayload(tt)
		return r, synthetics.TaskTypeDNS, err
	default:
		return nil, "", fmt.Errorf("unsupported task type: %v", trPayload.TaskType)
	}
}

func pingTaskResultsFromPayload(r *syntheticspb.TaskResults_Ping) (synthetics.TaskSpecificResults, error) {
	if r.Ping == nil {
		return nil, fmt.Errorf("ping task specific results are nil")
	}

	packetLoss, err := packetLossDataFromPayload(r.Ping.PacketLoss)
	if err != nil {
		return nil, err
	}

	latency, err := metricDataFromPayload(r.Ping.Latency)
	if err != nil {
		return nil, err
	}

	jitter, err := metricDataFromPayload(r.Ping.Jitter)
	if err != nil {
		return nil, err
	}

	dstIP := net.ParseIP(r.Ping.DstIp)
	if dstIP == nil {
		return nil, fmt.Errorf("cannot parse ping destination IP %q", r.Ping.DstIp)
	}

	return synthetics.PingResults{
		Target:     r.Ping.Target,
		PacketLoss: packetLoss,
		Latency:    latency,
		Jitter:     jitter,
		DstIP:      dstIP,
	}, nil
}

func httpTaskResultsFromPayload(r *syntheticspb.TaskResults_Http) (synthetics.TaskSpecificResults, error) {
	if r.Http == nil {
		return nil, fmt.Errorf("HTTP task specific results are nil")
	}

	target, err := url.Parse(r.Http.Target)
	if err != nil {
		return nil, err
	}

	latency, err := metricDataFromPayload(r.Http.Latency)
	if err != nil {
		return nil, err
	}

	response, err := httpResponseDataFromPayload(r.Http.Response)
	if err != nil {
		return nil, err
	}

	dstIP := net.ParseIP(r.Http.DstIp)
	if dstIP == nil {
		return nil, fmt.Errorf("cannot parse HTTP destination IP %q", r.Http.DstIp)
	}

	return synthetics.HTTPResults{
		Target:   *target,
		Latency:  latency,
		Response: response,
		DstIP:    dstIP,
	}, nil
}

func dnsTaskResultsFromPayload(r *syntheticspb.TaskResults_Dns) (synthetics.TaskSpecificResults, error) {
	if r.Dns == nil {
		return nil, fmt.Errorf("DNS task specific results are nil")
	}

	server := net.ParseIP(r.Dns.Server)
	if server == nil {
		return nil, fmt.Errorf("cannot parse DNS destination IP %q", r.Dns.Server)
	}

	latency, err := metricDataFromPayload(r.Dns.Latency)
	if err != nil {
		return nil, err
	}

	response, err := dnsResponseDataFromPayload(r.Dns.Response)
	if err != nil {
		return nil, err
	}

	return synthetics.DNSResults{
		Target:   r.Dns.Target,
		Server:   server,
		Latency:  latency,
		Response: response,
	}, nil
}

func packetLossDataFromPayload(d *syntheticspb.PacketLossData) (synthetics.PacketLossData, error) {
	if d == nil {
		return synthetics.PacketLossData{}, fmt.Errorf("packet loss data is nil")
	}

	return synthetics.PacketLossData{
		Current: d.Current,
		Health:  synthetics.Health(d.Health),
	}, nil
}

func metricDataFromPayload(d *syntheticspb.MetricData) (synthetics.MetricData, error) {
	if d == nil {
		return synthetics.MetricData{}, fmt.Errorf("metric data is nil")
	}

	return synthetics.MetricData{
		Current:       time.Duration(d.Current) * time.Microsecond,
		RollingAvg:    time.Duration(d.RollingAvg) * time.Microsecond,
		RollingStdDev: time.Duration(d.RollingStddev) * time.Microsecond,
		Health:        synthetics.Health(d.Health),
	}, nil
}

func httpResponseDataFromPayload(d *syntheticspb.HTTPResponseData) (synthetics.HTTPResponseData, error) {
	if d == nil {
		return synthetics.HTTPResponseData{}, fmt.Errorf("HTTP response data is nil")
	}

	return synthetics.HTTPResponseData{
		Status: d.Status,
		Size:   d.Size,
		Data:   d.Data,
	}, nil
}

func dnsResponseDataFromPayload(d *syntheticspb.DNSResponseData) (synthetics.DNSResponseData, error) {
	if d == nil {
		return synthetics.DNSResponseData{}, fmt.Errorf("DNS response data is nil")
	}

	return synthetics.DNSResponseData{
		Status: d.Status,
		Data:   d.Data,
	}, nil
}
