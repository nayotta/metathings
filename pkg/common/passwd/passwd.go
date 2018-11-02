package passwd_helper

import "golang.org/x/crypto/bcrypt"

func MustParsePassword(x string) string {
	buf, _ := bcrypt.GenerateFromPassword([]byte(x), bcrypt.DefaultCost)
	return string(buf)
}

func ValidatePassword(hash, passwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwd)) == nil
}
