package metathings_plugin_evaluator

type Evaluator interface {
	Eval(Data) error
}
