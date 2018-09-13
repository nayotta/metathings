package stream_manager

type StreamManager interface {
	NewStream(opt StreamOption, extra map[string]interface{}) (Stream, error)
	GetStream(id string) (Stream, error)
}

type StreamManagerFactory interface {
	Set(key string, val interface{}) StreamManagerFactory
	New() (StreamManager, error)
}

var new_stream_manager_factories = make(map[string]func() StreamManagerFactory)

func RegisterStreamManagerFactory(name string, fn func() StreamManagerFactory) {
	if _, ok := new_stream_manager_factories[name]; !ok {
		new_stream_manager_factories[name] = fn
	}
}

func NewStreamManagerFactory(name string) (StreamManagerFactory, error) {
	new_fn, ok := new_stream_manager_factories[name]
	if !ok {
		return nil, ErrUnregisteredStreamManagerFactory
	}
	return new_fn(), nil
}

func NewDefaultStreamManagerFactory() StreamManagerFactory {
	fty, _ := NewStreamManagerFactory("default")
	return fty
}
