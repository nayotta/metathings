package main

import (
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	mtp "github.com/nayotta/metathings/pkg/core/plugin"
)

type cameraServicePlugin struct{}

func (p *cameraServicePlugin) Run() error {
	return nil
}

func (p *cameraServicePlugin) Init(opts opt_helper.Option) error {
	return nil
}

func NewServicePlugin() mtp.ServicePlugin {
	return &cameraServicePlugin{}
}
