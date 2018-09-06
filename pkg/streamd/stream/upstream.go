package stream_manager

import (
	"time"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

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

type UpstreamMetadata struct {
	opt_helper.Option
}

func (self *UpstreamMetadata) SensorId() string {
	return self.GetString("sensor_id")
}

func (self *UpstreamMetadata) SensorName() string {
	return self.GetString("sensor_name")
}

func (self *UpstreamMetadata) CreatedAt() time.Time {
	return *self.Get("created_at").(*time.Time)
}

func (self *UpstreamMetadata) ArrviedAt() time.Time {
	return *self.Get("arrvied_at").(*time.Time)
}

type UpstreamData struct {
	opt_helper.Option
	metadata *UpstreamMetadata
}

func (self *UpstreamData) Metadata() *UpstreamMetadata {
	return self.metadata
}

func NewUpstreamData(data map[string]interface{}, metadata map[string]interface{}) *UpstreamData {
	usmd := &UpstreamMetadata{opt_helper.NewOptionMap(metadata)}
	usd := &UpstreamData{
		Option:   opt_helper.NewOptionMap(data),
		metadata: usmd,
	}

	return usd
}
