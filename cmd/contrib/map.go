package cmd_contrib

import "strings"

func ParseDepth1StringMap(in map[string]any) map[string]any {
	out := map[string]any{}

	for k, v := range in {
		ks := strings.Split(k, ".")
		if len(ks) != 2 {
			panic("unexpected key: " + k)
		}
		ko, ki := ks[0], ks[1]

		m, ok := out[ko].(map[string]any)
		if !ok {
			m = map[string]any{}
			out[ko] = m
		}
		m[ki] = v
	}

	return out
}
