package metathings_data_storage_sdk

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockDataStorage struct {
	mock.Mock
	query_result QueryResult
}

func (m *MockDataStorage) Write(ctx context.Context, measurement string, tags map[string]string, data map[string]interface{}) error {
	m.Called(ctx, measurement, tags, data)
	return nil
}

func (m *MockDataStorage) SetQueryResult(qr QueryResult) {
	m.query_result = qr
}

func (m *MockDataStorage) Query(ctx context.Context, measurement string, tags map[string]string, opts ...QueryOption) (QueryResult, error) {
	o := map[string]interface{}{}
	for _, opt := range opts {
		opt(o)
	}

	m.Called(ctx, measurement, tags, o)

	return m.query_result, nil
}
