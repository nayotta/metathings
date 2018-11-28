package metathings_component

import (
	"context"
	"errors"
	"reflect"

	"github.com/golang/protobuf/ptypes/any"

	pb "github.com/nayotta/metathings/pkg/proto/component"
)

var (
	ErrProcessUnimplemented = errors.New("process unimplemented")
)

type GrpcModuleWrapper struct {
	target interface{}
}

func (self *GrpcModuleWrapper) UnaryCall(ctx context.Context, req *pb.UnaryCallRequest) (*pb.UnaryCallResponse, error) {
	meth := req.GetMethod().GetValue()
	tgr := reflect.ValueOf(self.target)
	fn := tgr.MethodByName("PROCESS_GRPC_" + meth)
	if fn.Kind() != reflect.Func {
		return nil, ErrProcessUnimplemented
	}

	any_res, err := fn.Interface().(func(context.Context, *any.Any) (*any.Any, error))(ctx, req.Value)
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
	return &GrpcModuleWrapper{target: target}
}
