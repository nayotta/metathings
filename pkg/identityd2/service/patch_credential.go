package metathings_identityd2_service

import (
	"context"
	"errors"
	
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) PatchCredential(ctx context.Context, req *pb.PatchCredentialRequest) (*pb.PatchCredentialResponse, error) {
	var cred = &storage.Credential{}
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	if req.GetId() == nil || req.GetId().GetValue() == "" {
		err = errors.New("credential.id is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	idStr := req.GetId().GetValue()

	if req.GetAlias() != nil {
		aliasStr := req.GetAlias().GetValue()
		cred.Alias = &aliasStr
	}
	if req.GetDescription() != nil {
		descriptionStr := req.GetDescription().GetValue()
		cred.Description = &descriptionStr
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
