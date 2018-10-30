package cmd_contrib

type BaseOption struct {
	ConfigOption           `mapstructure:",squash"`
	StageOption            `mapstructure:",squash"`
	VerboseOption          `mapstructure:",squash"`
	LoggerOption           `mapstructure:",squash"`
	CredentialOption       `mapstructure:",squash"`
	ServiceEndpointsOption `mapstructure:",squash"`
}

func CreateBaseOption() BaseOption {
	return BaseOption{
		ServiceEndpointsOption: CreateServiceEndpointsOption(),
	}
}

type ServiceBaseOption struct {
	BaseOption                `mapstructure:",squash"`
	ServiceOption             `mapstructure:",squash"`
	ListenOption              `mapstructure:",squash"`
	TransportCredentialOption `mapstructure:",squash"`
	StorageOption             `mapstructure:",squash"`
}

func (self *ServiceBaseOption) GetStorage() StorageOptioner {
	return &self.StorageOption
}

func (self *ServiceBaseOption) GetTransportCredential() TransportCredentialOptioner {
	return &self.TransportCredentialOption
}

func CreateServiceBaseOption() ServiceBaseOption {
	return ServiceBaseOption{
		BaseOption: CreateBaseOption(),
	}
}
