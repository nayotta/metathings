package metathings_evaluatord_sdk

import (
	"encoding/json"
	"sync"
)

type Data interface {
	Iter() map[string]interface{}
	Get(string) interface{}
}

type data struct {
	buf []byte
	raw map[string]interface{}

	ensure_raw_once sync.Once
	ensure_buf_once sync.Once
}

func (d *data) ensure_raw() {
	d.ensure_raw_once.Do(func() {
		if d.raw != nil {
			return
		}

		d.raw = make(map[string]interface{})

		if d.buf == nil {
			d.buf = []byte(`{}`)
			return
		}

		json.Unmarshal(d.buf, &d.raw)
	})
}

func (d *data) ensure_buf() {
	d.ensure_buf_once.Do(func() {
		if d.buf != nil {
			return
		}

		if d.raw == nil {
			d.buf = []byte(`{}`)
			d.raw = make(map[string]interface{})
			return
		}

		d.buf, _ = json.Marshal(d.raw)
	})
}

func (d *data) Iter() map[string]interface{} {
	d.ensure_buf()
	out := make(map[string]interface{})
	json.Unmarshal(d.buf, &out)
	return out
}

func (d *data) Get(key string) interface{} {
	d.ensure_raw()
	val, ok := d.raw[key]
	if !ok {
		return nil
	}
	return val
}

func DataFromMap(raw map[string]interface{}) (Data, error) {
	return &data{raw: raw}, nil
}

func DataFromBytes(buf []byte) (Data, error) {
	return &data{buf: buf}, nil
}

type DataEncoder interface {
	Encode(Data) ([]byte, error)
}

type DataDecoder interface {
	Decode([]byte) (Data, error)
}

type JsonDataEncoder struct{}

func (*JsonDataEncoder) Encode(dat Data) ([]byte, error) {
	return json.Marshal(dat.Iter())
}

type JsonDataDecoder struct{}

func (*JsonDataDecoder) Decode(buf []byte) (Data, error) {
	m := make(map[string]interface{})

	if err := json.Unmarshal(buf, &m); err != nil {
		return nil, err
	}

	return DataFromMap(m)
}

var data_encoders map[string]DataEncoder
var data_encoders_once sync.Once

func registry_data_encoder(name string, enc DataEncoder) {
	data_encoders_once.Do(func() {
		data_encoders = make(map[string]DataEncoder)
	})

	data_encoders[name] = enc
}

func GetDataEncoder(name string) (DataEncoder, error) {
	enc, ok := data_encoders[name]
	if !ok {
		return nil, ErrUnsupportedDataEncoder
	}

	return enc, nil
}

var data_decoders map[string]DataDecoder
var data_decoders_once sync.Once

func registry_data_decoder(name string, dec DataDecoder) {
	data_decoders_once.Do(func() {
		data_decoders = make(map[string]DataDecoder)
	})

	data_decoders[name] = dec
}

func GetDataDecoder(name string) (DataDecoder, error) {
	dec, ok := data_decoders[name]
	if !ok {
		return nil, ErrUnsupportedDataDecoder
	}

	return dec, nil
}

func init() {
	registry_data_encoder("json", new(JsonDataEncoder))
	registry_data_decoder("json", new(JsonDataDecoder))
}
