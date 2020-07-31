package metathings_identityd2_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
	if tkn_s, err = self.validate_token(ctx, tkn); err != nil {
		return nil, err
	}

	self.logger.WithFields(log.Fields{
		"token":  tkn.GetText().GetValue()[:8],
		"entity": *tkn_s.EntityId,
		"domain": *tkn_s.DomainId,
	}).Debugf("check token")

	return &empty.Empty{}, nil
}
