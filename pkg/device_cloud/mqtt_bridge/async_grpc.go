package metathingsdevicecloudmqttbridge

import (
	pb "github.com/nayotta/metathings/pkg/proto/device_cloud"
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
