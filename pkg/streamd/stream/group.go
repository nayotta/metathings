package stream_manager

type Group interface {
	Id() string
	Inputs() []Input
	Outputs() []Output
}

type GroupFactory interface {
	Set(key string, val interface{}) GroupFactory
	New() (Group, error)
}

var new_group_factories = make(map[string]func() GroupFactory)

func RegisterGroupFactory(name string, fn func() GroupFactory) {
	if _, ok := new_group_factories[name]; !ok {
		new_group_factories[name] = fn
	}
}

func NewGroupFactory(name string) (GroupFactory, error) {
	new_fn, ok := new_group_factories[name]
	if !ok {
		return nil, ErrUnregisteredGroupFactory
	}
	return new_fn(), nil
}

func NewDefaultGroupFactory() GroupFactory {
	fty, _ := NewGroupFactory("default")
	return fty
}
