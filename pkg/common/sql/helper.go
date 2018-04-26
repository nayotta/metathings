package sql_helper

import "time"

type Metadata struct {
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
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
