package metathings_identityd2_authorizer

import (
	"context"

	"github.com/golang/protobuf/ptypes/wrappers"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
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

	new_ctx := context_helper.WithToken(ctx, context_helper.ExtractToken(ctx).GetText())
	req := &identityd_pb.AuthorizeTokenRequest{
		Object: &identityd_pb.OpEntity{Id: &wrappers.StringValue{Value: object}},
		Action: &identityd_pb.OpAction{Name: &wrappers.StringValue{Value: action}},
	}
	_, err = cli.AuthorizeToken(new_ctx, req)

	if err != nil {
		if gst, ok := status.FromError(err); ok {
			err = status.Errorf(gst.Code(), gst.Err().Error())
		} else {
			err = status.Errorf(codes.Internal, err.Error())
		}

		a.logger.WithError(err).Warningf("failed to authorize token")
		return err
	}

	return nil
}

func NewAuthorizer(cli_fty *client_helper.ClientFactory, logger log.FieldLogger) Authorizer {
	return &authorizer{
		logger:  logger,
		cli_fty: cli_fty,
	}
}
