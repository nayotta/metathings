package metathingsmqttdconnection

import (
	"math/rand"
	"strings"
)

func generateSession() int32 {
	return rand.Int31()
}

func getTopicDeviceID(topic string) string {
	strs := strings.Split(topic, "/")
	if len(strs) > 0 {
		return strs[0]
	}

	return ""
}

func getTopicType(topic string) string {
	strs := strings.Split(topic, "/")
	if len(strs) > 1 {
		return strs[1]
	}

	return ""
}
