// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package evaluatord

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 660 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0x6d, 0x2e, 0xbf, 0x1f, 0x9a, 0x96, 0x46, 0x59, 0x68, 0x0e, 0x46, 0x50, 0x54, 0xee, 0x2e,
	0xf4, 0x0f, 0x94, 0x22, 0x5a, 0x9a, 0xb4, 0xf4, 0xc2, 0x01, 0x9a, 0x9c, 0x90, 0x50, 0xe4, 0xc4,
	0x13, 0xd7, 0x22, 0xce, 0x26, 0xde, 0x4d, 0x24, 0x0b, 0x24, 0x4e, 0x48, 0x9c, 0xb8, 0x72, 0xe2,
	0x13, 0xf2, 0x25, 0xd0, 0xee, 0x7a, 0xd7, 0x76, 0x12, 0xb4, 0x76, 0x8e, 0xdd, 0x79, 0x6f, 0xde,
	0xdb, 0xb7, 0xe3, 0x69, 0xe0, 0x2e, 0xc3, 0x78, 0x1e, 0x0e, 0xd0, 0x9d, 0xc4, 0x94, 0x53, 0xf2,
	0xd8, 0x0b, 0xdd, 0x08, 0xb9, 0xc7, 0x6f, 0xc3, 0x71, 0xc0, 0x5c, 0x5d, 0xc4, 0xb9, 0x37, 0x9a,
	0x79, 0x9c, 0xc6, 0xbe, 0xf3, 0x20, 0xa0, 0x34, 0x18, 0xe1, 0xbe, 0xc4, 0xf7, 0x67, 0xc3, 0x7d,
	0x8c, 0x26, 0x3c, 0x51, 0x74, 0xa7, 0x39, 0x88, 0xd1, 0xe3, 0xd8, 0x33, 0x78, 0x7d, 0xee, 0xe3,
	0x08, 0x57, 0x9c, 0xef, 0x4c, 0x3c, 0x3e, 0xb8, 0x5d, 0x3a, 0xbe, 0x17, 0x20, 0x5f, 0xc6, 0x8e,
	0x42, 0x96, 0x3b, 0x65, 0xe9, 0xf1, 0x23, 0xcf, 0xf7, 0x7b, 0x8c, 0xce, 0xe2, 0x01, 0xb2, 0x1e,
	0xa7, 0x4b, 0xb4, 0x27, 0x31, 0x46, 0x74, 0x8e, 0x06, 0x32, 0x8c, 0x69, 0xb4, 0x04, 0xda, 0x5d,
	0xe8, 0xdd, 0xeb, 0x27, 0x29, 0x21, 0x05, 0x38, 0x12, 0xc0, 0x3d, 0xf6, 0x79, 0xb9, 0xb6, 0x2d,
	0xdc, 0x8a, 0x52, 0xfa, 0xf7, 0xc3, 0xe9, 0x0c, 0xe3, 0xa4, 0xc7, 0x38, 0x8d, 0xbd, 0x00, 0x05,
	0xdc, 0xc7, 0x2c, 0x62, 0x87, 0xa4, 0x19, 0xf1, 0x30, 0x42, 0xad, 0x4f, 0xd2, 0x7c, 0xf2, 0x67,
	0x0d, 0x95, 0x4d, 0xfe, 0xa8, 0x2e, 0x95, 0xf2, 0x18, 0x65, 0x4b, 0x9c, 0xe8, 0x3c, 0x1c, 0x91,
	0xc7, 0x80, 0x8e, 0x87, 0x61, 0x20, 0xf3, 0xc8, 0xc3, 0x77, 0xd3, 0x2c, 0x74, 0x59, 0x66, 0x91,
	0x03, 0x1c, 0xfc, 0x69, 0x40, 0xe3, 0xca, 0xbc, 0x75, 0x47, 0xbd, 0x3e, 0xf9, 0x51, 0x83, 0x7a,
	0x5b, 0x9a, 0x36, 0x35, 0x72, 0xe2, 0xda, 0x26, 0xc5, 0x5d, 0xa0, 0xdc, 0xe0, 0x74, 0x86, 0x8c,
	0x3b, 0x2f, 0xd7, 0x60, 0xb2, 0x09, 0x1d, 0x33, 0xdc, 0xdb, 0x20, 0x08, 0xf5, 0x4b, 0x19, 0x55,
	0x25, 0x27, 0x0b, 0x14, 0xed, 0xa4, 0xe9, 0xaa, 0x59, 0x76, 0xf5, 0x2c, 0xbb, 0x57, 0x62, 0x96,
	0xf7, 0x36, 0xc8, 0xf7, 0x1a, 0x6c, 0xbf, 0x17, 0xf1, 0x67, 0x32, 0x2f, 0xec, 0x32, 0x45, 0x86,
	0x56, 0x39, 0xa9, 0x4e, 0x34, 0xd7, 0xfd, 0x06, 0x5b, 0xd7, 0xc8, 0x33, 0x13, 0xc7, 0xf6, 0x5e,
	0x79, 0xbc, 0xb6, 0xf0, 0xbc, 0x2a, 0xcd, 0x18, 0x10, 0x41, 0xbc, 0x0b, 0x59, 0x56, 0x63, 0x65,
	0x82, 0x28, 0x32, 0x2a, 0x04, 0xb1, 0x48, 0x34, 0x3e, 0xa6, 0xb0, 0x73, 0xe1, 0xfb, 0x1d, 0xf5,
	0x0d, 0x77, 0x69, 0x96, 0xc8, 0x99, 0xbd, 0xe9, 0x4a, 0xa2, 0x7d, 0x06, 0xbe, 0x80, 0x73, 0x23,
	0x3f, 0x97, 0x94, 0xfc, 0x36, 0xa6, 0x51, 0xa6, 0xdb, 0xb6, 0xeb, 0xfe, 0x9b, 0x6d, 0x17, 0xff,
	0x5d, 0x83, 0x66, 0x31, 0x8c, 0x56, 0xa2, 0x3a, 0x91, 0xf3, 0xaa, 0x31, 0x6a, 0xa6, 0x56, 0x7d,
	0xb3, 0x7e, 0x03, 0xf3, 0x1e, 0x3f, 0x6b, 0xd0, 0x10, 0xa0, 0xae, 0xd8, 0x88, 0xc6, 0xda, 0x69,
	0xb9, 0xce, 0x05, 0x92, 0x76, 0xf5, 0x6a, 0x2d, 0xae, 0x31, 0x34, 0x81, 0xff, 0xaf, 0x51, 0x56,
	0xc9, 0xd3, 0x52, 0xd3, 0x2e, 0xa0, 0x5a, 0xfb, 0x59, 0x05, 0x86, 0x51, 0xfc, 0x55, 0x83, 0xfb,
	0x1f, 0xc4, 0xa6, 0xef, 0xa8, 0x45, 0xdf, 0x4a, 0x2e, 0xe5, 0x9a, 0x27, 0xaf, 0xed, 0xdd, 0x56,
	0xf1, 0xb4, 0x99, 0xb3, 0x75, 0xe9, 0xc6, 0xd9, 0x57, 0xd8, 0x54, 0x1b, 0xb4, 0x2b, 0x56, 0x3b,
	0x39, 0x2a, 0xbb, 0x70, 0x25, 0x5c, 0xdb, 0x38, 0xae, 0xc8, 0x32, 0xea, 0x9f, 0x60, 0x53, 0xed,
	0xdb, 0xd2, 0xea, 0x39, 0xb8, 0xfd, 0xcb, 0x48, 0x00, 0xe4, 0xba, 0x54, 0xdd, 0x0f, 0x4b, 0x2e,
	0xd7, 0x42, 0xf3, 0xa3, 0x6a, 0x24, 0x73, 0x33, 0x06, 0x77, 0xc4, 0x18, 0x48, 0xe1, 0x92, 0x23,
	0x93, 0x97, 0x3d, 0xa8, 0x42, 0x31, 0xa2, 0x09, 0x80, 0x9c, 0x7b, 0xf9, 0x3f, 0xbe, 0xcc, 0x7d,
	0x33, 0x74, 0x85, 0xfb, 0xe6, 0x49, 0x46, 0x3a, 0x84, 0xc6, 0x85, 0xef, 0xb7, 0xd5, 0x8f, 0x85,
	0x2e, 0x55, 0x17, 0x3f, 0x2d, 0xb5, 0x70, 0x8b, 0x24, 0xfb, 0xab, 0x32, 0x68, 0xaa, 0x75, 0x99,
	0x12, 0xc5, 0xba, 0x54, 0x7a, 0xe7, 0x65, 0x17, 0xed, 0x22, 0xd3, 0x2a, 0xda, 0xda, 0xfa, 0x08,
	0x59, 0x97, 0xfe, 0x7f, 0xb2, 0x7e, 0xf8, 0x37, 0x00, 0x00, 0xff, 0xff, 0x2a, 0xc8, 0x49, 0x53,
	0x12, 0x0b, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EvaluatordServiceClient is the client API for EvaluatordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EvaluatordServiceClient interface {
	CreateEvaluator(ctx context.Context, in *CreateEvaluatorRequest, opts ...grpc.CallOption) (*CreateEvaluatorResponse, error)
	DeleteEvaluator(ctx context.Context, in *DeleteEvaluatorRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	PatchEvaluator(ctx context.Context, in *PatchEvaluatorRequest, opts ...grpc.CallOption) (*PatchEvaluatorResponse, error)
	GetEvaluator(ctx context.Context, in *GetEvaluatorRequest, opts ...grpc.CallOption) (*GetEvaluatorResponse, error)
	ListEvaluators(ctx context.Context, in *ListEvaluatorsRequest, opts ...grpc.CallOption) (*ListEvaluatorsResponse, error)
	AddSourcesToEvaluator(ctx context.Context, in *AddSourcesToEvaluatorRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	RemoveSourcesFromEvaluator(ctx context.Context, in *RemoveSourcesFromEvaluatorRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ListEvaluatorsBySource(ctx context.Context, in *ListEvaluatorsBySourceRequest, opts ...grpc.CallOption) (*ListEvaluatorsBySourceResponse, error)
	// Task
	ListTasksBySource(ctx context.Context, in *ListTasksBySourceRequest, opts ...grpc.CallOption) (*ListTasksBySourceResponse, error)
	GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*GetTaskResponse, error)
	// Storage
	QueryStorageByDevice(ctx context.Context, in *QueryStorageByDeviceRequest, opts ...grpc.CallOption) (*QueryStorageByDeviceResponse, error)
	// Timer
	CreateTimer(ctx context.Context, in *CreateTimerRequest, opts ...grpc.CallOption) (*CreateTimerResponse, error)
	DeleteTimer(ctx context.Context, in *DeleteTimerRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	PatchTimer(ctx context.Context, in *PatchTimerRequest, opts ...grpc.CallOption) (*PatchTimerResponse, error)
	GetTimer(ctx context.Context, in *GetTimerRequest, opts ...grpc.CallOption) (*GetTimerResponse, error)
	ListTimers(ctx context.Context, in *ListTimersRequest, opts ...grpc.CallOption) (*ListTimersResponse, error)
	AddConfigsToTimer(ctx context.Context, in *AddConfigsToTimerRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	RemoveConfigsFromTimer(ctx context.Context, in *RemoveConfigsFromTimerRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type evaluatordServiceClient struct {
	cc *grpc.ClientConn
}

func NewEvaluatordServiceClient(cc *grpc.ClientConn) EvaluatordServiceClient {
	return &evaluatordServiceClient{cc}
}

func (c *evaluatordServiceClient) CreateEvaluator(ctx context.Context, in *CreateEvaluatorRequest, opts ...grpc.CallOption) (*CreateEvaluatorResponse, error) {
	out := new(CreateEvaluatorResponse)
	err := c.cc.Invoke(ctx, "/ai.metathings.service.evaluatord.EvaluatordService/CreateEvaluator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evaluatordServiceClient) DeleteEvaluator(ctx context.Context, in *DeleteEvaluatorRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ai.metathings.service.evaluatord.EvaluatordService/DeleteEvaluator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evaluatordServiceClient) PatchEvaluator(ctx context.Context, in *PatchEvaluatorRequest, opts ...grpc.CallOption) (*PatchEvaluatorResponse, error) {
	out := new(PatchEvaluatorResponse)
	err := c.cc.Invoke(ctx, "/ai.metathings.service.evaluatord.EvaluatordService/PatchEvaluator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evaluatordServiceClient) GetEvaluator(ctx context.Context, in *GetEvaluatorRequest, opts ...grpc.CallOption) (*GetEvaluatorResponse, error) {
	out := new(GetEvaluatorResponse)
	err := c.cc.Invoke(ctx, "/ai.metathings.service.evaluatord.EvaluatordService/GetEvaluator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evaluatordServiceClient) ListEvaluators(ctx context.Context, in *ListEvaluatorsRequest, opts ...grpc.CallOption) (*ListEvaluatorsResponse, error) {
	out := new(ListEvaluatorsResponse)
	err := c.cc.Invoke(ctx, "/ai.metathings.service.evaluatord.EvaluatordService/ListEvaluators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evaluatordServiceClient) AddSourcesToEvaluator(ctx context.Context, in *AddSourcesToEvaluatorRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ai.metathings.service.evaluatord.EvaluatordService/AddSourcesToEvaluator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evaluatordServiceClient) RemoveSourcesFromEvaluator(ctx context.Context, in *RemoveSourcesFromEvaluatorRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ai.metathings.service.evaluatord.EvaluatordService/RemoveSourcesFromEvaluator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evaluatordServiceClient) ListEvaluatorsBySource(ctx context.Context, in *ListEvaluatorsBySourceRequest, opts ...grpc.CallOption) (*ListEvaluatorsBySourceResponse, error) {
	out := new(ListEvaluatorsBySourceResponse)
	err := c.cc.Invoke(ctx, "/ai.metathings.service.evaluatord.EvaluatordService/ListEvaluatorsBySource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evaluatordServiceClient) ListTasksBySource(ctx context.Context, in *ListTasksBySourceRequest, opts ...grpc.CallOption) (*ListTasksBySourceResponse, error) {
	out := new(ListTasksBySourceResponse)
	err := c.cc.Invoke(ctx, "/ai.metathings.service.evaluatord.EvaluatordService/ListTasksBySource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evaluatordServiceClient) GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*GetTaskResponse, error) {
	out := new(GetTaskResponse)
	err := c.cc.Invoke(ctx, "/ai.metathings.service.evaluatord.EvaluatordService/GetTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evaluatordServiceClient) QueryStorageByDevice(ctx context.Context, in *QueryStorageByDeviceRequest, opts ...grpc.CallOption) (*QueryStorageByDeviceResponse, error) {
	out := new(QueryStorageByDeviceResponse)
	err := c.cc.Invoke(ctx, "/ai.metathings.service.evaluatord.EvaluatordService/QueryStorageByDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evaluatordServiceClient) CreateTimer(ctx context.Context, in *CreateTimerRequest, opts ...grpc.CallOption) (*CreateTimerResponse, error) {
	out := new(CreateTimerResponse)
	err := c.cc.Invoke(ctx, "/ai.metathings.service.evaluatord.EvaluatordService/CreateTimer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evaluatordServiceClient) DeleteTimer(ctx context.Context, in *DeleteTimerRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ai.metathings.service.evaluatord.EvaluatordService/DeleteTimer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evaluatordServiceClient) PatchTimer(ctx context.Context, in *PatchTimerRequest, opts ...grpc.CallOption) (*PatchTimerResponse, error) {
	out := new(PatchTimerResponse)
	err := c.cc.Invoke(ctx, "/ai.metathings.service.evaluatord.EvaluatordService/PatchTimer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evaluatordServiceClient) GetTimer(ctx context.Context, in *GetTimerRequest, opts ...grpc.CallOption) (*GetTimerResponse, error) {
	out := new(GetTimerResponse)
	err := c.cc.Invoke(ctx, "/ai.metathings.service.evaluatord.EvaluatordService/GetTimer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evaluatordServiceClient) ListTimers(ctx context.Context, in *ListTimersRequest, opts ...grpc.CallOption) (*ListTimersResponse, error) {
	out := new(ListTimersResponse)
	err := c.cc.Invoke(ctx, "/ai.metathings.service.evaluatord.EvaluatordService/ListTimers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evaluatordServiceClient) AddConfigsToTimer(ctx context.Context, in *AddConfigsToTimerRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ai.metathings.service.evaluatord.EvaluatordService/AddConfigsToTimer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evaluatordServiceClient) RemoveConfigsFromTimer(ctx context.Context, in *RemoveConfigsFromTimerRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ai.metathings.service.evaluatord.EvaluatordService/RemoveConfigsFromTimer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EvaluatordServiceServer is the server API for EvaluatordService service.
type EvaluatordServiceServer interface {
	CreateEvaluator(context.Context, *CreateEvaluatorRequest) (*CreateEvaluatorResponse, error)
	DeleteEvaluator(context.Context, *DeleteEvaluatorRequest) (*empty.Empty, error)
	PatchEvaluator(context.Context, *PatchEvaluatorRequest) (*PatchEvaluatorResponse, error)
	GetEvaluator(context.Context, *GetEvaluatorRequest) (*GetEvaluatorResponse, error)
	ListEvaluators(context.Context, *ListEvaluatorsRequest) (*ListEvaluatorsResponse, error)
	AddSourcesToEvaluator(context.Context, *AddSourcesToEvaluatorRequest) (*empty.Empty, error)
	RemoveSourcesFromEvaluator(context.Context, *RemoveSourcesFromEvaluatorRequest) (*empty.Empty, error)
	ListEvaluatorsBySource(context.Context, *ListEvaluatorsBySourceRequest) (*ListEvaluatorsBySourceResponse, error)
	// Task
	ListTasksBySource(context.Context, *ListTasksBySourceRequest) (*ListTasksBySourceResponse, error)
	GetTask(context.Context, *GetTaskRequest) (*GetTaskResponse, error)
	// Storage
	QueryStorageByDevice(context.Context, *QueryStorageByDeviceRequest) (*QueryStorageByDeviceResponse, error)
	// Timer
	CreateTimer(context.Context, *CreateTimerRequest) (*CreateTimerResponse, error)
	DeleteTimer(context.Context, *DeleteTimerRequest) (*empty.Empty, error)
	PatchTimer(context.Context, *PatchTimerRequest) (*PatchTimerResponse, error)
	GetTimer(context.Context, *GetTimerRequest) (*GetTimerResponse, error)
	ListTimers(context.Context, *ListTimersRequest) (*ListTimersResponse, error)
	AddConfigsToTimer(context.Context, *AddConfigsToTimerRequest) (*empty.Empty, error)
	RemoveConfigsFromTimer(context.Context, *RemoveConfigsFromTimerRequest) (*empty.Empty, error)
}

func RegisterEvaluatordServiceServer(s *grpc.Server, srv EvaluatordServiceServer) {
	s.RegisterService(&_EvaluatordService_serviceDesc, srv)
}

func _EvaluatordService_CreateEvaluator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEvaluatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluatordServiceServer).CreateEvaluator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.evaluatord.EvaluatordService/CreateEvaluator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluatordServiceServer).CreateEvaluator(ctx, req.(*CreateEvaluatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvaluatordService_DeleteEvaluator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEvaluatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluatordServiceServer).DeleteEvaluator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.evaluatord.EvaluatordService/DeleteEvaluator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluatordServiceServer).DeleteEvaluator(ctx, req.(*DeleteEvaluatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvaluatordService_PatchEvaluator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchEvaluatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluatordServiceServer).PatchEvaluator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.evaluatord.EvaluatordService/PatchEvaluator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluatordServiceServer).PatchEvaluator(ctx, req.(*PatchEvaluatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvaluatordService_GetEvaluator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEvaluatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluatordServiceServer).GetEvaluator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.evaluatord.EvaluatordService/GetEvaluator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluatordServiceServer).GetEvaluator(ctx, req.(*GetEvaluatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvaluatordService_ListEvaluators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEvaluatorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluatordServiceServer).ListEvaluators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.evaluatord.EvaluatordService/ListEvaluators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluatordServiceServer).ListEvaluators(ctx, req.(*ListEvaluatorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvaluatordService_AddSourcesToEvaluator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSourcesToEvaluatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluatordServiceServer).AddSourcesToEvaluator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.evaluatord.EvaluatordService/AddSourcesToEvaluator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluatordServiceServer).AddSourcesToEvaluator(ctx, req.(*AddSourcesToEvaluatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvaluatordService_RemoveSourcesFromEvaluator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveSourcesFromEvaluatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluatordServiceServer).RemoveSourcesFromEvaluator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.evaluatord.EvaluatordService/RemoveSourcesFromEvaluator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluatordServiceServer).RemoveSourcesFromEvaluator(ctx, req.(*RemoveSourcesFromEvaluatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvaluatordService_ListEvaluatorsBySource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEvaluatorsBySourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluatordServiceServer).ListEvaluatorsBySource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.evaluatord.EvaluatordService/ListEvaluatorsBySource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluatordServiceServer).ListEvaluatorsBySource(ctx, req.(*ListEvaluatorsBySourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvaluatordService_ListTasksBySource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTasksBySourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluatordServiceServer).ListTasksBySource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.evaluatord.EvaluatordService/ListTasksBySource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluatordServiceServer).ListTasksBySource(ctx, req.(*ListTasksBySourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvaluatordService_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluatordServiceServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.evaluatord.EvaluatordService/GetTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluatordServiceServer).GetTask(ctx, req.(*GetTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvaluatordService_QueryStorageByDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStorageByDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluatordServiceServer).QueryStorageByDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.evaluatord.EvaluatordService/QueryStorageByDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluatordServiceServer).QueryStorageByDevice(ctx, req.(*QueryStorageByDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvaluatordService_CreateTimer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTimerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluatordServiceServer).CreateTimer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.evaluatord.EvaluatordService/CreateTimer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluatordServiceServer).CreateTimer(ctx, req.(*CreateTimerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvaluatordService_DeleteTimer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTimerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluatordServiceServer).DeleteTimer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.evaluatord.EvaluatordService/DeleteTimer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluatordServiceServer).DeleteTimer(ctx, req.(*DeleteTimerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvaluatordService_PatchTimer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchTimerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluatordServiceServer).PatchTimer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.evaluatord.EvaluatordService/PatchTimer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluatordServiceServer).PatchTimer(ctx, req.(*PatchTimerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvaluatordService_GetTimer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTimerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluatordServiceServer).GetTimer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.evaluatord.EvaluatordService/GetTimer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluatordServiceServer).GetTimer(ctx, req.(*GetTimerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvaluatordService_ListTimers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTimersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluatordServiceServer).ListTimers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.evaluatord.EvaluatordService/ListTimers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluatordServiceServer).ListTimers(ctx, req.(*ListTimersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvaluatordService_AddConfigsToTimer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddConfigsToTimerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluatordServiceServer).AddConfigsToTimer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.evaluatord.EvaluatordService/AddConfigsToTimer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluatordServiceServer).AddConfigsToTimer(ctx, req.(*AddConfigsToTimerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvaluatordService_RemoveConfigsFromTimer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveConfigsFromTimerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluatordServiceServer).RemoveConfigsFromTimer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.service.evaluatord.EvaluatordService/RemoveConfigsFromTimer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluatordServiceServer).RemoveConfigsFromTimer(ctx, req.(*RemoveConfigsFromTimerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EvaluatordService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ai.metathings.service.evaluatord.EvaluatordService",
	HandlerType: (*EvaluatordServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEvaluator",
			Handler:    _EvaluatordService_CreateEvaluator_Handler,
		},
		{
			MethodName: "DeleteEvaluator",
			Handler:    _EvaluatordService_DeleteEvaluator_Handler,
		},
		{
			MethodName: "PatchEvaluator",
			Handler:    _EvaluatordService_PatchEvaluator_Handler,
		},
		{
			MethodName: "GetEvaluator",
			Handler:    _EvaluatordService_GetEvaluator_Handler,
		},
		{
			MethodName: "ListEvaluators",
			Handler:    _EvaluatordService_ListEvaluators_Handler,
		},
		{
			MethodName: "AddSourcesToEvaluator",
			Handler:    _EvaluatordService_AddSourcesToEvaluator_Handler,
		},
		{
			MethodName: "RemoveSourcesFromEvaluator",
			Handler:    _EvaluatordService_RemoveSourcesFromEvaluator_Handler,
		},
		{
			MethodName: "ListEvaluatorsBySource",
			Handler:    _EvaluatordService_ListEvaluatorsBySource_Handler,
		},
		{
			MethodName: "ListTasksBySource",
			Handler:    _EvaluatordService_ListTasksBySource_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _EvaluatordService_GetTask_Handler,
		},
		{
			MethodName: "QueryStorageByDevice",
			Handler:    _EvaluatordService_QueryStorageByDevice_Handler,
		},
		{
			MethodName: "CreateTimer",
			Handler:    _EvaluatordService_CreateTimer_Handler,
		},
		{
			MethodName: "DeleteTimer",
			Handler:    _EvaluatordService_DeleteTimer_Handler,
		},
		{
			MethodName: "PatchTimer",
			Handler:    _EvaluatordService_PatchTimer_Handler,
		},
		{
			MethodName: "GetTimer",
			Handler:    _EvaluatordService_GetTimer_Handler,
		},
		{
			MethodName: "ListTimers",
			Handler:    _EvaluatordService_ListTimers_Handler,
		},
		{
			MethodName: "AddConfigsToTimer",
			Handler:    _EvaluatordService_AddConfigsToTimer_Handler,
		},
		{
			MethodName: "RemoveConfigsFromTimer",
			Handler:    _EvaluatordService_RemoveConfigsFromTimer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}