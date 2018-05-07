package common

import uuid "github.com/satori/go.uuid"

func NewId() string {
	return uuid.NewV4().String()
}
