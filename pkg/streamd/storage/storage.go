package metathings_streamd_storage

import (
	"time"

	log "github.com/sirupsen/logrus"
)

type Stream struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	Name    *string `gorm:"column:name"`
	OwnerId *string `gorm:"column:owner_id"`
	State   *string `gorm:"column:state"`
	Pads    []Pad   `gorm:"-"`
}

type Pad struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	Name       *string `gorm:"column:name"`
	Type       *string `gorm:"column:type"`
	Processor  *string `gorm:"column:processor"`
	StreamId   *string `gorm:"column:stream_id"`
	NextPadIds *string `gorm:"column:next_pad_ids"`
	Config     *string `gorm:"column:config"`
}

type Storage interface {
	CreateStream(Stream) (Stream, error)
	DeleteStream(stm_id string) error
	PatchStream(stm_id string, stm Stream) (Stream, error)
	GetStream(stm_id string) (Stream, error)
	ListStreams(Stream) ([]Stream, error)
	ListStreamsForUser(owner_id string, stm Stream) ([]Stream, error)
}

func NewStorage(driver, uri string, logger log.FieldLogger) (Storage, error) {
	panic("unimplemented")
}
