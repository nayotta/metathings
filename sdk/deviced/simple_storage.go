package metathings_deviced_sdk

import (
	"context"
	"io"
	"sync"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

type SimpleStorage interface {
	Put(context.Context, *pb.Object, io.Reader) error
	Remove(context.Context, *pb.Object) error
	Rename(ctx context.Context, src *pb.Object, dst *pb.Object) error
	Get(context.Context, *pb.Object) (*pb.Object, error)
	GetContent(context.Context, *pb.Object) ([]byte, error)
	List(context.Context, *pb.Object, ...SimpleStorageListOption) ([]*pb.Object, error)
}

type SimpleStorageListOption func(map[string]interface{})

var (
	SimpleStorageListOption_SetRecursive = SetBool("recursive")
	SimpleStorageListOption_SetDepth     = SetInt("depth")
)

type SimpleStorageFactory func(...interface{}) (SimpleStorage, error)

var simple_storage_factories_once sync.Once
var simple_storage_factories map[string]SimpleStorageFactory

func register_simple_storage_factory(name string, fty SimpleStorageFactory) {
	simple_storage_factories_once.Do(func() {
		simple_storage_factories = make(map[string]SimpleStorageFactory)
	})

	simple_storage_factories[name] = fty
}

func NewSimpleStorage(name string, args ...interface{}) (SimpleStorage, error) {
	fty, ok := simple_storage_factories[name]
	if !ok {
		return nil, ErrUnsupportedSimpleStorageFactory
	}

	return fty(args...)
}

func ToSimpleStorage(v *SimpleStorage) func(string, interface{}) error {
	return func(key string, val interface{}) error {
		var ok bool

		if *v, ok = val.(SimpleStorage); !ok {
			return opt_helper.InvalidArgument(key)
		}

		return nil
	}
}
