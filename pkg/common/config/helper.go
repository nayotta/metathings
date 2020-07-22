package config_helper

import "errors"

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
