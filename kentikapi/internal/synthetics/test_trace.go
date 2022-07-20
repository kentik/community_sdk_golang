package synthetics

import (
	"errors"
	"fmt"
	"net"

	syntheticspb "github.com/kentik/api-schema-public/gen/go/kentik/synthetics/v202202"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/convert"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/kentik/community_sdk_golang/kentikapi/synthetics"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func getTraceForTestRequestToPayload(req synthetics.GetTraceForTestRequest) *syntheticspb.GetTraceForTestRequest {
	return &syntheticspb.GetTraceForTestRequest{
		Id:        req.TestID,
		StartTime: timestamppb.New(req.StartTime),
		EndTime:   timestamppb.New(req.EndTime),
		AgentIds:  req.AgentIDs,
		TargetIps: convert.IPsToStrings(req.Targets),
	}
}

type getTraceForTestResponse syntheticspb.GetTraceForTestResponse

func (r *getTraceForTestResponse) ToModel() (synthetics.GetTraceForTestResponse, error) {
	if r == nil {
		return synthetics.GetTraceForTestResponse{}, errors.New("response payload is nil")
	}

	nodes, err := nodesFromPayload(r.Nodes)
	if err != nil {
		return synthetics.GetTraceForTestResponse{}, err
	}

	paths, err := pathsFromPayload(r.Paths)
	if err != nil {
		return synthetics.GetTraceForTestResponse{}, err
	}

	return synthetics.GetTraceForTestResponse{
		Nodes: nodes,
		Paths: paths,
	}, nil
}

func nodesFromPayload(nodesP map[string]*syntheticspb.NetNode) (map[string]synthetics.NetworkNode, error) {
	nodes := make(map[models.ID]synthetics.NetworkNode, len(nodesP))
	for nodeID, nodeP := range nodesP {
		node, err := nodeFromPayload(nodeP)
		if err != nil {
			return nil, fmt.Errorf("node with ID %v: %w", nodeID, err)
		}

		nodes[nodeID] = node
	}

	return nodes, nil
}

func nodeFromPayload(n *syntheticspb.NetNode) (synthetics.NetworkNode, error) {
	if n == nil {
		return synthetics.NetworkNode{}, fmt.Errorf("node payload is nil")
	}

	ip := net.ParseIP(n.Ip)
	if ip == nil {
		return synthetics.NetworkNode{}, fmt.Errorf("cannot parse node IP %q", n.Ip)
	}

	return synthetics.NetworkNode{
		IP:       ip,
		ASN:      n.Asn,
		AsName:   n.AsName,
		Location: locationFromPayload(n.Location),
		DNSName:  n.DnsName,
		DeviceID: n.DeviceId,
		SiteID:   n.SiteId,
	}, nil
}

func locationFromPayload(l *syntheticspb.Location) *synthetics.Location {
	if l == nil {
		return nil
	}

	return &synthetics.Location{
		Latitude:  l.Latitude,
		Longitude: l.Longitude,
		Country:   l.Country,
		Region:    l.Region,
		City:      l.City,
	}
}

func pathsFromPayload(pathsP []*syntheticspb.Path) ([]synthetics.Path, error) {
	var paths []synthetics.Path
	for i, pathP := range pathsP {
		path, err := pathFromPayload(pathP)
		if err != nil {
			return nil, fmt.Errorf("path with index %v: %w", i, err)
		}

		paths = append(paths, path)
	}

	return paths, nil
}

func pathFromPayload(p *syntheticspb.Path) (synthetics.Path, error) {
	if p == nil {
		return synthetics.Path{}, errors.New("path is nil")
	}

	targetIP := net.ParseIP(p.TargetIp)
	if targetIP == nil {
		return synthetics.Path{}, fmt.Errorf("cannot parse target IP %q", p.TargetIp)
	}

	hopCount, err := statsFromPayload(p.HopCount)
	if err != nil {
		return synthetics.Path{}, err
	}

	traces, err := tracesFromPayload(p.Traces)
	if err != nil {
		return synthetics.Path{}, err
	}

	return synthetics.Path{
		Time:            p.Time.AsTime(),
		AgentID:         p.AgentId,
		TargetIP:        targetIP,
		HopCount:        hopCount,
		MaxASPathLength: p.MaxAsPathLength,
		Traces:          traces,
	}, nil
}

func statsFromPayload(s *syntheticspb.Stats) (synthetics.Stats, error) {
	if s == nil {
		return synthetics.Stats{}, errors.New("stats object is nil")
	}

	return synthetics.Stats{
		Average: s.Average,
		Min:     s.Min,
		Max:     s.Max,
	}, nil
}

func tracesFromPayload(tracesP []*syntheticspb.PathTrace) ([]synthetics.PathTrace, error) {
	var traces []synthetics.PathTrace
	for i, traceP := range tracesP {
		trace, err := traceFromPayload(traceP)
		if err != nil {
			return nil, fmt.Errorf("trace with index %v: %q", i, err)
		}

		traces = append(traces, trace)
	}

	return traces, nil
}

func traceFromPayload(p *syntheticspb.PathTrace) (synthetics.PathTrace, error) {
	if p == nil {
		return synthetics.PathTrace{}, errors.New("path trace is nil")
	}

	hops, err := hopsFromPayload(p.Hops)
	if err != nil {
		return synthetics.PathTrace{}, err
	}

	return synthetics.PathTrace{
		ASPath:     p.AsPath,
		IsComplete: p.IsComplete,
		Hops:       hops,
	}, nil
}

func hopsFromPayload(hopsP []*syntheticspb.TraceHop) ([]synthetics.TraceHop, error) {
	var hops []synthetics.TraceHop
	for i, hopP := range hopsP {
		hop, err := hopFromPayload(hopP)
		if err != nil {
			return nil, fmt.Errorf("hop with index %v: %w", i, err)
		}

		hops = append(hops, hop)
	}

	return hops, nil
}

func hopFromPayload(h *syntheticspb.TraceHop) (synthetics.TraceHop, error) {
	if h == nil {
		return synthetics.TraceHop{}, errors.New("trace hop is nil")
	}

	return synthetics.TraceHop{
		Latency: h.Latency,
		NodeID:  h.NodeId,
	}, nil
}
