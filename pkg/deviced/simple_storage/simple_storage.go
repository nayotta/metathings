package metathings_deviced_simple_storage

import (
	"io"
	"path"
	"time"

	storage "github.com/nayotta/metathings/pkg/deviced/storage"
)

type Object struct {
	Device       string
	Prefix       string
	Name         string
	Length       int64
	Etag         string
	LastModified time.Time
	Metadata     map[string]string
}

func (o *Object) FullName() string {
	return path.Join(o.Prefix, o.Name)
}

func NewObject(device, prefix, name string, metadata map[string]string) *Object {
	prefix = path.Join(prefix, path.Dir(name))
	name = path.Clean(path.Base(name))

	return &Object{
		Device:   device,
		Prefix:   prefix,
		Name:     name,
		Metadata: metadata,
	}
}

func new_object(device, prefix, name string, length int64, etag string, last_modified time.Time, metadata map[string]string) *Object {
	return &Object{
		Device:       device,
		Prefix:       prefix,
		Name:         name,
		Length:       length,
		Etag:         etag,
		LastModified: last_modified,
		Metadata:     metadata,
	}
}

type SimpleStorage interface {
	PutObject(dev *storage.Device, obj *Object, reader io.Reader) error
	RemoveObject(dev *storage.Device, obj *Object) error
	RenameObject(dev *storage.Device, src, dst *Object) error
	GetObject(dev *storage.Device, obj *Object) (*Object, error)
	GetObjectContent(dev *storage.Device, obj *Object) (chan []byte, error)
	ListObjects(dev *storage.Device, obj *Object) ([]*Object, error)
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
