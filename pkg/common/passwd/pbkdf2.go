package passwd_helper

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/rand"

	"golang.org/x/crypto/pbkdf2"
)

// from https://github.com/jpmens/mosquitto-auth-plug/blob/master/contrib/golang/mosquitto_pbkdf2.go

var (
	PBKDF2_RAND_LETTERS []byte
	PBKDF2_SEPARATOR    = "$"
	PBKDF2_TAG          = "PBKDF2"
	PBKDF2_ALGORITHM    = "sha256"
	PBKDF2_ITERATIONS   = 901
	PBKDF2_KEY_LENGTH   = 48
	PBKDF2_SALT_LENGTH  = 12
)

func MustParsePbkdf2(passwd string) string {
	salt := get_salt(PBKDF2_SALT_LENGTH)
	sha256_pwd := base64.StdEncoding.EncodeToString(pbkdf2.Key([]byte(passwd), salt, PBKDF2_ITERATIONS, PBKDF2_KEY_LENGTH, sha256.New))

	return fmt.Sprintf("%v%v%v%v%v%v%v%v%v",
		PBKDF2_TAG,
		PBKDF2_SEPARATOR,
		PBKDF2_ALGORITHM,
		PBKDF2_SEPARATOR,
		PBKDF2_ITERATIONS,
		PBKDF2_SEPARATOR,
		string(salt),
		PBKDF2_SEPARATOR,
		sha256_pwd)
}

func get_salt(length int) []byte {
	buf := make([]byte, length)

	for i := 0; i < length; i++ {
		buf[i] = PBKDF2_RAND_LETTERS[rand.Intn(len(PBKDF2_RAND_LETTERS))]
	}

	return buf
}

func init() {
	for i := 0; i < 10; i++ {
		PBKDF2_RAND_LETTERS = append(PBKDF2_RAND_LETTERS, byte(48+i))
	}
	for i := 0; i < 26; i++ {
		PBKDF2_RAND_LETTERS = append(PBKDF2_RAND_LETTERS, byte(65+i))
	}
	for i := 0; i < 26; i++ {
		PBKDF2_RAND_LETTERS = append(PBKDF2_RAND_LETTERS, byte(97+i))
	}
}
