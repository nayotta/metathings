package metathings_identityd2_service

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	var t *storage.Token
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	tkn_txt := req.GetToken().GetText()
	if tkn_txt == nil || tkn_txt.GetValue() == "" {
		err = errors.New("token.text is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	tkn_txt_str := tkn_txt.GetValue()

	if t, err = self.storage.GetTokenByText(tkn_txt_str); err != nil {
		self.logger.WithError(err).Warningf("failed to get token by text in storage")
		return nil, status.Errorf(codes.Unauthenticated, ErrUnauthenticated.Error())
	}

	res := &pb.ValidateTokenResponse{
		Token: copy_token(t),
	}

	self.logger.WithField("token_text", tkn_txt_str).Infof("validate token")

	return res, nil
}
