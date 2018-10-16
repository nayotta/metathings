package cmd_contrib

type ConfigOptioner interface {
	GetConfigP() *string
	GetConfig() string
	SetConfig(string)
}

type ConfigOption struct {
	Config string
}

func (self *ConfigOption) GetConfigP() *string {
	return &self.Config
}

func (self *ConfigOption) GetConfig() string {
	return self.Config
}

func (self *ConfigOption) SetConfig(cfg string) {
	self.Config = cfg
}
