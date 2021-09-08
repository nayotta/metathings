package cmd_contrib

import (
	log "github.com/sirupsen/logrus"

	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
)

func NewValidator(logger log.FieldLogger) identityd_validator.Validator {
	return identityd_validator.NewValidator(identityd_validator.Invokers{}, logger)
}
