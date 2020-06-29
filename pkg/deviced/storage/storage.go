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

	Modules     []*Module         `gorm:"-"`
	Flows       []*Flow           `gorm:"-"`
	Configs     []*Config         `gorm:"-"`
	ExtraHelper map[string]string `gorm:"-"`
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

type Config struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	Alias *string `gorm:"column:alias"`
	Body  *string `gorm:"column:body"`
}

type DeviceConfigMapping struct {
	CreatedAt time.Time

	DeviceId *string `gorm:"column:device_id"`
	ConfigId *string `gorm:"column:config_id"`
}

type FirmwareHub struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	Alias       *string `gorm:"column:alias"`
	Description *string `gorm:"column:description"`

	Devices             []*Device             `gorm:"-"`
	FirmwareDescriptors []*FirmwareDescriptor `gorm:"-"`
}

type DeviceFirmwareHubMapping struct {
	CreatedAt time.Time

	DeviceId      *string `gorm:"column:device_id"`
	FirmwareHubId *string `gorm:"column:firmware_hub_id"`
}

type FirmwareDescriptor struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	Name *string `gorm:"column:name"`

	FirmwareHubId *string `gorm:"column:firmware_hub_id"`
	ConfigId      *string `gorm:"column:config_id"`

	Config *Config `gorm:"-"`
}

type Storage interface {
	CreateDevice(context.Context, *Device) (*Device, error)
	DeleteDevice(ctx context.Context, id string) error
	PatchDevice(ctx context.Context, id string, device *Device) (*Device, error)
	GetDevice(ctx context.Context, id string) (*Device, error)
	ListDevices(context.Context, *Device) ([]*Device, error)
	GetDeviceByModuleId(ctx context.Context, id string) (*Device, error)

	CreateConfig(context.Context, *Config) (*Config, error)
	DeleteConfig(ctx context.Context, id string) error
	PatchConfig(ctx context.Context, id string, cfg *Config) (*Config, error)
	GetConfig(ctx context.Context, id string) (*Config, error)
	ListConfigs(context.Context, *Config) ([]*Config, error)
	AddConfigToDevice(ctx context.Context, dev_id, cfg_id string) error
	RemoveConfigFromDevice(ctx context.Context, dev_id, cfg_id string) error
	RemoveConfigFromDeviceByConfigId(ctx context.Context, cfg_id string) error
	ListConfigsByDeviceId(ctx context.Context, dev_id string) ([]*Config, error)

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

	CreateFirmwareHub(context.Context, *FirmwareHub) (*FirmwareHub, error)
	DeleteFirmwareHub(ctx context.Context, id string) error
	PatchFirmwareHub(ctx context.Context, id string, fh *FirmwareHub) (*FirmwareHub, error)
	GetFirmwareHub(ctx context.Context, id string) (*FirmwareHub, error)
	ListFirmwareHubs(ctx context.Context, frm_hub *FirmwareHub) ([]*FirmwareHub, error)
	AddDevicesToFirmwareHub(ctx context.Context, dev_ids []string, frm_hub_id string) error
	RemoveDevicesFromFirmwareHub(ctx context.Context, dev_ids []string, frm_hub_id string) error
	CreateFirmwareDescriptor(ctx context.Context, frm_desc *FirmwareDescriptor) error
	DeleteFirmwareDescriptor(ctx context.Context, frm_desc_id string) error
}

func NewStorage(driver, uri string, args ...interface{}) (Storage, error) {
	return NewStorageImpl(driver, uri, args...)
}
