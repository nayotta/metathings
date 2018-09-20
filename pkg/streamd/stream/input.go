package stream_manager

import (
	"encoding/json"
	"time"
)

type InputState int32

const (
	INPUT_STATE_UNKNOWN InputState = iota
	INPUT_STATE_RUNNING
	INPUT_STATE_STARTING
	INPUT_STATE_TERMINATING
	INPUT_STATE_STOP
	INPUT_STATE_OVERFLOW
)

type Input interface {
	Emitter
	Id() string
	Symbol() string
	Start() error
	Stop() error
	State() InputState
	Close()
}

type InputFactory interface {
	Set(key string, val interface{}) InputFactory
	New() (Input, error)
}

var new_input_factories = make(map[string]func() InputFactory)

func RegisterInputFactory(name string, fn func() InputFactory) {
	if _, ok := new_input_factories[name]; !ok {
		new_input_factories[name] = fn
	}
}

func NewInputFactory(name string) (InputFactory, error) {
	new_fn, ok := new_input_factories[name]
	if !ok {
		return nil, ErrUnregisteredInputFactory
	}
	return new_fn(), nil
}

type InputMetadata struct {
	StreamData
}

func (self *InputMetadata) SensorId() string {
	return self.AsString("sensor_id")
}

func (self *InputMetadata) SensorName() string {
	return self.AsString("sensor_name")
}

func (self *InputMetadata) From() string {
	return self.AsString("from")
}

func (self *InputMetadata) CreatedAt() time.Time {
	return self.AsTime("created_at")
}

func (self *InputMetadata) ArrviedAt() time.Time {
	return self.AsTime("arrvied_at")
}

func (self *InputMetadata) Data() StreamData {
	return self.StreamData
}

type InputData struct {
	StreamData
	metadata *InputMetadata
}

func (self *InputData) Metadata() *InputMetadata {
	return self.metadata
}

func (self *InputData) Data() StreamData {
	return self.StreamData
}

func NewInputData(data map[string]interface{}, metadata map[string]interface{}) *InputData {
	imd := &InputMetadata{NewStreamData(metadata)}
	id := &InputData{
		StreamData: NewStreamData(data),
		metadata:   imd,
	}

	return id
}

func UpstreamDataToInputData(data *UpstreamData) *InputData {
	return &InputData{
		StreamData: data.Data(),
		metadata:   &InputMetadata{data.Metadata().Data()},
	}
}

type InputDataCodec struct{}

func (self *InputDataCodec) Encode(value interface{}) ([]byte, error) {
	val, ok := value.(*InputData)
	if !ok {
		return nil, ErrInputDataCodec
	}

	mtd := val.Metadata().Data().Data()
	d := val.Data().Data()
	d["#metadata"] = mtd

	return json.Marshal(d)
}

func (self *InputDataCodec) Decode(data []byte) (interface{}, error) {
	d := map[string]interface{}{}
	err := json.Unmarshal(data, &d)
	if err != nil {
		return nil, err
	}

	mt_val, ok := d["#metadata"]
	if !ok {
		return nil, ErrInputDataCodec
	}

	delete(d, "#metadata")
	mt, ok := mt_val.(map[string]interface{})
	if !ok {
		return nil, ErrInputDataCodec
	}

	return NewInputData(d, mt), nil
}
