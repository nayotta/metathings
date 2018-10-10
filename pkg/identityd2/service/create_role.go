package metathings_identityd2_service

import (
	"context"
	"encoding/json"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) CreateRole(ctx context.Context, req *pb.CreateRoleRequest) (*pb.CreateRoleResponse, error) {
	var role *storage.Role
	var buf []byte
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	id := req.GetId().GetValue()
	if id == "" {
		id = id_helper.NewId()
	}

	dom_id := req.GetDomain().GetId().GetValue()
	if dom_id == "" {
		err = errors.New("domain.id is empty")
		self.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	description := req.GetDescription().GetValue()

	extra_map := pb_helper.ExtractStringMap(req.GetExtra())
	if buf, err = json.Marshal(extra_map); err != nil {
		err = errors.New("extra is bad argument")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	extra_str := string(buf)

	role = &storage.Role{
		Id:          &id,
		DomainId:    &dom_id,
		Name:        &req.Name.Value,
		Alias:       &req.Alias.Value,
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
