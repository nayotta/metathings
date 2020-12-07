package metathings_deviced_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/nayotta/metathings/proto/deviced"
)

func (self *MetathingsDevicedService) AuthorizeAddFlowsToFlowSet(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.AddFlowsToFlowSetRequest).GetFlowSet().GetId().GetValue(), "deviced:add_flows_to_flow_set")
}

func (self *MetathingsDevicedService) AddFlowsToFlowSet(ctx context.Context, req *pb.AddFlowsToFlowSetRequest) (*empty.Empty, error) {
	var err error

	logger := self.get_logger()

	flwst := req.GetFlowSet()
	flwst_id := flwst.GetId().GetValue()

	devs := req.GetDevices()
	flw_ids, err := self.get_flow_ids_by_devices(ctx, devs)
	if err != nil {
		logger.WithError(err).Errorf("failed to get flow ids by devices")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	logger = logger.WithFields(log.Fields{
		"flow_set": flwst_id,
		"flows":    flw_ids,
	})

	flwst_s, err := self.storage.GetFlowSet(ctx, flwst_id)
	if err != nil {
		logger.WithError(err).Errorf("failed to get flow set in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var flw_ids_expect []string
	for _, flw_id := range flw_ids {
		exists := false
		for _, flw := range flwst_s.Flows {
			if *flw.Id == flw_id {
				exists = true
				break
			}
		}
		if !exists {
			flw_ids_expect = append(flw_ids_expect, flw_id)
		}

	}

	for _, flw_id := range flw_ids_expect {
		if err = self.storage.AddFlowToFlowSet(ctx, flwst_id, flw_id); err != nil {
			logger.WithError(err).Errorf("failed to add flow to flow set")
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	// TODO(Peer): dynamic add flow to pushing flow.
	logger.Infof("add flows to flow set")

	return &empty.Empty{}, nil
}
