package metathings_toolkit_tag

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

type mongoTagToolkitTestSuite struct {
	suite.Suite
	opt   *MongoTagToolkitOption
	tagtk *MongoTagToolkit
}

func (s *mongoTagToolkitTestSuite) SetupTest() {
	opt := NewMongoTagToolkitOption()
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
	s.tagtk = &MongoTagToolkit{opt: s.opt}

	s.Nil(s.tagtk.connect())
	s.Nil(s.tagtk.get_collection().Drop(s.tagtk.context()))
	s.Nil(s.tagtk.Tag(id_test0, []string{tag_test0, tag_test1}))
	s.Nil(s.tagtk.Tag(id_test1, []string{tag_test1, tag_test2}))
}

func (s *mongoTagToolkitTestSuite) TestGet() {
	tags, err := s.tagtk.Get(id_test0)
	s.Nil(err)
	s.ElementsMatch([]string{tag_test0, tag_test1}, tags)

	_, err = s.tagtk.Get("unknown")
	s.Equal(ErrNotFound, err)
}

func (s *mongoTagToolkitTestSuite) TestQuery() {
	ids, err := s.tagtk.Query([]string{tag_test0})
	s.Nil(err)
	s.ElementsMatch([]string{id_test0}, ids)

	ids, err = s.tagtk.Query([]string{tag_test1})
	s.Nil(err)
	s.ElementsMatch([]string{id_test0, id_test1}, ids)

	ids, err = s.tagtk.Query([]string{tag_test2})
	s.Nil(err)
	s.ElementsMatch([]string{id_test1}, ids)

	ids, err = s.tagtk.Query([]string{})
	s.Nil(err)
	s.ElementsMatch([]string{}, ids)

	ids, err = s.tagtk.Query([]string{tag_test0, tag_test1})
	s.Nil(err)
	s.ElementsMatch([]string{id_test0}, ids)

	ids, err = s.tagtk.Query([]string{tag_test0, tag_test2})
	s.Nil(err)
	s.ElementsMatch([]string{}, ids)
}

func (s *mongoTagToolkitTestSuite) TestRemove() {
	s.Nil(s.tagtk.Remove(id_test0))

	_, err := s.tagtk.Get(id_test0)
	s.Equal(ErrNotFound, err)
}

func (s *mongoTagToolkitTestSuite) TestTag() {
	temp := "temp"
	tags := []string{tag_test0}
	s.Nil(s.tagtk.Tag(temp, []string{tag_test0}))
	ret_tags, err := s.tagtk.Get(temp)
	s.Nil(err)
	s.ElementsMatch(tags, ret_tags)
}

func (s *mongoTagToolkitTestSuite) TestUntag() {
	s.Nil(s.tagtk.Untag(id_test0, []string{tag_test0}))
	tags, err := s.tagtk.Get(id_test0)
	s.Nil(err)
	s.ElementsMatch(tags, []string{tag_test1})
}

func TestMongoTagToolkitTestSuite(t *testing.T) {
	suite.Run(t, new(mongoTagToolkitTestSuite))
}
