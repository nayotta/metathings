package metathings_deviced_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) AuthorizeRemoveFlowsFromFlowSet(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.RemoveFlowsFromFlowSetRequest).GetFlowSet().GetId().GetValue(), "deviced:remove_flows_from_flow_set")
}

func (self *MetathingsDevicedService) RemoveFlowsFromFlowSet(ctx context.Context, req *pb.RemoveFlowsFromFlowSetRequest) (*empty.Empty, error) {
	var err error

	flwst := req.GetFlowSet()
	flwst_id := flwst.GetId().GetValue()

	devs := req.GetDevices()
	flw_ids, err := self.get_flow_ids_by_devices(ctx, devs)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to get flow ids by devices")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	for _, flw_id := range flw_ids {
		if err = self.storage.RemoveFlowFromFlowSet(ctx, flwst_id, flw_id); err != nil {
			self.logger.WithError(err).Errorf("failed to remove flow from flow set")
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	// TODO(Peer): dynamic remove flow from pulling flow set.

	self.logger.WithFields(log.Fields{
		"flow_set": flwst_id,
		"flows":    flw_ids,
	}).Infof("remove flows from flow set")

	return &empty.Empty{}, nil
}
