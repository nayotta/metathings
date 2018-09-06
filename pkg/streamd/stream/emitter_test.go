package stream_manager

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type emitterTestSuite struct {
	suite.Suite
	emitter Emitter
	mock    *mock.Mock
	wg      *sync.WaitGroup
}

func (self *emitterTestSuite) SetupTest() {
	self.emitter = NewEmitter()
	self.mock = &mock.Mock{}
	self.wg = &sync.WaitGroup{}

	self.mock.On("func1", Event("test"), 1).Return()
	self.emitter.Listen(Event("test"), func(event Event, message interface{}) {
		self.mock.Called(event, message)
		self.wg.Done()
	})
	self.wg.Add(1)
}

func (self *emitterTestSuite) TestListen() {
	self.emitter.Listen(Event("test"), nil)
}

func (self *emitterTestSuite) TestEmit() {
	self.emitter.Emit(Event("test"), 1)
	self.wg.Wait()
	self.mock.AssertCalled(self.T(), "func1", Event("test"), 1)
	self.mock.AssertNumberOfCalls(self.T(), "func1", 1)
}

func (self *emitterTestSuite) TestNotEmit() {
	self.emitter.Emit(Event("notest"), 1)
	self.mock.AssertNotCalled(self.T(), "func1", Event("notest"), 1)
	self.mock.AssertNumberOfCalls(self.T(), "func1", 0)
}

func TestEmitterTestSuite(t *testing.T) {
	suite.Run(t, new(emitterTestSuite))
}
