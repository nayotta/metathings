package metathings_evaluatord_service

import (
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	log "github.com/sirupsen/logrus"

	afo_helper "github.com/nayotta/metathings/pkg/common/auth_func_overrider"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	storage "github.com/nayotta/metathings/pkg/evaluatord/storage"
	timer "github.com/nayotta/metathings/pkg/evaluatord/timer"
	identityd_authorizer "github.com/nayotta/metathings/pkg/identityd2/authorizer"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/evaluatord"
)

type MetathingsEvaluatordServiceOption struct{}

type MetathingsEvaluatordService struct {
	grpc_auth.ServiceAuthFuncOverride
	tknr          token_helper.Tokener
	cli_fty       *client_helper.ClientFactory
	opt           *MetathingsEvaluatordServiceOption
	logger        log.FieldLogger
	storage       storage.Storage
	task_storage  storage.TaskStorage
	timer_storage storage.TimerStorage
	timer_backend timer.TimerBackend
	authorizer    identityd_authorizer.Authorizer
	validator     identityd_validator.Validator
	tkvdr         token_helper.TokenValidator
}

func (srv *MetathingsEvaluatordService) get_logger() log.FieldLogger {
	return srv.logger
}

func (srv *MetathingsEvaluatordService) IsIgnoreMethod(md *grpc_helper.MethodDescription) bool {
	return false
}

func NewMetathingsEvaludatorService(
	opt *MetathingsEvaluatordServiceOption,
	logger log.FieldLogger,
	storage storage.Storage,
	task_storage storage.TaskStorage,
	authorizer identityd_authorizer.Authorizer,
	validator identityd_validator.Validator,
	tkvdr token_helper.TokenValidator,
	tknr token_helper.Tokener,
	cli_fty *client_helper.ClientFactory,
) (pb.EvaluatordServiceServer, error) {
	srv := &MetathingsEvaluatordService{
		opt:          opt,
		logger:       logger,
		storage:      storage,
		task_storage: task_storage,
		authorizer:   authorizer,
		validator:    validator,
		tkvdr:        tkvdr,
		tknr:         tknr,
		cli_fty:      cli_fty,
	}

	srv.ServiceAuthFuncOverride = afo_helper.NewAuthFuncOverrider(tkvdr, srv, logger)

	return srv, nil
}
