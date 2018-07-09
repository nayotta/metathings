package main

import (
	"context"

	"github.com/golang/protobuf/ptypes/any"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	mt_plugin "github.com/nayotta/metathings/pkg/cored/plugin"
	pb "github.com/nayotta/metathings/pkg/proto/servo"
)

type servoDispatcherPlugin struct {
	cli_fty *client_helper.ClientFactory
}

func (dp *servoDispatcherPlugin) Init(opt opt_helper.Option) error {
	cfgs := client_helper.NewDefaultServiceConfigs(opt.GetString("endpoint"))
	cli_fty, err := client_helper.NewClientFactory(cfgs, client_helper.WithInsecureOptionFunc())

	if err != nil {
		return err
	}

	dp.cli_fty = cli_fty

	return nil
}

var (
	unary_call_methods = map[string]func(pb.ServoServiceClient, context.Context, *any.Any) (*any.Any, error){}
)

func (dp *servoDispatcherPlugin) UnaryCall(method string, ctx context.Context, req *any.Any) (*any.Any, error) {
	cli, cfn, err := dp.cli_fty.NewServoServiceClient()
	if err != nil {
		return nil, err
	}
	defer cfn()

	fn, ok := unary_call_methods[method]
	if !ok {
		return nil, mt_plugin.ErrUnknownMethod
	}

	return fn(cli, ctx, req)
}

var (
	stream_call_methods = map[string]func(cli pb.ServoServiceClient, ctx context.Context, cbs ...func()) (mt_plugin.Stream, error){}
)

func (dp *servoDispatcherPlugin) StreamCall(method string, ctx context.Context) (mt_plugin.Stream, error) {
	cli, cfn, err := dp.cli_fty.NewServoServiceClient()
	if err != nil {
		return nil, err
	}

	fn, ok := stream_call_methods[method]
	if !ok {
		return nil, mt_plugin.ErrUnknownMethod
	}

	return fn(cli, ctx, cfn)
}

func NewDispatcherPlugin() mt_plugin.DispatcherPlugin {
	return &servoDispatcherPlugin{}
}
