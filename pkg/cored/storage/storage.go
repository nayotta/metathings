package metathings_cored_storage

import (
	"time"

	log "github.com/sirupsen/logrus"
)

type Core struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	Name        *string    `gorm:"column:name"`
	ProjectId   *string    `gorm:"column:project_id"`
	OwnerId     *string    `gorm:"column:owner_id"`
	State       *string    `gorm:"column:state"`
	HeartbeatAt *time.Time `gorm:"column:heartbeat_at"`

	Entities []Entity `gorm:"-"`
}

type Entity struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	CoreId      *string    `gorm:"column:core_id"`
	Name        *string    `gorm:"column:name"`
	ServiceName *string    `gorm:"column:service_name"`
	Endpoint    *string    `gorm:"column:endpoint"`
	State       *string    `gorm:"column:state"`
	HeartbeatAt *time.Time `gorm:"column:heartbeat_at"`
}

func (Entity) TableName() string {
	return "entities"
}

type CoreApplicationCredentialMapping struct {
	CoreId                  *string `gorm:"column:core_id"`
	ApplicationCredentialId *string `gorm:"column:application_credential_id"`
}

type Storage interface {
	CreateCore(core Core) (Core, error)
	DeleteCore(id string) error
	PatchCore(id string, core Core) (Core, error)
	GetCore(id string) (Core, error)
	ListCores(Core) ([]Core, error)
	ListCoresForUser(owner_id string, core Core) ([]Core, error)
	AssignCoreToApplicationCredential(core_id string, app_cred_id string) error
	GetAssignedCoreFromApplicationCredential(app_cred_id string) (Core, error)

	CreateEntity(Entity) (Entity, error)
	DeleteEntity(id string) error
	PatchEntity(id string, entity Entity) (Entity, error)
	GetEntity(id string) (Entity, error)
	ListEntities(Entity) ([]Entity, error)
	ListEntitiesForCore(core_id string, entity Entity) ([]Entity, error)
}

func NewStorage(driver, uri string, logger log.FieldLogger) (Storage, error) {
	return newStorageImpl(driver, uri, logger)
}
