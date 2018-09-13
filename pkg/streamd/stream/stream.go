package stream_manager

type StreamState int32

const (
	STREAM_STATE_UNKNOWN StreamState = iota
	STREAM_STATE_RUNNING
	STREAM_STATE_STARTING
	STREAM_STATE_TERMINATING
	STREAM_STATE_STOP
	STREAM_STATE_OVERFLOW
)

type Stream interface {
	Emitter
	Id() string
	Start() error
	Stop() error
	State() StreamState
	Sources() []Source
	Groups() []Group
	Close()
}

type StreamFactory interface {
	Set(key string, val interface{}) StreamFactory
	New() (Stream, error)
}

var new_stream_factorys = map[string]func() StreamFactory{}

func RegisterStreamFactory(name string, fn func() StreamFactory) {
	if _, ok := new_stream_factorys[name]; !ok {
		new_stream_factorys[name] = fn
	}
}

func NewStreamFactory(name string) (StreamFactory, error) {
	new_fn, ok := new_stream_factorys[name]
	if !ok {
		return nil, ErrUnregisteredStreamFactory
	}

	return new_fn(), nil
}
