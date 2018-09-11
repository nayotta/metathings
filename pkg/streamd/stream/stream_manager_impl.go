package stream_manager

type streamManagerImplOption struct{}

type streamManagerImpl struct{}

func (self *streamManagerImpl) NewStream(opts ...StreamOption) (Stream, error) {
	panic("unimplemented")
}

func (self *streamManagerImpl) GetStream(id string) (Stream, error) {
	panic("unimplemented")
}

func newStreamManagerImpl() (StreamManager, error) {
	panic("unimplemented")
}
