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

func (self *MetathingsIdentitydService) validate_token(ctx context.Context, tkn *pb.OpToken) (*storage.Token, error) {
	var err error
	var tkn_s *storage.Token

	tkn_txt := tkn.GetText()
	if tkn_txt == nil || tkn_txt.GetValue() == "" {
		err = errors.New("token.text is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	tkn_txt_str := tkn_txt.GetValue()

	if tkn_s, err = self.storage.GetTokenByText(ctx, tkn_txt_str); err != nil {
		self.logger.WithError(err).Errorf("failed to get token by text in storage")
		return nil, status.Errorf(codes.Unauthenticated, policy.ErrUnauthenticated.Error())
	}

	if self.is_invalid_token(tkn_s) {
		if err = self.revoke_token(ctx, *tkn_s.Id); err != nil {
			self.logger.WithError(err).Warningf("failed to revoke token")
		}
		return nil, status.Errorf(codes.Unauthenticated, policy.ErrUnauthenticated.Error())
	}

	if self.is_refreshable_token(tkn_s) {
		if err = self.refresh_token(ctx, tkn_s); err != nil {
			self.logger.WithError(err).Warningf("failed to refresh token")
		}
	}

	return tkn_s, nil
}

func (self *MetathingsIdentitydService) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	var tkn_s *storage.Token
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	tkn := req.GetToken()
	if tkn_s, err = self.validate_token(ctx, tkn); err != nil {
		return nil, err
	}

	res := &pb.ValidateTokenResponse{
		Token: copy_token(tkn_s),
	}

	self.logger.WithFields(log.Fields{
		"token":  tkn.GetText().GetValue()[:8],
		"entity": *tkn_s.EntityId,
		"domain": *tkn_s.DomainId,
	}).Debugf("validate token")

	return res, nil
}
