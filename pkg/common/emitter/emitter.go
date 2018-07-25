package emitter

type Emitter interface {
	OnEvent(func(Event) error)
	Trigger(Event) error
}

type emitter struct {
	listeners []func(Event) error
}

func (self *emitter) OnEvent(lis func(Event) error) {
	self.listeners = append(self.listeners, lis)
}

func (self *emitter) Trigger(evt Event) error {
	for _, lis := range self.listeners {
		if err := lis(evt); err != nil {
			return err
		}
	}
	return nil
}

func NewEmitter() Emitter {
	return &emitter{listeners: []func(Event) error{}}
}
