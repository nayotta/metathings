package nonce_helper

import "math/rand"

var (
	NONCE_LENGTH  = 8
	NONCE_LETTERS = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func GenerateNonce() string {
	buf := make([]byte, NONCE_LENGTH)
	for i := 0; i < NONCE_LENGTH; i++ {
		buf[i] = NONCE_LETTERS[rand.Intn(len(NONCE_LETTERS))]
	}
	return string(buf)
}
