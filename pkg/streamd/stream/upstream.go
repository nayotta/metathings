package stream_manager

type UpstreamState int32

const (
	UPSTREAM_STATE_UNKNOWN UpstreamState = iota
	UPSTREAM_STATE_RUNNING
	UPSTREAM_STATE_STARTING
	UPSTREAM_STATE_TERMINATING
	UPSTREAM_STATE_STOP
	UPSTREAM_STATE_OVERFLOW
)

type Upstream interface {
	Emitter
	Id() string
	Start() error
	Stop() error
	State() UpstreamState
	Close()
}

type UpstreamOption func(interface{})

type UpstreamFactory func(opts ...UpstreamOption) (Upstream, error)

var upstream_factorys = make(map[string]UpstreamFactory)

func RegisterUpstream(name string, fty UpstreamFactory) {
	if _, ok := upstream_factorys[name]; !ok {
		upstream_factorys[name] = fty
	}
}
