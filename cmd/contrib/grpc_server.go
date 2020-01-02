package cmd_contrib

import (
	"context"
	"net"

	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	log "github.com/sirupsen/logrus"
)

type NewGrpcServerParams struct {
	fx.In

	Lis    net.Listener
	Creds  credentials.TransportCredentials
	Logger log.FieldLogger
}

func NewGrpcServer(params NewGrpcServerParams, lc fx.Lifecycle) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_helper.UnaryServerInterceptor(params.Logger)),
		grpc.StreamInterceptor(grpc_helper.StreamServerInterceptor(params.Logger)),
	}

	if params.Creds != nil {
		opts = append(opts, grpc.Creds(params.Creds))
	}

	s := grpc.NewServer(opts...)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go s.Serve(params.Lis)
			params.Logger.Infof("grpc server started")
			return nil
		},
		OnStop: func(context.Context) error {
			s.Stop()
			params.Logger.Infof("grpc server stoped")
			return nil
		},
	})

	return s
}
