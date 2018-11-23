package option_helper

import (
	"errors"
	"fmt"
)

type Option interface {
	Set(string, interface{})
	Get(string) interface{}
	Keys() []string
	Contains(string) bool
	Update(Option)
	GetString(string) string
	GetStrings(string) []string
	GetInt(string) int
	GetUInt32(string) uint32
	GetUInt64(string) uint64
	GetInt32(string) int32
	GetInt64(string) int64
	GetFloat32(string) float32
	GetFloat64(string) float64
	GetBool(string) bool
	Data() map[string]interface{}
}

func NewOptionMap(o ...map[string]interface{}) Option {
	if len(o) > 0 && o[0] != nil {
		return option(o[0])
	}
	return option{}
}

func NewOptionArr(args ...interface{}) Option {
	m := map[string]interface{}{}
	for i := 0; i < len(args); i += 2 {
		m[args[i].(string)] = args[i+1]
	}
	return NewOptionMap(m)
}

var NewOption = NewOptionArr

type option map[string]interface{}

func (o option) Set(k string, v interface{}) {
	o[k] = v
}

func (o option) Get(k string) interface{} {
	v, ok := o[k]
	if !ok {
		return nil
	}
	return v
}

func (o option) Keys() []string {
	ks := make([]string, 0, len(o))
	for k, _ := range o {
		ks = append(ks, k)
	}
	return ks
}

func (o option) Contains(k string) bool {
	_, ok := o[k]
	return ok
}

func (o option) Update(x Option) {
	for _, k := range x.Keys() {
		if !o.Contains(k) {
			panic(fmt.Sprintf("%v not found", k))
		}

		vx := x.Get(k)
		vx_opt, ok := vx.(Option)
		if ok {
			o.Get(k).(Option).Update(vx_opt)
		} else {
			o.Set(k, vx)
		}
	}
}

func (o option) GetString(k string) string {
	v := o.Get(k)
	if v == nil {
		return ""
	}
	return v.(string)
}

func (o option) GetStrings(k string) []string {
	v := o.Get(k)
	if v == nil {
		return nil
	}
	return v.([]string)
}

func (o option) GetInt(k string) int {
	v := o.Get(k)
	if v == nil {
		return 0
	}
	return v.(int)
}

func (o option) GetUInt32(k string) uint32 {
	v := o.Get(k)
	if v == nil {
		return 0
	}
	return v.(uint32)
}

func (o option) GetUInt64(k string) uint64 {
	v := o.Get(k)
	if v == nil {
		return 0
	}
	return v.(uint64)
}

func (o option) GetInt32(k string) int32 {
	v := o.Get(k)
	if v == nil {
		return 0
	}
	return v.(int32)
}

func (o option) GetInt64(k string) int64 {
	v := o.Get(k)
	if v == nil {
		return 0
	}
	return v.(int64)
}

func (o option) GetFloat32(k string) float32 {
	v := o.Get(k)
	if v == nil {
		return 0
	}
	return v.(float32)
}

func (o option) GetFloat64(k string) float64 {
	v := o.Get(k)
	if v == nil {
		return 0
	}
	return v.(float64)
}

func (o option) GetBool(k string) bool {
	v := o.Get(k)
	if v == nil {
		return false
	}
	return v.(bool)
}

func (o option) Data() map[string]interface{} {
	return o
}

func Copy(x Option) Option {
	o := option{}
	for _, k := range x.Keys() {
		o.Set(k, x.Get(k))
	}
	return o
}

var (
	ErrInvalidArguments = errors.New("invalid arguments")
)

func Setopt(conds map[string]func(key string, val interface{}) error) func(...interface{}) error {
	return func(args ...interface{}) error {
		if len(args)%2 != 0 {
			return ErrInvalidArguments
		}

		for i := 0; i < len(args); i += 2 {
			key, ok := args[i].(string)
			if !ok {
				return ErrInvalidArguments
			}
			val := args[i+1]

			cond, ok := conds[key]
			if !ok {
				return ErrInvalidArguments
			}
			if err := cond(key, val); err != nil {
				return err
			}
		}

		return nil
	}
}
