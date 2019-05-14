package metathings_tagd_service

import (
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	log "github.com/sirupsen/logrus"

	afo_helper "github.com/nayotta/metathings/pkg/common/auth_func_overrider"
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	log_helper "github.com/nayotta/metathings/pkg/common/log"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	identityd_authorizer "github.com/nayotta/metathings/pkg/identityd2/authorizer"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/tagd"
	storage "github.com/nayotta/metathings/pkg/tagd/storage"
)

type MetathingsTagdService struct {
	grpc_auth.ServiceAuthFuncOverride
	*log_helper.GetLoggerer
	authorizer identityd_authorizer.Authorizer
	validator  identityd_validator.Validator
	logger     log.FieldLogger
	stor       storage.Storage
}

func (ts *MetathingsTagdService) IsIgnoreMethod(md *grpc_helper.MethodDescription) bool {
	return false
}

func NewMetathingsTagdService(
	logger log.FieldLogger,
	stor storage.Storage,
	authorizer identityd_authorizer.Authorizer,
	validator identityd_validator.Validator,
	tkvdr token_helper.TokenValidator,
) (pb.TagdServiceServer, error) {
	ts := &MetathingsTagdService{
		GetLoggerer: log_helper.NewGetLoggerer(logger),
		logger:      logger,
		stor:        stor,
		authorizer:  authorizer,
		validator:   validator,
	}
	ts.ServiceAuthFuncOverride = afo_helper.NewAuthFuncOverrider(tkvdr, ts, logger)

	return ts, nil
}
