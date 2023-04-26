package metathings_deviced_simple_storage

import (
	"bytes"
	"context"
	"path"
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

func (ts *minioSimpleStorageTestSuite) TestListObjects1() {
	x := NewObject(test_device_id, "", "")
	ys, err := ts.mss.ListObjects(x, &ListObjectsOption{
		Recursive: false,
	})
	ts.Nil(err)
	ts.Len(ys, 1)
	ts.Equal(test_device_id, ys[0].Device)
	ts.Equal(test_object_prefix, ys[0].Prefix)
	ts.Equal("", ys[0].Name)
}

func (ts *minioSimpleStorageTestSuite) TestListObjects2() {
	x := NewObject(test_device_id, test_object_prefix, "")
	ys, err := ts.mss.ListObjects(x, &ListObjectsOption{
		Recursive: false,
	})
	ts.Nil(err)
	ts.Len(ys, 1)
	ts.Equal(test_device_id, ys[0].Device)
	ts.Equal(test_object_prefix, ys[0].Prefix)
	ts.Equal(test_object_name, ys[0].Name)
}

func (ts *minioSimpleStorageTestSuite) TestListObjects3() {
	ts.removeAll()

	txt := "hello, world!"
	xa := NewObject(test_device_id, "x", "a.txt")
	xa.Length = int64(len(txt))
	err := ts.mss.PutObject(xa, strings.NewReader(txt))
	ts.Require().Nil(err)
	xb := NewObject(test_device_id, "x", "b.txt")
	xb.Length = xa.Length
	err = ts.mss.PutObject(xb, strings.NewReader(txt))
	ts.Require().Nil(err)
	xya := NewObject(test_device_id, "x/y", "a.txt")
	xya.Length = xa.Length
	err = ts.mss.PutObject(xya, strings.NewReader(txt))
	ts.Require().Nil(err)

	os, err := ts.mss.ListObjects(NewObject(test_device_id, "", ""), &ListObjectsOption{
		Recursive: true,
		Depth:     1,
	})
	ts.Nil(err)
	// x/
	ts.Len(os, 1)

	os, err = ts.mss.ListObjects(NewObject(test_device_id, "x", ""), &ListObjectsOption{
		Recursive: true,
		Depth:     1,
	})
	ts.Nil(err)
	// x/a.txt, x/b.txt, x/y/
	ts.Len(os, 3)

	os, err = ts.mss.ListObjects(NewObject(test_device_id, "", ""), &ListObjectsOption{
		Recursive: true,
		Depth:     16,
	})
	ts.Nil(err)
	// x/, x/a.txt, x/b.txt, x/y/, x/y/a.txt
	ts.Len(os, 5)
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
	err := ts.mc.MakeBucket(ts.mss.context(), ts.bucket, minio.MakeBucketOptions{})
	ts.Require().Nil(err)
	_, err = ts.mc.PutObject(ts.context(), ts.bucket, path.Join(test_device_id, test_object_prefix, test_object_name), strings.NewReader(test_object_content), int64(len(test_object_content)), minio.PutObjectOptions{})
	ts.Require().Nil(err)
}

func (ts *minioSimpleStorageTestSuite) TearDownTest() {
	ts.removeAll()
	ts.mc.RemoveBucket(ts.context(), ts.bucket)
}

func TestMinioSimpleStorageTestSuite(t *testing.T) {
	suite.Run(t, new(minioSimpleStorageTestSuite))
}

func (ts *minioSimpleStorageTestSuite) context() context.Context {
	return context.Background()
}

func (ts *minioSimpleStorageTestSuite) removeAll() {
	ois := make(chan minio.ObjectInfo)
	go func() {
		defer close(ois)
		for oi := range ts.mc.ListObjects(ts.context(), ts.bucket, minio.ListObjectsOptions{
			Recursive: true,
			Prefix:    "",
		}) {
			ts.Require().Nil(oi.Err)
			ois <- oi
		}
	}()
	errCh := ts.mc.RemoveObjects(ts.context(), ts.bucket, ois, minio.RemoveObjectsOptions{GovernanceBypass: true})
	for err := range errCh {
		ts.T().Logf("remove object: %v\n", err)
	}
}
