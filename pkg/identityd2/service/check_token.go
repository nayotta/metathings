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
	pb "github.com/nayotta/metathings/proto/identityd2"
)

func (self *MetathingsIdentitydService) CheckToken(ctx context.Context, req *pb.CheckTokenRequest) (*empty.Empty, error) {
	var dom_id_str string
	var tkn_s *storage.Token
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	tkn := req.GetToken()
	if dom_id_str = tkn.GetDomain().GetId().GetValue(); dom_id_str == "" {
		err = errors.New("token.domain.id is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	if tkn_s, err = self.validate_token(ctx, tkn); err != nil {
		return nil, err
	}

	if *tkn_s.Domain.Id != dom_id_str {
		err = policy.ErrUnauthenticated
		self.logger.WithError(err).Warningf("failed to match request domain id and token domain id")
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"token":  tkn.GetText().GetValue()[:8],
		"entity": *tkn_s.EntityId,
		"domain": *tkn_s.DomainId,
	}).Debugf("check token")

	return &empty.Empty{}, nil
}
