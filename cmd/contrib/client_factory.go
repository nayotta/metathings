package cmd_contrib

import client_helper "github.com/nayotta/metathings/pkg/common/client"

func NewClientFactory(opt ServiceEndpointsOptioner) (*client_helper.ClientFactory, error) {
	cfgs := client_helper.NewDefaultServiceConfigs(opt.GetServiceEndpoint(client_helper.DEFAULT_CONFIG).GetAddress())

	for i := client_helper.DEFAULT_CONFIG + 1; i < client_helper.OVERFLOW_CONFIG; i++ {
		if ep := opt.GetServiceEndpoint(i); ep.GetAddress() != "" {
			cfgs[i] = client_helper.ServiceConfig{ep.GetAddress()}
		}
	}

	return client_helper.NewClientFactory(cfgs,
		client_helper.WithInsecureOptionFunc(),
	)
}
