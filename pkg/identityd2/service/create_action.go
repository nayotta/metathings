package metathings_identityd2_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/proto/identityd2"
)

func (self *MetathingsIdentitydService) CreateAction(ctx context.Context, req *pb.CreateActionRequest) (*pb.CreateActionResponse, error) {
	var err error
	var act_s *storage.Action

	act := req.GetAction()

	id_str := id_helper.NewId()
	if act.GetId() != nil {
		id_str = act.GetId().GetValue()
	}

	desc_str := ""
	if act.GetDescription() != nil {
		desc_str = act.GetDescription().GetValue()
	}
	name_str := act.GetName().GetValue()
	alias_str := name_str
	if act.GetAlias() != nil {
		alias_str = act.GetAlias().GetValue()
	}

	act_s = &storage.Action{
		Id:          &id_str,
		Name:        &name_str,
		Alias:       &alias_str,
		Description: &desc_str,
	}

	if extra := act.GetExtra(); extra != nil {
		act_s.ExtraHelper = pb_helper.ExtractStringMapToString(extra)
	}

	act_s, err = self.storage.CreateAction(ctx, act_s)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to create action in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.CreateActionResponse{
		Action: copy_action(act_s),
	}

	self.logger.WithField("id", id_str).Infof("create action")

	return res, nil
}
