package cmd_contrib

import (
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"google.golang.org/grpc"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
)

type NewClientFactoryParams struct {
	fx.In

	Endpoints ServiceEndpointsOptioner
	Logger    log.FieldLogger
	Tracer    opentracing.Tracer `name:"opentracing_tracer" optional:"true"`
}

func NewClientFactory(p NewClientFactoryParams) (*client_helper.ClientFactory, error) {
	dep := p.Endpoints.GetServiceEndpoint(client_helper.DEFAULT_CONFIG)
	addr := dep.GetAddress()
	cred, err := NewClientTransportCredentials(dep)
	if err != nil {
		p.Logger.WithError(err).Debugf("failed to new client transport credentials")
		return nil, err
	}
	cfgs := client_helper.NewDefaultServiceConfigs(addr, cred)
	cfg_name := client_helper.DEFAULT_CONFIG.String()
	logger := p.Logger.WithFields(log.Fields{
		cfg_name + ".address":    addr,
		cfg_name + ".insecure":   dep.GetInsecure(),
		cfg_name + ".plain_text": dep.GetPlainText(),
		cfg_name + ".cert_file":  dep.GetCertFile(),
		cfg_name + ".key_file":   dep.GetKeyFile(),
	})

	for i := int32(client_helper.DEFAULT_CONFIG) + 1; i < int32(client_helper.OVERFLOW_CONFIG); i++ {
		typ := client_helper.ClientType(i)
		ep := p.Endpoints.GetServiceEndpoint(typ)
		addr := ep.GetAddress()
		cred, err := NewClientTransportCredentials(ep)
		if err != nil {
			logger.WithError(err).Debugf("failed to new client transport credentials")
			return nil, err
		}
		if addr != "" {
			cfgs.SetServiceConfig(typ, client_helper.ServiceConfig{
				Address:              addr,
				TransportCredentials: cred,
			})
			cfg_name := typ.String()
			logger = logger.WithFields(log.Fields{
				cfg_name + ".address":    addr,
				cfg_name + ".insecure":   ep.GetInsecure(),
				cfg_name + ".plain_text": ep.GetPlainText(),
				cfg_name + ".cert_file":  ep.GetCertFile(),
				cfg_name + ".key_file":   ep.GetKeyFile(),
			})
		}
	}
	logger.Debugf("new client factory")

	opts := client_helper.DefaultDialOption()
	if p.Tracer != nil {
		opts = append(opts,
			grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()),
			grpc.WithStreamInterceptor(grpc_opentracing.StreamClientInterceptor()))
	}

	return client_helper.NewClientFactory(cfgs, opts)
}
