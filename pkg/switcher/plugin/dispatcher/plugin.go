package main

import (
	"context"

	"github.com/golang/protobuf/ptypes/any"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	mt_plugin "github.com/nayotta/metathings/pkg/cored/plugin"
	pb "github.com/nayotta/metathings/pkg/proto/switcher"
)

type switcherDispatcherPlugin struct {
	cli_fty *client_helper.ClientFactory
}

func (dp *switcherDispatcherPlugin) Init(opt opt_helper.Option) error {
	cfgs := client_helper.NewDefaultServiceConfigs(opt.GetString("endpoint"))
	cli_fty, err := client_helper.NewClientFactory(cfgs, client_helper.WithInsecureOptionFunc())
	if err != nil {
		return err
	}

	dp.cli_fty = cli_fty

	return nil
}

var (
	unary_call_methods = map[string]func(pb.SwitcherServiceClient, context.Context, *any.Any) (*any.Any, error){}
)

func (dp *switcherDispatcherPlugin) UnaryCall(method string, ctx context.Context, req *any.Any) (*any.Any, error) {
	cli, cfn, err := dp.cli_fty.NewSwitcherServiceClient()
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

func (dp *switcherDispatcherPlugin) StreamCall(method string, ctx context.Context) (mt_plugin.Stream, error) {
	return nil, mt_plugin.ErrUnknownMethod
}

func NewDispatcherPlugin() mt_plugin.DispatcherPlugin {
	return &switcherDispatcherPlugin{}
}
