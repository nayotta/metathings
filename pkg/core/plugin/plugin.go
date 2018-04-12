package metathings_core_plugin

type Option struct {
	Config string
}

type CorePlugin interface {
	Init(opt Option) error
	Run() error
}
