package metathings_core_plugin

const METATHINGS_PLUGIN_PREFIX = "mtp"

type Option struct {
	Args []string
}

type CorePlugin interface {
	Init(opt Option) error
	Run() error
}
