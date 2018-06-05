package main

import (
	"context"

	"github.com/golang/protobuf/ptypes/any"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	mt_plugin "github.com/nayotta/metathings/pkg/core/plugin"
)

type cameraDispatcherPlugin struct {
	cli_fty *client_helper.ClientFactory
}

func (dp *cameraDispatcherPlugin) Init(opt opt_helper.Option) error {
	return nil
}

func (dp *cameraDispatcherPlugin) UnaryCall(method string, ctx context.Context, req *any.Any) (*any.Any, error) {
	return nil, nil
}

func (dp *cameraDispatcherPlugin) StreamCall(method string, ctx context.Context) (mt_plugin.Stream, error) {
	return nil, mt_plugin.ErrUnknownMethod
}

func NewDispatcherPlugin() mt_plugin.DispatcherPlugin {
	return &cameraDispatcherPlugin{}
}
