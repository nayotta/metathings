package stream_manager

type Group struct {
}

func (self *Group) Id() string {
	panic("unimplemented")
}

func (self *Group) Inputs() []Input {
	panic("unimplemented")
}

func (self *Group) GetInput(id string) (Input, error) {
	panic("unimplemented")
}

func (self *Group) Outputs() []Output {
	panic("unimplemented")
}

func (self *Group) GetOutput(id string) (Output, error) {
	panic("unimplemented")
}

type GroupOption func(interface{})

func NewGroup(opts ...GroupOption) (*Group, error) {
	panic("unimplemented")
}
