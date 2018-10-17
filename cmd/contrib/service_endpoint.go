package cmd_contrib

import (
	cli_helper "github.com/nayotta/metathings/pkg/common/client"
)

type ServiceEndpointOptioner interface {
	GetAddressP() *string
	GetAddress() string
	SetAddress(string)
}

type ServiceEndpointOption struct {
	Address string
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
	GetServiceEndpoint(int) ServiceEndpointOptioner
	SetServiceEndpoint(int, ServiceEndpointOptioner)
}

type ServiceEndpointsOption struct {
	eps map[int]ServiceEndpointOptioner
}

func (self *ServiceEndpointsOption) GetServiceEndpoint(srv int) ServiceEndpointOptioner {
	return self.eps[srv]
}

func (self *ServiceEndpointsOption) SetServiceEndpoint(srv int, ep ServiceEndpointOptioner) {
	self.eps[srv] = ep
}

func CreateServiceEndpointsOption() ServiceEndpointsOption {
	eps := make(map[int]ServiceEndpointOptioner)

	for i := cli_helper.DEFAULT_CONFIG; i < cli_helper.OVERFLOW_CONFIG; i++ {
		eps[i] = &ServiceEndpointOption{}
	}

	return ServiceEndpointsOption{
		eps: eps,
	}
}
