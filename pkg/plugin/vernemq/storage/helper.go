package metathings_plugin_vernemq_storage

import "encoding/json"

func ParseVernemqRedisKey(clientid, username string) string {
	keySlice := []any{"", clientid, username}
	buf, _ := json.Marshal(keySlice)
	return string(buf)
}
