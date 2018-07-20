// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package sensor

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

// Client API for SensorService service

type SensorServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Patch(ctx context.Context, in *PatchRequest, opts ...grpc.CallOption) (*PatchResponse, error)
	GetData(ctx context.Context, in *GetDataRequest, opts ...grpc.CallOption) (*GetDataResponse, error)
	ListData(ctx context.Context, in *ListDataRequest, opts ...grpc.CallOption) (*ListDataResponse, error)
}

type sensorServiceClient struct {
	cc *grpc.ClientConn
}

func NewSensorServiceClient(cc *grpc.ClientConn) SensorServiceClient {
	return &sensorServiceClient{cc}
}

func (c *sensorServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.sensor.SensorService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sensorServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.sensor.SensorService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sensorServiceClient) Patch(ctx context.Context, in *PatchRequest, opts ...grpc.CallOption) (*PatchResponse, error) {
	out := new(PatchResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.sensor.SensorService/Patch", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sensorServiceClient) GetData(ctx context.Context, in *GetDataRequest, opts ...grpc.CallOption) (*GetDataResponse, error) {
	out := new(GetDataResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.sensor.SensorService/GetData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sensorServiceClient) ListData(ctx context.Context, in *ListDataRequest, opts ...grpc.CallOption) (*ListDataResponse, error) {
	out := new(ListDataResponse)
	err := grpc.Invoke(ctx, "/ai.metathings.service.sensor.SensorService/ListData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SensorService service

type SensorServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Patch(context.Context, *PatchRequest) (*PatchResponse, error)
	GetData(context.Context, *GetDataRequest) (*GetDataResponse, error)
	ListData(context.Context, *ListDataRequest) (*ListDataResponse, error)
}

func RegisterSensorServiceServer(s *grpc.Server, srv SensorServiceServer) {
	s.RegisterService(&_SensorService_serviceDesc, srv)
}

func _SensorService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SensorServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.sensor.SensorService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SensorServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SensorService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SensorServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.sensor.SensorService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SensorServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SensorService_Patch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SensorServiceServer).Patch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.sensor.SensorService/Patch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SensorServiceServer).Patch(ctx, req.(*PatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SensorService_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SensorServiceServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.sensor.SensorService/GetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SensorServiceServer).GetData(ctx, req.(*GetDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SensorService_ListData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SensorServiceServer).ListData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.sensor.SensorService/ListData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SensorServiceServer).ListData(ctx, req.(*ListDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SensorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ai.metathings.service.sensor.SensorService",
	HandlerType: (*SensorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _SensorService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _SensorService_List_Handler,
		},
		{
			MethodName: "Patch",
			Handler:    _SensorService_Patch_Handler,
		},
		{
			MethodName: "GetData",
			Handler:    _SensorService_GetData_Handler,
		},
		{
			MethodName: "ListData",
			Handler:    _SensorService_ListData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_service_118ce9c19a6a9758) }

var fileDescriptor_service_118ce9c19a6a9758 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4a, 0x87, 0x40,
	0x10, 0xc7, 0x83, 0x7e, 0x99, 0x4d, 0x58, 0xb0, 0x47, 0xe9, 0xd4, 0xa9, 0x7f, 0xee, 0xa1, 0xde,
	0x20, 0x02, 0x2f, 0x1d, 0x22, 0x6f, 0x11, 0xc8, 0x6a, 0x83, 0x2e, 0x95, 0x6b, 0xce, 0xd4, 0x3b,
	0xf4, 0xd6, 0xb1, 0xab, 0x2b, 0x9e, 0xda, 0xbd, 0x39, 0xce, 0x67, 0xbe, 0x1f, 0x66, 0x58, 0xc8,
	0x08, 0xa7, 0x1f, 0xdd, 0xa2, 0x1c, 0x27, 0xc3, 0x46, 0x9c, 0x29, 0x2d, 0x3f, 0x91, 0x15, 0xf7,
	0x7a, 0xe8, 0x48, 0xfa, 0x26, 0xe1, 0x40, 0x66, 0xca, 0x8f, 0x3a, 0xe4, 0x19, 0xcc, 0xe1, 0x43,
	0x93, 0xff, 0x3e, 0x1e, 0x15, 0xb7, 0xfd, 0x52, 0x9c, 0x74, 0xc8, 0xf5, 0x9b, 0x62, 0xb5, 0xd4,
	0xa7, 0x16, 0xdc, 0xfc, 0xb8, 0xfd, 0xdd, 0x41, 0x56, 0xb9, 0xbc, 0x6a, 0x4e, 0x17, 0xaf, 0xb0,
	0x5f, 0x22, 0x8b, 0x0b, 0xf9, 0x9f, 0x5c, 0x96, 0xc8, 0xcf, 0xf8, 0xf5, 0x8d, 0xc4, 0xf9, 0x65,
	0x04, 0x49, 0xa3, 0x19, 0x08, 0xcf, 0xf7, 0x44, 0x0d, 0xbb, 0x47, 0x4d, 0x2c, 0x02, 0x43, 0x96,
	0xf1, 0xf9, 0x57, 0x31, 0xe8, 0x2a, 0x68, 0xe0, 0xe0, 0xc9, 0x1e, 0x40, 0x04, 0xc6, 0x1c, 0xe4,
	0x15, 0xd7, 0x51, 0xec, 0xea, 0xe8, 0xe1, 0xb0, 0x44, 0x7e, 0x50, 0xac, 0xc4, 0x4d, 0x70, 0x79,
	0x8b, 0x79, 0x4f, 0x11, 0x49, 0xaf, 0xa6, 0x77, 0x48, 0xed, 0x7e, 0x4e, 0x55, 0x84, 0xef, 0xb0,
	0x75, 0xc9, 0x58, 0xdc, 0xcb, 0xee, 0xd3, 0x97, 0x64, 0x6e, 0x36, 0x89, 0x7b, 0x1c, 0x77, 0x7f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x0c, 0x84, 0x17, 0x5c, 0x90, 0x02, 0x00, 0x00,
}