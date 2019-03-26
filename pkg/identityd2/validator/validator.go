package metathings_identityd2_validator

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	log "github.com/sirupsen/logrus"
)

type Providers []interface{}
type Invokers []interface{}

type Validator interface {
	Validate(ps Providers, is Invokers) error
}

type validator struct {
	logger   log.FieldLogger
	defaults Invokers
}

func (v *validator) Validate(ps Providers, is Invokers) error {
	is = append(v.defaults, is...)
	if err := policy_helper.ValidateChain(ps, is); err != nil {
		v.logger.WithError(err).Warningf("failed to validate request data")
		return status.Errorf(codes.InvalidArgument, err.Error())
	}

	return nil
}

func NewValidator(defaults Invokers, logger log.FieldLogger) Validator {
	return &validator{
		logger:   logger,
		defaults: defaults,
	}
}
