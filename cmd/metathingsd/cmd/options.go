package cmd

type _heartbeatOptions struct {
	CoreAliveTimeout   int `mapstructure:"core_alive_timeout"`
	EntityAliveTimeout int `mapstructure:"entity_alive_timeout"`
}
