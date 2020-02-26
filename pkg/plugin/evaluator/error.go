package metathings_plugin_evaluator

import "errors"

var (
	ErrUnsupportedOperatorName  = errors.New("unsupported operator name")
	ErrUnexpectedOperatorResult = errors.New("unexpected operator result")
)
