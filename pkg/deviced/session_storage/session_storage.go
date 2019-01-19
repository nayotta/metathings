package metathings_deviced_session_storage

import "time"

type SessionStorage interface {
	GetStartupSession(id string) (int32, error)
	SetStartupSessionIfNotExists(id string, sess int32, expire time.Duration) error
	RefreshStartupSession(id string, expire time.Duration) error
}

func NewSessionStorage(driver, uri string, args ...interface{}) (SessionStorage, error) {
	return NewSessionStorage(driver, uri, args...)
}
