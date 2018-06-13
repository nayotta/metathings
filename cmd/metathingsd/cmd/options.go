package cmd

type _serviceConfigOptions struct {
	Identityd _identitydServiceOptions
}

type _identitydServiceOptions struct {
	Address string
}

type _storageOptions struct {
	Driver string
	Uri    string
}

type _heartbeatOptions struct {
	CoreAliveTimeout   int `mapstructure:"core_alive_timeout"`
	EntityAliveTimeout int `mapstructure:"entity_alive_timeout"`
}
