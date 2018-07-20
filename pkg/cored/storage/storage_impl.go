package metathings_cored_storage

import (
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

var schemas = `
CREATE TABLE IF NOT EXISTS core (
    id VARCHAR(255),
    name VARCHAR(255),
    project_id VARCHAR(255),
    owner_id VARCHAR(255),
    state VARCHAR(255),
    heartbeat_at DATETIME,

    created_at DATETIME,
    updated_at DATETIME
);

CREATE TABLE IF NOT EXISTS entity (
    id VARCHAR(255),
    core_id VARCHAR(255),
    name VARCHAR(255),
    service_name VARCHAR(255),
    endpoint VARCHAR(255),
    state VARCHAR(255),
    heartbeat_at DATETIME,

    created_at DATETIME,
    updated_at DATETIME
);

CREATE TABLE IF NOT EXISTS core_app_cred_relationship (
    core_id VARCHAR(255),
    app_cred_id VARCHAR(255)
);
`

type storageImpl struct {
	logger log.FieldLogger
	db     *sqlx.DB
}

func (s *storageImpl) CreateCore(core Core) (Core, error) {
	c := Core{}

	now := time.Now()
	core.CreatedAt = now
	core.UpdatedAt = now
	core.HeartbeatAt = &now
	_, err := s.db.NamedExec(`
INSERT INTO core (id, name, project_id, owner_id, state, heartbeat_at, created_at, updated_at)
VALUES (:id, :name, :project_id, :owner_id, :state, :heartbeat_at, :created_at, :updated_at)`, &core)
	if err != nil {
		s.logger.WithError(err).Errorf("failed to create core")
		return c, err
	}

	s.db.Get(&c, "SELECT * FROM core WHERE id=$1", *core.Id)

	s.logger.WithField("core_id", *core.Id).Debugf("create core")
	return c, nil
}

func (s *storageImpl) DeleteCore(core_id string) error {
	tx, err := s.db.Begin()
	if err != nil {
		s.logger.WithError(err).Errorf("failed to begin tx")
		return err
	}
	tx.Exec("DELETE FROM entity WHERE core_id=$1", core_id)
	tx.Exec("DELETE FROM core WHERE id=$1", core_id)
	err = tx.Commit()
	if err != nil {
		s.logger.
			WithError(err).
			WithField("core_id", core_id).
			Errorf("failed to delete core")
		return err
	}

	s.logger.WithField("core_id", core_id).Debugf("delete core")
	return nil
}

func (s *storageImpl) PatchCore(core_id string, core Core) (Core, error) {
	values := []string{}
	arguments := []interface{}{}
	i := 1
	c := Core{}

	if core.Name != nil {
		values = append(values, fmt.Sprintf("name=$%v", i))
		arguments = append(arguments, *core.Name)
		i++
	}

	if core.State != nil {
		values = append(values, fmt.Sprintf("state=$%v", i))
		arguments = append(arguments, *core.State)
		i++
	}

	if core.HeartbeatAt != nil {
		values = append(values, fmt.Sprintf("heartbeat_at=$%v", i))
		arguments = append(arguments, *core.HeartbeatAt)
		i++
	}

	if len(values) > 0 {
		values = append(values, fmt.Sprintf("updated_at=$%v", i))
		arguments = append(arguments, time.Now())
		i++

		val := strings.Join(values, ", ")
		arguments = append(arguments, core_id)

		sql_str := "UPDATE core SET " + val + fmt.Sprintf(" WHERE id=$%v", i)
		s.logger.WithFields(log.Fields{
			"sql":  sql_str,
			"args": arguments,
		}).Debugf("execute sql")
		_, err := s.db.Exec(sql_str, arguments...)
		if err != nil {
			s.logger.WithError(err).
				WithField("core_id", core_id).
				Errorf("failed to patch core")
			return c, err
		}
		s.db.Get(&c, "SELECT * FROM core WHERE id=$1", core_id)
		s.logger.WithField("core_id", core_id).Debugf("update core")
		return c, nil
	}
	s.logger.WithField("core_id", core_id).Debugf("nothing changed when update core")
	return c, ErrNothingChanged
}

func (s *storageImpl) GetCore(core_id string) (Core, error) {
	c := Core{}
	err := s.db.Get(&c, "SELECT * FROM core WHERE id=$1", core_id)
	if err != nil {
		s.logger.WithError(err).
			WithField("core_id", core_id).
			Errorf("failed to get core")
		return c, err
	}

	s.logger.WithField("core_id", core_id).Debugf("get core")
	return c, nil
}

func (s *storageImpl) ListCores(core Core) ([]Core, error) {
	cs, err := s.list_cores(core)
	if err != nil {
		s.logger.WithError(err).
			Errorf("failed to list cores")
		return nil, err
	}
	s.logger.Debugf("list cores")
	return cs, nil
}

func (s *storageImpl) ListCoresForUser(owner_id string, core Core) ([]Core, error) {
	core.OwnerId = &owner_id
	cs, err := s.list_cores(core)
	if err != nil {
		s.logger.WithField("owner_id", owner_id).WithError(err).Errorf("failed to list cores for user")
		return nil, err
	}
	s.logger.WithField("owner_id", owner_id).Debugf("list cores for user")
	return cs, nil
}

func (s *storageImpl) list_cores(core Core) ([]Core, error) {
	var err error
	values := []string{}
	arguments := []interface{}{}
	i := 0
	cores := []Core{}

	if core.Name != nil {
		values = append(values, fmt.Sprintf("name=$%v", i))
		arguments = append(arguments, *core.Name)
		i++
	}

	if core.ProjectId != nil {
		values = append(values, fmt.Sprintf("project_id=$%v", i))
		arguments = append(arguments, *core.ProjectId)
		i++
	}

	if core.OwnerId != nil {
		values = append(values, fmt.Sprintf("owner_id=$%v", i))
		arguments = append(arguments, *core.OwnerId)
		i++
	}

	if core.State != nil {
		values = append(values, fmt.Sprintf("state=$%v", i))
		arguments = append(arguments, *core.State)
		i++
	}

	if len(values) == 0 {
		err = s.db.Select(&cores, "SELECT * FROM core")
	} else {
		val := strings.Join(values, " and ")
		sql_str := fmt.Sprintf("SELECT * FROM core WHERE %v", val)
		s.logger.WithFields(log.Fields{
			"sql":  sql_str,
			"args": arguments,
		}).Debugf("execute sql")
		err = s.db.Select(&cores, sql_str, arguments...)
	}
	if err != nil {
		return nil, err
	}

	return cores, nil
}

func (s *storageImpl) AssignCoreToApplicationCredential(core_id string, app_cred_id string) error {
	_, err := s.db.Exec("INSERT INTO core_app_cred_relationship (core_id, app_cred_id) VALUES ($1, $2)", core_id, app_cred_id)
	if err != nil {
		s.logger.WithError(err).Errorf("failed to assign core to application credential")
		return err
	}
	s.logger.WithFields(log.Fields{
		"core_id":                   core_id,
		"application_credential_id": app_cred_id,
	}).Debugf("assign core to application credential")
	return nil
}

func (s *storageImpl) GetAssignedCoreFromApplicationCredential(app_cred_id string) (Core, error) {
	c := Core{}
	err := s.db.Get(&c, `
SELECT c.*
FROM core AS c
JOIN core_app_cred_relationship AS r
ON c.id = r.core_id
WHERE r.app_cred_id = $1`, app_cred_id)
	if err != nil {
		s.logger.WithError(err).
			WithField("application_credential_id", app_cred_id).
			Errorf("failed to get assigned core from application credential")
		return c, err
	}
	return c, nil
}

func (s *storageImpl) CreateEntity(entity Entity) (Entity, error) {
	e := Entity{}

	now := time.Now()
	entity.CreatedAt = now
	entity.UpdatedAt = now
	entity.HeartbeatAt = &now
	_, err := s.db.NamedExec(`
INSERT INTO entity (id, core_id, name, service_name, endpoint, state, created_at, updated_at, heartbeat_at)
VALUES (:id, :core_id, :name, :service_name, :endpoint, :state, :created_at, :updated_at, :heartbeat_at)`, &entity)
	if err != nil {
		s.logger.WithError(err).Errorf("failed to create entity")
	}

	s.db.Get(&e, "SELECT * FROM entity WHERE id=$1", *entity.Id)
	s.logger.WithField("entity_id", *e.Id).Debugf("create entity")
	return e, nil
}

func (s *storageImpl) DeleteEntity(entity_id string) error {
	_, err := s.db.Exec("DELETE FROM entity WHERE id=$1", entity_id)
	if err != nil {
		s.logger.WithError(err).
			WithField("entity_id", entity_id).
			Errorf("failed to delete entity")
		return err
	}
	s.logger.WithField("entity_id", entity_id).Debugf("delete entity")
	return nil
}

func (s *storageImpl) PatchEntity(entity_id string, entity Entity) (Entity, error) {
	values := []string{}
	arguments := []interface{}{}
	i := 1
	e := Entity{}

	if entity.State != nil {
		values = append(values, fmt.Sprintf("state=$%v", i))
		arguments = append(arguments, *entity.State)
		i++
	}

	if entity.HeartbeatAt != nil {
		values = append(values, fmt.Sprintf("heartbeat_at=$%v", i))
		arguments = append(arguments, *entity.HeartbeatAt)
		i++
	}

	if len(values) > 0 {
		values = append(values, fmt.Sprintf("updated_at=$%v", i))
		arguments = append(arguments, time.Now())
		i++

		val := strings.Join(values, ", ")
		arguments = append(arguments, entity_id)

		sql_str := "UPDATE entity SET " + val + fmt.Sprintf(" WHERE id=$%v", i)
		s.logger.WithFields(log.Fields{
			"sql":  sql_str,
			"args": arguments,
		}).Debugf("execute sql")
		_, err := s.db.Exec(sql_str, arguments...)
		if err != nil {
			s.logger.WithError(err).
				WithField("entity_id", entity_id).
				Errorf("failed to patch entity")
			return e, err
		}
		s.db.Get(&e, "SELECT * FROM entity WHERE id=$1", entity_id)
		s.logger.WithField("entity_id", entity_id).Debugf("update entity")
		return e, nil

	}

	s.logger.WithField("entity_id", entity_id).Debugf("nothing changed when update entity")
	return Entity{}, ErrNothingChanged
}

func (s *storageImpl) GetEntity(entity_id string) (Entity, error) {
	e := Entity{}
	err := s.db.Get(&e, "SELECT * FROM entity WHERE id=$1", entity_id)
	if err != nil {
		s.logger.WithError(err).
			WithField("entity_id", entity_id).
			Errorf("failed to get entity")
		return e, err
	}

	s.logger.WithField("entity_id", entity_id).Debugf("get entity")
	return e, nil
}

func (s *storageImpl) ListEntities(entity Entity) ([]Entity, error) {
	es, err := s.list_entities(entity)
	if err != nil {
		s.logger.WithError(err).
			Errorf("failed to list entities")
		return nil, err
	}

	s.logger.Debugf("list entities")
	return es, nil
}

func (s *storageImpl) list_entities(entity Entity) ([]Entity, error) {
	var err error
	values := []string{}
	arguments := []interface{}{}
	i := 0
	entities := []Entity{}

	if entity.CoreId != nil {
		values = append(values, fmt.Sprintf("core_id=$%v", i))
		arguments = append(arguments, *entity.CoreId)
		i++
	}

	if entity.Name != nil {
		values = append(values, fmt.Sprintf("name=$%v", i))
		arguments = append(arguments, *entity.Name)
	}

	if entity.ServiceName != nil {
		values = append(values, fmt.Sprintf("service_name=$%v", i))
		arguments = append(arguments, *entity.ServiceName)
	}

	if entity.State != nil {
		values = append(values, fmt.Sprintf("state=$%v", i))
		arguments = append(arguments, *entity.State)
	}

	if len(values) == 0 {
		err = s.db.Select(&entities, "SELECT * FROM entity")
	} else {
		val := strings.Join(values, " and ")
		sql_str := fmt.Sprintf("SELECT * FROM entity WHERE %v", val)
		s.logger.WithFields(log.Fields{
			"sql":  sql_str,
			"args": arguments,
		}).Debugf("execute sql")
		err = s.db.Select(&entities, sql_str, arguments...)
	}
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (s *storageImpl) ListEntitiesForCore(core_id string, entity Entity) ([]Entity, error) {
	entity.CoreId = &core_id
	es, err := s.list_entities(entity)
	if err != nil {
		s.logger.WithError(err).Errorf("failed to list entities for core")
		return nil, err
	}

	s.logger.WithField("core_id", core_id).Debugf("list entities for core")
	return es, nil
}

func newStorageImpl(driver, uri string, logger log.FieldLogger) (*storageImpl, error) {
	if driver != "sqlite3" {
		logger.WithField("driver", driver).Errorf("not supprted driver")
		return nil, ErrUnknownStorageDriver
	}
	db, err := sqlx.Connect(driver, uri)
	if err != nil {
		logger.WithFields(log.Fields{
			"driver": driver,
			"uri":    uri,
		}).WithError(err).Errorf("failed to connect database")
		return nil, err
	}
	db.MustExec(schemas)
	return &storageImpl{
		logger: logger.WithField("#module", "storage"),
		db:     db,
	}, nil
}
