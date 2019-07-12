package metathings_device_cloud_service

import (
	"errors"
	"sync"

	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

var (
	ErrUnsupportedPushFrameToFlowChannelDriver = errors.New("unsupported push frame to flow channel driver")
)

type PushFrameToFlowChannel interface {
	Channel() <-chan *pb.OpFrame
	Close() error
}

type PushFrameToFlowChannelFactory interface {
	New(args ...interface{}) (PushFrameToFlowChannel, error)
}

var push_frame_to_flow_channel_factories map[string]PushFrameToFlowChannelFactory
var push_frame_to_flow_channel_factories_once sync.Once

func register_push_frame_to_flow_channel_factory(name string, fty PushFrameToFlowChannelFactory) {
	push_frame_to_flow_channel_factories_once.Do(func() {
		push_frame_to_flow_channel_factories = make(map[string]PushFrameToFlowChannelFactory)
	})
	push_frame_to_flow_channel_factories[name] = fty
}

func NewPushFrameToFlowChannel(name string, args ...interface{}) (PushFrameToFlowChannel, error) {
	fty, ok := push_frame_to_flow_channel_factories[name]
	if !ok {
		return nil, ErrUnsupportedPushFrameToFlowChannelDriver
	}

	return fty.New(args...)
}
