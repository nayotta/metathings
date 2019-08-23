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

func (self *MetathingsIdentitydService) ValidateDeleteCredential(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, credential_getter) {
				req := in.(*pb.DeleteCredentialRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_credential_id,
		},
	)
}

func (self *MetathingsIdentitydService) AuthorizeDeleteCredential(ctx context.Context, in interface{}) error {
	return self.authorize(ctx, in.(*pb.DeleteCredentialRequest).GetCredential().GetId().GetValue(), "identityd2:delete_credential")
}

func (self *MetathingsIdentitydService) DeleteCredential(ctx context.Context, req *pb.DeleteCredentialRequest) (*empty.Empty, error) {
	var tkn_s *storage.Token
	var tkns_s []*storage.Token
	var err error

	cred_id_str := req.GetCredential().GetId().GetValue()

	// ensure no token issued by credential
	tkn_s = &storage.Token{
		CredentialId: &cred_id_str,
	}
	if tkns_s, err = self.storage.ListTokens(tkn_s); err != nil {
		self.logger.WithError(err).Errorf("failed to list tokens in storage")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	if len(tkns_s) > 0 {
		for _, tkn_s := range tkns_s {
			if err = self.storage.DeleteToken(*tkn_s.Id); err != nil {
				self.logger.WithError(err).Errorf("failed to delete token in storage")
				return nil, status.Errorf(codes.Internal, err.Error())
			}
		}
		self.logger.WithField("credential_id", cred_id_str).Debugf("delete tokens by credential id")
	}

	if err = self.storage.DeleteCredential(cred_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to delete token in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if err = self.webhook.Trigger(map[string]interface{}{
		"action":     "delete_credential",
		"credential": &pb.Credential{Id: cred_id_str},
	}); err != nil {
		self.logger.WithError(err).Errorf("failed to trigger delete credential webhook")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"credential_id": cred_id_str,
	}).Infof("delete credential")

	return &empty.Empty{}, nil
}
