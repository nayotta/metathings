package metathings_deviced_storage

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type StorageImpl struct {
	db     *gorm.DB
	logger log.FieldLogger
}

func (self *StorageImpl) get_flow(id string) (*Flow, error) {
	var err error
	flw := &Flow{}

	if err = self.db.First(flw, "id = ?", id).Error; err != nil {
		return nil, err
	}

	flw.Device = &Device{Id: flw.DeviceId}

	return flw, nil
}

func (self *StorageImpl) list_flows_by_device_id(id string) ([]*Flow, error) {
	var err error
	var flws_t []*Flow

	if err = self.db.Select("id").Find(&flws_t, "device_id = ?", id).Error; err != nil {
		return nil, err
	}

	flws := []*Flow{}
	for _, f := range flws_t {
		if f, err = self.get_flow(*f.Id); err != nil {
			return nil, err
		}

		flws = append(flws, f)
	}

	return flws, nil
}

func (self *StorageImpl) list_flows(flw *Flow) ([]*Flow, error) {
	var err error
	var flws_t []*Flow

	f := &Flow{}
	if flw.Id != nil {
		f.Id = flw.Id
	}

	if flw.DeviceId != nil {
		f.DeviceId = flw.DeviceId
	}

	if flw.Name != nil {
		f.Name = flw.Name
	}

	if flw.Alias != nil {
		f.Alias = flw.Alias
	}

	if err = self.db.Select("id").Find(&flws_t, f).Error; err != nil {
		return nil, err
	}

	var flws []*Flow
	for _, f = range flws_t {
		if f, err = self.get_flow(*f.Id); err != nil {
			return nil, err
		}

		flws = append(flws, f)
	}

	return flws, nil
}

func (self *StorageImpl) get_module(id string) (*Module, error) {
	var err error
	mdl := &Module{}

	if err = self.db.First(mdl, "id = ?", id).Error; err != nil {
		return nil, err
	}

	mdl.Device = &Device{Id: mdl.DeviceId}

	return mdl, nil
}

func (self *StorageImpl) list_modules_by_device_id(id string) ([]*Module, error) {
	var err error
	var mdls_t []*Module

	if err = self.db.Select("id").Find(&mdls_t, "device_id = ?", id).Error; err != nil {
		return nil, err
	}

	mdls := []*Module{}
	for _, m := range mdls_t {
		if m, err = self.get_module(*m.Id); err != nil {
			return nil, err
		}

		mdls = append(mdls, m)
	}

	return mdls, nil
}

func (self *StorageImpl) list_modules(mdl *Module) ([]*Module, error) {
	var err error
	var mdls_t []*Module

	m := &Module{}
	if mdl.Id != nil {
		m.Id = mdl.Id
	}

	if mdl.State != nil {
		m.State = mdl.State
	}

	if mdl.DeviceId != nil {
		m.DeviceId = mdl.DeviceId
	}

	if mdl.Endpoint != nil {
		m.Endpoint = mdl.Endpoint
	}

	if mdl.Component != nil {
		m.Component = mdl.Component
	}

	if mdl.Name != nil {
		m.Name = mdl.Name
	}

	if mdl.Alias != nil {
		m.Alias = mdl.Alias
	}

	if err = self.db.Select("id").Find(&mdls_t, m).Error; err != nil {
		return nil, err
	}

	var mdls []*Module
	for _, m = range mdls_t {
		if m, err = self.get_module(*m.Id); err != nil {
			return nil, err
		}

		mdls = append(mdls, m)
	}

	return mdls, nil
}

func (self *StorageImpl) internal_get_device(dev *Device) (*Device, error) {
	var err error

	if dev.Modules, err = self.list_modules_by_device_id(*dev.Id); err != nil {
		return nil, err
	}

	if dev.Flows, err = self.list_flows_by_device_id(*dev.Id); err != nil {
		return nil, err
	}

	return dev, nil
}

