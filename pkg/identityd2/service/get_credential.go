package metathings_identityd2_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/identityd2"
)

func (self *MetathingsIdentitydService) ValidateGetCredential(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() credential_getter {
				req := in.(*pb.GetCredentialRequest)
				return req
			},
		},
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

	if cred_s, err = self.storage.GetCredential(ctx, id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to get credential in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.GetCredentialResponse{
		Credential: copy_credential(cred_s),
	}

	self.logger.WithField("id", id_str).Debugf("get credential")

	return res, nil
}
