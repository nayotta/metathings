package option_helper

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
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

func InvalidArgument(key string) error {
	return fmt.Errorf("invalid argument: %v", key)
}

type SetoptConds map[string]func(string, interface{}) error

func Setopt(conds SetoptConds) func(...interface{}) error {
	return func(args ...interface{}) error {
		if len(args)%2 != 0 {
			return InvalidArgument("arguments")
		}

		for i := 0; i < len(args); i += 2 {
			key, ok := args[i].(string)
			if !ok {
				return InvalidArgument("arguments")
			}
			val := args[i+1]

			cond, ok := conds[key]
			if !ok {
				return InvalidArgument(key)
			}
			if err := cond(key, val); err != nil {
				return err
			}
		}

		return nil
	}
}

func ToString(v *string) func(string, interface{}) error {
	return func(key string, val interface{}) error {
		var ok bool
		*v, ok = val.(string)
		if !ok {
			return InvalidArgument(key)
		}
		return nil
	}
}

func ToByte(v *byte) func(string, interface{}) error {
	return func(key string, val interface{}) error {
		var ok bool
		*v, ok = val.(byte)
		if !ok {
			return InvalidArgument(key)
		}
		return nil
	}
}

func ToBool(v *bool) func(string, interface{}) error {
	return func(key string, val interface{}) error {
		var ok bool
		*v, ok = val.(bool)
		if !ok {
			return InvalidArgument(key)
		}
		return nil
	}
}

func ToInt(v *int) func(string, interface{}) error {
	return func(key string, val interface{}) error {
		var ok bool
		*v, ok = val.(int)
		if !ok {
			return InvalidArgument(key)
		}
		return nil
	}
}

func ToInt32(v *int32) func(string, interface{}) error {
	return func(key string, val interface{}) error {
		var ok bool
		*v, ok = val.(int32)
		if !ok {
			return InvalidArgument(key)
		}
		return nil
	}
}

func ToInt64(v *int64) func(string, interface{}) error {
	return func(key string, val interface{}) error {
		var ok bool
		*v, ok = val.(int64)
		if !ok {
			return InvalidArgument(key)
		}
		return nil
	}
}

func ToDuration(v *time.Duration) func(string, interface{}) error {
	return func(key string, val interface{}) error {
		var ok bool
		*v, ok = val.(time.Duration)
		if !ok {
			return InvalidArgument(key)
		}
		return nil
	}
}

func ToLogger(v *log.FieldLogger) func(string, interface{}) error {
	return func(key string, val interface{}) error {
		var ok bool
		*v, ok = val.(log.FieldLogger)
		if !ok {
			return InvalidArgument(key)
		}
		return nil
	}
}
