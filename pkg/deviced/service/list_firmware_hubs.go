package metathings_deviced_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) AuthorizeListFirmwareHubs(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, self.opt.Domain, "deviced:list_firmware_hubs")
}

func (self *MetathingsDevicedService) ListFirmwareHubs(ctx context.Context, req *pb.ListFirmwareHubsRequest) (*pb.ListFirmwareHubsResponse, error) {
	var frm_hubs_s []*storage.FirmwareHub
	var err error

	frm_hub_s := &storage.FirmwareHub{}
	frm_hub := req.GetFirmwareHub()
	logger := self.logger

	if id := frm_hub.GetId(); id != nil {
		frm_hub_s.Id = &id.Value
	}

	if alias := frm_hub.GetAlias(); alias != nil {
		frm_hub_s.Alias = &alias.Value
	}

	if frm_hubs_s, err = self.storage.ListFirmwareHubs(ctx, frm_hub_s); err != nil {
		logger.WithError(err).Errorf("failed to list firmware hubs in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ListFirmwareHubsResponse{
		FirmwareHubs: copy_firmware_hubs(frm_hubs_s),
	}

	logger.Debugf("list firmware hubs")

	return res, nil
}
