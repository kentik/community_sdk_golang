package kentikapi_test

import (
	"context"
	"errors"
	"net"
	"testing"

	"github.com/AlekSi/pointer"
	cloudexportpb "github.com/kentik/api-schema-public/gen/go/kentik/cloud_export/v202101beta1"
	"github.com/kentik/community_sdk_golang/kentikapi"
	"github.com/kentik/community_sdk_golang/kentikapi/cloud"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/testutil"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	awsExportID   = "1001"
	azureExportID = "1002"
	gceExportID   = "1003"
	ibmExportID   = "1004"
)

func TestClient_Cloud_GetAllExports(t *testing.T) {
	tests := []struct {
		name            string
		response        listCEResponse
		expectedResult  *cloud.GetAllExportsResponse
		expectedError   bool
		errorPredicates []func(error) bool
	}{
		{
			name: "status InvalidArgument received",
			response: listCEResponse{
				err: status.Errorf(codes.InvalidArgument, codes.InvalidArgument.String()),
			},
			expectedError:   true,
			errorPredicates: []func(error) bool{kentikapi.IsInvalidRequestError},
		}, {
			name: "empty response received",
			response: listCEResponse{
				data: &cloudexportpb.ListCloudExportResponse{},
			},
			expectedResult: &cloud.GetAllExportsResponse{
				Exports:             nil,
				InvalidExportsCount: 0,
			},
		}, {
			name: "no exports received",
			response: listCEResponse{
				data: &cloudexportpb.ListCloudExportResponse{
					Exports:             []*cloudexportpb.CloudExport{},
					InvalidExportsCount: 0,
				},
			},
			expectedResult: &cloud.GetAllExportsResponse{
				Exports:             nil,
				InvalidExportsCount: 0,
			},
		}, {
			name: "4 exports received",
			response: listCEResponse{
				data: &cloudexportpb.ListCloudExportResponse{
					Exports: []*cloudexportpb.CloudExport{
						newAWSExportPayload(),
						newAzureExportPayload(),
						newGCEExportPayload(),
						newIBMExportPayload(),
					},
					InvalidExportsCount: 1,
				},
			},
			expectedResult: &cloud.GetAllExportsResponse{
				Exports: []cloud.Export{
					*newAWSExport(),
					*newAzureExport(),
					*newGCEExport(),
					*newIBMExport(),
				},
				InvalidExportsCount: 1,
			},
		}, {
			name: "2 exports received - one nil",
			response: listCEResponse{
				data: &cloudexportpb.ListCloudExportResponse{
					Exports: []*cloudexportpb.CloudExport{
						newAWSExportPayload(),
						nil,
					},
					InvalidExportsCount: 0,
				},
			},
			expectedError: true, // InvalidResponse
		},
	}
	//nolint:dupl
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			server := newSpyCloudExportServer(t, cloudExportResponses{
				listCEResponse: tt.response,
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
			result, err := client.Cloud.GetAllExports(context.Background())

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

			if assert.Equal(t, 1, len(server.requests.listCERequests), "invalid number of requests") {
				r := server.requests.listCERequests[0]
				assert.Equal(t, dummyAuthEmail, r.metadata.Get(authEmailKey)[0])
				assert.Equal(t, dummyAuthToken, r.metadata.Get(authAPITokenKey)[0])
				testutil.AssertProtoEqual(t, &cloudexportpb.ListCloudExportRequest{}, r.data)
			}

			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestClient_Cloud_GetExport(t *testing.T) {
	tests := []struct {
		name            string
		requestID       models.ID
		expectedRequest *cloudexportpb.GetCloudExportRequest
		response        getCEResponse
		expectedResult  *cloud.Export
		expectedError   bool
		errorPredicates []func(error) bool
	}{
		{
			name:            "status InvalidArgument received",
			requestID:       "13",
			expectedRequest: &cloudexportpb.GetCloudExportRequest{Id: "13"},
			response: getCEResponse{
				err: status.Errorf(codes.InvalidArgument, codes.InvalidArgument.String()),
			},
			expectedError:   true,
			errorPredicates: []func(error) bool{kentikapi.IsInvalidRequestError},
		}, {
			name:            "status NotFound received",
			requestID:       "13",
			expectedRequest: &cloudexportpb.GetCloudExportRequest{Id: "13"},
			response: getCEResponse{
				err: status.Errorf(codes.NotFound, codes.NotFound.String()),
			},
			expectedError:   true,
			errorPredicates: []func(error) bool{kentikapi.IsNotFoundError},
		}, {
			name:            "empty response received",
			requestID:       "13",
			expectedRequest: &cloudexportpb.GetCloudExportRequest{Id: "13"},
			response: getCEResponse{
				data: &cloudexportpb.GetCloudExportResponse{},
			},
			expectedError: true, // InvalidResponse
		}, {
			name:            "minimal AWS cloud export received",
			requestID:       "58192",
			expectedRequest: &cloudexportpb.GetCloudExportRequest{Id: "58192"},
			response: getCEResponse{
				// minimal AWS cloud export data returned by Kentik API
				data: &cloudexportpb.GetCloudExportResponse{
					Export: &cloudexportpb.CloudExport{
						Id:            "58192",
						Type:          cloudexportpb.CloudExportType_CLOUD_EXPORT_TYPE_KENTIK_MANAGED,
						Enabled:       true,
						Name:          "minimal-aws-export",
						Description:   "",
						ApiRoot:       "https://api.kentik.com",
						FlowDest:      "https://flow.kentik.com",
						PlanId:        "11467",
						CloudProvider: "aws",
						Properties: &cloudexportpb.CloudExport_Aws{
							Aws: &cloudexportpb.AwsProperties{
								Bucket:          "dummy-bucket",
								IamRoleArn:      "",
								Region:          "",
								DeleteAfterRead: false,
								MultipleBuckets: false,
							},
						},
						CurrentStatus: &cloudexportpb.Status{
							Status:       "",
							ErrorMessage: "",
						},
					},
				},
			},
			expectedResult: &cloud.Export{
				ID:          "58192",
				Type:        cloud.ExportTypeKentikManaged,
				Enabled:     pointer.ToBool(true),
				Name:        "minimal-aws-export",
				Description: "",
				PlanID:      "11467",
				Provider:    cloud.ProviderAWS,
				Properties: &cloud.AWSProperties{
					Bucket:          "dummy-bucket",
					IAMRoleARN:      "",
					Region:          "",
					DeleteAfterRead: pointer.ToBool(false),
					MultipleBuckets: pointer.ToBool(false),
				},
				CurrentStatus: &cloud.ExportStatus{
					Status:       "",
					ErrorMessage: "",
				},
			},
		}, {
			name:            "AWS cloud export received",
			requestID:       awsExportID,
			expectedRequest: &cloudexportpb.GetCloudExportRequest{Id: awsExportID},
			response: getCEResponse{
				data: &cloudexportpb.GetCloudExportResponse{Export: newAWSExportPayload()},
			},
			expectedResult: newAWSExport(),
		}, {
			name:            "Azure cloud export received",
			requestID:       azureExportID,
			expectedRequest: &cloudexportpb.GetCloudExportRequest{Id: azureExportID},
			response: getCEResponse{
				data: &cloudexportpb.GetCloudExportResponse{Export: newAzureExportPayload()},
			},
			expectedResult: newAzureExport(),
		}, {
			name:            "GCE cloud export received",
			requestID:       gceExportID,
			expectedRequest: &cloudexportpb.GetCloudExportRequest{Id: gceExportID},
			response: getCEResponse{
				data: &cloudexportpb.GetCloudExportResponse{Export: newGCEExportPayload()},
			},
			expectedResult: newGCEExport(),
		}, {
			name:            "IBM cloud export received",
			requestID:       ibmExportID,
			expectedRequest: &cloudexportpb.GetCloudExportRequest{Id: ibmExportID},
			response: getCEResponse{
				data: &cloudexportpb.GetCloudExportResponse{Export: newIBMExportPayload()},
			},
			expectedResult: newIBMExport(),
		},
	}
	//nolint:dupl
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			server := newSpyCloudExportServer(t, cloudExportResponses{
				getCEResponse: tt.response,
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
			result, err := client.Cloud.GetExport(context.Background(), tt.requestID)

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

			if assert.Equal(t, 1, len(server.requests.getCERequests), "invalid number of requests") {
				r := server.requests.getCERequests[0]
				assert.Equal(t, dummyAuthEmail, r.metadata.Get(authEmailKey)[0])
				assert.Equal(t, dummyAuthToken, r.metadata.Get(authAPITokenKey)[0])
				testutil.AssertProtoEqual(t, tt.expectedRequest, r.data)
			}

			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestClient_Cloud_CreateExport(t *testing.T) {
	tests := []struct {
		name            string
		request         *cloud.Export
		expectedRequest *cloudexportpb.CreateCloudExportRequest
		response        createCEResponse
		expectedResult  *cloud.Export
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
			request: newAWSExport(),
			expectedRequest: func() *cloudexportpb.CreateCloudExportRequest {
				r := &cloudexportpb.CreateCloudExportRequest{
					Export: newAWSExportPayload(),
				}
				// read-only fields should be omitted
				r.Export.CurrentStatus = nil
				return r
			}(),
			response: createCEResponse{
				data: &cloudexportpb.CreateCloudExportResponse{Export: nil},
			},
			expectedResult: nil,
			expectedError:  true, // InvalidResponse
		}, {
			name: "minimal AWS export created",
			request: cloud.NewAWSExport(cloud.AWSExportRequiredFields{
				Name:   "minimal-aws-export",
				PlanID: "11467",
				AWSProperties: cloud.AWSPropertiesRequiredFields{
					Bucket: "dummy-bucket",
				},
			}),
			expectedRequest: &cloudexportpb.CreateCloudExportRequest{
				Export: &cloudexportpb.CloudExport{
					Name:          "minimal-aws-export",
					PlanId:        "11467",
					CloudProvider: "aws",
					Properties: &cloudexportpb.CloudExport_Aws{
						Aws: &cloudexportpb.AwsProperties{
							Bucket: "dummy-bucket",
						},
					},
				},
			},
			response: createCEResponse{
				// minimal AWS cloud export data returned by Kentik API
				data: &cloudexportpb.CreateCloudExportResponse{
					Export: &cloudexportpb.CloudExport{
						Id:            "58192",
						Type:          cloudexportpb.CloudExportType_CLOUD_EXPORT_TYPE_KENTIK_MANAGED,
						Enabled:       true,
						Name:          "minimal-aws-export",
						Description:   "",
						ApiRoot:       "https://api.kentik.com",
						FlowDest:      "https://flow.kentik.com",
						PlanId:        "11467",
						CloudProvider: "aws",
						Properties: &cloudexportpb.CloudExport_Aws{
							Aws: &cloudexportpb.AwsProperties{
								Bucket:          "dummy-bucket",
								IamRoleArn:      "",
								Region:          "",
								DeleteAfterRead: false,
								MultipleBuckets: false,
							},
						},
						CurrentStatus: &cloudexportpb.Status{
							Status:       "",
							ErrorMessage: "",
						},
					},
				},
			},
			expectedResult: &cloud.Export{
				ID:          "58192",
				Type:        cloud.ExportTypeKentikManaged,
				Enabled:     pointer.ToBool(true),
				Name:        "minimal-aws-export",
				Description: "",
				PlanID:      "11467",
				Provider:    cloud.ProviderAWS,
				Properties: &cloud.AWSProperties{
					Bucket:          "dummy-bucket",
					IAMRoleARN:      "",
					Region:          "",
					DeleteAfterRead: pointer.ToBool(false),
					MultipleBuckets: pointer.ToBool(false),
				},
				CurrentStatus: &cloud.ExportStatus{
					Status:       "",
					ErrorMessage: "",
				},
			},
		}, {
			name:    "AWS export created",
			request: newAWSExport(),
			expectedRequest: func() *cloudexportpb.CreateCloudExportRequest {
				r := &cloudexportpb.CreateCloudExportRequest{
					Export: newAWSExportPayload(),
				}
				// read-only fields should be omitted
				r.Export.CurrentStatus = nil
				return r
			}(),
			response: createCEResponse{
				data: &cloudexportpb.CreateCloudExportResponse{
					Export: newAWSExportPayload(),
				},
			},
			expectedResult: newAWSExport(),
		}, {
			name:    "Azure export created",
			request: newAzureExport(),
			expectedRequest: func() *cloudexportpb.CreateCloudExportRequest {
				r := &cloudexportpb.CreateCloudExportRequest{
					Export: newAzureExportPayload(),
				}
				// read-only fields should be omitted
				r.Export.CurrentStatus = nil
				return r
			}(),
			response: createCEResponse{
				data: &cloudexportpb.CreateCloudExportResponse{
					Export: newAzureExportPayload(),
				},
			},
			expectedResult: newAzureExport(),
		}, {
			name:    "GCE export created",
			request: newGCEExport(),
			expectedRequest: func() *cloudexportpb.CreateCloudExportRequest {
				r := &cloudexportpb.CreateCloudExportRequest{
					Export: newGCEExportPayload(),
				}
				// read-only fields should be omitted
				r.Export.CurrentStatus = nil
				return r
			}(),
			response: createCEResponse{
				data: &cloudexportpb.CreateCloudExportResponse{
					Export: newGCEExportPayload(),
				},
			},
			expectedResult: newGCEExport(),
		}, {
			name:    "IBM export created",
			request: newIBMExport(),
			expectedRequest: func() *cloudexportpb.CreateCloudExportRequest {
				r := &cloudexportpb.CreateCloudExportRequest{
					Export: newIBMExportPayload(),
				}
				// read-only fields should be omitted
				r.Export.CurrentStatus = nil
				return r
			}(),
			response: createCEResponse{
				data: &cloudexportpb.CreateCloudExportResponse{
					Export: newIBMExportPayload(),
				},
			},
			expectedResult: newIBMExport(),
		}, {
			name: "AWS export with missing AWSProperties.Bucket field",
			request: cloud.NewAWSExport(cloud.AWSExportRequiredFields{
				Name:          "invalid-aws-export",
				PlanID:        "11467",
				AWSProperties: cloud.AWSPropertiesRequiredFields{},
			}),
			expectedResult:  nil,
			expectedError:   true,
			errorPredicates: []func(error) bool{kentikapi.IsInvalidRequestError},
		},
	}
	//nolint:dupl
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			server := newSpyCloudExportServer(t, cloudExportResponses{
				createCEResponse: tt.response,
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
			result, err := client.Cloud.CreateExport(context.Background(), tt.request)

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
				t, 1, len(server.requests.createCERequests), "invalid number of requests",
			) {
				r := server.requests.createCERequests[0]
				assert.Equal(t, dummyAuthEmail, r.metadata.Get(authEmailKey)[0])
				assert.Equal(t, dummyAuthToken, r.metadata.Get(authAPITokenKey)[0])
				testutil.AssertProtoEqual(t, tt.expectedRequest, r.data)
			}
			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestClient_Cloud_UpdateExport(t *testing.T) {
	tests := []struct {
		name            string
		request         *cloud.Export
		expectedRequest *cloudexportpb.UpdateCloudExportRequest
		response        updateCEResponse
		expectedResult  *cloud.Export
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
			request: newAWSExport(),
			expectedRequest: func() *cloudexportpb.UpdateCloudExportRequest {
				r := &cloudexportpb.UpdateCloudExportRequest{
					Export: newAWSExportPayload(),
				}
				// read-only fields should be omitted
				r.Export.CurrentStatus = nil
				return r
			}(),
			response: updateCEResponse{
				data: &cloudexportpb.UpdateCloudExportResponse{Export: nil},
			},
			expectedResult: nil,
			expectedError:  true, // InvalidResponse
		}, {
			name:    "AWS export updated",
			request: newAWSExport(),
			expectedRequest: func() *cloudexportpb.UpdateCloudExportRequest {
				r := &cloudexportpb.UpdateCloudExportRequest{
					Export: newAWSExportPayload(),
				}
				// read-only fields should be omitted
				r.Export.CurrentStatus = nil
				return r
			}(),
			response: updateCEResponse{
				data: &cloudexportpb.UpdateCloudExportResponse{
					Export: newAWSExportPayload(),
				},
			},
			expectedResult: newAWSExport(),
		}, {
			name: "AWS export with missing AWSProperties.Bucket field",
			request: cloud.NewAWSExport(cloud.AWSExportRequiredFields{
				Name:          "invalid-aws-export",
				PlanID:        "11467",
				AWSProperties: cloud.AWSPropertiesRequiredFields{},
			}),
			expectedResult:  nil,
			expectedError:   true,
			errorPredicates: []func(error) bool{kentikapi.IsInvalidRequestError},
		},
	}
	//nolint:dupl
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			server := newSpyCloudExportServer(t, cloudExportResponses{
				updateCEResponse: tt.response,
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
			result, err := client.Cloud.UpdateExport(context.Background(), tt.request)

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
				t, 1, len(server.requests.updateCERequests), "invalid number of requests",
			) {
				r := server.requests.updateCERequests[0]
				assert.Equal(t, dummyAuthEmail, r.metadata.Get(authEmailKey)[0])
				assert.Equal(t, dummyAuthToken, r.metadata.Get(authAPITokenKey)[0])
				testutil.AssertProtoEqual(t, tt.expectedRequest, r.data)
			}

			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestClient_Cloud_DeleteExport(t *testing.T) {
	tests := []struct {
		name            string
		requestID       string
		expectedRequest *cloudexportpb.DeleteCloudExportRequest
		response        deleteCEResponse
		expectedError   bool
		errorPredicates []func(error) bool
	}{
		{
			name:            "status InvalidArgument received",
			requestID:       "13",
			expectedRequest: &cloudexportpb.DeleteCloudExportRequest{Id: "13"},
			response: deleteCEResponse{
				data: &cloudexportpb.DeleteCloudExportResponse{},
				err:  status.Errorf(codes.InvalidArgument, codes.InvalidArgument.String()),
			},
			expectedError:   true,
			errorPredicates: []func(error) bool{kentikapi.IsInvalidRequestError},
		}, {
			name:            "resource deleted",
			requestID:       "13",
			expectedRequest: &cloudexportpb.DeleteCloudExportRequest{Id: "13"},
			response: deleteCEResponse{
				data: &cloudexportpb.DeleteCloudExportResponse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			server := newSpyCloudExportServer(t, cloudExportResponses{
				deleteCEResponse: tt.response,
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
			err = client.Cloud.DeleteExport(context.Background(), tt.requestID)

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

			if assert.Equal(t, 1, len(server.requests.deleteCERequests), "invalid number of requests") {
				r := server.requests.deleteCERequests[0]
				assert.Equal(t, dummyAuthEmail, r.metadata.Get(authEmailKey)[0])
				assert.Equal(t, dummyAuthToken, r.metadata.Get(authAPITokenKey)[0])
				testutil.AssertProtoEqual(t, tt.expectedRequest, r.data)
			}
		})
	}
}

type spyCloudExportServer struct {
	cloudexportpb.UnimplementedCloudExportAdminServiceServer
	server *grpc.Server

	url  string
	done chan struct{}
	t    testing.TB
	// responses to return to the client
	responses cloudExportResponses

	// requests spied by the server
	requests cloudExportRequests
}

type cloudExportRequests struct {
	listCERequests   []listCERequest
	getCERequests    []getCERequest
	createCERequests []createCERequest
	updateCERequests []updateCERequest
	deleteCERequests []deleteCERequest
}

type listCERequest struct {
	metadata metadata.MD
	data     *cloudexportpb.ListCloudExportRequest
}

type getCERequest struct {
	metadata metadata.MD
	data     *cloudexportpb.GetCloudExportRequest
}

type createCERequest struct {
	metadata metadata.MD
	data     *cloudexportpb.CreateCloudExportRequest
}

type updateCERequest struct {
	metadata metadata.MD
	data     *cloudexportpb.UpdateCloudExportRequest
}

type deleteCERequest struct {
	metadata metadata.MD
	data     *cloudexportpb.DeleteCloudExportRequest
}

type cloudExportResponses struct {
	listCEResponse   listCEResponse
	getCEResponse    getCEResponse
	createCEResponse createCEResponse
	updateCEResponse updateCEResponse
	deleteCEResponse deleteCEResponse
}

type listCEResponse struct {
	data *cloudexportpb.ListCloudExportResponse
	err  error
}

type getCEResponse struct {
	data *cloudexportpb.GetCloudExportResponse
	err  error
}

type createCEResponse struct {
	data *cloudexportpb.CreateCloudExportResponse
	err  error
}

type updateCEResponse struct {
	data *cloudexportpb.UpdateCloudExportResponse
	err  error
}

type deleteCEResponse struct {
	data *cloudexportpb.DeleteCloudExportResponse
	err  error
}

func newSpyCloudExportServer(t testing.TB, r cloudExportResponses) *spyCloudExportServer {
	return &spyCloudExportServer{
		done:      make(chan struct{}),
		t:         t,
		responses: r,
	}
}

func (s *spyCloudExportServer) Start() {
	l, err := net.Listen("tcp", "localhost:0")
	require.NoError(s.t, err)

	s.url = l.Addr().String()
	s.server = grpc.NewServer()
	cloudexportpb.RegisterCloudExportAdminServiceServer(s.server, s)

	go func() {
		err = s.server.Serve(l)
		if !errors.Is(err, grpc.ErrServerStopped) {
			assert.NoError(s.t, err)
		}
		s.done <- struct{}{}
	}()
}

// Stop blocks until the server is stopped.
func (s *spyCloudExportServer) Stop() {
	s.server.GracefulStop()
	<-s.done
}

func (s *spyCloudExportServer) ListCloudExport(
	ctx context.Context, req *cloudexportpb.ListCloudExportRequest,
) (*cloudexportpb.ListCloudExportResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	s.requests.listCERequests = append(s.requests.listCERequests, listCERequest{
		metadata: md,
		data:     req,
	})

	return s.responses.listCEResponse.data, s.responses.listCEResponse.err
}

func (s *spyCloudExportServer) GetCloudExport(
	ctx context.Context, req *cloudexportpb.GetCloudExportRequest,
) (*cloudexportpb.GetCloudExportResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	s.requests.getCERequests = append(s.requests.getCERequests, getCERequest{
		metadata: md,
		data:     req,
	})

	return s.responses.getCEResponse.data, s.responses.getCEResponse.err
}

func (s *spyCloudExportServer) CreateCloudExport(
	ctx context.Context, req *cloudexportpb.CreateCloudExportRequest,
) (*cloudexportpb.CreateCloudExportResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	s.requests.createCERequests = append(s.requests.createCERequests, createCERequest{
		metadata: md,
		data:     req,
	})

	return s.responses.createCEResponse.data, s.responses.createCEResponse.err
}

func (s *spyCloudExportServer) UpdateCloudExport(
	ctx context.Context, req *cloudexportpb.UpdateCloudExportRequest,
) (*cloudexportpb.UpdateCloudExportResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	s.requests.updateCERequests = append(s.requests.updateCERequests, updateCERequest{
		metadata: md,
		data:     req,
	})

	return s.responses.updateCEResponse.data, s.responses.updateCEResponse.err
}

func (s *spyCloudExportServer) DeleteCloudExport(
	ctx context.Context, req *cloudexportpb.DeleteCloudExportRequest,
) (*cloudexportpb.DeleteCloudExportResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	s.requests.deleteCERequests = append(s.requests.deleteCERequests, deleteCERequest{
		metadata: md,
		data:     req,
	})

	return s.responses.deleteCEResponse.data, s.responses.deleteCEResponse.err
}

func newAWSExport() *cloud.Export {
	ce := newExport()
	ce.ID = awsExportID
	ce.Provider = cloud.ProviderAWS
	ce.Properties = &cloud.AWSProperties{
		Bucket:          "dummy-bucket",
		IAMRoleARN:      "arn:aws:iam::003740049406:role/trafficTerraformIngestRole",
		Region:          "us-east-2",
		DeleteAfterRead: pointer.ToBool(true),
		MultipleBuckets: pointer.ToBool(false),
	}
	return ce
}

func newAzureExport() *cloud.Export {
	ce := newExport()
	ce.ID = azureExportID
	ce.Provider = cloud.ProviderAzure
	ce.Properties = &cloud.AzureProperties{
		Location:                 "dummy-location",
		ResourceGroup:            "dummy-rg",
		StorageAccount:           "dummy-sa",
		SubscriptionID:           "dummy-sid",
		SecurityPrincipalEnabled: pointer.ToBool(true),
	}
	return ce
}

func newGCEExport() *cloud.Export {
	ce := newExport()
	ce.ID = gceExportID
	ce.Provider = cloud.ProviderGCE
	ce.Properties = &cloud.GCEProperties{
		Project:      "dummy-project",
		Subscription: "dummy-subscription",
	}
	return ce
}

func newIBMExport() *cloud.Export {
	ce := newExport()
	ce.ID = ibmExportID
	ce.Provider = cloud.ProviderIBM
	ce.Properties = &cloud.IBMProperties{
		Bucket: "dummy-bucket",
	}
	return ce
}

func newExport() *cloud.Export {
	return &cloud.Export{
		Type:        cloud.ExportTypeKentikManaged,
		Enabled:     pointer.ToBool(true),
		Name:        "full-export",
		Description: "Export with all fields set", // including read-only fields
		PlanID:      "11467",
		BGP: &cloud.BGPProperties{
			ApplyBGP:       pointer.ToBool(true),
			UseBGPDeviceID: "dummy-device-id",
			DeviceBGPType:  "dummy-device-bgp-type",
		},
		CurrentStatus: &cloud.ExportStatus{
			Status:               "ERROR",
			ErrorMessage:         "BucketRegionError: incorrect region ...",
			FlowFound:            pointer.ToBool(false),
			APIAccess:            pointer.ToBool(true),
			StorageAccountAccess: pointer.ToBool(false),
		},
	}
}

func newAWSExportPayload() *cloudexportpb.CloudExport {
	ce := newExportPayload()
	ce.Id = awsExportID
	ce.CloudProvider = "aws"
	ce.Properties = &cloudexportpb.CloudExport_Aws{
		Aws: &cloudexportpb.AwsProperties{
			Bucket:          "dummy-bucket",
			IamRoleArn:      "arn:aws:iam::003740049406:role/trafficTerraformIngestRole",
			Region:          "us-east-2",
			DeleteAfterRead: true,
			MultipleBuckets: false,
		},
	}
	return ce
}

func newAzureExportPayload() *cloudexportpb.CloudExport {
	ce := newExportPayload()
	ce.Id = azureExportID
	ce.CloudProvider = "azure"
	ce.Properties = &cloudexportpb.CloudExport_Azure{
		Azure: &cloudexportpb.AzureProperties{
			Location:                 "dummy-location",
			ResourceGroup:            "dummy-rg",
			StorageAccount:           "dummy-sa",
			SubscriptionId:           "dummy-sid",
			SecurityPrincipalEnabled: true,
		},
	}
	return ce
}

func newGCEExportPayload() *cloudexportpb.CloudExport {
	ce := newExportPayload()
	ce.Id = gceExportID
	ce.CloudProvider = "gce"
	ce.Properties = &cloudexportpb.CloudExport_Gce{
		Gce: &cloudexportpb.GceProperties{
			Project:      "dummy-project",
			Subscription: "dummy-subscription",
		},
	}
	return ce
}

func newIBMExportPayload() *cloudexportpb.CloudExport {
	ce := newExportPayload()
	ce.Id = ibmExportID
	ce.CloudProvider = "ibm"
	ce.Properties = &cloudexportpb.CloudExport_Ibm{
		Ibm: &cloudexportpb.IbmProperties{
			Bucket: "dummy-bucket",
		},
	}
	return ce
}

// newExportPayload returns payload with all fields set.
// ApiRoot and FlowDest are going to be removed from the API and are omitted.
func newExportPayload() *cloudexportpb.CloudExport {
	return &cloudexportpb.CloudExport{
		Type:        cloudexportpb.CloudExportType_CLOUD_EXPORT_TYPE_KENTIK_MANAGED,
		Enabled:     true,
		Name:        "full-export",
		Description: "Export with all fields set", // including read-only fields
		PlanId:      "11467",
		Bgp: &cloudexportpb.BgpProperties{
			ApplyBgp:       true,
			UseBgpDeviceId: "dummy-device-id",
			DeviceBgpType:  "dummy-device-bgp-type",
		},
		CurrentStatus: &cloudexportpb.Status{
			Status:               "ERROR",
			ErrorMessage:         "BucketRegionError: incorrect region ...",
			FlowFound:            &wrapperspb.BoolValue{Value: false},
			ApiAccess:            &wrapperspb.BoolValue{Value: true},
			StorageAccountAccess: &wrapperspb.BoolValue{Value: false},
		},
	}
}
