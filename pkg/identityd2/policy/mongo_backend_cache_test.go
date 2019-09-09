package metathings_identityd2_policy

import (
	"testing"

	"github.com/stretchr/testify/suite"

	log_helper "github.com/nayotta/metathings/pkg/common/log"
	test_helper "github.com/nayotta/metathings/pkg/common/test"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
)

type mongoBackendCacheTestSuite struct {
	suite.Suite
	cb *MongoBackendCache
}

func (s *mongoBackendCacheTestSuite) SetupTest() {
	mgo_uri := test_helper.GetTestMongoUri()
	mgo_db := test_helper.GetTestMongoDatabase()
	mgo_coll := test_helper.GetTestMongoCollection()
	logger, _ := log_helper.NewLogger("test", "debug")

	cb, err := NewBackendCache(
		"mongo",
		"mongo_uri", mgo_uri,
		"mongo_database", mgo_db,
		"mongo_collection", mgo_coll,
		"logger", logger,
	)
	s.Nil(err)

	s.cb = cb.(*MongoBackendCache)
	s.cb.mgo_coll.Drop(s.cb.context())
	err = s.cb.Set(test_subject, test_object, test_action, true)
	s.Nil(err)

	err = s.cb.Set(test_subject, test_object, test_action2, false)
	s.Nil(err)
}

func (s *mongoBackendCacheTestSuite) TestGet() {
	ret, err := s.cb.Get(test_subject, test_object, test_action)
	s.Nil(err)
	s.True(ret)

	_, err = s.cb.Get(test_subject2, test_object2, test_action2)
	s.Equal(ErrNoCached, err)

	_, err = s.cb.Get(test_subject2, test_object2, test_action)
	s.Equal(ErrNoCached, err)
}

func (s *mongoBackendCacheTestSuite) TestSet() {
	err := s.cb.Set(test_subject2, test_object2, test_action2, false)
	s.Nil(err)

	ret, err := s.cb.Get(test_subject2, test_object2, test_action2)
	s.Nil(err)
	s.False(ret)
}

func (s *mongoBackendCacheTestSuite) TestRemoveBySubject() {
	err := s.cb.Remove("subject", test_subject)
	s.Nil(err)

	_, err = s.cb.Get(test_subject, test_object, test_action)
	s.Equal(ErrNoCached, err)

	_, err = s.cb.Get(test_subject, test_object, test_action2)
	s.Equal(ErrNoCached, err)
}

func (s *mongoBackendCacheTestSuite) TestRemoveByObject() {
	err := s.cb.Remove("object", test_object)
	s.Nil(err)

	_, err = s.cb.Get(test_subject, test_object, test_action)
	s.Equal(ErrNoCached, err)

	_, err = s.cb.Get(test_subject, test_object, test_action2)
	s.Equal(ErrNoCached, err)
}

func (s *mongoBackendCacheTestSuite) TestRemoveByAction() {
	err := s.cb.Remove("action", test_action)
	s.Nil(err)

	_, err = s.cb.Get(test_subject, test_object, test_action)
	s.Equal(ErrNoCached, err)

	ret, err := s.cb.Get(test_subject, test_object, test_action2)
	s.Nil(err)
	s.False(ret)
}

func TestMongoBackendCacheTestSuite(t *testing.T) {
	suite.Run(t, new(mongoBackendCacheTestSuite))
}

func init() {
	test_action = &storage.Action{
		Id:   &test_action_id,
		Name: &test_action_name,
	}
	test_action2 = &storage.Action{
		Id:   &test_action2_id,
		Name: &test_action2_name,
	}
	test_subject = &storage.Entity{
		Id: &test_subject_id,
		Groups: []*storage.Group{
			test_group,
		},
	}
	test_subject2 = &storage.Entity{
		Id: &test_subject2_id,
		Groups: []*storage.Group{
			test_group,
		},
	}
	test_object = &storage.Entity{Id: &test_object_id}
	test_object2 = &storage.Entity{Id: &test_object2_id}
}
