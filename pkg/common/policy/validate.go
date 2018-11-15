package policy_helper

import (
	"context"

	"go.uber.org/fx"
)

type Validator interface {
	Validate() error
}

func ValidateChain(providers []interface{}, invokers []interface{}) error {
	var err error

	app := fx.New(
		fx.Provide(providers...),
		fx.Invoke(invokers...),
	)

	if err = app.Start(context.Background()); err != nil {
		return err
	}

	return nil
}

func ValidateValidator(v Validator) error {
	return v.Validate()
}
