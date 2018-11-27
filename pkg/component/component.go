package metathings_component

type Component interface {
	Name() string
	RunModule(args ...interface{}) error
}

type NewComponent func() (Component, error)
