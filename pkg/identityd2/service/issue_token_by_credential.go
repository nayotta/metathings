package metathings_identityd2_service

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	passwd_helper "github.com/nayotta/metathings/pkg/common/passwd"
	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	policy "github.com/nayotta/metathings/pkg/identityd2/policy"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) ValidateIssueTokenByCredential(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, credential_getter, *pb.IssueTokenByCredentialRequest) {
				req := in.(*pb.IssueTokenByCredentialRequest)
				return req, req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_credential_id,
			func(req *pb.IssueTokenByCredentialRequest) error {
				if req.GetTimestamp() == nil {
					return errors.New("timestamp is empty")
				}

				if req.GetNonce() == nil {
					return errors.New("nonce is empty")
				}

				if req.GetHmac() == nil {
					return errors.New("hmac is empty")
				}

				return nil
			},
		},
	)
}

func (self *MetathingsIdentitydService) IssueTokenByCredential(ctx context.Context, req *pb.IssueTokenByCredentialRequest) (*pb.IssueTokenByCredentialResponse, error) {
	var cred_s *storage.Credential
	var err error

	cred := req.GetCredential()
	dom_id_str := ""
	dom := cred.GetDomain()
	if dom != nil && dom.GetId() != nil {
		dom_id_str = dom.GetId().GetValue()
	}

	cred_id := cred.GetId()
	cred_id_str := cred_id.GetValue()

	timestamp := pb_helper.ToTime(*req.GetTimestamp())
	nonce := req.GetNonce().GetValue()
	hmac := req.GetHmac().GetValue()

	if cred_s, err = self.storage.GetCredential(cred_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to find credential by id in storage")
		return nil, status.Errorf(codes.Unauthenticated, policy.ErrUnauthenticated.Error())
	}

	if dom_id_str != "" && !domain_in_credential(cred_s, dom_id_str) {
		err = policy.ErrUnauthenticated
		self.logger.WithError(err).Errorf("failed to match request domain id and credential domain id")
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	if !passwd_helper.ValidateHmac(hmac, *cred_s.Secret, cred_id_str, timestamp, nonce) {
		err = policy.ErrUnauthenticated
		self.logger.WithError(err).Warningf("failed to validate secret")
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	tkn := new_token(&dom_id_str, cred_s.EntityId, cred_s.Id, self.opt.TokenExpire)
	if tkn, err = self.storage.CreateToken(tkn); err != nil {
		self.logger.WithError(err).Errorf("failed to create token in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.IssueTokenByCredentialResponse{
		Token: copy_token(tkn),
	}

	self.logger.WithFields(log.Fields{
		"credential_id": *cred_s.Id,
		"entity_id":     *cred_s.EntityId,
		"domain_id":     dom_id_str,
	}).Infof("issue token by credential")

	return res, nil
}
