package metathings_deviced_flow

import (
	"time"

	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

type FlowFilter struct {
	BeginAt time.Time
	EndAt   time.Time
}

type Flow interface {
	Id() string
	Device() string
	PushFrame(*pb.Frame) error
	PullFrame() <-chan *pb.Frame
	QueryFrame(...*FlowFilter) ([]*pb.Frame, error)
	Close() error
}

func NewFlow(args ...interface{}) (Flow, error) {
	panic("unimplemented")
}
