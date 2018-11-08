package metathings_deviced_connection

import (
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

type Connection interface {
	Err() error
	Wait() chan bool
}

type ConnectionCenter interface {
	BuildConnection(*storage.Device, pb.DevicedService_ConnectServer) (Connection, error)
	UnaryCall(*storage.Device, *pb.OpUnaryCallValue) (*pb.UnaryCallValue, error)
}

type cc struct {
}

func (self *cc) BuildConnection(dev *storage.Device, stm pb.DevicedService_ConnectServer) (Connection, error) {
	panic("unimplemented")
}

func (self *cc) UnaryCall(*storage.Device, *pb.OpUnaryCallValue) (*pb.UnaryCallValue, error) {
	panic("unimplemented")
}

func NewConnectionCenter() (ConnectionCenter, error) {
	return &cc{}, nil
}
