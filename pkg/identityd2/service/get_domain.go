package metathings_identityd2_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/identityd2"
)

func (self *MetathingsIdentitydService) ValidateGetDomain(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() domain_getter {
				req := in.(*pb.GetDomainRequest)
				return req
			},
		},
		identityd_validator.Invokers{
			ensure_get_domain_id,
		},
	)
}

func (self *MetathingsIdentitydService) AuthorizeGetDomain(ctx context.Context, in interface{}) error {
	return self.authorize(ctx, in.(*pb.GetDomainRequest).GetDomain().GetId().GetValue(), "identityd2:get_domain")
}

func (self *MetathingsIdentitydService) GetDomain(ctx context.Context, req *pb.GetDomainRequest) (*pb.GetDomainResponse, error) {
	var dom_s *storage.Domain
	var err error

	id := req.GetDomain().GetId().GetValue()
	dom_s, err = self.storage.GetDomain(ctx, id)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to get domain in storage")
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	res := &pb.GetDomainResponse{
		Domain: copy_domain(dom_s),
	}

	self.logger.WithField("id", id).Debugf("get domain")

	return res, nil
}
