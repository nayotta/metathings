package main

import (
	"net"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	plugin "github.com/bigdatagz/metathings/pkg/core/plugin"
	service "github.com/bigdatagz/metathings/pkg/echo/service"
	pb "github.com/bigdatagz/metathings/pkg/proto/echo"
)

var config struct {
	Bind string
}

type echoServicePlugin struct{}

func (p *echoServicePlugin) Run() error {
	lis, err := net.Listen("tcp", config.Bind)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	srv := service.NewEchoService()

	pb.RegisterEchoServiceServer(s, srv)
	log.Infof("echo(core) service listen on %v", config.Bind)
	return s.Serve(lis)
}

func (p *echoServicePlugin) Init(opt plugin.Option) error {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("echo")
	v.AddConfigPath(opt.Config)
	err := v.ReadInConfig()
	if err != nil {
		return err
	}

	err = v.Unmarshal(&config)
	if err != nil {
		return err
	}

	return nil
}

func NewPlugin() plugin.CorePlugin {
	return &echoServicePlugin{}
}
