package metathings_identityd2_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) ListRoles(ctx context.Context, req *pb.ListRolesRequest) (*pb.ListRolesResponse, error) {
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	rol := &storage.Role{}

	if id := req.GetId(); id != nil {
		rol.Id = &id.Value
	}
	if domain := req.GetDomain(); domain != nil {
		if domainID := domain.GetId(); domainID != nil {
			rol.DomainId = &domainID.Value
		}
	}
	if name := req.GetName(); name != nil {
		rol.Name = &name.Value
	}
	if alias := req.GetAlias(); alias != nil {
		rol.Alias = &alias.Value
	}

	rols, err := self.storage.ListRoles(rol)
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
