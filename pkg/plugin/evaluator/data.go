package metathings_plugin_evaluator

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
