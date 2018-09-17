package metathings_streamd_storage

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type storageImpl struct {
	db     *gorm.DB
	logger log.FieldLogger
}

func (self *storageImpl) CreateStream(stm Stream) (Stream, error) {
	tx := self.db.Begin()
	self.tx_create_stream(tx, stm)
	err := tx.Commit().Error
	if err != nil {
		return Stream{}, err
	}

	stm, err = self.get_stream(*stm.Id)
	if err != nil {
		return Stream{}, err
	}

	return stm, nil
}

func (self *storageImpl) tx_create_stream(tx *gorm.DB, stm Stream) {
	tx.Create(&stm)
	self.tx_create_sources(tx, stm.Sources)
	self.tx_create_groups(tx, stm.Groups)
}

func (self *storageImpl) tx_create_sources(tx *gorm.DB, sources []Source) {
	for _, src := range sources {
		self.tx_create_source(tx, src)
	}
}

func (self *storageImpl) tx_create_groups(tx *gorm.DB, groups []Group) {
	for _, grp := range groups {
		self.tx_create_group(tx, grp)
	}
}

func (self *storageImpl) tx_create_source(tx *gorm.DB, source Source) {
	tx.Create(&source)
	self.tx_create_upstream(tx, source.Upstream)
}

func (self *storageImpl) tx_create_upstream(tx *gorm.DB, upstream Upstream) {
	tx.Create(&upstream)
}

func (self *storageImpl) tx_create_group(tx *gorm.DB, group Group) {
	tx.Create(&group)

	self.tx_create_inputs(tx, group.Inputs)
	self.tx_create_outputs(tx, group.Outputs)
}

func (self *storageImpl) tx_create_inputs(tx *gorm.DB, inputs []Input) {
	for _, in := range inputs {
		self.tx_create_input(tx, in)
	}
}

func (self *storageImpl) tx_create_input(tx *gorm.DB, input Input) {
	tx.Create(&input)
}

func (self *storageImpl) tx_create_outputs(tx *gorm.DB, outputs []Output) {
	for _, out := range outputs {
		self.tx_create_output(tx, out)
	}
}

func (self *storageImpl) tx_create_output(tx *gorm.DB, output Output) {
	tx.Create(&output)
}

func (self *storageImpl) get_stream(stm_id string) (Stream, error) {
	var stm Stream
	err := self.db.First(&stm, "id = ?", stm_id).Error
	if err != nil {
		return stm, err
	}

	stm.Sources, err = self.get_sources_by_stream_id(stm_id)
	if err != nil {
		return stm, err
	}

	stm.Groups, err = self.get_groups_by_stream_id(stm_id)
	if err != nil {
		return stm, err
	}

	return stm, nil
}

func (self *storageImpl) get_sources_by_stream_id(stm_id string) ([]Source, error) {
	var srcs_t []struct{ Id *string }
	err := self.db.Model(Source{}).Select("id").Where("stream_id = ?", stm_id).Find(&srcs_t).Error
	if err != nil {
		return nil, err
	}

	var sources []Source
	for _, src_t := range srcs_t {
		source, err := self.get_source(*src_t.Id)
		if err != nil {
			return nil, err
		}
		sources = append(sources, source)
	}

	return sources, nil
}

func (self *storageImpl) get_source(src_id string) (Source, error) {
	panic("unimplemented")
}

func (self *storageImpl) get_groups_by_stream_id(stm_id string) ([]Group, error) {
	var grps_t []struct{ Id *string }
	err := self.db.Model(Group{}).Select("id").Where("stream_id = ?", stm_id).Find(&grps_t).Error
	if err != nil {
		return nil, err
	}

	var groups []Group
	for _, grp_t := range grps_t {
		group, err := self.get_group(*grp_t.Id)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}

	return groups, nil
}

func (self *storageImpl) get_group(grp_id string) (Group, error) {
	panic("unimplemented")
}

