package metathings_identityd2_service

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"

	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	policy "github.com/nayotta/metathings/pkg/identityd2/policy"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

type MetathingsIdentitydServiceOption struct {
	TokenExpire time.Duration
}

type MetathingsIdentitydService struct {
	opt       *MetathingsIdentitydServiceOption
	logger    log.FieldLogger
	storage   storage.Storage
	validator validator.Validator
	backend   policy.Backend
}

var (
	ignore_methods = []string{
		"IssueTokenByToken",
		"IssueTokenByPassword",
		"IssueTokenByCredential",
	}
)

func (self *MetathingsIdentitydService) is_ignore_method(md *grpc_helper.MethodDescription) bool {
	for _, m := range ignore_methods {
		if md.Method == m {
			return true
		}
	}
	return false
}

func (self *MetathingsIdentitydService) revoke_token(tkn_id string) error {
	var err error

	if err = self.storage.DeleteToken(tkn_id); err != nil {
		self.logger.WithError(err).WithField("id", tkn_id).Warningf("failed to delete token in storage")
		return err
	}

	return nil
}

func (self *MetathingsIdentitydService) is_invalid_token(tkn *storage.Token) bool {
	now := time.Now()
	if tkn.ExpiresAt.Sub(now) < 0 {
		self.logger.WithFields(log.Fields{
			"token":      *tkn.Text,
			"expired_at": *tkn.ExpiresAt,
			"now":        now,
		}).Debugf("token expired")
		return true
	}

	return false
}

func (self *MetathingsIdentitydService) is_refreshable_token(tkn *storage.Token) bool {
	return tkn.ExpiresAt.Sub(time.Now()) < time.Duration(.25*float64(self.opt.TokenExpire))
}

func (self *MetathingsIdentitydService) refresh_token(tkn *storage.Token) error {
	tkn_id := *tkn.Id
	expires_at := time.Now().Add(self.opt.TokenExpire)

	return self.storage.RefreshToken(tkn_id, expires_at)
}

func (self *MetathingsIdentitydService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	var tkn *storage.Token
	var tkn_txt string
	var new_ctx context.Context
	var err error
	var md *grpc_helper.MethodDescription

	if md, err = grpc_helper.ParseMethodDescription(fullMethodName); err != nil {
		self.logger.WithError(err).Warningf("failed to parse method description")
		return ctx, err
	}
	if self.is_ignore_method(md) {
		return ctx, nil
	}

	if tkn_txt, err = grpc_helper.GetTokenFromContext(ctx); err != nil {
		self.logger.WithError(err).Warningf("failed to get token from context")
		return ctx, err
	}

	if tkn, err = self.storage.GetTokenByText(tkn_txt); err != nil {
		self.logger.WithError(err).Warningf("failed to get token in storage")
		return ctx, err
	}

	if self.is_invalid_token(tkn) {
		if err = self.revoke_token(*tkn.Id); err != nil {
			self.logger.WithError(err).Warningf("failed to revoke token")
		}
		return ctx, policy.ErrUnauthenticated
	}

	if self.is_refreshable_token(tkn) {
		if err = self.refresh_token(tkn); err != nil {
			self.logger.WithError(err).Warningf("failed to refresh token")
		}
	}

	new_ctx = context.WithValue(ctx, "token", copy_token(tkn))

	self.logger.WithFields(log.Fields{
		"method":    md.Method,
		"entity_id": *tkn.EntityId,
		"domain_id": *tkn.DomainId,
	}).Debugf("authorize token")

	return new_ctx, nil
}

func (self *MetathingsIdentitydService) ListRolesForEntity(context.Context, *pb.ListRolesForEntityRequest) (*pb.ListRolesForEntityResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) ShowEntity(context.Context, *empty.Empty) (*pb.ShowEntityResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) ShowGroups(context.Context, *empty.Empty) (*pb.ShowGroupsResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) ListCredentialsForEntity(context.Context, *pb.ListCredentialsForEntityRequest) (*pb.ListCredentialsForEntityResponse, error) {
	panic("unimplemented")
}

func NewMetathingsIdentitydService(
	opt *MetathingsIdentitydServiceOption,
	logger log.FieldLogger,
	storage storage.Storage,
	validator validator.Validator,
	backend policy.Backend,
) (pb.IdentitydServiceServer, error) {
	return &MetathingsIdentitydService{
		opt:       opt,
		logger:    logger,
		storage:   storage,
		validator: validator,
		backend:   backend,
	}, nil
}
