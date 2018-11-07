package metathings_deviced_storage

import (
	"time"

	identityd_storage "github.com/nayotta/metathings/pkg/identityd2/storage"
)

type Device struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time
	EntityId  *string `gorm:"column:entity_id"`

	Kind  *string `gorm:"column:kind"`
	State *string `gorm:"column:state"`
	Name  *string `gorm:"column:name"`
	Alias *string `gorm:"column:alias"`

	Modules []*Module `gorm:"-"`
	Entity  *identityd_storage.Entity
}

type Module struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time
	EntityId  *string `gorm:"column:entity_id"`

	State    *string `gorm:"column:state"`
	DeviceId *string `gorm:"column:device_id"`
	Endpoint *string `gorm:"column:endpoint"`
	Name     *string `gorm:"column:name"`
	Alias    *string `gorm:"column:alias"`

	Device *Device `gorm:"-"`
	Entity *identityd_storage.Entity
}

type Storage interface {
	CreateDevice(*Device) (*Device, error)
	DeleteDevice(id string) error
	PatchDevice(id string, device *Device) (*Device, error)
	GetDevice(id string) (*Device, error)
	ListDevices(*Device) ([]*Device, error)
	GetDeviceByEntityId(ent_id string) (*Device, error)

	CreateModule(*Module) (*Module, error)
	DeleteModule(id string) error
	PatchModule(id string, module *Module) (*Module, error)
	GetModule(id string) (*Module, error)
	ListModules(*Module) ([]*Module, error)
}

func NewStorage(driver, uri string, args ...interface{}) (Storage, error) {
	panic("unimplemented")
}
