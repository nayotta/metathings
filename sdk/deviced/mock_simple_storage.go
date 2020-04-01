package metathings_deviced_sdk

import (
	"context"
	"io"
	"io/ioutil"

	"github.com/stretchr/testify/mock"

	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

type MockSimpleStorage struct {
	mock.Mock

	get_result         *pb.Object
	get_content_result []byte
	list_result        []*pb.Object
}

func (m *MockSimpleStorage) SetGetResult(obj *pb.Object) {
	m.get_result = obj
}

func (m *MockSimpleStorage) SetGetContentResult(buf []byte) {
	m.get_content_result = buf
}

func (m *MockSimpleStorage) SetListResult(objs []*pb.Object) {
	m.list_result = objs
}

func (m *MockSimpleStorage) Put(ctx context.Context, obj *pb.Object, rd io.Reader) error {
	buf, _ := ioutil.ReadAll(rd)
	m.Called(ctx, pb_helper.ToStringMap(obj), string(buf))
	return nil
}

func (m *MockSimpleStorage) Remove(ctx context.Context, obj *pb.Object) error {
	m.Called(ctx, pb_helper.ToStringMap(obj))
	return nil
}

func (m *MockSimpleStorage) Rename(ctx context.Context, src, dst *pb.Object) error {
	m.Called(ctx, pb_helper.ToStringMap(src), pb_helper.ToStringMap(dst))
	return nil
}

func (m *MockSimpleStorage) Get(ctx context.Context, obj *pb.Object) (*pb.Object, error) {
	m.Called(ctx, pb_helper.ToStringMap(obj))
	return m.get_result, nil
}

func (m *MockSimpleStorage) GetContent(ctx context.Context, obj *pb.Object) ([]byte, error) {
	m.Called(ctx, pb_helper.ToStringMap(obj))
	return m.get_content_result, nil
}

func (m *MockSimpleStorage) List(ctx context.Context, obj *pb.Object, opts ...SimpleStorageListOption) ([]*pb.Object, error) {
	o := pb_helper.ToStringMap(obj)
	for _, opt := range opts {
		opt(o)
	}
	m.Called(ctx, o)
	return m.list_result, nil
}
