package resources

import (
	"context"
	"fmt"

	cloudexportpb "github.com/kentik/api-schema-public/gen/go/kentik/cloud_export/v202101beta1"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"google.golang.org/grpc"
)

// CloudExportsAPI aggregates Cloud Exports API methods.
type CloudExportsAPI struct {
	client cloudexportpb.CloudExportAdminServiceClient
}

// NewCloudExportsAPI creates new CloudExportsAPI.
func NewCloudExportsAPI(cc grpc.ClientConnInterface) *CloudExportsAPI {
	return &CloudExportsAPI{
		client: cloudexportpb.NewCloudExportAdminServiceClient(cc),
	}
}

// GetAll lists Cloud Exports.
func (a *CloudExportsAPI) GetAll(ctx context.Context) (*models.GetAllCloudExportsResponse, error) {
	response, err := a.client.ListCloudExport(ctx, &cloudexportpb.ListCloudExportRequest{})
	if err != nil {
		return nil, err
	}

	return (*api_payloads.ListCloudExportsResponse)(response).ToModel(), nil
}

// Get retrieves Cloud Export with given ID.
func (a *CloudExportsAPI) Get(ctx context.Context, id models.ID) (*models.CloudExport, error) {
	response, err := a.client.GetCloudExport(ctx, &cloudexportpb.GetCloudExportRequest{Id: id})
	if err != nil {
		return nil, err
	}

	obj := api_payloads.CloudExportFromPayload(response.GetExport())
	if obj == nil {
		return nil, fmt.Errorf("get cloud export with ID %v: API returned no data and no error", id)
	}

	return obj, nil
}

// Create creates new Cloud Export.
func (a *CloudExportsAPI) Create(ctx context.Context, ce *models.CloudExport) (*models.CloudExport, error) {
	// TODO(dfurman): add request validation
	payload, err := api_payloads.CloudExportToPayload(ce)
	if err != nil {
		return nil, err
	}

	response, err := a.client.CreateCloudExport(ctx, &cloudexportpb.CreateCloudExportRequest{
		Export: payload,
	})
	if err != nil {
		return nil, err
	}

	return api_payloads.CloudExportFromPayload(response.GetExport()), nil
}

// Update updates the Cloud Export.
func (a *CloudExportsAPI) Update(ctx context.Context, ce *models.CloudExport) (*models.CloudExport, error) {
	// TODO(dfurman): add request validation
	payload, err := api_payloads.CloudExportToPayload(ce)
	if err != nil {
		return nil, err
	}

	response, err := a.client.UpdateCloudExport(ctx, &cloudexportpb.UpdateCloudExportRequest{
		Export: payload,
	})
	if err != nil {
		return nil, err
	}

	return api_payloads.CloudExportFromPayload(response.GetExport()), nil
}

// Delete removes Cloud Export with given ID.
func (a *CloudExportsAPI) Delete(ctx context.Context, id models.ID) error {
	_, err := a.client.DeleteCloudExport(ctx, &cloudexportpb.DeleteCloudExportRequest{Id: id})
	return err
}
