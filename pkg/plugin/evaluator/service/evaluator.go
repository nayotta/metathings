package main

type Evaluator interface {
	Eval(Data) error
}
