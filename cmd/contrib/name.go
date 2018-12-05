package cmd_contrib

type NameOptioner interface {
	GetNameP() *string
	GetName() string
	SetName(string)
}

type NameOption struct {
	Name string `mapstructure:"name"`
}

func (self *NameOption) GetNameP() *string {
	return &self.Name
}

func (self *NameOption) GetName() string {
	return self.Name
}

func (self *NameOption) SetName(name string) {
	self.Name = name
}
