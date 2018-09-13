package stream_manager

type Source interface {
	Id() string
	Upstream() Upstream
}

type SourceFactory interface {
	Set(key string, val interface{}) SourceFactory
	New() (Source, error)
}

var new_source_factories = make(map[string]func() SourceFactory)

func RegisterSourceFactory(name string, fn func() SourceFactory) {
	if _, ok := new_source_factories[name]; !ok {
		new_source_factories[name] = fn
	}
}

func NewSourceFactory(name string) (SourceFactory, error) {
	new_fn, ok := new_source_factories[name]
	if !ok {
		return nil, ErrUnregisteredSourceFactory
	}
	return new_fn(), nil
}

func NewDefaultSourceFactory() SourceFactory {
	fty, _ := NewSourceFactory("default")
	return fty
}
