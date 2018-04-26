package metathings_core_storage

import (
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

var schemas = `
CREATE TABLE core (
    created_at datetime,
    updated_at datetime,

    id varchar(255),
    name varchar(255),
    project_id varchar(255),
    owner_id varchar(255),
    state varchar(255)
);

CREATE TABLE entity (
    created_at datetime,
    updated_at datetime,

    id varchar(255),
    core_id varchar(255),
    name varchar(255),
    service_name varchar(255),
    endpoint varchar(255),
    state varchar(255)
);

CREATE TABLE core_app_cred_relationship (
    core_id varchar(255),
    app_cred_id varchar(255)
);
`

type storageImpl struct {
	logger log.FieldLogger
	db     *sqlx.DB
}

func (s *storageImpl) CreateCore(core Core) (Core, error) {
	c := Core{}

	core.InitializedAtNow()
	_, err := s.db.NamedExec("INSERT INTO core (id, name, project_id, owner_id, state) VALUES (:id, :name, :project_id, :owner_id, :state)", &core)
	if err != nil {
		s.logger.WithError(err).Errorf("failed to create core")
		return c, err
	}

	s.db.Get(&c, "SELECT * FROM core WHERE id=$1", core.Id)

	s.logger.WithField("core_id", *core.Id).Infof("create core")
	return c, nil
}

func (s *storageImpl) DeleteCore(core_id string) error {
	_, err := s.db.Exec("DELETE FROM core WHERE id=$1", core_id)
	if err != nil {
		s.logger.WithError(err).
			WithField("core_id", core_id).
			Errorf("failed to delete core")
		return err
	}

	s.logger.WithField("core_id", core_id).Infof("delete core")
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
		i += 1
	}

	if core.State != nil {
		values = append(values, fmt.Sprintf("state=$%v", i))
		arguments = append(arguments, *core.State)
		i += 1
	}

	if len(values) > 0 {
		values = append(values, fmt.Sprintf("updated_at=%v", i))
		arguments = append(arguments, time.Now())
		i += 1

		val := strings.Join(values, ", ")
		arguments = append(arguments, core_id)

		_, err := s.db.Exec("UPDATE core SET "+val+fmt.Sprintf(" WHERE id=$%v", i),
			arguments...)
		if err != nil {
			s.logger.WithError(err).
				WithField("core_id", core_id).
				Errorf("failed to patch core")
			return c, err
		}
		s.db.Get(&c, "SELECT * FROM core WHERE id=$1", core_id)
		s.logger.WithField("core_id", core_id).Infof("update core")
		return c, nil
	}
	s.logger.WithField("core_id", core_id).Debugf("nothing changed when update core")
	return c, NothingChanged
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

func (s *storageImpl) ListCores(_ Core) ([]Core, error) {
	cores := []Core{}
	err := s.db.Select(&cores, "SELECT * FROM core")
	if err != nil {
		s.logger.WithError(err).
			Errorf("failed to list cores")
		return cores, err
	}

	s.logger.Debugf("list cores")
	return cores, nil
}

func (s *storageImpl) ListCoresForUser(owner_id string, _ Core) ([]Core, error) {
	cores := []Core{}
	err := s.db.Select(&cores, "SELECT * FROM core WHERE owner_id=$1", owner_id)
	if err != nil {
		s.logger.WithError(err).
			WithField("owner_id", owner_id).
			Errorf("failed to list cores for user")
		return cores, err
	}

	s.logger.WithField("owner_id", owner_id).Debugf("list cores for user")
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
	}).Infof("assign core to application credential")
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

	e.InitializedAtNow()
	_, err := s.db.NamedExec("INSERT INTO entity (id, core_id, name, service_name, endpoint, state) VALUES (:id, :core_id, :name, :service_name, :endpoint, :state)", &entity)
	if err != nil {
		s.logger.WithError(err).Errorf("failed to create entity")
	}

	s.db.Get(&e, "SELECT * FROM entity WHERE id=$1", entity.Id)
	s.logger.WithField("entity_id", *e.Id).Infof("create entity")
	return e, nil
}

func (s *storageImpl) DeleteEntity(entity_id string) error {
	_, err := s.db.Exec("DELETE FROM core WHERE id=$1", entity_id)
	if err != nil {
		s.logger.WithError(err).
			WithField("entity_id", entity_id).
			Errorf("failed to delete entity")
		return err
	}
	s.logger.WithField("entity_id", entity_id).Infof("delete entity")
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
		i += 1
	}

	if len(values) > 0 {
		values = append(values, fmt.Sprintf("updated_at=%v", i))
		arguments = append(arguments, time.Now())
		i += 1

		val := strings.Join(values, ", ")
		arguments = append(arguments, entity_id)

		_, err := s.db.Exec("UPDATE entity SET "+val+fmt.Sprintf(" WHERE id=$%v", i),
			arguments...)
		if err != nil {
			s.logger.WithError(err).
				WithField("entity_id", entity_id).
				Errorf("failed to patch entity")
			return e, err
		}
		s.db.Get(&e, "SELECT * FROM entity WHERE id=$1", entity_id)
		s.logger.WithField("entity_id", entity_id).Infof("update entity")
		return e, nil

	}

	s.logger.WithField("entity_id", entity_id).Debugf("nothing changed when update entity")
	return Entity{}, NothingChanged
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

func (s *storageImpl) ListEntities(_ Entity) ([]Entity, error) {
	entities := []Entity{}
	err := s.db.Select(&entities, "SELECT * FROM entity")
	if err != nil {
		s.logger.WithError(err).
			Errorf("failed to list entities")
		return entities, err
	}
	s.logger.Debugf("list entities")
	return entities, nil
}

func (s *storageImpl) ListEntitiesForCore(core_id string, _ Entity) ([]Entity, error) {
	entities := []Entity{}
	err := s.db.Select(&entities, "SELECT * FROM entity WHERE core_id=$1", core_id)
	if err != nil {
		s.logger.WithError(err).
			Errorf("failed to list entities for core")
		return entities, err
	}
	s.logger.Debugf("list entities for core")
	return entities, nil
}

func newStorageImpl(dbpath string, logger log.FieldLogger) (*storageImpl, error) {
	db, err := sqlx.Connect("sqlite3", dbpath)
	if err != nil {
		logger.WithError(err).Errorf("failed to connect database")
		return nil, err
	}
	db.MustExec(schemas)
	return &storageImpl{
		logger: logger.WithField("#module", "storage"),
		db:     db,
	}, nil
}
