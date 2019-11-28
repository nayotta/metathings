package metathings_deviced_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) AuthorizeAddFlowsToFlowSet(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.AddFlowsToFlowSetRequest).GetFlowSet().GetId().GetValue(), "deviced:add_flows_to_flow_set")
}

func (self *MetathingsDevicedService) AddFlowsToFlowSet(ctx context.Context, req *pb.AddFlowsToFlowSetRequest) (*empty.Empty, error) {
	var err error

	flws := req.GetFlows()
	flwst := req.GetFlowSet()

	flwst_id := flwst.GetId().GetValue()
	flw_ids := []string{}

	for _, flw := range flws {
		flw_id := flw.GetId().GetValue()
		flw_ids = append(flw_ids, flw_id)
		if err = self.storage.AddFlowToFlowSet(flwst_id, flw_id); err != nil {
			self.logger.WithError(err).Errorf("failed to add flow to flow set")
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	// TODO(Peer): dynamic add flow to pulling flow set.

	self.logger.WithFields(log.Fields{
		"flow_set": flwst_id,
		"flows":    flw_ids,
	}).Infof("add flows to flow set")

	return &empty.Empty{}, nil
}
