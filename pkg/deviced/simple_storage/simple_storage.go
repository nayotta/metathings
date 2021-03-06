package metathings_deviced_simple_storage

import (
	"io"
	"path"
	"time"

	file_helper "github.com/nayotta/metathings/pkg/common/file"
)

type Object struct {
	Device       string
	Prefix       string
	Name         string
	Length       int64
	Etag         string
	LastModified time.Time
}

func (o *Object) FullName() string {
	return path.Join(o.Prefix, o.Name)
}

func NewObject(device, prefix, name string) *Object {
	prefix = path.Join(prefix, path.Dir(name))
	name = path.Clean(path.Base(name))

	return &Object{
		Device: device,
		Prefix: prefix,
		Name:   name,
	}
}

func new_object(device, prefix, name string, length int64, etag string, last_modified time.Time) *Object {
	return &Object{
		Device:       device,
		Prefix:       prefix,
		Name:         name,
		Length:       length,
		Etag:         etag,
		LastModified: last_modified,
	}
}

type PutObjectAsyncOption struct {
	SHA1      string
	ChunkSize int64
}

type ListObjectsOption struct {
	Recursive bool
	Depth     int
}

type SimpleStorage interface {
	PutObject(obj *Object, reader io.Reader) error
	PutObjectAsync(obj *Object, opt *PutObjectAsyncOption) (*file_helper.FileSyncer, error)
	RemoveObject(obj *Object) error
	RenameObject(src, dst *Object) error
	GetObject(obj *Object) (*Object, error)
	GetObjectContent(obj *Object) (chan []byte, error)
	GetObjectContentSync(obj *Object) ([]byte, error)
	ListObjects(obj *Object, opt *ListObjectsOption) ([]*Object, error)
}

type SimpleStorageFactory func(...interface{}) (SimpleStorage, error)

var simple_storage_factories = make(map[string]SimpleStorageFactory)

func register_simple_storage_factory(name string, fty SimpleStorageFactory) {
	simple_storage_factories[name] = fty
}

func NewSimpleStorage(name string, args ...interface{}) (SimpleStorage, error) {
	var fty SimpleStorageFactory
	var ok bool

	if fty, ok = simple_storage_factories[name]; !ok {
		return nil, ErrInvalidSimpleStorageDriver
	}

	return fty(args...)
}
