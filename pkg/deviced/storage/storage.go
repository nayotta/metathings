package metathings_deviced_storage

import (
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
	CreateDevice(*Device) (*Device, error)
	DeleteDevice(id string) error
	PatchDevice(id string, device *Device) (*Device, error)
	GetDevice(id string) (*Device, error)
	ListDevices(*Device) ([]*Device, error)
	GetDeviceByModuleId(id string) (*Device, error)

	CreateModule(*Module) (*Module, error)
	DeleteModule(id string) error
	PatchModule(id string, module *Module) (*Module, error)
	GetModule(id string) (*Module, error)
	ListModules(*Module) ([]*Module, error)

	CreateFlow(*Flow) (*Flow, error)
	DeleteFlow(id string) error
	PatchFlow(id string, flow *Flow) (*Flow, error)
	GetFlow(id string) (*Flow, error)
	ListFlows(*Flow) ([]*Flow, error)

	CreateFlowSet(*FlowSet) (*FlowSet, error)
	DeleteFlowSet(id string) error
	PatchFlowSet(id string, flow_set *FlowSet) (*FlowSet, error)
	GetFlowSet(id string) (*FlowSet, error)
	ListFlowSets(*FlowSet) ([]*FlowSet, error)
	AddFlowToFlowSet(flow_set_id, flow_id string) error
	RemoveFlowFromFlowSet(flow_set_id, flow_id string) error
}

func NewStorage(driver, uri string, args ...interface{}) (Storage, error) {
	return NewStorageImpl(driver, uri, args...)
}
