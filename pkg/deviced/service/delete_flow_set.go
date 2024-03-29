package metathings_deviced_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/deviced"
)

func (self *MetathingsDevicedService) ValidateDeleteFlowSet(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() flow_set_getter {
				req := in.(*pb.DeleteFlowSetRequest)
				return req
			},
		},
		identityd_validator.Invokers{ensure_get_flow_set_id},
	)
}

func (self *MetathingsDevicedService) AuthorizeDeleteFlowSet(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.DeleteFlowSetRequest).GetFlowSet().GetId().GetValue(), "deviced:delete_flow_set")
}

func (self *MetathingsDevicedService) DeleteFlowSet(ctx context.Context, req *pb.DeleteFlowSetRequest) (*empty.Empty, error) {
	var flwst *storage.FlowSet
	var err error

	flwst_id_str := req.GetFlowSet().GetId().GetValue()
	logger := self.get_logger().WithField("flow_set", flwst_id_str)

	if flwst, err = self.storage.GetFlowSet(ctx, flwst_id_str); err != nil {
		logger.WithError(err).Errorf("failed to get flow set in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	for _, flw := range flwst.Flows {
		if err = self.storage.RemoveFlowFromFlowSet(ctx, flwst_id_str, *flw.Id); err != nil {
			logger.WithError(err).Errorf("failed to remove flow from flow set")
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	if err = self.storage.DeleteFlowSet(ctx, flwst_id_str); err != nil {
		logger.WithError(err).Errorf("failed to delete flow set in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	logger.Infof("delete flow set")

	return &empty.Empty{}, nil
}
