package main

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/empty"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	mt_plugin "github.com/nayotta/metathings/pkg/core/plugin"
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

func (dp *switcherDispatcherPlugin) UnaryCall(method string, ctx context.Context, req *any.Any) (*any.Any, error) {
	cli, cfn, err := dp.cli_fty.NewSwitcherServiceClient()
	if err != nil {
		return nil, err
	}
	defer cfn()

	switch method {
	case "Get":
		res, err := cli.Get(ctx, &empty.Empty{})
		if err != nil {
			return nil, err
		}

		res1, err := ptypes.MarshalAny(res)
		if err != nil {
			return nil, err
		}

		return res1, nil
	case "Turn":
		req1 := &pb.TurnRequest{}
		err = ptypes.UnmarshalAny(req, req1)
		if err != nil {
			return nil, err
		}

		res, err := cli.Turn(ctx, req1)
		if err != nil {
			return nil, err
		}

		res1, err := ptypes.MarshalAny(res)
		if err != nil {
			return nil, err
		}

		return res1, nil
	default:
		return nil, mt_plugin.ErrUnknownMethod
	}
}

func NewDispatcherPlugin() mt_plugin.DispatcherPlugin {
	return &switcherDispatcherPlugin{}
}