func (self *StorageImpl) get_device(id string) (*Device, error) {
	var err error
	dev := &Device{}

	if err = self.db.First(dev, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if dev, err = self.internal_get_device(dev); err != nil {
		return nil, err
	}

	return dev, nil
}

func (self *StorageImpl) get_device_by_module_id(id string) (*Device, error) {
	var err error
	var mdl Module

	if err = self.db.Select("device_id").First(&mdl, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return self.get_device(*mdl.DeviceId)
}

func (self *StorageImpl) list_devices(dev *Device) ([]*Device, error) {
	var err error
	var devs_t []*Device

	d := &Device{}
	if dev.Id != nil {
		d.Id = dev.Id
	}

	if dev.Kind != nil {
		d.Kind = dev.Kind
	}

	if dev.State != nil {
		d.State = dev.State
	}

	if dev.Name != nil {
		d.Name = dev.Name
	}

	if dev.Alias != nil {
		d.Alias = dev.Alias
	}

	if err = self.db.Select("id").Find(&devs_t, d).Error; err != nil {
		return nil, err
	}

	var devs []*Device
	for _, d = range devs_t {
		if d, err = self.get_device(*d.Id); err != nil {
			return nil, err
		}

		devs = append(devs, d)
	}

	return devs, nil
}

func (self *StorageImpl) CreateDevice(dev *Device) (*Device, error) {
	var err error

	if err = self.db.Create(dev).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to create device")
		return nil, err
	}

	if dev, err = self.get_device(*dev.Id); err != nil {
		self.logger.WithError(err).Debugf("failed to get device")
		return nil, err
	}

	self.logger.WithField("id", *dev.Id).Debugf("create device")

	return dev, nil
}

func (self *StorageImpl) DeleteDevice(id string) error {
	if err := self.db.Delete(&Device{}, "id = ?", id).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to delete device")
		return err
	}

	self.logger.WithField("id", id).Debugf("delete device")

	return nil
}

func (self *StorageImpl) PatchDevice(id string, device *Device) (*Device, error) {
	var err error
	var dev Device

	if device.Alias != nil {
		dev.Alias = device.Alias
	}

	if device.State != nil {
		dev.State = device.State
	}

	if device.HeartbeatAt != nil {
		dev.HeartbeatAt = device.HeartbeatAt
	}

	if err = self.db.Model(&Device{Id: &id}).Update(dev).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to patch device")
		return nil, err
	}

	if device, err = self.get_device(id); err != nil {
		self.logger.WithError(err).Debugf("failed to get device")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("patch device")

	return device, nil
}

func (self *StorageImpl) GetDevice(id string) (*Device, error) {
	var dev *Device
	var err error

	if dev, err = self.get_device(id); err != nil {
		self.logger.WithError(err).Debugf("failed to get device")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("get device")

	return dev, nil
}

func (self *StorageImpl) ListDevices(dev *Device) ([]*Device, error) {
	var devs []*Device
	var err error

	if devs, err = self.list_devices(dev); err != nil {
		self.logger.WithError(err).Debugf("failed to list devices")
		return nil, err
	}

	self.logger.Debugf("list devices")

	return devs, nil
}

func (self *StorageImpl) GetDeviceByModuleId(id string) (*Device, error) {
	var dev *Device
	var err error

	if dev, err = self.get_device_by_module_id(id); err != nil {
		self.logger.WithError(err).Debugf("failed to get device by module id")
		return nil, err
	}

	return dev, nil
}

func (self *StorageImpl) CreateModule(mdl *Module) (*Module, error) {
	var err error

	if err = self.db.Create(mdl).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to create module")
		return nil, err
	}

	if mdl, err = self.get_module(*mdl.Id); err != nil {
		self.logger.WithError(err).Debugf("failed to get module")
		return nil, err
	}

	self.logger.WithField("id", *mdl.Id).Debugf("create module")

	return mdl, nil
}

func (self *StorageImpl) DeleteModule(id string) error {
	var err error

	if err = self.db.Delete(&Module{}, "id = ?", id).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to delete module")
		return err
	}

	self.logger.WithField("id", id).Debugf("delete module")

	return nil
}

