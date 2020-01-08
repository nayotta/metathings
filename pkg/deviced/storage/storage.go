package metathings_deviced_storage

import (
	"context"
	"time"
)

type Device struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	HeartbeatAt *time.Time
	Kind        *string `gorm:"column:kind"`
	State       *string `gorm:"column:state"`
	Name        *string `gorm:"column:name"`
	Alias       *string `gorm:"column:alias"`

	Modules []*Module `gorm:"-"`
	Flows   []*Flow   `gorm:"-"`
}

type Flow struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	DeviceId *string `gorm:"column:device_id"`
	Name     *string `gorm:"column:name"`
	Alias    *string `gorm:"column:alias"`

	Device *Device `gorm:"-"`
}

type Module struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	HeartbeatAt *time.Time
	State       *string `gorm:"column:state"`
	DeviceId    *string `gorm:"column:device_id"`
	Endpoint    *string `gorm:"column:endpoint"`
	Component   *string `gorm:"column:component"`
	Name        *string `gorm:"column:name"`
	Alias       *string `gorm:"column:alias"`

	Device *Device `gorm:"-"`
}

type Storage interface {
	CreateDevice(context.Context, *Device) (*Device, error)
	DeleteDevice(ctx context.Context, id string) error
	PatchDevice(ctx context.Context, id string, device *Device) (*Device, error)
	GetDevice(ctx context.Context, id string) (*Device, error)
	ListDevices(context.Context, *Device) ([]*Device, error)
	GetDeviceByModuleId(ctx context.Context, id string) (*Device, error)

	CreateModule(context.Context, *Module) (*Module, error)
	DeleteModule(ctx context.Context, id string) error
	PatchModule(ctx context.Context, id string, module *Module) (*Module, error)
	GetModule(ctx context.Context, id string) (*Module, error)
	ListModules(context.Context, *Module) ([]*Module, error)

	CreateFlow(context.Context, *Flow) (*Flow, error)
	DeleteFlow(ctx context.Context, id string) error
	PatchFlow(ctx context.Context, id string, flow *Flow) (*Flow, error)
	GetFlow(ctx context.Context, id string) (*Flow, error)
	ListFlows(context.Context, *Flow) ([]*Flow, error)
}

func NewStorage(driver, uri string, args ...interface{}) (Storage, error) {
	return NewStorageImpl(driver, uri, args...)
}
