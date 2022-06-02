package api_payloads

import (
	"fmt"

	syntheticspb "github.com/kentik/api-schema-public/gen/go/kentik/synthetics/v202202"
	kentikerrors "github.com/kentik/community_sdk_golang/kentikapi/internal/errors"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type ListSyntheticsAgentsResponse syntheticspb.ListAgentsResponse

func (r *ListSyntheticsAgentsResponse) ToModel() (*models.GetAllSyntheticsAgentsResponse, error) {
	if r == nil {
		return nil, nil
	}

	agents, err := syntheticsAgentsFromPayload(r.Agents)
	if err != nil {
		return nil, err
	}

	return &models.GetAllSyntheticsAgentsResponse{
		Agents:             agents,
		InvalidAgentsCount: r.InvalidCount,
	}, nil
}

func syntheticsAgentsFromPayload(agents []*syntheticspb.Agent) ([]models.SyntheticsAgent, error) {
	var result []models.SyntheticsAgent
	for i, a := range agents {
		agent, err := SyntheticsAgentFromPayload(a)
		if err != nil {
			return nil, fmt.Errorf("agent with index %v: %w", i, err)
		}
		result = append(result, *agent)
	}
	return result, nil
}

// SyntheticsAgentFromPayload converts synthetics agent payload to model.
func SyntheticsAgentFromPayload(a *syntheticspb.Agent) (*models.SyntheticsAgent, error) {
	if a == nil {
		return nil, kentikerrors.New(kentikerrors.InvalidResponse, "agent response payload is nil")
	}

	if a.Id == "" {
		return nil, kentikerrors.New(kentikerrors.InvalidResponse, "empty agent ID in response payload")
	}

	return &models.SyntheticsAgent{
		Status:             models.AgentStatus(a.Status.String()),
		Alias:              a.Alias,
		SiteID:             a.SiteId,
		LocalIP:            a.LocalIp,
		IPFamily:           models.IPFamily(a.Family.String()),
		CloudProvider:      cloudProviderFromPayload(a.CloudProvider),
		CloudRegion:        a.CloudRegion,
		ID:                 a.Id,
		Type:               models.AgentType(a.Type),
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
		ImplementationType: models.AgentImplementationType(a.AgentImpl.String()),
		LastAuthed:         a.LastAuthed.AsTime(),
		TestIDs:            a.TestIds,
	}, nil
}

func cloudProviderFromPayload(cp string) models.CloudProvider {
	// Agents API uses "gcp" value for GCE provider
	if cp == "gcp" {
		return models.CloudProviderGCE
	}
	return models.CloudProvider(cp)
}

// SyntheticsAgentToPayload converts synthetics agent from model to payload. It sets only ID, SiteName and
// read-write fields.
func SyntheticsAgentToPayload(a *models.SyntheticsAgent) (*syntheticspb.Agent, error) {
	if a == nil {
		return nil, nil
	}

	return &syntheticspb.Agent{
		Id:            a.ID,
		SiteName:      a.SiteName, // read-only, but required for update to work
		Status:        syntheticspb.AgentStatus(syntheticspb.AgentStatus_value[string(a.Status)]),
		Alias:         a.Alias,
		SiteId:        a.SiteID,
		LocalIp:       a.LocalIP,
		Family:        syntheticspb.IPFamily(syntheticspb.IPFamily_value[string(a.IPFamily)]),
		CloudProvider: cloudProviderToPayload(a.CloudProvider),
		CloudRegion:   a.CloudRegion,
	}, nil
}

func cloudProviderToPayload(cp models.CloudProvider) string {
	// Agents API uses "gcp" value for GCE provider
	if cp == models.CloudProviderGCE {
		return "gcp"
	}
	return string(cp)
}
