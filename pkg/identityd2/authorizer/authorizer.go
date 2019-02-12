package metathings_identityd2_authorizer

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy "github.com/nayotta/metathings/pkg/identityd2/policy"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
)

type Authorizer interface {
	Authorize(ctx context.Context, obj, act string) error
}

type authorizer struct {
	logger   log.FieldLogger
	enforcer policy.Enforcer
}

func (a *authorizer) Authorize(ctx context.Context, obj, act string) error {
	var err error

	tkn := ctx.Value("token").(*storage.Token)

	var groups []string
	for _, g := range tkn.Groups {
		groups = append(groups, *g.Id)
	}

	if err = a.enforcer.Enforce(*tkn.DomainId, groups, *tkn.EntityId, obj, act); err != nil {
		if err == policy.ErrPermissionDenied {
			a.logger.WithFields(log.Fields{
				"subject": *tkn.EntityId,
				"domain":  *tkn.DomainId,
				"groups":  groups,
				"object":  obj,
				"action":  act,
			}).Warningf("denied to do #action")
			return status.Errorf(codes.PermissionDenied, err.Error())
		} else {
			a.logger.WithError(err).Errorf("failed to enforce")
			return status.Errorf(codes.Internal, err.Error())
		}
	}
	return nil
}

func NewAuthorizer(enforcer policy.Enforcer, logger log.FieldLogger) Authorizer {
	return &authorizer{
		logger:   logger,
		enforcer: enforcer,
	}
}