func (self *storageImpl) DeleteStream(stm_id string) error {
	stm, err := self.get_stream(stm_id)
	if err != nil {
		return err
	}

	tx := self.db.Begin()
	self.tx_delete_stream(tx, stm)
	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

func (self *storageImpl) tx_delete_stream(tx *gorm.DB, stm Stream) {
	self.tx_delete_sources(tx, stm.Sources)
	self.tx_delete_groups(tx, stm.Groups)
}

func (self *storageImpl) tx_delete_sources(tx *gorm.DB, sources []Source) {
	for _, src := range sources {
		self.tx_delete_source(tx, src)
	}
}

func (self *storageImpl) tx_delete_groups(tx *gorm.DB, groups []Group) {
	for _, grp := range groups {
		self.tx_delete_group(tx, grp)
	}
}

func (self *storageImpl) tx_delete_source(tx *gorm.DB, source Source) {
	self.tx_delete_upstream_by_source_id(tx, *source.Id)
	tx.Delete(Source{}, "id = ?", *source.Id)
}

func (self *storageImpl) tx_delete_upstream_by_source_id(tx *gorm.DB, upstm_id string) {
	tx.Delete(Upstream{}, "upstream_id = ?", upstm_id)
}

func (self *storageImpl) tx_delete_group(tx *gorm.DB, group Group) {
	self.tx_delete_inputs_by_group_id(tx, *group.Id)
	self.tx_delete_outputs_by_group_id(tx, *group.Id)
	tx.Delete(Group{}, "id = ?", *group.Id)
}

func (self *storageImpl) tx_delete_inputs_by_group_id(tx *gorm.DB, grp_id string) {
	tx.Delete(Input{}, "group_id = ?", grp_id)
}

func (self *storageImpl) tx_delete_outputs_by_group_id(tx *gorm.DB, grp_id string) {
	tx.Delete(Output{}, "group_id = ?", grp_id)
}

func (self *storageImpl) PatchStream(stm_id string, stm Stream) (Stream, error) {
	var s Stream

	if stm.Name != nil {
		s.Name = stm.Name
	}

	if stm.State != nil {
		s.State = stm.State
	}

	err := self.db.Model(Stream{}).Where("id = ?", stm_id).Updates(s).Error
	if err != nil {
		return Stream{}, err
	}

	stm, err = self.get_stream(stm_id)
	if err != nil {
		return Stream{}, err
	}

	return stm, nil
}

func (self *storageImpl) GetStream(stm_id string) (Stream, error) {
	stm, err := self.get_stream(stm_id)
	if err != nil {
		return Stream{}, err
	}

	return stm, nil
}

func (self *storageImpl) list_streams(stm Stream) ([]Stream, error) {
	var stms_t []struct{ Id *string }
	err := self.db.Select("id").Find(&stms_t, stm).Error
	if err != nil {
		return nil, err
	}

	var streams []Stream
	for _, s := range stms_t {
		stm, err := self.get_stream(*s.Id)
		if err != nil {
			return nil, err
		}
		streams = append(streams, stm)
	}

	return streams, nil
}

func (self *storageImpl) ListStreams(stm Stream) ([]Stream, error) {
	streams, err := self.list_streams(stm)
	if err != nil {
		return nil, err
	}
	return streams, nil
}

func (self *storageImpl) ListStreamsForUser(owner_id string, stm Stream) ([]Stream, error) {
	stm.OwnerId = &owner_id
	streams, err := self.list_streams(stm)
	if err != nil {
		return nil, err
	}
	return streams, nil
}

func newStorageImpl(driver, uri string, logger log.FieldLogger) (Storage, error) {
	db, err := gorm.Open(driver, uri)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Stream{})
	db.AutoMigrate(&Source{})
	db.AutoMigrate(&Group{})
	db.AutoMigrate(&Upstream{})
	db.AutoMigrate(&Input{})
	db.AutoMigrate(&Output{})

	return &storageImpl{
		logger: logger.WithField("#module", "storage"),
		db:     db,
	}, nil
}
