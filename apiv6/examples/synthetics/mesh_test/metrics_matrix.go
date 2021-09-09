package main

import (
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/synthetics"
)

// metricsMatrix holds "fromAgent" -> "toAgent" connection metrics
type metricsMatrix struct {
	agents []string
	cells  map[string]map[string]*synthetics.V202101beta1MeshMetrics
}

func newMetricsMatrix(mesh []synthetics.V202101beta1MeshResponse) metricsMatrix {
	// fill agents
	agents := []string{}
	for _, agent := range mesh {
		agents = append(agents, agent.GetAlias())
	}

	// fill matrix cells
	cells := make(map[string]map[string]*synthetics.V202101beta1MeshMetrics)
	for _, fromAgent := range mesh {
		cells[fromAgent.GetAlias()] = make(map[string]*synthetics.V202101beta1MeshMetrics)
		for _, toAgent := range *fromAgent.Columns {
			cells[fromAgent.GetAlias()][toAgent.GetAlias()] = toAgent.Metrics
		}
	}
	return metricsMatrix{agents: agents, cells: cells}
}

func (m metricsMatrix) getMetric(fromAgent string, toAgent string) (*synthetics.V202101beta1MeshMetrics, bool) {
	toAgents, ok := m.cells[fromAgent]
	if !ok {
		return nil, false
	}

	metric, ok := toAgents[toAgent]
	if !ok {
		return nil, false
	}

	return metric, true
}
