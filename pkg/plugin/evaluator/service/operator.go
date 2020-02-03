package main

import "sync"

type Operator interface {
	Run(Data, Config) (Data, error)
	Close() error
}

type OperatorFactory func(args ...interface{}) (Operator, error)

var operator_factories map[string]OperatorFactory
var operator_factories_once sync.Once

func registry_operator_factory(name string, fty OperatorFactory) {
	operator_factories_once.Do(func() {
		operator_factories = make(map[string]OperatorFactory)
	})

	operator_factories[name] = fty
}
