package metathings_identityd2_service

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*pb.CreateGroupResponse, error) {
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	id_str := id_helper.NewId()
	if req.GetId() != nil && req.GetId().GetValue() != "" {
		id_str = req.GetId().GetValue()
	}

	dom := req.GetDomain()
	if dom == nil || dom.GetId() == nil || dom.GetId().GetValue() == "" {
		err = errors.New("domain.id is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	dom_id_str := dom.GetId().GetValue()

	desc_str := ""
	if req.GetDescription() != nil {
		desc_str = req.GetDescription().GetValue()
	}

	extra_str := must_parse_extra(req.Extra)

	if err = self.enforcer.AddGroup(dom_id_str, id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to add group in enforcer")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	grp := &storage.Group{
		Id:          &id_str,
		DomainId:    &dom_id_str,
		Name:        &req.Name.Value,
		Alias:       &req.Alias.Value,
		Description: &desc_str,
		Extra:       &extra_str,
	}

	if grp, err = self.storage.CreateGroup(grp); err != nil {
		self.logger.WithError(err).Errorf("failed to create group in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.CreateGroupResponse{
		Group: copy_group(grp),
	}

	self.logger.WithField("id", id_str).Infof("create group")

	return res, nil
}
