package metathings_tagd_storage

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	id_test0 = "id-test0"
	id_test1 = "id-test1"

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
	opt.Uri = os.Getenv("MTT_MONGO_URI")
	if opt.Uri == "" {
		opt.Uri = "mongodb://127.0.0.1:27107"
	}
	opt.Database = os.Getenv("MTT_MONGO_DATABASE")
	if opt.Database == "" {
		opt.Database = "test"
	}
	opt.Collection = os.Getenv("MTT_MONGO_COLLECTION")
	if opt.Collection == "" {
		opt.Collection = "metathings-testing"
	}

	s.opt = opt
	s.stor = &MongoStorage{opt: s.opt}

	s.Nil(s.stor.connect())
	s.Nil(s.stor.get_collection().Drop(s.stor.context()))
	s.Nil(s.stor.Tag(id_test0, []string{tag_test0, tag_test1}))
	s.Nil(s.stor.Tag(id_test1, []string{tag_test1, tag_test2}))
}

func (s *mongoStorageTestSuite) TestGet() {
	tags, err := s.stor.Get(id_test0)
	s.Nil(err)
	s.ElementsMatch([]string{tag_test0, tag_test1}, tags)

	_, err = s.stor.Get("unknown")
	s.Equal(ErrNotFound, err)
}

func (s *mongoStorageTestSuite) TestQuery() {
	ids, err := s.stor.Query([]string{tag_test0})
	s.Nil(err)
	s.ElementsMatch([]string{id_test0}, ids)

	ids, err = s.stor.Query([]string{tag_test1})
	s.Nil(err)
	s.ElementsMatch([]string{id_test0, id_test1}, ids)

	ids, err = s.stor.Query([]string{tag_test2})
	s.Nil(err)
	s.ElementsMatch([]string{id_test1}, ids)

	ids, err = s.stor.Query([]string{})
	s.Nil(err)
	s.ElementsMatch([]string{}, ids)

	ids, err = s.stor.Query([]string{tag_test0, tag_test1})
	s.Nil(err)
	s.ElementsMatch([]string{id_test0}, ids)

	ids, err = s.stor.Query([]string{tag_test0, tag_test2})
	s.Nil(err)
	s.ElementsMatch([]string{}, ids)
}

func (s *mongoStorageTestSuite) TestRemove() {
	s.Nil(s.stor.Remove(id_test0))

	_, err := s.stor.Get(id_test0)
	s.Equal(ErrNotFound, err)
}

func (s *mongoStorageTestSuite) TestTag() {
	temp := "temp"
	tags := []string{tag_test0}
	s.Nil(s.stor.Tag(temp, []string{tag_test0}))
	ret_tags, err := s.stor.Get(temp)
	s.Nil(err)
	s.ElementsMatch(tags, ret_tags)
}

func (s *mongoStorageTestSuite) TestUntag() {
	s.Nil(s.stor.Untag(id_test0, []string{tag_test0}))
	tags, err := s.stor.Get(id_test0)
	s.Nil(err)
	s.ElementsMatch(tags, []string{tag_test1})
}

func TestMongoStorageTestSuite(t *testing.T) {
	suite.Run(t, new(mongoStorageTestSuite))
}
