package metathings_deviced_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) ValidateFlowSet(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, flow_set_getter) {
				req := in.(*pb.GetFlowSetRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{ensure_get_flow_set_id},
	)
}

func (self *MetathingsDevicedService) AuthorizeGetFlowSet(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.GetFlowSetRequest).GetFlowSet().GetId().GetValue(), "deviced:get_flow_set")
}

func (self *MetathingsDevicedService) GetFlowSet(ctx context.Context, req *pb.GetFlowSetRequest) (*pb.GetFlowSetResponse, error) {
	var flwst_s *storage.FlowSet
	var err error

	flwst_id_str := req.GetFlowSet().GetId().GetValue()
	if flwst_s, err = self.storage.GetFlowSet(ctx, flwst_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to get flow set in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.GetFlowSetResponse{
		FlowSet: copy_flow_set(flwst_s),
	}

	self.logger.WithField("id", flwst_id_str).Debugf("get flow set")

	return res, nil
}
