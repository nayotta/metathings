package protobuf_helper

import (
	"time"

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
		Nanos:   int32(t.UnixNano() % 1000000000),
	}
}

func ToTime(t timestamp.Timestamp) time.Time {
	return time.Unix(t.Seconds, int64(t.Nanos))
}

func ExtractStringMap(xs map[string]*wrappers.StringValue) map[string]interface{} {
	ys := make(map[string]interface{})

	for k, v := range xs {
		ys[k] = v.GetValue()
	}

	return ys
}
