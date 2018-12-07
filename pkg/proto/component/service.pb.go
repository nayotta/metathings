// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package component

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// Client API for ModuleService service

type ModuleServiceClient interface {
	UnaryCall(ctx context.Context, in *UnaryCallRequest, opts ...grpc.CallOption) (*UnaryCallResponse, error)
	StreamCall(ctx context.Context, opts ...grpc.CallOption) (ModuleService_StreamCallClient, error)
}

type moduleServiceClient struct {
	cc *grpc.ClientConn
}

func NewModuleServiceClient(cc *grpc.ClientConn) ModuleServiceClient {
	return &moduleServiceClient{cc}
}

func (c *moduleServiceClient) UnaryCall(ctx context.Context, in *UnaryCallRequest, opts ...grpc.CallOption) (*UnaryCallResponse, error) {
	out := new(UnaryCallResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.component.ModuleService/UnaryCall", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moduleServiceClient) StreamCall(ctx context.Context, opts ...grpc.CallOption) (ModuleService_StreamCallClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ModuleService_serviceDesc.Streams[0], c.cc, "/ai.metathings.component.ModuleService/StreamCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &moduleServiceStreamCallClient{stream}
	return x, nil
}

type ModuleService_StreamCallClient interface {
	Send(*StreamCallRequest) error
	Recv() (*StreamCallResponse, error)
	grpc.ClientStream
}

type moduleServiceStreamCallClient struct {
	grpc.ClientStream
}

func (x *moduleServiceStreamCallClient) Send(m *StreamCallRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *moduleServiceStreamCallClient) Recv() (*StreamCallResponse, error) {
	m := new(StreamCallResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ModuleService service

type ModuleServiceServer interface {
	UnaryCall(context.Context, *UnaryCallRequest) (*UnaryCallResponse, error)
	StreamCall(ModuleService_StreamCallServer) error
}

func RegisterModuleServiceServer(s *grpc.Server, srv ModuleServiceServer) {
	s.RegisterService(&_ModuleService_serviceDesc, srv)
}

func _ModuleService_UnaryCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnaryCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModuleServiceServer).UnaryCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.component.ModuleService/UnaryCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModuleServiceServer).UnaryCall(ctx, req.(*UnaryCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModuleService_StreamCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ModuleServiceServer).StreamCall(&moduleServiceStreamCallServer{stream})
}

type ModuleService_StreamCallServer interface {
	Send(*StreamCallResponse) error
	Recv() (*StreamCallRequest, error)
	grpc.ServerStream
}

type moduleServiceStreamCallServer struct {
	grpc.ServerStream
}

func (x *moduleServiceStreamCallServer) Send(m *StreamCallResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *moduleServiceStreamCallServer) Recv() (*StreamCallRequest, error) {
	m := new(StreamCallRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ModuleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ai.metathings.component.ModuleService",
	HandlerType: (*ModuleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryCall",
			Handler:    _ModuleService_UnaryCall_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamCall",
			Handler:       _ModuleService_StreamCall_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_service_ee43eead3470e0f3) }

var fileDescriptor_service_ee43eead3470e0f3 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4f, 0xcc, 0xd4, 0xcb, 0x4d,
	0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x4b, 0xce, 0xcf, 0x2d, 0xc8, 0xcf, 0x4b,
	0xcd, 0x2b, 0x91, 0x12, 0x28, 0xcd, 0x4b, 0x2c, 0xaa, 0x8c, 0x4f, 0x4e, 0xcc, 0xc9, 0x81, 0x28,
	0x95, 0x12, 0x2c, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0x45, 0x12, 0x32, 0x7a, 0xc4, 0xc8, 0xc5, 0xeb,
	0x9b, 0x9f, 0x52, 0x9a, 0x93, 0x1a, 0x0c, 0x31, 0x55, 0x28, 0x85, 0x8b, 0x33, 0x14, 0xa4, 0xd1,
	0x39, 0x31, 0x27, 0x47, 0x48, 0x53, 0x0f, 0x87, 0xe9, 0x7a, 0x70, 0x35, 0x41, 0xa9, 0x85, 0xa5,
	0xa9, 0xc5, 0x25, 0x52, 0x5a, 0xc4, 0x28, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x55, 0x62, 0x10,
	0xca, 0xe6, 0xe2, 0x0a, 0x06, 0x3b, 0x06, 0x6c, 0x0d, 0x6e, 0xbd, 0x08, 0x45, 0x30, 0x7b, 0xb4,
	0x89, 0x52, 0x0b, 0xb3, 0x48, 0x83, 0xd1, 0x80, 0xd1, 0x89, 0x3b, 0x8a, 0x13, 0xae, 0x2a, 0x89,
	0x0d, 0xec, 0x71, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x35, 0xf5, 0x9d, 0xcc, 0x47, 0x01,
	0x00, 0x00,
}