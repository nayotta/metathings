package metathings_deviced_simple_storage

import (
	"bytes"
	"context"
	"io"
	"path"
	"path/filepath"
	"strings"

	minio "github.com/minio/minio-go/v7"
	logging "github.com/sirupsen/logrus"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	file_helper "github.com/nayotta/metathings/pkg/common/file"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type MinioSimpleStorageOption struct {
	Bucket          string
	ReadBufferSize  int
	WriteBufferSize int
}

func NewMinioSimpleStorageOption() *MinioSimpleStorageOption {
	return &MinioSimpleStorageOption{
		ReadBufferSize:  4 * 1024 * 1024,
		WriteBufferSize: 4 * 1024 * 1024,
	}
}

type MinioSimpleStorage struct {
	minioClient *minio.Client
	opt         *MinioSimpleStorageOption
	logger      logging.FieldLogger
}

func new_minio_simple_storage(args ...any) (SimpleStorage, error) {
	var logger logging.FieldLogger
	var minioClient *minio.Client
	opt := NewMinioSimpleStorageOption()

	err := opt_helper.Setopt(map[string]func(string, any) error{
		"bucket":            opt_helper.ToString(&opt.Bucket),
		"read_buffer_size":  opt_helper.ToInt(&opt.ReadBufferSize),
		"write_buffer_size": opt_helper.ToInt(&opt.WriteBufferSize),
		"minio_client":      client_helper.ToMinioClient(&minioClient),
		"logger":            opt_helper.ToLogger(&logger),
	}, opt_helper.SetSkip(true))(args...)
	if err != nil {
		return nil, err
	}

	return &MinioSimpleStorage{
		opt:         opt,
		minioClient: minioClient,
		logger:      logger,
	}, nil
}

func (mss *MinioSimpleStorage) PutObject(obj *Object, rd io.Reader) error {
	logger := mss.GetLoggerWithObject(obj).WithField("#method", "PutObject")
	ctx := mss.context()
	fp := mss.join_path(obj)

	_, err := mss.minioClient.PutObject(ctx, mss.minioBucket(), fp, rd, obj.Length, minio.PutObjectOptions{})
	if err != nil {
		logger.WithError(err).Debugf("failed to put object to minio")
		return err
	}

	logger.Tracef("put object")

	return nil
}

func (mss *MinioSimpleStorage) PutObjectAsync(obj *Object, opt *PutObjectAsyncOption) (file_helper.FileSyncer, error) {
	logger := mss.GetLoggerWithObject(obj).WithField("#method", "PutObjectAsync")
	ctx := mss.context()

	rd, wr := io.Pipe()

	fs := file_helper.NewSequenceFileSyncer(wr, obj.Length, opt.SHA1, int64(mss.opt.WriteBufferSize))
	_, err := mss.minioClient.PutObject(ctx, mss.minioBucket(), mss.join_path(obj), rd, obj.Length, minio.PutObjectOptions{})
	if err != nil {
		logger.WithError(err).Debugf("failed to put object async to minio")
		return nil, err
	}

	logger.Tracef("put object async")

	return fs, nil
}

func (mss *MinioSimpleStorage) RemoveObject(obj *Object) error {
	logger := mss.GetLoggerWithObject(obj).WithField("#method", "RemoveObject")
	ctx := mss.context()
	fp := mss.join_path(obj)

	err := mss.minioClient.RemoveObject(ctx, mss.minioBucket(), fp, minio.RemoveObjectOptions{ForceDelete: true})
	if err != nil {
		logger.WithError(err).Debugf("failed to remove object from minio")
		return err
	}

	logger.Tracef("remove object")

	return nil
}

func (mss *MinioSimpleStorage) RenameObject(src, dst *Object) error {
	logger := mss.loggerWithObject(
		mss.loggerWithObject(mss.GetLogger(), "source.", src),
		"destination.", dst).WithField("#method", "RenameObject")
	ctx := mss.context()

	_, err := mss.minioClient.CopyObject(ctx, minio.CopyDestOptions{
		Bucket: mss.minioBucket(),
		Object: mss.join_path(dst),
	}, minio.CopySrcOptions{
		Bucket: mss.minioBucket(),
		Object: mss.join_path(src),
	})
	if err != nil {
		logger.WithError(err).Debugf("failed to copy source to destination")
		return err
	}

	if err = mss.minioClient.RemoveObject(ctx, mss.minioBucket(), mss.join_path(src), minio.RemoveObjectOptions{}); err != nil {
		logger.WithError(err).Debugf("failed to remove source object")
	} else {
		logger.Tracef("rename object")
	}

	return nil
}

func (mss *MinioSimpleStorage) GetObject(x *Object) (y *Object, err error) {
	logger := mss.GetLoggerWithObject(x).WithField("#method", "GetObject")
	ctx := mss.context()

	obj, err := mss.minioClient.GetObject(ctx, mss.minioBucket(), mss.join_path(x), minio.GetObjectOptions{})
	if err != nil {
		logger.WithError(err).Debugf("failed to get objectc from minio")
		return nil, err
	}

	info, err := obj.Stat()
	if err != nil {
		logger.WithError(err).Debugf("failed to get object stat from minio")
		return nil, err
	}

	y, err = mss.new_object_from_minio_object_info(info)
	if err != nil {
		logger.WithError(err).Debugf("failed to new object from minio object stat")
		return nil, err
	}

	logger.Tracef("get object")

	return
}

func (mss *MinioSimpleStorage) GetObjectContent(obj *Object) (chan []byte, error) {
	return mss.get_object_content(obj)
}

func (mss *MinioSimpleStorage) GetObjectContentSync(obj *Object) ([]byte, error) {
	var bb bytes.Buffer

	ch, err := mss.get_object_content(obj)
	if err != nil {
		return nil, err
	}

	for buf := range ch {
		bb.Write(buf)
	}

	return bb.Bytes(), nil
}

// ignore option:depth in minio backend
// dont list directory in minio backend
func (mss *MinioSimpleStorage) ListObjects(obj *Object, opt *ListObjectsOption) ([]*Object, error) {
	logger := mss.GetLoggerWithObject(obj).WithFields(logging.Fields{
		"#method":   "ListObjects",
		"recursive": opt.Recursive,
	})
	ctx := mss.context()

	ois := mss.minioClient.ListObjects(ctx, mss.minioBucket(), minio.ListObjectsOptions{
		Prefix:    obj.Prefix,
		Recursive: opt.Recursive,
	})

	var objs []*Object
	for oi := range ois {
		obj, err := mss.new_object_from_minio_object_info(oi)
		if err != nil {
			logger.WithError(err).Debugf("failed to new object from minio object info")
			return nil, err
		}
		objs = append(objs, obj)
	}

	logger.Tracef("list objects")

	return objs, nil
}

func (mss *MinioSimpleStorage) GetLogger() logging.FieldLogger {
	return mss.logger.WithFields(logging.Fields{
		"#instance": "MinioSimpleStorage",
		"bucket":    mss.opt.Bucket,
	})
}

func (mss *MinioSimpleStorage) loggerWithObject(logger logging.FieldLogger, objectPrefix string, object *Object) logging.FieldLogger {
	fp := mss.join_path(object)
	return logger.WithFields(logging.Fields{
		objectPrefix + "device": object.Device,
		objectPrefix + "prefix": filepath.Dir(fp),
		objectPrefix + "file":   filepath.Base(fp),
	})
}

func (mss *MinioSimpleStorage) GetLoggerWithObject(obj *Object) logging.FieldLogger {
	return mss.loggerWithObject(mss.GetLogger(), "", obj)
}

func (mss *MinioSimpleStorage) minioBucket() string {
	return mss.opt.Bucket
}

func (mss *MinioSimpleStorage) join_path(obj *Object) string {
	return path.Join(obj.Device, obj.FullName())
}

func (mss *MinioSimpleStorage) context() context.Context {
	return context.Background()
}

func (mss *MinioSimpleStorage) new_object_from_minio_object_info(oi minio.ObjectInfo) (*Object, error) {
	ss := strings.SplitN(oi.Key, "/", 2)
	if len(ss) != 2 {
		return nil, ErrObjectNotFound
	}

	device := ss[0]
	prefix := path.Dir(ss[1])
	base := path.Base(ss[1])

	return new_object(device, prefix, base, oi.Size, oi.ETag, oi.LastModified), nil
}

func (mss *MinioSimpleStorage) get_object_content(obj *Object) (chan []byte, error) {
	logger := mss.GetLoggerWithObject(obj).WithField("#method", "get_object_content")
	ctx := mss.context()

	minioObject, err := mss.minioClient.GetObject(ctx, mss.minioBucket(), mss.join_path(obj), minio.GetObjectOptions{})
	if err != nil {
		logger.WithError(err).Debugf("failed to get object from minio")
		return nil, err
	}

	ch := make(chan []byte)
	go func() {
		defer close(ch)
		for {
			slice := make([]byte, mss.opt.ReadBufferSize)
			n, err := minioObject.Read(slice)
			if n > 0 {
				ch <- slice[:n]
			}

			if err != nil || n == 0 {
				break
			}
		}
	}()

	logger.Tracef("get object content")

	return ch, nil
}

func init() {
	register_simple_storage_factory("minio", new_minio_simple_storage)
}
