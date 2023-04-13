package metathings_deviced_simple_storage

import (
	"io"

	minio "github.com/minio/minio-go/v7"
	log "github.com/sirupsen/logrus"

	file_helper "github.com/nayotta/metathings/pkg/common/file"
)

type MinioSimpleStorageOption struct {
}

type MinioSimpleStorage struct {
	minioClient *minio.Client
	opt         *MinioSimpleStorageOption
	logger      log.FieldLogger
}

func (mss *MinioSimpleStorage) PutObject(obj *Object, rd io.Reader) error {
	panic("unimplemented")
}

func (mss *MinioSimpleStorage) PutObjectAsync(obj *Object, opt *PutObjectAsyncOption) (*file_helper.FileSyncer, error) {
	panic("unimplemented")
}

func (mss *MinioSimpleStorage) RemoveObject(obj *Object) error {
	panic("unimplemented")
}

func (mss *MinioSimpleStorage) RenameObject(src, dst *Object) error {
	panic("unimplemented")
}

func (mss *MinioSimpleStorage) GetObject(obj *Object) (*Object, error) {
	panic("unimplemented")
}

func (mss *MinioSimpleStorage) GetObjectContent(obj *Object) (chan []byte, error) {
	panic("unimplemented")
}

func (mss *MinioSimpleStorage) GetObjectContentSync(obj *Object) ([]byte, error) {
	panic("unimplemented")
}

func (mss *MinioSimpleStorage) ListObjects(obj *Object, opt *ListObjectsOption) ([]*Object, error) {
	panic("unimplemented")
}
