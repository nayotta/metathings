package stream_manager

import (
	log "github.com/sirupsen/logrus"
)

type groupImplOption struct {
	id      string
	inputs  []*InputOption
	outputs []*OutputOption

	sym_tbl SymbolTable
	logger  log.FieldLogger
	brokers []string
}

type groupImpl struct {
	logger  log.FieldLogger
	opt     *groupImplOption
	inputs  []Input
	outputs []Output
}

func (self *groupImpl) Id() string {
	return self.opt.id
}

func (self *groupImpl) Inputs() []Input {
	return self.inputs
}

func (self *groupImpl) Outputs() []Output {
	return self.outputs
}

type groupImplFactory struct {
	opt *groupImplOption
}

func (self *groupImplFactory) Set(key string, val interface{}) GroupFactory {
	switch key {
	case "logger":
		self.opt.logger = val.(log.FieldLogger)
	case "symbol_table":
		self.opt.sym_tbl = val.(SymbolTable)
	case "brokers":
		self.opt.brokers = val.([]string)
	case "option":
		opt := val.(*GroupOption)
		self.opt.id = opt.Id
		self.opt.inputs = opt.Inputs
		self.opt.outputs = opt.Outputs
	}

	return self
}

func (self *groupImplFactory) New() (Group, error) {
	inputs := []Input{}
	for _, in_opt := range self.opt.inputs {
		fty, err := NewInputFactory(in_opt.Name)
		if err != nil {
			return nil, err
		}

		input, err := fty.Set("logger", self.opt.logger).
			Set("symbol_table", self.opt.sym_tbl).
			Set("brokers", self.opt.brokers).
			Set("option", in_opt).
			New()
		if err != nil {
			return nil, err
		}

		inputs = append(inputs, input)
	}

	outputs := []Output{}
	for _, out_opt := range self.opt.outputs {
		fty, err := NewOutputFactory(out_opt.Name)
		if err != nil {
			return nil, err
		}

		output, err := fty.Set("logger", self.opt.logger).
			Set("symbol_table", self.opt.sym_tbl).
			Set("brokers", self.opt.brokers).
			Set("option", out_opt).
			New()
		if err != nil {
			return nil, err
		}

		outputs = append(outputs, output)
	}

	grp := &groupImpl{
		logger: self.opt.logger.WithFields(log.Fields{
			"id":         self.opt.id,
			"#component": "group:default",
		}),
		opt:     self.opt,
		inputs:  inputs,
		outputs: outputs,
	}

	return grp, nil
}

func init() {
	RegisterGroupFactory("default", func() GroupFactory { return &groupImplFactory{opt: &groupImplOption{}} })
}
