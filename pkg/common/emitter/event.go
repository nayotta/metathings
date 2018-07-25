package emitter

type Event interface {
	Name() string
	Data() interface{}
}

type event struct {
	name string
	data interface{}
}

func (self *event) Name() string {
	return self.name
}

func (self *event) Data() interface{} {
	return self.data
}

func NewEvent(name string, data interface{}) Event {
	return &event{
		name: name,
		data: data,
	}
}
