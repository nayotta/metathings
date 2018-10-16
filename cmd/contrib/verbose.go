package cmd_contrib

type VerboseOptioner interface {
	GetVerboseP() *bool
	GetVerbose() bool
	SetVerbose(bool)
}

type VerboseOption struct {
	Verbose bool
}

func (self *VerboseOption) GetVerboseP() *bool {
	return &self.Verbose
}

func (self *VerboseOption) GetVerbose() bool {
	return self.Verbose
}

func (self *VerboseOption) SetVerbose(verbose bool) {
	self.Verbose = verbose
}
