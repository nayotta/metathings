package cmd_contrib

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
	GetServiceEndpoint(string) ServiceEndpointsOptioner
	SetServiceEndpoint(string, ServiceEndpointOptioner)
}

type ServiceEndpointsOption struct {
	eps map[string]ServiceEndpointOptioner
}

func (self *ServiceEndpointsOption) GetServiceEndpoint(name string) ServiceEndpointOptioner {
	return self.eps[name]
}

func (self *ServiceEndpointsOption) SetServiceEndpoint(name string, ep ServiceEndpointOptioner) {
	self.eps[name] = ep
}
