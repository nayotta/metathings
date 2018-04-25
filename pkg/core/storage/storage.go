package metathings_core_storage

import (
	log "github.com/sirupsen/logrus"

	sql_helper "github.com/bigdatagz/metathings/pkg/common/sql"
)

type Core struct {
	sql_helper.Metadata
	Id        *string
	Name      *string
	ProjectId *string `db:"project_id"`
	OwnerId   *string `db:"owner_id"`
	State     *string
}

type Storage interface {
	CreateCore(core Core) (Core, error)
	DeleteCore(core_id string) error
	PatchCore(core_id string, core Core) (Core, error)
	GetCore(core_id string) (Core, error)
	ListCores(Core) ([]Core, error)
	AssignCoreToApplicationCredential(core_id string, app_cred_id string) error
	GetAssignedCoreFromApplicationCredential(app_cred_id string) (Core, error)
}

func NewStorage(dbpath string, logger log.FieldLogger) (Storage, error) {
	return newStorageImpl(dbpath, logger)
}
