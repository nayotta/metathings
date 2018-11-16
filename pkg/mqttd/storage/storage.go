package metathingsmqttdstorage

import (
	"time"
)

// Device device struct
type Device struct {
	ID        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	State *string `gorm:"column:state"`
	Name  *string `gorm:"column:name"`
	Alias *string `gorm:"column:alias"`
}

// Storage storage interface
type Storage interface {
	CreateDevice(*Device) (*Device, error)
	DeleteDevice(id string) error
	PatchDevice(id string, device *Device) (*Device, error)
	GetDevice(id string) (*Device, error)
	ListDevices(*Device) ([]*Device, error)
}

// NewStorage new storage
func NewStorage(driver, uri string, args ...interface{}) (Storage, error) {
	return NewStorageImpl(driver, uri, args...)
}
