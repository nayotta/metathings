package metathings_deviced_storage

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type StorageImpl struct {
	db     *gorm.DB
	logger log.FieldLogger
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

func (self *StorageImpl) PatchModule(id string, module *Module) (*Module, error) {
	var err error
	var mdl Module

	if module.Alias != nil {
		mdl.Alias = module.Alias
	}

	if module.State != nil {
		mdl.State = module.State
	}

	if module.HeartbeatAt != nil {
		mdl.HeartbeatAt = module.HeartbeatAt
	}

	if err = self.db.Model(&Module{Id: &id}).Update(mdl).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to patch module")
		return nil, err
	}

	if module, err = self.get_module(id); err != nil {
		self.logger.WithError(err).Debugf("failed to get module")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("patch module")

	return module, nil
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
