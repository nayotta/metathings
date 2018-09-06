package stream_manager

type Source struct {
}

func (self *Source) Upstream() Upstream {
	panic("unimplemented")
}

func (self *Source) Id() string {
	panic("unimplemented")
}

type SourceOption func(o interface{})

func NewSource(opts ...SourceOption) (*Source, error) {
	panic("unimplemented")
}
