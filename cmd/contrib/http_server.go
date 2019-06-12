package cmd_contrib

import (
	"context"
	"net"
	"net/http"

	"go.uber.org/fx"
)

type NewHttpServerParams struct {
	fx.In

	Lis net.Listener
}

func NewHttpServer(params NewHttpServerParams, lc fx.Lifecycle) *http.Server {
	s := &http.Server{}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go s.Serve(params.Lis)
			return nil
		},
		OnStop: func(context.Context) error {
			s.Close()
			return nil
		},
	})

	return s
}
