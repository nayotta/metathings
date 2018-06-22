// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package camerad

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

// Client API for CameradService service

type CameradServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Patch(ctx context.Context, in *PatchRequest, opts ...grpc.CallOption) (*PatchResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	ListForUser(ctx context.Context, in *ListForUserRequest, opts ...grpc.CallOption) (*ListForUserResponse, error)
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error)
	// for camera entity call
	Callback(ctx context.Context, in *CallbackRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type cameradServiceClient struct {
	cc *grpc.ClientConn
}

func NewCameradServiceClient(cc *grpc.ClientConn) CameradServiceClient {
	return &cameradServiceClient{cc}
}

func (c *cameradServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.camerad.CameradService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cameradServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/ai.metathings.service.camerad.CameradService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cameradServiceClient) Patch(ctx context.Context, in *PatchRequest, opts ...grpc.CallOption) (*PatchResponse, error) {
	out := new(PatchResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.camerad.CameradService/Patch", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cameradServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.camerad.CameradService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cameradServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.camerad.CameradService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cameradServiceClient) ListForUser(ctx context.Context, in *ListForUserRequest, opts ...grpc.CallOption) (*ListForUserResponse, error) {
	out := new(ListForUserResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.camerad.CameradService/ListForUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cameradServiceClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.camerad.CameradService/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cameradServiceClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error) {
	out := new(StopResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.camerad.CameradService/Stop", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cameradServiceClient) Callback(ctx context.Context, in *CallbackRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/ai.metathings.service.camerad.CameradService/Callback", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CameradService service

type CameradServiceServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Delete(context.Context, *DeleteRequest) (*empty.Empty, error)
	Patch(context.Context, *PatchRequest) (*PatchResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	ListForUser(context.Context, *ListForUserRequest) (*ListForUserResponse, error)
	Start(context.Context, *StartRequest) (*StartResponse, error)
	Stop(context.Context, *StopRequest) (*StopResponse, error)
	// for camera entity call
	Callback(context.Context, *CallbackRequest) (*empty.Empty, error)
}

func RegisterCameradServiceServer(s *grpc.Server, srv CameradServiceServer) {
	s.RegisterService(&_CameradService_serviceDesc, srv)
}

func _CameradService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CameradServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.camerad.CameradService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CameradServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CameradService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CameradServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.camerad.CameradService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CameradServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CameradService_Patch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CameradServiceServer).Patch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.camerad.CameradService/Patch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CameradServiceServer).Patch(ctx, req.(*PatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CameradService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CameradServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.camerad.CameradService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CameradServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CameradService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CameradServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.camerad.CameradService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CameradServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CameradService_ListForUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListForUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CameradServiceServer).ListForUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.camerad.CameradService/ListForUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CameradServiceServer).ListForUser(ctx, req.(*ListForUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CameradService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CameradServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.camerad.CameradService/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CameradServiceServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CameradService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CameradServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.camerad.CameradService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CameradServiceServer).Stop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CameradService_Callback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CameradServiceServer).Callback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.camerad.CameradService/Callback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CameradServiceServer).Callback(ctx, req.(*CallbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CameradService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ai.metathings.service.camerad.CameradService",
	HandlerType: (*CameradServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CameradService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CameradService_Delete_Handler,
		},
		{
			MethodName: "Patch",
			Handler:    _CameradService_Patch_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CameradService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _CameradService_List_Handler,
		},
		{
			MethodName: "ListForUser",
			Handler:    _CameradService_ListForUser_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _CameradService_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _CameradService_Stop_Handler,
		},
		{
			MethodName: "Callback",
			Handler:    _CameradService_Callback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_service_1a28121d3573c15a) }

var fileDescriptor_service_1a28121d3573c15a = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x4e, 0x02, 0x31,
	0x10, 0xc6, 0x49, 0x14, 0x90, 0xa2, 0x1c, 0x6a, 0xe2, 0x01, 0xe3, 0x85, 0x9b, 0x82, 0x25, 0xe2,
	0x1b, 0x88, 0xca, 0xc5, 0x03, 0x11, 0xbd, 0x78, 0x90, 0x94, 0xdd, 0x61, 0xd9, 0xb8, 0xd0, 0xb5,
	0x9d, 0x25, 0xf1, 0x19, 0x7c, 0x69, 0xd3, 0xed, 0x9f, 0x78, 0xb2, 0xf5, 0x36, 0xb3, 0xfd, 0x7d,
	0xfd, 0xa6, 0x5f, 0xbb, 0xe4, 0x44, 0x81, 0xdc, 0xe7, 0x09, 0xb0, 0x52, 0x0a, 0x14, 0xf4, 0x82,
	0xe7, 0x6c, 0x0b, 0xc8, 0x71, 0x93, 0xef, 0x32, 0xc5, 0xdc, 0x62, 0xc2, 0xb7, 0x20, 0x79, 0xda,
	0x3f, 0xcf, 0x84, 0xc8, 0x0a, 0x18, 0xd7, 0xf0, 0xaa, 0x5a, 0x8f, 0x61, 0x5b, 0xe2, 0x97, 0xd1,
	0xf6, 0x8f, 0x13, 0x09, 0x1c, 0xc1, 0x75, 0x29, 0x14, 0xe0, 0xbb, 0x6e, 0xc9, 0x31, 0xd9, 0xd8,
	0xa6, 0x93, 0x01, 0xda, 0x92, 0x14, 0xb9, 0x72, 0xf5, 0xa9, 0xae, 0x97, 0x6b, 0x21, 0x97, 0x95,
	0x02, 0xe9, 0x84, 0x0a, 0xb9, 0xf4, 0xb4, 0x42, 0x51, 0xda, 0xba, 0x97, 0xf0, 0xa2, 0x58, 0xf1,
	0xe4, 0xc3, 0xf4, 0x93, 0xef, 0x36, 0xe9, 0x4d, 0xcd, 0x98, 0x0b, 0x33, 0x35, 0xcd, 0x48, 0x6b,
	0x5a, 0x8f, 0x44, 0x47, 0xec, 0xcf, 0x73, 0x31, 0x83, 0x3d, 0xc3, 0x67, 0x05, 0x0a, 0xfb, 0xd7,
	0x91, 0xb4, 0x2a, 0xc5, 0x4e, 0xc1, 0xa0, 0x41, 0xe7, 0xa4, 0x75, 0x5f, 0x9f, 0x36, 0x68, 0x64,
	0x30, 0x67, 0x74, 0xc6, 0x4c, 0x9e, 0xcc, 0xe5, 0xc9, 0x1e, 0x74, 0x9e, 0x83, 0x06, 0x4d, 0x49,
	0x73, 0xae, 0x13, 0xa3, 0xc3, 0xc0, 0x86, 0x35, 0xe5, 0xf6, 0x1b, 0xc5, 0xc1, 0x7e, 0xee, 0x77,
	0x72, 0x30, 0x03, 0xa4, 0x97, 0x01, 0xd9, 0x0c, 0xd0, 0x39, 0x5c, 0xc5, 0xa0, 0x7e, 0x7f, 0x4e,
	0x0e, 0x9f, 0x72, 0x85, 0x34, 0xa4, 0xd2, 0x90, 0x73, 0x18, 0x46, 0xb1, 0xde, 0x62, 0x4f, 0xba,
	0xfa, 0xcb, 0xa3, 0x90, 0xaf, 0x0a, 0x24, 0xbd, 0x89, 0x50, 0x5b, 0xd6, 0x19, 0x4e, 0xfe, 0x23,
	0xf1, 0xbe, 0x29, 0x69, 0x2e, 0xf4, 0xcb, 0x0c, 0x5e, 0x50, 0x4d, 0xc5, 0x5e, 0x90, 0x85, 0x7f,
	0x07, 0xb8, 0x40, 0x51, 0x06, 0x03, 0xd4, 0x50, 0x6c, 0x80, 0x86, 0xf5, 0x16, 0x2f, 0xe4, 0x68,
	0x6a, 0xff, 0x24, 0xca, 0x42, 0x0f, 0xdf, 0x82, 0xc1, 0xf7, 0x7b, 0xd7, 0x79, 0x6b, 0x5b, 0xd1,
	0xaa, 0x55, 0x2f, 0xde, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xcb, 0xf8, 0x2b, 0x06, 0x6a, 0x04,
	0x00, 0x00,
}
