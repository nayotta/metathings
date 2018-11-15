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

func (self *MetathingsIdentitydService) RevokeToken(ctx context.Context, req *pb.RevokeTokenRequest) (*empty.Empty, error) {
	var tkn_s *storage.Token
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	tkn := req.GetToken()
	if tkn.GetText() == nil {
		err = errors.New("token.text is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	tkn_txt_str := tkn.GetText().GetValue()

	if tkn_s, err = self.storage.GetTokenByText(tkn_txt_str); err != nil {
		self.logger.WithError(err).Warningf("faield to find token by text in storage")
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}

	if err = self.revoke_token(*tkn_s.Id); err != nil {
		self.logger.WithError(err).Errorf("failed to revoke token")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"token_id":  *tkn_s.Id,
		"entity_id": *tkn_s.EntityId,
		"domain_id": *tkn_s.DomainId,
	}).Infof("revoke token")

	return &empty.Empty{}, nil
}
