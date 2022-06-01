package cloud

import (
	"context"

	cloudexportpb "github.com/kentik/api-schema-public/gen/go/kentik/cloud_export/v202101beta1"
	"github.com/kentik/community_sdk_golang/kentikapi/cloud"
	kentikerrors "github.com/kentik/community_sdk_golang/kentikapi/internal/errors"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"google.golang.org/grpc"
)

// API aggregates Cloud Exports API methods.
type API struct {
	client cloudexportpb.CloudExportAdminServiceClient
}

// NewAPI creates new API.
func NewAPI(cc grpc.ClientConnInterface) *API {
	return &API{
		client: cloudexportpb.NewCloudExportAdminServiceClient(cc),
	}
}

// GetAllExports lists cloud exports.
func (a *API) GetAllExports(ctx context.Context) (*cloud.GetAllExportsResponse, error) {
	response, err := a.client.ListCloudExport(ctx, &cloudexportpb.ListCloudExportRequest{})
	if err != nil {
		return nil, kentikerrors.StatusErrorFromGRPC(err)
	}

	return (*listExportsResponse)(response).ToModel()
}

// GetExport retrieves cloud export with given ID.
func (a *API) GetExport(ctx context.Context, id models.ID) (*cloud.Export, error) {
	response, err := a.client.GetCloudExport(ctx, &cloudexportpb.GetCloudExportRequest{Id: id})
	if err != nil {
		return nil, kentikerrors.StatusErrorFromGRPC(err)
	}

	return exportFromPayload(response.GetExport())
}

// CreateExport creates new cloud export.
func (a *API) CreateExport(ctx context.Context, ce *cloud.Export) (*cloud.Export, error) {
	if err := validateCreateExportRequest(ce); err != nil {
		return nil, err
	}
	payload, err := exportToPayload(ce)
	if err != nil {
		return nil, kentikerrors.StatusErrorFromGRPC(err)
	}

	response, err := a.client.CreateCloudExport(ctx, &cloudexportpb.CreateCloudExportRequest{
		Export: payload,
	})
	if err != nil {
		return nil, kentikerrors.StatusErrorFromGRPC(err)
	}

	return exportFromPayload(response.GetExport())
}

// UpdateExport updates the cloud export.
func (a *API) UpdateExport(ctx context.Context, ce *cloud.Export) (*cloud.Export, error) {
	if err := validateExportUpdateRequest(ce); err != nil {
		return nil, err
	}
	payload, err := exportToPayload(ce)
	if err != nil {
		return nil, kentikerrors.StatusErrorFromGRPC(err)
	}

	response, err := a.client.UpdateCloudExport(ctx, &cloudexportpb.UpdateCloudExportRequest{
		Export: payload,
	})
	if err != nil {
		return nil, kentikerrors.StatusErrorFromGRPC(err)
	}

	return exportFromPayload(response.GetExport())
}

// DeleteExport removes cloud export with given ID.
func (a *API) DeleteExport(ctx context.Context, id models.ID) error {
	_, err := a.client.DeleteCloudExport(ctx, &cloudexportpb.DeleteCloudExportRequest{Id: id})
	return kentikerrors.StatusErrorFromGRPC(err)
}
