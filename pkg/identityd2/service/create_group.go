package metathings_identityd2_service

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) ValidateCreateGroup(ctx context.Context, in interface{}) error {
	return self.validate_chain([]interface{}{
		func() (policy_helper.Validator, group_getter) {
			req := in.(*pb.CreateGroupRequest)
			return req, req
		},
	}, []interface{}{
		func(x group_getter) error {
			grp := x.GetGroup()

			if grp.GetDomain() == nil || grp.GetDomain().GetId() == nil {
				return errors.New("group.domain.id is empty")
			}

			if grp.GetName() == nil {
				return errors.New("group.name is empty")
			}

			return nil
		},
	})
}

func (self *MetathingsIdentitydService) CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*pb.CreateGroupResponse, error) {
	var err error

	grp := req.GetGroup()
	id_str := id_helper.NewId()
	if grp.GetId() != nil {
		id_str = grp.GetId().GetValue()
	}

	dom_id_str := grp.GetDomain().GetId().GetValue()
	desc_str := ""
	if grp.GetDescription() != nil {
		desc_str = grp.GetDescription().GetValue()
	}
	extra_str := must_parse_extra(grp.GetExtra())
	name_str := grp.GetName().GetValue()
	alias_str := name_str
	if grp.GetAlias() != nil {
		alias_str = grp.GetAlias().GetValue()
	}

	if err = self.enforcer.AddGroup(dom_id_str, id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to add group in enforcer")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if err = self.enforcer.AddObjectToKind(id_str, KIND_GROUP); err != nil {
		self.logger.WithError(err).Errorf("failed to add group to kind in enforcer")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	grp_s := &storage.Group{
		Id:          &id_str,
		DomainId:    &dom_id_str,
		Name:        &name_str,
		Alias:       &alias_str,
		Description: &desc_str,
		Extra:       &extra_str,
	}

	if grp_s, err = self.storage.CreateGroup(grp_s); err != nil {
		self.logger.WithError(err).Errorf("failed to create group in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.CreateGroupResponse{
		Group: copy_group(grp_s),
	}

	self.logger.WithField("id", id_str).Infof("create group")

	return res, nil
}
