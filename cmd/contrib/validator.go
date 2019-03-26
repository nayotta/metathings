package cmd_contrib

import (
	log "github.com/sirupsen/logrus"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
)

func NewValidator(logger log.FieldLogger) identityd_validator.Validator {
	return identityd_validator.NewValidator(identityd_validator.Invokers{policy_helper.ValidateValidator}, logger)
}
