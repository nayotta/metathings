package stream_manager

import "sync"

type Event string
type EventCallback func(event Event, message interface{})

const (
	START_EVENT = Event("start")
	STOP_EVENT  = Event("stop")
	ERROR_EVENT = Event("error")
)

type Emitter interface {
	Listen(event Event, cb EventCallback)
	Once(event Event, cb EventCallback)
	Emit(event Event, message interface{})
}

type emitter struct {
	lck      *sync.Mutex
	cbs      map[Event][]EventCallback
	once_cbs map[Event][]EventCallback
}

func (self *emitter) Listen(event Event, cb EventCallback) {
	self.lck.Lock()
	defer self.lck.Unlock()

	if _, ok := self.cbs[event]; !ok {
		self.cbs[event] = []EventCallback{}
	}

	self.cbs[event] = append(self.cbs[event], cb)
}

func (self *emitter) Once(event Event, cb EventCallback) {
	self.lck.Lock()
	defer self.lck.Unlock()

	if _, ok := self.once_cbs[event]; !ok {
		self.once_cbs[event] = []EventCallback{}
	}

	self.once_cbs[event] = append(self.once_cbs[event], cb)
}

func (self *emitter) Emit(event Event, message interface{}) {
	self.lck.Lock()
	defer self.lck.Unlock()

	for _, cb := range self.cbs[event] {
		go cb(event, message)
	}

	for _, cb := range self.once_cbs[event] {
		go cb(event, message)
	}

	delete(self.once_cbs, event)
}

func NewEmitter() Emitter {
	return &emitter{
		lck:      &sync.Mutex{},
		cbs:      make(map[Event][]EventCallback),
		once_cbs: make(map[Event][]EventCallback),
	}
}
