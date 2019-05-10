package metathings_identityd2_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) ValidateCreateDomain(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, domain_getter) {
				req := in.(*pb.CreateDomainRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_domain_parent_id,
			ensure_get_domain_name,
		},
	)
}

func (self *MetathingsIdentitydService) CreateDomain(ctx context.Context, req *pb.CreateDomainRequest) (*pb.CreateDomainResponse, error) {
	var dom_s *storage.Domain
	var err error

	dom := req.GetDomain()

	id_str := dom.GetId().GetValue()
	if id_str == "" {
		id_str = id_helper.NewId()
	}
	parent_id_str := dom.GetParent().GetId().GetValue()
	extra_str := must_parse_extra(dom.GetExtra())
	name_str := dom.GetName().GetValue()
	alias_str := name_str
	if dom.GetAlias() != nil {
		alias_str = dom.GetAlias().GetValue()
	}

	dom_s = &storage.Domain{
		Id:       &id_str,
		Name:     &name_str,
		Alias:    &alias_str,
		ParentId: &parent_id_str,
		Extra:    &extra_str,
	}

	if dom_s, err = self.storage.CreateDomain(dom_s); err != nil {
		self.logger.WithError(err).Errorf("failed to create domain in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.CreateDomainResponse{
		Domain: copy_domain(dom_s),
	}

	self.logger.WithField("id", id_str).Infof("create domain")

	return res, nil
}
