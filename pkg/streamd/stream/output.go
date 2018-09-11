package stream_manager

import (
	"encoding/json"
	"time"
)

type OutputState int32

const (
	OUTPUT_STATE_UNKNOWN OutputState = iota
	OUTPUT_STATE_RUNNING
	OUTPUT_STATE_STARTING
	OUTPUT_STATE_TERMINATING
	OUTPUT_STATE_STOP
	OUTPUT_STATE_OVERFLOW
)

type Output interface {
	Emitter
	Id() string
	Symbol() string
	Start() error
	Stop() error
	State() OutputState
	Close()
}

type OutputOption func(interface{})

type OutputFactory func(opts ...OutputOption) (Output, error)

var output_factorys = make(map[string]OutputFactory)

func RegisterOutput(name string, fty OutputFactory) {
	if _, ok := output_factorys[name]; !ok {
		output_factorys[name] = fty
	}
}

type OutputMetadata struct {
	StreamData
}

func (self *OutputMetadata) SensorId() string {
	return self.AsString("sensor_id")
}

func (self *OutputMetadata) SensorName() string {
	return self.AsString("sensor_name")
}

func (self *OutputMetadata) From() string {
	return self.AsString("from")
}

func (self *OutputMetadata) CreatedAt() time.Time {
	return self.AsTime("created_at")
}

func (self *OutputMetadata) ArrviedAt() time.Time {
	return self.AsTime("arrvied_at")
}

func (self *OutputMetadata) Data() StreamData {
	return self.StreamData
}

type OutputData struct {
	StreamData
	metadata *OutputMetadata
}

func (self *OutputData) Metadata() *OutputMetadata {
	return self.metadata
}

func (self *OutputData) Data() StreamData {
	return self.StreamData
}

func NewOutputData(data map[string]interface{}, metadata map[string]interface{}) *OutputData {
	omd := &OutputMetadata{NewStreamData(metadata)}
	od := &OutputData{
		StreamData: NewStreamData(data),
		metadata:   omd,
	}

	return od
}

func UpstreamDataToOutputData(data *UpstreamData) *OutputData {
	return &OutputData{
		StreamData: data.Data(),
		metadata:   &OutputMetadata{data.Metadata().Data()},
	}
}

func InputDataToOutputData(data *InputData) *OutputData {
	return &OutputData{
		StreamData: data.Data(),
		metadata:   &OutputMetadata{data.Metadata().Data()},
	}
}

type OutputDataCodec struct{}

func (self *OutputDataCodec) Encode(value interface{}) ([]byte, error) {
	val, ok := value.(*OutputData)
	if !ok {
		return nil, ErrOutputDataCodec
	}

	mtd := val.Metadata().Data()
	d := val.Data().Data()
	d["#metadata"] = mtd

	return json.Marshal(d)
}

func (self *OutputDataCodec) Decode(data []byte) (interface{}, error) {
	d := map[string]interface{}{}
	err := json.Unmarshal(data, &d)
	if err != nil {
		return nil, err
	}

	mt_val, ok := d["#metadata"]
	if !ok {
		return nil, ErrOutputDataCodec
	}

	delete(d, "#metadata")
	mt, ok := mt_val.(map[string]interface{})
	if !ok {
		return nil, ErrOutputDataCodec
	}

	return NewOutputData(d, mt), nil
}
