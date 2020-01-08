package cmd_contrib

import (
	"context"
	"io"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/opentracing-contrib/go-gorilla/gorilla"
	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type NewHttpServerParams struct {
	fx.In

	Lis    net.Listener
	Logger log.FieldLogger
	Router *mux.Router
	Tracer opentracing.Tracer `name:"opentracing_tracer" optional:"true"`
	Closer io.Closer          `name:"opentracing_closer" optional:"true"`
}

func NewHttpServer(params NewHttpServerParams, lc fx.Lifecycle) error {
	s := &http.Server{
		Handler: params.Router,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			if params.Tracer != nil {
				if err := params.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
					if h := route.GetHandler(); h != nil {
						route.Handler(gorilla.Middleware(params.Tracer, h))
					}
					return nil
				}); err != nil {
					return err
				}
			}

			go s.Serve(params.Lis)
			params.Logger.Infof("http server started")
			return nil
		},
		OnStop: func(context.Context) error {
			s.Close()
			if params.Closer != nil {
				params.Closer.Close()
			}
			params.Logger.Infof("http server stoped")
			return nil
		},
	})

	return nil
}
