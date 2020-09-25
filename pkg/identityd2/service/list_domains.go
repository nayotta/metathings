package metathings_identityd2_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/proto/identityd2"
)

func (self *MetathingsIdentitydService) ListDomains(ctx context.Context, req *pb.ListDomainsRequest) (*pb.ListDomainsResponse, error) {
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	dom_req := req.GetDomain()
	dom := &storage.Domain{}

	if dom_req.GetId() != nil && dom_req.GetId().GetValue() != "" {
		dom.Id = &dom_req.Id.Value
	}
	if dom_req.GetName() != nil && dom_req.GetName().GetValue() != "" {
		dom.Name = &dom_req.Name.Value
	}
	if dom_req.GetAlias() != nil && dom_req.GetAlias().GetValue() != "" {
		dom.Alias = &dom_req.Alias.Value
	}

	doms, err := self.storage.ListDomains(ctx, dom)
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
