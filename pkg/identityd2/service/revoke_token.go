package metathings_identityd2_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) ValidateRevokeToken(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, token_getter) {
				req := in.(*pb.RevokeTokenRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_token_text,
		},
	)
}

func (self *MetathingsIdentitydService) AuthorizeRevokeToken(ctx context.Context, in interface{}) error {
	var tkn_s *storage.Token
	var err error

	tkn_txt_str := in.(*pb.RevokeTokenRequest).GetToken().GetText().GetValue()
	if tkn_s, err = self.storage.GetTokenByText(tkn_txt_str); err != nil {
		self.logger.WithError(err).Warningf("faield to find token by text in storage")
		return err
	}

	return self.authorize(ctx, *tkn_s.Id, "revoke_token")
}

func (self *MetathingsIdentitydService) RevokeToken(ctx context.Context, req *pb.RevokeTokenRequest) (*empty.Empty, error) {
	var tkn_s *storage.Token
	var err error

	tkn_txt_str := req.GetToken().GetText().GetValue()
	if tkn_s, err = self.storage.GetTokenByText(tkn_txt_str); err != nil {
		self.logger.WithError(err).Warningf("faield to find token by text in storage")
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}

	if err = self.revoke_token(*tkn_s.Id); err != nil {
		self.logger.WithError(err).Errorf("failed to revoke token")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"token_id":  *tkn_s.Id,
		"entity_id": *tkn_s.EntityId,
		"domain_id": *tkn_s.DomainId,
	}).Infof("revoke token")

	return &empty.Empty{}, nil
}
