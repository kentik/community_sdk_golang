package kentikapi_test

import (
	"context"
	"io/ioutil"
	"net"
	"path/filepath"
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

func TestClient_Synthetics_GetTraceForTest(t *testing.T) {
	tests := []struct {
		name            string
		response        getTraceForTestResponse
		checkResult     func(*testing.T, synthetics.GetTraceForTestResponse)
		expectedError   bool
		errorPredicates []func(error) bool
	}{
		{
			name: "status InvalidArgument received",
			response: getTraceForTestResponse{
				err: status.Errorf(codes.InvalidArgument, codes.InvalidArgument.String()),
			},
			checkResult: func(t *testing.T, result synthetics.GetTraceForTestResponse) {
				assert.Equal(t, synthetics.GetTraceForTestResponse{}, result)
			},
			expectedError:   true,
			errorPredicates: []func(error) bool{kentikapi.IsInvalidRequestError},
		}, {
			name: "empty response received",
			response: getTraceForTestResponse{
				data: &syntheticspb.GetTraceForTestResponse{},
			},
			checkResult: func(t *testing.T, result synthetics.GetTraceForTestResponse) {
				assert.Equal(t,
					synthetics.GetTraceForTestResponse{
						Nodes: map[models.ID]synthetics.NetworkNode{},
						Paths: nil,
					},
					result,
				)
			},
			expectedError: false,
		}, {
			name: "trace results received",
			response: getTraceForTestResponse{
				data: getTraceForTestResponseFromJSONFile(
					t,
					filepath.Join("internal", "synthetics", "testdata", "get_test_for_trace_response.json"),
				),
			},
			checkResult: checkGetTraceForTestResult,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			server := newSpySyntheticsServer(t, syntheticsResponses{
				getTraceForTestResponse: tt.response,
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
			result, err := client.Synthetics.GetTraceForTest(context.Background(), newGetTraceForTestsRequest())

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

			if assert.Equal(t, 1, len(server.requests.getTraceForTestRequests), "invalid number of requests") {
				r := server.requests.getTraceForTestRequests[0]
				assert.Equal(t, dummyAuthEmail, r.metadata.Get(authEmailKey)[0])
				assert.Equal(t, dummyAuthToken, r.metadata.Get(authAPITokenKey)[0])
				testutil.AssertProtoEqual(t,
					newGetTraceForTestRequestPayload(),
					r.data,
				)
			}

			tt.checkResult(t, result)
		})
	}
}

func getTraceForTestResponseFromJSONFile(t *testing.T, filePath string) *syntheticspb.GetTraceForTestResponse {
	payloadJSON, err := ioutil.ReadFile(filePath) // nolint: gosec
	require.NoError(t, err)

	var r syntheticspb.GetTraceForTestResponse
	err = protojson.Unmarshal(payloadJSON, &r)
	require.NoError(t, err)
	return &r
}

func newGetTraceForTestsRequest() synthetics.GetTraceForTestRequest {
	return synthetics.GetTraceForTestRequest{
		TestID:    "40785",
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

func newGetTraceForTestRequestPayload() *syntheticspb.GetTraceForTestRequest {
	return &syntheticspb.GetTraceForTestRequest{
		Id:        "40785",
		StartTime: timestamppb.New(time.Date(2022, time.July, 19, 14, 11, 22, 123456789, time.UTC)),
		EndTime:   timestamppb.New(time.Date(2022, time.July, 19, 20, 11, 22, 123456789, time.UTC)),
		AgentIds:  []string{"598", "608", "631"},
		TargetIps: []string{"23.92.28.124", "54.152.163.64", "95.179.136.58"},
	}
}

func checkGetTraceForTestResult(t *testing.T, result synthetics.GetTraceForTestResponse) {
	if assert.Len(t, result.Nodes, 76) {
		assert.Equal(t,
			synthetics.NetworkNode{
				IP:     net.ParseIP("23.92.28.124"),
				ASN:    63949,
				AsName: "Akamai (Linode),US",
				Location: &synthetics.Location{
					Latitude:  33.749001,
					Longitude: -84.387978,
					Country:   "United States",
					Region:    "Georgia",
					City:      "Atlanta",
				},
				DNSName:  "li661-124.members.linode.com",
				DeviceID: "",
				SiteID:   "",
			},
			result.Nodes["23.92.28.124"],
		)
	}

	if assert.Len(t, result.Paths, 11) {
		checkTraceResultFirstPath(t, result.Paths[0])
	}
}

func checkTraceResultFirstPath(t *testing.T, path synthetics.Path) {
	traces := path.Traces
	path.Traces = nil

	assert.Equal(t,
		synthetics.Path{
			Time:     time.Date(2022, time.July, 19, 14, 0, 0, 0, time.UTC),
			AgentID:  "631",
			TargetIP: net.ParseIP("23.92.28.124"),
			HopCount: synthetics.Stats{
				Average: 21,
				Min:     0,
				Max:     29,
			},
			MaxASPathLength: 10,
			Traces:          nil, // set to nil above
		},
		path,
	)

	if !assert.Len(t, traces, 3) {
		return
	}

	assert.Equal(t,
		synthetics.PathTrace{
			ASPath:     []int32{20473, 1299, 1299, 1299, 1299, 1299, 1299, 1299, 1299, 63949},
			IsComplete: true,
			Hops: []synthetics.TraceHop{
				{Latency: 0, NodeID: ""},
				{Latency: 101423, NodeID: "45.76.40.33"},
				{Latency: 0, NodeID: ""},
				{Latency: 0, NodeID: ""},
				{Latency: 673, NodeID: "62.115.58.193"},
				{Latency: 1506, NodeID: "62.115.116.124"},
				{Latency: 8137, NodeID: "62.115.134.97"},
				{Latency: 93897, NodeID: "62.115.112.242"},
				{Latency: 90701, NodeID: "62.115.136.200"},
				{Latency: 107125, NodeID: "62.115.137.59"},
				{Latency: 116351, NodeID: "62.115.134.10"},
				{Latency: 108991, NodeID: "62.115.190.69"},
				{Latency: 0, NodeID: ""},
				{Latency: 0, NodeID: ""},
				{Latency: 0, NodeID: ""},
				{Latency: 109005, NodeID: "23.92.28.124"},
			},
		},
		traces[0],
	)
}
