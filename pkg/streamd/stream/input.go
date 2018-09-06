package stream_manager

type InputState int32

const (
	INPUT_STATE_UNKNOWN InputState = iota
	INPUT_STATE_RUNNING
	INPUT_STATE_STARTING
	INPUT_STATE_TERMINATING
	INPUT_STATE_STOP
	INPUT_STATE_OVERFLOW
)

type Input interface {
	Emitter
	Id() string
	Symbol() string
	Start() error
	Stop() error
	State() InputState
	Close()
}

type InputOption func(interface{})

type InputFactory func(opts ...InputOption) (Input, error)

var input_factorys = make(map[string]InputFactory)

func RegisterInput(name string, fty InputFactory) {
	if _, ok := input_factorys[name]; !ok {
		input_factorys[name] = fty
	}
}
