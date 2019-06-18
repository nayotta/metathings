package cmd_contrib

import (
	"context"
	"net"
	"net/http"

	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type NewHttpServerParams struct {
	fx.In

	Lis    net.Listener
	Logger log.FieldLogger
}

func NewHttpServer(params NewHttpServerParams, lc fx.Lifecycle) *http.Server {
	s := &http.Server{}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go s.Serve(params.Lis)
			params.Logger.Infof("http server started")
			return nil
		},
		OnStop: func(context.Context) error {
			s.Close()
			params.Logger.Infof("http server stoped")
			return nil
		},
	})

	return s
}
