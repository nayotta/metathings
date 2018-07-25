package emitter

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type emitterTestSuite struct {
	suite.Suite
	emitter Emitter
}

func (suite *emitterTestSuite) SetupTest() {
	suite.emitter = NewEmitter()
}

func (suite *emitterTestSuite) TestEmitter() {
	suite.emitter.OnEvent(func(evt Event) error {
		suite.Equal("test", evt.Name())
		suite.Equal("test", evt.Data().(string))
		return nil
	})

	suite.emitter.Trigger(NewEvent("test", "test"))
}

func TestEmitterTestSuite(t *testing.T) {
	suite.Run(t, new(emitterTestSuite))
}
