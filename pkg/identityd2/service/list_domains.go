package metathings_identityd2_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) ListDomains(ctx context.Context, req *pb.ListDomainsRequest) (*pb.ListDomainsResponse, error) {
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	dom := &storage.Domain{}

	if req.GetId() != nil && req.GetId().GetValue() != "" {
		idStr := req.GetId().GetValue()
		dom.Id = &idStr
	}
	if req.GetName() != nil && req.GetName().GetValue() != "" {
		nameStr := req.GetName().GetValue()
		dom.Name = &nameStr
	}
	if req.GetAlias() != nil && req.GetAlias().GetValue() != "" {
		aliasStr := req.GetAlias().GetValue()
		dom.Alias = &aliasStr
	}

	doms, err := self.storage.ListDomains(dom)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to list domains in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ListDomainsResponse{}

	for _, dom = range doms {
		res.Domains = append(res.Domains, copy_domain(dom))
	}

	return res, nil
}
