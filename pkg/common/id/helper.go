package common

import (
	"math/rand"
	"strings"

	uuid "github.com/satori/go.uuid"
)

func NewId() string {
	id, _ := uuid.NewV4()
	return strings.Replace(id.String(), "-", "", -1)
}

func NewUint64Id() uint64 {
	return rand.Uint64()
}
