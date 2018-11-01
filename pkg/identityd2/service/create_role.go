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

func (self *MetathingsIdentitydService) CreateRole(ctx context.Context, req *pb.CreateRoleRequest) (*pb.CreateRoleResponse, error) {
	var role *storage.Role
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	id := req.GetId().GetValue()
	if id == "" {
		id = id_helper.NewId()
	}

	dom_id := req.GetDomain().GetId().GetValue()
	if dom_id == "" {
		err = errors.New("domain.id is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	description := req.GetDescription().GetValue()
	extra_str := must_parse_extra(req.GetExtra())
	name_str := req.Name.Value
	alias_str := req.Alias.Value

	if err = self.enforcer.AddObjectToKind(id, KIND_ROLE); err != nil {
		self.logger.WithError(err).Errorf("failed to add object to kind in enforcer")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	role = &storage.Role{
		Id:          &id,
		DomainId:    &dom_id,
		Name:        &name_str,
		Alias:       &alias_str,
		Description: &description,
		Extra:       &extra_str,
	}

	role, err = self.storage.CreateRole(role)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to create role in storage")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	res := &pb.CreateRoleResponse{
		Role: copy_role(role),
	}

	self.logger.WithField("id", *role.Id).Infof("create role")

	return res, nil
}
