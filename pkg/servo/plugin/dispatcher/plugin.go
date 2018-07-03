package main

import (
	"context"

	"github.com/golang/protobuf/ptypes/any"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	mt_plugin "github.com/nayotta/metathings/pkg/cored/plugin"
	pb "github.com/nayotta/metathings/pkg/proto/servo"
)

type servoDispatcherPlugin struct {
}

func (dp *servoDispatcherPlugin) Init(opt opt_helper.Option) error {
	return nil
}

var (
	unary_call_methods = map[string]func(pb.ServoServiceClient, context.Context, *any.Any) (*any.Any, error){}
)

func (dp *servoDispatcherPlugin) UnaryCall(method string, ctx context.Context, req *any.Any) (*any.Any, error) {
	return nil, nil
}

var (
	stream_call_methods = map[string]func(cli pb.ServoServiceClient, ctx context.Context, cbs ...func()) (mt_plugin.Stream, error){}
)

func (dp *servoDispatcherPlugin) StreamCall(method string, ctx context.Context) (mt_plugin.Stream, error) {
	return nil, nil
}

func NewDispatcherPlugin() mt_plugin.DispatcherPlugin {
	return &servoDispatcherPlugin{}
}
