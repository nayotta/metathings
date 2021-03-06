package metathings_deviced_simple_storage

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"

	log_helper "github.com/nayotta/metathings/pkg/common/log"
)

const (
	_TEST_LARGE_OBJECT_CONTENT_SIZE = 513
	_TEST_LARGE_OBJECT_CONTENT_MASK = 42
)

var (
	test_device_id      = "test_device_id"
	test_object_prefix  = "test_object_prefix"
	test_object_name    = "test_object_name.txt"
	test_object_content = "test object content"

	test_large_object_prefix  = "test_large_object_prefix"
	test_large_object_name    = "test_large_object_name"
	test_large_object_content = make([]byte, _TEST_LARGE_OBJECT_CONTENT_SIZE)

	test_object = &Object{
		Device: test_device_id,
		Prefix: test_object_prefix,
		Name:   test_object_name,
	}

	test_large_object = &Object{
		Device: test_device_id,
		Prefix: test_large_object_prefix,
		Name:   test_large_object_name,
	}
)

type fileSimpleStorageTestSuite struct {
	suite.Suite
	fss *FileSimpleStorage
}

func (s *fileSimpleStorageTestSuite) SetupTest() {
	logger, err := log_helper.NewLogger("test", "debug")
	s.Nil(err)

	home, err := ioutil.TempDir("", "")
	s.Nil(err)

	fss, err := new_file_simple_storage("home", home, "logger", logger)
	s.Nil(err)

	s.fss = fss.(*FileSimpleStorage)

	err = s.fss.PutObject(test_object, strings.NewReader(test_object_content))
	s.Nil(err)
}

func (s *fileSimpleStorageTestSuite) TearDownTest() {
	err := os.RemoveAll(s.fss.opt.Home)
	s.Nil(err)
}

func (s *fileSimpleStorageTestSuite) TestGetObjectContent() {
	ch, err := s.fss.GetObjectContent(test_object)
	s.Nil(err)

	content := ""
	for buf, ok := <-ch; ok; buf, ok = <-ch {
		content += string(buf)
	}

	s.Equal(test_object_content, content)
}

func (s *fileSimpleStorageTestSuite) TestGetObjectContentSync() {
	content, err := s.fss.GetObjectContentSync(test_object)
	s.Nil(err)

	s.Equal(test_object_content, string(content))
}

func (s *fileSimpleStorageTestSuite) TestGetObjectContentWithLargeFile() {
	for i := 0; i < _TEST_LARGE_OBJECT_CONTENT_SIZE; i++ {
		test_large_object_content[i] = byte(i % (_TEST_LARGE_OBJECT_CONTENT_MASK + 1))
	}

	err := s.fss.PutObject(test_large_object, bytes.NewReader(test_large_object_content))
	s.Require().Nil(err)

	ch, err := s.fss.GetObjectContent(test_large_object)
	s.Require().Nil(err)

	var sb strings.Builder
	for buf, ok := <-ch; ok; buf, ok = <-ch {
		sb.Write(buf)
	}

	s.Equal(test_large_object_content, []byte(sb.String()))
}

func (s *fileSimpleStorageTestSuite) TestGetObject() {
	obj, err := s.fss.GetObject(test_object)
	s.Nil(err)
	s.Equal(test_device_id, obj.Device)
	s.Equal(test_object_prefix, obj.Prefix)
	s.Equal(test_object_name, obj.Name)
	s.Equal(int64(len(test_object_content)), obj.Length)
}

func (s *fileSimpleStorageTestSuite) TestPutObject() {
	txt := "hello, world"
	obj := NewObject(test_device_id, "", "test/test/test.txt")
	err := s.fss.PutObject(obj, strings.NewReader(txt))
	s.Nil(err)

	obj1, err := s.fss.GetObject(obj)
	s.Nil(err)

	s.Equal("test/test", obj1.Prefix)
	s.Equal("test.txt", obj1.Name)
	s.Equal(int64(len(txt)), obj1.Length)
}

func (s *fileSimpleStorageTestSuite) TestRemoveObject() {
	err := s.fss.RemoveObject(test_object)
	s.Nil(err)

	_, err = s.fss.GetObject(test_object)
	s.NotNil(err)
}

func (s *fileSimpleStorageTestSuite) TestRenameObject() {
	dst := NewObject(test_device_id, "dst", "test.txt")
	err := s.fss.RenameObject(test_object, dst)
	s.Nil(err)

	_, err = s.fss.GetObject(test_object)
	s.NotNil(err)

	_, err = s.fss.GetObject(dst)
	s.Nil(err)
}

func (s *fileSimpleStorageTestSuite) TestListObjects() {
	fltr := NewObject(test_device_id, test_object_prefix, "")
	objs, err := s.fss.ListObjects(fltr, &ListObjectsOption{})
	s.Nil(err)
	s.Len(objs, 1)
}

func TestFileSimpleStorageTestSuite(t *testing.T) {
	suite.Run(t, new(fileSimpleStorageTestSuite))
}
