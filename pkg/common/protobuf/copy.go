package protobuf_helper

import "encoding/json"

func CopyExtra(x *string) map[string]string {
	if x == nil {
		return map[string]string{}
	}

	y := map[string]string{}
	if err := json.Unmarshal([]byte(*x), &y); err != nil {
		y = map[string]string{}
	}

	return y
}

func CopyString(x *string) string {
	if x == nil {
		return ""
	}

	return *x
}
