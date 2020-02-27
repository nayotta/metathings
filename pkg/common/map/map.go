package map_helper

import (
	"strconv"
	"strings"

	"github.com/spf13/cast"
)

type Seeker map[string]interface{}

func (s Seeker) Find(key string) interface{} {
	var ok bool
	paths := strings.Split(key, ".")

	cur := interface{}(map[string]interface{}(s))
	for _, path := range paths {
		if len(path) == 0 {
			return nil
		}

		if len(path) > 2 && path[0] == '[' && path[len(path)-1] == ']' {
			idx, err := strconv.Atoi(path[1 : len(path)-1])
			if err != nil {
				return nil
			}

			slice, err := cast.ToSliceE(cur)
			if err != nil {
				return nil
			}

			if idx >= len(slice) {
				return nil
			}

			cur = slice[idx]
		} else {
			m, err := cast.ToStringMapE(cur)
			if err != nil {
				return nil
			}

			cur, ok = m[path]
			if !ok {
				return nil
			}
		}

	}

	return cur
}
