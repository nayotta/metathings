package metathings_identityd2_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) ValidateAddObjectToGroup(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, object_getter, group_getter) {
				req := in.(*pb.AddObjectToGroupRequest)
				return req, req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_object_id,
			ensure_get_group_id,
		},
	)
}

func (self *MetathingsIdentitydService) AuthorizeAddObjectToGroup(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.AddObjectToGroupRequest).GetObject().GetId().GetValue(), "add_object_to_group")
}

func (self *MetathingsIdentitydService) AddObjectToGroup(ctx context.Context, req *pb.AddObjectToGroupRequest) (*empty.Empty, error) {
	var err error

	grp_id_str := req.GetGroup().GetId().GetValue()
	obj_id_str := req.GetObject().GetId().GetValue()

	if err = self.storage.AddObjectToGroup(grp_id_str, obj_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to add object to group in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"object_id": obj_id_str,
		"group_id":  grp_id_str,
	}).Infof("add object to group")

	return &empty.Empty{}, nil
}
