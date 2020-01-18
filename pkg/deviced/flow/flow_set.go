package metathings_deviced_flow

import (
	"sync"

	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

type FlowSetOption struct {
	FlowSetId string
}

type FlowSetFrame struct {
	Device *pb.Device `json:"device"`
	Frame  *pb.Frame  `json:"frame"`
}

type FlowSet interface {
	Id() string
	PushFrame(*FlowSetFrame) error
	PullFrame() (<-chan *FlowSetFrame, chan struct{})
	Close() error
}

type FlowSetFactory interface {
	New(*FlowSetOption) (FlowSet, error)
}

var flow_set_factories map[string]func(...interface{}) (FlowSetFactory, error)
var register_flow_set_factory_once sync.Once

func registry_flow_set_factory(name string, fty func(...interface{}) (FlowSetFactory, error)) {
	register_flow_set_factory_once.Do(func() {
		flow_set_factories = make(map[string]func(...interface{}) (FlowSetFactory, error))
	})

	flow_set_factories[name] = fty
}

func NewFlowSetFactory(name string, args ...interface{}) (FlowSetFactory, error) {
	fty, ok := flow_set_factories[name]
	if !ok {
		return nil, ErrUnknownFlowSetFactory
	}

	return fty(args...)
}
