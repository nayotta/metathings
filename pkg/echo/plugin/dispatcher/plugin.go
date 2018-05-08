package main

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"

	client_helper "github.com/bigdatagz/metathings/pkg/common/client"
	opt_helper "github.com/bigdatagz/metathings/pkg/common/option"
	mt_plugin "github.com/bigdatagz/metathings/pkg/core/plugin"
	pb "github.com/bigdatagz/metathings/pkg/proto/echo"
)

type echoDispatcherPlugin struct {
	cli_fty *client_helper.ClientFactory
}

func (dp *echoDispatcherPlugin) Init(opts opt_helper.Option) error {
	cfgs := client_helper.NewDefaultServiceConfigs(opts.GetString("endpoint"))
	cli_fty, err := client_helper.NewClientFactory(cfgs, client_helper.WithInsecureOptionFunc())
	if err != nil {
		return err
	}
	dp.cli_fty = cli_fty

	return nil
}

func (dp *echoDispatcherPlugin) UnaryCall(method string, ctx context.Context, req *any.Any) (*any.Any, error) {
	switch method {
	case "Echo":
		cli, fn, err := dp.cli_fty.NewEchoServiceClient()
		if err != nil {
			return nil, err
		}
		defer fn()

		req1 := new(pb.EchoRequest)
		err = ptypes.UnmarshalAny(req, req1)
		if err != nil {
			return nil, err
		}

		res, err := cli.Echo(ctx, req1)
		if err != nil {
			return nil, err
		}

		res1, err := ptypes.MarshalAny(res)
		if err != nil {
			return nil, err
		}

		return res1, nil
	}
	return nil, nil
}

func NewDispatcherPlugin() mt_plugin.DispatcherPlugin {
	return &echoDispatcherPlugin{}
}
