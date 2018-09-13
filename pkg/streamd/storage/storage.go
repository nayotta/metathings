package metathings_streamd_storage

import (
	"time"

	log "github.com/sirupsen/logrus"
)

type Stream struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	Name    *string  `gorm:"column:name"`
	OwnerId *string  `gorm:"column:owner_id"`
	State   *string  `gorm:"column:state"`
	Sources []Source `gorm:"-"`
	Groups  []Group  `gorm:"-"`
}

type Source struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	StreamId *string  `gorm:"column:stream_id"`
	Upstream Upstream `gorm:"-"`
}

type Group struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	StreamId *string  `gorm:"column:stream_id"`
	Inputs   []Input  `gorm:"-"`
	Outputs  []Output `gorm:"-"`
}

type Upstream struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	SourceId *string `gorm:"column:source_id"`
	Name     *string `gorm:"column:name"`
	Alias    *string `gorm:"column:alias"`
	Config   *string `gorm:"column:config"`
}

type Input struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	GroupId *string `gorm:"column:group_id"`
	Name    *string `gorm:"column:name"`
	Alias   *string `gorm:"column:alias"`
	Config  *string `gorm:"column:config"`
}

type Output struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	GroupId *string `gorm:"column:group_id"`
	Name    *string `gorm:"column:name"`
	Alias   *string `gorm:"column:alias"`
	Config  *string `gorm:"column:config"`
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
