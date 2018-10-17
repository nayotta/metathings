package cmd_contrib

type ServiceOptioner interface {
	GetServiceNameP() *string
	GetServiceName() string
	SetServiceName(string)
}

type ServiceOption struct {
	ServiceName string `mapstructure:"service_name"`
}

func (self *ServiceOption) GetServiceNameP() *string {
	return &self.ServiceName
}

func (self *ServiceOption) GetServiceName() string {
	return self.ServiceName
}

func (self *ServiceOption) SetServiceName(srv_name string) {
	self.ServiceName = srv_name
}
