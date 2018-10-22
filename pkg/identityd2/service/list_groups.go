package metathings_identityd2_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) ListGroups(ctx context.Context, req *pb.ListGroupsRequest) (*pb.ListGroupsResponse, error) {
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	grp := &storage.Group{}

	if id := req.GetId(); id != nil {
		grp.Id = &id.Value
	}
	if domain := req.GetDomain(); domain != nil {
		if domainID := domain.GetId(); domainID != nil {
			grp.DomainId = &domainID.Id.Value
		}
	}
	if name := req.GetName(); name != nil {
		grp.Name = &name.Value
	}
	if alias := req.GetAlias(); alias != nil {
		grp.Alias = &alias.Value
	}

	if grps, err := self.storage.ListGroups(grp); err != nil {
		self.logger.WithError(err).Errorf("failed to list groups in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.GetGroupResponse{}

	for grp = range grps {
		res.Group = append(res.Group, copy_group(grp)),
	}

	return res, nil
}
