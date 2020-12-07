package metathings_identityd2_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/nayotta/metathings/proto/identityd2"
)

func (self *MetathingsIdentitydService) AuthorizeAddActionToRole(ctx context.Context, in interface{}) error {
	return self.authorize(ctx, in.(*pb.AddActionToRoleRequest).GetRole().GetId().GetValue(), "identityd2:add_action_to_role")
}

func (self *MetathingsIdentitydService) AddActionToRole(ctx context.Context, req *pb.AddActionToRoleRequest) (*empty.Empty, error) {
	var err error

	act := req.GetAction()
	rol := req.GetRole()

	act_id := act.GetId().GetValue()
	rol_id := rol.GetId().GetValue()

	err = self.storage.AddActionToRole(ctx, rol_id, act_id)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to add action to role")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"action": act_id,
		"role":   rol_id,
	}).Infof("add action to role")

	return &empty.Empty{}, nil
}
