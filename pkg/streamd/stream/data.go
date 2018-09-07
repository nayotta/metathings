package stream_manager

import (
	"strconv"
	"time"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type StreamData interface {
	opt_helper.Option
	AsString(string) string
	AsInt(string) int
	AsInt32(string) int32
	AsUInt32(string) uint32
	AsInt64(string) int64
	AsUInt64(string) uint64
	AsFloat32(string) float32
	AsFloat64(string) float64
	AsTime(string) time.Time
}

type streamData struct {
	opt_helper.Option
}

func (self *streamData) AsString(k string) string {
	return self.GetString(k)
}

func (self *streamData) as_int64(k string, bit_size int) int64 {
	v := self.GetString(k)
	if v == "" {
		return 0
	}

	x, err := strconv.ParseInt(v, 10, bit_size)
	if err != nil {
		return 0
	}

	return x
}

func (self *streamData) as_uint64(k string, bit_size int) uint64 {
	v := self.GetString(k)
	if v == "" {
		return 0
	}

	x, err := strconv.ParseUint(v, 10, bit_size)
	if err != nil {
		return 0
	}

	return x
}

func (self *streamData) AsInt(k string) int {
	return int(self.as_int64(k, 32))
}

func (self *streamData) AsInt32(k string) int32 {
	return int32(self.as_int64(k, 32))
}

func (self *streamData) AsInt64(k string) int64 {
	return self.as_int64(k, 64)
}

func (self *streamData) AsUInt32(k string) uint32 {
	return uint32(self.as_uint64(k, 32))
}

func (self *streamData) AsUInt64(k string) uint64 {
	return self.as_uint64(k, 64)
}

func (self *streamData) as_float64(k string, bit_size int) float64 {
	v := self.GetString(k)
	if v == "" {
		return 0
	}

	x, err := strconv.ParseFloat(v, bit_size)
	if err != nil {
		return 0
	}

	return x
}

func (self *streamData) AsTime(k string) time.Time {
	v := self.as_int64(k, 64)
	sec := v / 10e9
	nano := v % 10e9
	return time.Unix(sec, nano)
}

func (self *streamData) AsFloat32(k string) float32 {
	return float32(self.as_float64(k, 32))
}

func (self *streamData) AsFloat64(k string) float64 {
	return self.as_float64(k, 64)
}

func NewStreamData(data map[string]interface{}) StreamData {
	return &streamData{opt_helper.NewOptionMap(data)}
}
