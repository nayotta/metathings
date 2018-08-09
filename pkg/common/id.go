package common

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

func NewId() string {
	return strings.Replace(uuid.NewV4().String(), "-", "", -1)
}
