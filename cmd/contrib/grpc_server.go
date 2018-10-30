package cmd_contrib

import (
	"context"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
)

type NewGrpcServerParams struct {
	fx.In

	Lis   net.Listener
	Creds credentials.TransportCredentials
}

func NewGrpcServer(params NewGrpcServerParams, lc fx.Lifecycle) *grpc.Server {

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_helper.UnaryServerInterceptor()),
		grpc.StreamInterceptor(grpc_auth.StreamServerInterceptor(nil)),
	}

	if params.Creds != nil {
		opts = append(opts, grpc.Creds(params.Creds))
	}

	s := grpc.NewServer(opts...)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go s.Serve(params.Lis)
			return nil
		},
		OnStop: func(context.Context) error {
			s.Stop()
			return nil
		},
	})

	return s
}
