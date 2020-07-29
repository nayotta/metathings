package metathings_sms_sdk

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockSmsSender struct {
	mock.Mock
}

func (m *MockSmsSender) SendSms(ctx context.Context, id string, numbers []string, arguments map[string]string) error {
	m.Called(ctx, id, numbers, arguments)
	return nil
}
