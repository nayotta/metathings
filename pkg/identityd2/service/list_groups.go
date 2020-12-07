package metathings_identityd2_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/proto/identityd2"
)

func (self *MetathingsIdentitydService) ListGroups(ctx context.Context, req *pb.ListGroupsRequest) (*pb.ListGroupsResponse, error) {
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	grp_req := req.GetGroup()
	grp := &storage.Group{}

	if id := grp_req.GetId(); id != nil {
		grp.Id = &id.Value
	}
	if domain := grp_req.GetDomain(); domain != nil {
		if domainID := domain.GetId(); domainID != nil {
			grp.DomainId = &domainID.Value
		}
	}
	if name := grp_req.GetName(); name != nil {
		grp.Name = &name.Value
	}
	if alias := grp_req.GetAlias(); alias != nil {
		grp.Alias = &alias.Value
	}

	grps, err := self.storage.ListGroups(ctx, grp)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to list groups in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ListGroupsResponse{}

	for _, grp = range grps {
		res.Groups = append(res.Groups, copy_group(grp))
	}

	self.logger.Debugf("list groups")

	return res, nil
}
