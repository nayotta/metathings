package metathings_deviced_simple_storage

import (
	"io"
	"path"
	"time"

	storage "github.com/nayotta/metathings/pkg/deviced/storage"
)

type Object struct {
	Prefix       string
	Name         string
	Length       int64
	Etag         string
	LastModified time.Time
	Metadata     map[string]string
}

func NewObject(prefix string, name string, metadata map[string]string) *Object {
	prefix = path.Join(prefix, path.Dir(name))
	name = path.Clean(path.Base(name))

	return &Object{
		Prefix:   prefix,
		Name:     name,
		Metadata: metadata,
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

var simple_storage_factories map[string]SimpleStorageFactory

func NewSimpleStorage(name string, args ...interface{}) (SimpleStorage, error) {
	var fty SimpleStorageFactory
	var ok bool

	if fty, ok = simple_storage_factories[name]; !ok {
		return nil, ErrInvalidSimpleStorageDriver
	}

	return fty(args...)
}

func init() {
	simple_storage_factories = make(map[string]SimpleStorageFactory)
}
