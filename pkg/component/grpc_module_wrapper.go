package metathings_component

import (
	"context"
	"errors"
	"reflect"

	"github.com/golang/protobuf/ptypes/any"
	log "github.com/sirupsen/logrus"

	pb "github.com/nayotta/metathings/pkg/proto/component"
)

var (
	ErrHandleUnimplemented = errors.New("handle unimplemented")
)

type GrpcModuleWrapper struct {
	logger              log.FieldLogger
	target              interface{}
	unary_method_cache  map[string]func(context.Context, *any.Any) (*any.Any, error)
	stream_method_cache map[string]func(pb.ModuleService_StreamCallServer) error
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

func (self *GrpcModuleWrapper) lookup_stream_method(meth string) (func(pb.ModuleService_StreamCallServer) error, error) {
	fn, ok := self.stream_method_cache[meth]
	if ok {
		return fn, nil
	}

	tgr := reflect.ValueOf(self.target)
	ref_fn := tgr.MethodByName("HANDLE_GRPC_" + meth)
	if ref_fn.Kind() != reflect.Func {
		return nil, ErrHandleUnimplemented
	}

	fn, ok = ref_fn.Interface().(func(pb.ModuleService_StreamCallServer) error)
	if !ok {
		return nil, ErrHandleUnimplemented
	}

	return fn, nil
}

func (self *GrpcModuleWrapper) UnaryCall(ctx context.Context, req *pb.UnaryCallRequest) (*pb.UnaryCallResponse, error) {
	meth := req.GetMethod().GetValue()
	fn, err := self.lookup_unary_method(meth)
	if err != nil {
		self.logger.WithError(err).WithField("method", meth).Errorf("failed to lookup unary method")
		return nil, err
	}
	self.logger.WithField("method", meth).Debugf("lookup unary method")

	any_res, err := fn(ctx, req.Value)
	if err != nil {
		self.logger.WithError(err).Errorf("unary call error")
		return nil, err
	}

	res := &pb.UnaryCallResponse{
		Method: meth,
		Value:  any_res,
	}

	self.logger.Debugf("unary call success")

	return res, nil
}

func (self *GrpcModuleWrapper) StreamCall(upstm pb.ModuleService_StreamCallServer) error {
	var req *pb.StreamCallRequest
	var err error

	if req, err = upstm.Recv(); err != nil {
		self.logger.WithError(err).Errorf("failed to receive config message")
		return err
	}
	self.logger.Debugf("receive config message")

	cfg := req.GetConfig()
	if cfg == nil {
		err = ErrInvalidArguments
		self.logger.WithError(err).Errorf("faield to get config from message")
		return err
	}

	meth := cfg.GetMethod().GetValue()
	fn, err := self.lookup_stream_method(meth)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to lookup stream method")
		return err
	}
	self.logger.WithField("method", meth).Debugf("lookup stream method")

	err = fn(upstm)
	if err != nil {
		self.logger.WithError(err).Errorf("stream call error")
		return err
	}

	self.logger.Debugf("stream call success")

	return nil
}

func NewGrpcModuleWrapper(target interface{}, logger log.FieldLogger) *GrpcModuleWrapper {
	return &GrpcModuleWrapper{
		logger: logger.WithFields(log.Fields{
			"#driver": "grpc",
			"#module": "module-wrapper",
		}),
		target:              target,
		unary_method_cache:  make(map[string]func(context.Context, *any.Any) (*any.Any, error)),
		stream_method_cache: make(map[string]func(pb.ModuleService_StreamCallServer) error),
	}
}
