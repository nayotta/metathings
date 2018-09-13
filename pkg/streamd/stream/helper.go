package stream_manager

import "strings"

func split_and_trim(x string) []string {
	y := []string{}
	for _, t := range strings.Split(x, ",") {
		y = append(y, strings.Trim(t, " "))
	}
	return y
}

func group_by_prefix(x map[string]string, prefix string) map[string]string {
	y := map[string]string{}

	for k, v := range x {
		if strings.HasPrefix(k, prefix) {
			y[strings.Trim(k, prefix)] = v
		}
	}

	return y
}
