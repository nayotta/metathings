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

type Entity struct {
	sql_helper.Metadata
	Id          *string
	CoreId      *string
	Name        *string
	ServiceName *string `db:"service_name"`
	Endpoint    *string
	State       *string
}

type Storage interface {
	CreateCore(core Core) (Core, error)
	DeleteCore(core_id string) error
	PatchCore(core_id string, core Core) (Core, error)
	GetCore(core_id string) (Core, error)
	ListCores(Core) ([]Core, error)
	ListCoresForUser(owner_id string, core Core) ([]Core, error)
	AssignCoreToApplicationCredential(core_id string, app_cred_id string) error
	GetAssignedCoreFromApplicationCredential(app_cred_id string) (Core, error)

	CreateEntity(Entity) (Entity, error)
	DeleteEntity(entity_id string) error
	PatchEntity(entity_id string, entity Entity) (Entity, error)
	GetEntity(entity_id string) (Entity, error)
	ListEntities(Entity) ([]Entity, error)
	ListEntitiesForCore(core_id string, entity Entity) ([]Entity, error)
}

func NewStorage(dbpath string, logger log.FieldLogger) (Storage, error) {
	return newStorageImpl(dbpath, logger)
}
