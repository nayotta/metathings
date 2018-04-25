package sql_helper

import "time"

type Metadata struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (md Metadata) CreatedAtNow() {
	md.CreatedAt = time.Now()
}

func (md Metadata) UpdatedAtNow() {
	md.UpdatedAt = time.Now()
}

func (md Metadata) InitializedAtNow() {
	md.CreatedAtNow()
	md.UpdatedAtNow()
}
