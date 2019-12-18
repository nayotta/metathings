package cmd_contrib

import (
	client_helper "github.com/nayotta/metathings/pkg/common/client"
)

type ServiceEndpointOptioner interface {
	TransportCredentialOptioner
	GetAddressP() *string
	GetAddress() string
	SetAddress(string)
}

type ServiceEndpointOption struct {
	TransportCredentialOption `mapstructure:",squash"`
	Address                   string
}

func (self *ServiceEndpointOption) GetAddressP() *string {
	return &self.Address
}

func (self *ServiceEndpointOption) GetAddress() string {
	return self.Address
}

func (self *ServiceEndpointOption) SetAddress(addr string) {
	self.Address = addr
}

type ServiceEndpointsOptioner interface {
	GetServiceEndpoint(interface{}) ServiceEndpointOptioner
}

type ServiceEndpointsOption struct {
	ServiceEndpoint map[string]*ServiceEndpointOption `mapstructure:"service_endpoint"`
}

func (self *ServiceEndpointsOption) GetServiceEndpoint(srv interface{}) ServiceEndpointOptioner {
	switch typ := srv.(type) {
	case string:
		return self.ServiceEndpoint[typ]
	case client_helper.ClientType:
		return self.ServiceEndpoint[typ.String()]
	}
	panic("unexpected type")
}

func CreateServiceEndpointsOption() ServiceEndpointsOption {
	eps := make(map[string]*ServiceEndpointOption)

	for i := int32(client_helper.DEFAULT_CONFIG); i < int32(client_helper.OVERFLOW_CONFIG); i++ {
		eps[client_helper.ClientType(i).String()] = &ServiceEndpointOption{}
	}

	return ServiceEndpointsOption{
		ServiceEndpoint: eps,
	}
}

func NewServiceEndpointsOptionWithTransportCredentialOption(x ServiceEndpointsOptioner, y TransportCredentialOptioner) ServiceEndpointsOptioner {
	epsz := make(map[string]*ServiceEndpointOption)

	for i := int32(client_helper.DEFAULT_CONFIG); i < int32(client_helper.OVERFLOW_CONFIG); i++ {
		typ := client_helper.ClientType(i)
		epx := x.GetServiceEndpoint(typ)
		epz := &ServiceEndpointOption{
			Address: epx.GetAddress(),
			TransportCredentialOption: TransportCredentialOption{
				Insecure:  y.GetInsecure(),
				PlainText: y.GetPlainText(),
				CertFile:  y.GetCertFile(),
				KeyFile:   y.GetKeyFile(),
			},
		}
		epsz[typ.String()] = epz
	}

	z := &ServiceEndpointsOption{ServiceEndpoint: epsz}
	return z
}
