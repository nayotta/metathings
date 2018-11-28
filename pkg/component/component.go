package metathings_component

const METATHINGS_COMPONENT_PREFIX = "mtc"

type Component interface {
	Name() string
	NewModule(args []string) (Module, error)
}

type NewComponent func() (Component, error)
