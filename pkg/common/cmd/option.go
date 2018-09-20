package cmd_helper

import (
	constant_helper "github.com/nayotta/metathings/pkg/common/constant"
)

type LogOptions struct {
	Level string
}

type ApplicationCredentialOptions struct {
	Id     string
	Secret string
}

type TokenOptions struct {
	Token string
}

type ServiceConfigOption struct {
	Address string
}

type ServiceConfigOptions struct {
	CoreAgentd  ServiceConfigOption `mapstructure:"core_agentd"`
	Cored       ServiceConfigOption
	Identityd   ServiceConfigOption
	Camerad     ServiceConfigOption
	Sensord     ServiceConfigOption
	Streamd     ServiceConfigOption
	Metathingsd ServiceConfigOption
}

type StorageOptions struct {
	Driver string
	Uri    string
}

type EndpointOptions struct {
	Type string
	Host string
}

type RootOptions struct {
	Config                string
	Stage                 string
	Verbose               bool
	Log                   LogOptions
	ApplicationCredential ApplicationCredentialOptions `mapstructure:"application_credential"`
	ServiceConfig         ServiceConfigOptions         `mapstructure:"service_config"`
}

func InitServiceConfigOptions(dst, src *ServiceConfigOptions) {
	init_service_config(&dst.Metathingsd, &src.Metathingsd, constant_helper.CONSTANT_METATHINGSD_DEFAULT_HOST)
	init_service_config(&dst.Cored, &src.Cored, dst.Metathingsd.Address)
	init_service_config(&dst.Identityd, &src.Identityd, dst.Metathingsd.Address)
	init_service_config(&dst.Camerad, &src.Camerad, dst.Metathingsd.Address)
	init_service_config(&dst.Sensord, &src.Sensord, dst.Metathingsd.Address)
	init_service_config(&dst.Streamd, &src.Streamd, dst.Metathingsd.Address)
}

func init_service_config(dst, src *ServiceConfigOption, constant string) {
	// 1. get address from config file
	addr := dst.Address
	if addr == "" {
		// 2. if address is empty, get address from command line
		addr = src.Address
		if addr == "" {
			// 3. if address is empty, get address from constant
			addr = constant
		}
	}
	// 4. set address to config option.
	dst.Address = addr
}
