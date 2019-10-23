package cmd_contrib

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
)

func NewClientFactory(eps ServiceEndpointsOptioner, creds credentials.TransportCredentials, logger log.FieldLogger) (*client_helper.ClientFactory, error) {
	addr := eps.GetServiceEndpoint(client_helper.DEFAULT_CONFIG).GetAddress()
	cfgs := client_helper.NewDefaultServiceConfigs(addr)
	logger = logger.WithField(client_helper.DEFAULT_CONFIG.String(), addr)

	for i := int32(client_helper.DEFAULT_CONFIG) + 1; i < int32(client_helper.OVERFLOW_CONFIG); i++ {
		typ := client_helper.ClientType(i)
		ep := eps.GetServiceEndpoint(typ)
		addr = ep.GetAddress()
		if addr != "" {
			cfgs.SetServiceConfig(typ, client_helper.ServiceConfig{Address: addr})
			logger = logger.WithField(typ.String(), addr)
		}
	}

	opts := client_helper.DefaultDialOption()
	if creds == nil {
		opts = append(opts, grpc.WithInsecure())
	} else {
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	logger.WithField("WithInsecureOptionFunc", true).Debugf("new client factory")

	return client_helper.NewClientFactory(cfgs, opts)
}
