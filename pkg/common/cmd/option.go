package cmd_helper

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
