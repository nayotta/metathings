package metathings_identityd2_service

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
	log "github.com/sirupsen/logrus"
)

func (self *MetathingsIdentitydService) IssueTokenByCredential(ctx context.Context, req *pb.IssueTokenByCredentialRequest) (*pb.IssueTokenByCredentialResponse, error) {
	var cred_s *storage.Credential
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	cred := req.GetCredential()

	dom := cred.GetDomain()
	if dom == nil || dom.GetId() == nil {
		err = errors.New("credential.domain.id is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	dom_id_str := dom.GetId().GetValue()

	cred_id := cred.GetId()
	if cred_id == nil {
		err = errors.New("credential.id is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
	}

	if cred_s, err = self.storage.GetCredential(cred_id.GetValue()); err != nil {
		self.logger.WithError(err).Errorf("failed to find credential by id in storage")
		return nil, status.Errorf(codes.Unauthenticated, ErrUnauthenticated.Error())
	}

	if *cred_s.Domain.Id != dom_id_str {
		err = ErrUnauthenticated
		self.logger.WithError(err).Errorf("failed to match request domain id and credential domain id")
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
