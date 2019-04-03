package config_helper

import "errors"

var (
	ErrInvalidArgument = errors.New("invalid argument")
)

func ParseConfigOption(k string, m map[string]interface{}) (string, []interface{}, error) {
	var key string
	var val interface{}
	var name string
	var ok bool

	var y []interface{}

	if val, ok = m[k]; !ok {
		return "", nil, ErrInvalidArgument
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

	return name, y, nil
}
