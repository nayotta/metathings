package metathings_identityd2_service

import (
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) GetCredential(ctx context.Context, req *pb.GetCredentialRequest) (*pb.GetCredentialResponse, error) {
	var cred *storage.Credential
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	credId := req.GetId()
	if credId == nil || credId.GetValue == "" {
		err = errors.New("credential.id is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	credIdStr := credId.GetValue

	if cred, err = self.storage.GetCredential(credIdStr); err != nil {
		self.logger.WithError(err).Errorf("failed to get credential in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.GetCredentialResponse{
		Credential: copy_credential(cred)
	}

	self.logger.WithField("id", credIdStr).Infof("get credential")

	return res, nil
}
