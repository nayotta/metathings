package metathings_identityd2_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) ValidatePatchCredential(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, credential_getter) {
				req := in.(*pb.PatchCredentialRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_credential_id,
		},
	)
}

func (self *MetathingsIdentitydService) AuthorizePatchCredential(ctx context.Context, in interface{}) error {
	return self.authorize(ctx, in.(*pb.PatchCredentialRequest).GetCredential().GetId().GetValue(), "patch_credential")
}

func (self *MetathingsIdentitydService) PatchCredential(ctx context.Context, req *pb.PatchCredentialRequest) (*pb.PatchCredentialResponse, error) {
	var err error

	cred_req := req.GetCredential()
	cred := &storage.Credential{}

	idStr := cred_req.GetId().GetValue()

	if cred_req.GetAlias() != nil {
		cred.Alias = &cred_req.Alias.Value
	}
	if cred_req.GetDescription() != nil {
		cred.Description = &cred_req.Description.Value
	}

	if cred, err = self.storage.PatchCredential(idStr, cred); err != nil {
		self.logger.WithError(err).Errorf("failed to patch credential in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.PatchCredentialResponse{
		Credential: copy_credential(cred),
	}

	self.logger.WithField("id", idStr).Infof("patch credential")

	return res, nil
}
