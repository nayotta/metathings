package metathings_deviced_flow

import (
	"sync"
	"time"

	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

type FlowFilter struct {
	BeginAt time.Time
	EndAt   time.Time
}

type FlowOption struct {
	FlowId   string
	DeviceId string
}

type Flow interface {
	Id() string
	Device() string
	PushFrame(*pb.Frame) error
<<<<<<< HEAD
	PullFrame() (<-chan *pb.Frame, chan struct{})
=======
	PullFrame() <-chan *pb.Frame
>>>>>>> v1.1.21.1
	QueryFrame(...*FlowFilter) ([]*pb.Frame, error)
	Err() error
	Close() error
}

type FlowFactory interface {
	New(*FlowOption) (Flow, error)
}

var flow_factories map[string]func(...interface{}) (FlowFactory, error)
var register_flow_factory_once sync.Once

func register_flow_factory(name string, fty func(...interface{}) (FlowFactory, error)) {
	register_flow_factory_once.Do(func() {
		flow_factories = make(map[string]func(...interface{}) (FlowFactory, error))
	})

	flow_factories[name] = fty
}

func NewFlowFactory(name string, args ...interface{}) (FlowFactory, error) {
	fty, ok := flow_factories[name]
	if !ok {
		return nil, ErrUnknownFlowFactory
	}

	return fty(args...)
}
