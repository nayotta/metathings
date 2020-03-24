package metathings_data_storage_sdk

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockDataStorage struct {
	mock.Mock
}

func (m *MockDataStorage) Write(ctx context.Context, measurement string, tags map[string]string, data map[string]interface{}) error {
	m.Called(ctx, measurement, tags, data)
	return nil
}
