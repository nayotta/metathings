package metathings_tagd_service

import (
	log "github.com/sirupsen/logrus"

	log_helper "github.com/nayotta/metathings/pkg/common/log"
	identityd_authorizer "github.com/nayotta/metathings/pkg/identityd2/authorizer"
	identityd_policy "github.com/nayotta/metathings/pkg/identityd2/policy"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/tagd"
	storage "github.com/nayotta/metathings/pkg/tagd/storage"
)

type MetathingsTagdService struct {
	*log_helper.GetLoggerer
	enforcer   identityd_policy.Enforcer
	authorizer identityd_authorizer.Authorizer
	validator  identityd_validator.Validator
	logger     log.FieldLogger
	stor       storage.Storage
}

func NewMetathingsTagdService(
	logger log.FieldLogger,
	stor storage.Storage,
	enforcer identityd_policy.Enforcer,
	authorizer identityd_authorizer.Authorizer,
	validator identityd_validator.Validator,
) (pb.TagdServiceServer, error) {
	ts := &MetathingsTagdService{
		GetLoggerer: log_helper.NewGetLoggerer(logger),
		logger:      logger,
		stor:        stor,
		enforcer:    enforcer,
		authorizer:  authorizer,
		validator:   validator,
	}
	return ts, nil
}
