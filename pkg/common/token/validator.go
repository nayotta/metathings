package token_helper

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"

	app_cred_mgr "github.com/nayotta/metathings/pkg/common/application_credential_manager"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	identityd_pb "github.com/nayotta/metathings/pkg/proto/identityd"
)

type TokenValidator interface {
	Validate(token string) (*identityd_pb.Token, error)
}

type tokenValidator struct {
	app_cred_mgr app_cred_mgr.ApplicationCredentialManager
	cli_fty      *client_helper.ClientFactory
	logger       log.FieldLogger
}

func (vdr *tokenValidator) Validate(token string) (*identityd_pb.Token, error) {
	ctx := context.Background()
	md := metadata.Pairs(
		"authorization-subject", "mt "+token,
		"authorization", vdr.app_cred_mgr.GetToken(),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)

	cli, closeFn, err := vdr.cli_fty.NewIdentityServiceClient()
	if err != nil {
		return nil, err
	}
	defer closeFn()

	res, err := cli.ValidateToken(ctx, &empty.Empty{})
	if err != nil {
		return nil, err
	}

	return res.Token, nil
}

func NewTokenValidator(
	app_cred_mgr app_cred_mgr.ApplicationCredentialManager,
	cli_fty *client_helper.ClientFactory,
	logger log.FieldLogger) TokenValidator {
	return &tokenValidator{
		app_cred_mgr: app_cred_mgr,
		cli_fty:      cli_fty,
		logger:       logger,
	}
}
