package metathings_camerad_storage

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"
)

var (
	empty_camera = Camera{}
)

type storageImpl struct {
	db     *gorm.DB
	logger log.FieldLogger
}

func (self *storageImpl) get_camera(id string) (Camera, error) {
	var cam Camera
	err := self.db.Where("id = ?", id).First(&cam).Error
	if err != nil {
		return empty_camera, err
	}

	return cam, nil
}

func (self *storageImpl) CreateCamera(cam Camera) (Camera, error) {
	err := self.db.Create(&cam).Error
	if err != nil {
		return empty_camera, err
	}

	cam, err = self.get_camera(*cam.Id)
	if err != nil {
		return empty_camera, err
	}

	self.logger.WithField("id", *cam.Id).Debugf("create camera")
	return cam, nil
}

func (self *storageImpl) DeleteCamera(id string) error {
	err := self.db.Delete(&Camera{}, "id = ?", id).Error
	if err != nil {
		return err
	}

	self.logger.WithField("id", id).Debugf("delete camera")
	return nil
}

func (self *storageImpl) PatchCamera(cam_id string, cam Camera) (Camera, error) {
	var c Camera

	if cam.Name != nil {
		c.Name = cam.Name
	}

	if cam.State != nil {
		c.State = cam.State
	}

	if cam.Url != nil {
		c.Url = cam.Url
	}

	if cam.Device != nil {
		c.Device = cam.Device
	}

	if cam.Width != nil && cam.Height != nil {
		c.Width = cam.Width
		c.Height = cam.Height
	}

	if cam.Bitrate != nil {
		c.Bitrate = cam.Bitrate
	}

	if cam.Framerate != nil {
		c.Framerate = cam.Framerate
	}

	err := self.db.Model(&Camera{}).Where("id = ?", cam_id).Patchs(c).Error
	if err != nil {
		return empty_camera, err
	}

	cam, err = self.get_camera(cam_id)
	if err != nil {
		return empty_camera, err
	}

	self.logger.WithField("id", cam_id).Debugf("patch camera")
	return cam, nil
}

func (self *storageImpl) GetCamera(cam_id string) (Camera, error) {
	cam, err := self.get_camera(cam_id)
	if err != nil {
		return empty_camera, err
	}

	self.logger.WithField("id", cam_id).Debugf("get camera")
	return cam, nil
}

func (self *storageImpl) list_cameras(cam Camera) ([]Camera, error) {
	var cams []Camera
	err := self.db.Find(cams, &cam).Error
	if err != nil {
		return nil, err
	}

	return cams, nil
}

func (self *storageImpl) ListCameras(cam Camera) ([]Camera, error) {
	cams, err := self.list_cameras(cam)
	if err != nil {
		return nil, err
	}

	self.logger.Debugf("list cameras")
	return cams, nil
}

func (self *storageImpl) ListCamerasForUser(owner_id string, cam Camera) ([]Camera, error) {
	cam.OwnerId = &owner_id
	cams, err := self.list_cameras(cam)
	if err != nil {
		return nil, err
	}

	self.logger.Debugf("list cameras for user")
	return cams, nil
}

func newStorageImpl(driver, uri string, logger log.FieldLogger) (*storageImpl, error) {
	db, err := gorm.Open(driver, uri)
	if err != nil {
		return nil, err
	}

	return &storageImpl{
		logger: logger.WithField("#module", "storage"),
		db:     db,
	}, nil
}
