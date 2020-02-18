package metathings_plugin_evaluator

import "context"

type Evaluator interface {
	Eval(context.Context, Data) error
}
