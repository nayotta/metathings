package protobuf_helper

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/golang/protobuf/ptypes/wrappers"
)

func Now() timestamp.Timestamp {
	now := time.Now()
	return FromTime(now)
}

func FromTime(t time.Time) timestamp.Timestamp {
	return timestamp.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.UnixNano() % 1e9),
	}
}

func ToTime(t timestamp.Timestamp) time.Time {
	return time.Unix(t.Seconds, int64(t.Nanos))
}

func FromTimestamp(ts int64) timestamp.Timestamp {
	return timestamp.Timestamp{
		Seconds: ts / 1e9,
		Nanos:   int32(ts % 1e9),
	}
}

func ExtractStringMap(xs map[string]*wrappers.StringValue) map[string]interface{} {
	ys := make(map[string]interface{})

	for k, v := range xs {
		ys[k] = v.GetValue()
	}

	return ys
}

func ExtractStringMapToString(xs map[string]*wrappers.StringValue) map[string]string {
	ys := make(map[string]string)

	for k, v := range xs {
		ys[k] = v.GetValue()
	}

	return ys
}

func MustParseExtra(x map[string]*wrappers.StringValue) string {
	var buf []byte
	var err error

	if x == nil {
		return `{}`
	}

	extra_map := ExtractStringMap(x)
	if buf, err = json.Marshal(extra_map); err != nil {
		return `{}`
	}

	return string(buf)
}

type Enumer struct {
	prefix string
	names  map[int32]string
	values map[string]int32
}

func (self Enumer) ToString(x int32) string {
	s, ok := self.names[x]
	if !ok {
		return "unknown"
	}
	return strings.TrimPrefix(strings.ToLower(s), self.prefix+"_")
}

func (self Enumer) ToValue(x string) int32 {
	i, ok := self.values[strings.ToUpper(self.prefix+"_"+x)]
	if !ok {
		return 0
	}
	return i
}

func NewEnumer(prefix string, names map[int32]string, values map[string]int32) Enumer {
	return Enumer{
		prefix: strings.ToLower(prefix),
		names:  names,
		values: values,
	}
}

func ToStringMap(msg proto.Message) map[string]interface{} {
	buf, _ := new(jsonpb.Marshaler).MarshalToString(msg)
	y := make(map[string]interface{})
	json.Unmarshal([]byte(buf), &y)
	return y
}
