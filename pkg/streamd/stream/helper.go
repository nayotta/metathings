package stream_manager

import "strings"

func split_and_trim(x string) []string {
	ys := []string{}
	for _, t := range strings.Split(x, ",") {
		y := strings.Trim(t, " ")
		if y != "" {
			ys = append(ys, y)
		}
	}
	return ys
}

func group_by_prefix(x map[string]string, prefix string) map[string]string {
	y := map[string]string{}

	for k, v := range x {
		if strings.HasPrefix(k, prefix) {
			y[k[len(prefix):]] = v
		}
	}

	return y
}
