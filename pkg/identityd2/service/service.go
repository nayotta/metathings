package metathings_identityd2_service

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	policy "github.com/nayotta/metathings/pkg/identityd2/policy"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MetathingsIdentitydServiceOption struct {
	TokenExpire time.Duration
}

type MetathingsIdentitydService struct {
	grpc_helper.AuthorizationTokenParser

	opt      *MetathingsIdentitydServiceOption
	logger   log.FieldLogger
	storage  storage.Storage
	enforcer policy.Enforcer
}

var (
	ignore_methods = []string{
		"IssueTokenByToken",
		"IssueTokenByPassword",
		"IssueTokenByCredential",
	}
)

func (self *MetathingsIdentitydService) enforce(ctx context.Context, obj, act string) error {
	var err error

	tkn := ctx.Value("token").(*storage.Token)

	var groups []string
	for _, g := range tkn.Groups {
		groups = append(groups, *g.Id)
	}

	if err = self.enforcer.Enforce(*tkn.DomainId, groups, *tkn.EntityId, obj, act); err != nil {
		if err == ErrPermissionDenied {
			self.logger.WithFields(log.Fields{
				"subject": *tkn.EntityId,
				"domain":  *tkn.DomainId,
				"groups":  groups,
				"object":  obj,
				"action":  act,
			}).Warningf("denied to do #action")
			return status.Errorf(codes.PermissionDenied, ErrPermissionDenied.Error())
		} else {
			self.logger.WithError(err).Errorf("failed to enforce")
			return status.Errorf(codes.Internal, err.Error())
		}
	}
	return nil
}

func (self *MetathingsIdentitydService) validate_chain(providers []interface{}, invokers []interface{}) error {
	default_invokers := []interface{}{policy_helper.ValidateValidator}
	invokers = append(default_invokers, invokers...)
	if err := policy_helper.ValidateChain(
		providers,
		invokers,
	); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return status.Errorf(codes.InvalidArgument, err.Error())
	}

	return nil
}

func (self *MetathingsIdentitydService) is_ignore_method(md *grpc_helper.MethodDescription) bool {
	for _, m := range ignore_methods {
		if md.Method == m {
			return true
		}
	}
	return false
}

func (self *MetathingsIdentitydService) add_token_to_kind_in_enforcer(tkn_id string) error {
	var err error

	if err = self.enforcer.AddObjectToKind(tkn_id, KIND_TOKEN); err != nil {
		self.logger.WithError(err).Errorf("failed to add token to kind in enforcer")
		return status.Errorf(codes.Internal, err.Error())
	}

	return nil
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

	if tkn_txt, err = self.GetTokenFromContext(ctx); err != nil {
		self.logger.WithError(err).Warningf("failed to get token from context")
		return ctx, err
	}

	if tkn, err = self.storage.GetTokenByText(tkn_txt); err != nil {
		self.logger.WithError(err).Warningf("failed to get token in storage")
		return ctx, err
	}

	new_ctx = context.WithValue(ctx, "token", tkn)

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

func (self *MetathingsIdentitydService) ListEntities(context.Context, *pb.ListEntitiesRequest) (*pb.ListEntitiesResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) ShowEntity(context.Context, *empty.Empty) (*pb.ShowEntityResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) ListGroupsForEntity(context.Context, *pb.ListGroupsForEntityRequest) (*pb.ListGroupsForEntityResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) ShowGroups(context.Context, *empty.Empty) (*pb.ShowGroupsResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) ListCredentialsForEntity(context.Context, *pb.ListCredentialsForEntityRequest) (*pb.ListCredentialsForEntityResponse, error) {
	panic("unimplemented")
}

func NewMetathingsIdentitydService(
	enforcor policy.Enforcer,
	opt *MetathingsIdentitydServiceOption,
	logger log.FieldLogger,
	storage storage.Storage,
) (*MetathingsIdentitydService, error) {
	return &MetathingsIdentitydService{
		opt:      opt,
		logger:   logger,
		storage:  storage,
		enforcer: enforcor,
	}, nil
}
