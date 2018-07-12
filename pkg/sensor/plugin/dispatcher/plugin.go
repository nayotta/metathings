package main

import (
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/any"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	mt_plugin "github.com/nayotta/metathings/pkg/cored/plugin"
	pb "github.com/nayotta/metathings/pkg/proto/sensor"
)

type sensorDispatcherPlugin struct {
	cli_fty *client_helper.ClientFactory
}

func (dp *sensorDispatcherPlugin) Init(opt opt_helper.Option) error {
	return errors.New("unimplemented")
}

var (
	unary_call_methods = map[string]func(pb.SensorServiceClient, context.Context, *any.Any) (*any.Any, error){}
)

func (dp *sensorDispatcherPlugin) UnaryCall(method string, ctx context.Context, req *any.Any) (*any.Any, error) {
	return nil, errors.New("unimplemented")
}

var (
	stream_call_methods = map[string]func(pb.SensorServiceClient, context.Context, ...func()) (mt_plugin.Stream, error){}
)

func (dp sensorDispatcherPlugin) StreamCall(method string, ctx context.Context) (mt_plugin.Stream, error) {
	return nil, errors.New("unimplemented")
}

func NewDispatcherPlugin() mt_plugin.DispatcherPlugin {
	return &sensorDispatcherPlugin{}
}