func (self *StorageImpl) PatchModule(id string, mdl *Module) (*Module, error) {
	var err error
	var m Module

	if mdl.Alias != nil {
		m.Alias = mdl.Alias
	}

	if mdl.State != nil {
		m.State = mdl.State
	}

	if mdl.HeartbeatAt != nil {
		m.HeartbeatAt = mdl.HeartbeatAt
	}

	if err = self.db.Model(&Module{Id: &id}).Update(m).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to patch module")
		return nil, err
	}

	if mdl, err = self.get_module(id); err != nil {
		self.logger.WithError(err).Debugf("failed to get module")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("patch module")

	return mdl, nil
}

func (self *StorageImpl) GetModule(id string) (*Module, error) {
	var mdl *Module
	var err error

	if mdl, err = self.get_module(id); err != nil {
		self.logger.WithError(err).Debugf("failed to get module")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("get module")

	return mdl, nil
}

func (self *StorageImpl) ListModules(mdl *Module) ([]*Module, error) {
	var mdls []*Module
	var err error

	if mdls, err = self.list_modules(mdl); err != nil {
		self.logger.WithError(err).Debugf("failed to list modules")
		return nil, err
	}

	self.logger.Debugf("list modules")

	return mdls, nil
}

func (self *StorageImpl) CreateFlow(flw *Flow) (*Flow, error) {
	var err error

	if err = self.db.Create(flw).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to create flow")
		return nil, err
	}

	if flw, err = self.get_flow(*flw.Id); err != nil {
		self.logger.WithError(err).Debugf("failed to get flow")
		return nil, err
	}

	self.logger.WithField("id", *flw.Id).Debugf("create flow")

	return flw, nil
}

func (self *StorageImpl) DeleteFlow(id string) error {
	var err error

	if err = self.db.Delete(&Flow{}, "id = ?", id).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to delete flow")
		return err
	}

	self.logger.WithField("id", id).Debugf("delete flow")

	return nil
}

func (self *StorageImpl) PatchFlow(id string, flw *Flow) (*Flow, error) {
	var err error
	var f Flow

	if flw.Alias != nil {
		f.Alias = flw.Alias
	}

	if err = self.db.Model(&Flow{Id: &id}).Update(f).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to patch flow")
		return nil, err
	}

	if flw, err = self.get_flow(id); err != nil {
		self.logger.WithError(err).Debugf("failed to get flow")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("patch flow")

	return flw, nil
}

func (self *StorageImpl) GetFlow(id string) (*Flow, error) {
	var flw *Flow
	var err error

	if flw, err = self.get_flow(id); err != nil {
		self.logger.WithError(err).Debugf("failed to get flow")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("get flow")

	return flw, nil
}

func (self *StorageImpl) ListFlows(flw *Flow) ([]*Flow, error) {
	var flws []*Flow
	var err error

	if flws, err = self.list_flows(flw); err != nil {
		self.logger.WithError(err).Debugf("failed to list flows")
		return nil, err
	}

	self.logger.Debugf("list flows")

	return flws, nil
}

func init_args(s *StorageImpl, args ...interface{}) error {
	var key string
	var ok bool

	if len(args)%2 != 0 {
		return InvalidArgument
	}

	for i := 0; i < len(args); i += 2 {
		key, ok = args[i].(string)
		if !ok {
			return InvalidArgument
		}

		switch key {
		case "logger":
			s.logger, ok = args[i+1].(log.FieldLogger)
			if !ok {
				return InvalidArgument
			}
		}
	}

	return nil
}

func new_db(s *StorageImpl, driver, uri string) error {
	var db *gorm.DB
	var err error

	if db, err = gorm.Open(driver, uri); err != nil {
		return err
	}
	s.db = db

	return nil
}

func init_db(s *StorageImpl) error {
	s.db.AutoMigrate(
		&Device{},
		&Module{},
		&Flow{},
	)

	return nil
}

func NewStorageImpl(driver, uri string, args ...interface{}) (*StorageImpl, error) {
	var err error

	s := &StorageImpl{}
	if err = init_args(s, args...); err != nil {
		return nil, err
	}

	if err = new_db(s, driver, uri); err != nil {
		return nil, err
	}

	if err = init_db(s); err != nil {
		return nil, err
	}

	return s, nil
}
