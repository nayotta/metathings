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

func (self *MetathingsIdentitydService) DeleteCredential(ctx context.Context, req *pb.DeleteCredentialRequest) (*empty.Empty, error) {
	var tkn_s *storage.Token
	var tkns_s []*storage.Token
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
	cred_id_str := cred.GetId().GetValue()

	// ensure no token issued by credential
	tkn_s = &storage.Token{
		CredentialId: &cred_id_str,
	}
	if tkns_s, err = self.storage.ListTokens(tkn_s); err != nil {
		self.logger.WithError(err).Errorf("failed to list tokens in storage")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	if len(tkns_s) > 0 {
		err = errors.New("token issued by credential")
		self.logger.WithError(err).Warningf("token issued by credential")
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}

	if err = self.storage.DeleteCredential(cred_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to delete token in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"credential_id": cred_id_str,
	}).Infof("delete credential")

	return &empty.Empty{}, nil
}
