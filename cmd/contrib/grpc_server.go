package cmd_contrib

import (
	"context"
	"io"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
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
	Tracer opentracing.Tracer `name:"opentracing_tracer" optional:"true"`
	Closer io.Closer          `name:"opentracing_closer" optional:"true"`
}

func NewGrpcServer(params NewGrpcServerParams, lc fx.Lifecycle) *grpc.Server {
	var unary_server_interceptors []grpc.UnaryServerInterceptor
	var stream_server_interceptors []grpc.StreamServerInterceptor

	if params.Tracer != nil {
		unary_server_interceptors = append(unary_server_interceptors, grpc_opentracing.UnaryServerInterceptor(
			grpc_opentracing.WithTracer(params.Tracer),
		))
		stream_server_interceptors = append(stream_server_interceptors, grpc_opentracing.StreamServerInterceptor(
			grpc_opentracing.WithTracer(params.Tracer),
		))
	}
	unary_server_interceptors = append(unary_server_interceptors, grpc_helper.UnaryServerInterceptor(params.Logger))
	stream_server_interceptors = append(stream_server_interceptors, grpc_helper.StreamServerInterceptor(params.Logger))

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(unary_server_interceptors...)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(stream_server_interceptors...)),
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
			if params.Closer != nil {
				params.Closer.Close()
			}
			params.Logger.Infof("grpc server stoped")
			return nil
		},
	})

	return s
}
