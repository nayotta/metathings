package metathings_camerad_storage

import (
	"time"

	log "github.com/sirupsen/logrus"
)

type Camera struct {
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`

	Id         *string
	Name       *string
	CoreId     *string `db:"core_id"`
	EntityName *string `db:"entity_name"`
	OwnerId    *string `db:"owner_id"`
	State      *string

	Url       *string
	Device    *string
	Width     *uint32
	Height    *uint32
	Bitrate   *uint32
	Framerate *uint32
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
