package metathings_deviced_sdk

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockCaller struct {
	mock.Mock

	unary_call_result map[string]interface{}
}

func (m *MockCaller) UnaryCall(ctx context.Context, device, module, method string, arguments map[string]interface{}) (map[string]interface{}, error) {
	m.Called(ctx, device, module, method, arguments)
	return m.unary_call_result, nil
}
