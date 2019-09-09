package metathings_tagd_storage

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"

	test_helper "github.com/nayotta/metathings/pkg/common/test"
)

const (
	ns_test0 = "ns-test0"
	ns_test1 = "ns-test1"

	id_unknown = "id-unknown"
	id_test0   = "id-test0"
	id_test1   = "id-test1"

	tag_test0 = "tag-test0"
	tag_test1 = "tag-test1"
	tag_test2 = "tag-test2"
)

type mongoStorageTestSuite struct {
	suite.Suite
	opt  *MongoStorageOption
	stor *MongoStorage
}

func (s *mongoStorageTestSuite) SetupTest() {
	opt := NewMongoStorageOption()
	opt.Uri = test_helper.GetTestMongoUri()
	opt.Database = test_helper.GetTestMongoDatabase()
	opt.Collection = test_helper.GetTestMongoCollection()

	s.opt = opt
	s.stor = &MongoStorage{opt: s.opt, logger: log.New()}

	s.Nil(s.stor.connect())
	s.Nil(s.stor.get_collection().Drop(s.stor.context()))
	s.Nil(s.stor.Tag(ns_test0, id_test0, []string{tag_test0, tag_test1}))
	s.Nil(s.stor.Tag(ns_test0, id_test1, []string{tag_test1, tag_test2}))
}

func (s *mongoStorageTestSuite) TestGet() {
	tags, err := s.stor.Get(ns_test0, id_test0)
	s.Nil(err)
	s.ElementsMatch([]string{tag_test0, tag_test1}, tags)

	_, err = s.stor.Get(ns_test0, id_unknown)
	s.Equal(ErrNotFound, err)

	_, err = s.stor.Get(ns_test1, id_unknown)
	s.Equal(ErrNotFound, err)

	_, err = s.stor.Get(ns_test1, id_test0)
	s.Equal(ErrNotFound, err)
}

func (s *mongoStorageTestSuite) TestQuery() {
	ids, err := s.stor.Query(ns_test0, []string{tag_test0})
	s.Nil(err)
	s.ElementsMatch([]string{id_test0}, ids)

	ids, err = s.stor.Query(ns_test0, []string{tag_test1})
	s.Nil(err)
	s.ElementsMatch([]string{id_test0, id_test1}, ids)

	ids, err = s.stor.Query(ns_test0, []string{tag_test2})
	s.Nil(err)
	s.ElementsMatch([]string{id_test1}, ids)

	ids, err = s.stor.Query(ns_test0, []string{})
	s.Nil(err)
	s.ElementsMatch([]string{}, ids)

	ids, err = s.stor.Query(ns_test0, []string{tag_test0, tag_test1})
	s.Nil(err)
	s.ElementsMatch([]string{id_test0}, ids)

	ids, err = s.stor.Query(ns_test0, []string{tag_test0, tag_test2})
	s.Nil(err)
	s.ElementsMatch([]string{}, ids)
}

func (s *mongoStorageTestSuite) TestRemove() {
	s.Nil(s.stor.Remove(ns_test0, id_test0))

	_, err := s.stor.Get(ns_test0, id_test0)
	s.Equal(ErrNotFound, err)
}

func (s *mongoStorageTestSuite) TestTag() {
	temp := "temp"
	tags := []string{tag_test0}
	s.Nil(s.stor.Tag(ns_test0, temp, []string{tag_test0}))
	ret_tags, err := s.stor.Get(ns_test0, temp)
	s.Nil(err)
	s.ElementsMatch(tags, ret_tags)
}

func (s *mongoStorageTestSuite) TestUntag() {
	s.Nil(s.stor.Untag(ns_test0, id_test0, []string{tag_test0}))
	tags, err := s.stor.Get(ns_test0, id_test0)
	s.Nil(err)
	s.ElementsMatch(tags, []string{tag_test1})
}

func TestMongoStorageTestSuite(t *testing.T) {
	suite.Run(t, new(mongoStorageTestSuite))
}
