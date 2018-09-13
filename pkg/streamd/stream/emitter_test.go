package stream_manager

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type emitterTestSuite struct {
	suite.Suite
	emitter0 Emitter
	mock0    *mock.Mock
	mock1    *mock.Mock
	wg0      *sync.WaitGroup
	wg1      *sync.WaitGroup
}

func (self *emitterTestSuite) SetupTest() {
	self.emitter0 = NewEmitter()
	self.mock0 = &mock.Mock{}
	self.mock1 = &mock.Mock{}
	self.wg0 = &sync.WaitGroup{}
	self.wg1 = &sync.WaitGroup{}

	self.mock0.On("func1", Event("test"), 1).Return()
	self.emitter0.Listen(Event("test"), func(event Event, message interface{}) {
		self.mock0.Called(event, message)
		self.wg0.Done()
	})
	self.wg0.Add(1)

	self.mock1.On("func2", Event("test1"), 1).Return()
	self.emitter0.Once(Event("test1"), func(event Event, message interface{}) {
		self.mock1.Called(event, message)
		self.wg1.Done()
	})
	self.wg1.Add(1)
}

func (self *emitterTestSuite) TestListen() {
	self.emitter0.Listen(Event("test"), nil)
}

func (self *emitterTestSuite) TestEmit() {
	self.emitter0.Emit(Event("test"), 1)
	self.wg0.Wait()
	self.mock0.AssertCalled(self.T(), "func1", Event("test"), 1)
	self.mock0.AssertNumberOfCalls(self.T(), "func1", 1)
}

func (self *emitterTestSuite) TestNotEmit() {
	self.emitter0.Emit(Event("notest"), 1)
	self.mock0.AssertNotCalled(self.T(), "func1", Event("notest"), 1)
	self.mock0.AssertNumberOfCalls(self.T(), "func1", 0)
}

func (self *emitterTestSuite) TestOnce() {
	self.emitter0.Emit(Event("test1"), 1)
	self.wg1.Wait()
	self.mock1.AssertCalled(self.T(), "func2", Event("test1"), 1)
	self.mock1.AssertNumberOfCalls(self.T(), "func2", 1)
}

func (self *emitterTestSuite) TestOnceWithEmitedTwice() {
	self.emitter0.Emit(Event("test1"), 1)
	self.emitter0.Emit(Event("test1"), 1)
	self.wg1.Wait()
	self.mock1.AssertNumberOfCalls(self.T(), "func2", 1)
}

func TestEmitterTestSuite(t *testing.T) {
	suite.Run(t, new(emitterTestSuite))
}
