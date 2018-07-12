package main

import (
	"errors"

	"github.com/nayotta/viper"

	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	mtp "github.com/nayotta/metathings/pkg/cored/plugin"
)

type _sensorDriverOption struct {
	Name string
}

type _rootOptions struct {
	cmd_helper.RootOptions `mapstructure:",squash"`
	Listen                 string
	Endpoint               cmd_helper.EndpointOptions
	Name                   string
	DriverDescriptor       string
}

var (
	root_opts *_rootOptions
	v         *viper.Viper
)

type sensorServicePlugin struct{}

func (p *sensorServicePlugin) Run() error {
	return errors.New("unimplemented")
}

func (p *sensorServicePlugin) Init(opts opt_helper.Option) error {
	return errors.New("unimplemented")
}

func NewServicePlugin() mtp.ServicePlugin {
	return &sensorServicePlugin{}
}
