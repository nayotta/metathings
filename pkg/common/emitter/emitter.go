package emitter

type Emitter interface {
	OnEvent(func(Event) error)
	Trigger(Event) error
}

type emitter struct {
	callbacks []func(Event) error
}

func (self *emitter) OnEvent(cb func(Event) error) {
	self.callbacks = append(self.callbacks, cb)
}

func (self *emitter) Trigger(evt Event) error {
	for _, cb := range self.callbacks {
		if err := cb(evt); err != nil {
			return err
		}
	}
	return nil
}

func NewEmitter() Emitter {
	return &emitter{callbacks: []func(Event) error{}}
}
