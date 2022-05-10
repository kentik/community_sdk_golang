package resources

import (
	"context"
	"fmt"
	"log"

	syntheticspb "github.com/kentik/api-schema-public/gen/go/kentik/synthetics/v202202"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"google.golang.org/grpc"
)

// SyntheticsAPI aggregates Synthetics API methods.
type SyntheticsAPI struct {
	Agents AgentsAPI
}

// NewSyntheticsAPI creates new SyntheticsAPI.
func NewSyntheticsAPI(cc grpc.ClientConnInterface) *SyntheticsAPI {
	return &SyntheticsAPI{
		Agents: AgentsAPI{
			client: syntheticspb.NewSyntheticsAdminServiceClient(cc),
		},
	}
}

// AgentsAPI aggregates Synthetics Agents API methods.
type AgentsAPI struct {
	client syntheticspb.SyntheticsAdminServiceClient
}

// GetAll lists synthetics agents.
func (a *AgentsAPI) GetAll(ctx context.Context) (*models.GetAllSyntheticsAgentsResponse, error) {
	response, err := a.client.ListAgents(ctx, &syntheticspb.ListAgentsRequest{})
	if err != nil {
		return nil, err
	}

	return (*api_payloads.ListSyntheticsAgentsResponse)(response).ToModel()
}

// Get retrieves synthetics agent with given ID.
func (a AgentsAPI) Get(ctx context.Context, id models.ID) (*models.SyntheticsAgent, error) {
	response, err := a.client.GetAgent(ctx, &syntheticspb.GetAgentRequest{Id: id})
	if err != nil {
		return nil, err
	}

	return api_payloads.SyntheticsAgentFromPayload(response.GetAgent())
}

// Update updates the synthetics agent.
func (a *AgentsAPI) Update(ctx context.Context, agent *models.SyntheticsAgent) (*models.SyntheticsAgent, error) {
	payload, err := api_payloads.SyntheticsAgentToPayload(agent)
	if err != nil {
		return nil, err
	}

	response, err := a.client.UpdateAgent(ctx, &syntheticspb.UpdateAgentRequest{
		Agent: payload,
	})
	if err != nil {
		return nil, err
	}

	return api_payloads.SyntheticsAgentFromPayload(response.GetAgent())
}

// Delete removes synthetics agent with given ID.
func (a *AgentsAPI) Delete(ctx context.Context, id models.ID) error {
	_, err := a.client.DeleteAgent(ctx, &syntheticspb.DeleteAgentRequest{Id: id})
	return err
}

// Activate activates pending (waiting) synthetics agent with given ID.
func (a AgentsAPI) Activate(ctx context.Context, id models.ID) (*models.SyntheticsAgent, error) {
	agent, err := a.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if agent.Status != models.AgentStatusWait {
		log.Printf("Agent %q is not pending (status: %v) - skipping activation", agent.ID, agent.Status)
		return agent, nil
	}

	agent.Status = models.AgentStatusOK
	agent, err = a.Update(ctx, agent)
	if err != nil {
		return nil, err
	}

	if agent.Status != models.AgentStatusOK {
		return nil, fmt.Errorf("failed to activate the agent %q (status: %v)", id, agent.Status)
	}

	return agent, nil
}

// Deactivate deactivates the active client with given ID.
func (a AgentsAPI) Deactivate(ctx context.Context, id models.ID) (*models.SyntheticsAgent, error) {
	agent, err := a.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if agent.Status != models.AgentStatusOK {
		log.Printf("Agent %q is not active (status: %v) - skipping deactivation", agent.ID, agent.Status)
		return agent, nil
	}

	agent.Status = models.AgentStatusWait
	agent, err = a.Update(ctx, agent)
	if err != nil {
		return nil, err
	}

	if agent.Status != models.AgentStatusWait {
		return nil, fmt.Errorf("failed to deactivate the agent %q (status: %v)", id, agent.Status)
	}

	return agent, nil
}
