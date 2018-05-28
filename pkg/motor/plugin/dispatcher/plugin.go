package main

import (
	"context"

	"github.com/golang/protobuf/ptypes/any"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	mt_plugin "github.com/nayotta/metathings/pkg/core/plugin"
)

type motorDispatcherPlugin struct{}

func (dp *motorDispatcherPlugin) Init(opt opt_helper.Option) error {
	return nil
}

func (dp *motorDispatcherPlugin) UnaryCall(method string, ctx context.Context, req *any.Any) (*any.Any, error) {
	return nil, nil
}

func (dp *motorDispatcherPlugin) StreamCall(method string, ctx context.Context) (mt_plugin.Stream, error) {
	return nil, nil
}

func NewDispatcherPlugin() mt_plugin.DispatcherPlugin {
	return &motorDispatcherPlugin{}
}
