package grpc_helper

import (
	"context"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

type unary_test_server struct {
	mock.Mock
}

func (self *unary_test_server) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	self.Called(ctx, fullMethodName)
	return ctx, nil
}

func (self *unary_test_server) ValidateFuncTest(ctx context.Context, in interface{}) error {
	self.Called(ctx, in)
	return nil
}

func (self *unary_test_server) AuthorizeFuncTest(ctx context.Context, in interface{}) error {
	self.Called(ctx, in)
	return nil
}

type unaryServerInterceptorTestSuite struct {
	suite.Suite

	server       *unary_test_server
	ctx          context.Context
	interceptor  grpc.UnaryServerInterceptor
	req          *struct{}
	full_method  string
	info         *grpc.UnaryServerInfo
	handler      grpc.UnaryHandler
	handler_mock mock.Mock
}

func (self *unaryServerInterceptorTestSuite) SetupTest() {
	self.server = &unary_test_server{}
	self.ctx = context.Background()
	self.interceptor = UnaryServerInterceptor(logrus.New())
	self.req = &struct{}{}
	self.full_method = "/ai.metathings.service.testd.TestdService/FuncTest"
	self.info = &grpc.UnaryServerInfo{
		Server:     self.server,
		FullMethod: self.full_method,
	}
	self.handler_mock = mock.Mock{}
	self.handler = func(ctx context.Context, in interface{}) (interface{}, error) {
		self.handler_mock.Called(ctx, in)
		return in, nil
	}

	self.server.On("AuthFuncOverride", self.ctx, self.full_method).Return(self.ctx, nil)
	self.server.On("ValidateFuncTest", self.ctx, self.req).Return(nil)
	self.server.On("AuthorizeFuncTest", self.ctx, self.req).Return(nil)
	self.handler_mock.On("func1", self.ctx, self.req).Return(self.req, nil)
}

func (self *unaryServerInterceptorTestSuite) TestInterceptor() {
	res, err := self.interceptor(self.ctx, self.req, self.info, self.handler)

	self.Nil(err)
	self.Equal(res, self.req)

	self.handler_mock.AssertCalled(self.T(), "func1", self.ctx, self.req)
	self.server.AssertCalled(self.T(), "AuthFuncOverride", self.ctx, self.full_method)
	self.server.AssertCalled(self.T(), "ValidateFuncTest", self.ctx, self.req)
	self.server.AssertCalled(self.T(), "AuthorizeFuncTest", self.ctx, self.req)
}

func TestUnaryServerInterceptorTestSuite(t *testing.T) {
	suite.Run(t, new(unaryServerInterceptorTestSuite))
}
