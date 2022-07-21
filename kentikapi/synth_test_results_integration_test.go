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
	"github.com/kentik/community_sdk_golang/kentikapi/synthetics"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestClient_Synthetics_GetResultsForTests(t *testing.T) {
	tests := []struct {
		name            string
		response        getResultsForTestsResponse
		expectedResult  []synthetics.TestResults
		expectedError   bool
		errorPredicates []func(error) bool
	}{
		{
			name: "status InvalidArgument received",
			response: getResultsForTestsResponse{
				err: status.Errorf(codes.InvalidArgument, codes.InvalidArgument.String()),
			},
			expectedResult:  nil,
			expectedError:   true,
			errorPredicates: []func(error) bool{kentikapi.IsInvalidRequestError},
		}, {
			name: "empty response received",
			response: getResultsForTestsResponse{
				data: &syntheticspb.GetResultsForTestsResponse{},
			},
			expectedResult: nil,
			expectedError:  false,
		}, {
			name: "test results with empty agent results list received",
			response: getResultsForTestsResponse{
				data: &syntheticspb.GetResultsForTestsResponse{
					Results: []*syntheticspb.TestResults{
						{
							TestId: "40785",
							Time:   timestamppb.New(time.Date(2022, time.July, 19, 15, 0, 0, 0, time.UTC)),
							Health: "healthy",
							Agents: nil,
						},
					},
				},
			},
			expectedResult: []synthetics.TestResults{
				{
					TestID: "40785",
					Time:   time.Date(2022, time.July, 19, 15, 0, 0, 0, time.UTC),
					Health: synthetics.HealthHealthy,
					Agents: nil,
				},
			},
		}, {
			name: "test results with nil agent results received",
			response: getResultsForTestsResponse{
				data: &syntheticspb.GetResultsForTestsResponse{
					Results: []*syntheticspb.TestResults{
						{
							TestId: "40785",
							Time:   timestamppb.New(time.Date(2022, time.July, 19, 15, 0, 0, 0, time.UTC)),
							Health: "healthy",
							Agents: []*syntheticspb.AgentResults{nil},
						},
					},
				},
			},
			expectedResult: []synthetics.TestResults{
				{
					TestID: "40785",
					Time:   time.Date(2022, time.July, 19, 15, 0, 0, 0, time.UTC),
					Health: synthetics.HealthHealthy,
					// gRPC bindings initialize syntheticspb.AgentResults anyway
					Agents: []synthetics.AgentResults{{
						AgentID: "",
						Health:  "",
						Tasks:   nil,
					}},
				},
			},
		}, {
			name: "3x3 network mesh test results received",
			response: getResultsForTestsResponse{
				data: testResultsPayloadFromJSON(t, testResultsResponseJSON),
			},
			expectedResult: newTestResults(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			server := newSpySyntheticsServer(t, syntheticsResponses{
				getResultsForTestsResponse: tt.response,
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
			result, err := client.Synthetics.GetResultsForTests(context.Background(), newGetResultsForTestsRequest())

			// assert
			t.Logf("Got result: %+v, err: %v", result, err)
			if tt.expectedError {
				assert.Error(t, err)
				for _, isErr := range tt.errorPredicates {
					assert.True(t, isErr(err))
				}
			} else {
				assert.NoError(t, err)
			}

			if assert.Equal(t, 1, len(server.requests.getResultsForTestsRequests), "invalid number of requests") {
				r := server.requests.getResultsForTestsRequests[0]
				assert.Equal(t, dummyAuthEmail, r.metadata.Get(authEmailKey)[0])
				assert.Equal(t, dummyAuthToken, r.metadata.Get(authAPITokenKey)[0])
				testutil.AssertProtoEqual(t,
					newGetResultsForTestsRequestPayload(),
					r.data,
				)
			}

			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func testResultsPayloadFromJSON(t *testing.T, payloadJSON string) *syntheticspb.GetResultsForTestsResponse {
	var r syntheticspb.GetResultsForTestsResponse
	err := protojson.Unmarshal([]byte(payloadJSON), &r)
	require.NoError(t, err)
	return &r
}

func newGetResultsForTestsRequest() synthetics.GetResultsForTestsRequest {
	return synthetics.GetResultsForTestsRequest{
		TestIDs:   []string{"40785"},
		StartTime: time.Date(2022, time.July, 19, 14, 11, 22, 123456789, time.UTC),
		EndTime:   time.Date(2022, time.July, 19, 20, 11, 22, 123456789, time.UTC),
		AgentIDs:  []models.ID{"598", "608", "631"},
		Targets: []net.IP{
			net.ParseIP("23.92.28.124"),
			net.ParseIP("54.152.163.64"),
			net.ParseIP("95.179.136.58"),
		},
	}
}

func newGetResultsForTestsRequestPayload() *syntheticspb.GetResultsForTestsRequest {
	return &syntheticspb.GetResultsForTestsRequest{
		Ids:       []string{"40785"},
		StartTime: timestamppb.New(time.Date(2022, time.July, 19, 14, 11, 22, 123456789, time.UTC)),
		EndTime:   timestamppb.New(time.Date(2022, time.July, 19, 20, 11, 22, 123456789, time.UTC)),
		AgentIds:  []string{"598", "608", "631"},
		Targets:   []string{"23.92.28.124", "54.152.163.64", "95.179.136.58"},
	}
}

const testResultsResponseJSON = `
{
    "results": [
        {
            "testId": "40785",
            "health": "warning",
            "agents": [
                {
                    "agentId": "598",
                    "health": "warning",
                    "tasks": [
                        {
                            "ping": {
                                "target": "23.92.28.124",
                                "packetLoss": {
                                    "current": 0.66,
                                    "health": "warning"
                                },
                                "latency": {
                                    "current": 13499,
                                    "rollingAvg": 13473,
                                    "rollingStddev": 36,
                                    "health": "healthy"
                                },
                                "jitter": {
                                    "current": 391,
                                    "rollingAvg": 381,
                                    "rollingStddev": 13,
                                    "health": "healthy"
                                },
                                "dstIp": "23.92.28.124"
                            },
                            "health": "healthy"
                        },
                        {
                            "ping": {
                                "target": "95.179.136.58",
                                "packetLoss": {
                                    "current": 0,
                                    "health": "healthy"
                                },
                                "latency": {
                                    "current": 93065,
                                    "rollingAvg": 93031,
                                    "rollingStddev": 48,
                                    "health": "healthy"
                                },
                                "jitter": {
                                    "current": 289,
                                    "rollingAvg": 257,
                                    "rollingStddev": 46,
                                    "health": "healthy"
                                },
                                "dstIp": "95.179.136.58"
                            },
                            "health": "healthy"
                        }
                    ]
                },
                {
                    "agentId": "608",
                    "health": "healthy",
                    "tasks": [
                        {
                            "ping": {
                                "target": "54.152.163.64",
                                "packetLoss": {
                                    "current": 0,
                                    "health": "healthy"
                                },
                                "latency": {
                                    "current": 13517,
                                    "rollingAvg": 13482,
                                    "rollingStddev": 49,
                                    "health": "healthy"
                                },
                                "jitter": {
                                    "current": 554,
                                    "rollingAvg": 484,
                                    "rollingStddev": 99,
                                    "health": "healthy"
                                },
                                "dstIp": "54.152.163.64"
                            },
                            "health": "healthy"
                        },
                        {
                            "ping": {
                                "target": "95.179.136.58",
                                "packetLoss": {
                                    "current": 0,
                                    "health": "healthy"
                                },
                                "latency": {
                                    "current": 109076,
                                    "rollingAvg": 109045,
                                    "rollingStddev": 43,
                                    "health": "healthy"
                                },
                                "jitter": {
                                    "current": 383,
                                    "rollingAvg": 351,
                                    "rollingStddev": 44,
                                    "health": "healthy"
                                },
                                "dstIp": "95.179.136.58"
                            },
                            "health": "healthy"
                        }
                    ]
                },
                {
                    "agentId": "631",
                    "health": "healthy",
                    "tasks": [
                        {
                            "ping": {
                                "target": "23.92.28.124",
                                "packetLoss": {
                                    "current": 0,
                                    "health": "healthy"
                                },
                                "latency": {
                                    "current": 109007,
                                    "rollingAvg": 109015,
                                    "rollingStddev": 10,
                                    "health": "healthy"
                                },
                                "jitter": {
                                    "current": 225,
                                    "rollingAvg": 226,
                                    "rollingStddev": 2,
                                    "health": "healthy"
                                },
                                "dstIp": "23.92.28.124"
                            },
                            "health": "healthy"
                        },
                        {
                            "ping": {
                                "target": "54.152.163.64",
                                "packetLoss": {
                                    "current": 0,
                                    "health": "healthy"
                                },
                                "latency": {
                                    "current": 93085,
                                    "rollingAvg": 93045,
                                    "rollingStddev": 55,
                                    "health": "healthy"
                                },
                                "jitter": {
                                    "current": 408,
                                    "rollingAvg": 308,
                                    "rollingStddev": 140,
                                    "health": "healthy"
                                },
                                "dstIp": "54.152.163.64"
                            },
                            "health": "healthy"
                        }
                    ]
                }
            ],
            "time": "2022-07-19T15:00:00Z"
        }
    ]
}
`

func newTestResults() []synthetics.TestResults {
	// nolint: dupl // ping results are similar but different
	return []synthetics.TestResults{
		{
			TestID: "40785",
			Time:   time.Date(2022, time.July, 19, 15, 0, 0, 0, time.UTC),
			Health: synthetics.HealthWarning,
			Agents: []synthetics.AgentResults{
				{
					AgentID: "598",
					Health:  synthetics.HealthWarning,
					Tasks: []synthetics.TaskResults{
						{
							Health:   synthetics.HealthHealthy,
							TaskType: synthetics.TaskTypePing,
							Task: synthetics.PingResults{
								Target: "23.92.28.124",
								PacketLoss: synthetics.PacketLossData{
									Current: 0.66,
									Health:  synthetics.HealthWarning,
								},
								Latency: synthetics.MetricData{
									Current:       13499 * time.Microsecond,
									RollingAvg:    13473 * time.Microsecond,
									RollingStdDev: 36 * time.Microsecond,
									Health:        synthetics.HealthHealthy,
								},
								Jitter: synthetics.MetricData{
									Current:       391 * time.Microsecond,
									RollingAvg:    381 * time.Microsecond,
									RollingStdDev: 13 * time.Microsecond,
									Health:        synthetics.HealthHealthy,
								},
								DstIP: net.ParseIP("23.92.28.124"),
							},
						},
						{
							Health:   synthetics.HealthHealthy,
							TaskType: synthetics.TaskTypePing,
							Task: synthetics.PingResults{
								Target: "95.179.136.58",
								PacketLoss: synthetics.PacketLossData{
									Current: 0,
									Health:  synthetics.HealthHealthy,
								},
								Latency: synthetics.MetricData{
									Current:       93065 * time.Microsecond,
									RollingAvg:    93031 * time.Microsecond,
									RollingStdDev: 48 * time.Microsecond,
									Health:        synthetics.HealthHealthy,
								},
								Jitter: synthetics.MetricData{
									Current:       289 * time.Microsecond,
									RollingAvg:    257 * time.Microsecond,
									RollingStdDev: 46 * time.Microsecond,
									Health:        synthetics.HealthHealthy,
								},
								DstIP: net.ParseIP("95.179.136.58"),
							},
						},
					},
				},
				{
					AgentID: "608",
					Health:  synthetics.HealthHealthy,
					Tasks: []synthetics.TaskResults{
						{
							Health:   synthetics.HealthHealthy,
							TaskType: synthetics.TaskTypePing,
							Task: synthetics.PingResults{
								Target: "54.152.163.64",
								PacketLoss: synthetics.PacketLossData{
									Current: 0,
									Health:  synthetics.HealthHealthy,
								},
								Latency: synthetics.MetricData{
									Current:       13517 * time.Microsecond,
									RollingAvg:    13482 * time.Microsecond,
									RollingStdDev: 49 * time.Microsecond,
									Health:        synthetics.HealthHealthy,
								},
								Jitter: synthetics.MetricData{
									Current:       554 * time.Microsecond,
									RollingAvg:    484 * time.Microsecond,
									RollingStdDev: 99 * time.Microsecond,
									Health:        synthetics.HealthHealthy,
								},
								DstIP: net.ParseIP("54.152.163.64"),
							},
						},
						{
							Health:   synthetics.HealthHealthy,
							TaskType: synthetics.TaskTypePing,
							Task: synthetics.PingResults{
								Target: "95.179.136.58",
								PacketLoss: synthetics.PacketLossData{
									Current: 0,
									Health:  synthetics.HealthHealthy,
								},
								Latency: synthetics.MetricData{
									Current:       109076 * time.Microsecond,
									RollingAvg:    109045 * time.Microsecond,
									RollingStdDev: 43 * time.Microsecond,
									Health:        synthetics.HealthHealthy,
								},
								Jitter: synthetics.MetricData{
									Current:       383 * time.Microsecond,
									RollingAvg:    351 * time.Microsecond,
									RollingStdDev: 44 * time.Microsecond,
									Health:        synthetics.HealthHealthy,
								},
								DstIP: net.ParseIP("95.179.136.58"),
							},
						},
					},
				},
				{
					AgentID: "631",
					Health:  synthetics.HealthHealthy,
					Tasks: []synthetics.TaskResults{
						{
							Health:   synthetics.HealthHealthy,
							TaskType: synthetics.TaskTypePing,
							Task: synthetics.PingResults{
								Target: "23.92.28.124",
								PacketLoss: synthetics.PacketLossData{
									Current: 0,
									Health:  synthetics.HealthHealthy,
								},
								Latency: synthetics.MetricData{
									Current:       109007 * time.Microsecond,
									RollingAvg:    109015 * time.Microsecond,
									RollingStdDev: 10 * time.Microsecond,
									Health:        synthetics.HealthHealthy,
								},
								Jitter: synthetics.MetricData{
									Current:       225 * time.Microsecond,
									RollingAvg:    226 * time.Microsecond,
									RollingStdDev: 2 * time.Microsecond,
									Health:        synthetics.HealthHealthy,
								},
								DstIP: net.ParseIP("23.92.28.124"),
							},
						},
						{
							Health:   synthetics.HealthHealthy,
							TaskType: synthetics.TaskTypePing,
							Task: synthetics.PingResults{
								Target: "54.152.163.64",
								PacketLoss: synthetics.PacketLossData{
									Current: 0,
									Health:  synthetics.HealthHealthy,
								},
								Latency: synthetics.MetricData{
									Current:       93085 * time.Microsecond,
									RollingAvg:    93045 * time.Microsecond,
									RollingStdDev: 55 * time.Microsecond,
									Health:        synthetics.HealthHealthy,
								},
								Jitter: synthetics.MetricData{
									Current:       408 * time.Microsecond,
									RollingAvg:    308 * time.Microsecond,
									RollingStdDev: 140 * time.Microsecond,
									Health:        synthetics.HealthHealthy,
								},
								DstIP: net.ParseIP("54.152.163.64"),
							},
						},
					},
				},
			},
		},
	}
}
