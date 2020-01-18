package metathings_deviced_service

import (
	"context"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (self *MetathingsDevicedService) ValidatePatchFlowSet(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, flow_set_getter) {
				req := in.(*pb.PatchFlowSetRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{ensure_get_flow_set_id},
	)
}

func (self *MetathingsDevicedService) AuthorizePatchFlowSet(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.PatchFlowSetRequest).GetFlowSet().GetId().GetValue(), "deviced:patch_flow_set")
}

func (self *MetathingsDevicedService) PatchFlowSet(ctx context.Context, req *pb.PatchFlowSetRequest) (*pb.PatchFlowSetResponse, error) {
	flwst_s := &storage.FlowSet{}
	var err error

	flwst := req.GetFlowSet()
	flwst_id_str := flwst.GetId().GetValue()

	alias := flwst.GetAlias()
	if alias != nil {
		flwst_s.Alias = &alias.Value
	}

	if flwst_s, err = self.storage.PatchFlowSet(ctx, flwst_id_str, flwst_s); err != nil {
		self.logger.WithError(err).Errorf("failed to patch flow set in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.PatchFlowSetResponse{
		FlowSet: copy_flow_set(flwst_s),
	}

	self.logger.WithField("id", flwst_id_str).Infof("patch flow set")

	return res, nil
}
