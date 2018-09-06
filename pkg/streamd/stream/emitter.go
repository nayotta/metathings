package stream_manager

type Event string
type EventCallback func(event Event, message interface{})

const (
	START_EVENT = Event("start")
	STOP_EVENT  = Event("stop")
	ERROR_EVENT = Event("error")
)

type Emitter interface {
	Listen(event Event, cb EventCallback)
	Emit(event Event, message interface{})
}

type emitter struct {
	cbs map[Event][]EventCallback
}

func (self *emitter) Listen(event Event, cb EventCallback) {
	if _, ok := self.cbs[event]; !ok {
		self.cbs[event] = []EventCallback{}
	}

	self.cbs[event] = append(self.cbs[event], cb)
}

func (self *emitter) Emit(event Event, message interface{}) {
	var cbs []EventCallback
	var ok bool

	if cbs, ok = self.cbs[event]; !ok {
		return
	}

	go func() {
		for _, cb := range cbs {
			cb(event, message)
		}
	}()
}

func NewEmitter() Emitter {
	return &emitter{
		cbs: make(map[Event][]EventCallback),
	}
}
