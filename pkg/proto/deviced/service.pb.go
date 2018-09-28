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
	Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Stream(ctx context.Context, opts ...grpc.CallOption) (DevicedService_StreamClient, error)
	ListDevicesForUser(ctx context.Context, in *ListDevicesForUserRequest, opts ...grpc.CallOption) (*ListDevicesForUserResponse, error)
	UnaryCall(ctx context.Context, in *UnaryCallRequest, opts ...grpc.CallOption) (*UnaryCallResponse, error)
	StreamCall(ctx context.Context, opts ...grpc.CallOption) (DevicedService_StreamCallClient, error)
	GrantDeviceToGroup(ctx context.Context, in *GrantDeviceToGroupRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	RevokeDeviceFromGroup(ctx context.Context, in *RevokeDeviceFromGroupRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GrantDeviceToUser(ctx context.Context, in *GrantDeviceToUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	RevokeDeviceFromUser(ctx context.Context, in *RevokeDeviceFromUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
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

func (c *devicedServiceClient) Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/ai.metathings.service.deviced.DevicedService/Heartbeat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicedServiceClient) Stream(ctx context.Context, opts ...grpc.CallOption) (DevicedService_StreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DevicedService_serviceDesc.Streams[0], c.cc, "/ai.metathings.service.deviced.DevicedService/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &devicedServiceStreamClient{stream}
	return x, nil
}

type DevicedService_StreamClient interface {
	Send(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ClientStream
}

type devicedServiceStreamClient struct {
	grpc.ClientStream
}

func (x *devicedServiceStreamClient) Send(m *StreamResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *devicedServiceStreamClient) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *devicedServiceClient) ListDevicesForUser(ctx context.Context, in *ListDevicesForUserRequest, opts ...grpc.CallOption) (*ListDevicesForUserResponse, error) {
	out := new(ListDevicesForUserResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.deviced.DevicedService/ListDevicesForUser", in, out, c.cc, opts...)
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

func (c *devicedServiceClient) GrantDeviceToGroup(ctx context.Context, in *GrantDeviceToGroupRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/ai.metathings.service.deviced.DevicedService/GrantDeviceToGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicedServiceClient) RevokeDeviceFromGroup(ctx context.Context, in *RevokeDeviceFromGroupRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/ai.metathings.service.deviced.DevicedService/RevokeDeviceFromGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicedServiceClient) GrantDeviceToUser(ctx context.Context, in *GrantDeviceToUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/ai.metathings.service.deviced.DevicedService/GrantDeviceToUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicedServiceClient) RevokeDeviceFromUser(ctx context.Context, in *RevokeDeviceFromUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/ai.metathings.service.deviced.DevicedService/RevokeDeviceFromUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
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
	Heartbeat(context.Context, *HeartbeatRequest) (*empty.Empty, error)
	Stream(DevicedService_StreamServer) error
	ListDevicesForUser(context.Context, *ListDevicesForUserRequest) (*ListDevicesForUserResponse, error)
	UnaryCall(context.Context, *UnaryCallRequest) (*UnaryCallResponse, error)
	StreamCall(DevicedService_StreamCallServer) error
	GrantDeviceToGroup(context.Context, *GrantDeviceToGroupRequest) (*empty.Empty, error)
	RevokeDeviceFromGroup(context.Context, *RevokeDeviceFromGroupRequest) (*empty.Empty, error)
	GrantDeviceToUser(context.Context, *GrantDeviceToUserRequest) (*empty.Empty, error)
	RevokeDeviceFromUser(context.Context, *RevokeDeviceFromUserRequest) (*empty.Empty, error)
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

func _DevicedService_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DevicedServiceServer).Stream(&devicedServiceStreamServer{stream})
}

type DevicedService_StreamServer interface {
	Send(*StreamRequest) error
	Recv() (*StreamResponse, error)
	grpc.ServerStream
}

type devicedServiceStreamServer struct {
	grpc.ServerStream
}

func (x *devicedServiceStreamServer) Send(m *StreamRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *devicedServiceStreamServer) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DevicedService_ListDevicesForUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDevicesForUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicedServiceServer).ListDevicesForUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.deviced.DevicedService/ListDevicesForUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicedServiceServer).ListDevicesForUser(ctx, req.(*ListDevicesForUserRequest))
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

func _DevicedService_GrantDeviceToGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrantDeviceToGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicedServiceServer).GrantDeviceToGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.deviced.DevicedService/GrantDeviceToGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicedServiceServer).GrantDeviceToGroup(ctx, req.(*GrantDeviceToGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicedService_RevokeDeviceFromGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeDeviceFromGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicedServiceServer).RevokeDeviceFromGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.deviced.DevicedService/RevokeDeviceFromGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicedServiceServer).RevokeDeviceFromGroup(ctx, req.(*RevokeDeviceFromGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicedService_GrantDeviceToUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrantDeviceToUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicedServiceServer).GrantDeviceToUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.deviced.DevicedService/GrantDeviceToUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicedServiceServer).GrantDeviceToUser(ctx, req.(*GrantDeviceToUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicedService_RevokeDeviceFromUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeDeviceFromUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicedServiceServer).RevokeDeviceFromUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.deviced.DevicedService/RevokeDeviceFromUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicedServiceServer).RevokeDeviceFromUser(ctx, req.(*RevokeDeviceFromUserRequest))
	}
	return interceptor(ctx, in, info, handler)
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
			MethodName: "ListDevicesForUser",
			Handler:    _DevicedService_ListDevicesForUser_Handler,
		},
		{
			MethodName: "UnaryCall",
			Handler:    _DevicedService_UnaryCall_Handler,
		},
		{
			MethodName: "GrantDeviceToGroup",
			Handler:    _DevicedService_GrantDeviceToGroup_Handler,
		},
		{
			MethodName: "RevokeDeviceFromGroup",
			Handler:    _DevicedService_RevokeDeviceFromGroup_Handler,
		},
		{
			MethodName: "GrantDeviceToUser",
			Handler:    _DevicedService_GrantDeviceToUser_Handler,
		},
		{
			MethodName: "RevokeDeviceFromUser",
			Handler:    _DevicedService_RevokeDeviceFromUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _DevicedService_Stream_Handler,
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

func init() { proto.RegisterFile("service.proto", fileDescriptor_service_7e04fa9b03e18c9e) }

var fileDescriptor_service_7e04fa9b03e18c9e = []byte{
	// 546 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x73, 0xd3, 0x30,
	0x10, 0x4d, 0x2e, 0x65, 0xa2, 0xa6, 0x40, 0xc5, 0xc7, 0xc1, 0x99, 0x72, 0xc8, 0x89, 0x03, 0x38,
	0x69, 0x7a, 0xe0, 0xeb, 0x46, 0x4b, 0xc3, 0x81, 0x03, 0xd3, 0x50, 0x66, 0x60, 0x98, 0xf1, 0x28,
	0xc9, 0xc6, 0xf6, 0xd4, 0xb6, 0x8c, 0x24, 0x27, 0x93, 0xdf, 0xc0, 0x95, 0x1f, 0xcc, 0x58, 0x1f,
	0xae, 0x1c, 0x1a, 0x64, 0x73, 0xca, 0x64, 0xf5, 0xde, 0xbe, 0xdd, 0xa7, 0x67, 0xa1, 0x23, 0x0e,
	0x6c, 0x1d, 0x2f, 0xc0, 0xcf, 0x19, 0x15, 0x14, 0x9f, 0x90, 0xd8, 0x4f, 0x41, 0x10, 0x11, 0xc5,
	0x59, 0xc8, 0x7d, 0x73, 0xb8, 0x84, 0xf2, 0x67, 0xe9, 0x0d, 0x42, 0x4a, 0xc3, 0x04, 0x46, 0x12,
	0x3c, 0x2f, 0x56, 0x23, 0x48, 0x73, 0xb1, 0x55, 0x5c, 0xef, 0xd1, 0x82, 0x01, 0x11, 0x10, 0x28,
	0xb0, 0x29, 0x2e, 0x21, 0x81, 0xdd, 0x22, 0xce, 0x89, 0x58, 0x44, 0xf5, 0xda, 0xc3, 0x10, 0xc4,
	0x0e, 0x2a, 0x89, 0xb9, 0x29, 0x71, 0x5d, 0x3b, 0xe6, 0x11, 0xdd, 0xd4, 0x61, 0x0f, 0x22, 0x20,
	0x4c, 0xcc, 0x81, 0x08, 0x5d, 0xe8, 0x73, 0xc1, 0x80, 0xa4, 0xfa, 0xdf, 0xc0, 0xee, 0x12, 0xac,
	0x28, 0x0b, 0x0a, 0x0e, 0xcc, 0x88, 0x16, 0x19, 0x61, 0xdb, 0x60, 0x41, 0x92, 0xa4, 0x12, 0x90,
	0x64, 0xbb, 0x34, 0x08, 0x19, 0xc9, 0x4c, 0x8b, 0x40, 0xd0, 0x20, 0x64, 0xb4, 0xc8, 0xf5, 0xe1,
	0x33, 0x06, 0x6b, 0x7a, 0x63, 0xf6, 0x0b, 0x56, 0x8c, 0xa6, 0xb5, 0x73, 0x6f, 0x97, 0x6c, 0xa9,
	0x9f, 0xdc, 0xc1, 0xbd, 0x3d, 0x9e, 0xfc, 0x3e, 0x42, 0xf7, 0x2f, 0x94, 0xf1, 0x33, 0x75, 0x0f,
	0x78, 0x8b, 0xfa, 0xe7, 0xd2, 0x64, 0x55, 0xc7, 0x13, 0xff, 0x9f, 0xf7, 0xe5, 0xdb, 0xe0, 0x2b,
	0xf8, 0x59, 0x00, 0x17, 0xde, 0x59, 0x2b, 0x0e, 0xcf, 0x69, 0xc6, 0x61, 0xd8, 0xc1, 0x3f, 0x50,
	0xff, 0x42, 0x5e, 0x65, 0x43, 0x69, 0x1b, 0x6c, 0xa4, 0x9f, 0xfa, 0x2a, 0x3f, 0xbe, 0xc9, 0x8f,
	0xff, 0xa1, 0xcc, 0xcf, 0xb0, 0x83, 0xd7, 0xe8, 0xf0, 0x73, 0x99, 0x09, 0xdd, 0xfc, 0xd4, 0xd1,
	0xdc, 0xc2, 0x9a, 0xde, 0x93, 0x36, 0x94, 0x6a, 0xab, 0x1c, 0xf5, 0xa6, 0x20, 0xb4, 0xea, 0xc8,
	0xd1, 0xa2, 0x42, 0x1a, 0xcd, 0x71, 0x73, 0x42, 0xa5, 0xb8, 0x46, 0x87, 0x9f, 0x62, 0xae, 0xeb,
	0xdc, 0xb9, 0xa9, 0x85, 0x6d, 0xba, 0x69, 0x8d, 0x52, 0xe9, 0x7e, 0x43, 0x68, 0x16, 0xd1, 0x8d,
	0x5e, 0x75, 0xcf, 0x4d, 0x78, 0xae, 0x71, 0x6e, 0x5b, 0x58, 0xad, 0xbf, 0xa2, 0xde, 0x47, 0xf3,
	0x0d, 0x3a, 0x4d, 0xac, 0x90, 0xee, 0x50, 0xdc, 0xa0, 0x83, 0x99, 0xfc, 0x1a, 0xf1, 0x4b, 0xd7,
	0x58, 0x12, 0x66, 0x46, 0xf2, 0x5e, 0x34, 0x84, 0xcb, 0x01, 0x86, 0x9d, 0xe7, 0xdd, 0x71, 0x17,
	0xff, 0xea, 0x22, 0x6c, 0x39, 0x77, 0x49, 0xd9, 0x35, 0x07, 0x86, 0x5f, 0x37, 0x37, 0x5b, 0x53,
	0xcc, 0x5e, 0x6f, 0xfe, 0x83, 0x69, 0xe7, 0xf2, 0xba, 0x7c, 0x9a, 0xce, 0x49, 0x92, 0x38, 0x2d,
	0xad, 0x90, 0x4d, 0x73, 0x69, 0x11, 0x2a, 0xc5, 0x0d, 0x42, 0xca, 0x16, 0x29, 0x39, 0x6e, 0xe4,
	0xa0, 0xad, 0x79, 0xda, 0x82, 0x61, 0x44, 0xa5, 0xf1, 0x11, 0xc2, 0xd3, 0xf2, 0x8d, 0x54, 0x5e,
	0x7c, 0xa1, 0xd3, 0xf2, 0xf5, 0x74, 0xfa, 0xfe, 0x37, 0xc5, 0x9d, 0xa7, 0x0c, 0x3d, 0xb9, 0x92,
	0x2f, 0xae, 0xe2, 0x5d, 0x32, 0x9a, 0x2a, 0xb1, 0x77, 0x0e, 0xb1, 0x3b, 0x59, 0x6e, 0xbd, 0x15,
	0x3a, 0xae, 0x8d, 0x29, 0x03, 0xf5, 0xaa, 0xcd, 0x62, 0x76, 0x9e, 0xf6, 0xeb, 0x24, 0xe8, 0xf1,
	0xee, 0x84, 0x52, 0xea, 0x6d, 0xcb, 0xb5, 0x1a, 0xa9, 0xbd, 0xef, 0x7d, 0xbf, 0xa7, 0x1b, 0xcc,
	0x0f, 0xe4, 0xe1, 0xd9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xe1, 0x95, 0x2f, 0x45, 0x08,
	0x00, 0x00,
}
