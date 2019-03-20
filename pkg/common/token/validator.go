package token_helper

import (
	"context"

	"github.com/golang/protobuf/ptypes/wrappers"
	log "github.com/sirupsen/logrus"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	identityd2_pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

type TokenValidator interface {
	Validate(token string) (*identityd2_pb.Token, error)
}

type identityd2TokenValidator struct {
	tknr    Tokener
	cli_fty *client_helper.ClientFactory
	logger  log.FieldLogger
}

func (self *identityd2TokenValidator) Validate(token string) (*identityd2_pb.Token, error) {
	ctx := context_helper.WithToken(context.Background(), self.tknr.GetToken())

	cli, cfn, err := self.cli_fty.NewIdentityd2ServiceClient()
	if err != nil {
		return nil, err
	}
	defer cfn()

	req := &identityd2_pb.ValidateTokenRequest{
		Token: &identityd2_pb.OpToken{
			Text: &wrappers.StringValue{Value: token},
		},
	}
	res, err := cli.ValidateToken(ctx, req)
	if err != nil {
		return nil, err
	}

	return res.Token, nil
}

func NewTokenValidator(
	tknr Tokener,
	cli_fty *client_helper.ClientFactory,
	logger log.FieldLogger) TokenValidator {
	return &identityd2TokenValidator{
		tknr:    tknr,
		cli_fty: cli_fty,
		logger:  logger,
	}
}
