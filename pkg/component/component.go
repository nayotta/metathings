package metathings_component

type Component interface {
	Name() string
	NewModule(args ...interface{}) (Module, error)
}

type NewComponent func(args ...interface{}) (Component, error)
