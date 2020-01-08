package metathings_identityd2_service

import (
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy "github.com/nayotta/metathings/pkg/identityd2/policy"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) CheckToken(ctx context.Context, req *pb.CheckTokenRequest) (*empty.Empty, error) {
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

	if tkn.GetDomain() == nil || tkn.GetDomain().GetId() == nil {
		err = errors.New("token.domain.id is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	dom_id_str := tkn.GetDomain().GetId().GetValue()

	if tkn_s, err = self.storage.GetTokenByText(ctx, tkn_txt_str); err != nil {
		self.logger.WithError(err).Errorf("failed to find token by text in storage")
		return nil, status.Errorf(codes.Unauthenticated, policy.ErrUnauthenticated.Error())
	}

	if *tkn_s.Domain.Id != dom_id_str {
		err = policy.ErrUnauthenticated
		self.logger.WithError(err).Warningf("failed to match request domain id and token domain id")
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"token_id":  *tkn_s.Id,
		"domain_id": dom_id_str,
		"entity_id": *tkn_s.EntityId,
	}).Debugf("check token")

	return &empty.Empty{}, nil
}
