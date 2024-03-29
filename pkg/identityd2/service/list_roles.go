package metathings_identityd2_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/proto/identityd2"
)

func (self *MetathingsIdentitydService) ListRoles(ctx context.Context, req *pb.ListRolesRequest) (*pb.ListRolesResponse, error) {
	var err error

	role_req := req.GetRole()
	rol := &storage.Role{}

	if id := role_req.GetId(); id != nil {
		rol.Id = &id.Value
	}
	if name := role_req.GetName(); name != nil {
		rol.Name = &name.Value
	}
	if alias := role_req.GetAlias(); alias != nil {
		rol.Alias = &alias.Value
	}

	rols, err := self.storage.ListRoles(ctx, rol)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to list roles in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ListRolesResponse{}

	for _, rol = range rols {
		res.Roles = append(res.Roles, copy_role(rol))
	}

	self.logger.Debugf("list roles")

	return res, nil
}
