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
	Extra       *string `gorm:"column:extra"`

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

type FlowSet struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	Name  *string `gorm:"column:name"`
	Alias *string `gorm:"column:alias"`

	Flows []*Flow `gorm:"-"`
}

type FlowFlowSetMapping struct {
	CreatedAt time.Time

	FlowId    *string `gorm:"flow_id"`
	FlowSetId *string `gorm:"flow_set_id"`
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

	CreateFlowSet(context.Context, *FlowSet) (*FlowSet, error)
	DeleteFlowSet(ctx context.Context, id string) error
	PatchFlowSet(ctx context.Context, id string, flow_set *FlowSet) (*FlowSet, error)
	GetFlowSet(ctx context.Context, id string) (*FlowSet, error)
	ListFlowSets(context.Context, *FlowSet) ([]*FlowSet, error)
	ListViewFlowSetsByFlowId(ctx context.Context, id string) ([]*FlowSet, error)
	AddFlowToFlowSet(ctx context.Context, flow_set_id, flow_id string) error
	RemoveFlowFromFlowSet(ctx context.Context, flow_set_id, flow_id string) error
}

func NewStorage(driver, uri string, args ...interface{}) (Storage, error) {
	return NewStorageImpl(driver, uri, args...)
}
