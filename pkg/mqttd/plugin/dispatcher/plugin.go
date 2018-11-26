package main

import (
	"context"

	"github.com/golang/protobuf/ptypes/any"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	mt_plugin "github.com/nayotta/metathings/pkg/cored/plugin"
	pb "github.com/nayotta/metathings/pkg/proto/mqttd"
)

type mqttdDispatcherPlugin struct {
	cliFty *client_helper.ClientFactory
}

func (dp *mqttdDispatcherPlugin) Init(opt opt_helper.Option) error {
	cfgs := client_helper.NewDefaultServiceConfigs(opt.GetString("endpoint"))
	cliFty, err := client_helper.NewClientFactory(cfgs, client_helper.WithInsecureOptionFunc())
	if err != nil {
		return err
	}

	dp.cliFty = cliFty

	return nil
}

var (
	unaryCallMethods = map[string]func(pb.MqttdServiceClient, context.Context, *any.Any) (*any.Any, error){}
)

func (dp *mqttdDispatcherPlugin) UnaryCall(method string, ctx context.Context, req *any.Any) (*any.Any, error) {
	cli, cfn, err := dp.cliFty.NewMqttdServiceClient()
	if err != nil {
		return nil, err
	}
	defer cfn()

	fn, ok := unaryCallMethods[method]
	if !ok {
		return nil, mt_plugin.ErrUnknownMethod
	}
	return fn(cli, ctx, req)
}

func (dp *mqttdDispatcherPlugin) StreamCall(method string, ctx context.Context) (mt_plugin.Stream, error) {
	return nil, mt_plugin.ErrUnknownMethod
}

func NewDispatcherPlugin() mt_plugin.DispatcherPlugin {
	return &mqttdDispatcherPlugin{}
}
