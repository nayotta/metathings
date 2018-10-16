package common

import uuid "github.com/satori/go.uuid"

func NewId() string {
	uuid_ret, _ := uuid.NewV4()
	return uuid_ret.String()
}
