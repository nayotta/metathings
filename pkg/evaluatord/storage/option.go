package metathings_evaluatord_storage

import (
	"time"

	"github.com/spf13/cast"
	"github.com/stretchr/objx"
)

func to_time_string(v interface{}) (string, bool) {
	switch v.(type) {
	case time.Time:
		return cast.ToTime(v).Format(time.RFC3339), true
	case time.Duration:
		return cast.ToDuration(v).String(), true
	case string:
		s := cast.ToString(v)
		_, err := time.Parse(time.RFC3339, s)
		if err != nil {
			_, err := cast.ToDurationE(s)
			if err != nil {
				return "", false
			}
		}
		return s, true
	default:
		return "", false
	}

}

func SetRangeStartOption(start interface{}) func(objx.Map) {
	return func(o objx.Map) {
		if s, ok := to_time_string(start); ok {
			o.Set("start", s)
		}
	}
}

func SetRangeStopOption(stop interface{}) func(objx.Map) {
	return func(o objx.Map) {
		if s, ok := to_time_string(stop); ok {
			o.Set("stop", s)
		}
	}
}

type ListTasksBySourceOption func(objx.Map)
