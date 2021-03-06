package common

import (
	"strings"

	uuid "github.com/satori/go.uuid"

	rand_helper "github.com/nayotta/metathings/pkg/common/rand"
)

func NewId() string {
	return strings.Replace(uuid.NewV4().String(), "-", "", -1)
}

func NewUint64Id() uint64 {
	return rand_helper.Uint64()
}

func NewNamedId(x string) string {
	return strings.Replace(uuid.NewV5(uuid.Nil, x).String(), "-", "", -1)
}
