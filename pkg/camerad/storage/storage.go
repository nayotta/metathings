package metathings_camerad_storage

import (
	"time"

	log "github.com/sirupsen/logrus"
)

type Camera struct {
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`

	Id *string // Core Entity ID

}

type Storage interface {
}

func NewStorage(driver, uri string, logger log.FieldLogger) (Storage, error) {
	return nil, nil
}
