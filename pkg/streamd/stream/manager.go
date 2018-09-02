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
	Close()
}

type StreamOption func(interface{})

type StreamManager interface {
	NewStream(opts ...StreamOption) Stream
	GetStream(id string) Stream
}

type StreamManagerFactory func() (StreamManager, error)

var stream_manager_factorys = make(map[string]StreamManagerFactory)

func RegisterStreamManager(name string, fty StreamManagerFactory) {
	if _, ok := stream_manager_factorys[name]; !ok {
		stream_manager_factorys[name] = fty
	}
}
