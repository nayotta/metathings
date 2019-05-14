package passwd_helper

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func MustParsePassword(x string) string {
	buf, _ := bcrypt.GenerateFromPassword([]byte(x), bcrypt.DefaultCost)
	return string(buf)
}

func ValidatePassword(hash, passwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwd)) == nil
}

func parse_hmac(key []byte, id string, timestamp time.Time, nonce int64) []byte {
	h := hmac.New(sha256.New, key)
	txt := fmt.Sprintf("%v%v%v", id, timestamp.UnixNano(), nonce)
	h.Write([]byte(txt))
	return h.Sum(nil)
}

func MustParseHmac(key, id string, timestamp time.Time, nonce int64) string {
	bkey, _ := base64.StdEncoding.DecodeString(key)
	return base64.StdEncoding.EncodeToString(parse_hmac(bkey, id, timestamp, nonce))
}

func ValidateHmac(hash, key, id string, timestamp time.Time, nonce int64) bool {
	bhash, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		return false
	}

	bkey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return false
	}

	bhmac := parse_hmac(bkey, id, timestamp, nonce)

	return bytes.Equal(bhash, bhmac)
}
