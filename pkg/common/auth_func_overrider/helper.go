package auth_func_overrider

import (
	"context"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	log "github.com/sirupsen/logrus"

	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	identityd_pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

type IgnoreMethoder interface {
	IsIgnoreMethod(md *grpc_helper.MethodDescription) bool
}

type authFuncOverrider struct {
	logger log.FieldLogger
	tkvdr  token_helper.TokenValidator
	igmthr IgnoreMethoder
}

func (afo *authFuncOverrider) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	var tkn *identityd_pb.Token
	var tkn_txt string
	var new_ctx context.Context
	var err error
	var md *grpc_helper.MethodDescription

	if md, err = grpc_helper.ParseMethodDescription(fullMethodName); err != nil {
		afo.logger.WithError(err).Warningf("failed to parse method description")
		return ctx, err
	}

	if afo.igmthr.IsIgnoreMethod(md) {
		return ctx, nil
	}

	if tkn_txt, err = grpc_helper.GetTokenFromContext(ctx); err != nil {
		afo.logger.WithError(err).Warningf("failed to get token from context")
		return ctx, err
	}

	if tkn, err = afo.tkvdr.Validate(ctx, tkn_txt); err != nil {
		afo.logger.WithError(err).Warningf("failed to validate token in identity service")
		return ctx, err
	}

	new_ctx = context.WithValue(ctx, "token", tkn)

	afo.logger.WithFields(log.Fields{
		"method":    md.Method,
		"entity_id": tkn.Entity.Id,
		"domain_id": tkn.Domain.Id,
	}).Debugf("authorize token")

	return new_ctx, nil
}

func NewAuthFuncOverrider(
	tkvdr token_helper.TokenValidator,
	igmthr IgnoreMethoder,
	logger log.FieldLogger,
) grpc_auth.ServiceAuthFuncOverride {
	return &authFuncOverrider{
		logger: logger,
		tkvdr:  tkvdr,
		igmthr: igmthr,
	}
}
