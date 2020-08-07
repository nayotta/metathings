package metathings_deviced_sdk

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockFlow struct {
	mock.Mock
}

func (m *MockFlow) PushFrame(ctx context.Context, device, flow string, data interface{}) error {
	m.Called(ctx, device, flow, data)
	return nil
}
