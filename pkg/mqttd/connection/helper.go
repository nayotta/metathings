package metathingsmqttdconnection

import "math/rand"

func generateSession() int32 {
	return rand.Int31()
}
