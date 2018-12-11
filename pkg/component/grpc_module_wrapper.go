package metathings_component

import (
	"context"
	"errors"
	"reflect"

	"github.com/golang/protobuf/ptypes/any"
	"google.golang.org/grpc"

	pb "github.com/nayotta/metathings/pkg/proto/component"
)

var (
	ErrHandleUnimplemented = errors.New("handle unimplemented")
)

type GrpcModuleStream interface {
	Send(*any.Any) error
	Recv() (*any.Any, error)
	grpc.ServerStream
}

type GrpcModuleWrapper struct {
	target              interface{}
	unary_method_cache  map[string]func(context.Context, *any.Any) (*any.Any, error)
	stream_method_cache map[string]func(pb.ModuleService_StreamCallServer) (GrpcModuleStream, error)
}

func (self *GrpcModuleWrapper) lookup_unary_method(meth string) (func(context.Context, *any.Any) (*any.Any, error), error) {
	fn, ok := self.unary_method_cache[meth]
	if ok {
		return fn, nil
	}

	tgr := reflect.ValueOf(self.target)
	ref_fn := tgr.MethodByName("HANDLE_GRPC_" + meth)
	if ref_fn.Kind() != reflect.Func {
		return nil, ErrHandleUnimplemented
	}

	fn, ok = ref_fn.Interface().(func(context.Context, *any.Any) (*any.Any, error))
	if !ok {
		return nil, ErrHandleUnimplemented
	}

	self.unary_method_cache[meth] = fn

	return fn, nil
}

func (self *GrpcModuleWrapper) lookup_stream_method(meth string) (func(pb.ModuleService_StreamCallServer) (GrpcModuleStream, error), error) {
	fn, ok := self.stream_method_cache[meth]
	if ok {
		return fn, nil
	}

	tgr := reflect.ValueOf(self.target)
	ref_fn := tgr.MethodByName("HANDLE_GRPC_" + meth)
	if ref_fn.Kind() != reflect.Func {
		return nil, ErrHandleUnimplemented
	}

	fn, ok = ref_fn.Interface().(func(pb.ModuleService_StreamCallServer) (GrpcModuleStream, error))
	if !ok {
		return nil, ErrHandleUnimplemented
	}

	return fn, nil
}

func (self *GrpcModuleWrapper) UnaryCall(ctx context.Context, req *pb.UnaryCallRequest) (*pb.UnaryCallResponse, error) {
	meth := req.GetMethod().GetValue()
	fn, err := self.lookup_unary_method(meth)
	if err != nil {
		return nil, err
	}

	any_res, err := fn(ctx, req.Value)
	if err != nil {
		return nil, err
	}

	res := &pb.UnaryCallResponse{
		Method: meth,
		Value:  any_res,
	}

	return res, nil
}

func (self *GrpcModuleWrapper) StreamCall(upstm pb.ModuleService_StreamCallServer) error {
	var req *pb.StreamCallRequest
	var err error

	if req, err = upstm.Recv(); err != nil {
		return err
	}

	cfg := req.GetConfig()
	if cfg == nil {
		return ErrInvalidArguments
	}

	meth := cfg.GetMethod().GetValue()

	fn, err := self.lookup_stream_method(meth)
	if err != nil {
		return err
	}

	downstm, err := fn(upstm)
	if err != nil {
		return err
	}

	go self.up2down(upstm, downstm)
	go self.down2up(upstm, downstm)

	return nil
}

func (self *GrpcModuleWrapper) up2down(upstm pb.ModuleService_StreamCallServer, downstm GrpcModuleStream) {
	var upreq *pb.StreamCallRequest
	var downreq *any.Any
	var err error

	for {
		if upreq, err = upstm.Recv(); err != nil {
			return
		}

		downreq = upreq.GetData().GetValue()
		if err = downstm.Send(downreq); err != nil {
			return
		}
	}
}

func (self *GrpcModuleWrapper) down2up(upstm pb.ModuleService_StreamCallServer, downstm GrpcModuleStream) {
	var upres *pb.StreamCallResponse
	var downres *any.Any
	var err error

	for {
		if downres, err = downstm.Recv(); err != nil {
			return
		}

		upres = &pb.StreamCallResponse{
			Response: &pb.StreamCallResponse_Data{
				Data: &pb.StreamCallDataResponse{Value: downres},
			},
		}

		if err = upstm.Send(upres); err != nil {
			return
		}
	}
}

func NewGrpcModuleWrapper(target interface{}) *GrpcModuleWrapper {
	return &GrpcModuleWrapper{
		target:              target,
		unary_method_cache:  make(map[string]func(context.Context, *any.Any) (*any.Any, error)),
		stream_method_cache: make(map[string]func(pb.ModuleService_StreamCallServer) (GrpcModuleStream, error)),
	}
}
