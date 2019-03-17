package metathings_identityd2_authorizer

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	identityd_pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

type Authorizer interface {
	Authorize(ctx context.Context, obj, act string) error
}

type authorizer struct {
	logger  log.FieldLogger
	cli_fty *client_helper.ClientFactory
}

func (a *authorizer) Authorize(ctx context.Context, object, action string) error {
	cli, cfn, err := a.cli_fty.NewIdentityd2ServiceClient()
	if err != nil {
		a.logger.WithError(err).Errorf("failed to connect to identityd2")
		return status.Errorf(codes.Internal, err.Error())
	}
	defer cfn()

	req := &identityd_pb.AuthorizeTokenRequest{}
	_, err = cli.AuthorizeToken(ctx, req)
	if err != nil {
		a.logger.WithError(err).Warningf("permission denied")
		return status.Errorf(codes.PermissionDenied, err.Error())
	}

	return nil
}

func NewAuthorizer(cli_fty *client_helper.ClientFactory, logger log.FieldLogger) Authorizer {
	return &authorizer{
		logger:  logger,
		cli_fty: cli_fty,
	}
}
