package main

import (
	"context"

	"github.com/golang/protobuf/ptypes/any"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	mt_plugin "github.com/nayotta/metathings/pkg/core/plugin"
	pb "github.com/nayotta/metathings/pkg/proto/motor"
)

type motorDispatcherPlugin struct {
	cli_fty *client_helper.ClientFactory
}

func (dp *motorDispatcherPlugin) Init(opt opt_helper.Option) error {
	cfgs := client_helper.NewDefaultServiceConfigs(opt.GetString("endpoint"))
	cli_fty, err := client_helper.NewClientFactory(cfgs, client_helper.WithInsecureOptionFunc())
	if err != nil {
		return err
	}

	dp.cli_fty = cli_fty

	return nil
}

var (
	unary_call_methods = map[string]func(pb.MotorServiceClient, context.Context, *any.Any) (*any.Any, error){}
)

func (dp *motorDispatcherPlugin) UnaryCall(method string, ctx context.Context, req *any.Any) (*any.Any, error) {
	cli, cfn, err := dp.cli_fty.NewMotorServiceClient()
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
	stream_call_methods = map[string]func(cli pb.MotorServiceClient, ctx context.Context, cbs ...func()) (mt_plugin.Stream, error){}
)

func (dp *motorDispatcherPlugin) StreamCall(method string, ctx context.Context) (mt_plugin.Stream, error) {
	cli, cfn, err := dp.cli_fty.NewMotorServiceClient()
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
	return &motorDispatcherPlugin{}
}
