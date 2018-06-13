package metathings_core_storage

import (
	"time"

	log "github.com/sirupsen/logrus"
)

type Core struct {
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`

	Id          *string
	Name        *string
	ProjectId   *string `db:"project_id"`
	OwnerId     *string `db:"owner_id"`
	State       *string
	HeartbeatAt *time.Time `db:"heartbeat_at"`
}

type Entity struct {
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`

	Id          *string
	CoreId      *string `db:"core_id"`
	Name        *string
	ServiceName *string `db:"service_name"`
	Endpoint    *string
	State       *string
	HeartbeatAt *time.Time `db:"heartbeat_at"`
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

func NewStorage(driver, uri string, logger log.FieldLogger) (Storage, error) {
	return newStorageImpl(driver, uri, logger)
}
