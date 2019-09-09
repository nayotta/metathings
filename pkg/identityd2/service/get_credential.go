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

func (self *MetathingsIdentitydService) ValidateGetCredential(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, credential_getter) {
				req := in.(*pb.GetCredentialRequest)
				return req, req
			}},
		identityd_validator.Invokers{
			ensure_get_credential_id,
		},
	)
}

func (self *MetathingsIdentitydService) AuthorizeGetCredential(ctx context.Context, in interface{}) error {
	return self.authorize(ctx, in.(*pb.GetCredentialRequest).GetCredential().GetId().GetValue(), "identityd2:get_credential")
}

func (self *MetathingsIdentitydService) GetCredential(ctx context.Context, req *pb.GetCredentialRequest) (*pb.GetCredentialResponse, error) {
	var cred_s *storage.Credential
	var err error

	id_str := req.GetCredential().GetId().GetValue()

	if cred_s, err = self.storage.GetCredential(id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to get credential in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.GetCredentialResponse{
		Credential: copy_credential(cred_s),
	}

	self.logger.WithField("id", id_str).Debugf("get credential")

	return res, nil
}
