package metathings_identityd2_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/identityd2"
)

func (self *MetathingsIdentitydService) ValidatePatchAction(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() action_getter {
				req := in.(*pb.PatchActionRequest)
				return req
			},
		},
		identityd_validator.Invokers{ensure_get_action_id},
	)
}

func (self *MetathingsIdentitydService) AuthorizePatchAction(ctx context.Context, in interface{}) error {
	return self.authorize(ctx, in.(*pb.PatchActionRequest).GetAction().GetId().GetValue(), "identityd2:patch_action")
}

func (self *MetathingsIdentitydService) PatchAction(ctx context.Context, req *pb.PatchActionRequest) (*pb.PatchActionResponse, error) {
	var err error

	act_req := req.GetAction()
	act := &storage.Action{}

	idStr := act_req.GetId().GetValue()

	if act_req.GetAlias() != nil {
		act.Alias = &act_req.Alias.Value
	}
	if act_req.GetDescription() != nil {
		act.Description = &act_req.Description.Value
	}

	if extra := act_req.GetExtra(); extra != nil {
		act.ExtraHelper = pb_helper.ExtractStringMapToString(extra)
	}

	if act, err = self.storage.PatchAction(ctx, idStr, act); err != nil {
		self.logger.WithError(err).Errorf("failed to patch action in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.PatchActionResponse{
		Action: copy_action(act),
	}

	self.logger.WithField("id", idStr).Infof("patch action")

	return res, nil
}
