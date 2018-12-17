package metathingsdevicecloudmqttbridge

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
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

// EncodeDownPath EncodeDownPath
func EncodeDownPath(deviceID string) string {
	return fmt.Sprintf("%s/down/", deviceID)
}

// newSessionID
func newSessionID() (string, int) {
	rand.Seed(time.Now().UnixNano())
	sessionID := rand.Intn(899999) + 100000

	return strconv.Itoa(sessionID), sessionID
}
