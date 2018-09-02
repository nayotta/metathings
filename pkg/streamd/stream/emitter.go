package stream_manager

type Emitter interface {
	Listen(event string, message interface{})
	Emit(event string, message interface{})
}
