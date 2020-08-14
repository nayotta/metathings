package config_helper

import (
	"errors"
	"strings"
)

var (
	ErrInvalidArgument     = errors.New("invalid argument")
	ErrExpectedKeyNotFound = errors.New("expected key not found")
)

func ParseConfigOption(k string, m map[string]interface{}, a ...interface{}) (string, []interface{}, error) {
	var key string
	var val interface{}
	var name string
	var ok bool

	var y []interface{}

	if val, ok = m[k]; !ok {
		return "", nil, ErrExpectedKeyNotFound
	}

	if name, ok = val.(string); !ok {
		return "", nil, ErrInvalidArgument
	}

	for key, val = range m {
		if key == k {
			continue
		}

		y = append(y, key, val)
	}

	y = append(y, a...)

	return name, y, nil
}

func FlattenConfigOption(m map[string]interface{}, a ...interface{}) []interface{} {
	var y []interface{}
	for k, v := range m {
		y = append(y, k, v)
	}

	y = append(y, a...)

	return y
}

func FoldConfigOption(xs []interface{}, t string) (map[string]interface{}, error) {
	y := map[string]interface{}{}
	for i := 0; i < len(xs); i += 2 {
		k, ok := xs[i].(string)
		if !ok {
			return nil, ErrInvalidArgument
		}
		v := xs[i+1]
		ks := strings.SplitN(k, ".", 2)
		if len(ks) > 1 && ks[0] == t {
			y[ks[1]] = v
		}
	}
	return y, nil
}
