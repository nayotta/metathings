package metathings_component

import (
	"context"
	"errors"
	"reflect"

	"github.com/golang/protobuf/ptypes/any"

	pb "github.com/nayotta/metathings/pkg/proto/component"
)

var (
	ErrHandleUnimplemented = errors.New("handle unimplemented")
)

type GrpcModuleWrapper struct {
	target             interface{}
	unary_method_cache map[string]func(context.Context, *any.Any) (*any.Any, error)
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

func (self *GrpcModuleWrapper) StreamCall(pb.ModuleService_StreamCallServer) error {
	panic("unimplemented")
}

func NewGrpcModuleWrapper(target interface{}) *GrpcModuleWrapper {
	return &GrpcModuleWrapper{
		target:             target,
		unary_method_cache: make(map[string]func(context.Context, *any.Any) (*any.Any, error)),
	}
}