// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf3 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CoreService service

type CoreServiceClient interface {
	CreateCore(ctx context.Context, in *CreateCoreRequest, opts ...grpc.CallOption) (*CreateCoreResponse, error)
	DeleteCore(ctx context.Context, in *DeleteCoreRequest, opts ...grpc.CallOption) (*google_protobuf3.Empty, error)
	GetCore(ctx context.Context, in *GetCoreRequest, opts ...grpc.CallOption) (*GetCoreResponse, error)
	ListCores(ctx context.Context, in *ListCoresRequest, opts ...grpc.CallOption) (*ListCoresResponse, error)
	Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*google_protobuf3.Empty, error)
	Pipeline(ctx context.Context, opts ...grpc.CallOption) (CoreService_PipelineClient, error)
	ListCoresForUser(ctx context.Context, in *ListCoresForUserRequest, opts ...grpc.CallOption) (*ListCoresForUserResponse, error)
	SendUnaryCall(ctx context.Context, in *SendUnaryCallRequest, opts ...grpc.CallOption) (*SendUnaryCallResponse, error)
}

type coreServiceClient struct {
	cc *grpc.ClientConn
}

func NewCoreServiceClient(cc *grpc.ClientConn) CoreServiceClient {
	return &coreServiceClient{cc}
}

func (c *coreServiceClient) CreateCore(ctx context.Context, in *CreateCoreRequest, opts ...grpc.CallOption) (*CreateCoreResponse, error) {
	out := new(CreateCoreResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.core.CoreService/CreateCore", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) DeleteCore(ctx context.Context, in *DeleteCoreRequest, opts ...grpc.CallOption) (*google_protobuf3.Empty, error) {
	out := new(google_protobuf3.Empty)
	err := grpc.Invoke(ctx, "/ai.metathings.service.core.CoreService/DeleteCore", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) GetCore(ctx context.Context, in *GetCoreRequest, opts ...grpc.CallOption) (*GetCoreResponse, error) {
	out := new(GetCoreResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.core.CoreService/GetCore", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) ListCores(ctx context.Context, in *ListCoresRequest, opts ...grpc.CallOption) (*ListCoresResponse, error) {
	out := new(ListCoresResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.core.CoreService/ListCores", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*google_protobuf3.Empty, error) {
	out := new(google_protobuf3.Empty)
	err := grpc.Invoke(ctx, "/ai.metathings.service.core.CoreService/Heartbeat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) Pipeline(ctx context.Context, opts ...grpc.CallOption) (CoreService_PipelineClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_CoreService_serviceDesc.Streams[0], c.cc, "/ai.metathings.service.core.CoreService/Pipeline", opts...)
	if err != nil {
		return nil, err
	}
	x := &coreServicePipelineClient{stream}
	return x, nil
}

type CoreService_PipelineClient interface {
	Send(*PipelineOutStream) error
	Recv() (*PipelineInStream, error)
	grpc.ClientStream
}

type coreServicePipelineClient struct {
	grpc.ClientStream
}

func (x *coreServicePipelineClient) Send(m *PipelineOutStream) error {
	return x.ClientStream.SendMsg(m)
}

func (x *coreServicePipelineClient) Recv() (*PipelineInStream, error) {
	m := new(PipelineInStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *coreServiceClient) ListCoresForUser(ctx context.Context, in *ListCoresForUserRequest, opts ...grpc.CallOption) (*ListCoresForUserResponse, error) {
	out := new(ListCoresForUserResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.core.CoreService/ListCoresForUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) SendUnaryCall(ctx context.Context, in *SendUnaryCallRequest, opts ...grpc.CallOption) (*SendUnaryCallResponse, error) {
	out := new(SendUnaryCallResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.core.CoreService/SendUnaryCall", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CoreService service

type CoreServiceServer interface {
	CreateCore(context.Context, *CreateCoreRequest) (*CreateCoreResponse, error)
	DeleteCore(context.Context, *DeleteCoreRequest) (*google_protobuf3.Empty, error)
	GetCore(context.Context, *GetCoreRequest) (*GetCoreResponse, error)
	ListCores(context.Context, *ListCoresRequest) (*ListCoresResponse, error)
	Heartbeat(context.Context, *HeartbeatRequest) (*google_protobuf3.Empty, error)
	Pipeline(CoreService_PipelineServer) error
	ListCoresForUser(context.Context, *ListCoresForUserRequest) (*ListCoresForUserResponse, error)
	SendUnaryCall(context.Context, *SendUnaryCallRequest) (*SendUnaryCallResponse, error)
}

func RegisterCoreServiceServer(s *grpc.Server, srv CoreServiceServer) {
	s.RegisterService(&_CoreService_serviceDesc, srv)
}

func _CoreService_CreateCore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).CreateCore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.core.CoreService/CreateCore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).CreateCore(ctx, req.(*CreateCoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_DeleteCore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).DeleteCore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.core.CoreService/DeleteCore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).DeleteCore(ctx, req.(*DeleteCoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_GetCore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).GetCore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.core.CoreService/GetCore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).GetCore(ctx, req.(*GetCoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_ListCores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).ListCores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.core.CoreService/ListCores",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).ListCores(ctx, req.(*ListCoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.core.CoreService/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).Heartbeat(ctx, req.(*HeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_Pipeline_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CoreServiceServer).Pipeline(&coreServicePipelineServer{stream})
}

type CoreService_PipelineServer interface {
	Send(*PipelineInStream) error
	Recv() (*PipelineOutStream, error)
	grpc.ServerStream
}

type coreServicePipelineServer struct {
	grpc.ServerStream
}

func (x *coreServicePipelineServer) Send(m *PipelineInStream) error {
	return x.ServerStream.SendMsg(m)
}

func (x *coreServicePipelineServer) Recv() (*PipelineOutStream, error) {
	m := new(PipelineOutStream)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CoreService_ListCoresForUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCoresForUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).ListCoresForUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.core.CoreService/ListCoresForUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).ListCoresForUser(ctx, req.(*ListCoresForUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_SendUnaryCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendUnaryCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).SendUnaryCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.core.CoreService/SendUnaryCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).SendUnaryCall(ctx, req.(*SendUnaryCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ai.metathings.service.core.CoreService",
	HandlerType: (*CoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCore",
			Handler:    _CoreService_CreateCore_Handler,
		},
		{
			MethodName: "DeleteCore",
			Handler:    _CoreService_DeleteCore_Handler,
		},
		{
			MethodName: "GetCore",
			Handler:    _CoreService_GetCore_Handler,
		},
		{
			MethodName: "ListCores",
			Handler:    _CoreService_ListCores_Handler,
		},
		{
			MethodName: "Heartbeat",
			Handler:    _CoreService_Heartbeat_Handler,
		},
		{
			MethodName: "ListCoresForUser",
			Handler:    _CoreService_ListCoresForUser_Handler,
		},
		{
			MethodName: "SendUnaryCall",
			Handler:    _CoreService_SendUnaryCall_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Pipeline",
			Handler:       _CoreService_Pipeline_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x5f, 0x4f, 0xea, 0x30,
	0x18, 0xc6, 0x21, 0x39, 0xe1, 0x1c, 0x7a, 0x82, 0x62, 0x13, 0x4d, 0x9c, 0x77, 0x5c, 0x19, 0x85,
	0x82, 0xe2, 0x27, 0x10, 0xff, 0x26, 0x26, 0x1a, 0x17, 0x6e, 0xbc, 0x59, 0xca, 0xf6, 0x32, 0x6a,
	0xba, 0x76, 0xb6, 0x1d, 0x09, 0x57, 0x7e, 0x57, 0x3f, 0x89, 0xd9, 0x9f, 0x0e, 0xb9, 0x70, 0x8c,
	0x4b, 0x9e, 0xfe, 0x9e, 0xf7, 0xe9, 0xfb, 0x94, 0xa1, 0x8e, 0x06, 0xb5, 0x64, 0x3e, 0x90, 0x58,
	0x49, 0x23, 0xb1, 0x43, 0x19, 0x89, 0xc0, 0x50, 0xb3, 0x60, 0x22, 0xd4, 0xc4, 0x1e, 0xfa, 0x52,
	0x81, 0x73, 0x12, 0x4a, 0x19, 0x72, 0x18, 0x66, 0xe4, 0x2c, 0x99, 0x0f, 0x21, 0x8a, 0xcd, 0x2a,
	0x37, 0x3a, 0x07, 0xbe, 0x02, 0x6a, 0xc0, 0x4b, 0x49, 0x2b, 0x05, 0xc0, 0x61, 0x53, 0xda, 0x0b,
	0xc1, 0xfc, 0xfc, 0xdd, 0xe5, 0x4c, 0xe7, 0x82, 0x2e, 0x94, 0xfd, 0x05, 0x50, 0x65, 0x66, 0x40,
	0x8d, 0xb5, 0xc4, 0x2c, 0x06, 0xce, 0x84, 0xb5, 0x1c, 0xaf, 0x2d, 0xde, 0x5c, 0x2a, 0x2f, 0xd1,
	0xa0, 0x8a, 0xa3, 0x43, 0x0d, 0x22, 0xf0, 0x12, 0x41, 0xd5, 0xca, 0xf3, 0x29, 0xe7, 0xb9, 0x7c,
	0xf9, 0xd5, 0x42, 0xff, 0x27, 0x52, 0x81, 0x9b, 0x2f, 0x83, 0x23, 0x84, 0x26, 0xd9, 0x65, 0x53,
	0x11, 0x0f, 0xc8, 0xef, 0x2b, 0x93, 0x35, 0xf7, 0x0a, 0x1f, 0x09, 0x68, 0xe3, 0x90, 0xba, 0xb8,
	0x8e, 0xa5, 0xd0, 0xd0, 0x6b, 0xe0, 0x29, 0x42, 0x37, 0x59, 0x11, 0xdb, 0xe3, 0xd6, 0x9c, 0x8d,
	0x3b, 0x22, 0x79, 0xe9, 0xc4, 0x96, 0x4e, 0x6e, 0xd3, 0xd2, 0x7b, 0x0d, 0x1c, 0xa0, 0xbf, 0xf7,
	0x60, 0xb2, 0x99, 0x67, 0x55, 0x33, 0x0b, 0xc8, 0x0e, 0x3c, 0xaf, 0xc5, 0x96, 0x97, 0x7f, 0x47,
	0xed, 0x27, 0xa6, 0x33, 0x55, 0xe3, 0x7e, 0x95, 0xb7, 0xc4, 0x6c, 0xd2, 0xa0, 0x26, 0x5d, 0x66,
	0xb9, 0xa8, 0xfd, 0x60, 0x1f, 0xbf, 0x3a, 0xab, 0xc4, 0xb6, 0xd7, 0x14, 0xa1, 0x7f, 0x2f, 0xc5,
	0x1f, 0xa8, 0xba, 0x7b, 0x4b, 0x3d, 0x27, 0xc6, 0x35, 0x0a, 0x68, 0xe4, 0xf4, 0xeb, 0xe0, 0x8f,
	0x22, 0xa7, 0x7b, 0x8d, 0xd3, 0xe6, 0xa8, 0x89, 0x3f, 0x51, 0xb7, 0x5c, 0xed, 0x4e, 0xaa, 0xa9,
	0x06, 0x85, 0xc7, 0xb5, 0x8a, 0x28, 0x68, 0xbb, 0xd1, 0xd5, 0x6e, 0xa6, 0xb2, 0xc4, 0x25, 0xea,
	0xb8, 0x20, 0x82, 0x69, 0xfa, 0x11, 0x4c, 0x28, 0xe7, 0x78, 0x54, 0x35, 0x68, 0x03, 0xb5, 0xd1,
	0x17, 0x3b, 0x38, 0x6c, 0xee, 0x75, 0xeb, 0xed, 0x4f, 0x7a, 0x3e, 0x6b, 0x65, 0x2f, 0x30, 0xfe,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xa4, 0x33, 0xbf, 0x58, 0x04, 0x00, 0x00,
}