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

type StreamManagerOption func(interface{})
type NewStreamOption func(interface{})

type StreamManager interface {
	NewStream(opts ...NewStreamOption) (Stream, error)
	GetStream(id string) (Stream, error)
}

type StreamManagerFactory func(...StreamManagerOption) (StreamManager, error)

var stream_manager_factorys = make(map[string]StreamManagerFactory)

func RegisterStreamManager(name string, fty StreamManagerFactory) {
	if _, ok := stream_manager_factorys[name]; !ok {
		stream_manager_factorys[name] = fty
	}
}
