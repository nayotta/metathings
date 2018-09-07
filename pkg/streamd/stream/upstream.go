package stream_manager

import (
	"time"
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
	StreamData
}

func (self *UpstreamMetadata) SensorId() string {
	return self.AsString("sensor_id")
}

func (self *UpstreamMetadata) SensorName() string {
	return self.AsString("sensor_name")
}

func (self *UpstreamMetadata) CreatedAt() time.Time {
	return self.AsTime("created_at")
}

func (self *UpstreamMetadata) ArrviedAt() time.Time {
	return self.AsTime("arrvied_at")
}

func (self *UpstreamMetadata) Data() StreamData {
	return self.StreamData
}

type UpstreamData struct {
	StreamData
	metadata *UpstreamMetadata
}

func (self *UpstreamData) Metadata() *UpstreamMetadata {
	return self.metadata
}

func (self *UpstreamData) Data() StreamData {
	return self.StreamData
}

func NewUpstreamData(data map[string]interface{}, metadata map[string]interface{}) *UpstreamData {
	usmd := &UpstreamMetadata{NewStreamData(metadata)}
	usd := &UpstreamData{
		StreamData: NewStreamData(data),
		metadata:   usmd,
	}

	return usd
}
