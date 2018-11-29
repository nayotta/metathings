package metathings_device_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/nayotta/metathings/pkg/proto/device"
)

func (self *MetathingsDeviceServiceImpl) Heartbeat(ctx context.Context, req *pb.HeartbeatRequest) (*empty.Empty, error) {
	op_mdl := req.GetModule()
	component := op_mdl.GetComponent().GetValue()
	name := op_mdl.GetName().GetValue()

	mdl, err := self.mdl_db.Lookup(component, name)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to lookup module")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	mdl.Heartbeat()

	self.logger.WithFields(log.Fields{
		"component": component,
		"name":      name,
	}).Debugf("module heartbeat")

	return &empty.Empty{}, nil
}
