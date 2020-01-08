package metathings_identityd2_service

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy "github.com/nayotta/metathings/pkg/identityd2/policy"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
	log "github.com/sirupsen/logrus"
)

func (self *MetathingsIdentitydService) IssueTokenByToken(ctx context.Context, req *pb.IssueTokenByTokenRequest) (*pb.IssueTokenByTokenResponse, error) {
	var tkn_s *storage.Token
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	tkn_req := req.GetToken()
	dom := tkn_req.GetDomain()
	if dom == nil || dom.GetId() == nil {
		err = errors.New("token.domain.id is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	dom_id_str := dom.GetId().GetValue()

	if tkn_req.GetText() == nil {
		err = errors.New("token.text is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	tkn_txt_str := tkn_req.GetText().GetValue()

	if tkn_s, err = self.storage.GetTokenByText(ctx, tkn_txt_str); err != nil {
		self.logger.WithError(err).Warningf("failed to find token by text in storage")
		return nil, status.Errorf(codes.Unauthenticated, policy.ErrUnauthenticated.Error())
	}

	if *tkn_s.DomainId != dom_id_str {
		err = policy.ErrUnauthenticated
		self.logger.WithError(err).Errorf("failed to match request domain id and token domain id")
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	tkn := new_token(tkn_s.DomainId, tkn_s.EntityId, tkn_s.CredentialId, NONEXPIRATION)
	if tkn_s, err = self.storage.CreateToken(ctx, tkn); err != nil {
		self.logger.WithError(err).Errorf("failed to create token in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.IssueTokenByTokenResponse{
		Token: copy_token(tkn),
	}

	self.logger.WithFields(log.Fields{
		"entity_id": *tkn_s.EntityId,
		"domain_id": dom_id_str,
	}).Infof("issue token by token")

	return res, nil
}
