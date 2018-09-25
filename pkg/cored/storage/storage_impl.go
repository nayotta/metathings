package metathings_cored_storage

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"
)

var (
	empty_core   = Core{}
	empty_entity = Entity{}
)

type storageImpl struct {
	db     *gorm.DB
	logger log.FieldLogger
}

func (self *storageImpl) get_core(id string) (Core, error) {
	var core Core
	err := self.db.Where("id = ?", id).First(&core).Error
	if err != nil {
		return empty_core, err
	}

	core.Entities, err = self.get_entities_by_core_id(id)
	if err != nil {
		return empty_core, err
	}

	return core, nil
}

func (self *storageImpl) get_entities_by_core_id(core_id string) ([]Entity, error) {
	var entities []Entity
	err := self.db.Where("core_id = ?", core_id).Find(&entities).Error
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func (self *storageImpl) CreateCore(core Core) (Core, error) {
	now := time.Now()
	core.HeartbeatAt = &now

	err := self.db.Create(&core).Error
	if err != nil {
		return empty_core, err
	}

	core, err = self.get_core(*core.Id)
	if err != nil {
		return empty_core, err
	}

	self.logger.WithField("id", *core.Id).Debugf("create core")
	return core, nil
}

func (self *storageImpl) tx_delete_entities_by_core_id(tx *gorm.DB, core_id string) {
	tx.Delete(&Entity{}, "core_id = ?", core_id)
}

func (self *storageImpl) tx_delete_core(tx *gorm.DB, id string) {
	tx.Delete(&Core{}, "id = ?", id)
	self.tx_delete_entities_by_core_id(tx, id)
}

func (self *storageImpl) DeleteCore(id string) error {
	tx := self.db.Begin()
	self.tx_delete_core(tx, id)
	err := tx.Commit().Error
	if err != nil {
		return err
	}

	self.logger.WithField("id", id).Debugf("delete core")
	return nil
}

func (self *storageImpl) PatchCore(id string, core Core) (Core, error) {
	var c Core

	if core.Name != nil {
		c.Name = core.Name
	}

	if core.State != nil {
		c.State = core.State
	}

	if core.HeartbeatAt != nil {
		c.HeartbeatAt = core.HeartbeatAt
	}

	err := self.db.Model(&Core{}).Where("id = ?", id).Updates(c).Error
	if err != nil {
		return empty_core, err
	}

	core, err = self.get_core(id)
	if err != nil {
		return empty_core, err
	}

	self.logger.WithField("id", id).Debugf("patch core")
	return c, nil
}

func (self *storageImpl) GetCore(id string) (Core, error) {
	core, err := self.get_core(id)
	if err != nil {
		return empty_core, err
	}

	self.logger.WithField("id", id).Debugf("get core")
	return core, nil
}

func (self *storageImpl) list_cores(core Core) ([]Core, error) {
	var cores_t []Core
	err := self.db.Select("id").Find(&cores_t, core).Error
	if err != nil {
		return nil, err
	}

	var cores []Core
	for _, c := range cores_t {
		core, err := self.get_core(*c.Id)
		if err != nil {
			return nil, err
		}
		cores = append(cores, core)
	}

	return cores, nil
}

func (self *storageImpl) ListCores(core Core) ([]Core, error) {
	cores, err := self.list_cores(core)
	if err != nil {
		return nil, err
	}

	self.logger.Debugf("list cores")
	return cores, nil
}

func (self *storageImpl) ListCoresForUser(owner_id string, core Core) ([]Core, error) {
	core.OwnerId = &owner_id
	cores, err := self.list_cores(core)
	if err != nil {
		return nil, err
	}

	self.logger.Debugf("list cores for user")
	return cores, nil
}

func (self *storageImpl) AssignCoreToApplicationCredential(core_id string, app_cred_id string) error {
	m := CoreApplicationCredentialMapping{
		CoreId:                  &core_id,
		ApplicationCredentialId: &app_cred_id,
	}

	err := self.db.Create(&m).Error
	if err != nil {
		return err
	}

	self.logger.WithFields(log.Fields{
		"core_id":                   core_id,
		"application_credential_id": app_cred_id,
	}).Debugf("assign core to application credential")
	return nil
}

func (self *storageImpl) GetAssignedCoreFromApplicationCredential(app_cred_id string) (Core, error) {
	var m CoreApplicationCredentialMapping
	err := self.db.Where("application_credential_id = ?", app_cred_id).Find(&m).Error
	if err != nil {
		return empty_core, err
	}

	var c Core
	err = self.db.Where("id = ?", *m.CoreId).First(&c).Error
	if err != nil {
		return empty_core, err
	}

	return c, nil
}

func (self *storageImpl) get_entity(id string) (Entity, error) {
	var e Entity
	err := self.db.Where("id = ?", id).First(&e).Error
	if err != nil {
		return empty_entity, err
	}

	return e, nil
}

func (self *storageImpl) CreateEntity(entity Entity) (Entity, error) {
	now := time.Now()
	entity.HeartbeatAt = &now

	err := self.db.Create(&entity).Error
	if err != nil {
		return empty_entity, err
	}

	entity, err = self.get_entity(*entity.Id)
	if err != nil {
		return empty_entity, err
	}

	self.logger.WithFields(log.Fields{
		"id":      *entity.Id,
		"core_id": *entity.CoreId,
	}).Debugf("create entity")
	return entity, nil
}

func (self *storageImpl) DeleteEntity(id string) error {
	err := self.db.Delete(&Entity{}, "id = ?", id).Error
	if err != nil {
		return err
	}

	self.logger.WithField("id", id).Debugf("delete entity")
	return err
}

func (self *storageImpl) PatchEntity(id string, entity Entity) (Entity, error) {
	var e Entity

	if entity.Name != nil {
		e.Name = entity.Name
	}

	if entity.State != nil {
		e.State = entity.State
	}

	if entity.HeartbeatAt != nil {
		e.HeartbeatAt = entity.HeartbeatAt
	}

	err := self.db.Model(&Entity{}).Where("id = ?", id).Updates(e).Error
	if err != nil {
		return empty_entity, err
	}

	entity, err = self.get_entity(id)
	if err != nil {
		return empty_entity, err
	}

	self.logger.WithField("id", id).Debugf("patch entity")
	return entity, nil
}

func (self *storageImpl) GetEntity(id string) (Entity, error) {
	entity, err := self.get_entity(id)
	if err != nil {
		return empty_entity, err
	}

	self.logger.WithField("id", id).Debugf("get entity")
	return entity, nil
}

func (self *storageImpl) list_entities(entity Entity) ([]Entity, error) {
	var ents_t []Entity
	err := self.db.Select("id").Find(&ents_t, entity).Error
	if err != nil {
		return nil, err
	}

	var entities []Entity
	for _, e := range ents_t {
		ent, err := self.get_entity(*e.Id)
		if err != nil {
			return nil, err
		}

		entities = append(entities, ent)
	}

	return entities, nil
}

func (self *storageImpl) ListEntities(ent Entity) ([]Entity, error) {
	entities, err := self.list_entities(ent)
	if err != nil {
		return nil, err
	}

	self.logger.Debugf("list entities")
	return entities, nil
}

func (self *storageImpl) ListEntitiesForCore(core_id string, ent Entity) ([]Entity, error) {
	ent.CoreId = &core_id
	entities, err := self.list_entities(ent)
	if err != nil {
		return nil, err
	}

	self.logger.Debugf("list entities for core")
	return entities, nil
}

func newStorageImpl(driver, uri string, logger log.FieldLogger) (*storageImpl, error) {
	db, err := gorm.Open(driver, uri)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Core{})
	db.AutoMigrate(&Entity{})
	db.AutoMigrate(&CoreApplicationCredentialMapping{})

	return &storageImpl{
		db:     db,
		logger: logger,
	}, nil
}
