// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package deviced

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DevicedService service

type DevicedServiceClient interface {
	CreateDevice(ctx context.Context, in *CreateDeviceRequest, opts ...grpc.CallOption) (*CreateDeviceResponse, error)
	DeleteDevice(ctx context.Context, in *DeleteDeviceRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	PatchDevice(ctx context.Context, in *PatchDeviceRequest, opts ...grpc.CallOption) (*PatchDeviceResponse, error)
	GetDevice(ctx context.Context, in *GetDeviceRequest, opts ...grpc.CallOption) (*GetDeviceResponse, error)
	ListDevices(ctx context.Context, in *ListDevicesRequest, opts ...grpc.CallOption) (*ListDevicesResponse, error)
	// internal device only
	ShowDevice(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ShowDeviceResponse, error)
	Connect(ctx context.Context, opts ...grpc.CallOption) (DevicedService_ConnectClient, error)
	Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UnaryCall(ctx context.Context, in *UnaryCallRequest, opts ...grpc.CallOption) (*UnaryCallResponse, error)
	StreamCall(ctx context.Context, opts ...grpc.CallOption) (DevicedService_StreamCallClient, error)
}

type devicedServiceClient struct {
	cc *grpc.ClientConn
}

func NewDevicedServiceClient(cc *grpc.ClientConn) DevicedServiceClient {
	return &devicedServiceClient{cc}
}

func (c *devicedServiceClient) CreateDevice(ctx context.Context, in *CreateDeviceRequest, opts ...grpc.CallOption) (*CreateDeviceResponse, error) {
	out := new(CreateDeviceResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.deviced.DevicedService/CreateDevice", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicedServiceClient) DeleteDevice(ctx context.Context, in *DeleteDeviceRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/ai.metathings.service.deviced.DevicedService/DeleteDevice", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicedServiceClient) PatchDevice(ctx context.Context, in *PatchDeviceRequest, opts ...grpc.CallOption) (*PatchDeviceResponse, error) {
	out := new(PatchDeviceResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.deviced.DevicedService/PatchDevice", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicedServiceClient) GetDevice(ctx context.Context, in *GetDeviceRequest, opts ...grpc.CallOption) (*GetDeviceResponse, error) {
	out := new(GetDeviceResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.deviced.DevicedService/GetDevice", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicedServiceClient) ListDevices(ctx context.Context, in *ListDevicesRequest, opts ...grpc.CallOption) (*ListDevicesResponse, error) {
	out := new(ListDevicesResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.deviced.DevicedService/ListDevices", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicedServiceClient) ShowDevice(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ShowDeviceResponse, error) {
	out := new(ShowDeviceResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.deviced.DevicedService/ShowDevice", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicedServiceClient) Connect(ctx context.Context, opts ...grpc.CallOption) (DevicedService_ConnectClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DevicedService_serviceDesc.Streams[0], c.cc, "/ai.metathings.service.deviced.DevicedService/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &devicedServiceConnectClient{stream}
	return x, nil
}

type DevicedService_ConnectClient interface {
	Send(*ConnectResponse) error
	Recv() (*ConnectRequest, error)
	grpc.ClientStream
}

type devicedServiceConnectClient struct {
	grpc.ClientStream
}

func (x *devicedServiceConnectClient) Send(m *ConnectResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *devicedServiceConnectClient) Recv() (*ConnectRequest, error) {
	m := new(ConnectRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *devicedServiceClient) Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/ai.metathings.service.deviced.DevicedService/Heartbeat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicedServiceClient) UnaryCall(ctx context.Context, in *UnaryCallRequest, opts ...grpc.CallOption) (*UnaryCallResponse, error) {
	out := new(UnaryCallResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.deviced.DevicedService/UnaryCall", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicedServiceClient) StreamCall(ctx context.Context, opts ...grpc.CallOption) (DevicedService_StreamCallClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DevicedService_serviceDesc.Streams[1], c.cc, "/ai.metathings.service.deviced.DevicedService/StreamCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &devicedServiceStreamCallClient{stream}
	return x, nil
}

type DevicedService_StreamCallClient interface {
	Send(*StreamCallRequest) error
	Recv() (*StreamCallResponse, error)
	grpc.ClientStream
}

type devicedServiceStreamCallClient struct {
	grpc.ClientStream
}

func (x *devicedServiceStreamCallClient) Send(m *StreamCallRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *devicedServiceStreamCallClient) Recv() (*StreamCallResponse, error) {
	m := new(StreamCallResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for DevicedService service

type DevicedServiceServer interface {
	CreateDevice(context.Context, *CreateDeviceRequest) (*CreateDeviceResponse, error)
	DeleteDevice(context.Context, *DeleteDeviceRequest) (*empty.Empty, error)
	PatchDevice(context.Context, *PatchDeviceRequest) (*PatchDeviceResponse, error)
	GetDevice(context.Context, *GetDeviceRequest) (*GetDeviceResponse, error)
	ListDevices(context.Context, *ListDevicesRequest) (*ListDevicesResponse, error)
	// internal device only
	ShowDevice(context.Context, *empty.Empty) (*ShowDeviceResponse, error)
	Connect(DevicedService_ConnectServer) error
	Heartbeat(context.Context, *HeartbeatRequest) (*empty.Empty, error)
	UnaryCall(context.Context, *UnaryCallRequest) (*UnaryCallResponse, error)
	StreamCall(DevicedService_StreamCallServer) error
}

func RegisterDevicedServiceServer(s *grpc.Server, srv DevicedServiceServer) {
	s.RegisterService(&_DevicedService_serviceDesc, srv)
}

func _DevicedService_CreateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicedServiceServer).CreateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.deviced.DevicedService/CreateDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicedServiceServer).CreateDevice(ctx, req.(*CreateDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicedService_DeleteDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicedServiceServer).DeleteDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.deviced.DevicedService/DeleteDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicedServiceServer).DeleteDevice(ctx, req.(*DeleteDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicedService_PatchDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicedServiceServer).PatchDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.deviced.DevicedService/PatchDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicedServiceServer).PatchDevice(ctx, req.(*PatchDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicedService_GetDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicedServiceServer).GetDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.deviced.DevicedService/GetDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicedServiceServer).GetDevice(ctx, req.(*GetDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicedService_ListDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicedServiceServer).ListDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.deviced.DevicedService/ListDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicedServiceServer).ListDevices(ctx, req.(*ListDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicedService_ShowDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicedServiceServer).ShowDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.deviced.DevicedService/ShowDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicedServiceServer).ShowDevice(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicedService_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DevicedServiceServer).Connect(&devicedServiceConnectServer{stream})
}

type DevicedService_ConnectServer interface {
	Send(*ConnectRequest) error
	Recv() (*ConnectResponse, error)
	grpc.ServerStream
}

type devicedServiceConnectServer struct {
	grpc.ServerStream
}

func (x *devicedServiceConnectServer) Send(m *ConnectRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *devicedServiceConnectServer) Recv() (*ConnectResponse, error) {
	m := new(ConnectResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DevicedService_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicedServiceServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.deviced.DevicedService/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicedServiceServer).Heartbeat(ctx, req.(*HeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicedService_UnaryCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnaryCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicedServiceServer).UnaryCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.deviced.DevicedService/UnaryCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicedServiceServer).UnaryCall(ctx, req.(*UnaryCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicedService_StreamCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DevicedServiceServer).StreamCall(&devicedServiceStreamCallServer{stream})
}

type DevicedService_StreamCallServer interface {
	Send(*StreamCallResponse) error
	Recv() (*StreamCallRequest, error)
	grpc.ServerStream
}

type devicedServiceStreamCallServer struct {
	grpc.ServerStream
}

func (x *devicedServiceStreamCallServer) Send(m *StreamCallResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *devicedServiceStreamCallServer) Recv() (*StreamCallRequest, error) {
	m := new(StreamCallRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DevicedService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ai.metathings.service.deviced.DevicedService",
	HandlerType: (*DevicedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDevice",
			Handler:    _DevicedService_CreateDevice_Handler,
		},
		{
			MethodName: "DeleteDevice",
			Handler:    _DevicedService_DeleteDevice_Handler,
		},
		{
			MethodName: "PatchDevice",
			Handler:    _DevicedService_PatchDevice_Handler,
		},
		{
			MethodName: "GetDevice",
			Handler:    _DevicedService_GetDevice_Handler,
		},
		{
			MethodName: "ListDevices",
			Handler:    _DevicedService_ListDevices_Handler,
		},
		{
			MethodName: "ShowDevice",
			Handler:    _DevicedService_ShowDevice_Handler,
		},
		{
			MethodName: "Heartbeat",
			Handler:    _DevicedService_Heartbeat_Handler,
		},
		{
			MethodName: "UnaryCall",
			Handler:    _DevicedService_UnaryCall_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _DevicedService_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamCall",
			Handler:       _DevicedService_StreamCall_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_service_f88b37b534d6558a) }

var fileDescriptor_service_f88b37b534d6558a = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcd, 0x4e, 0xf2, 0x40,
	0x14, 0x85, 0xcd, 0xc7, 0xd7, 0x0b, 0x28, 0x8e, 0x89, 0x8b, 0x1a, 0x37, 0xac, 0xdc, 0x38, 0xfc,
	0xbd, 0x81, 0x60, 0x74, 0xe1, 0xc2, 0x48, 0x34, 0xd1, 0x98, 0x90, 0xa1, 0x5c, 0xdb, 0x26, 0xa5,
	0x53, 0x3b, 0x03, 0x84, 0x47, 0xf5, 0x6d, 0x4c, 0x3b, 0x9d, 0xa1, 0x90, 0xe8, 0x94, 0x15, 0xe1,
	0x70, 0x7e, 0x38, 0xf7, 0xde, 0x81, 0xb6, 0xc0, 0x74, 0x1d, 0x7a, 0x48, 0x93, 0x94, 0x4b, 0x4e,
	0xae, 0x58, 0x48, 0x97, 0x28, 0x99, 0x0c, 0xc2, 0xd8, 0x17, 0x54, 0xff, 0xb8, 0xc0, 0xec, 0x63,
	0xe1, 0x5e, 0xfa, 0x9c, 0xfb, 0x11, 0xf6, 0x72, 0xf2, 0x7c, 0xf5, 0xd9, 0xc3, 0x65, 0x22, 0xb7,
	0x4a, 0xeb, 0x9e, 0x7b, 0x29, 0x32, 0x89, 0x33, 0x45, 0xd6, 0xe0, 0x02, 0x23, 0x3c, 0x04, 0x49,
	0xc2, 0xa4, 0x17, 0xec, 0x63, 0x1d, 0x1f, 0xe5, 0x01, 0x2b, 0x0a, 0x85, 0x86, 0x44, 0x81, 0x9d,
	0x89, 0x80, 0x6f, 0xf6, 0x69, 0x6d, 0x8f, 0xc7, 0x31, 0x7a, 0xb2, 0xf8, 0x7a, 0x1a, 0x20, 0x4b,
	0xe5, 0x1c, 0x99, 0x06, 0x3a, 0xab, 0x98, 0xa5, 0xdb, 0x99, 0xc7, 0xa2, 0xc8, 0x98, 0xc8, 0x14,
	0xd9, 0xb2, 0x04, 0x0d, 0xbf, 0xff, 0xc3, 0xc9, 0x44, 0x95, 0x9c, 0xaa, 0xce, 0x64, 0x0b, 0xad,
	0x71, 0x5e, 0x48, 0xe1, 0x64, 0x48, 0xff, 0x9c, 0x0d, 0x2d, 0x93, 0x9f, 0xf1, 0x6b, 0x85, 0x42,
	0xba, 0xa3, 0xa3, 0x34, 0x22, 0xe1, 0xb1, 0xc0, 0x6e, 0x8d, 0x7c, 0x40, 0x6b, 0x92, 0x8f, 0xad,
	0x62, 0x74, 0x99, 0xac, 0xa3, 0x2f, 0xa8, 0xda, 0x15, 0xd5, 0xbb, 0xa2, 0x77, 0xd9, 0xae, 0xba,
	0x35, 0xb2, 0x86, 0xe6, 0x53, 0x36, 0xff, 0xc2, 0x7c, 0x60, 0x31, 0x2f, 0x71, 0xb5, 0xf7, 0xf0,
	0x18, 0x89, 0x69, 0x95, 0x80, 0x73, 0x8f, 0xb2, 0x48, 0xed, 0x59, 0x2c, 0x0c, 0x53, 0x67, 0xf6,
	0xab, 0x0b, 0x4c, 0xe2, 0x1a, 0x9a, 0x8f, 0xa1, 0x28, 0x70, 0x61, 0x6d, 0x5a, 0xe2, 0x56, 0x6d,
	0xba, 0x27, 0x31, 0xb9, 0x6f, 0x00, 0xd3, 0x80, 0x6f, 0x8a, 0xaa, 0xbf, 0x6c, 0xc2, 0xb5, 0xfd,
	0x9d, 0x9d, 0x45, 0xc9, 0x3a, 0x86, 0xc6, 0x58, 0xdd, 0x3b, 0xa1, 0xb6, 0xe3, 0x52, 0x3c, 0x2d,
	0x76, 0x6f, 0xaa, 0xf2, 0xf3, 0xea, 0xdd, 0xda, 0x75, 0xbd, 0x5f, 0x27, 0xaf, 0xe0, 0x3c, 0xe8,
	0x07, 0x65, 0x5d, 0x9a, 0x61, 0xda, 0x8f, 0x30, 0x01, 0xe7, 0x25, 0x7b, 0x97, 0x63, 0x16, 0x45,
	0x56, 0x5f, 0xc3, 0xac, 0x7a, 0x0c, 0x25, 0x81, 0x99, 0xdc, 0x06, 0x60, 0x9a, 0xbf, 0xfb, 0x3c,
	0xd2, 0xe6, 0xb0, 0xa3, 0xea, 0xcc, 0xc1, 0x11, 0x0a, 0x1d, 0x9a, 0x8d, 0xf0, 0xd6, 0x79, 0x6f,
	0x14, 0x9c, 0xf9, 0xbf, 0x7c, 0x0e, 0xa3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x41, 0xc7, 0xc8,
	0x2c, 0x76, 0x05, 0x00, 0x00,
}
