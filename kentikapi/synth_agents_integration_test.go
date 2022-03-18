package kentikapi_test

import (
	"context"
	"net"
	"testing"
	"time"

	syntheticspb "github.com/kentik/api-schema-public/gen/go/kentik/synthetics/v202202"
	"github.com/kentik/community_sdk_golang/kentikapi"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/testutil"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestClient_GetAllSyntheticsAgents(t *testing.T) {
	tests := []struct {
		name              string
		response          listAgentsResponse
		expectedResult    *models.GetAllSyntheticsAgentsResponse
		expectedError     bool
		expectedErrorCode *codes.Code
	}{
		{
			name: "status InvalidArgument received",
			response: listAgentsResponse{
				err: status.Errorf(codes.InvalidArgument, codes.InvalidArgument.String()),
			},
			expectedError:     true,
			expectedErrorCode: codePtr(codes.InvalidArgument),
		}, {
			name: "empty response received",
			response: listAgentsResponse{
				data: &syntheticspb.ListAgentsResponse{},
			},
			expectedResult: &models.GetAllSyntheticsAgentsResponse{
				Agents:             nil,
				InvalidAgentsCount: 0,
			},
		}, {
			name: "no exports received",
			response: listAgentsResponse{
				data: &syntheticspb.ListAgentsResponse{
					Agents:       []*syntheticspb.Agent{},
					InvalidCount: 0,
				},
			},
			expectedResult: &models.GetAllSyntheticsAgentsResponse{
				Agents:             nil,
				InvalidAgentsCount: 0,
			},
		}, {
			name: "2 agents received",
			response: listAgentsResponse{
				data: &syntheticspb.ListAgentsResponse{
					Agents: []*syntheticspb.Agent{
						newWarsawAgentPayload(),
						newMoscowAgentPayload(),
					},
					InvalidCount: 1,
				},
			},
			expectedResult: &models.GetAllSyntheticsAgentsResponse{
				Agents: []models.SyntheticsAgent{
					*newWarsawAgent(),
					*newMoscowAgent(),
				},
				InvalidAgentsCount: 1,
			},
		}, {
			name: "2 exports received - one empty",
			response: listAgentsResponse{
				data: &syntheticspb.ListAgentsResponse{
					Agents: []*syntheticspb.Agent{
						newWarsawAgentPayload(),
						nil,
					},
					InvalidCount: 0,
				},
			},
			// empty response fails validation
			expectedError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			server := newSpySyntheticsServer(t, syntheticsResponses{
				listAgentsResponse: tt.response,
			})
			server.Start()
			defer server.Stop()

			client, err := kentikapi.NewClient(
				kentikapi.WithAPIURL("http://"+server.url),
				kentikapi.WithCredentials(dummyAuthEmail, dummyAuthToken),
				kentikapi.WithLogPayloads(),
			)
			require.NoError(t, err)

			// act
			result, err := client.Synthetics.Agents.GetAll(context.Background())

			// assert
			t.Logf("Got result: %+v, err: %v", result, err)
			if tt.expectedError {
				assert.Error(t, err)
				if tt.expectedErrorCode != nil {
					s, ok := status.FromError(err)
					assert.True(t, ok)
					assert.Equal(t, *tt.expectedErrorCode, s.Code())
				}
			} else {
				assert.NoError(t, err)
			}

			if assert.Equal(t, 1, len(server.requests.listAgentsRequests), "invalid number of requests") {
				r := server.requests.listAgentsRequests[0]
				assert.Equal(t, dummyAuthEmail, r.metadata.Get(authEmailKey)[0])
				assert.Equal(t, dummyAuthToken, r.metadata.Get(authAPITokenKey)[0])
				testutil.AssertProtoEqual(t, &syntheticspb.ListAgentsRequest{}, r.data)
			}

			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestClient_GetSyntheticsAgent(t *testing.T) {
	tests := []struct {
		name              string
		requestID         models.ID
		expectedRequest   *syntheticspb.GetAgentRequest
		response          getAgentResponse
		expectedResult    *models.SyntheticsAgent
		expectedError     bool
		expectedErrorCode *codes.Code
	}{
		{
			name:            "status InvalidArgument received",
			requestID:       "13",
			expectedRequest: &syntheticspb.GetAgentRequest{Id: "13"},
			response: getAgentResponse{
				err: status.Errorf(codes.InvalidArgument, codes.InvalidArgument.String()),
			},
			expectedError:     true,
			expectedErrorCode: codePtr(codes.InvalidArgument),
		}, {
			name:            "status NotFound received",
			requestID:       "13",
			expectedRequest: &syntheticspb.GetAgentRequest{Id: "13"},
			response: getAgentResponse{
				err: status.Errorf(codes.NotFound, codes.NotFound.String()),
			},
			expectedError:     true,
			expectedErrorCode: codePtr(codes.NotFound),
		}, {
			name:            "empty response received",
			requestID:       "13",
			expectedRequest: &syntheticspb.GetAgentRequest{Id: "13"},
			response: getAgentResponse{
				data: &syntheticspb.GetAgentResponse{},
			},
			expectedError: true,
		}, {
			name:            "agent returned",
			requestID:       warsawAgentID,
			expectedRequest: &syntheticspb.GetAgentRequest{Id: warsawAgentID},
			response: getAgentResponse{
				data: &syntheticspb.GetAgentResponse{Agent: newWarsawAgentPayload()},
			},
			expectedResult: newWarsawAgent(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			server := newSpySyntheticsServer(t, syntheticsResponses{
				getAgentResponse: tt.response,
			})
			server.Start()
			defer server.Stop()

			client, err := kentikapi.NewClient(
				kentikapi.WithAPIURL("http://"+server.url),
				kentikapi.WithCredentials(dummyAuthEmail, dummyAuthToken),
				kentikapi.WithLogPayloads(),
			)
			require.NoError(t, err)

			// act
			result, err := client.Synthetics.Agents.Get(context.Background(), tt.requestID)

			// assert
			t.Logf("Got result: %+v, err: %v", result, err)
			if tt.expectedError {
				assert.Error(t, err)
				if tt.expectedErrorCode != nil {
					s, ok := status.FromError(err)
					assert.True(t, ok)
					assert.Equal(t, *tt.expectedErrorCode, s.Code())
				}
			} else {
				assert.NoError(t, err)
			}

			if assert.Equal(t, 1, len(server.requests.getAgentRequests), "invalid number of requests") {
				r := server.requests.getAgentRequests[0]
				assert.Equal(t, dummyAuthEmail, r.metadata.Get(authEmailKey)[0])
				assert.Equal(t, dummyAuthToken, r.metadata.Get(authAPITokenKey)[0])
				testutil.AssertProtoEqual(t, tt.expectedRequest, r.data)
			}

			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestClient_UpdateSyntheticsAgent(t *testing.T) {
	tests := []struct {
		name              string
		request           *models.SyntheticsAgent
		expectedRequest   *syntheticspb.UpdateAgentRequest
		response          updateAgentResponse
		expectedResult    *models.SyntheticsAgent
		expectedError     bool
		expectedErrorCode *codes.Code
	}{
		{
			name:            "nil request, status InvalidArgument received",
			request:         nil,
			expectedRequest: &syntheticspb.UpdateAgentRequest{Agent: nil},
			response: updateAgentResponse{
				err: status.Errorf(codes.InvalidArgument, codes.InvalidArgument.String()),
			},
			expectedResult:    nil,
			expectedError:     true,
			expectedErrorCode: codePtr(codes.InvalidArgument),
		}, {
			name:    "empty response received",
			request: newWarsawAgent(),
			expectedRequest: &syntheticspb.UpdateAgentRequest{
				Agent: newWarsawAgentUpdatePayload(),
			},
			response: updateAgentResponse{
				data: &syntheticspb.UpdateAgentResponse{Agent: nil},
			},
			expectedResult: nil,
			expectedError:  true,
		}, {
			name:    "agent updated",
			request: newWarsawAgent(),
			expectedRequest: &syntheticspb.UpdateAgentRequest{
				Agent: newWarsawAgentUpdatePayload(),
			},
			response: updateAgentResponse{
				data: &syntheticspb.UpdateAgentResponse{
					Agent: newWarsawAgentPayload(),
				},
			},
			expectedResult: newWarsawAgent(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			server := newSpySyntheticsServer(t, syntheticsResponses{
				updateAgentResponse: tt.response,
			})
			server.Start()
			defer server.Stop()

			client, err := kentikapi.NewClient(
				kentikapi.WithAPIURL("http://"+server.url),
				kentikapi.WithCredentials(dummyAuthEmail, dummyAuthToken),
				kentikapi.WithLogPayloads(),
			)
			require.NoError(t, err)

			// act
			result, err := client.Synthetics.Agents.Update(context.Background(), tt.request)

			// assert
			t.Logf("Got err: %v", err)
			if tt.expectedError {
				assert.Error(t, err)
				if tt.expectedErrorCode != nil {
					s, ok := status.FromError(err)
					assert.True(t, ok)
					assert.Equal(t, *tt.expectedErrorCode, s.Code())
				}
			} else {
				assert.NoError(t, err)
			}

			if assert.Equal(t, 1, len(server.requests.updateAgentRequests), "invalid number of requests") {
				r := server.requests.updateAgentRequests[0]
				assert.Equal(t, dummyAuthEmail, r.metadata.Get(authEmailKey)[0])
				assert.Equal(t, dummyAuthToken, r.metadata.Get(authAPITokenKey)[0])
				testutil.AssertProtoEqual(t, tt.expectedRequest, r.data)
			}

			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestClient_DeleteSyntheticsAgent(t *testing.T) {
	tests := []struct {
		name              string
		requestID         string
		expectedRequest   *syntheticspb.DeleteAgentRequest
		response          deleteAgentResponse
		expectedError     bool
		expectedErrorCode *codes.Code
	}{
		{
			name:            "status InvalidArgument received",
			requestID:       "13",
			expectedRequest: &syntheticspb.DeleteAgentRequest{Id: "13"},
			response: deleteAgentResponse{
				data: &syntheticspb.DeleteAgentResponse{},
				err:  status.Errorf(codes.InvalidArgument, codes.InvalidArgument.String()),
			},
			expectedError:     true,
			expectedErrorCode: codePtr(codes.InvalidArgument),
		}, {
			name:            "resource deleted",
			requestID:       "13",
			expectedRequest: &syntheticspb.DeleteAgentRequest{Id: "13"},
			response: deleteAgentResponse{
				data: &syntheticspb.DeleteAgentResponse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			server := newSpySyntheticsServer(t, syntheticsResponses{
				deleteAgentResponse: tt.response,
			})
			server.Start()
			defer server.Stop()

			client, err := kentikapi.NewClient(
				kentikapi.WithAPIURL("http://"+server.url),
				kentikapi.WithCredentials(dummyAuthEmail, dummyAuthToken),
				kentikapi.WithLogPayloads(),
			)
			require.NoError(t, err)

			// act
			err = client.Synthetics.Agents.Delete(context.Background(), tt.requestID)

			// assert
			t.Logf("Got err: %v", err)
			if tt.expectedError {
				assert.Error(t, err)
				if tt.expectedErrorCode != nil {
					s, ok := status.FromError(err)
					assert.True(t, ok)
					assert.Equal(t, *tt.expectedErrorCode, s.Code())
				}
			} else {
				assert.NoError(t, err)
			}

			if assert.Equal(t, 1, len(server.requests.deleteAgentRequests), "invalid number of requests") {
				r := server.requests.deleteAgentRequests[0]
				assert.Equal(t, dummyAuthEmail, r.metadata.Get(authEmailKey)[0])
				assert.Equal(t, dummyAuthToken, r.metadata.Get(authAPITokenKey)[0])
				testutil.AssertProtoEqual(t, tt.expectedRequest, r.data)
			}
		})
	}
}

//nolint:dupl
func TestClient_ActivateSyntheticsAgent(t *testing.T) {
	tests := []struct {
		name              string
		requestID         models.ID
		expectedRequest   *syntheticspb.UpdateAgentRequest
		responses         syntheticsResponses
		expectedResult    *models.SyntheticsAgent
		expectedError     bool
		expectedErrorCode *codes.Code
	}{
		{
			name:      "pending agent activated",
			requestID: warsawAgentID,
			expectedRequest: &syntheticspb.UpdateAgentRequest{
				Agent: agentPayloadWithStatus(newWarsawAgentUpdatePayload(), syntheticspb.AgentStatus_AGENT_STATUS_OK),
			},
			responses: syntheticsResponses{
				getAgentResponse: getAgentResponse{
					data: &syntheticspb.GetAgentResponse{
						Agent: agentPayloadWithStatus(newWarsawAgentPayload(), syntheticspb.AgentStatus_AGENT_STATUS_WAIT),
					},
				},
				updateAgentResponse: updateAgentResponse{
					data: &syntheticspb.UpdateAgentResponse{
						Agent: agentPayloadWithStatus(newWarsawAgentPayload(), syntheticspb.AgentStatus_AGENT_STATUS_OK),
					},
				},
			},
			expectedResult: agentWithStatus(newWarsawAgent(), models.AgentStatusOK),
		}, {
			name:      "return error when status not active after activation",
			requestID: warsawAgentID,
			expectedRequest: &syntheticspb.UpdateAgentRequest{
				Agent: agentPayloadWithStatus(newWarsawAgentUpdatePayload(), syntheticspb.AgentStatus_AGENT_STATUS_OK),
			},
			responses: syntheticsResponses{
				getAgentResponse: getAgentResponse{
					data: &syntheticspb.GetAgentResponse{
						Agent: agentPayloadWithStatus(newWarsawAgentPayload(), syntheticspb.AgentStatus_AGENT_STATUS_WAIT),
					},
				},
				updateAgentResponse: updateAgentResponse{
					data: &syntheticspb.UpdateAgentResponse{
						// not active/ok
						Agent: agentPayloadWithStatus(newWarsawAgentPayload(), syntheticspb.AgentStatus_AGENT_STATUS_DELETED),
					},
				},
			},
			expectedResult: nil,
			expectedError:  true,
		}, {
			name:            "skip activation for active agent",
			requestID:       warsawAgentID,
			expectedRequest: nil,
			responses: syntheticsResponses{
				getAgentResponse: getAgentResponse{
					data: &syntheticspb.GetAgentResponse{
						Agent: agentPayloadWithStatus(newWarsawAgentPayload(), syntheticspb.AgentStatus_AGENT_STATUS_OK),
					},
				},
			},
			expectedResult: agentWithStatus(newWarsawAgent(), models.AgentStatusOK),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			server := newSpySyntheticsServer(t, tt.responses)
			server.Start()
			defer server.Stop()

			client, err := kentikapi.NewClient(
				kentikapi.WithAPIURL("http://"+server.url),
				kentikapi.WithCredentials(dummyAuthEmail, dummyAuthToken),
				kentikapi.WithLogPayloads(),
			)
			require.NoError(t, err)

			// act
			result, err := client.Synthetics.Agents.Activate(context.Background(), tt.requestID)

			// assert
			t.Logf("Got err: %v", err)
			if tt.expectedError {
				assert.Error(t, err)
				if tt.expectedErrorCode != nil {
					s, ok := status.FromError(err)
					assert.True(t, ok)
					assert.Equal(t, *tt.expectedErrorCode, s.Code())
				}
			} else {
				assert.NoError(t, err)
			}

			// verify update request only; get request verification is omitted
			if tt.expectedRequest == nil {
				assert.Equal(t, 0, len(server.requests.updateAgentRequests), "invalid number of requests")
			} else if assert.Equal(t, 1, len(server.requests.updateAgentRequests), "invalid number of requests") {
				testutil.AssertProtoEqual(t, tt.expectedRequest, server.requests.updateAgentRequests[0].data)
			}

			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

//nolint:dupl
func TestClient_DeactivateSyntheticsAgent(t *testing.T) {
	tests := []struct {
		name              string
		requestID         models.ID
		expectedRequest   *syntheticspb.UpdateAgentRequest
		responses         syntheticsResponses
		expectedResult    *models.SyntheticsAgent
		expectedError     bool
		expectedErrorCode *codes.Code
	}{
		{
			name:      "active agent deactivated",
			requestID: warsawAgentID,
			expectedRequest: &syntheticspb.UpdateAgentRequest{
				Agent: agentPayloadWithStatus(newWarsawAgentUpdatePayload(), syntheticspb.AgentStatus_AGENT_STATUS_WAIT),
			},
			responses: syntheticsResponses{
				getAgentResponse: getAgentResponse{
					data: &syntheticspb.GetAgentResponse{
						Agent: agentPayloadWithStatus(newWarsawAgentPayload(), syntheticspb.AgentStatus_AGENT_STATUS_OK),
					},
				},
				updateAgentResponse: updateAgentResponse{
					data: &syntheticspb.UpdateAgentResponse{
						Agent: agentPayloadWithStatus(newWarsawAgentPayload(), syntheticspb.AgentStatus_AGENT_STATUS_WAIT),
					},
				},
			},
			expectedResult: agentWithStatus(newWarsawAgent(), models.AgentStatusWait),
		}, {
			name:      "return error when status not pending (waiting) after deactivation",
			requestID: warsawAgentID,
			expectedRequest: &syntheticspb.UpdateAgentRequest{
				Agent: agentPayloadWithStatus(newWarsawAgentUpdatePayload(), syntheticspb.AgentStatus_AGENT_STATUS_WAIT),
			},
			responses: syntheticsResponses{
				getAgentResponse: getAgentResponse{
					data: &syntheticspb.GetAgentResponse{
						Agent: agentPayloadWithStatus(newWarsawAgentPayload(), syntheticspb.AgentStatus_AGENT_STATUS_OK),
					},
				},
				updateAgentResponse: updateAgentResponse{
					data: &syntheticspb.UpdateAgentResponse{
						// not pending/wait
						Agent: agentPayloadWithStatus(newWarsawAgentPayload(), syntheticspb.AgentStatus_AGENT_STATUS_DELETED),
					},
				},
			},
			expectedResult: nil,
			expectedError:  true,
		}, {
			name:            "skip deactivation for pending (waiting) agent",
			requestID:       warsawAgentID,
			expectedRequest: nil,
			responses: syntheticsResponses{
				getAgentResponse: getAgentResponse{
					data: &syntheticspb.GetAgentResponse{
						Agent: agentPayloadWithStatus(newWarsawAgentPayload(), syntheticspb.AgentStatus_AGENT_STATUS_WAIT),
					},
				},
			},
			expectedResult: agentWithStatus(newWarsawAgent(), models.AgentStatusWait),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			server := newSpySyntheticsServer(t, tt.responses)
			server.Start()
			defer server.Stop()

			client, err := kentikapi.NewClient(
				kentikapi.WithAPIURL("http://"+server.url),
				kentikapi.WithCredentials(dummyAuthEmail, dummyAuthToken),
				kentikapi.WithLogPayloads(),
			)
			require.NoError(t, err)

			// act
			result, err := client.Synthetics.Agents.Deactivate(context.Background(), tt.requestID)

			// assert
			t.Logf("Got err: %v", err)
			if tt.expectedError {
				assert.Error(t, err)
				if tt.expectedErrorCode != nil {
					s, ok := status.FromError(err)
					assert.True(t, ok)
					assert.Equal(t, *tt.expectedErrorCode, s.Code())
				}
			} else {
				assert.NoError(t, err)
			}

			// verify update request only; get request verification is omitted
			if tt.expectedRequest == nil {
				assert.Equal(t, 0, len(server.requests.updateAgentRequests), "invalid number of requests")
			} else if assert.Equal(t, 1, len(server.requests.updateAgentRequests), "invalid number of requests") {
				testutil.AssertProtoEqual(t, tt.expectedRequest, server.requests.updateAgentRequests[0].data)
			}

			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

type spySyntheticsServer struct {
	syntheticspb.UnimplementedSyntheticsAdminServiceServer
	server *grpc.Server

	url  string
	done chan struct{}
	t    testing.TB
	// responses to return to the client
	responses syntheticsResponses

	// requests spied by the server
	requests syntheticsRequests
}

type syntheticsRequests struct {
	listAgentsRequests  []listAgentsRequest
	getAgentRequests    []getAgentRequest
	updateAgentRequests []updateAgentRequest
	deleteAgentRequests []deleteAgentRequest
}

type listAgentsRequest struct {
	metadata metadata.MD
	data     *syntheticspb.ListAgentsRequest
}

type getAgentRequest struct {
	metadata metadata.MD
	data     *syntheticspb.GetAgentRequest
}

type updateAgentRequest struct {
	metadata metadata.MD
	data     *syntheticspb.UpdateAgentRequest
}

type deleteAgentRequest struct {
	metadata metadata.MD
	data     *syntheticspb.DeleteAgentRequest
}

type syntheticsResponses struct {
	listAgentsResponse  listAgentsResponse
	getAgentResponse    getAgentResponse
	updateAgentResponse updateAgentResponse
	deleteAgentResponse deleteAgentResponse
}

type listAgentsResponse struct {
	data *syntheticspb.ListAgentsResponse
	err  error
}

type getAgentResponse struct {
	data *syntheticspb.GetAgentResponse
	err  error
}

type updateAgentResponse struct {
	data *syntheticspb.UpdateAgentResponse
	err  error
}

type deleteAgentResponse struct {
	data *syntheticspb.DeleteAgentResponse
	err  error
}

func newSpySyntheticsServer(t testing.TB, r syntheticsResponses) *spySyntheticsServer {
	return &spySyntheticsServer{
		done:      make(chan struct{}),
		t:         t,
		responses: r,
	}
}

func (s *spySyntheticsServer) Start() {
	l, err := net.Listen("tcp", "localhost:0")
	require.NoError(s.t, err)

	s.url = l.Addr().String()
	s.server = grpc.NewServer()
	syntheticspb.RegisterSyntheticsAdminServiceServer(s.server, s)

	go func() {
		err = s.server.Serve(l)
		assert.NoError(s.t, err)
		s.done <- struct{}{}
	}()
}

// Stop blocks until the server is stopped.
func (s *spySyntheticsServer) Stop() {
	s.server.GracefulStop()
	<-s.done
}

func (s *spySyntheticsServer) ListAgents(
	ctx context.Context, req *syntheticspb.ListAgentsRequest,
) (*syntheticspb.ListAgentsResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	s.requests.listAgentsRequests = append(s.requests.listAgentsRequests, listAgentsRequest{
		metadata: md,
		data:     req,
	})

	return s.responses.listAgentsResponse.data, s.responses.listAgentsResponse.err
}

func (s *spySyntheticsServer) GetAgent(
	ctx context.Context, req *syntheticspb.GetAgentRequest,
) (*syntheticspb.GetAgentResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	s.requests.getAgentRequests = append(s.requests.getAgentRequests, getAgentRequest{
		metadata: md,
		data:     req,
	})

	return s.responses.getAgentResponse.data, s.responses.getAgentResponse.err
}

func (s *spySyntheticsServer) UpdateAgent(
	ctx context.Context, req *syntheticspb.UpdateAgentRequest,
) (*syntheticspb.UpdateAgentResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	s.requests.updateAgentRequests = append(s.requests.updateAgentRequests, updateAgentRequest{
		metadata: md,
		data:     req,
	})

	return s.responses.updateAgentResponse.data, s.responses.updateAgentResponse.err
}

func (s *spySyntheticsServer) DeleteAgent(
	ctx context.Context, req *syntheticspb.DeleteAgentRequest,
) (*syntheticspb.DeleteAgentResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	s.requests.deleteAgentRequests = append(s.requests.deleteAgentRequests, deleteAgentRequest{
		metadata: md,
		data:     req,
	})

	return s.responses.deleteAgentResponse.data, s.responses.deleteAgentResponse.err
}

func agentWithStatus(agent *models.SyntheticsAgent, status models.AgentStatus) *models.SyntheticsAgent {
	agent.Status = status
	return agent
}

func newWarsawAgent() *models.SyntheticsAgent {
	return &models.SyntheticsAgent{
		Status:             models.AgentStatusOK,
		Alias:              "Warsaw, Poland",
		SiteID:             "2137",
		LocalIP:            "10.10.10.10",
		IPFamily:           models.IPFamilyDual,
		CloudProvider:      models.CloudProviderGCE,
		CloudRegion:        "europe-central2",
		ID:                 warsawAgentID,
		Type:               models.AgentTypeGlobal,
		SiteName:           "gce-europe-central2",
		IP:                 "34.118.69.79",
		ASN:                396982,
		Latitude:           52.22977,
		Longitude:          21.01178,
		City:               "Warsaw",
		Region:             "Mazowieckie",
		Country:            "PL",
		Version:            "1.2.0",
		OS:                 "Linux probe-1-waw-1 4.9.0-18-amd64 #1 SMP Debian 4.9.303-1 (2022-03-07) x86_64",
		ImplementationType: models.AgentImplementationTypeRust,
		LastAuthed:         time.Date(2020, time.July, 9, 21, 37, 0, 826*1000000, time.UTC),
		TestIDs:            []string{"13", "133", "1337"},
	}
}

func newMoscowAgent() *models.SyntheticsAgent {
	return &models.SyntheticsAgent{
		Status:             models.AgentStatusOK,
		Alias:              "Moscow, Russia",
		SiteID:             "0",
		LocalIP:            "",
		IPFamily:           models.IPFamilyV4,
		CloudProvider:      "",
		CloudRegion:        "",
		ID:                 "7667",
		Type:               models.AgentTypePrivate,
		SiteName:           "Tencent,CN (132203)",
		IP:                 "162.62.11.117",
		ASN:                132203,
		Latitude:           55.75222,
		Longitude:          37.61556,
		City:               "Moscow",
		Region:             "Moskva",
		Country:            "RU",
		Version:            "0.3.9",
		OS:                 "",
		ImplementationType: models.AgentImplementationTypeNode,
		LastAuthed:         time.Date(2022, time.February, 24, 6, 48, 0, 0, time.UTC),
		TestIDs:            nil,
	}
}

func agentPayloadWithStatus(a *syntheticspb.Agent, s syntheticspb.AgentStatus) *syntheticspb.Agent {
	a.Status = s
	return a
}

// newWarsawAgentUpdatePayload returns Warsaw agent payload without read-only fields (except ID).
func newWarsawAgentUpdatePayload() *syntheticspb.Agent {
	return &syntheticspb.Agent{
		Id:            warsawAgentID,
		Status:        syntheticspb.AgentStatus_AGENT_STATUS_OK,
		Alias:         "Warsaw, Poland",
		Family:        syntheticspb.IPFamily_IP_FAMILY_DUAL,
		SiteId:        "2137",
		LocalIp:       "10.10.10.10",
		CloudRegion:   "europe-central2",
		CloudProvider: "gcp",
	}
}

func newWarsawAgentPayload() *syntheticspb.Agent {
	return &syntheticspb.Agent{
		Id:            warsawAgentID,
		SiteName:      "gce-europe-central2",
		Status:        syntheticspb.AgentStatus_AGENT_STATUS_OK,
		Alias:         "Warsaw, Poland",
		Type:          "global",
		Os:            "Linux probe-1-waw-1 4.9.0-18-amd64 #1 SMP Debian 4.9.303-1 (2022-03-07) x86_64",
		Ip:            "34.118.69.79",
		Lat:           52.22977,
		Long:          21.01178,
		LastAuthed:    timestamppb.New(time.Date(2020, time.July, 9, 21, 37, 0, 826*1000000, time.UTC)),
		Family:        syntheticspb.IPFamily_IP_FAMILY_DUAL,
		Asn:           396982,
		SiteId:        "2137",
		Version:       "1.2.0",
		City:          "Warsaw",
		Region:        "Mazowieckie",
		Country:       "PL",
		TestIds:       []string{"13", "133", "1337"},
		LocalIp:       "10.10.10.10",
		CloudRegion:   "europe-central2",
		CloudProvider: "gcp",
		AgentImpl:     syntheticspb.ImplementType_IMPLEMENT_TYPE_RUST,
	}
}

func newMoscowAgentPayload() *syntheticspb.Agent {
	return &syntheticspb.Agent{
		Id:            "7667",
		SiteName:      "Tencent,CN (132203)",
		Status:        syntheticspb.AgentStatus_AGENT_STATUS_OK,
		Alias:         "Moscow, Russia",
		Type:          "private",
		Os:            "",
		Ip:            "162.62.11.117",
		Lat:           55.75222,
		Long:          37.61556,
		LastAuthed:    timestamppb.New(time.Date(2022, time.February, 24, 6, 48, 0, 0, time.UTC)),
		Family:        syntheticspb.IPFamily_IP_FAMILY_V4,
		Asn:           132203,
		SiteId:        "0",
		Version:       "0.3.9",
		City:          "Moscow",
		Region:        "Moskva",
		Country:       "RU",
		TestIds:       []string{},
		LocalIp:       "",
		CloudRegion:   "",
		CloudProvider: "",
		AgentImpl:     syntheticspb.ImplementType_IMPLEMENT_TYPE_NODE,
	}
}
