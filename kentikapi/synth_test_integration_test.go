package kentikapi_test

import (
	"context"
	"net"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/AlekSi/pointer"
	syntheticspb "github.com/kentik/api-schema-public/gen/go/kentik/synthetics/v202202"
	"github.com/kentik/community_sdk_golang/kentikapi"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/testutil"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/kentik/community_sdk_golang/kentikapi/synthetics"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	ipTestID          = "1001"
	networkGridTestID = "1002"
	hostnameTestID    = "1003"
	agentTestID       = "1004"
	networkMeshTestID = "1005"
	flowTestID        = "1006"
	urlTestID         = "1007"
	pageLoadTestID    = "1008"
	dnsTestID         = "1009"
	dnsGridTestID     = "1010"
	bgpMonitorTestID  = "1011"
	transactionTestID = "1012"
)

func TestClient_Synthetics_GetAllTests(t *testing.T) {
	tests := []struct {
		name            string
		response        listTestsResponse
		expectedResult  *synthetics.GetAllTestsResponse
		expectedError   bool
		errorPredicates []func(error) bool
	}{
		{
			name: "status InvalidArgument received",
			response: listTestsResponse{
				err: status.Errorf(codes.InvalidArgument, codes.InvalidArgument.String()),
			},
			expectedError:   true,
			errorPredicates: []func(error) bool{kentikapi.IsInvalidRequestError},
		}, {
			name: "empty response received",
			response: listTestsResponse{
				data: &syntheticspb.ListTestsResponse{},
			},
			expectedResult: &synthetics.GetAllTestsResponse{
				Tests:             nil,
				InvalidTestsCount: 0,
			},
		}, {
			name: "no tests received",
			response: listTestsResponse{
				data: &syntheticspb.ListTestsResponse{
					Tests:        []*syntheticspb.Test{},
					InvalidCount: 0,
				},
			},
			expectedResult: &synthetics.GetAllTestsResponse{
				Tests:             nil,
				InvalidTestsCount: 0,
			},
		}, {
			name: "multiple tests received",
			response: listTestsResponse{
				data: &syntheticspb.ListTestsResponse{
					Tests: []*syntheticspb.Test{
						newIPTestPayload(),
						newNetworkGridTestPayload(),
						newHostnameTestPayload(),
						newAgentTestPayload(),
						newNetworkMeshTestPayload(),
						newFlowTestPayload(),
						newURLTestPayload(),
						newPageLoadTestPayload(),
						newDNSTestPayload(),
						newDNSGridTestPayload(),
						newBGPMonitorTestPayload(),
						newTransactionTestPayload(),
					},
					InvalidCount: 1,
				},
			},
			expectedResult: &synthetics.GetAllTestsResponse{
				Tests: []synthetics.Test{
					*newIPTest(),
					*newNetworkGridTest(),
					*newHostnameTest(),
					*newAgentTest(),
					*newNetworkMeshTest(),
					*newFlowTest(),
					*newURLTest(),
					*newPageLoadTest(),
					*newDNSTest(),
					*newDNSGridTest(),
					// BGP monitor test should be silently ignored
					// transaction test should be silently ignored
				},
				InvalidTestsCount: 1,
			},
		}, {
			name: "2 tests received - one nil",
			response: listTestsResponse{
				data: &syntheticspb.ListTestsResponse{
					Tests: []*syntheticspb.Test{
						newIPTestPayload(),
						nil,
					},
					InvalidCount: 0,
				},
			},
			expectedError: true, // InvalidResponse
		},
	}
	//nolint:dupl
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			server := newSpySyntheticsServer(t, syntheticsResponses{
				listTestsResponse: tt.response,
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
			result, err := client.Synthetics.GetAllTests(context.Background())

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

			if assert.Equal(t, 1, len(server.requests.listTestsRequests), "invalid number of requests") {
				r := server.requests.listTestsRequests[0]
				assert.Equal(t, dummyAuthEmail, r.metadata.Get(authEmailKey)[0])
				assert.Equal(t, dummyAuthToken, r.metadata.Get(authAPITokenKey)[0])
				testutil.AssertProtoEqual(t, &syntheticspb.ListTestsRequest{}, r.data)
			}

			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestClient_Synthetics_GetTest(t *testing.T) {
	tests := []struct {
		name            string
		requestID       models.ID
		expectedRequest *syntheticspb.GetTestRequest
		response        getTestResponse
		expectedResult  *synthetics.Test
		expectedError   bool
		errorPredicates []func(error) bool
	}{
		{
			name:            "status InvalidArgument received",
			requestID:       "13",
			expectedRequest: &syntheticspb.GetTestRequest{Id: "13"},
			response: getTestResponse{
				err: status.Errorf(codes.InvalidArgument, codes.InvalidArgument.String()),
			},
			expectedError:   true,
			errorPredicates: []func(error) bool{kentikapi.IsInvalidRequestError},
		}, {
			name:            "status NotFound received",
			requestID:       "13",
			expectedRequest: &syntheticspb.GetTestRequest{Id: "13"},
			response: getTestResponse{
				err: status.Errorf(codes.NotFound, codes.NotFound.String()),
			},
			expectedError:   true,
			errorPredicates: []func(error) bool{kentikapi.IsNotFoundError},
		}, {
			name:            "empty response received",
			requestID:       "13",
			expectedRequest: &syntheticspb.GetTestRequest{Id: "13"},
			response: getTestResponse{
				data: &syntheticspb.GetTestResponse{},
			},
			expectedError: true, // InvalidResponse
		}, {
			name:            "ip test returned",
			requestID:       ipTestID,
			expectedRequest: &syntheticspb.GetTestRequest{Id: ipTestID},
			response: getTestResponse{
				data: &syntheticspb.GetTestResponse{Test: newIPTestPayload()},
			},
			expectedResult: newIPTest(),
		}, {
			name:            "network grid test returned",
			requestID:       networkGridTestID,
			expectedRequest: &syntheticspb.GetTestRequest{Id: networkGridTestID},
			response: getTestResponse{
				data: &syntheticspb.GetTestResponse{Test: newNetworkGridTestPayload()},
			},
			expectedResult: newNetworkGridTest(),
		}, {
			name:            "hostname test returned",
			requestID:       hostnameTestID,
			expectedRequest: &syntheticspb.GetTestRequest{Id: hostnameTestID},
			response: getTestResponse{
				data: &syntheticspb.GetTestResponse{Test: newHostnameTestPayload()},
			},
			expectedResult: newHostnameTest(),
		}, {
			name:            "agent test returned",
			requestID:       agentTestID,
			expectedRequest: &syntheticspb.GetTestRequest{Id: agentTestID},
			response: getTestResponse{
				data: &syntheticspb.GetTestResponse{Test: newAgentTestPayload()},
			},
			expectedResult: newAgentTest(),
		}, {
			name:            "network mesh test returned",
			requestID:       networkMeshTestID,
			expectedRequest: &syntheticspb.GetTestRequest{Id: networkMeshTestID},
			response: getTestResponse{
				data: &syntheticspb.GetTestResponse{Test: newNetworkMeshTestPayload()},
			},
			expectedResult: newNetworkMeshTest(),
		}, {
			name:            "flow test returned",
			requestID:       flowTestID,
			expectedRequest: &syntheticspb.GetTestRequest{Id: flowTestID},
			response: getTestResponse{
				data: &syntheticspb.GetTestResponse{Test: newFlowTestPayload()},
			},
			expectedResult: newFlowTest(),
		}, {
			name:            "URL test returned",
			requestID:       urlTestID,
			expectedRequest: &syntheticspb.GetTestRequest{Id: urlTestID},
			response: getTestResponse{
				data: &syntheticspb.GetTestResponse{Test: newURLTestPayload()},
			},
			expectedResult: newURLTest(),
		}, {
			name:            "page load test returned",
			requestID:       pageLoadTestID,
			expectedRequest: &syntheticspb.GetTestRequest{Id: pageLoadTestID},
			response: getTestResponse{
				data: &syntheticspb.GetTestResponse{Test: newPageLoadTestPayload()},
			},
			expectedResult: newPageLoadTest(),
		}, {
			name:            "DNS test returned",
			requestID:       dnsTestID,
			expectedRequest: &syntheticspb.GetTestRequest{Id: dnsTestID},
			response: getTestResponse{
				data: &syntheticspb.GetTestResponse{Test: newDNSTestPayload()},
			},
			expectedResult: newDNSTest(),
		}, {
			name:            "DNS grid test returned",
			requestID:       dnsGridTestID,
			expectedRequest: &syntheticspb.GetTestRequest{Id: dnsGridTestID},
			response: getTestResponse{
				data: &syntheticspb.GetTestResponse{Test: newDNSGridTestPayload()},
			},
			expectedResult: newDNSGridTest(),
		}, {
			name:            "BGP monitor test returned",
			requestID:       bgpMonitorTestID,
			expectedRequest: &syntheticspb.GetTestRequest{Id: bgpMonitorTestID},
			response: getTestResponse{
				data: &syntheticspb.GetTestResponse{Test: newBGPMonitorTestPayload()},
			},
			expectedResult: nil,
			expectedError:  true, // InvalidResponse
		}, {
			name:            "transaction test returned",
			requestID:       transactionTestID,
			expectedRequest: &syntheticspb.GetTestRequest{Id: transactionTestID},
			response: getTestResponse{
				data: &syntheticspb.GetTestResponse{Test: newTransactionTestPayload()},
			},
			expectedResult: nil,
			expectedError:  true, // InvalidResponse
		},
	}
	//nolint:dupl
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			server := newSpySyntheticsServer(t, syntheticsResponses{
				getTestResponse: tt.response,
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
			result, err := client.Synthetics.GetTest(context.Background(), tt.requestID)

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

			if assert.Equal(t, 1, len(server.requests.getTestRequests), "invalid number of requests") {
				r := server.requests.getTestRequests[0]
				assert.Equal(t, dummyAuthEmail, r.metadata.Get(authEmailKey)[0])
				assert.Equal(t, dummyAuthToken, r.metadata.Get(authAPITokenKey)[0])
				testutil.AssertProtoEqual(t, tt.expectedRequest, r.data)
			}

			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestClient_Synthetics_CreateTest(t *testing.T) {
	tests := []struct {
		name            string
		request         *synthetics.Test
		expectedRequest *syntheticspb.CreateTestRequest
		response        createTestResponse
		expectedResult  *synthetics.Test
		expectedError   bool
		errorPredicates []func(error) bool
	}{
		{
			name:            "nil request",
			request:         nil,
			expectedRequest: nil,
			expectedResult:  nil,
			expectedError:   true,
			errorPredicates: []func(error) bool{kentikapi.IsInvalidRequestError},
		}, {
			name:    "empty response received",
			request: newIPTest(),
			expectedRequest: &syntheticspb.CreateTestRequest{
				Test: testWithoutReadOnlyFields(newIPTestPayload()),
			},
			response: createTestResponse{
				data: &syntheticspb.CreateTestResponse{Test: nil},
			},
			expectedResult: nil,
			expectedError:  true, // InvalidResponse
		}, {
			name: "minimal hostname test created",
			request: synthetics.NewHostnameTest(synthetics.HostnameTestRequiredFields{
				BasePingTraceTestRequiredFields: synthetics.BasePingTraceTestRequiredFields{
					BaseTestRequiredFields: synthetics.BaseTestRequiredFields{
						Name:     "minimal-hostname-test",
						AgentIDs: []string{"817", "818", "819"},
					},
					Ping: synthetics.PingSettingsRequiredFields{
						Timeout:  10 * time.Second,
						Count:    10,
						Protocol: synthetics.PingProtocolTCP,
						Port:     65535,
					},
					Traceroute: synthetics.TracerouteSettingsRequiredFields{
						Timeout:  59999 * time.Millisecond,
						Count:    5,
						Delay:    100 * time.Millisecond,
						Protocol: synthetics.TracerouteProtocolUDP,
						Limit:    255,
					},
				},
				Definition: synthetics.TestDefinitionHostnameRequiredFields{
					Target: "www.example.com",
				},
			}),
			expectedRequest: &syntheticspb.CreateTestRequest{
				Test: &syntheticspb.Test{
					Name:   "minimal-hostname-test",
					Type:   "hostname",
					Status: syntheticspb.TestStatus_TEST_STATUS_ACTIVE,
					Settings: &syntheticspb.TestSettings{
						Definition: &syntheticspb.TestSettings_Hostname{
							Hostname: &syntheticspb.HostnameTest{
								Target: "www.example.com",
							},
						},
						AgentIds: []string{"817", "818", "819"},
						Tasks:    []string{"ping", "traceroute"},
						HealthSettings: &syntheticspb.HealthSettings{
							UnhealthySubtestThreshold: 1,
						},
						Ping: &syntheticspb.TestPingSettings{
							Count:    10,
							Protocol: "tcp",
							Port:     65535,
							Timeout:  10000,
							Delay:    0,
						},
						Trace: &syntheticspb.TestTraceSettings{
							Count:    5,
							Protocol: "udp",
							Port:     33434,
							Timeout:  59999,
							Limit:    255,
							Delay:    100,
						},
						Period: 60,
						Family: syntheticspb.IPFamily_IP_FAMILY_DUAL,
					},
				},
			},
			response: createTestResponse{
				// a minimal hostname test data returned by Kentik API
				data: &syntheticspb.CreateTestResponse{
					Test: &syntheticspb.Test{
						Name:   "minimal-hostname-test",
						Type:   "hostname",
						Status: syntheticspb.TestStatus_TEST_STATUS_ACTIVE,
						Settings: &syntheticspb.TestSettings{
							Definition: &syntheticspb.TestSettings_Hostname{
								Hostname: &syntheticspb.HostnameTest{
									Target: "www.example.com",
								},
							},
							AgentIds: []string{"817", "818", "819"},
							Tasks:    []string{"ping", "traceroute"},
							HealthSettings: &syntheticspb.HealthSettings{
								UnhealthySubtestThreshold: 1,
								Activation: &syntheticspb.ActivationSettings{
									GracePeriod: "2",
									TimeUnit:    "m",
									TimeWindow:  "5",
									Times:       "3",
								},
							},
							Ping: &syntheticspb.TestPingSettings{
								Count:    10,
								Protocol: "tcp",
								Port:     65535,
								Timeout:  10000,
								Delay:    0,
							},
							Trace: &syntheticspb.TestTraceSettings{
								Count:    5,
								Protocol: "udp",
								Port:     33434,
								Timeout:  59999,
								Limit:    255,
								Delay:    100,
							},
							Period: 60,
							Family: syntheticspb.IPFamily_IP_FAMILY_DUAL,
						},
						Id:    hostnameTestID,
						Cdate: timestamppb.New(time.Date(2022, time.April, 6, 9, 43, 39, 324*1000000, time.UTC)),
						Edate: timestamppb.New(time.Date(2022, time.April, 6, 9, 43, 39, 835*1000000, time.UTC)),
						CreatedBy: &syntheticspb.UserInfo{
							Id:       "4321",
							Email:    "joe.doe@example.com",
							FullName: "Joe Doe",
						},
						LastUpdatedBy: nil,
					},
				},
			},
			expectedResult: &synthetics.Test{
				Name:       "minimal-hostname-test",
				Type:       synthetics.TestTypeHostname,
				Status:     synthetics.TestStatusActive,
				UpdateDate: pointer.ToTime(time.Date(2022, time.April, 6, 9, 43, 39, 835*1000000, time.UTC)),
				Settings: synthetics.TestSettings{
					Definition: &synthetics.TestDefinitionHostname{
						Target: "www.example.com",
					},
					AgentIDs: []string{"817", "818", "819"},
					Period:   60 * time.Second,
					Family:   synthetics.IPFamilyDual,
					Health: synthetics.HealthSettings{
						UnhealthySubtestThreshold: 1,
						AlarmActivation: &synthetics.AlarmActivationSettings{
							TimeWindow:  5 * time.Minute,
							Times:       3,
							GracePeriod: 2,
						},
					},
					Ping: &synthetics.PingSettings{
						Timeout:  10 * time.Second,
						Count:    10,
						Delay:    0,
						Protocol: synthetics.PingProtocolTCP,
						Port:     65535,
					},
					Traceroute: &synthetics.TracerouteSettings{
						Timeout:  59999 * time.Millisecond,
						Count:    5,
						Delay:    100 * time.Millisecond,
						Protocol: synthetics.TracerouteProtocolUDP,
						Port:     33434,
						Limit:    255,
					},
					Tasks: []synthetics.TaskType{synthetics.TaskTypePing, synthetics.TaskTypeTraceroute},
				},
				ID:         hostnameTestID,
				CreateDate: time.Date(2022, time.April, 6, 9, 43, 39, 324*1000000, time.UTC),
				CreatedBy: synthetics.UserInfo{
					ID:       "4321",
					Email:    "joe.doe@example.com",
					FullName: "Joe Doe",
				},
				LastUpdatedBy: nil,
			},
		}, {
			name:    "IP test created",
			request: newIPTest(),
			expectedRequest: &syntheticspb.CreateTestRequest{
				Test: testWithoutReadOnlyFields(newIPTestPayload()),
			},
			response: createTestResponse{
				data: &syntheticspb.CreateTestResponse{
					Test: newIPTestPayload(),
				},
			},
			expectedResult: newIPTest(),
		}, {
			name:    "network grid test created",
			request: newNetworkGridTest(),
			expectedRequest: &syntheticspb.CreateTestRequest{
				Test: testWithoutReadOnlyFields(newNetworkGridTestPayload()),
			},
			response: createTestResponse{
				data: &syntheticspb.CreateTestResponse{
					Test: newNetworkGridTestPayload(),
				},
			},
			expectedResult: newNetworkGridTest(),
		}, {
			name:    "hostname test created",
			request: newHostnameTest(),
			expectedRequest: &syntheticspb.CreateTestRequest{
				Test: testWithoutReadOnlyFields(newHostnameTestPayload()),
			},
			response: createTestResponse{
				data: &syntheticspb.CreateTestResponse{
					Test: newHostnameTestPayload(),
				},
			},
			expectedResult: newHostnameTest(),
		}, {
			name:    "agent test created",
			request: newAgentTest(),
			expectedRequest: &syntheticspb.CreateTestRequest{
				Test: testWithoutReadOnlyFields(newAgentTestPayload()),
			},
			response: createTestResponse{
				data: &syntheticspb.CreateTestResponse{
					Test: newAgentTestPayload(),
				},
			},
			expectedResult: newAgentTest(),
		}, {
			name:    "network mesh test created",
			request: newNetworkMeshTest(),
			expectedRequest: &syntheticspb.CreateTestRequest{
				Test: testWithoutReadOnlyFields(newNetworkMeshTestPayload()),
			},
			response: createTestResponse{
				data: &syntheticspb.CreateTestResponse{
					Test: newNetworkMeshTestPayload(),
				},
			},
			expectedResult: newNetworkMeshTest(),
		}, {
			name:    "flow test created",
			request: newFlowTest(),
			expectedRequest: &syntheticspb.CreateTestRequest{
				Test: testWithoutReadOnlyFields(newFlowTestPayload()),
			},
			response: createTestResponse{
				data: &syntheticspb.CreateTestResponse{
					Test: newFlowTestPayload(),
				},
			},
			expectedResult: newFlowTest(),
		}, {
			name:    "URL test created",
			request: newURLTest(),
			expectedRequest: &syntheticspb.CreateTestRequest{
				Test: testWithoutReadOnlyFields(newURLTestPayload()),
			},
			response: createTestResponse{
				data: &syntheticspb.CreateTestResponse{
					Test: newURLTestPayload(),
				},
			},
			expectedResult: newURLTest(),
		}, {
			name:    "page load test created",
			request: newPageLoadTest(),
			expectedRequest: &syntheticspb.CreateTestRequest{
				Test: testWithoutReadOnlyFields(newPageLoadTestPayload()),
			},
			response: createTestResponse{
				data: &syntheticspb.CreateTestResponse{
					Test: newPageLoadTestPayload(),
				},
			},
			expectedResult: newPageLoadTest(),
		}, {
			name:    "DNS test created",
			request: newDNSTest(),
			expectedRequest: &syntheticspb.CreateTestRequest{
				Test: testWithoutReadOnlyFields(newDNSTestPayload()),
			},
			response: createTestResponse{
				data: &syntheticspb.CreateTestResponse{
					Test: newDNSTestPayload(),
				},
			},
			expectedResult: newDNSTest(),
		}, {
			name:    "DNS grid test created",
			request: newDNSGridTest(),
			expectedRequest: &syntheticspb.CreateTestRequest{
				Test: testWithoutReadOnlyFields(newDNSGridTestPayload()),
			},
			response: createTestResponse{
				data: &syntheticspb.CreateTestResponse{
					Test: newDNSGridTestPayload(),
				},
			},
			expectedResult: newDNSGridTest(),
		}, {
			name:            "BGP monitor test passed",
			request:         newBGPMonitorTest(),
			expectedRequest: nil,
			expectedResult:  nil,
			expectedError:   true,
			errorPredicates: []func(error) bool{kentikapi.IsInvalidRequestError},
		}, {
			name:            "transaction test passed",
			request:         newTransactionTest(),
			expectedRequest: nil,
			expectedResult:  nil,
			expectedError:   true,
			errorPredicates: []func(error) bool{kentikapi.IsInvalidRequestError},
		},
	}
	//nolint:dupl
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			server := newSpySyntheticsServer(t, syntheticsResponses{
				createTestResponse: tt.response,
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
			result, err := client.Synthetics.CreateTest(context.Background(), tt.request)

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

			if tt.expectedRequest != nil && assert.Equal(
				t, 1, len(server.requests.createTestRequests), "invalid number of requests",
			) {
				r := server.requests.createTestRequests[0]
				assert.Equal(t, dummyAuthEmail, r.metadata.Get(authEmailKey)[0])
				assert.Equal(t, dummyAuthToken, r.metadata.Get(authAPITokenKey)[0])
				testutil.AssertProtoEqual(t, tt.expectedRequest, r.data)
			}
			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestClient_Synthetics_UpdateTest(t *testing.T) {
	tests := []struct {
		name            string
		request         *synthetics.Test
		expectedRequest *syntheticspb.UpdateTestRequest
		response        updateTestResponse
		expectedResult  *synthetics.Test
		expectedError   bool
		errorPredicates []func(error) bool
	}{
		{
			name:            "nil request",
			request:         nil,
			expectedRequest: nil,
			expectedResult:  nil,
			expectedError:   true,
			errorPredicates: []func(error) bool{kentikapi.IsInvalidRequestError},
		}, {
			name:    "empty response received",
			request: newHostnameTest(),
			expectedRequest: &syntheticspb.UpdateTestRequest{
				Test: testWithoutReadOnlyFields(newHostnameTestPayload()),
			},
			response: updateTestResponse{
				data: &syntheticspb.UpdateTestResponse{Test: nil},
			},
			expectedResult: nil,
			expectedError:  true, // InvalidResponse
		}, {
			name:    "test updated",
			request: newHostnameTest(),
			expectedRequest: &syntheticspb.UpdateTestRequest{
				Test: testWithoutReadOnlyFields(newHostnameTestPayload()),
			},
			response: updateTestResponse{
				data: &syntheticspb.UpdateTestResponse{
					Test: newHostnameTestPayload(),
				},
			},
			expectedResult: newHostnameTest(),
		}, {
			name:            "BGP monitor test passed",
			request:         newBGPMonitorTest(),
			expectedRequest: nil,
			expectedResult:  nil,
			expectedError:   true,
			errorPredicates: []func(error) bool{kentikapi.IsInvalidRequestError},
		}, {
			name:            "transaction test passed",
			request:         newTransactionTest(),
			expectedRequest: nil,
			expectedResult:  nil,
			expectedError:   true,
			errorPredicates: []func(error) bool{kentikapi.IsInvalidRequestError},
		},
	}
	// nolint: dupl
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			server := newSpySyntheticsServer(t, syntheticsResponses{
				updateTestResponse: tt.response,
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
			result, err := client.Synthetics.UpdateTest(context.Background(), tt.request)

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

			if tt.expectedRequest != nil && assert.Equal(
				t, 1, len(server.requests.updateTestRequests), "invalid number of requests",
			) {
				r := server.requests.updateTestRequests[0]
				assert.Equal(t, dummyAuthEmail, r.metadata.Get(authEmailKey)[0])
				assert.Equal(t, dummyAuthToken, r.metadata.Get(authAPITokenKey)[0])
				testutil.AssertProtoEqual(t, tt.expectedRequest, r.data)
			}

			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestClient_Synthetics_DeleteTest(t *testing.T) {
	tests := []struct {
		name            string
		requestID       string
		expectedRequest *syntheticspb.DeleteTestRequest
		response        deleteTestResponse
		expectedError   bool
		errorPredicates []func(error) bool
	}{
		{
			name:            "status InvalidArgument received",
			requestID:       "13",
			expectedRequest: &syntheticspb.DeleteTestRequest{Id: "13"},
			response: deleteTestResponse{
				data: &syntheticspb.DeleteTestResponse{},
				err:  status.Errorf(codes.InvalidArgument, codes.InvalidArgument.String()),
			},
			expectedError:   true,
			errorPredicates: []func(error) bool{kentikapi.IsInvalidRequestError},
		}, {
			name:            "resource deleted",
			requestID:       "13",
			expectedRequest: &syntheticspb.DeleteTestRequest{Id: "13"},
			response: deleteTestResponse{
				data: &syntheticspb.DeleteTestResponse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			server := newSpySyntheticsServer(t, syntheticsResponses{
				deleteTestResponse: tt.response,
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
			err = client.Synthetics.DeleteTest(context.Background(), tt.requestID)

			// assert
			t.Logf("Got err: %v", err)
			if tt.expectedError {
				assert.Error(t, err)
				for _, isErr := range tt.errorPredicates {
					assert.True(t, isErr(err))
				}
			} else {
				assert.NoError(t, err)
			}

			if assert.Equal(t, 1, len(server.requests.deleteTestRequests), "invalid number of requests") {
				r := server.requests.deleteTestRequests[0]
				assert.Equal(t, dummyAuthEmail, r.metadata.Get(authEmailKey)[0])
				assert.Equal(t, dummyAuthToken, r.metadata.Get(authAPITokenKey)[0])
				testutil.AssertProtoEqual(t, tt.expectedRequest, r.data)
			}
		})
	}
}

func TestClient_Synthetics_SetTestStatus(t *testing.T) {
	tests := []struct {
		name            string
		requestID       string
		requestStatus   synthetics.TestStatus
		expectedRequest *syntheticspb.SetTestStatusRequest
		response        setTestStatusResponse
		expectedError   bool
		errorPredicates []func(error) bool
	}{
		{
			name:          "status InvalidArgument received",
			requestID:     "13",
			requestStatus: synthetics.TestStatusDeleted,
			expectedRequest: &syntheticspb.SetTestStatusRequest{
				Id:     "13",
				Status: syntheticspb.TestStatus_TEST_STATUS_DELETED,
			},
			response: setTestStatusResponse{
				data: &syntheticspb.SetTestStatusResponse{},
				err:  status.Errorf(codes.InvalidArgument, codes.InvalidArgument.String()),
			},
			expectedError:   true,
			errorPredicates: []func(error) bool{kentikapi.IsInvalidRequestError},
		}, {
			name:          "status set",
			requestID:     "13",
			requestStatus: synthetics.TestStatusPaused,
			expectedRequest: &syntheticspb.SetTestStatusRequest{
				Id:     "13",
				Status: syntheticspb.TestStatus_TEST_STATUS_PAUSED,
			},
			response: setTestStatusResponse{
				data: &syntheticspb.SetTestStatusResponse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			server := newSpySyntheticsServer(t, syntheticsResponses{
				setTestStatusResponse: tt.response,
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
			err = client.Synthetics.SetTestStatus(context.Background(), tt.requestID, tt.requestStatus)

			// assert
			t.Logf("Got err: %v", err)
			if tt.expectedError {
				assert.Error(t, err)
				for _, isErr := range tt.errorPredicates {
					assert.True(t, isErr(err))
				}
			} else {
				assert.NoError(t, err)
			}

			if assert.Equal(t, 1, len(server.requests.setTestStatusRequests), "invalid number of requests") {
				r := server.requests.setTestStatusRequests[0]
				assert.Equal(t, dummyAuthEmail, r.metadata.Get(authEmailKey)[0])
				assert.Equal(t, dummyAuthToken, r.metadata.Get(authAPITokenKey)[0])
				testutil.AssertProtoEqual(t, tt.expectedRequest, r.data)
			}
		})
	}
}

func testWithoutReadOnlyFields(test *syntheticspb.Test) *syntheticspb.Test {
	test.Cdate = nil
	test.CreatedBy = nil
	test.LastUpdatedBy = nil
	return test
}

func newIPTest() *synthetics.Test {
	t := newTest()
	t.Name = "ip-test"
	t.Type = synthetics.TestTypeIP
	t.ID = ipTestID
	t.Settings.Definition = &synthetics.TestDefinitionIP{
		Targets: []net.IP{net.ParseIP("192.0.2.213"), net.ParseIP("2001:db8:dead:beef:dead:beef:dead:beef")},
	}
	return t
}

func newIPTestPayload() *syntheticspb.Test {
	t := newTestPayload()
	t.Name = "ip-test"
	t.Type = "ip"
	t.Id = ipTestID
	t.Settings.Definition = &syntheticspb.TestSettings_Ip{
		Ip: &syntheticspb.IpTest{
			Targets: []string{"192.0.2.213", "2001:db8:dead:beef:dead:beef:dead:beef"},
		},
	}
	return t
}

func newNetworkGridTest() *synthetics.Test {
	t := newTest()
	t.Name = "network-grid-test"
	t.Type = synthetics.TestTypeNetworkGrid
	t.ID = networkGridTestID
	t.Settings.Definition = &synthetics.TestDefinitionNetworkGrid{
		Targets: []net.IP{net.ParseIP("192.0.2.213"), net.ParseIP("2001:db8:dead:beef:dead:beef:dead:beef")},
	}
	return t
}

func newNetworkGridTestPayload() *syntheticspb.Test {
	t := newTestPayload()
	t.Name = "network-grid-test"
	t.Type = "network_grid"
	t.Id = networkGridTestID
	t.Settings.Definition = &syntheticspb.TestSettings_NetworkGrid{
		NetworkGrid: &syntheticspb.IpTest{
			Targets: []string{"192.0.2.213", "2001:db8:dead:beef:dead:beef:dead:beef"},
		},
	}
	return t
}

func newHostnameTest() *synthetics.Test {
	t := newTest()
	t.Name = "hostname-test"
	t.Type = synthetics.TestTypeHostname
	t.ID = hostnameTestID
	t.Settings.Definition = &synthetics.TestDefinitionHostname{
		Target: "www.example.com",
	}
	return t
}

func newHostnameTestPayload() *syntheticspb.Test {
	t := newTestPayload()
	t.Name = "hostname-test"
	t.Type = "hostname"
	t.Id = hostnameTestID
	t.Settings.Definition = &syntheticspb.TestSettings_Hostname{
		Hostname: &syntheticspb.HostnameTest{
			Target: "www.example.com",
		},
	}
	return t
}

func newAgentTest() *synthetics.Test {
	t := newTest()
	t.Name = "agent-test"
	t.Type = synthetics.TestTypeAgent
	t.ID = agentTestID
	t.Settings.Definition = &synthetics.TestDefinitionAgent{
		Target:     "dummy-agent-id",
		UseLocalIP: true,
	}
	return t
}

func newAgentTestPayload() *syntheticspb.Test {
	t := newTestPayload()
	t.Name = "agent-test"
	t.Type = "agent"
	t.Id = agentTestID
	t.Settings.Definition = &syntheticspb.TestSettings_Agent{
		Agent: &syntheticspb.AgentTest{
			Target:     "dummy-agent-id",
			UseLocalIp: true,
		},
	}
	return t
}

func newNetworkMeshTest() *synthetics.Test {
	t := newTest()
	t.Name = "network-mesh-test"
	t.Type = synthetics.TestTypeNetworkMesh
	t.ID = networkMeshTestID
	t.Settings.Definition = &synthetics.TestDefinitionNetworkMesh{
		UseLocalIP: true,
	}
	return t
}

func newNetworkMeshTestPayload() *syntheticspb.Test {
	t := newTestPayload()
	t.Name = "network-mesh-test"
	t.Type = "network_mesh"
	t.Id = networkMeshTestID
	t.Settings.Definition = &syntheticspb.TestSettings_NetworkMesh{
		NetworkMesh: &syntheticspb.NetworkMeshTest{
			UseLocalIp: true,
		},
	}
	return t
}

func newFlowTest() *synthetics.Test {
	t := newTest()
	t.Name = "flow-test"
	t.Type = synthetics.TestTypeFlow
	t.ID = flowTestID
	t.Settings.Definition = &synthetics.TestDefinitionFlow{
		Type:                  synthetics.FlowTestTypeCity,
		Target:                "Warsaw",
		TargetRefreshInterval: 168 * time.Hour,
		MaxIPTargets:          10,
		MaxProviders:          3,
		Direction:             synthetics.DirectionSrc,
		InetDirection:         synthetics.DirectionDst,
	}
	return t
}

func newFlowTestPayload() *syntheticspb.Test {
	t := newTestPayload()
	t.Name = "flow-test"
	t.Type = "flow"
	t.Id = flowTestID
	t.Settings.Definition = &syntheticspb.TestSettings_Flow{
		Flow: &syntheticspb.FlowTest{
			Target:                      "Warsaw",
			TargetRefreshIntervalMillis: 604800000,
			MaxProviders:                3,
			MaxIpTargets:                10,
			Type:                        "city",
			InetDirection:               "dst",
			Direction:                   "src",
		},
	}
	return t
}

func newURLTest() *synthetics.Test {
	t := newTest()
	t.Name = "url-test"
	t.Type = synthetics.TestTypeURL
	t.ID = urlTestID
	t.Settings.Definition = &synthetics.TestDefinitionURL{
		Target: url.URL{
			Scheme:   "https",
			Host:     "www.example.com:443",
			RawQuery: "dummy=query",
		},
		Timeout: time.Minute,
		Method:  http.MethodGet,
		Headers: map[string]string{
			"dummy-key-1": "dummy-value-1",
			"dummy-key-2": "dummy-value-2",
		},
		Body:            "dummy-body",
		IgnoreTLSErrors: true,
	}
	return t
}

func newURLTestPayload() *syntheticspb.Test {
	t := newTestPayload()
	t.Name = "url-test"
	t.Type = "url"
	t.Id = urlTestID
	t.Settings.Definition = &syntheticspb.TestSettings_Url{
		Url: &syntheticspb.UrlTest{
			Target:  "https://www.example.com:443?dummy=query",
			Timeout: 60000,
			Method:  "GET",
			Headers: map[string]string{
				"dummy-key-1": "dummy-value-1",
				"dummy-key-2": "dummy-value-2",
			},
			Body:            "dummy-body",
			IgnoreTlsErrors: true,
		},
	}
	return t
}

func newPageLoadTest() *synthetics.Test {
	t := newTest()
	t.Name = "page-load-test"
	t.Type = synthetics.TestTypePageLoad
	t.ID = pageLoadTestID
	t.Settings.Definition = &synthetics.TestDefinitionPageLoad{
		Target: url.URL{
			Scheme:   "https",
			Host:     "www.example.com:443",
			RawQuery: "dummy=query",
		},
		Timeout: time.Minute,
		Headers: map[string]string{
			"dummy-key-1": "dummy-value-1",
			"dummy-key-2": "dummy-value-2",
		},
		CSSSelectors: map[string]string{
			"dummy-key-1": "dummy-selector-1",
			"dummy-key-2": "dummy-selector-2",
		},
		IgnoreTLSErrors: true,
	}
	return t
}

func newPageLoadTestPayload() *syntheticspb.Test {
	t := newTestPayload()
	t.Name = "page-load-test"
	t.Type = "page_load"
	t.Id = pageLoadTestID
	t.Settings.Definition = &syntheticspb.TestSettings_PageLoad{
		PageLoad: &syntheticspb.PageLoadTest{
			Target:  "https://www.example.com:443?dummy=query",
			Timeout: 60000,
			Headers: map[string]string{
				"dummy-key-1": "dummy-value-1",
				"dummy-key-2": "dummy-value-2",
			},
			IgnoreTlsErrors: true,
			CssSelectors: map[string]string{
				"dummy-key-1": "dummy-selector-1",
				"dummy-key-2": "dummy-selector-2",
			},
		},
	}
	return t
}

func newDNSTest() *synthetics.Test {
	t := newTest()
	t.Name = "dns-test"
	t.Type = synthetics.TestTypeDNS
	t.ID = dnsTestID
	t.Settings.Definition = &synthetics.TestDefinitionDNS{
		Target:     "www.example.com",
		Timeout:    time.Minute,
		RecordType: synthetics.DNSRecordAAAA,
		Servers:    []net.IP{net.ParseIP("192.0.2.213"), net.ParseIP("2001:db8:dead:beef:dead:beef:dead:beef")},
		Port:       53,
	}
	return t
}

func newDNSTestPayload() *syntheticspb.Test {
	t := newTestPayload()
	t.Name = "dns-test"
	t.Type = "dns"
	t.Id = dnsTestID
	t.Settings.Definition = &syntheticspb.TestSettings_Dns{
		Dns: &syntheticspb.DnsTest{
			Target:     "www.example.com",
			Timeout:    60000,
			RecordType: syntheticspb.DNSRecord_DNS_RECORD_AAAA,
			Servers:    []string{"192.0.2.213", "2001:db8:dead:beef:dead:beef:dead:beef"},
			Port:       53,
		},
	}
	return t
}

func newDNSGridTest() *synthetics.Test {
	t := newTest()
	t.Name = "dns-grid-test"
	t.Type = synthetics.TestTypeDNSGrid
	t.ID = dnsGridTestID
	t.Settings.Definition = &synthetics.TestDefinitionDNSGrid{
		Target:     "www.example.com",
		Timeout:    time.Minute,
		RecordType: synthetics.DNSRecordAAAA,
		Servers:    []net.IP{net.ParseIP("192.0.2.213"), net.ParseIP("2001:db8:dead:beef:dead:beef:dead:beef")},
		Port:       53,
	}
	return t
}

func newDNSGridTestPayload() *syntheticspb.Test {
	t := newTestPayload()
	t.Name = "dns-grid-test"
	t.Type = "dns_grid"
	t.Id = dnsGridTestID
	t.Settings.Definition = &syntheticspb.TestSettings_DnsGrid{
		DnsGrid: &syntheticspb.DnsTest{
			Target:     "www.example.com",
			Timeout:    60000,
			RecordType: syntheticspb.DNSRecord_DNS_RECORD_AAAA,
			Servers:    []string{"192.0.2.213", "2001:db8:dead:beef:dead:beef:dead:beef"},
			Port:       53,
		},
	}
	return t
}

func newBGPMonitorTest() *synthetics.Test {
	t := newTest()
	t.Name = "bgp-monitor-test"
	t.Type = "bgp_monitor"
	t.ID = bgpMonitorTestID
	return t
}

func newBGPMonitorTestPayload() *syntheticspb.Test {
	t := newTestPayload()
	t.Name = "bgp-monitor-test"
	t.Type = "bgp_monitor"
	t.Id = bgpMonitorTestID
	return t
}

func newTransactionTest() *synthetics.Test {
	t := newTest()
	t.Name = "transaction-test"
	t.Type = "transaction"
	t.ID = transactionTestID
	return t
}

func newTransactionTestPayload() *syntheticspb.Test {
	t := newTestPayload()
	t.Name = "transaction-test"
	t.Type = "transaction"
	t.Id = transactionTestID
	return t
}

func newTest() *synthetics.Test {
	return &synthetics.Test{
		Name:       "dummy-test",
		Type:       "unknown-type",
		Status:     synthetics.TestStatusActive,
		UpdateDate: pointer.ToTime(time.Date(2022, time.April, 8, 7, 26, 51, 505*1000000, time.UTC)),
		Settings: synthetics.TestSettings{
			Definition:           nil,
			AgentIDs:             []string{"817", "818", "819"},
			Period:               60 * time.Second,
			Family:               synthetics.IPFamilyDual,
			NotificationChannels: []string{"7143fa58", "14f23014"},
			Health: synthetics.HealthSettings{
				LatencyCritical:           1 * time.Millisecond,
				LatencyWarning:            2 * time.Millisecond,
				LatencyCriticalStdDev:     3 * time.Millisecond,
				LatencyWarningStdDev:      4 * time.Millisecond,
				JitterCritical:            5 * time.Millisecond,
				JitterWarning:             6 * time.Millisecond,
				JitterCriticalStdDev:      7 * time.Millisecond,
				JitterWarningStdDev:       8 * time.Millisecond,
				PacketLossCritical:        9,
				PacketLossWarning:         10,
				HTTPLatencyCritical:       11 * time.Millisecond,
				HTTPLatencyWarning:        12 * time.Millisecond,
				HTTPLatencyCriticalStdDev: 13 * time.Millisecond,
				HTTPLatencyWarningStdDev:  14 * time.Millisecond,
				HTTPValidCodes:            []uint32{http.StatusOK, http.StatusCreated},
				DNSValidCodes:             []uint32{1, 2, 3},
				UnhealthySubtestThreshold: 42,
				AlarmActivation: &synthetics.AlarmActivationSettings{
					TimeWindow:  5 * time.Minute,
					Times:       3,
					GracePeriod: 2,
				},
			},
			Ping: &synthetics.PingSettings{
				Timeout:  3 * time.Second,
				Count:    5,
				Delay:    1 * time.Millisecond,
				Protocol: synthetics.PingProtocolTCP,
				Port:     443,
			},
			Traceroute: &synthetics.TracerouteSettings{
				Timeout:  22500 * time.Millisecond,
				Count:    3,
				Delay:    1 * time.Millisecond,
				Protocol: synthetics.TracerouteProtocolUDP,
				Port:     33434,
				Limit:    30,
			},
			Tasks: []synthetics.TaskType{synthetics.TaskTypePing, synthetics.TaskTypeTraceroute},
		},
		ID:         "dummy-id",
		CreateDate: time.Date(2022, time.April, 6, 9, 43, 39, 324*1000000, time.UTC),
		CreatedBy: synthetics.UserInfo{
			ID:       "4321",
			Email:    "joe.doe@example.com",
			FullName: "Joe Doe",
		},
		LastUpdatedBy: &synthetics.UserInfo{
			ID:       "4321",
			Email:    "joe.doe@example.com",
			FullName: "Joe Doe",
		},
	}
}

func newTestPayload() *syntheticspb.Test {
	return &syntheticspb.Test{
		Id:     "dummy-id",
		Name:   "dummy-test",
		Type:   "dummy-type",
		Status: syntheticspb.TestStatus_TEST_STATUS_ACTIVE,
		Settings: &syntheticspb.TestSettings{
			Definition: nil,
			AgentIds:   []string{"817", "818", "819"},
			Tasks:      []string{"ping", "traceroute"},
			HealthSettings: &syntheticspb.HealthSettings{
				LatencyCritical:           1,
				LatencyWarning:            2,
				PacketLossCritical:        9,
				PacketLossWarning:         10,
				JitterCritical:            5,
				JitterWarning:             6,
				HttpLatencyCritical:       11,
				HttpLatencyWarning:        12,
				HttpValidCodes:            []uint32{200, 201},
				DnsValidCodes:             []uint32{1, 2, 3},
				LatencyCriticalStddev:     3,
				LatencyWarningStddev:      4,
				JitterCriticalStddev:      7,
				JitterWarningStddev:       8,
				HttpLatencyCriticalStddev: 13,
				HttpLatencyWarningStddev:  14,
				UnhealthySubtestThreshold: 42,
				Activation: &syntheticspb.ActivationSettings{
					GracePeriod: "2",
					TimeUnit:    "m",
					TimeWindow:  "5",
					Times:       "3",
				},
			},
			Ping: &syntheticspb.TestPingSettings{
				Count:    5,
				Protocol: "tcp",
				Port:     443,
				Timeout:  3000,
				Delay:    1,
			},
			Trace: &syntheticspb.TestTraceSettings{
				Count:    3,
				Protocol: "udp",
				Port:     33434,
				Timeout:  22500,
				Limit:    30,
				Delay:    1,
			},
			Period:               60,
			Family:               syntheticspb.IPFamily_IP_FAMILY_DUAL,
			NotificationChannels: []string{"7143fa58", "14f23014"},
		},
		Cdate: timestamppb.New(time.Date(2022, time.April, 6, 9, 43, 39, 324*1000000, time.UTC)),
		Edate: timestamppb.New(time.Date(2022, time.April, 8, 7, 26, 51, 505*1000000, time.UTC)),
		CreatedBy: &syntheticspb.UserInfo{
			Id:       "4321",
			Email:    "joe.doe@example.com",
			FullName: "Joe Doe",
		},
		LastUpdatedBy: &syntheticspb.UserInfo{
			Id:       "4321",
			Email:    "joe.doe@example.com",
			FullName: "Joe Doe",
		},
	}
}
