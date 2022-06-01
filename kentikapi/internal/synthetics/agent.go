package synthetics

import (
	"fmt"

	syntheticspb "github.com/kentik/api-schema-public/gen/go/kentik/synthetics/v202202"
	"github.com/kentik/community_sdk_golang/kentikapi/cloud"
	kentikerrors "github.com/kentik/community_sdk_golang/kentikapi/internal/errors"
	"github.com/kentik/community_sdk_golang/kentikapi/synthetics"
)

type listAgentsResponse syntheticspb.ListAgentsResponse

func (r *listAgentsResponse) ToModel() (*synthetics.GetAllAgentsResponse, error) {
	if r == nil {
		return nil, nil
	}

	agents, err := agentsFromPayload(r.Agents)
	if err != nil {
		return nil, err
	}

	return &synthetics.GetAllAgentsResponse{
		Agents:             agents,
		InvalidAgentsCount: r.InvalidCount,
	}, nil
}

func agentsFromPayload(agents []*syntheticspb.Agent) ([]synthetics.Agent, error) {
	var result []synthetics.Agent
	for i, a := range agents {
		agent, err := agentFromPayload(a)
		if err != nil {
			return nil, fmt.Errorf("agent with index %v: %w", i, err)
		}
		result = append(result, *agent)
	}
	return result, nil
}

// agentFromPayload converts synthetics agent payload to model.
func agentFromPayload(a *syntheticspb.Agent) (*synthetics.Agent, error) {
	if a == nil {
		return nil, kentikerrors.New(kentikerrors.InvalidResponse, "agent response payload is nil")
	}

	if a.Id == "" {
		return nil, kentikerrors.New(kentikerrors.InvalidResponse, "empty agent ID in response payload")
	}

	return &synthetics.Agent{
		Status:             synthetics.AgentStatus(a.Status.String()),
		Alias:              a.Alias,
		SiteID:             a.SiteId,
		LocalIP:            a.LocalIp,
		IPFamily:           synthetics.IPFamily(a.Family.String()),
		CloudProvider:      cloudProviderFromPayload(a.CloudProvider),
		CloudRegion:        a.CloudRegion,
		ID:                 a.Id,
		Type:               synthetics.AgentType(a.Type),
		SiteName:           a.SiteName,
		IP:                 a.Ip,
		ASN:                a.Asn,
		Latitude:           a.Lat,
		Longitude:          a.Long,
		City:               a.City,
		Region:             a.Region,
		Country:            a.Country,
		Version:            a.Version,
		OS:                 a.Os,
		ImplementationType: synthetics.AgentImplementationType(a.AgentImpl.String()),
		LastAuthed:         a.LastAuthed.AsTime(),
		TestIDs:            a.TestIds,
	}, nil
}

func cloudProviderFromPayload(cp string) cloud.Provider {
	// Agents API uses "gcp" value for GCE provider
	if cp == "gcp" {
		return cloud.ProviderGCE
	}
	return cloud.Provider(cp)
}

// agentToPayload converts synthetics agent from model to payload. It sets only ID, SiteName and
// read-write fields.
//nolint:unparam // TODO: return InvalidRequest error on empty request
func agentToPayload(a *synthetics.Agent) (*syntheticspb.Agent, error) {
	if a == nil {
		return nil, nil
	}

	return &syntheticspb.Agent{
		Id:            a.ID,
		SiteName:      a.SiteName, // read-only, but required for an update to work
		Status:        syntheticspb.AgentStatus(syntheticspb.AgentStatus_value[string(a.Status)]),
		Alias:         a.Alias,
		SiteId:        a.SiteID,
		LocalIp:       a.LocalIP,
		Family:        syntheticspb.IPFamily(syntheticspb.IPFamily_value[string(a.IPFamily)]),
		CloudProvider: cloudProviderToPayload(a.CloudProvider),
		CloudRegion:   a.CloudRegion,
	}, nil
}

func cloudProviderToPayload(cp cloud.Provider) string {
	// Agents API uses "gcp" value for GCE provider
	if cp == cloud.ProviderGCE {
		return "gcp"
	}
	return string(cp)
}
