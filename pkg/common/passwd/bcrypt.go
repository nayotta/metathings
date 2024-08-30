package passwd_helper

import "golang.org/x/crypto/bcrypt"

func MustParseBcrypt(passwd string) string {
	encrypted, _ := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	return string(encrypted)
}
