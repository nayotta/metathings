package metathings_identityd2_service

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) GetCredential(ctx context.Context, req *pb.GetCredentialRequest) (*pb.GetCredentialResponse, error) {
	var cred_s *storage.Credential
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	cred := req.GetCredential()
	if cred.GetId() == nil {
		err = errors.New("credential.id is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	id_str := cred.GetId().GetValue()

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
