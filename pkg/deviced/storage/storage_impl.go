package metathings_deviced_storage

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type StorageImpl struct {
	db     *gorm.DB
	logger log.FieldLogger
}

func (self *StorageImpl) get_device(id string) (*Device, error) {
	panic("unimplemented")
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
	panic("unimplemented")
}

func (self *StorageImpl) PatchDevice(id string, device *Device) (*Device, error) {
	panic("unimplemented")
}

func (self *StorageImpl) GetDevice(id string) (*Device, error) {
	panic("unimplemented")
}

func (self *StorageImpl) ListDevices(*Device) ([]*Device, error) {
	panic("unimplemented")
}

func (self *StorageImpl) GetDeviceByEntityId(ent_id string) (*Device, error) {
	panic("unimplemented")
}

func (self *StorageImpl) CreateModule(*Module) (*Module, error) {
	panic("unimplemented")
}

func (self *StorageImpl) DeleteModule(id string) error {
	panic("unimplemented")
}

func (self *StorageImpl) PatchModule(id string, module *Module) (*Module, error) {
	panic("unimplemented")
}

func (self *StorageImpl) GetModule(id string) (*Module, error) {
	panic("unimplemented")
}

func (self *StorageImpl) ListModules(*Module) ([]*Module, error) {
	panic("unimplemented")
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

	if err = new_db(s, driver, uri); err != nil {
		return nil, err
	}

	if err = init_db(s); err != nil {
		return nil, err
	}

	return s, nil
}
