package metathings_identityd2_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) ValidateGetAction(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, action_getter) {
				req := in.(*pb.GetActionRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{ensure_get_action_id},
	)
}

func (self *MetathingsIdentitydService) AuthorizeGetAction(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.GetActionRequest).GetAction().GetId().GetValue(), "get_action")
}

func (self *MetathingsIdentitydService) GetAction(ctx context.Context, req *pb.GetActionRequest) (*pb.GetActionResponse, error) {
	var err error
	var act_s *storage.Action

	act := req.GetAction()
	id_str := act.GetId().GetValue()

	if act_s, err = self.storage.GetAction(id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to get action in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.GetActionResponse{
		Action: copy_action(act_s),
	}

	self.logger.WithField("id", id_str).Debugf("get action")

	return res, nil
}
