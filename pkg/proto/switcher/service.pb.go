// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package switcher

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

// Client API for SwitcherService service

type SwitcherServiceClient interface {
	Info(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*InfoResponse, error)
	Toggle(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ToggleResponse, error)
	TurnOn(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TurnOnResponse, error)
	TurnOff(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TurnOffResponse, error)
}

type switcherServiceClient struct {
	cc *grpc.ClientConn
}

func NewSwitcherServiceClient(cc *grpc.ClientConn) SwitcherServiceClient {
	return &switcherServiceClient{cc}
}

func (c *switcherServiceClient) Info(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.switcher.SwitcherService/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *switcherServiceClient) Toggle(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ToggleResponse, error) {
	out := new(ToggleResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.switcher.SwitcherService/Toggle", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *switcherServiceClient) TurnOn(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TurnOnResponse, error) {
	out := new(TurnOnResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.switcher.SwitcherService/TurnOn", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *switcherServiceClient) TurnOff(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TurnOffResponse, error) {
	out := new(TurnOffResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.switcher.SwitcherService/TurnOff", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SwitcherService service

type SwitcherServiceServer interface {
	Info(context.Context, *empty.Empty) (*InfoResponse, error)
	Toggle(context.Context, *empty.Empty) (*ToggleResponse, error)
	TurnOn(context.Context, *empty.Empty) (*TurnOnResponse, error)
	TurnOff(context.Context, *empty.Empty) (*TurnOffResponse, error)
}

func RegisterSwitcherServiceServer(s *grpc.Server, srv SwitcherServiceServer) {
	s.RegisterService(&_SwitcherService_serviceDesc, srv)
}

func _SwitcherService_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwitcherServiceServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.switcher.SwitcherService/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwitcherServiceServer).Info(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwitcherService_Toggle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwitcherServiceServer).Toggle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.switcher.SwitcherService/Toggle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwitcherServiceServer).Toggle(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwitcherService_TurnOn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwitcherServiceServer).TurnOn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.switcher.SwitcherService/TurnOn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwitcherServiceServer).TurnOn(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwitcherService_TurnOff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwitcherServiceServer).TurnOff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.switcher.SwitcherService/TurnOff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwitcherServiceServer).TurnOff(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _SwitcherService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ai.metathings.service.switcher.SwitcherService",
	HandlerType: (*SwitcherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _SwitcherService_Info_Handler,
		},
		{
			MethodName: "Toggle",
			Handler:    _SwitcherService_Toggle_Handler,
		},
		{
			MethodName: "TurnOn",
			Handler:    _SwitcherService_TurnOn_Handler,
		},
		{
			MethodName: "TurnOff",
			Handler:    _SwitcherService_TurnOff_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_service_f5e2816be63b407e) }

var fileDescriptor_service_f5e2816be63b407e = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4b, 0xcc, 0xd4, 0xcb, 0x4d,
	0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x83, 0x49, 0x16, 0x97, 0x67, 0x96, 0x24,
	0x67, 0xa4, 0x16, 0x49, 0x49, 0xa7, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x83, 0x55, 0x27, 0x95,
	0xa6, 0xe9, 0xa7, 0xe6, 0x16, 0x94, 0x54, 0x42, 0x34, 0x4b, 0x71, 0x65, 0xe6, 0xa5, 0xe5, 0x43,
	0xd9, 0x3c, 0x25, 0xf9, 0xe9, 0xe9, 0x39, 0x50, 0x63, 0xa5, 0x78, 0x4b, 0x4a, 0x8b, 0xf2, 0xe2,
	0xf3, 0xf3, 0xa0, 0x5c, 0x3e, 0x08, 0x37, 0x2d, 0x0d, 0xc2, 0x37, 0xba, 0xcf, 0xc4, 0xc5, 0x1f,
	0x0c, 0xb5, 0x22, 0x18, 0x62, 0xa5, 0x90, 0x1f, 0x17, 0x8b, 0x67, 0x5e, 0x5a, 0xbe, 0x90, 0x98,
	0x1e, 0xc4, 0x4a, 0x3d, 0x98, 0x95, 0x7a, 0xae, 0x20, 0x2b, 0xa5, 0x74, 0xf4, 0xf0, 0x3b, 0x55,
	0x0f, 0xa4, 0x3b, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x89, 0x41, 0x28, 0x88, 0x8b,
	0x2d, 0x04, 0xec, 0x24, 0x9c, 0x26, 0xea, 0x11, 0x32, 0x11, 0xa2, 0x1f, 0xcd, 0xcc, 0xd2, 0xa2,
	0x3c, 0xff, 0x3c, 0x0a, 0xcc, 0x04, 0xeb, 0x47, 0x32, 0x33, 0x84, 0x8b, 0x1d, 0x2c, 0x96, 0x96,
	0x86, 0xd3, 0x50, 0x7d, 0xa2, 0x0c, 0x4d, 0x4b, 0x43, 0x98, 0xea, 0xc4, 0x15, 0xc5, 0x01, 0x93,
	0x4d, 0x62, 0x03, 0x1b, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x80, 0xd4, 0x0a, 0xf9, 0xfb,
	0x01, 0x00, 0x00,
}
