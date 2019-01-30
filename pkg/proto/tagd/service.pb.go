// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package tagd

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0x0b, 0x6b, 0x17, 0x4c, 0xf5, 0x92, 0x83, 0x87, 0xf5, 0xd6, 0x83, 0xda, 0xcb, 0x14,
	0xf4, 0x0d, 0x04, 0xd9, 0x83, 0x78, 0xb0, 0xae, 0x20, 0x1e, 0x84, 0xd4, 0x8e, 0xe3, 0x82, 0xbb,
	0xd9, 0x26, 0x93, 0x82, 0xef, 0xe4, 0x43, 0x4a, 0x92, 0x06, 0x4f, 0x66, 0x7b, 0xcb, 0x9f, 0xdf,
	0xfc, 0xf2, 0x7d, 0x11, 0xa7, 0x16, 0xcd, 0xae, 0x7d, 0x47, 0x18, 0x8c, 0x66, 0x2d, 0x2b, 0xd5,
	0x42, 0x87, 0xac, 0xf8, 0xb3, 0xed, 0xc9, 0x42, 0xba, 0x64, 0x45, 0x9b, 0xea, 0x9c, 0xb4, 0xa6,
	0x2f, 0x5c, 0x06, 0x72, 0xed, 0x3e, 0x96, 0xd8, 0x0d, 0xfc, 0x1d, 0x07, 0xab, 0x63, 0x56, 0xb4,
	0x5f, 0xce, 0x5c, 0xff, 0xb7, 0x39, 0x31, 0xd8, 0xe9, 0x1d, 0x26, 0x8a, 0x90, 0x13, 0xb5, 0x75,
	0x68, 0xf6, 0xd3, 0xd7, 0x3f, 0x85, 0x98, 0x35, 0x8a, 0x36, 0x4f, 0xf1, 0x3d, 0x59, 0x8b, 0xa2,
	0x51, 0x24, 0x2f, 0xe0, 0xff, 0x38, 0xd0, 0x28, 0x5a, 0xe1, 0xd6, 0xa1, 0xe5, 0xea, 0x0c, 0x62,
	0x34, 0x48, 0xd1, 0xe0, 0xce, 0x47, 0x9b, 0x4f, 0xe4, 0xbd, 0x98, 0x3e, 0xfb, 0x34, 0xf2, 0x2a,
	0xa7, 0x0a, 0xc8, 0xb8, 0xec, 0x41, 0x94, 0xab, 0xd0, 0x46, 0x2e, 0x72, 0xb6, 0xc8, 0x8c, 0xeb,
	0x5e, 0x44, 0x51, 0x23, 0xe7, 0x4b, 0xd6, 0xc8, 0x49, 0x74, 0x39, 0xca, 0xd9, 0x41, 0xf7, 0x16,
	0xe7, 0x13, 0xf9, 0x26, 0xa6, 0x8f, 0xfe, 0x77, 0xf3, 0xad, 0x03, 0x92, 0xec, 0x8b, 0x03, 0xc8,
	0xe4, 0xbf, 0x2d, 0x5f, 0x8f, 0xfc, 0xf9, 0xba, 0x0c, 0x9d, 0x6e, 0x7e, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x78, 0x71, 0x9f, 0x21, 0x45, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TagdServiceClient is the client API for TagdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TagdServiceClient interface {
	Tag(ctx context.Context, in *TagRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Untag(ctx context.Context, in *UntagRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type tagdServiceClient struct {
	cc *grpc.ClientConn
}

func NewTagdServiceClient(cc *grpc.ClientConn) TagdServiceClient {
	return &tagdServiceClient{cc}
}

func (c *tagdServiceClient) Tag(ctx context.Context, in *TagRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ai.metathings.service.tagd.TagdService/Tag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagdServiceClient) Untag(ctx context.Context, in *UntagRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ai.metathings.service.tagd.TagdService/Untag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagdServiceClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ai.metathings.service.tagd.TagdService/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagdServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/ai.metathings.service.tagd.TagdService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagdServiceClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/ai.metathings.service.tagd.TagdService/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagdServiceServer is the server API for TagdService service.
type TagdServiceServer interface {
	Tag(context.Context, *TagRequest) (*empty.Empty, error)
	Untag(context.Context, *UntagRequest) (*empty.Empty, error)
	Remove(context.Context, *RemoveRequest) (*empty.Empty, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
}

func RegisterTagdServiceServer(s *grpc.Server, srv TagdServiceServer) {
	s.RegisterService(&_TagdService_serviceDesc, srv)
}

func _TagdService_Tag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagdServiceServer).Tag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.tagd.TagdService/Tag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagdServiceServer).Tag(ctx, req.(*TagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagdService_Untag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UntagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagdServiceServer).Untag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.tagd.TagdService/Untag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagdServiceServer).Untag(ctx, req.(*UntagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagdService_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagdServiceServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.tagd.TagdService/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagdServiceServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagdService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagdServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.tagd.TagdService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagdServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagdService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagdServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.tagd.TagdService/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagdServiceServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TagdService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ai.metathings.service.tagd.TagdService",
	HandlerType: (*TagdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Tag",
			Handler:    _TagdService_Tag_Handler,
		},
		{
			MethodName: "Untag",
			Handler:    _TagdService_Untag_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _TagdService_Remove_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _TagdService_Get_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _TagdService_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
