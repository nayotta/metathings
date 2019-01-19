package metathings_deviced_session_storage

import (
	"time"
)

type SessionStorageImpl struct {
}

func (self *SessionStorageImpl) GetStartupSession(id string) (int32, error) {
	panic("unimplemented")
}

func (self *SessionStorageImpl) SetStartupSessionIfNotExists(id string, sess int32, expire time.Duration) error {
	panic("unimplemented")
}

func (self *SessionStorageImpl) RefreshStartupSession(id string, expire time.Duration) error {
	panic("unimplemented")
}

func NewSessionStorageImpl(driver, uri string, args ...interface{}) (*SessionStorageImpl, error) {
	panic("unimplemented")
}
