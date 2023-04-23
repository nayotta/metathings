package metathings_deviced_simple_storage

import (
	"bytes"
	"context"
	"strings"
	"testing"
	"time"

	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	logging "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"

	test_helper "github.com/nayotta/metathings/pkg/common/test"
)

type minioSimpleStorageTestSuite struct {
	mc     *minio.Client
	mss    *MinioSimpleStorage
	bucket string
	suite.Suite
}

func (ts *minioSimpleStorageTestSuite) TestGetObjectContent() {
	bufs, err := ts.mss.GetObjectContent(test_object)
	ts.Nil(err)

	var bb bytes.Buffer
	for buf := range bufs {
		bb.Write(buf)
	}

	ts.Equal(test_object_content, bb.String())
}

func (ts *minioSimpleStorageTestSuite) TestGetObjectContentSync() {
	content, err := ts.mss.GetObjectContentSync(test_object)
	ts.Nil(err)
	ts.Equal(test_object_content, string(content))
}

func (ts *minioSimpleStorageTestSuite) TestGetObjectContentWithLargeFile() {
	for i := 0; i < _TEST_LARGE_OBJECT_CONTENT_SIZE; i++ {
		test_large_object_content[i] = byte(i % (_TEST_LARGE_OBJECT_CONTENT_MASK + 1))
	}

	err := ts.mss.PutObject(test_large_object, bytes.NewReader(test_large_object_content))
	ts.Require().Nil(err)

	bufs, err := ts.mss.GetObjectContent(test_large_object)
	ts.Require().Nil(err)

	var bb bytes.Buffer
	for buf := range bufs {
		bb.Write(buf)
	}

	ts.Equal(test_large_object_content, bb.Bytes())
}

func (ts *minioSimpleStorageTestSuite) TestGetObject() {
	obj, err := ts.mss.GetObject(test_object)
	ts.Nil(err)
	ts.Equal(test_device_id, obj.Device)
	ts.Equal(test_object_prefix, obj.Prefix)
	ts.Equal(test_object_name, obj.Name)
	ts.Equal(len(test_object_content), int(obj.Length))
}

func (ts *minioSimpleStorageTestSuite) TestPutObject() {
	txt := "hello, world"
	obj := new_object(test_device_id, "", "x/y/z/test.txt", int64(len(txt)), "", time.Time{})

	err := ts.mss.PutObject(obj, strings.NewReader(txt))
	ts.Nil(err)

	obj1, err := ts.mss.GetObject(obj)
	ts.Nil(err)
	ts.Equal("x/y/z", obj1.Prefix)
	ts.Equal("test.txt", obj1.Name)
	ts.Equal(len(txt), int(obj1.Length))
}

func (ts *minioSimpleStorageTestSuite) TestRemoveObject() {
	err := ts.mss.RemoveObject(test_object)
	ts.Nil(err)

	_, err = ts.mss.GetObject(test_object)
	ts.NotNil(err)
}

func (ts *minioSimpleStorageTestSuite) TestRenameObject() {
	dst := NewObject(test_device_id, "dst", "test.txt")
	err := ts.mss.RenameObject(test_object, dst)
	ts.Nil(err)

	_, err = ts.mss.GetObject(test_object)
	ts.NotNil(err)

	_, err = ts.mss.GetObject(dst)
	ts.Nil(err)
}

func (ts *minioSimpleStorageTestSuite) SetupSuite() {
	endpoint := test_helper.GetTestMinioEndpoint()
	id := test_helper.GetTestMinioID()
	secret := test_helper.GetTestMinioSecret()
	token := test_helper.GetTestMinioToken()
	secure := test_helper.GetTestMinioSecure()
	minioBucket := test_helper.GetTestMinioBucket()
	rdBufSz := test_helper.GetTestMinioReadBufferSize()
	wrBufSz := test_helper.GetTestMinioWriteBufferSize()

	ts.bucket = minioBucket

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(id, secret, token),
		Secure: secure,
	})
	ts.Require().Nil(err)
	ts.mc = client

	logger := logging.New()
	logger.SetLevel(logging.TraceLevel)

	mss, err := new_minio_simple_storage(
		"minio_endpoint", endpoint,
		"minio_id", id,
		"minio_secret", secret,
		"minio_token", token,
		"minio_secure", secure,
		"minio_bucket", minioBucket,

		"read_buffer_size", rdBufSz,
		"write_buffer_size", wrBufSz,

		"logger", logging.NewEntry(logger),
	)
	ts.Require().Nil(err)
	ts.mss = mss.(*MinioSimpleStorage)
}

func (ts *minioSimpleStorageTestSuite) SetupTest() {
	ts.mc.MakeBucket(ts.mss.context(), ts.bucket, minio.MakeBucketOptions{})
	err := ts.mss.PutObject(test_object, strings.NewReader(test_object_content))
	ts.Nil(err)
}

func (ts *minioSimpleStorageTestSuite) TearDownTest() {
	ts.mc.RemoveBucket(ts.context(), ts.bucket)
}

func TestMinioSimpleStorageTestSuite(t *testing.T) {
	suite.Run(t, new(minioSimpleStorageTestSuite))
}

func (ts *minioSimpleStorageTestSuite) context() context.Context {
	return context.Background()
}
