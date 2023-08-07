package metathings_deviced_simple_storage

import (
	"bytes"
	"context"
	"io"
	"path"
	"path/filepath"
	"strings"

	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	logging "github.com/sirupsen/logrus"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	file_helper "github.com/nayotta/metathings/pkg/common/file"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type MinioSimpleStorageOption struct {
	MinioEndpoint string
	MinioID       string
	MinioSecret   string
	MinioToken    string
	MinioSecure   bool
	MinioBucket   string

	ReadBufferSize int
}

func NewMinioSimpleStorageOption() *MinioSimpleStorageOption {
	return &MinioSimpleStorageOption{
		ReadBufferSize: 1 * 1024 * 1024,
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
		"minio_endpoint":   opt_helper.ToString(&opt.MinioEndpoint),
		"minio_id":         opt_helper.ToString(&opt.MinioID),
		"minio_secret":     opt_helper.ToString(&opt.MinioSecret),
		"minio_token":      opt_helper.ToString(&opt.MinioToken),
		"minio_secure":     opt_helper.ToBool(&opt.MinioSecure),
		"minio_bucket":     opt_helper.ToString(&opt.MinioBucket),
		"read_buffer_size": opt_helper.ToInt(&opt.ReadBufferSize),
		"minio_client":     client_helper.ToMinioClient(&minioClient),
		"logger":           opt_helper.ToLogger(&logger),
	}, opt_helper.SetSkip(true))(args...)
	if err != nil {
		return nil, err
	}

	if minioClient == nil {
		minioClient, err = minio.New(opt.MinioEndpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(opt.MinioID, opt.MinioSecret, opt.MinioToken),
			Secure: opt.MinioSecure,
		})
		if err != nil {
			return nil, err
		}
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

	rd, wr := WrapedPipe(mss.logger)
	fs := file_helper.NewSequenceFileSyncer(wr, obj.Length, opt.SHA1, opt.ChunkSize)
	go mss.minioClient.PutObject(ctx, mss.minioBucket(), mss.join_path(obj), rd, obj.Length, minio.PutObjectOptions{})

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

func (mss *MinioSimpleStorage) ListObjects(x *Object, opt *ListObjectsOption) (ys []*Object, err error) {
	logger := mss.GetLoggerWithObject(x).WithFields(logging.Fields{
		"#method":   "ListObjects",
		"recursive": opt.Recursive,
		"depth":     opt.Depth,
	})

	depth := opt.Depth
	if !opt.Recursive {
		depth = 1
	}

	ys, err = mss.list_objects(mss.context(), mss.minioClient, mss.minioBucket(), x, depth)
	if err != nil {
		logger.WithError(err).Debugf("failed to list objects")
		return nil, err
	}

	logger.Tracef("list objects")
	return ys, nil
}

func (mss *MinioSimpleStorage) GetLogger() logging.FieldLogger {
	return mss.logger.WithFields(logging.Fields{
		"#instance": "MinioSimpleStorage",
		"bucket":    mss.opt.MinioBucket,
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
	return mss.opt.MinioBucket
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
	if prefix == "." {
		prefix = ""
	}
	base := ""
	if len(ss[1]) != 0 && ss[1][len(ss[1])-1] != '/' {
		base = path.Base(ss[1])
	}

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

func (mss *MinioSimpleStorage) list_objects(ctx context.Context, mc *minio.Client, bkt string, x *Object, d int) (ys []*Object, err error) {
	ys = []*Object{}

	if d <= 0 {
		return
	}

	ois := mc.ListObjects(ctx, bkt, minio.ListObjectsOptions{
		Prefix:    mss.join_path(x) + "/",
		Recursive: false,
	})
	for oi := range ois {
		y, err := mss.new_object_from_minio_object_info(oi)
		if err != nil {
			return nil, err
		}

		if y.Name == "." {
			continue
		}

		ys = append(ys, y)

		if mss.is_directory(y) {
			zs, err := mss.list_objects(ctx, mc, bkt, y, d-1)
			if err != nil {
				return nil, err
			}
			ys = append(ys, zs...)
		}
	}

	return ys, nil
}

func (mss *MinioSimpleStorage) is_directory(obj *Object) bool {
	return obj.Etag == ""
}

func init() {
	register_simple_storage_factory("minio", new_minio_simple_storage)
}
