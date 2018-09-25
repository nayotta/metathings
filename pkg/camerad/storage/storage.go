package metathings_camerad_storage

import (
	"time"

	log "github.com/sirupsen/logrus"
)

type Camera struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	Name                    *string `gorm:"column:name"`
	CoreId                  *string `gorm:"column:core_id"`
	EntityName              *string `gorm:"column:entity_name"`
	OwnerId                 *string `gorm:"column:owner_id"`
	ApplicationCredentialId *string `gorm:"column:application_credential_id"`
	State                   *string `gorm:"column:state"`

	Url       *string `gorm:"column:url"`
	Device    *string `gorm:"column:device"`
	Width     *uint32 `gorm:"width"`
	Height    *uint32 `gorm:"height"`
	Bitrate   *uint32 `gorm:"bitrate"`
	Framerate *uint32 `gorm:"framerate"`
}

type Storage interface {
	CreateCamera(Camera) (Camera, error)
	DeleteCamera(cam_id string) error
	PatchCamera(cam_id string, cam Camera) (Camera, error)
	GetCamera(cam_id string) (Camera, error)
	ListCameras(Camera) ([]Camera, error)
	ListCamerasForUser(owner_id string, cam Camera) ([]Camera, error)
}

func NewStorage(driver, uri string, logger log.FieldLogger) (Storage, error) {
	return newStorageImpl(driver, uri, logger)
}
