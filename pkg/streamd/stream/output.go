package stream_manager

type OutputState int32

const (
	OUTPUT_STATE_UNKNOWN OutputState = iota
	OUTPUT_STATE_RUNNING
	OUTPUT_STATE_STARTING
	OUTPUT_STATE_TERMINATING
	OUTPUT_STATE_STOP
	OUTPUT_STATE_OVERFLOW
)

type Output interface {
	Emitter
	Id() string
	Symbol() string
	Start() error
	Stop() error
	State() OutputState
	Close()
}

type OutputOption func(interface{})

type OutputFactory func(opts ...OutputOption) (Output, error)

var output_factorys = make(map[string]OutputFactory)

func RegisterOutput(name string, fty OutputFactory) {
	if _, ok := output_factorys[name]; !ok {
		output_factorys[name] = fty
	}
}
