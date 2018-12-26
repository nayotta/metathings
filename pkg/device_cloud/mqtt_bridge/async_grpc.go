package metathingsdevicecloudmqttbridge

import (
	pb "github.com/nayotta/metathings/pkg/proto/device_cloud"
	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

// AsyncDeviceCloudGrpc AsyncDeviceCloudGrpc
type AsyncDeviceCloudGrpc interface {
	Recv() chan *pb.StreamCallRequest
}

// AsyncDeviceCloudGrpcImpl AsyncDeviceCloudGrpcImpl
type AsyncDeviceCloudGrpcImpl struct {
	Stream pb.DeviceCloudService_StreamCallServer
}

// Recv Recv
func (that *AsyncDeviceCloudGrpcImpl) Recv() chan *pb.StreamCallRequest {
	ch := make(chan *pb.StreamCallRequest)
	go func() {
		for {
			buf, err := that.Stream.Recv()
			if err != nil {
				ch <- nil
				return
			}
			ch <- buf
		}
	}()
	return ch
}

// NewAsyncDeviceCloudGrpc NewAsyncDeviceCloudGrpc
func NewAsyncDeviceCloudGrpc(stream pb.DeviceCloudService_StreamCallServer) AsyncDeviceCloudGrpc {
	return &AsyncDeviceCloudGrpcImpl{
		Stream: stream,
	}
}

// AsyncDevicedGrpc AsyncDevicedGrpc
type AsyncDevicedGrpc interface {
	Recv() chan *deviced_pb.StreamCallRequest
}

// AsyncDevicedGrpcImpl AsyncDevicedGrpcImpl
type AsyncDevicedGrpcImpl struct {
	Stream deviced_pb.DevicedService_StreamCallServer
}

// Recv Recv
func (that *AsyncDevicedGrpcImpl) Recv() chan *deviced_pb.StreamCallRequest {
	ch := make(chan *deviced_pb.StreamCallRequest)
	go func() {
		for {
			buf, err := that.Stream.Recv()
			if err != nil {
				ch <- nil
				return
			}
			ch <- buf
		}
	}()
	return ch
}

// NewAsyncDevicedGrpc NewAsyncDevicedGrpc
func NewAsyncDevicedGrpc(stream deviced_pb.DevicedService_StreamCallServer) AsyncDevicedGrpc {
	return &AsyncDevicedGrpcImpl{
		Stream: stream,
	}
}
