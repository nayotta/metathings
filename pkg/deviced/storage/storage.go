package metathings_deviced_storage

import (
	"fmt"
	"time"
)

type Device struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	Kind  *string `gorm:"column:kind"`
	State *string `gorm:"column:state"`
	Name  *string `gorm:"column:name"`
	Alias *string `gorm:"column:alias"`

	Modules []*Module `gorm:"-"`
}

type Module struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	State    *string `gorm:"column:state"`
	DeviceId *string `gorm:"column:device_id"`
	Endpoint *string `gorm:"column:endpoint"`
	Name     *string `gorm:"column:name"`
	Alias    *string `gorm:"column:alias"`

	Device *Device `gorm:"-"`
}

func (self *Module) String() string {
	return fmt.Sprintf("{id=%v created_at=%v updated_at=%v state=%v device_id=%v endpoint=%v name=%v alias=%v}", *self.Id, self.CreatedAt, self.UpdatedAt, *self.State, *self.DeviceId, *self.Endpoint, *self.Name, *self.Alias)
}

type Storage interface {
	CreateDevice(*Device) (*Device, error)
	DeleteDevice(id string) error
	PatchDevice(id string, device *Device) (*Device, error)
	GetDevice(id string) (*Device, error)
	ListDevices(*Device) ([]*Device, error)

	CreateModule(*Module) (*Module, error)
	DeleteModule(id string) error
	PatchModule(id string, module *Module) (*Module, error)
	GetModule(id string) (*Module, error)
	ListModules(*Module) ([]*Module, error)
}

func NewStorage(driver, uri string, args ...interface{}) (Storage, error) {
	return NewStorageImpl(driver, uri, args...)
}
