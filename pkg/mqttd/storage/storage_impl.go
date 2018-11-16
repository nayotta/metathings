package metathingsmqttdstorage

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// StorageImpl storage
type StorageImpl struct {
	db     *gorm.DB
	logger log.FieldLogger
}

func (storImpl *StorageImpl) getDevice(id string) (*Device, error) {
	var err error
	var dev *Device

	if err = storImpl.db.First(&dev, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return dev, nil
}

func (storImpl *StorageImpl) listDevices(dev *Device) ([]*Device, error) {
	var err error
	var devsT []*Device

	d := &Device{}
	if dev.ID != nil {
		d.ID = dev.ID
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

	if err = storImpl.db.Select("id").Find(&devsT, d).Error; err != nil {
		return nil, err
	}

	var devs []*Device
	for _, d = range devsT {
		if d, err = storImpl.getDevice(*d.ID); err != nil {
			return nil, err
		}

		devs = append(devs, d)
	}

	return devs, nil
}

// CreateDevice create device
func (storImpl *StorageImpl) CreateDevice(dev *Device) (*Device, error) {
	var err error

	if err = storImpl.db.Create(dev).Error; err != nil {
		storImpl.logger.WithError(err).Debugf("failed to create device")
		return nil, err
	}

	if dev, err = storImpl.getDevice(*dev.ID); err != nil {
		storImpl.logger.WithError(err).Debugf("failed to get device")
		return nil, err
	}

	storImpl.logger.WithField("id", *dev.ID).Debugf("create device")

	return dev, nil
}

// DeleteDevice delete device
func (storImpl *StorageImpl) DeleteDevice(id string) error {
	if err := storImpl.db.Delete(&Device{}, "id = ?", id).Error; err != nil {
		storImpl.logger.WithError(err).Debugf("failed to delete device")
		return err
	}

	storImpl.logger.WithField("id", id).Debugf("delete device")

	return nil
}

// PatchDevice patch device
func (storImpl *StorageImpl) PatchDevice(id string, device *Device) (*Device, error) {
	var err error
	var dev Device

	if device.Alias != nil {
		dev.Alias = device.Alias
	}

	if device.State != nil {
		dev.State = device.State
	}

	if err = storImpl.db.Model(&Device{ID: &id}).Update(dev).Error; err != nil {
		storImpl.logger.WithError(err).Debugf("failed to patch device")
		return nil, err
	}

	if device, err = storImpl.getDevice(id); err != nil {
		storImpl.logger.WithError(err).Debugf("failed to get device")
		return nil, err
	}

	storImpl.logger.WithField("id", id).Debugf("patch device")

	return device, nil
}

// GetDevice get device
func (storImpl *StorageImpl) GetDevice(id string) (*Device, error) {
	var dev *Device
	var err error

	if dev, err = storImpl.getDevice(id); err != nil {
		storImpl.logger.WithError(err).Debugf("failed to get device")
		return nil, err
	}

	storImpl.logger.WithField("id", id).Debugf("get device")

	return dev, nil
}

// ListDevices list devices
func (storImpl *StorageImpl) ListDevices(dev *Device) ([]*Device, error) {
	var devs []*Device
	var err error

	if devs, err = storImpl.listDevices(dev); err != nil {
		storImpl.logger.WithError(err).Debugf("failed to list devices")
		return nil, err
	}

	storImpl.logger.Debugf("list devices")

	return devs, nil
}

func newDb(s *StorageImpl, driver, uri string) error {
	var db *gorm.DB
	var err error

	if db, err = gorm.Open(driver, uri); err != nil {
		return err
	}
	s.db = db

	return nil
}

func initDb(s *StorageImpl) error {
	s.db.AutoMigrate(
		&Device{},
	)

	return nil
}

// NewStorageImpl new storage impl
func NewStorageImpl(driver, uri string, args ...interface{}) (*StorageImpl, error) {
	var err error

	s := &StorageImpl{}

	if err = newDb(s, driver, uri); err != nil {
		return nil, err
	}

	if err = initDb(s); err != nil {
		return nil, err
	}

	return s, nil
}
