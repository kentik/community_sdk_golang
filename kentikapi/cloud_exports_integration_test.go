package kentikapi_test

import (
	"context"
	"net"
	"testing"

	"github.com/AlekSi/pointer"
	cloudexportpb "github.com/kentik/api-schema-public/gen/go/kentik/cloud_export/v202101beta1"
	"github.com/kentik/community_sdk_golang/kentikapi"
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
	awsCloudExportID   = "1001"
	azureCloudExportID = "1002"
	gceCloudExportID   = "1003"
	ibmCloudExportID   = "1004"
)

func TestClient_GetAllCloudExports(t *testing.T) {
	tests := []struct {
		name              string
		response          listCEResponse
		expectedResult    *models.GetAllCloudExportsResponse
		expectedError     bool
		expectedErrorCode *codes.Code
	}{
		{
			name: "status InvalidArgument received",
			response: listCEResponse{
				err: status.Errorf(codes.InvalidArgument, codes.InvalidArgument.String()),
			},
			expectedError:     true,
			expectedErrorCode: codePtr(codes.InvalidArgument),
		}, {
			name: "empty response received",
			response: listCEResponse{
				data: &cloudexportpb.ListCloudExportResponse{},
			},
			expectedResult: &models.GetAllCloudExportsResponse{
				CloudExports:             nil,
				InvalidCloudExportsCount: 0,
			},
		}, {
			name: "no exports received",
			response: listCEResponse{
				data: &cloudexportpb.ListCloudExportResponse{
					Exports:             []*cloudexportpb.CloudExport{},
					InvalidExportsCount: 0,
				},
			},
			expectedResult: &models.GetAllCloudExportsResponse{
				CloudExports:             nil,
				InvalidCloudExportsCount: 0,
			},
		}, {
			name: "4 exports received",
			response: listCEResponse{
				data: &cloudexportpb.ListCloudExportResponse{
					Exports: []*cloudexportpb.CloudExport{
						newFullAWSCloudExportPayload(),
						newFullAzureCloudExportPayload(),
						newFullGCECloudExportPayload(),
						newFullIBMCloudExportPayload(),
					},
					InvalidExportsCount: 1,
				},
			},
			expectedResult: &models.GetAllCloudExportsResponse{
				CloudExports: []models.CloudExport{
					*newFullAWSCloudExport(),
					*newFullAzureCloudExport(),
					*newFullGCECloudExport(),
					*newFullIBMCloudExport(),
				},
				InvalidCloudExportsCount: 1,
			},
		}, {
			name: "2 exports received - one empty",
			response: listCEResponse{
				data: &cloudexportpb.ListCloudExportResponse{
					Exports: []*cloudexportpb.CloudExport{
						newFullAWSCloudExportPayload(),
						nil,
					},
					InvalidExportsCount: 0,
				},
			},
			expectedResult: &models.GetAllCloudExportsResponse{
				CloudExports: []models.CloudExport{
					*newFullAWSCloudExport(),
					// client receives initialized, empty CE here
					{
						Type:    models.CloudExportTypeUnspecified,
						Enabled: pointer.ToBool(false),
					},
				},
				InvalidCloudExportsCount: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			server := newSpyCloudExportServer(t, cloudExportResponses{
				listCEResponse: tt.response,
			})
			server.Start()
			defer server.Stop()

			client, err := kentikapi.NewClient(kentikapi.Config{
				CloudExportHostPort: server.url,
				AuthToken:           dummyAuthToken,
				AuthEmail:           dummyAuthEmail,
				DisableTLS:          true,
			})
			require.NoError(t, err)

			// act
			result, err := client.CloudExports.GetAll(context.Background())

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

func TestClient_GetCloudExport(t *testing.T) {
	tests := []struct {
		name              string
		requestID         models.ID
		expectedRequest   *cloudexportpb.GetCloudExportRequest
		response          getCEResponse
		expectedResult    *models.CloudExport
		expectedError     bool
		expectedErrorCode *codes.Code
	}{
		{
			name:            "status InvalidArgument received",
			requestID:       "13",
			expectedRequest: &cloudexportpb.GetCloudExportRequest{Id: "13"},
			response: getCEResponse{
				err: status.Errorf(codes.InvalidArgument, codes.InvalidArgument.String()),
			},
			expectedError:     true,
			expectedErrorCode: codePtr(codes.InvalidArgument),
		}, {
			name:            "status NotFound received",
			requestID:       "13",
			expectedRequest: &cloudexportpb.GetCloudExportRequest{Id: "13"},
			response: getCEResponse{
				err: status.Errorf(codes.NotFound, codes.NotFound.String()),
			},
			expectedError:     true,
			expectedErrorCode: codePtr(codes.NotFound),
		}, {
			name:            "empty response received",
			requestID:       "13",
			expectedRequest: &cloudexportpb.GetCloudExportRequest{Id: "13"},
			response: getCEResponse{
				data: &cloudexportpb.GetCloudExportResponse{},
			},
			expectedError: true,
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
			expectedResult: &models.CloudExport{
				ID:            "58192",
				Type:          models.CloudExportTypeKentikManaged,
				Enabled:       pointer.ToBool(true),
				Name:          "minimal-aws-export",
				Description:   "",
				PlanID:        "11467",
				CloudProvider: models.CloudProviderAWS,
				AWSProperties: &models.AWSProperties{
					Bucket:          "dummy-bucket",
					IAMRoleARN:      "",
					Region:          "",
					DeleteAfterRead: pointer.ToBool(false),
					MultipleBuckets: pointer.ToBool(false),
				},
				CurrentStatus: &models.CloudExportStatus{
					Status:       "",
					ErrorMessage: "",
				},
			},
		}, {
			name:            "AWS cloud export received",
			requestID:       awsCloudExportID,
			expectedRequest: &cloudexportpb.GetCloudExportRequest{Id: awsCloudExportID},
			response: getCEResponse{
				data: &cloudexportpb.GetCloudExportResponse{Export: newFullAWSCloudExportPayload()},
			},
			expectedResult: newFullAWSCloudExport(),
		}, {
			name:            "Azure cloud export received",
			requestID:       azureCloudExportID,
			expectedRequest: &cloudexportpb.GetCloudExportRequest{Id: azureCloudExportID},
			response: getCEResponse{
				data: &cloudexportpb.GetCloudExportResponse{Export: newFullAzureCloudExportPayload()},
			},
			expectedResult: newFullAzureCloudExport(),
		}, {
			name:            "GCE cloud export received",
			requestID:       gceCloudExportID,
			expectedRequest: &cloudexportpb.GetCloudExportRequest{Id: gceCloudExportID},
			response: getCEResponse{
				data: &cloudexportpb.GetCloudExportResponse{Export: newFullGCECloudExportPayload()},
			},
			expectedResult: newFullGCECloudExport(),
		}, {
			name:            "IBM cloud export received",
			requestID:       ibmCloudExportID,
			expectedRequest: &cloudexportpb.GetCloudExportRequest{Id: ibmCloudExportID},
			response: getCEResponse{
				data: &cloudexportpb.GetCloudExportResponse{Export: newFullIBMCloudExportPayload()},
			},
			expectedResult: newFullIBMCloudExport(),
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

			client, err := kentikapi.NewClient(kentikapi.Config{
				CloudExportHostPort: server.url,
				AuthToken:           dummyAuthToken,
				AuthEmail:           dummyAuthEmail,
				DisableTLS:          true,
			})
			require.NoError(t, err)

			// act
			result, err := client.CloudExports.Get(context.Background(), tt.requestID)

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

func TestClient_CreateCloudExport(t *testing.T) {
	tests := []struct {
		name              string
		request           *models.CloudExport
		expectedRequest   *cloudexportpb.CreateCloudExportRequest
		response          createCEResponse
		expectedResult    *models.CloudExport
		expectedError     bool
		expectedErrorCode *codes.Code
	}{
		{
			name:            "nil request, status InvalidArgument received",
			request:         nil,
			expectedRequest: &cloudexportpb.CreateCloudExportRequest{Export: nil},
			response: createCEResponse{
				data: &cloudexportpb.CreateCloudExportResponse{},
				err:  status.Errorf(codes.InvalidArgument, codes.InvalidArgument.String()),
			},
			expectedResult:    nil,
			expectedError:     true,
			expectedErrorCode: codePtr(codes.InvalidArgument),
		}, {
			name:    "empty response received",
			request: newFullAWSCloudExport(),
			expectedRequest: func() *cloudexportpb.CreateCloudExportRequest {
				r := &cloudexportpb.CreateCloudExportRequest{
					Export: newFullAWSCloudExportPayload(),
				}
				// read-only fields should be omitted
				r.Export.CurrentStatus = nil
				return r
			}(),
			response: createCEResponse{
				data: &cloudexportpb.CreateCloudExportResponse{Export: nil},
			},
			expectedResult: nil,
			expectedError:  true,
		}, {
			name: "minimal AWS export created",
			request: models.NewAWSCloudExport(models.CloudExportAWSRequiredFields{
				Name:   "minimal-aws-export",
				PlanID: "11467",
				AWSProperties: models.AWSPropertiesRequiredFields{
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
			expectedResult: &models.CloudExport{
				ID:            "58192",
				Type:          models.CloudExportTypeKentikManaged,
				Enabled:       pointer.ToBool(true),
				Name:          "minimal-aws-export",
				Description:   "",
				PlanID:        "11467",
				CloudProvider: models.CloudProviderAWS,
				AWSProperties: &models.AWSProperties{
					Bucket:          "dummy-bucket",
					IAMRoleARN:      "",
					Region:          "",
					DeleteAfterRead: pointer.ToBool(false),
					MultipleBuckets: pointer.ToBool(false),
				},
				CurrentStatus: &models.CloudExportStatus{
					Status:       "",
					ErrorMessage: "",
				},
			},
		}, {
			name:    "full AWS export created",
			request: newFullAWSCloudExport(),
			expectedRequest: func() *cloudexportpb.CreateCloudExportRequest {
				r := &cloudexportpb.CreateCloudExportRequest{
					Export: newFullAWSCloudExportPayload(),
				}
				// read-only fields should be omitted
				r.Export.CurrentStatus = nil
				return r
			}(),
			response: createCEResponse{
				data: &cloudexportpb.CreateCloudExportResponse{
					Export: newFullAWSCloudExportPayload(),
				},
			},
			expectedResult: newFullAWSCloudExport(),
		}, {
			name:    "full Azure export created",
			request: newFullAzureCloudExport(),
			expectedRequest: func() *cloudexportpb.CreateCloudExportRequest {
				r := &cloudexportpb.CreateCloudExportRequest{
					Export: newFullAzureCloudExportPayload(),
				}
				// read-only fields should be omitted
				r.Export.CurrentStatus = nil
				return r
			}(),
			response: createCEResponse{
				data: &cloudexportpb.CreateCloudExportResponse{
					Export: newFullAzureCloudExportPayload(),
				},
			},
			expectedResult: newFullAzureCloudExport(),
		}, {
			name:    "full GCE export created",
			request: newFullGCECloudExport(),
			expectedRequest: func() *cloudexportpb.CreateCloudExportRequest {
				r := &cloudexportpb.CreateCloudExportRequest{
					Export: newFullGCECloudExportPayload(),
				}
				// read-only fields should be omitted
				r.Export.CurrentStatus = nil
				return r
			}(),
			response: createCEResponse{
				data: &cloudexportpb.CreateCloudExportResponse{
					Export: newFullGCECloudExportPayload(),
				},
			},
			expectedResult: newFullGCECloudExport(),
		}, {
			name:    "full IBM export created",
			request: newFullIBMCloudExport(),
			expectedRequest: func() *cloudexportpb.CreateCloudExportRequest {
				r := &cloudexportpb.CreateCloudExportRequest{
					Export: newFullIBMCloudExportPayload(),
				}
				// read-only fields should be omitted
				r.Export.CurrentStatus = nil
				return r
			}(),
			response: createCEResponse{
				data: &cloudexportpb.CreateCloudExportResponse{
					Export: newFullIBMCloudExportPayload(),
				},
			},
			expectedResult: newFullIBMCloudExport(),
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

			client, err := kentikapi.NewClient(kentikapi.Config{
				CloudExportHostPort: server.url,
				AuthToken:           dummyAuthToken,
				AuthEmail:           dummyAuthEmail,
				DisableTLS:          true,
			})
			require.NoError(t, err)

			// act
			result, err := client.CloudExports.Create(context.Background(), tt.request)

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

			if assert.Equal(t, 1, len(server.requests.createCERequests), "invalid number of requests") {
				r := server.requests.createCERequests[0]
				assert.Equal(t, dummyAuthEmail, r.metadata.Get(authEmailKey)[0])
				assert.Equal(t, dummyAuthToken, r.metadata.Get(authAPITokenKey)[0])
				testutil.AssertProtoEqual(t, tt.expectedRequest, r.data)
			}

			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestClient_UpdateCloudExport(t *testing.T) {
	tests := []struct {
		name              string
		request           *models.CloudExport
		expectedRequest   *cloudexportpb.UpdateCloudExportRequest
		response          updateCEResponse
		expectedResult    *models.CloudExport
		expectedError     bool
		expectedErrorCode *codes.Code
	}{
		{
			name:            "nil request, status InvalidArgument received",
			request:         nil,
			expectedRequest: &cloudexportpb.UpdateCloudExportRequest{Export: nil},
			response: updateCEResponse{
				data: &cloudexportpb.UpdateCloudExportResponse{},
				err:  status.Errorf(codes.InvalidArgument, codes.InvalidArgument.String()),
			},
			expectedResult:    nil,
			expectedError:     true,
			expectedErrorCode: codePtr(codes.InvalidArgument),
		}, {
			name:    "empty response received",
			request: newFullAWSCloudExport(),
			expectedRequest: func() *cloudexportpb.UpdateCloudExportRequest {
				r := &cloudexportpb.UpdateCloudExportRequest{
					Export: newFullAWSCloudExportPayload(),
				}
				// read-only fields should be omitted
				r.Export.CurrentStatus = nil
				return r
			}(),
			response: updateCEResponse{
				data: &cloudexportpb.UpdateCloudExportResponse{Export: nil},
			},
			expectedResult: nil,
			expectedError:  true,
		}, {
			name:    "full AWS export updated",
			request: newFullAWSCloudExport(),
			expectedRequest: func() *cloudexportpb.UpdateCloudExportRequest {
				r := &cloudexportpb.UpdateCloudExportRequest{
					Export: newFullAWSCloudExportPayload(),
				}
				// read-only fields should be omitted
				r.Export.CurrentStatus = nil
				return r
			}(),
			response: updateCEResponse{
				data: &cloudexportpb.UpdateCloudExportResponse{
					Export: newFullAWSCloudExportPayload(),
				},
			},
			expectedResult: newFullAWSCloudExport(),
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

			client, err := kentikapi.NewClient(kentikapi.Config{
				CloudExportHostPort: server.url,
				AuthToken:           dummyAuthToken,
				AuthEmail:           dummyAuthEmail,
				DisableTLS:          true,
			})
			require.NoError(t, err)

			// act
			result, err := client.CloudExports.Update(context.Background(), tt.request)

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

			if assert.Equal(t, 1, len(server.requests.updateCERequests), "invalid number of requests") {
				r := server.requests.updateCERequests[0]
				assert.Equal(t, dummyAuthEmail, r.metadata.Get(authEmailKey)[0])
				assert.Equal(t, dummyAuthToken, r.metadata.Get(authAPITokenKey)[0])
				testutil.AssertProtoEqual(t, tt.expectedRequest, r.data)
			}

			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestClient_DeleteCloudExport(t *testing.T) {
	tests := []struct {
		name              string
		requestID         string
		expectedRequest   *cloudexportpb.DeleteCloudExportRequest
		response          deleteCEResponse
		expectedError     bool
		expectedErrorCode *codes.Code
	}{
		{
			name:            "status InvalidArgument received",
			requestID:       "13",
			expectedRequest: &cloudexportpb.DeleteCloudExportRequest{Id: "13"},
			response: deleteCEResponse{
				data: &cloudexportpb.DeleteCloudExportResponse{},
				err:  status.Errorf(codes.InvalidArgument, codes.InvalidArgument.String()),
			},
			expectedError:     true,
			expectedErrorCode: codePtr(codes.InvalidArgument),
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

			client, err := kentikapi.NewClient(kentikapi.Config{
				CloudExportHostPort: server.url,
				AuthToken:           dummyAuthToken,
				AuthEmail:           dummyAuthEmail,
				DisableTLS:          true,
			})
			require.NoError(t, err)

			// act
			err = client.CloudExports.Delete(context.Background(), tt.requestID)

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
		assert.NoError(s.t, err)
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

func newFullAWSCloudExport() *models.CloudExport {
	ce := newFullCloudExport()
	ce.ID = awsCloudExportID
	ce.CloudProvider = models.CloudProviderAWS
	ce.AWSProperties = &models.AWSProperties{
		Bucket:          "dummy-bucket",
		IAMRoleARN:      "arn:aws:iam::003740049406:role/trafficTerraformIngestRole",
		Region:          "us-east-2",
		DeleteAfterRead: pointer.ToBool(true),
		MultipleBuckets: pointer.ToBool(false),
	}
	return ce
}

func newFullAzureCloudExport() *models.CloudExport {
	ce := newFullCloudExport()
	ce.ID = azureCloudExportID
	ce.CloudProvider = models.CloudProviderAzure
	ce.AzureProperties = &models.AzureProperties{
		Location:                 "dummy-location",
		ResourceGroup:            "dummy-rg",
		StorageAccount:           "dummy-sa",
		SubscriptionID:           "dummy-sid",
		SecurityPrincipalEnabled: pointer.ToBool(true),
	}
	return ce
}

func newFullGCECloudExport() *models.CloudExport {
	ce := newFullCloudExport()
	ce.ID = gceCloudExportID
	ce.CloudProvider = models.CloudProviderGCE
	ce.GCEProperties = &models.GCEProperties{
		Project:      "dummy-project",
		Subscription: "dummy-subscription",
	}
	return ce
}

func newFullIBMCloudExport() *models.CloudExport {
	ce := newFullCloudExport()
	ce.ID = ibmCloudExportID
	ce.CloudProvider = models.CloudProviderIBM
	ce.IBMProperties = &models.IBMProperties{
		Bucket: "dummy-bucket",
	}
	return ce
}

func newFullCloudExport() *models.CloudExport {
	return &models.CloudExport{
		Type:        models.CloudExportTypeKentikManaged,
		Enabled:     pointer.ToBool(true),
		Name:        "full-export",
		Description: "Export with all fields set", // including read-only fields
		PlanID:      "11467",
		BGP: &models.BGPProperties{
			ApplyBGP:       pointer.ToBool(true),
			UseBGPDeviceID: "dummy-device-id",
			DeviceBGPType:  "dummy-device-bgp-type",
		},
		CurrentStatus: &models.CloudExportStatus{
			Status:               "ERROR",
			ErrorMessage:         "BucketRegionError: incorrect region ...",
			FlowFound:            pointer.ToBool(false),
			APIAccess:            pointer.ToBool(true),
			StorageAccountAccess: pointer.ToBool(false),
		},
	}
}

func newFullAWSCloudExportPayload() *cloudexportpb.CloudExport {
	ce := newFullCloudExportPayload()
	ce.Id = awsCloudExportID
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

func newFullAzureCloudExportPayload() *cloudexportpb.CloudExport {
	ce := newFullCloudExportPayload()
	ce.Id = azureCloudExportID
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

func newFullGCECloudExportPayload() *cloudexportpb.CloudExport {
	ce := newFullCloudExportPayload()
	ce.Id = gceCloudExportID
	ce.CloudProvider = "gce"
	ce.Properties = &cloudexportpb.CloudExport_Gce{
		Gce: &cloudexportpb.GceProperties{
			Project:      "dummy-project",
			Subscription: "dummy-subscription",
		},
	}
	return ce
}

func newFullIBMCloudExportPayload() *cloudexportpb.CloudExport {
	ce := newFullCloudExportPayload()
	ce.Id = ibmCloudExportID
	ce.CloudProvider = "ibm"
	ce.Properties = &cloudexportpb.CloudExport_Ibm{
		Ibm: &cloudexportpb.IbmProperties{
			Bucket: "dummy-bucket",
		},
	}
	return ce
}

// newFullCloudExportPayload returns payload with all fields set.
// ApiRoot and FlowDest are going to be removed from the API and are omitted.
func newFullCloudExportPayload() *cloudexportpb.CloudExport {
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

func codePtr(c codes.Code) *codes.Code {
	return &c
}
