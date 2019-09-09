package metathings_identityd2_service

import (
	"context"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (self *MetathingsIdentitydService) ValidatePatchDomain(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, domain_getter) {
				req := in.(*pb.PatchDomainRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_domain_id,
		},
	)
}

func (self *MetathingsIdentitydService) AuthorizePatchDomain(ctx context.Context, in interface{}) error {
	return self.authorize(ctx, in.(*pb.PatchDomainRequest).GetDomain().GetId().GetValue(), "identityd2:patch_domain")
}

func (self *MetathingsIdentitydService) PatchDomain(ctx context.Context, req *pb.PatchDomainRequest) (*pb.PatchDomainResponse, error) {
	var err error

	dom_req := req.GetDomain()
	dom := &storage.Domain{}

	idStr := dom_req.GetId().GetValue()
	if dom_req.GetAlias() != nil {
		dom.Alias = &dom_req.Alias.Value
	}
	if dom_req.GetExtra() != nil {
		extraStr := must_parse_extra(dom_req.GetExtra())
		dom.Extra = &extraStr
	}

	if dom, err = self.storage.PatchDomain(idStr, dom); err != nil {
		self.logger.WithError(err).Errorf("failed to patch domain in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.PatchDomainResponse{
		Domain: copy_domain(dom),
	}

	self.logger.WithField("id", idStr).Infof("patch domain")

	return res, nil
}
