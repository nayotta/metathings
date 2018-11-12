// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package motor

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
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

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4e, 0xcc, 0xd4, 0xcb, 0x4d,
	0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x83, 0x49, 0xe6, 0xe6, 0x97, 0xe4, 0x17,
	0x49, 0x71, 0xe5, 0x64, 0x16, 0x97, 0x40, 0x14, 0x4a, 0x71, 0xa6, 0xa7, 0xc2, 0x98, 0x3c, 0xc5,
	0x25, 0x45, 0xa9, 0x89, 0xb9, 0x10, 0x9e, 0xd1, 0x66, 0x26, 0x2e, 0x1e, 0x5f, 0x90, 0xf2, 0x60,
	0x88, 0x5e, 0xa1, 0x58, 0x2e, 0x16, 0x9f, 0xcc, 0xe2, 0x12, 0x21, 0x0d, 0x3d, 0x3c, 0x66, 0xeb,
	0x81, 0x94, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x48, 0x69, 0x12, 0xa1, 0xb2, 0xb8, 0x20,
	0x3f, 0xaf, 0x38, 0x55, 0x89, 0x41, 0x28, 0x8a, 0x8b, 0xd9, 0x3d, 0xb5, 0x44, 0x48, 0x1d, 0xaf,
	0x1e, 0xf7, 0x54, 0xb8, 0xe1, 0x1a, 0x84, 0x15, 0xc2, 0xcd, 0xce, 0xe0, 0x62, 0x0b, 0x06, 0xfb,
	0x4d, 0x48, 0x1b, 0xaf, 0x2e, 0x88, 0x22, 0xa8, 0x0d, 0xc5, 0x52, 0xc4, 0x29, 0x86, 0xd9, 0xa2,
	0xc1, 0x68, 0xc0, 0xe8, 0xc4, 0x1e, 0xc5, 0x0a, 0x96, 0x4d, 0x62, 0x03, 0x87, 0xa2, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0xd7, 0x58, 0xe2, 0x85, 0x98, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MotorServiceClient is the client API for MotorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MotorServiceClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Stream(ctx context.Context, opts ...grpc.CallOption) (MotorService_StreamClient, error)
}

type motorServiceClient struct {
	cc *grpc.ClientConn
}

func NewMotorServiceClient(cc *grpc.ClientConn) MotorServiceClient {
	return &motorServiceClient{cc}
}

func (c *motorServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/ai.metathings.service.motor.MotorService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *motorServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/ai.metathings.service.motor.MotorService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *motorServiceClient) Stream(ctx context.Context, opts ...grpc.CallOption) (MotorService_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MotorService_serviceDesc.Streams[0], "/ai.metathings.service.motor.MotorService/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &motorServiceStreamClient{stream}
	return x, nil
}

type MotorService_StreamClient interface {
	Send(*StreamRequests) error
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type motorServiceStreamClient struct {
	grpc.ClientStream
}

func (x *motorServiceStreamClient) Send(m *StreamRequests) error {
	return x.ClientStream.SendMsg(m)
}

func (x *motorServiceStreamClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MotorServiceServer is the server API for MotorService service.
type MotorServiceServer interface {
	List(context.Context, *ListRequest) (*ListResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Stream(MotorService_StreamServer) error
}

func RegisterMotorServiceServer(s *grpc.Server, srv MotorServiceServer) {
	s.RegisterService(&_MotorService_serviceDesc, srv)
}

func _MotorService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MotorServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.motor.MotorService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MotorServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MotorService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MotorServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.motor.MotorService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MotorServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MotorService_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MotorServiceServer).Stream(&motorServiceStreamServer{stream})
}

type MotorService_StreamServer interface {
	Send(*StreamResponse) error
	Recv() (*StreamRequests, error)
	grpc.ServerStream
}

type motorServiceStreamServer struct {
	grpc.ServerStream
}

func (x *motorServiceStreamServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *motorServiceStreamServer) Recv() (*StreamRequests, error) {
	m := new(StreamRequests)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MotorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ai.metathings.service.motor.MotorService",
	HandlerType: (*MotorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _MotorService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _MotorService_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _MotorService_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}
